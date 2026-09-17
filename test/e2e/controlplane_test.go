package e2e

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	k8s "k8s.io/client-go/kubernetes"

	"github.com/example/controlplane-operator/pkg/failover"
	"github.com/example/controlplane-operator/pkg/healthcheck"
	"github.com/example/controlplane-operator/pkg/replication"
)

var _ = Describe("ControlPlane E2E", func() {
	var (
		ctx           context.Context
		kubeClient    k8s.Interface
		failoverMed   *failover.Mediator
		replTopology  *replication.Topology
		healthChecker *healthcheck.Checker
	)

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		kubeClient, err = getKubeClient()
		Expect(err).NotTo(HaveOccurred())

		failoverMed = failover.NewMediator(kubeClient)
		replTopology = replication.NewTopology(kubeClient)
		healthChecker = healthcheck.NewChecker(kubeClient)
	})

	Describe("Validacion de infraestructura del control plane", func() {
		It("debe verificar que todos los clusters objetivo esten alcanzables", func() {
			clusters, err := replTopology.GetClusters(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(clusters).ToNot(BeEmpty())

			for _, cluster := range clusters {
				healthy, err := healthChecker.IsClusterHealthy(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(healthy).To(BeTrue(), "Cluster %s debe estar saludable", cluster.Name)
			}
		})

		It("debe validar la configuracion de ArgoCD en cada cluster", func() {
			clusters, err := replTopology.GetClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				argoHealthy, err := healthChecker.CheckArgoCD(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(argoHealthy).To(BeTrue(), "ArgoCD debe estar operativo en %s", cluster.Name)
			}
		})

		It("debe verificar la conectividad con Crossplane", func() {
			clusters, err := replTopology.GetClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				crossplaneReady, err := healthChecker.CheckCrossplane(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(crossplaneReady).To(BeTrue(), "Crossplane debe estar listo en %s", cluster.Name)
			}
		})
	})

	Describe("Gestion de secrets entre regiones", func() {
		It("debe replicar secrets de Vault a todos los clusters", func() {
			clusters, err := replTopology.GetClustersByRegion(ctx, "us-east-1")
			Expect(err).NotTo(HaveOccurred())

			secretPath := "secret/data/controlplane/credentials"
			for _, cluster := range clusters {
				replicated, err := healthChecker.VerifySecretReplicated(ctx, cluster, secretPath)
				Expect(err).NotTo(HaveOccurred())
				Expect(replicated).To(BeTrue(), "Secret debe estar replicado en %s", cluster.Name)
			}
		})

		It("debe rotar secrets automaticamente cada 24 horas", func() {
			lastRotation, err := healthChecker.GetLastSecretRotation(ctx)
			Expect(err).NotTo(HaveOccurred())

			now := time.Now()
			expectedNextRotation := lastRotation.Add(24 * time.Hour)
			Expect(now).To(BeTemporally("<=", expectedNextRotation),
				"Rotacion debe ocurrir cada 24 horas")
		})

		It("debe validar integridad de secrets replicados", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(clusters)).To(BeNumerically(">=", 2), "Se necesitan al menos 2 clusters")

			secretChecksums := make(map[string]string)
			for _, cluster := range clusters {
				checksum, err := healthChecker.ComputeSecretChecksum(ctx, cluster, "secret/data/controlplane/tls")
				Expect(err).NotTo(HaveOccurred())
				secretChecksums[cluster.Name] = checksum
			}

			var checksums []string
			for _, cs := range secretChecksums {
				checksums = append(checksums, cs)
			}
			Expect(checksums).To(HaveLen(1), "Todos los clusters deben tener el mismo checksum")
		})
	})

	Describe("Replicacion de estado del control plane", func() {
		It("debe sincronizar el estado de CRDs entre clusters", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				synced, err := replTopology.VerifyCRDSync(ctx, cluster, "controlplanes.controlplane.example.com")
				Expect(err).NotTo(HaveOccurred())
				Expect(synced).To(BeTrue(), "CRD debe estar sincronizado en %s", cluster.Name)
			}
		})

		It("debe mantener consistencia de recursos entre regiones", func() {
			regions := []string{"us-east-1", "us-west-2", "eu-west-1", "ap-southeast-1"}

			for _, region := range regions {
				clusters, err := replTopology.GetClustersByRegion(ctx, region)
				Expect(err).NotTo(HaveOccurred())
				Expect(clusters).ToNot(BeEmpty(), "Region %s debe tener clusters", region)

				for _, cluster := range clusters {
					consistent, err := replTopology.VerifyResourceConsistency(ctx, cluster)
					Expect(err).NotTo(HaveOccurred())
					Expect(consistent).To(BeTrue(), "Recursos deben ser consistentes en %s", cluster.Name)
				}
			}
		})
	})

	Describe("Failover automatico del control plane", func() {
		It("debe detectar fallo de primario y promover secundario", func() {
			primary, err := failoverMed.GetPrimary(ctx)
			Expect(err).NotTo(HaveOccurred())

			failed := failoverMed.SimulateFailure(ctx, primary)
			Expect(failed).To(BeTrue())

			Eventually(func() bool {
				newPrimary, err := failoverMed.GetPrimary(ctx)
				if err != nil {
					return false
				}
				return newPrimary.Name != primary.Name
			}, 60*time.Second, 5*time.Second).Should(BeTrue(),
				"Debe promover un nuevo primario cuando el actual falla")
		})

		It("debe actualizar Route53 con el nuevo primario", func() {
			newPrimary, err := failoverMed.GetPrimary(ctx)
			Expect(err).NotTo(HaveOccurred())

			route53Updated, err := healthChecker.VerifyRoute53Update(ctx, newPrimary)
			Expect(err).NotTo(HaveOccurred())
			Expect(route53Updated).To(BeTrue(), "Route53 debe apuntar al nuevo primario")
		})

		It("debe notificar a todos los clusters del cambio de primario", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			newPrimary, err := failoverMed.GetPrimary(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				notified, err := healthChecker.VerifyFailoverNotification(ctx, cluster, newPrimary)
				Expect(err).NotTo(HaveOccurred())
				Expect(notified).To(BeTrue(), "Cluster %s debe estar notificado", cluster.Name)
			}
		})
	})
})