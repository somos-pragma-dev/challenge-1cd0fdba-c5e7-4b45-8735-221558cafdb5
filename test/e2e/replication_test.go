package e2e

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/example/controlplane-operator/pkg/replication"
)

var _ = Describe("Replication Topology E2E", func() {
	var (
		ctx          context.Context
		replTopology *replication.Topology
	)

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		kubeClient, err = getKubeClient()
		Expect(err).NotTo(HaveOccurred())

		replTopology = replication.NewTopology(kubeClient)
	})

	Describe("Topologia de replicacion entre clusters", func() {
		It("debe obtener todos los clusters registrados", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(clusters).ToNot(BeEmpty(), "Debe existir al menos un cluster registrado")
		})

		It("debe organizar clusters por region", func() {
			regions := []string{"us-east-1", "us-west-2", "eu-west-1", "ap-southeast-1"}

			for _, region := range regions {
				clusters, err := replTopology.GetClustersByRegion(ctx, region)
				Expect(err).NotTo(HaveOccurred())
				Expect(clusters).ToNot(BeEmpty(),
					"Region %s debe tener clusters registrados", region)
			}
		})

		It("debe identificar cluster primario y secundarios", func() {
			primary, err := replTopology.GetPrimaryCluster(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(primary).ToNot(BeNil())
			Expect(primary.IsPrimary).To(BeTrue())

			secondaries, err := replTopology.GetSecondaryClusters(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(secondaries).ToNot(BeEmpty())

			for _, sec := range secondaries {
				Expect(sec.IsPrimary).To(BeFalse())
			}
		})
	})

	Describe("Replicacion de estado entre regiones", func() {
		It("debe establecer conexion de replicacion entre clusters", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(clusters).To(HaveLen(12),
				"Debe haber 12 clusters en la topologia")

			for i := 0; i < len(clusters)-1; i++ {
				connected, err := replTopology.TestReplicationLink(ctx, clusters[i], clusters[i+1])
				Expect(err).NotTo(HaveOccurred())
				Expect(connected).To(BeTrue(),
					"Debe haber enlace de replicacion entre %s y %s",
					clusters[i].Name, clusters[i+1].Name)
			}
		})

		It("debe sincronizar CustomResources entre clusters", func() {
			crName := "test-controlplane"
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				synced, err := replTopology.VerifyCRSync(ctx, cluster, crName)
				Expect(err).NotTo(HaveOccurred())
				Expect(synced).To(BeTrue(),
					"CR %s debe estar sincronizado en %s", crName, cluster.Name)
			}
		})

		It("debe mantener consistencia eventual con maximo 5 segundos de divergencia", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(clusters).To(HaveLen(12))

			for _, cluster := range clusters {
				divergence, err := replTopology.MeasureStateDivergence(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(divergence.Seconds()).To(BeNumerically("<", 5),
					"Divergencia maxima de 5 segundos en %s", cluster.Name)
			}
		})

		It("debe resolver conflictos de escritura con last-write-wins", func() {
			cluster, err := replTopology.GetPrimaryCluster(ctx)
			Expect(err).NotTo(HaveOccurred())

			conflictResolved, err := replTopology.ResolveWriteConflict(ctx, cluster, "last-write-wins")
			Expect(err).NotTo(HaveOccurred())
			Expect(conflictResolved).To(BeTrue())
		})
	})

	Describe("Consistencia de datos cross-region", func() {
		It("debe verificar que el estado de etcd es consistente", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				consistent, err := replTopology.VerifyEtcdConsistency(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(consistent).To(BeTrue(),
					"Estado de etcd debe ser consistente en %s", cluster.Name)
			}
		})

		It("debe validar quórum distribuido", func() {
			quorumStatus, err := replTopology.CheckDistributedQuorum(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(quorumStatus.HasQuorum).To(BeTrue())
			Expect(quorumStatus.RequiredVotes).To(BeNumerically(">", len(quorumStatus.Voters)/2))
		})

		It("debe detectar y resolver divergencias de estado", func() {
			divergences, err := replTopology.DetectStateDivergences(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, div := range divergences {
				resolved, err := replTopology.ResolveDivergence(ctx, div)
				Expect(err).NotTo(HaveOccurred())
				Expect(resolved).To(BeTrue(), "Divergencia %s debe resolverse", div.ID)
			}
		})

		It("debe replicar configuraciones de RBAC entre clusters", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				rbacSynced, err := replTopology.VerifyRBACSync(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(rbacSynced).To(BeTrue(),
					"Configuracion RBAC debe estar sincronizada en %s", cluster.Name)
			}
		})
	})

	Describe("Latencia de replicacion", func() {
		It("debe medir latencia de replicacion entre regiones", func() {
			regions := []string{"us-east-1", "us-west-2", "eu-west-1", "ap-southeast-1"}

			for i := 0; i < len(regions)-1; i++ {
				latency, err := replTopology.MeasureReplicationLatency(ctx, regions[i], regions[i+1])
				Expect(err).NotTo(HaveOccurred())
				Expect(latency.Milliseconds()).To(BeNumerically("<", 500),
					"Latencia entre %s y %s debe ser menor a 500ms",
					regions[i], regions[i+1])
			}
		})

		It("debe cumplir con SLA de replicacion", func() {
			slaMet, err := replTopology.VerifySLA(ctx, 5*60) // 5 minutes
			Expect(err).NotTo(HaveOccurred())
			Expect(slaMet).To(BeTrue(), "Debe cumplir SLA de replicacion de 5 minutos")
		})
	})
})