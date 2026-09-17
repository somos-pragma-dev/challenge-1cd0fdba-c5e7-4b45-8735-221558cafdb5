# Diseño de control plane para clusters Kubernetes distribuidos

Debes diseñar una topología de control plane que gestione 12 clusters Kubernetes distribuidos en 4 regiones AWS. El sistema usa ArgoCD para GitOps, Crossplane para provisionar infraestructura declarativa, y una lógica custom de failover. Debes justificar la elección de herramientas, diseñar la replicación de estado, el modelo de RBAC federado, la estrategia de secrets, y la coreografía de disaster recovery.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | control plane multi-cluster kubernetes con GitOps y disaster recovery cross-region |
| **Nivel** | master-l2 |
| **Tipo** | theoretical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Exploración del dominio y restricciones

**Objetivo:** Identificar y documentar las restricciones y ambigüedades del sistema actual.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identifica las restricciones clave del sistema, como el throughput de 4 000 transacciones por segundo con un SLA del 99.9% y la trazabilidad de 5 años.
- Documenta las ambigüedades encontradas y cómo podrían impactar el diseño del control plane.

**Entregable:** Documento que describe las restricciones y ambigüedades del sistema.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo distinguir entre restricciones relevantes y triviales.

</details>

### Fase 2: Evaluación de herramientas y trade-offs

**Objetivo:** Evaluar y justificar la elección de ArgoCD vs Flux vs Rancher Fleet para el control plane.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Evalúa las ventajas y desventajas de usar ArgoCD, Flux y Rancher Fleet en el contexto de tu control plane.
- Justifica tu elección basándote en las necesidades del sistema y las restricciones identificadas.

**Entregable:** Registro de decisiones que documenta la elección de la herramienta, incluyendo contexto, fuerzas, opciones con pros/contras, decisión y consecuencias.

<details>
<summary>Pistas de conocimiento</summary>

- Considera la facilidad de uso, la integración con otras herramientas, la comunidad y el soporte.

</details>

### Fase 3: Diseño de la topología de replicación de estado

**Objetivo:** Diseñar la topología de replicación de estado entre los clusters.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Diseña la topología de replicación de estado que garantice la consistencia y disponibilidad del sistema.
- Considera los posibles modos de falla y cómo evitar el split-brain durante los failovers.

**Entregable:** Diagrama de la topología de replicación de estado y documento que describe el diseño.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el uso de herramientas como etcd o Consul para la gestión de estado.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es un control plane en el contexto de Kubernetes?
- **paraQueSirve**: ¿Para qué sirve ArgoCD en el contexto de GitOps?
- **comoSeUsa**: ¿Cómo se usa Crossplane para provisionar infraestructura declarativa?
- **erroresComunes**: ¿Cuáles son los errores comunes al diseñar un sistema de disaster recovery cross-region?
- **queDecisionesImplica**: ¿Qué decisiones implica la elección de una herramienta para el control plane?

## Criterios de Evaluacion

- Identificar y documentar las restricciones y ambigüedades del sistema.
- Evaluar y justificar la elección de ArgoCD vs Flux vs Rancher Fleet.
- Diseñar la topología de replicación de estado entre los clusters.
- Justificar la elección de herramientas y estrategias para el disaster recovery.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT). Devuelve un ZIP con el proyecto. Sirve si no tenes un agente en el IDE.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
el comando de build o arranque canonico del stack elegido
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
