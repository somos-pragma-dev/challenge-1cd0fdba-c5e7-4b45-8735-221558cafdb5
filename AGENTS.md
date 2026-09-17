# AGENTS.md

Instrucciones para el agente de IA que abra este repositorio (Claude Code, Cursor, Codex, Copilot, Gemini). Se cargan solas: no hay que pegar nada en ningun chat.

## Que es este repositorio

Es el codigo base de un reto de aprendizaje de Pragma: **Diseño de control plane para clusters Kubernetes distribuidos**.

| | |
|---|---|
| Tema | control plane multi-cluster kubernetes con GitOps y disaster recovery cross-region |
| Nivel | master-l2 |
| Chapter | Generico |
| Especialidad | Inferido del contexto |
| Stack | Go / Kubernetes Operator SDK v1.34 |
| Patron arquitectonico | control plane basado en operadores Kubernetes con patrón de mediador para coreografía de disaster recovery |
| Tiempo estimado | 8 horas |

## Tu tarea

Dejar este proyecto en estado **verificable**: que el comando de verificacion corra sin errores. Escribi los archivos en disco, en este repositorio. No generes ZIPs ni archivos adjuntos.

En orden:

1. Corre `el comando de build o arranque canonico del stack elegido` y mira que falla.
2. Completa lo que falte de la lista de abajo: manifiesto de dependencias, punto de entrada, capa de interfaz y las capas del patron declarado.
3. Arregla SOLO los errores que impiden compilar o arrancar.
4. Volve a correr `el comando de build o arranque canonico del stack elegido` hasta que pase.
5. Pará ahí.

## Regla dura: las fases son trabajo del humano

**PROHIBIDO implementar los entregables de las fases.** El valor del reto esta en que la persona los resuelva. Tu trabajo es que tenga un proyecto que arranca; el hueco pedagogico se queda como esta.

No resuelvas nada de esto:

- **Fase 1 — Exploración del dominio y restricciones**: Documento que describe las restricciones y ambigüedades del sistema.
- **Fase 2 — Evaluación de herramientas y trade-offs**: Registro de decisiones que documenta la elección de la herramienta, incluyendo contexto, fuerzas, opciones con pros/contras, decisión y consecuencias.
- **Fase 3 — Diseño de la topología de replicación de estado**: Diagrama de la topología de replicación de estado y documento que describe el diseño.

Distincion operativa:

- **Arreglar** (si): import faltante, tipo que no existe, dependencia sin declarar, error de sintaxis, archivo referenciado que no existe.
- **No tocar** (no): logica de negocio incompleta, validaciones ausentes, secretos hardcodeados, APIs deprecadas que funcionan, concurrencia insegura, patrones mejorables. Eso es lo que la persona tiene que encontrar.

## Lo que falta y tenes que completar

### 1. Boilerplate del stack (1)

Sin esto el proyecto no compila ni arranca. **Es tu trabajo crearlo**, y no toca nada de lo pedagogico: es andamiaje del stack.

- [ ] **Capa de interfaz (controller/handler)** — Sin una capa de interfaz explicita, no hay forma de invocar la logica de negocio desde afuera del proceso.

### 2. Archivos que la arquitectura declara (3 de 20)

La propuesta arquitectonica del reto los lista y no llegaron al repo. Crealos con implementacion real, respetando la capa en la que viven:

- [ ] `controllers/controlplane_controller.go`
- [ ] `pkg/aws/config.go`
- [ ] `pkg/aws/healthcheck_handler.go`

### Presentes (17)

- `go.mod`
- `main.go`
- `api/v1alpha1/controlplane_types.go`
- `pkg/argocd/client.go`
- `pkg/crossplane/provider.go`
- `pkg/vault/secrets.go`
- `pkg/failover/mediator.go`
- `pkg/healthcheck/checker.go`
- `pkg/rbac/federator.go`
- `pkg/replication/topology.go`
- `pkg/aws/dto.go`
- `config/crd/bases/controlplane.example.com_controlplanes.yaml`
- `config/rbac/role.yaml`
- `config/samples/controlplane_v1alpha1_controlplane.yaml`
- `test/e2e/controlplane_test.go`
- `test/e2e/failover_test.go`
- `test/e2e/replication_test.go`

### Capas del patron declarado

Cada una tiene que existir como directorio real con al menos un archivo. Codigo plano en la raiz no satisface el patron.

- `api/v1alpha1`
- `controllers`
- `pkg/argocd`
- `pkg/crossplane`
- `pkg/vault`
- `pkg/failover`
- `pkg/healthcheck`
- `pkg/rbac`
- `pkg/replication`
- `config/crd`
- `config/rbac`
- `config/samples`
- `hack`
- `test/e2e`

## Verificacion

```bash
el comando de build o arranque canonico del stack elegido
```

Ese comando pasando es la definicion de "terminado" para vos.

## Convenciones que tenes que respetar

- Un solo ecosistema: no declares librerias de otro lenguaje ni mezcles gestores de paquetes.
- Toda libreria que uses tiene que estar declarada en el manifiesto de dependencias.
- Todo import declarado tiene que usarse; todo tipo usado tiene que existir o venir de una dependencia declarada.
- El patron es **control plane basado en operadores Kubernetes con patrón de mediador para coreografía de disaster recovery**: los contratos (interfaces, puertos) los define la capa interna y los implementa la externa, nunca al revés.
- Los archivos que crees llevan implementacion real, no stubs: sin `TODO`, sin cuerpos vacios, sin `// getters y setters`.

## Contexto del candidato

Sirve para calibrar el nivel del codigo, no para resolver las fases.

- Brecha que el reto ataca: Arquitectura master-l2 de un control plane que gestiona 12 clusters Kubernetes distribuidos en 4 regiones AWS. Usa ArgoCD para GitOps con sync policies por entorno, Crossplane para provisionar infraestructura declarativa (RDS, S3, IAM roles) desde el repo de Git, y una lógica custom de failover que promueve un cluster standby cuando el primary de una región falla su health check por 3 minutos consecutivos. El developer master-l2 debe diseñar la topología de replicación de estado entre clusters, justificar la elección de ArgoCD vs Flux vs Rancher Fleet, el modelo de RBAC federado, la estrategia de secrets con Vault + external-secrets, y la coreografía del disaster recovery incluyendo cómo evita el split-brain durante failovers.

---

*Generado por Challenge Generator — Pragma. `README.md` tiene el enunciado completo del reto para la persona. `PROMPT_MEJORA.md` es la variante para pegar en un chat, si se prefiere ese flujo.*
