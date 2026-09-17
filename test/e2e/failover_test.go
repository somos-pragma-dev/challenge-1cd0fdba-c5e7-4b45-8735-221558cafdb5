package e2e

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/example/controlplane-operator/pkg/failover"
	"github.com/example/controlplane-operator/pkg/replication"
)

var _ = Describe("Failover y Disaster Recovery E2E", func() {
	var (
		ctx          context.Context
		failoverMed  *failover.Mediator
		replTopology *replication.Topology
	)

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		kubeClient, err = getKubeClient()
		Expect(err).NotTo(HaveOccurred())

		failoverMed = failover.NewMediator(kubeClient)
		replTopology = replication.NewTopology(kubeClient)
	})

	Describe("Logica de failover", func() {
		Context("cuando el cluster primario falla", func() {
			It("debe iniciar procedimiento de failover automaticamente", func() {
				initialPrimary, err := failoverMed.GetPrimary(ctx)
				Expect(err).NotTo(HaveOccurred())

				failoverMed.SimulateNetworkPartition(ctx, initialPrimary)

				Eventually(func() string {
					newPrimary, _ := failoverMed.GetPrimary(ctx)
					if newPrimary != nil {
						return newPrimary.Name
					}
					return ""
				}, 120*time.Second, 10*time.Second).ShouldNot(Equal(initialPrimary.Name),
					"Debe haber un nuevo primario diferente al original")
			})

			It("debe seleccionar el candidato con mayor quorum", func() {
				candidates, err := failoverMed.GetCandidates(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(candidates).ToNot(BeEmpty())

				bestCandidate := failoverMed.SelectByQuorum(candidates)
				Expect(bestCandidate).ToNot(BeNil())
				Expect(bestCandidate.QuorumVotes).To(BeNumerically(">=", len(candidates)/2+1))
			})

			It("debe actualizar el estado de failover en etcd", func() {
				failoverState, err := failoverMed.GetFailoverState(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(failoverState.Status).To(Equal("in-progress"))

				Eventually(func() string {
					state, _ := failoverMed.GetFailoverState(ctx)
					if state != nil {
						return state.Status
					}
					return ""
				}, 90*time.Second, 5*time.Second).Should(Equal("completed"))
			})
		})

		Context("cuando ocurre split-brain", func() {
			It("debe detectar conflicto de liderazgo", func() {
				twoPrimaries, err := failoverMed.DetectSplitBrain(ctx)
				Expect(err).NotTo(HaveOccurred())
				if twoPrimaries {
					conflictResolved, err := failoverMed.ResolveSplitBrain(ctx)
					Expect(err).NotTo(HaveOccurred())
					Expect(conflictResolved).To(BeTrue())
				}
			})

			It("debe resolver split-brain con timestamp-based resolution", func() {
				resolution, err := failoverMed.ResolveWithTimestamp(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(resolution.Primary.Name).ToNot(BeEmpty())

				singlePrimary, err := failoverMed.VerifySinglePrimary(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(singlePrimary).To(BeTrue(), "Debe existir un unico primario tras resolucion")
			})

			It("debe prevenir split-brain con lease de etcd", func() {
				leaseAcquired, err := failoverMed.TryAcquireLease(ctx, "controlplane-leader")
				Expect(err).NotTo(HaveOccurred())
				Expect(leaseAcquired).To(BeTrue())

				expireErr := failoverMed.ExpireLease(ctx, "controlplane-leader")
				Expect(expireErr).ToNot(HaveOccurred())

				secondLease, err := failoverMed.TryAcquireLease(ctx, "controlplane-leader")
				Expect(err).NotTo(HaveOccurred())
				Expect(secondLease).To(BeTrue(), "Segundo lease debe poder adquirirse tras expiration")
			})
		})
	})

	Describe("Coreografia de disaster recovery", func() {
		It("debe ejecutar secuencia de recovery en orden correcto", func() {
			recoverySteps := []string{
				"detect_failure",
				"promote_secondary",
				"update_dns",
				"sync_state",
				"verify_health",
				"notify_clusters",
			}

			executed, err := failoverMed.ExecuteRecoverySequence(ctx, recoverySteps)
			Expect(err).NotTo(HaveOccurred())
			Expect(executed).To(HaveLen(len(recoverySteps)))

			for i, step := range recoverySteps {
				Expect(executed[i].Name).To(Equal(step))
				Expect(executed[i].Status).To(Equal("success"))
			}
		})

		It("debe rollbackear si verificacion falla", func() {
			originalPrimary, err := failoverMed.GetPrimary(ctx)
			Expect(err).NotTo(HaveOccurred())

			failoverMed.SimulateHealthCheckFailure(ctx)

			rolledBack, err := failoverMed.RollbackOnFailure(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(rolledBack).To(BeTrue())

			currentPrimary, err := failoverMed.GetPrimary(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(currentPrimary.Name).To(Equal(originalPrimary.Name))
		})

		It("debe mantener RPO menor a 30 segundos", func() {
			rpo, err := failoverMed.MeasureRPO(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(rpo.Seconds()).To(BeNumerically("<", 30),
				"RPO debe ser menor a 30 segundos")
		})

		It("debe registrar eventos de recovery para auditoria", func() {
			events, err := failoverMed.GetRecoveryEvents(ctx)
			Expect(err).NotTo(HaveOccurred())

			if len(events) > 0 {
				latestEvent := events[len(events)-1]
				Expect(latestEvent.Timestamp).ToNot(BeZero())
				Expect(latestEvent.Type).To(BeElementOf("failover", "rollback", "manual_switch"))
			}
		})
	})

	Describe("Validacion post-failover", func() {
		It("debe verificar que todos los servicios estan operativos", func() {
			newPrimary, err := failoverMed.GetPrimary(ctx)
			Expect(err).NotTo(HaveOccurred())

			servicesHealthy, err := failoverMed.VerifyServicesPostFailover(ctx, newPrimary)
			Expect(err).NotTo(HaveOccurred())
			Expect(servicesHealthy).To(BeTrue(),
				"Todos los servicios deben estar operativos tras failover")
		})

		It("debe validar consistencia de datos post-failover", func() {
			clusters, err := replTopology.GetAllClusters(ctx)
			Expect(err).NotTo(HaveOccurred())

			for _, cluster := range clusters {
				consistent, err := replTopology.ValidateDataConsistency(ctx, cluster)
				Expect(err).NotTo(HaveOccurred())
				Expect(consistent).To(BeTrue(), "Datos deben ser consistentes en %s", cluster.Name)
			}
		})
	})
})