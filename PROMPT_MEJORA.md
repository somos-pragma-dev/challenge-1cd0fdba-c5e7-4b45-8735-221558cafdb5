# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener un ZIP con el proyecto completo y arrancable.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Lo que le falta a este proyecto

Esto NO lo tenes que adivinar: salio de comparar el proyecto contra la arquitectura declarada del reto y de un analisis estatico del codigo. Completalo TODO.

### Boilerplate del stack que falta

Sin esto no compila ni arranca. Es andamiaje, no toca nada de lo pedagogico:

- **Capa de interfaz (controller/handler)** — Sin una capa de interfaz explicita, no hay forma de invocar la logica de negocio desde afuera del proceso.

### Archivos que la arquitectura del reto declara y no estan

Creálos con implementacion real, en la capa que les corresponde:

- `controllers/controlplane_controller.go`
- `pkg/aws/config.go`
- `pkg/aws/healthcheck_handler.go`

## Como saber que terminaste

```bash
el comando de build o arranque canonico del stack elegido
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
Arquitectura master-l2 de un control plane que gestiona 12 clusters Kubernetes distribuidos en 4 regiones AWS. Usa ArgoCD para GitOps con sync policies por entorno, Crossplane para provisionar infraestructura declarativa (RDS, S3, IAM roles) desde el repo de Git, y una lógica custom de failover que promueve un cluster standby cuando el primary de una región falla su health check por 3 minutos consecutivos. El developer master-l2 debe diseñar la topología de replicación de estado entre clusters, justificar la elección de ArgoCD vs Flux vs Rancher Fleet, el modelo de RBAC federado, la estrategia de secrets con Vault + external-secrets, y la coreografía del disaster recovery incluyendo cómo evita el split-brain durante failovers.

### Reto
- Tema: control plane multi-cluster kubernetes con GitOps y disaster recovery cross-region
- Seniority: master-l2
- Tipo: theoretical
- Título: Diseño de control plane para clusters Kubernetes distribuidos
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Exploración del dominio y restricciones — objetivo: Identificar y documentar las restricciones y ambigüedades del sistema actual. — entregable (NO resolver): Documento que describe las restricciones y ambigüedades del sistema.
- Fase 2: Evaluación de herramientas y trade-offs — objetivo: Evaluar y justificar la elección de ArgoCD vs Flux vs Rancher Fleet para el control plane. — entregable (NO resolver): Registro de decisiones que documenta la elección de la herramienta, incluyendo contexto, fuerzas, opciones con pros/contras, decisión y consecuencias.
- Fase 3: Diseño de la topología de replicación de estado — objetivo: Diseñar la topología de replicación de estado entre los clusters. — entregable (NO resolver): Diagrama de la topología de replicación de estado y documento que describe el diseño.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/example/controlplane-operator

go 1.22

require (
	github.com/argoproj/argo-cd/v2 v2.9.0
	github.com/aws/aws-sdk-go v1.44.0
	github.com/crossplane/crossplane-runtime v1.13.0
	github.com/hashicorp/vault/api v1.11.0
	github.com/onsi/ginkgo/v2 v2.13.0
	github.com/onsi/gomega v1.27.0
	k8s.io/apimachinery v0.28.0
	sigs.k8s.io/controller-runtime v0.17.0
	sigs.k8s.io/controller-tools/cmd/controller-gen v0.13.0
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.27 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.5 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.0 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/Microsoft/hcsshim v0.9.3 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20211112122917-428f8e5eb3dd // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bfcad3ed1 // indirect
	github.com/argoproj/argo-cd/v2 v2.9.0 // indirect
	github.com/argoproj/gitops-engine v0.7.1-0.20230704120346-27f9c4b044c5 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/crossplane/crossplane-runtime v1.13.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/docker/cli v20.10.17+incompatible // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.17+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.4 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/zapr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-containerregistry v0.12.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/shlex v0.0.0-20181106134648-c34317bd91bf // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d907240c4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/vault/api v1.11.0 // indirect
	github.com/hashicorp/vault/sdk v0.10.0 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/term v0.0.0-20221205130635-1ae3c6bda3e4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/morikuni/aec v1.0.0-20220314085400-cf5945bf2d2e // indirect
	github.com/onsi/ginkgo/v2 v2.13.0 // indirect
	github.com/onsi/gomega v1.27.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc2 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/russross/blackfriday v1.6.0 // indirect
	github.com/sanity-io/litter v1.5.5 // indirect
	github.com/sergi/go-diff v1.3.1-0.20230801104853-5e72ba4da035 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/theupdateframework/notary v0.7.0 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/vbatts/tar-split v0.11.2 // indirect
	github.com/xanzy/go-gitlab v0.73.1 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/v3 v3.5.7 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.14.0 // indirect
	go.opentelemetry.io/otel/metric v1.14.0 // indirect
	go.opentelemetry.io/otel/sdk v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/api v0.110.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db087b2931 // indirect
	google.golang.org/grpc v1.51.1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.28.0 // indirect
	k8s.io/apiextensions-apiserver v0.28.0 // indirect
	k8s.io/apimachinery v0.28.0 // indirect
	k8s.io/apiserver v0.28.0 // indirect
	k8s.io/cli-runtime v0.28.0 // indirect
	k8s.io/client-go v0.28.0 // indirect
	k8s.io/cluster-bootstrap v0.28.0 // indirect
	k8s.io/component-base v0.28.0 // indirect
	k8s.io/component-helpers v0.28.0 // indirect
	k8s.io/klog/v2 v2.90.1 // indirect
	k8s.io/kube-aggregator v0.28.0 // indirect
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280 // indirect
	k8s.io/kubectl v0.28.0 // indirect
	k8s.io/kubelet v0.28.0 // indirect
	k8s.io/legacy-cloud-providers v0.28.0 // indirect
	k8s.io/utils v0.0.0-20221108220202-28ad6200f2b3 // indirect
	oras.land/oras-go v1.2.0 // indirect
	rigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/controller-runtime v0.17.0 // indirect
	sigs.k8s.io/controller-tools v0.13.0 // indirect
	sigs.k8s.io/kustomize/api v0.13.2-0.20221102041619-24115becce9a // indirect
	sigs.k8s.io/kustomize/kyaml v0.13.2-0.20221102041619-24115becce9a // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.3.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// === ARCHIVO: main.go ===
package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	controlplanev1alpha1 "github.com/example/controlplane-operator/api/v1alpha1"
	"github.com/example/controlplane-operator/controllers"
	"github.com/example/controlplane-operator/pkg/argocd"
	"github.com/example/controlplane-operator/pkg/aws"
	"github.com/example/controlplane-operator/pkg/crossplane"
	"github.com/example/controlplane-operator/pkg/failover"
	"github.com/example/controlplane-operator/pkg/healthcheck"
	"github.com/example/controlplane-operator/pkg/rbac"
	"github.com/example/controlplane-operator/pkg/replication"
	"github.com/example/controlplane-operator/pkg/vault"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = controlplanev1alpha1.AddToScheme(scheme)
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	var clusterConfigPath string

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&clusterConfigPath, "cluster-config", "/etc/controlplane/clusters.yaml",
		"Path to the cluster topology configuration file.")

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "controlplane-operator-lock",
	})
	if err != nil {
		setupLog.Error(err, "Unable to start manager")
		os.Exit(1)
	}

	clusterConfig, err := loadClusterConfiguration(clusterConfigPath)
	if err != nil {
		setupLog.Error(err, "Failed to load cluster configuration")
		os.Exit(1)
	}

	argocdClient, err := argocd.NewClient(clusterConfig.ArgoCD)
	if err != nil {
		setupLog.Error(err, "Unable to initialize ArgoCD client")
		os.Exit(1)
	}

	crossplaneProvider, err := crossplane.NewProvider(clusterConfig.Crossplane)
	if err != nil {
		setupLog.Error(err, "Unable to initialize Crossplane provider")
		os.Exit(1)
	}

	vaultClient, err := vault.NewClient(clusterConfig.Vault)
	if err != nil {
		setupLog.Error(err, "Unable to initialize Vault client")
		os.Exit(1)
	}

	awsConfig, err := aws.NewConfig(clusterConfig.AWS)
	if err != nil {
		setupLog.Error(err, "Unable to initialize AWS configuration")
		os.Exit(1)
	}

	topology, err := replication.NewTopology(clusterConfig.Clusters)
	if err != nil {
		setupLog.Error(err, "Unable to initialize replication topology")
		os.Exit(1)
	}

	failoverMediator := failover.NewMediator(topology, argocdClient, vaultClient)

	healthChecker := healthcheck.NewChecker(argocdClient, awsConfig)

	rbacFederator, err := rbac.NewFederator(mgr.GetClient(), clusterConfig.Clusters)
	if err != nil {
		setupLog.Error(err, "Unable to initialize RBAC federator")
		os.Exit(1)
	}

	if err = (&controllers.ControlPlaneReconciler{
		Client:            mgr.GetClient(),
		Scheme:           mgr.GetScheme(),
		ArgoCD:            argocdClient,
		Crossplane:       crossplaneProvider,
		Vault:             vaultClient,
		AWSConfig:         awsConfig,
		Topology:          topology,
		FailoverMediator:  failoverMediator,
		HealthChecker:     healthChecker,
		RbacFederator:     rbacFederator,
		ClusterConfig:     clusterConfig,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller")
		os.Exit(1)
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "Unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "Unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("Starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "Problem running manager")
		os.Exit(1)
	}
}

type ClusterConfiguration struct {
	Clusters   []ClusterInfo `json:"clusters"`
	ArgoCD     ArgoCDConfig  `json:"argocd"`
	Crossplane CrossplaneConfig `json:"crossplane"`
	Vault      VaultConfig   `json:"vault"`
	AWS        AWSConfig     `json:"aws"`
}

type ClusterInfo struct {
	Name     string `json:"name"`
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
	Priority int    `json:"priority"`
}

type ArgoCDConfig struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CrossplaneConfig struct {
	Provider string `json:"provider"`
	Region   string `json:"region"`
}

type VaultConfig struct {
	Address  string `json:"address"`
	Token    string `json:"token"`
	Path     string `json:"path"`
}

type AWSConfig struct {
	Region    string `json:"region"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func loadClusterConfiguration(configPath string) (*ClusterConfiguration, error) {
	return &ClusterConfiguration{
		Clusters: []ClusterInfo{
			{Name: "primary-us-east-1", Region: "us-east-1", Endpoint: "https://kube.us-east-1.example.com", Priority: 1},
			{Name: "primary-us-west-2", Region: "us-west-2", Endpoint: "https://kube.us-west-2.example.com", Priority: 2},
			{Name: "primary-eu-west-1", Region: "eu-west-1", Endpoint: "https://kube.eu-west-1.example.com", Priority: 3},
			{Name: "primary-ap-southeast-1", Region: "ap-southeast-1", Endpoint: "https://kube.ap-southeast-1.example.com", Priority: 4},
		},
		ArgoCD: ArgoCDConfig{
			Server:   "argocd.example.com",
			Username: "admin",
			Password: "changeme",
		},
		Crossplane: CrossplaneConfig{
			Provider: "aws",
			Region:   "us-east-1",
		},
		Vault: VaultConfig{
			Address: "https://vault.example.com",
			Token:   "root",
			Path:    "secret/data/controlplane",
		},
		AWS: AWSConfig{
			Region:    "us-east-1",
			AccessKey: "",
			SecretKey: "",
		},
	}, nil
}


// === ARCHIVO: api/v1alpha1/controlplane_types.go ===
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ControlPlaneSpec define la especificación del control plane multi-cluster
type ControlPlaneSpec struct {
	// Clusters contiene la configuración de todos los clusters gestionados
	// +kubebuilder:validation:MinItems=1
	Clusters []ClusterSpec `json:"clusters"`

	// ArgoCD define la configuración de ArgoCD para GitOps
	// +optional
	ArgoCD *ArgoCDSpec `json:"argocd,omitempty"`

	// Crossplane define la configuración de Crossplane para infraestructura
	// +optional
	Crossplane *CrossplaneSpec `json:"crossplane,omitempty"`

	// Vault define la configuración de Vault para secrets
	// +optional
	Vault *VaultSpec `json:"vault,omitempty"`

	// AWS define la configuración de AWS para salud de regiones
	// +optional
	AWS *AWSSpec `json:"aws,omitempty"`

	// Replication define la estrategia de replicación de estado
	// +optional
	Replication *ReplicationSpec `json:"replication,omitempty"`

	// Failover define la configuración de disaster recovery
	// +optional
	Failover *FailoverSpec `json:"failover,omitempty"`
}

// ClusterSpec define la especificación de un cluster individual
type ClusterSpec struct {
	// Name es el nombre identificador del cluster
	// +kubebuilder:validation:Pattern=^[a-z0-9-]+$
	Name string `json:"name"`

	// Region es la región AWS donde reside el cluster
	// +kubebuilder:validation:Pattern=^[a-z]{2}-[a-z]+-[0-9]+$
	Region string `json:"region"`

	// Endpoint es la URL del API server del cluster
	// +kubebuilder:validation:Format=uri
	Endpoint string `json:"endpoint"`

	// IsMain indica si este cluster es el principal (primary)
	// +optional
	// +kubebuilder:default=false
	IsMain bool `json:"isMain,omitempty"`

	// IsDrTarget indica si este cluster es objetivo de disaster recovery
	// +optional
	// +kubebuilder:default=false
	IsDrTarget bool `json:"isDrTarget,omitempty"`

	// Priority define la prioridad del cluster para failover
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Priority int `json:"priority,omitempty"`

	// Labels contiene etiquetas adicionales para el cluster
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// ArgoCDSpec define la configuración de ArgoCD
type ArgoCDSpec struct {
	// Enabled indica si ArgoCD está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Version es la versión de ArgoCD a desplegar
	// +optional
	// +kubebuilder:default="v2.9.0"
	Version string `json:"version,omitempty"`

	// Namespace es el namespace donde ArgoCD será instalado
	// +optional
	// +kubebuilder:default="argocd"
	Namespace string `json:"namespace,omitempty"`

	// RepoURL es la URL del repositorio Git con las definiciones de aplicaciones
	// +kubebuilder:validation:Format=uri
	RepoURL string `json:"repoURL"`

	// RepoBranch es la rama del repositorio a usar
	// +optional
	// +kubebuilder:default="main"
	RepoBranch string `json:"repoBranch,omitempty"`

	// SyncPolicy define la política de sincronización automática
	// +optional
	SyncPolicy *ArgoCDSyncPolicy `json:"syncPolicy,omitempty"`

	// Clusters es la lista de clusters registrados en ArgoCD
	// +optional
	Clusters []ArgoCDClusterRef `json:"clusters,omitempty"`
}

// ArgoCDSyncPolicy define la política de sincronización de ArgoCD
type ArgoCDSyncPolicy struct {
	// AutoSync habilita la sincronización automática
	// +optional
	// +kubebuilder:default=true
	AutoSync bool `json:"autoSync,omitempty"`

	// RetryConfiguration contiene la configuración de reintentos
	// +optional
	RetryConfiguration *ArgoCDRetryConfig `json:"retryConfiguration,omitempty"`

	// PrunePropagationPolicy define cómo se propagan las eliminaciones
	// +optional
	// +kubebuilder:validation:Enum=foreground;background;orphan
	// +kubebuilder:default="foreground"
	PrunePropagationPolicy string `json:"prunePropagationPolicy,omitempty"`

	// SelfHeal habilita la auto-recuperación
	// +optional
	// +kubebuilder:default=true
	SelfHeal bool `json:"selfHeal,omitempty"`
}

// ArgoCDRetryConfig define la configuración de reintentos
type ArgoCDRetryConfig struct {
	// Limit es el número máximo de reintentos
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=3
	Limit int `json:"limit,omitempty"`

	// Backoff define la estrategia de backoff
	// +optional
	Backoff *ArgoCDRetryBackoff `json:"backoff,omitempty"`
}

// ArgoCDRetryBackoff define el backoff para reintentos
type ArgoCDRetryBackoff struct {
	// Duration es la duración inicial del backoff
	// +optional
	// +kubebuilder:default="5s"
	Duration string `json:"duration,omitempty"`

	// Factor es el factor multiplicador del backoff
	// +optional
	// +kubebuilder:validation:Minimum=1.0
	// +kubebuilder:default=2.0
	Factor float64 `json:"factor,omitempty"`

	// MaxDuration es la duración máxima del backoff
	// +optional
	// +kubebuilder:default="3m"
	MaxDuration string `json:"maxDuration,omitempty"`
}

// ArgoCDClusterRef referencia un cluster en ArgoCD
type ArgoCDClusterRef struct {
	// Name es el nombre del cluster en ArgoCD
	Name string `json:"name"`

	// ConfigServer es la URL del servidor de configuración
	// +optional
	ConfigServer string `json:"configServer,omitempty"`

	// Labels contiene etiquetas para el cluster en ArgoCD
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// CrossplaneSpec define la configuración de Crossplane
type CrossplaneSpec struct {
	// Enabled indica si Crossplane está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Version es la versión de Crossplane a desplegar
	// +optional
	// +kubebuilder:default="v1.13.0"
	Version string `json:"version,omitempty"`

	// ProviderConfigs contiene las configuraciones de proveedores
	// +optional
	ProviderConfigs []CrossplaneProviderConfig `json:"providerConfigs,omitempty"`

	// CompositionRevisions contiene las revisiones de composiciones
	// +optional
	CompositionRevisions []CompositionRevisionSpec `json:"compositionRevisions,omitempty"`
}

// CrossplaneProviderConfig define la configuración de un proveedor
type CrossplaneProviderConfig struct {
	// Name es el nombre del provider config
	Name string `json:"name"`

	// Region es la región del proveedor
	Region string `json:"region"`

	// CredentialsSecretRef referencia el secret con credenciales
	CredentialsSecretRef SecretRef `json:"credentialsSecretRef"`
}

// CompositionRevisionSpec define una revisión de composición
type CompositionRevisionSpec struct {
	// Name es el nombre de la composición
	Name string `json:"name"`

	// Revision es el número de revisión
	// +kubebuilder:validation:Minimum=1
	Revision int `json:"revision"`
}

// VaultSpec define la configuración de Vault
type VaultSpec struct {
	// Enabled indica si Vault está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Address es la dirección del servidor Vault
	// +kubebuilder:validation:Format=uri
	Address string `json:"address"`

	// Namespace es el namespace de Vault (para Vault Enterprise)
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// AuthMethod es el método de autenticación
	// +optional
	// +kubebuilder:validation:Enum=kubernetes;aws;azure;gcp;token
	// +kubebuilder:default="kubernetes"
	AuthMethod string `json:"authMethod,omitempty"`

	// SecretsEngine es el motor de secrets a usar
	// +optional
	// +kubebuilder:validation:Enum=kv-v2;kv-v1;dynamic
	// +kubebuilder:default="kv-v2"
	SecretsEngine string `json:"secretsEngine,omitempty"`

	// Paths contiene las rutas de secrets a replicar
	// +optional
	Paths []string `json:"paths,omitempty"`

	// ReplicationInterval es el intervalo de replicación de secrets
	// +optional
	// +kubebuilder:default="5m"
	ReplicationInterval string `json:"replicationInterval,omitempty"`
}

// AWSSpec define la configuración de AWS
type AWSSpec struct {
	// Regions contiene la configuración por región
	Regions []AWSRegionConfig `json:"regions"`

	// HealthCheckInterval es el intervalo de health checks
	// +optional
	// +kubebuilder:default="30s"
	HealthCheckInterval string `json:"healthCheckInterval,omitempty"`

	// FailoverThreshold define el umbral para triggering de failover
	// +optional
	FailoverThreshold *AWSFailoverThreshold `json:"failoverThreshold,omitempty"`
}

// AWSRegionConfig define la configuración de una región AWS
type AWSRegionConfig struct {
	// Name es el nombre de la región (ej. us-east-1)
	Name string `json:"name"`

	// IsPrimary indica si es la región primaria
	// +optional
	IsPrimary bool `json:"isPrimary,omitempty"`

	// HealthCheckConfig contiene la configuración de health check
	// +optional
	HealthCheckConfig *AWSHealthCheckConfig `json:"healthCheckConfig,omitempty"`
}

// AWSHealthCheckConfig define la configuración de health check
type AWSHealthCheckConfig struct {
	// Endpoint es el endpoint a verificar
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// Timeout es el timeout del health check
	// +optional
	// +kubebuilder:default="5s"
	Timeout string `json:"timeout,omitempty"`

	// Interval es el intervalo entre health checks
	// +optional
	// +kubebuilder:default="10s"
	Interval string `json:"interval,omitempty"`
}

// AWSFailoverThreshold define los umbrales para failover
type AWSFailoverThreshold struct {
	// MinHealthyClusters es el número mínimo de clusters saludables para no hacer failover
	// +kubebuilder:validation:Minimum=1
	MinHealthyClusters int `json:"minHealthyClusters"`

	// MaxLatencyMs es la latencia máxima aceptable en milisegundos
	// +kubebuilder:validation:Minimum=0
	MaxLatencyMs int `json:"maxLatencyMs"`

	// ErrorRateThreshold es el umbral de tasa de errores (0.0 - 1.0)
	// +kubebuilder:validation:Minimum=0.0
	// +kubebuilder:validation:Maximum=1.0
	ErrorRateThreshold float64 `json:"errorRateThreshold"`
}

// ReplicationSpec define la configuración de replicación de estado
type ReplicationSpec struct {
	// Backend es el backend de replicación a usar
	// +optional
	// +kubebuilder:validation:Enum=etcd;consul;dynamodb
	// +kubebuilder:default="etcd"
	Backend string `json:"backend,omitempty"`

	// Endpoints son los endpoints de los clusters de replicación
	// +kubebuilder:validation:MinItems=3
	Endpoints []string `json:"endpoints"`

	// QuorumSize es el tamaño del quórum para consistencia
	// +optional
	// +kubebuilder:validation:Minimum=1
	QuorumSize int `json:"quorumSize,omitempty"`

	// SyncInterval es el intervalo de sincronización
	// +optional
	// +kubebuilder:default="1s"
	SyncInterval string `json:"syncInterval,omitempty"`

	// ConsistencyMode define el modo de consistencia
	// +optional
	// +kubebuilder:validation:Enum=eventual;strong
	// +kubebuilder:default="eventual"
	ConsistencyMode string `json:"consistencyMode,omitempty"`
}

// FailoverSpec define la configuración de disaster recovery
type FailoverSpec struct {
	// Enabled indica si el failover está habilitado
	// +optional
	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// Strategy define la estrategia de failover
	// +optional
	// +kubebuilder:validation:Enum=automatic;manual;semi-automatic
	// +kubebuilder:default="automatic"
	Strategy string `json:"strategy,omitempty"`

	// DetectionTimeout es el timeout para detección de fallos
	// +optional
	// +kubebuilder:default="30s"
	DetectionTimeout string `json:"detectionTimeout,omitempty"`

	// PromotionTimeout es el timeout para promoción de nuevo primary
	// +optional
	// +kubebuilder:default="5m"
	PromotionTimeout string `json:"promotionTimeout,omitempty"`

	// HealthCheckRetries es el número de reintentos para health checks
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=3
	HealthCheckRetries int `json:"healthCheckRetries,omitempty"`

	// PrePromotionHooks son hooks a ejecutar antes de la promoción
	// +optional
	PrePromotionHooks []string `json:"prePromotionHooks,omitempty"`

	// PostPromotionHooks son hooks a ejecutar después de la promoción
	// +optional
	PostPromotionHooks []string `json:"postPromotionHooks,omitempty"`
}

// SecretRef referencia un Secret de Kubernetes
type SecretRef struct {
	// Name es el nombre del secret
	Name string `json:"name"`

	// Namespace es el namespace del secret
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// ControlPlaneStatus define el estado observado del control plane
type ControlPlaneStatus struct {
	// ObservedGeneration es la generación observada
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Clusters contiene el estado de cada cluster
	// +optional
	Clusters []ClusterStatus `json:"clusters,omitempty"`

	// ArgoCD contiene el estado de ArgoCD
	// +optional
	ArgoCD *ArgoCDStatus `json:"argocd,omitempty"`

	// Crossplane contiene el estado de Crossplane
	// +optional
	Crossplane *CrossplaneStatus `json:"crossplane,omitempty"`

	// Vault contiene el estado de Vault
	// +optional
	Vault *VaultStatus `json:"vault,omitempty"`

	// Conditions contiene las condiciones actuales
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// FailoverStatus contiene el estado del failover
	// +optional
	FailoverStatus *FailoverStatus `json:"failoverStatus,omitempty"`
}

// ClusterStatus contiene el estado de un cluster
type ClusterStatus struct {
	// Name es el nombre del cluster
	Name string `json:"name"`

	// Region es la región del cluster
	Region string `json:"region"`

	// Ready indica si el cluster está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Health indica el estado de salud del cluster
	// +optional
	// +kubebuilder:validation:Enum=healthy;degraded;unhealthy;unknown
	Health string `json:"health,omitempty"`

	// LastHeartbeat es la última vez que se recibió heartbeat
	// +optional
	LastHeartbeat *metav1.Time `json:"lastHeartbeat,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// ArgoCDStatus contiene el estado de ArgoCD
type ArgoCDStatus struct {
	// Ready indica si ArgoCD está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Version es la versión desplegada
	// +optional
	Version string `json:"version,omitempty"`

	// SyncStatus es el estado de sincronización
	// +optional
	SyncStatus string `json:"syncStatus,omitempty"`

	// RegisteredClusters es el número de clusters registrados
	// +optional
	RegisteredClusters int `json:"registeredClusters,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// CrossplaneStatus contiene el estado de Crossplane
type CrossplaneStatus struct {
	// Ready indica si Crossplane está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Version es la versión desplegada
	// +optional
	Version string `json:"version,omitempty"`

	// ProviderConfigsReady es el número de provider configs listos
	// +optional
	ProviderConfigsReady int `json:"providerConfigsReady,omitempty"`

	// CompositionsReady es el número de composiciones listas
	// +optional
	CompositionsReady int `json:"compositionsReady,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// VaultStatus contiene el estado de Vault
type VaultStatus struct {
	// Ready indica si Vault está listo
	// +optional
	Ready bool `json:"ready,omitempty"`

	// Address es la dirección de Vault
	// +optional
	Address string `json:"address,omitempty"`

	// Sealed indica si Vault está sellado
	// +optional
	Sealed bool `json:"sealed,omitempty"`

	// ReplicationStatus es el estado de replicación de secrets
	// +optional
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// FailoverStatus contiene el estado del failover
type FailoverStatus struct {
	// Active indica si hay un failover activo
	// +optional
	Active bool `json:"active,omitempty"`

	// SourceCluster es el cluster de origen (el que falló)
	// +optional
	SourceCluster string `json:"sourceCluster,omitempty"`

	// TargetCluster es el cluster objetivo (el nuevo primary)
	// +optional
	TargetCluster string `json:"targetCluster,omitempty"`

	// StartedAt es cuando empezó el failover
	// +optional
	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	// CompletedAt es cuando completó el failover
	// +optional
	CompletedAt *metav1.Time `json:"completedAt,omitempty"`

	// State es el estado actual del failover
	// +optional
	// +kubebuilder:validation:Enum=detecting;promoting;replicating;completed;failed
	State string `json:"state,omitempty"`

	// Message contiene mensajes adicionales
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=controlplanes,scope=Cluster,shortName=cp
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Clusters",type="integer",JSONPath=".status.clusters\.size()"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ControlPlane es el recurso que representa la configuración del control plane multi-cluster
type ControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec es la especificación del control plane
	Spec ControlPlaneSpec `json:"spec,omitempty"`

	// Status es el estado observado del control plane
	// +optional
	Status ControlPlaneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPlaneList contiene una lista de ControlPlane
type ControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ControlPlane{}, &ControlPlaneList{})
}

// Helper functions para condiciones
const (
	ConditionArgoCDReady     = "ArgoCDReady"
	ConditionCrossplaneReady = "CrossplaneReady"
	ConditionVaultReady      = "VaultReady"
	ConditionClustersReady   = "ClustersReady"
	ConditionFailoverReady   = "FailoverReady"
	ConditionReplicating     = "Replicating"
)

// IsReady verifica si el control plane está listo
func (cp *ControlPlane) IsReady() bool {
	for _, cond := range cp.Status.Conditions {
		if cond.Type == "Ready" && cond.Status == metav1.ConditionTrue {
			return true
		}
	}
	return false
}

// GetActiveFailover retorna el failover activo si existe
func (cp *ControlPlane) GetActiveFailover() *FailoverStatus {
	if cp.Status.FailoverStatus != nil && cp.Status.FailoverStatus.Active {
		return cp.Status.FailoverStatus
	}
	return nil
}

// GetPrimaryCluster retorna el cluster primario
func (cp *ControlPlane) GetPrimaryCluster() *ClusterSpec {
	for i := range cp.Spec.Clusters {
		if cp.Spec.Clusters[i].IsMain {
			return &cp.Spec.Clusters[i]
		}
	}
	return nil
}

// GetClustersByRegion retorna los clusters de una región específica
func (cp *ControlPlane) GetClustersByRegion(region string) []ClusterSpec {
	var result []ClusterSpec
	for _, c := range cp.Spec.Clusters {
		if c.Region == region {
			result = append(result, c)
		}
	}
	return result
}


// === ARCHIVO: pkg/argocd/client.go ===
package argocd

import (
	"context"
	"fmt"
	"time"

	argoclient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	appclient "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	clusterclient "github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultArgoCDTimeout = 30 * time.Second
	SyncRetryLimit       = 3
	SyncRetryInterval    = 5 * time.Second
)

type SyncPolicy struct {
	AutoSync          bool     `json:"autoSync"`
	SyncOptions       []string `json:"syncOptions"`
	SelfHeal          bool     `json:"selfHeal"`
	PrunePropagation  string   `json:"prunePropagation"`
	DryRun            bool     `json:"dryRun"`
}

type ClusterFederationConfig struct {
	SourceCluster    string
	TargetClusters   []string
	SyncPolicy       SyncPolicy
	ProjectName      string
	NamespaceMapping map[string]string
}

type ArgoCDClient struct {
	client    argoclient.Client
	opts      argoclient.ClientOptions
	namespace string
}

func NewArgoCDClient(config ArgoCDConfig, namespace string) (*ArgoCDClient, error) {
	if config.Server == "" {
		return nil, fmt.Errorf("ArgoCD server address is required")
	}

	opts := argoclient.ClientOptions{
		ServerAddr: config.Server,
		AuthToken:  config.Token,
		Insecure:   config.Insecure,
		TLSClientConfig: argoclient.TLSClientConfig{
			Insecure: config.Insecure,
		},
	}

	return &ArgoCDClient{
		opts:      opts,
		namespace: namespace,
	}, nil
}

func (a *ArgoCDClient) Connect(ctx context.Context) error {
	client, err := argoclient.NewClient(&a.opts)
	if err != nil {
		return fmt.Errorf("failed to create ArgoCD client: %w", err)
	}
	a.client = client
	return nil
}

func (a *ArgoCDClient) SyncApplication(ctx context.Context, appName string, syncPolicy SyncPolicy) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	appClient, err := a.client.NewApplicationClient()
	if err != nil {
		return fmt.Errorf("failed to create application client: %w", err)
	}

	req := &appclient.ApplicationSyncRequest{
		Name:         &appName,
		DryRun:       syncPolicy.DryRun,
		SyncOptions:  v1alpha1.SyncOptions(syncPolicy.SyncOptions),
		Prune:        true,
		RetryStrategy: &v1alpha1.RetryStrategy{
			Limit: int64(SyncRetryLimit),
		},
	}

	syncResp, err := appClient.Sync(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to sync application %s: %w", appName, err)
	}

	log.Log.Info("Application sync initiated", "app", appName, "syncResult", syncResp.Result)
	return nil
}

func (a *ArgoCDClient) GetApplicationState(ctx context.Context, appName string) (*v1alpha1.ApplicationState, error) {
	if a.client == nil {
		return nil, fmt.Errorf("ArgoCD client not connected")
	}

	appClient, err := a.client.NewApplicationClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create application client: %w", err)
	}

	app, err := appClient.Get(ctx, &appclient.ApplicationQuery{
		Name:    &appName,
		Project: a.namespace,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get application %s: %w", appName, err)
	}

	return &app.Status, nil
}

func (a *ArgoCDClient) RegisterCluster(ctx context.Context, cluster *v1alpha1.Cluster) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	clusterClient, err := a.client.NewClusterClient()
	if err != nil {
		return fmt.Errorf("failed to create cluster client: %w", err)
	}

	_, err = clusterClient.Create(ctx, &clusterclient.ClusterCreateRequest{
		Cluster: cluster,
	})
	if err != nil {
		return fmt.Errorf("failed to register cluster %s: %w", cluster.Server, err)
	}

	log.Log.Info("Cluster registered in ArgoCD", "server", cluster.Server, "name", cluster.Name)
	return nil
}

func (a *ArgoCDClient) FederateClusters(ctx context.Context, config ClusterFederationConfig) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	appClient, err := a.client.NewApplicationClient()
	if err != nil {
		return fmt.Errorf("failed to create application client: %w", err)
	}

	for _, targetCluster := range config.TargetClusters {
		app := &v1alpha1.Application{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s", config.ProjectName, targetCluster),
				Namespace: a.namespace,
			},
			Spec: v1alpha1.ApplicationSpec{
				Project: config.ProjectName,
				Source: v1alpha1.ApplicationSource{
					RepoURL:        "https://github.com/example/controlplane-repo",
					Path:           config.NamespaceMapping[targetCluster],
					targetRevision: "main",
				},
				Destination: v1alpha1.ApplicationDestination{
					Server:    targetCluster,
					Namespace: "default",
				},
				SyncPolicy: &v1alpha1.SyncPolicy{
					Automated: &v1alpha1.SyncPolicyAutomated{
						SelfHeal: config.SyncPolicy.SelfHeal,
						Prune:    true,
					},
				},
			},
		}

		_, err = appClient.Create(ctx, &appclient.ApplicationCreateRequest{
			Application: app,
		})
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				return fmt.Errorf("failed to create federated application for cluster %s: %w", targetCluster, err)
			}
			log.Log.Info("Application already exists, updating", "cluster", targetCluster)
		}

		log.Log.Info("Cluster federated in ArgoCD", "source", config.SourceCluster, "target", targetCluster)
	}

	return nil
}

func (a *ArgoCDClient) SetAutoSync(ctx context.Context, appName string, policy SyncPolicy) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	return nil
}

func (a *ArgoCDClient) GetClusterList(ctx context.Context) (*v1alpha1.ClusterList, error) {
	if a.client == nil {
		return nil, fmt.Errorf("ArgoCD client not connected")
	}

	clusterClient, err := a.client.NewClusterClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create cluster client: %w", err)
	}

	return clusterClient.List(ctx, &clusterclient.ClusterQuery{})
}

func (a *ArgoCDClient) DeleteCluster(ctx context.Context, server string) error {
	if a.client == nil {
		return fmt.Errorf("ArgoCD client not connected")
	}

	clusterClient, err := a.client.NewClusterClient()
	if err != nil {
		return fmt.Errorf("failed to create cluster client: %w", err)
	}

	err = clusterClient.Delete(ctx, &clusterclient.ClusterQuery{Server: server})
	if err != nil {
		return fmt.Errorf("failed to delete cluster %s: %w", server, err)
	}

	log.Log.Info("Cluster deleted from ArgoCD", "server", server)
	return nil
}

type ArgoCDConfig struct {
	Server   string
	Token    string
	Insecure bool
}

func (a *ArgoCDClient) Reconcile(ctx context.Context, req types.NamespacedName) error {
	log.Log.Info("Reconciling ArgoCD resources", "name", req.Name)

	return nil
}

var _ client.Reconciler = &ArgoCDClient{}

// === ARCHIVO: pkg/crossplane/provider.go ===
package crossplane

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/iam"
	crossplanev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	ProvisionTimeout = 15 * time.Minute
	PollInterval     = 30 * time.Second
)

type ProvisionedResource struct {
	ID        string
	Type      string
	Endpoint  string
	Status    string
	CreatedAt time.Time
}

type RDSConfig struct {
	InstanceClass    string
	Engine           string
	EngineVersion    string
	AllocatedStorage int64
	MultiAZ          bool
	StorageEncrypted bool
	BackupRetention  int64
}

type S3Config struct {
	BucketName    string
	Region        string
	Versioning    bool
	Encryption    string
	PublicAccess  bool
	LifecyclePolicy *LifecyclePolicy
}

type LifecyclePolicy struct {
	Transitions []LifecycleTransition
}

type LifecycleTransition struct {
	Days          int
	StorageClass  string
}

type IAMRoleConfig struct {
	RoleName        string
	PolicyARNs      []string
	AssumeRolePolicy string
	Tags            map[string]string
}

type CrossplaneProvider struct {
	client       client.Client
	awsConfig    *AWSConfig
	resourceDefs map[string]resource.Managed
}

type AWSConfig struct {
	Region    string
	AccountID string
}

func NewCrossplaneProvider(k8sClient client.Client, awsCfg *AWSConfig) *CrossplaneProvider {
	return &CrossplaneProvider{
		client:    k8sClient,
		awsConfig: awsCfg,
		resourceDefs: make(map[string]resource.Managed),
	}
}

func (cp *CrossplaneProvider) ProvisionRDS(ctx context.Context, name string, cfg RDSConfig) (*ProvisionedResource, error) {
	log.Log.Info("Provisioning RDS instance", "name", name, "class", cfg.InstanceClass)

	svc := rds.New(cp.newAWSSession())

	input := &rds.CreateDBInstanceInput{
		DBInstanceIdentifier: aws.String(name),
		DBInstanceClass:      aws.String(cfg.InstanceClass),
		Engine:               aws.String(cfg.Engine),
		EngineVersion:        aws.String(cfg.EngineVersion),
		AllocatedStorage:     aws.Int64(cfg.AllocatedStorage),
		MultiAZ:              aws.Bool(cfg.MultiAZ),
		StorageEncrypted:     aws.Bool(cfg.StorageEncrypted),
		BackupRetentionPeriod: aws.Int64(cfg.BackupRetention),
		MasterUsername:       aws.String("admin"),
		MasterUserPassword:   aws.String(generatePassword()),
		PubliclyAccessible:   aws.Bool(false),
		tags: []*rds.Tag{
			{Key: aws.String("managed-by"), Value: aws.String("crossplane")},
			{Key: aws.String("environment"), Value: aws.String("production")},
		},
	}

	result, err := svc.CreateDBInstance(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create RDS instance: %w", err)
	}

	pr := &ProvisionedResource{
		ID:        *result.DBInstance.DBInstanceArn,
		Type:      "RDS",
		Endpoint:  *result.DBInstance.Endpoint.Address,
		Status:    *result.DBInstance.DBInstanceStatus,
		CreatedAt: time.Now(),
	}

	log.Log.Info("RDS instance provisioned", "arn", pr.ID, "endpoint", pr.Endpoint)
	return pr, nil
}

func (cp *CrossplaneProvider) ProvisionS3(ctx context.Context, name string, cfg S3Config) (*ProvisionedResource, error) {
	log.Log.Info("Provisioning S3 bucket", "name", name, "region", cfg.Region)

	svc := s3.New(cp.newAWSSession())

	input := &s3.CreateBucketInput{
		Bucket: aws.String(name),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(cfg.Region),
		},
	}

	_, err := svc.CreateBucket(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 bucket: %w", err)
	}

	if cfg.Versioning {
		err = svc.PutBucketVersioning(&s3.PutBucketVersioningInput{
			Bucket: aws.String(name),
			VersioningConfiguration: &s3.VersioningConfiguration{
				Status: aws.String(s3.BucketVersioningStatusEnabled),
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to enable versioning: %w", err)
		}
	}

	if cfg.Encryption != "" {
		err = svc.PutBucketEncryption(&s3.PutBucketEncryptionInput{
			Bucket: aws.String(name),
			ServerSideEncryptionConfiguration: &s3.ServerSideEncryptionConfiguration{
				Rules: []*s3.ServerSideEncryptionRule{
					{
						ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
							SSEAlgorithm: aws.String(cfg.Encryption),
						},
					},
				},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to enable encryption: %w", err)
		}
	}

	pr := &ProvisionedResource{
		ID:        name,
		Type:      "S3",
		Endpoint:  fmt.Sprintf("s3://%s", name),
		Status:    "Available",
		CreatedAt: time.Now(),
	}

	log.Log.Info("S3 bucket provisioned", "bucket", name)
	return pr, nil
}

func (cp *CrossplaneProvider) ProvisionIAMRole(ctx context.Context, name string, cfg IAMRoleConfig) (*ProvisionedResource, error) {
	log.Log.Info("Provisioning IAM role", "name", name)

	svc := iam.New(cp.newAWSSession())

	input := &iam.CreateRoleInput{
		RoleName:                 aws.String(name),
		AssumeRolePolicyDocument: aws.String(cfg.AssumeRolePolicy),
		tags: []*iam.Tag{
			{Key: aws.String("managed-by"), Value: aws.String("crossplane")},
		},
	}

	role, err := svc.CreateRole(input)
	if err != nil {
		return nil, fmt.Errorf("failed to create IAM role: %w", err)
	}

	for _, policyArn := range cfg.PolicyARNs {
		_, err = svc.AttachRolePolicy(&iam.AttachRolePolicyInput{
			RoleName:  aws.String(name),
			PolicyArn: aws.String(policyArn),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to attach policy %s: %w", policyArn, err)
		}
	}

	pr := &ProvisionedResource{
		ID:        *role.Role.Arn,
		Type:      "IAM",
		Endpoint:  *role.Role.Arn,
		Status:    "Active",
		CreatedAt: time.Now(),
	}

	log.Log.Info("IAM role provisioned", "arn", pr.ID)
	return pr, nil
}

func (cp *CrossplaneProvider) DeleteResource(ctx context.Context, resourceType, identifier string) error {
	log.Log.Info("Deleting resource", "type", resourceType, "id", identifier)

	switch resourceType {
	case "RDS":
		return cp.deleteRDS(ctx, identifier)
	case "S3":
		return cp.deleteS3(ctx, identifier)
	case "IAM":
		return cp.deleteIAM(ctx, identifier)
	default:
		return fmt.Errorf("unknown resource type: %s", resourceType)
	}
}

func (cp *CrossplaneProvider) deleteRDS(ctx context.Context, identifier string) error {
	svc := rds.New(cp.newAWSSession())
	_, err := svc.DeleteDBInstance(&rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(identifier),
		SkipFinalSnapshot:     aws.Bool(true),
		DeleteAutomatedBackups: aws.Bool(true),
	})
	return err
}

func (cp *CrossplaneProvider) deleteS3(ctx context.Context, identifier string) error {
	svc := s3.New(cp.newAWSSession())
	_, err := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(identifier),
	})
	return err
}

func (cp *CrossplaneProvider) deleteIAM(ctx context.Context, identifier string) error {
	svc := iam.New(cp.newAWSSession())
	_, err := svc.DeleteRole(&iam.DeleteRoleInput{
		RoleName: aws.String(identifier),
	})
	return err
}

func (cp *CrossplaneProvider) newAWSSession() *aws.Config {
	return aws.NewConfig().WithRegion(cp.awsConfig.Region)
}

func generatePassword() string {
	return "temporary-password-change-me"
}

type CrossplaneConfig struct {
	ProviderName string
	Region       string
}

func (cp *CrossplaneProvider) Reconcile(ctx context.Context) error {
	log.Log.Info("Reconciling Crossplane resources")
	return nil
}

var _ client.Reconciler = &CrossplaneProvider{}

// === ARCHIVO: pkg/vault/secrets.go ===
package vault

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	vaultapi "github.com/hashicorp/vault/api"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultVaultPath = "secret/data/controlplane"
	MaxRetries       = 3
	RetryInterval    = 2 * time.Second
)

type SecretRotation struct {
	Enabled        bool
	Interval       time.Duration
	GracePeriod    time.Duration
	RotationPolicy string
}

type SecretReplication struct {
	SourceRegion    string
	TargetRegions   []string
	ReplicationMode string
	FailoverEnabled bool
}

type VaultSecret struct {
	Key       string
	Value     string
	Version   int
	CreatedAt time.Time
	ExpiresAt *time.Time
}

type ExternalSecretSpec struct {
	SecretStore    string
	RemoteKey      string
	RefreshInterval time.Duration
	Data           []RemoteRef
}

type RemoteRef struct {
	RemoteKey string
	Property  string
}

type VaultClient struct {
	client      *vaultapi.Client
	config      *VaultConfig
	secretCache map[string]*VaultSecret
}

type VaultConfig struct {
	Address       string
	Token         string
	Namespace     string
	MountPath     string
	KVVersion     string
	TLSConfig     TLSConfig
	RetryConfig   RetryConfig
}

type TLSConfig struct {
	CACert     string
	CAPem      string
	ClientCert string
	ClientKey string
	SkipVerify bool
}

type RetryConfig struct {
	MaxRetries int
	Interval   time.Duration
}

func NewVaultClient(config VaultConfig) (*VaultClient, error) {
	if config.Address == "" {
		return nil, fmt.Errorf("Vault address is required")
	}

	vaultConfig := &vaultapi.Config{
		Address:    config.Address,
		MaxRetries: config.RetryConfig.MaxRetries,
	}

	if config.TLSConfig.CACert != "" {
		vaultConfig.CAFile = config.TLSConfig.CACert
	}

	if config.TLSConfig.SkipVerify {
		vaultConfig.TLSConfig.Insecure = true
	}

	client, err := vaultapi.NewClient(vaultConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}

	if config.Token != "" {
		client.SetToken(config.Token)
	}

	if config.Namespace != "" {
		client.SetNamespace(config.Namespace)
	}

	return &VaultClient{
		client:      client,
		config:      &config,
		secretCache: make(map[string]*VaultSecret),
	}, nil
}

func (v *VaultClient) WriteSecret(ctx context.Context, path string, data map[string]interface{}) error {
	fullPath := v.buildPath(path)

	secret, err := v.client.KVv2(v.config.MountPath).Put(ctx, fullPath, data)
	if err != nil {
		return fmt.Errorf("failed to write secret to Vault: %w", err)
	}

	log.Log.Info("Secret written to Vault", "path", fullPath, "version", secret.Version)
	return nil
}

func (v *VaultClient) ReadSecret(ctx context.Context, path string) (map[string]interface{}, error) {
	fullPath := v.buildPath(path)

	secret, err := v.client.KVv2(v.config.MountPath).Get(ctx, fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret from Vault: %w", err)
	}

	data := make(map[string]interface{})
	for k, val := range secret.Data {
		data[k] = val
	}

	log.Log.Info("Secret read from Vault", "path", fullPath)
	return data, nil
}

func (v *VaultClient) DeleteSecret(ctx context.Context, path string) error {
	fullPath := v.buildPath(path)

	err := v.client.KVv2(v.config.MountPath).Delete(ctx, fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete secret from Vault: %w", err)
	}

	log.Log.Info("Secret deleted from Vault", "path", fullPath)
	return nil
}

func (v *VaultClient) ListSecrets(ctx context.Context, path string) ([]string, error) {
	fullPath := v.buildPath(path)

	secrets, err := v.client.KVv2(v.config.MountPath).List(ctx, fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list secrets from Vault: %w", err)
	}

	return secrets.Keys, nil
}

func (v *VaultClient) RotateSecret(ctx context.Context, path string, rotation SecretRotation) error {
	log.Log.Info("Initiating secret rotation", "path", path, "interval", rotation.Interval)

	data, err := v.ReadSecret(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to read secret for rotation: %w", err)
	}

	data["_rotation_timestamp"] = time.Now().Unix()
	data["_rotation_version"] = data["_rotation_version"].(int) + 1

	err = v.WriteSecret(ctx, path, data)
	if err != nil {
		return fmt.Errorf("failed to write rotated secret: %w", err)
	}

	log.Log.Info("Secret rotated successfully", "path", path)
	return nil
}

func (v *VaultClient) ReplicateSecret(ctx context.Context, path string, replication SecretReplication) error {
	sourcePath := v.buildPathInRegion(path, replication.SourceRegion)

	sourceData, err := v.client.KVv2(v.config.MountPath).Get(ctx, sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source secret: %w", err)
	}

	for _, targetRegion := range replication.TargetRegions {
		targetPath := v.buildPathInRegion(path, targetRegion)

		err = v.client.KVv2(v.config.MountPath).Put(ctx, targetPath, sourceData.Data)
		if err != nil {
			log.Log.Error(err, "Failed to replicate secret to region", "region", targetRegion)
			if !replication.FailoverEnabled {
				return fmt.Errorf("failed to replicate secret to %s: %w", targetRegion, err)
			}
		}

		log.Log.Info("Secret replicated", "from", replication.SourceRegion, "to", targetRegion)
	}

	return nil
}

func (v *VaultClient) buildPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}
	return filepath.Join(DefaultVaultPath, path)
}

func (v *VaultClient) buildPathInRegion(path string, region string) string {
	basePath := v.buildPath(path)
	parts := strings.Split(basePath, "/")
	if len(parts) > 0 {
		parts = append([]string{region}, parts...)
	}
	return strings.Join(parts, "/")
}

func (v *VaultClient) GetSecretVersion(ctx context.Context, path string, version int) (map[string]interface{}, error) {
	fullPath := v.buildPath(path)

	secret, err := v.client.KVv2(v.config.MountPath).GetVersion(ctx, fullPath, version)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret version: %w", err)
	}

	data := make(map[string]interface{})
	for k, val := range secret.Data {
		data[k] = val
	}

	return data, nil
}

func (v *VaultClient) UnwrapToken(ctx context.Context, token string) (map[string]interface{}, error) {
	secret, err := v.client.Logical().Unwrap(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to unwrap token: %w", err)
	}

	return secret.Data, nil
}

func (v *VaultClient) HealthCheck(ctx context.Context) error {
	_, err := v.client.Sys().Health()
	return err
}

type ExternalSecretReconciler struct {
	client     *VaultClient
	k8sClient  client.Client
	namespace  string
}

func NewExternalSecretReconciler(vaultClient *VaultClient, k8sClient client.Client, namespace string) *ExternalSecretReconciler {
	return &ExternalSecretReconciler{
		client:    vaultClient,
		k8sClient: k8sClient,
		namespace: namespace,
	}
}

func (r *ExternalSecretReconciler) Reconcile(ctx context.Context, req types.NamespacedName) error {
	log.Log.Info("Reconciling external secret", "name", req.Name, "namespace", req.Namespace)

	data, err := r.client.ReadSecret(ctx, req.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Log.Info("Secret not found in Vault, skipping", "name", req.Name)
			return nil
		}
		return fmt.Errorf("failed to read secret: %w", err)
	}

	log.Log.Info("External secret reconciled", "name", req.Name, "keys", len(data))
	return nil
}

var _ client.Reconciler = &ExternalSecretReconciler{}

type VaultConfig struct {
	Address     string
	Token       string
	Namespace   string
	MountPath   string
}


// === ARCHIVO: pkg/failover/mediator.go ===
package failover

import (
	"context"
	"fmt"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/example/controlplane-operator/pkg/healthcheck"
)

const (
	DefaultFailureThreshold    = 3
	DefaultFailureWindow       = 3 * time.Minute
	PromotionTimeout           = 30 * time.Second
	HealthCheckInterval        = 30 * time.Second
)

type ClusterRole string

const (
	RolePrimary   ClusterRole = "primary"
	RoleStandby   ClusterRole = "standby"
	RolePromoting ClusterRole = "promoting"
)

type FailureRecord struct {
	ClusterID    string
	FailureCount int
	FirstFailure time.Time
	LastFailure  time.Time
}

type Mediator struct {
	client         client.Client
	clusterStates  map[string]ClusterState
	failureRecords map[string]*FailureRecord
	mutex          sync.RWMutex
	config         *MediatorConfig
}

type MediatorConfig struct {
	FailureThreshold int
	FailureWindow    time.Duration
	QuorumSize       int
	EnableSplitBrain bool
}

type ClusterState struct {
	ClusterID     string
	Role          ClusterRole
	Region        string
	IsHealthy     bool
	LastHealth    time.Time
	PromotedAt    *time.Time
	FailoverCount int
}

func NewMediator(cli client.Client, cfg *MediatorConfig) *Mediator {
	if cfg == nil {
		cfg = &MediatorConfig{
			FailureThreshold: DefaultFailureThreshold,
			FailureWindow:    DefaultFailureWindow,
			QuorumSize:       2,
			EnableSplitBrain: false,
		}
	}
	return &Mediator{
		client:         cli,
		clusterStates:  make(map[string]ClusterState),
		failureRecords: make(map[string]*FailureRecord),
		config:         cfg,
	}
}

func (m *Mediator) RegisterCluster(ctx context.Context, clusterID, region string, initialRole ClusterRole) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.clusterStates[clusterID] = ClusterState{
		ClusterID: clusterID,
		Role:      initialRole,
		Region:    region,
		IsHealthy: true,
		LastHealth: time.Now(),
	}

	m.failureRecords[clusterID] = &FailureRecord{
		ClusterID:    clusterID,
		FailureCount: 0,
	}

	log.FromContext(ctx).Info("Cluster registered in mediator",
		"cluster", clusterID,
		"region", region,
		"role", initialRole)

	return nil
}

func (m *Mediator) HandleHealthCheckResult(ctx context.Context, clusterID string, healthy bool) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	logger := log.FromContext(ctx)
	state, exists := m.clusterStates[clusterID]
	if !exists {
		return fmt.Errorf("cluster %s not registered", clusterID)
	}

	record := m.failureRecords[clusterID]

	if healthy {
		state.IsHealthy = true
		state.LastHealth = time.Now()
		m.clusterStates[clusterID] = state

		if record.FailureCount > 0 {
			logger.Info("Cluster recovered", "cluster", clusterID, "previous_failures", record.FailureCount)
			record.FailureCount = 0
			record.FirstFailure = time.Time{}
		}
		return nil
	}

	state.IsHealthy = false
	m.clusterStates[clusterID] = state

	now := time.Now()
	if record.FailureCount == 0 {
		record.FirstFailure = now
	}
	record.LastFailure = now
	record.FailureCount++

	logger.Info("Health check failed",
		"cluster", clusterID,
		"failure_count", record.FailureCount,
		"threshold", m.config.FailureThreshold)

	if m.shouldTriggerFailover(ctx, clusterID, record) {
		return m.initiateFailover(ctx, clusterID)
	}

	return nil
}

func (m *Mediator) shouldTriggerFailover(ctx context.Context, clusterID string, record *FailureRecord) bool {
	if record.FailureCount < m.config.FailureThreshold {
		return false
	}

	windowExpired := time.Since(record.FirstFailure) > m.config.FailureWindow
	if !windowExpired {
		return false
	}

	logger := log.FromContext(ctx)
	logger.Info("Failure threshold reached",
		"cluster", clusterID,
		"failures", record.FailureCount,
		"window", m.config.FailureWindow)

	return true
}

func (m *Mediator) initiateFailover(ctx context.Context, failedClusterID string) error {
	logger := log.FromContext(ctx)

	failedState := m.clusterStates[failedClusterID]
	if failedState.Role != RolePrimary {
		logger.Info("Failover skipped - cluster is not primary", "cluster", failedClusterID, "role", failedState.Role)
		return nil
	}

	if !m.config.EnableSplitBrain {
		if !m.hasQuorum(ctx, failedState.Region) {
			return fmt.Errorf("cannot failover: insufficient quorum in region %s", failedState.Region)
		}
	}

	standbyClusters := m.getStandbyClustersInRegion(failedState.Region, failedClusterID)
	if len(standbyClusters) == 0 {
		return fmt.Errorf("no standby clusters available in region %s", failedState.Region)
	}

	promotedCluster := standbyClusters[0]
	promotedState := m.clusterStates[promotedCluster]

	logger.Info("Initiating failover",
		"from", failedClusterID,
		"to", promotedCluster,
		"region", failedState.Region)

	promotedState.Role = RolePromoting
	m.clusterStates[promotedCluster] = promotedState

	if err := m.promoteCluster(ctx, promotedCluster); err != nil {
		logger.Error(err, "Failed to promote cluster", "cluster", promotedCluster)
		promotedState.Role = RoleStandby
		m.clusterStates[promotedCluster] = promotedState
		return err
	}

	now := time.Now()
	promotedState.Role = RolePrimary
	promotedState.PromotedAt = &now
	promotedState.FailoverCount++
	m.clusterStates[promotedCluster] = promotedState

	failedState.Role = RoleStandby
	m.clusterStates[failedClusterID] = failedState

	m.failureRecords[failedClusterID].FailureCount = 0

	logger.Info("Failover completed",
		"new_primary", promotedCluster,
		"old_primary", failedClusterID,
		"failover_count", promotedState.FailoverCount)

	return nil
}

func (m *Mediator) promoteCluster(ctx context.Context, clusterID string) error {
	logger := log.FromContext(ctx)

	logger.Info("Executing promotion sequence", "cluster", clusterID)

	time.Sleep(100 * time.Millisecond)

	logger.Info("Promotion completed", "cluster", clusterID)
	return nil
}

func (m *Mediator) getStandbyClustersInRegion(region, excludeID string) []string {
	var standbys []string
	for id, state := range m.clusterStates {
		if id != excludeID && state.Region == region && state.Role == RoleStandby && state.IsHealthy {
			standbys = append(standbys, id)
		}
	}
	return standbys
}

func (m *Mediator) hasQuorum(ctx context.Context, region string) bool {
	healthyCount := 0
	for _, state := range m.clusterStates {
		if state.Region == region && state.IsHealthy {
			healthyCount++
		}
	}
	return healthyCount >= m.config.QuorumSize
}

func (m *Mediator) GetClusterState(clusterID string) (ClusterState, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	state, ok := m.clusterStates[clusterID]
	return state, ok
}

func (m *Mediator) GetAllClusterStates() map[string]ClusterState {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	copyStates := make(map[string]ClusterState, len(m.clusterStates))
	for k, v := range m.clusterStates {
		copyStates[k] = v
	}
	return copyStates
}

func (m *Mediator) GetPrimaryCluster(region string) (string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	for id, state := range m.clusterStates {
		if state.Region == region && state.Role == RolePrimary && state.IsHealthy {
			return id, true
		}
	}
	return "", false
}

func (m *Mediator) RunHealthCheckLoop(ctx context.Context, checker healthcheck.ClusterHealthChecker) error {
	ticker := time.NewTicker(HealthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			m.mutex.RLock()
			clusterIDs := make([]string, 0, len(m.clusterStates))
			for id := range m.clusterStates {
				clusterIDs = append(clusterIDs, id)
			}
			m.mutex.RUnlock()

			for _, clusterID := range clusterIDs {
				healthy, err := checker.CheckClusterHealth(ctx, clusterID)
				if err != nil {
					log.FromContext(ctx).Error(err, "Health check error", "cluster", clusterID)
					continue
				}
				if err := m.HandleHealthCheckResult(ctx, clusterID, healthy); err != nil {
					log.FromContext(ctx).Error(err, "Failed to handle health result", "cluster", clusterID)
				}
			}
		}
	}
}

func (m *Mediator) ListClustersByRole(role ClusterRole) []ClusterState {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var result []ClusterState
	for _, state := range m.clusterStates {
		if state.Role == role {
			result = append(result, state)
		}
	}
	return result
}

func (m *Mediator) ListClustersByRegion(region string) []ClusterState {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var result []ClusterState
	for _, state := range m.clusterStates {
		if state.Region == region {
			result = append(result, state)
		}
	}
	return result
}

var _ healthcheck.ClusterHealthChecker = (*Mediator)(nil)

// === ARCHIVO: pkg/healthcheck/checker.go ===
package healthcheck

import (
	"context"
	"fmt"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultTimeout       = 30 * time.Second
	DefaultRetries       = 3
	DefaultRetryInterval = 5 * time.Second
	ComponentAPIServer   = "api-server"
	ComponentEtcd        = "etcd"
	ComponentScheduler   = "scheduler"
	ComponentController  = "controller-manager"
	ComponentProxy       = "kube-proxy"
)

type CheckResult struct {
	ClusterID    string
	Timestamp    time.Time
	Healthy      bool
	Components   map[string]ComponentStatus
	LatencyMs    int64
	ErrorMessage string
}

type ComponentStatus struct {
	Name      string
	Healthy   bool
	LatencyMs int64
	Message   string
}

type ClusterHealthChecker interface {
	CheckClusterHealth(ctx context.Context, clusterID string) (bool, error)
	GetClusterStatus(ctx context.Context, clusterID string) (*CheckResult, error)
	CheckComponent(ctx context.Context, clusterID, component string) (*ComponentStatus, error)
}

type HealthCheckConfig struct {
	Timeout           time.Duration
	Retries           int
	RetryInterval     time.Duration
	EnableDeepCheck   bool
	CheckComponents   []string
}

type DistributedChecker struct {
	client       client.Client
	config       *HealthCheckConfig
	resultCache  map[string]*CheckResult
	cacheMutex   sync.RWMutex
	regionGroups map[string][]string
}

func NewDistributedChecker(cli client.Client, cfg *HealthCheckConfig) *DistributedChecker {
	if cfg == nil {
		cfg = &HealthCheckConfig{
			Timeout:         DefaultTimeout,
			Retries:         DefaultRetries,
			RetryInterval:   DefaultRetryInterval,
			EnableDeepCheck: true,
			CheckComponents: []string{ComponentAPIServer, ComponentEtcd, ComponentScheduler, ComponentController},
		}
	}

	return &DistributedChecker{
		client:       cli,
		config:       cfg,
		resultCache:  make(map[string]*CheckResult),
		regionGroups: make(map[string][]string),
	}
}

func (d *DistributedChecker) RegisterClusterToRegion(ctx context.Context, clusterID, region string) error {
	d.regionGroups[region] = append(d.regionGroups[region], clusterID)
	log.FromContext(ctx).Info("Cluster registered for health checks", "cluster", clusterID, "region", region)
	return nil
}

func (d *DistributedChecker) CheckClusterHealth(ctx context.Context, clusterID string) (bool, error) {
	result, err := d.GetClusterStatus(ctx, clusterID)
	if err != nil {
		return false, err
	}
	return result.Healthy, nil
}

func (d *DistributedChecker) GetClusterStatus(ctx context.Context, clusterID string) (*CheckResult, error) {
	d.cacheMutex.RLock()
	if cached, ok := d.resultCache[clusterID]; ok && time.Since(cached.Timestamp) < d.config.Timeout {
		d.cacheMutex.RUnlock()
		return cached, nil
	}
	d.cacheMutex.RUnlock()

	logger := log.FromContext(ctx)
	startTime := time.Now()

	result := &CheckResult{
		ClusterID:  clusterID,
		Timestamp:  startTime,
		Components: make(map[string]ComponentStatus),
	}

	if d.config.EnableDeepCheck {
		for _, component := range d.config.CheckComponents {
			compStatus, err := d.CheckComponent(ctx, clusterID, component)
			if err != nil {
				logger.Error(err, "Component check failed", "cluster", clusterID, "component", component)
				result.Components[component] = ComponentStatus{
					Name:    component,
					Healthy: false,
					Message: err.Error(),
				}
				continue
			}
			result.Components[component] = *compStatus
		}
	} else {
		apiHealthy, err := d.checkAPIServer(ctx, clusterID)
		if err != nil {
			result.Healthy = false
			result.ErrorMessage = err.Error()
			d.updateCache(clusterID, result)
			return result, nil
		}
		result.Components[ComponentAPIServer] = ComponentStatus{
			Name:    ComponentAPIServer,
			Healthy: apiHealthy,
		}
	}

	result.Healthy = d.isClusterHealthy(result)
	result.LatencyMs = time.Since(startTime).Milliseconds()

	d.updateCache(clusterID, result)

	logger.Info("Health check completed",
		"cluster", clusterID,
		"healthy", result.Healthy,
		"latency_ms", result.LatencyMs)

	return result, nil
}

func (d *DistributedChecker) CheckComponent(ctx context.Context, clusterID, component string) (*ComponentStatus, error) {
	startTime := time.Now()

	switch component {
	case ComponentAPIServer:
		healthy, err := d.checkAPIServer(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "API server responding",
		}, err
	case ComponentEtcd:
		healthy, err := d.checkEtcd(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "etcd cluster healthy",
		}, err
	case ComponentScheduler:
		healthy, err := d.checkScheduler(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "scheduler operational",
		}, err
	case ComponentController:
		healthy, err := d.checkControllerManager(ctx, clusterID)
		return &ComponentStatus{
			Name:      component,
			Healthy:   healthy,
			LatencyMs: time.Since(startTime).Milliseconds(),
			Message:   "controller manager running",
		}, err
	default:
		return nil, fmt.Errorf("unknown component: %s", component)
	}
}

func (d *DistributedChecker) checkAPIServer(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(10 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) checkEtcd(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(20 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) checkScheduler(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(15 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) checkControllerManager(ctx context.Context, clusterID string) (bool, error) {
	time.Sleep(15 * time.Millisecond)
	return true, nil
}

func (d *DistributedChecker) isClusterHealthy(result *CheckResult) bool {
	if len(result.Components) == 0 {
		return false
	}
	for _, comp := range result.Components {
		if !comp.Healthy {
			return false
		}
	}
	return true
}

func (d *DistributedChecker) updateCache(clusterID string, result *CheckResult) {
	d.cacheMutex.Lock()
	defer d.cacheMutex.Unlock()
	d.resultCache[clusterID] = result
}

func (d *DistributedChecker) GetCachedResult(clusterID string) (*CheckResult, bool) {
	d.cacheMutex.RLock()
	defer d.cacheMutex.RUnlock()
	result, ok := d.resultCache[clusterID]
	return result, ok
}

func (d *DistributedChecker) InvalidateCache(clusterID string) {
	d.cacheMutex.Lock()
	defer d.cacheMutex.Unlock()
	delete(d.resultCache, clusterID)
}

func (d *DistributedChecker) GetRegionStatus(ctx context.Context, region string) (*CheckResult, error) {
	d.cacheMutex.RLock()
	clusterIDs := d.regionGroups[region]
	d.cacheMutex.RUnlock()

	if len(clusterIDs) == 0 {
		return nil, fmt.Errorf("no clusters registered in region %s", region)
	}

	regionResult := &CheckResult{
		ClusterID:  region,
		Timestamp:  time.Now(),
		Components: make(map[string]ComponentStatus),
		Healthy:    true,
	}

	healthyClusters := 0
	for _, clusterID := range clusterIDs {
		result, err := d.GetClusterStatus(ctx, clusterID)
		if err != nil {
			regionResult.Healthy = false
			continue
		}
		if result.Healthy {
			healthyClusters++
		}
		regionResult.Components[clusterID] = ComponentStatus{
			Name:    clusterID,
			Healthy: result.Healthy,
			Message: fmt.Sprintf("%d/%d components healthy",
				len(result.Components), len(result.Components)),
		}
	}

	if healthyClusters == 0 {
		regionResult.Healthy = false
		regionResult.ErrorMessage = "no healthy clusters in region"
	}

	return regionResult, nil
}

func (d *DistributedChecker) GetAllRegionsStatus(ctx context.Context) (map[string]*CheckResult, error) {
	results := make(map[string]*CheckResult)

	d.cacheMutex.RLock()
	regions := make([]string, 0, len(d.regionGroups))
	for region := range d.regionGroups {
		regions = append(regions, region)
	}
	d.cacheMutex.RUnlock()

	for _, region := range regions {
		result, err := d.GetRegionStatus(ctx, region)
		if err != nil {
			log.FromContext(ctx).Error(err, "Failed to get region status", "region", region)
			continue
		}
		results[region] = result
	}

	return results, nil
}

func (d *DistributedChecker) CheckClusterWithRetry(ctx context.Context, clusterID string) (bool, error) {
	var lastErr error
	for i := 0; i < d.config.Retries; i++ {
		healthy, err := d.CheckClusterHealth(ctx, clusterID)
		if err == nil {
			return healthy, nil
		}
		lastErr = err
		log.FromContext(ctx).Info("Health check retry",
			"cluster", clusterID,
			"attempt", i+1,
			"max", d.config.Retries,
			"error", err)

		if i < d.config.Retries-1 {
			time.Sleep(d.config.RetryInterval)
		}
	}
	return false, lastErr
}

// === ARCHIVO: pkg/rbac/federator.go ===
package rbac

import (
	"context"
	"fmt"
	"sync"
	"time"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	DefaultSyncInterval  = 5 * time.Minute
	MaxRetryAttempts    = 3
	PropagationTimeout  = 30 * time.Second
)

type SyncDirection string

const (
	SyncDirectionPush SyncDirection = "push"
	SyncDirectionPull SyncDirection = "pull"
	SyncDirectionBoth SyncDirection = "bidirectional"
)

type Permission struct {
	Subjects    []rbacv1.Subject
	RoleRef     rbacv1.RoleRef
	Resources   []string
	Verbs       []string
	ClusterWide bool
}

type SyncStatus struct {
	ClusterID    string
	Timestamp    time.Time
	Success      bool
	Permissions  int
	ErrorMessage string
}

type FederatorConfig struct {
	SyncInterval     time.Duration
	Direction        SyncDirection
	EnableAutoSync   bool
	ConflictStrategy ConflictResolution
}

type ConflictResolution string

const (
	ConflictOverwrite ConflictResolution = "overwrite"
	ConflictMerge     ConflictResolution = "merge"
	ConflictFail      ConflictResolution = "fail"
)

type RBACFederator struct {
	client     client.Client
	config     *FederatorConfig
	policies   map[string]*Permission
	policyMux  sync.RWMutex
	syncStatus map[string]*SyncStatus
	statusMux  sync.RWMutex
	clusters   map[string]ClusterEndpoint
}

type ClusterEndpoint struct {
	ClusterID string
	Region    string
	Endpoint  string
	APIClient client.Client
}

func NewRBACFederator(cli client.Client, cfg *FederatorConfig) *RBACFederator {
	if cfg == nil {
		cfg = &FederatorConfig{
			SyncInterval:     DefaultSyncInterval,
			Direction:        SyncDirectionPush,
			EnableAutoSync:   true,
			ConflictStrategy: ConflictMerge,
		}
	}

	return &RBACFederator{
		client:     cli,
		config:     cfg,
		policies:   make(map[string]*Permission),
		syncStatus: make(map[string]*SyncStatus),
		clusters:   make(map[string]ClusterEndpoint),
	}
}

func (f *RBACFederator) RegisterCluster(ctx context.Context, endpoint ClusterEndpoint) error {
	f.statusMux.Lock()
	defer f.statusMuncUnlock()

	f.clusters[endpoint.ClusterID] = endpoint
	log.FromContext(ctx).Info("Cluster registered for RBAC federation",
		"cluster", endpoint.ClusterID,
		"region", endpoint.Region)

	return nil
}

func (f *RBACFederator) DefineGlobalPolicy(ctx context.Context, policyID string, perm Permission) error {
	f.policyMux.Lock()
	defer f.policyMux.Unlock()

	if len(perm.Verbs) == 0 {
		return fmt.Errorf("permission %s must have at least one verb", policyID)
	}

	if len(perm.Resources) == 0 {
		return fmt.Errorf("permission %s must have at least one resource", policyID)
	}

	f.policies[policyID] = &perm

	log.FromContext(ctx).Info("Global policy defined",
		"policy", policyID,
		"resources", perm.Resources,
		"verbs", perm.Verbs,
		"clusterWide", perm.ClusterWide)

	return nil
}

func (f *RBACFederator) SyncToCluster(ctx context.Context, clusterID string) error {
	f.statusMux.Lock()
	endpoint, exists := f.clusters[clusterID]
	f.statusMux.Unlock()

	if !exists {
		return fmt.Errorf("cluster %s not registered", clusterID)
	}

	logger := log.FromContext(ctx).WithValues("target_cluster", clusterID)

	f.policyMux.RLock()
	policies := make([]*Permission, 0, len(f.policies))
	for _, p := range f.policies {
		policies = append(policies, p)
	}
	f.policyMux.RUnlock()

	successCount := 0
	var lastErr error

	for _, policy := range policies {
		err := f.applyPolicyToCluster(ctx, endpoint, policy)
		if err != nil {
			logger.Error(err, "Failed to apply policy", "policy", policy)
			lastErr = err
			continue
		}
		successCount++
	}

	status := &SyncStatus{
		ClusterID:   clusterID,
		Timestamp:   time.Now(),
		Success:     lastErr == nil,
		Permissions: successCount,
	}
	if lastErr != nil {
		status.ErrorMessage = lastErr.Error()
	}

	f.statusMux.Lock()
	f.syncStatus[clusterID] = status
	f.statusMux.Unlock()

	logger.Info("Sync completed",
		"success", status.Success,
		"permissions", successCount)

	return lastErr
}

func (f *RBACFederator) applyPolicyToCluster(ctx context.Context, endpoint ClusterEndpoint, policy *Permission) error {
	roleName := "federated-role"

	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleName,
			Namespace: "default",
			Labels: map[string]string{
				"federated":   "true",
				"managed-by": "controlplane-operator",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups:     []string{"*"},
				Resources:     policy.Resources,
				Verbs:         policy.Verbs,
				ResourceNames: nil,
			},
		},
	}

	err := endpoint.APIClient.Create(ctx, role)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create role: %w", err)
	}

	if errors.IsAlreadyExists(err) {
		existing := &rbacv1.Role{}
		if err := endpoint.APIClient.Get(ctx, client.ObjectKey{Name: roleName, Namespace: "default"}, existing); err != nil {
			return fmt.Errorf("failed to get existing role: %w", err)
		}

		if f.config.ConflictStrategy == ConflictOverwrite {
			existing.Rules = role.Rules
			return endpoint.APIClient.Update(ctx, existing)
		} else if f.config.ConflictStrategy == ConflictMerge {
			existing.Rules = mergeRules(existing.Rules, role.Rules)
			return endpoint.APIClient.Update(ctx, existing)
		}
	}

	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "federated-role-binding",
			Namespace: "default",
			Labels: map[string]string{
				"federated":   "true",
				"managed-by": "controlplane-operator",
			},
		},
		Subjects: policy.Subjects,
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.GroupName,
			Kind:     "Role",
			Name:     roleName,
		},
	}

	err = endpoint.APIClient.Create(ctx, roleBinding)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to create role binding: %w", err)
	}

	return nil
}

func mergeRules(existing, new []rbacv1.PolicyRule) []rbacv1.PolicyRule {
	ruleMap := make(map[string]rbacv1.PolicyRule)

	for _, r := range existing {
		key := fmt.Sprintf("%v:%v", r.APIGroups, r.Resources)
		ruleMap[key] = r
	}

	for _, r := range new {
		key := fmt.Sprintf("%v:%v", r.APIGroups, r.Resources)
		if existingRule, ok := ruleMap[key]; ok {
			mergedVerbs := uniqueMerge(existingRule.Verbs, r.Verbs)
			existingRule.Verbs = mergedVerbs
			ruleMap[key] = existingRule
		} else {
			ruleMap[key] = r
		}
	}

	result := make([]rbacv1.PolicyRule, 0, len(ruleMap))
	for _, r := range ruleMap {
		result = append(result, r)
	}

	return result
}

func uniqueMerge(a, b []string) []string {
	existing := make(map[string]bool)
	result := append([]string{}, a...)

	for _, v := range b {
		if !existing[v] {
			existing[v] = true
			result = append(result, v)
		}
	}

	return result
}

func (f *RBACFederator) SyncAllClusters(ctx context.Context) error {
	f.statusMux.RLock()
	clusterIDs := make([]string, 0, len(f.clusters))
	for id := range f.clusters {
		clusterIDs = append(clusterIDs, id)
	}
	f.statusMux.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(clusterIDs))

	for _, clusterID := range clusterIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if err := f.SyncToCluster(ctx, id); err != nil {
				errChan <- fmt.Errorf("cluster %s: %w", id, err)
			}
		}(clusterID)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("sync failures: %v", errors)
	}

	return nil
}

func (f *RBACFederator) GetSyncStatus(clusterID string) (*SyncStatus, bool) {
	f.statusMux.RLock()
	defer f.statusMux.RUnlock()
	status, ok := f.syncStatus[clusterID]
	return status, ok
}

func (f *RBACFederator) GetAllSyncStatuses() map[string]*SyncStatus {
	f.statusMux.RLock()
	defer f.statusMux.RUnlock()

	copyStatus := make(map[string]*SyncStatus, len(f.syncStatus))
	for k, v := range f.syncStatus {
		copyStatus[k] = v
	}

	return copyStatus
}

func (f *RBACFederator) GetPolicy(policyID string) (*Permission, bool) {
	f.policyMux.RLock()
	defer f.policyMux.RUnlock()
	policy, ok := f.policies[policyID]
	return policy, ok
}

func (f *RBACFederator) ListPolicies() []*Permission {
	f.policyMux.RLock()
	defer f.policyMux.RUnlock()

	policies := make([]*Permission, 0, len(f.policies))
	for _, p := range f.policies {
		policies = append(policies, p)
	}

	return policies
}

func (f *RBACFederator) SyncRegion(ctx context.Context, region string) error {
	f.statusMux.RLock()
	var regionClusters []string
	for id, endpoint := range f.clusters {
		if endpoint.Region == region {
			regionClusters = append(regionClusters, id)
		}
	}
	f.statusMux.RUnlock()

	if len(regionClusters) == 0 {
		return fmt.Errorf("no clusters found in region %s", region)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(regionClusters))

	for _, clusterID := range regionClusters {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if err := f.SyncToCluster(ctx, id); err != nil {
				errChan <- err
			}
		}(clusterID)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("region sync failures: %v", errors)
	}

	return nil
}

func (f *RBACFederator) RunAutoSyncLoop(ctx context.Context) error {
	if !f.config.EnableAutoSync {
		return nil
	}

	ticker := time.NewTicker(f.config.SyncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			log.FromContext(ctx).Info("Starting auto-sync cycle")
			if err := f.SyncAllClusters(ctx); err != nil {
				log.FromContext(ctx).Error(err, "Auto-sync failed")
			}
		}
	}
}


// === ARCHIVO: pkg/replication/topology.go ===
package replication

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	controlplanev1alpha1 "github.com/example/controlplane-operator/api/v1alpha1"
)

type TopologyConfig struct {
	ClusterID       string
	Region          string
	EtcdEndpoints   []string
	ConsulEndpoints []string
	QuorumSize      int
	Timeout         time.Duration
}

type ClusterNode struct {
	ID        string
	Region    string
	Endpoint  string
	IsLeader  bool
	LastSeen  time.Time
	Health    bool
}

type ReplicationTopology struct {
	config        TopologyConfig
	nodes         map[string]*ClusterNode
	etcdClient    *clientv3.Client
	mu            sync.RWMutex
	watcherCancel context.CancelFunc
}

func NewReplicationTopology(cfg TopologyConfig) (*ReplicationTopology, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}
	if cfg.QuorumSize == 0 {
		cfg.QuorumSize = 3
	}

	rt := &ReplicationTopology{
		config: cfg,
		nodes:  make(map[string]*ClusterNode),
	}

	if len(cfg.EtcdEndpoints) > 0 {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   cfg.EtcdEndpoints,
			DialTimeout: cfg.Timeout,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create etcd client: %w", err)
		}
		rt.etcdClient = cli
	}

	return rt, nil
}

func (rt *ReplicationTopology) Start(ctx context.Context) error {
	logger := log.FromContext(ctx)

	if rt.etcdClient == nil {
		return fmt.Errorf("etcd client not initialized")
	}

	rt.watchClusterState(ctx)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := rt.syncClusterState(ctx); err != nil {
				logger.Error(err, "failed to sync cluster state")
			}
		}
	}
}

func (rt *ReplicationTopology) watchClusterState(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	rt.watcherCancel = cancel

	rchan := rt.etcdClient.Watch(ctx, "/clusters/", clientv3.WithPrefix())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case wresp := <-rchan:
				if wresp.Err() != nil {
					log.Log.Error(wresp.Err(), "watch error on cluster state")
					continue
				}
				for _, ev := range wresp.Events {
					rt.handleClusterEvent(ev)
				}
			}
		}
	}()
}

func (rt *ReplicationTopology) handleClusterEvent(ev *clientv3.Event) {
	rt.mu.Lock()
	defer rt.mu.Unlock()

	switch ev.Type {
	case clientv3.EventTypePut:
		nodeID := string(ev.Kv.Key)
		rt.nodes[nodeID] = &ClusterNode{
			ID:       nodeID,
			LastSeen: time.Now(),
			Health:   true,
		}
	case clientv3.EventTypeDelete:
		nodeID := string(ev.Kv.Key)
		delete(rt.nodes, nodeID)
	}
}

func (rt *ReplicationTopology) syncClusterState(ctx context.Context) error {
	logger := log.FromContext(ctx)

	resp, err := rt.etcdClient.Get(ctx, "/clusters/", clientv3.WithPrefix())
	if err != nil {
		return fmt.Errorf("failed to get cluster state: %w", err)
	}

	rt.mu.Lock()
	now := time.Now()
	for _, kv := range resp.Kvs {
		nodeID := string(kv.Key)
		if _, ok := rt.nodes[nodeID]; !ok {
			rt.nodes[nodeID] = &ClusterNode{ID: nodeID}
		}
		rt.nodes[nodeID].LastSeen = now
		rt.nodes[nodeID].Health = true
	}
	rt.mu.Unlock()

	logger.Info("cluster state synced", "nodeCount", len(rt.nodes))
	return nil
}

func (rt *ReplicationTopology) GetHealthyNodes(ctx context.Context) []*ClusterNode {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	var healthy []*ClusterNode
	for _, node := range rt.nodes {
		if node.Health && time.Since(node.LastSeen) < 2*time.Minute {
			healthy = append(healthy, node)
		}
	}
	return healthy
}

func (rt *ReplicationTopology) CheckQuorum(ctx context.Context) (bool, error) {
	healthy := rt.GetHealthyNodes(ctx)
	return len(healthy) >= rt.config.QuorumSize, nil
}

func (rt *ReplicationTopology) ElectLeader(ctx context.Context) (string, error) {
	healthy := rt.GetHealthyNodes(ctx)
	if len(healthy) == 0 {
		return "", fmt.Errorf("no healthy nodes available for leader election")
	}

	sort.Slice(healthy, func(i, j int) bool {
		return healthy[i].LastSeen.Before(healthy[j].LastSeen)
	})

	leader := healthy[0]
	rt.mu.Lock()
	rt.nodes[leader.ID].IsLeader = true
	rt.mu.Unlock()

	return leader.ID, nil
}

func (rt *ReplicationTopology) ReplicateState(ctx context.Context, key string, value []byte) error {
	hasQuorum, err := rt.CheckQuorum(ctx)
	if err != nil {
		return err
	}
	if !hasQuorum {
		return fmt.Errorf("cannot replicate: no quorum (need %d, have %d)", 
			rt.config.QuorumSize, len(rt.GetHealthyNodes(ctx)))
	}

	ctx, cancel := context.WithTimeout(ctx, rt.config.Timeout)
	defer cancel()

	_, err = rt.etcdClient.Put(ctx, key, string(value))
	return err
}

func (rt *ReplicationTopology) GetReplicatedState(ctx context.Context, key string) ([]byte, error) {
	resp, err := rt.etcdClient.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("failed to get replicated state: %w", err)
	}
	if len(resp.Kvs) == 0 {
		return nil, nil
	}
	return resp.Kvs[0].Value, nil
}

func (rt *ReplicationTopology) Stop() {
	if rt.watcherCancel != nil {
		rt.watcherCancel()
	}
	if rt.etcdClient != nil {
		rt.etcdClient.Close()
	}
}

// Replicator define el contrato para implementar estrategias de replicación
type Replicator interface {
	Replicate(ctx context.Context, state controlplanev1alpha1.ControlPlane) error
	GetReplicatedState(ctx context.Context, key string) ([]byte, error)
	CheckQuorum(ctx context.Context) (bool, error)
}

// EtcdReplicator implementa la replicación usando etcd como store distribuido
type EtcdReplicator struct {
	topology *ReplicationTopology
}

func NewEtcdReplicator(topology *ReplicationTopology) *EtcdReplicator {
	return &EtcdReplicator{topology: topology}
}

func (r *EtcdReplicator) Replicate(ctx context.Context, cp controlplanev1alpha1.ControlPlane) error {
	key := fmt.Sprintf("/clusters/%s/controlplane", cp.Name)
	value := []byte(cp.Status.State)
	return r.topology.ReplicateState(ctx, key, value)
}

func (r *EtcdReplicator) GetReplicatedState(ctx context.Context, key string) ([]byte, error) {
	return r.topology.GetReplicatedState(ctx, key)
}

func (r *EtcdReplicator) CheckQuorum(ctx context.Context) (bool, error) {
	return r.topology.CheckQuorum(ctx)
}

// GetClusterInfo obtiene información del cluster desde el CRD
func GetClusterInfo(ctx context.Context, c client.Client, nn types.NamespacedName) (*ClusterInfo, error) {
	cp := &controlplanev1alpha1.ControlPlane{}
	err := c.Get(ctx, nn, cp)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return &ClusterInfo{
		Name:      cp.Name,
		Namespace: cp.Namespace,
		Region:    cp.Spec.Region,
		State:     string(cp.Status.State),
	}, nil
}

// ClusterInfo representa la información básica de un cluster
type ClusterInfo struct {
	Name      string
	Namespace string
	Region    string
	State     string
}

// SyncClusterResources sincroniza recursos entre clusters usando la topología
func SyncClusterResources(ctx context.Context, topology *ReplicationTopology, resources []client.Object) error {
	hasQuorum, err := topology.CheckQuorum(ctx)
	if err != nil {
		return err
	}
	if !hasQuorum {
		return fmt.Errorf("cannot sync resources: quorum not reached")
	}

	for _, res := range resources {
		key := fmt.Sprintf("/clusters/%s/%s/%s", res.GetName(), res.GetObjectKind().GroupVersionKind().Kind, res.GetNamespace())
		value, err := client.MarshallJSON(res)
		if err != nil {
			return fmt.Errorf("failed to marshal resource: %w", err)
		}
		if err := topology.ReplicateState(ctx, key, value); err != nil {
			return fmt.Errorf("failed to replicate resource %s: %w", res.GetName(), err)
		}
	}
	return nil
}


// === ARCHIVO: pkg/aws/dto.go ===
package aws

import (
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/route53"
)

type HealthCheckMessage struct {
	ClusterID       string            `json:"cluster_id"`
	Region          string            `json:"region"`
	Timestamp       time.Time         `json:"timestamp"`
	Status          HealthStatus      `json:"status"`
	Components      []ComponentStatus `json:"components"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusDegraded  HealthStatus = "degraded"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
	HealthStatusUnknown   HealthStatus = "unknown"
)

type ComponentStatus struct {
	Name               string                 `json:"name"`
	Type               ComponentType          `json:"type"`
	Status             HealthStatus           `json:"status"`
	Message            string                 `json:"message,omitempty"`
	LastCheck          time.Time              `json:"last_check"`
	AvailabilityZones  []string               `json:"availability_zones,omitempty"`
	TargetGroupHealth  *TargetGroupHealth     `json:"target_group_health,omitempty"`
	Route53HealthCheck *Route53HealthCheckDTO `json:"route53_health,omitempty"`
}

type ComponentType string

const (
	ComponentTypeAPIServer    ComponentType = "api_server"
	ComponentTypeETCD         ComponentType = "etcd"
	ComponentTypeWorkerNode   ComponentType = "worker_node"
	ComponentTypeLoadBalancer ComponentType = "load_balancer"
	ComponentTypeRoute53      ComponentType = "route53"
	ComponentTypeCrossplane   ComponentType = "crossplane"
	ComponentTypeArgoCD       ComponentType = "argocd"
)

type TargetGroupHealth struct {
	TargetGroupARN  string              `json:"target_group_arn"`
	TargetGroupName string              `json:"target_group_name"`
	LoadBalancerARN string              `json:"load_balancer_arn"`
	HealthyTargets  int64               `json:"healthy_targets"`
	UnhealthyTargets int64              `json:"unhealthy_targets"`
	TotalTargets    int64               `json:"total_targets"`
	TargetStatuses  []TargetStatusDetail `json:"target_statuses"`
}

type TargetStatusDetail struct {
	TargetID  string         `json:"target_id"`
	Port      int64          `json:"port"`
	Health    string         `json:"health"`
	Reason    string         `json:"reason,omitempty"`
}

type Route53HealthCheckDTO struct {
	HealthCheckID        string   `json:"health_check_id"`
	HealthCheckStatus    string   `json:"health_check_status"`
	HealthCheckType      string   `json:"health_check_type"`
	FailureReason        string   `json:"failure_reason,omitempty"`
	Regions              []string `json:"regions"`
	IPAddress            string   `json:"ip_address,omitempty"`
	FullyQualifiedDomain string   `json:"fqdn,omitempty"`
}

type FailoverEvent struct {
	EventID            string               `json:"event_id"`
	EventType          FailoverEventType    `json:"event_type"`
	Timestamp          time.Time            `json:"timestamp"`
	SourceCluster      string               `json:"source_cluster"`
	SourceRegion       string               `json:"source_region"`
	TargetCluster      string               `json:"target_cluster,omitempty"`
	TargetRegion       string               `json:"target_region,omitempty"`
	Trigger            FailoverTrigger      `json:"trigger"`
	Severity           FailoverSeverity     `json:"severity"`
	State              FailoverState        `json:"state"`
	AffectedComponents []string             `json:"affected_components"`
	Actions            []FailoverAction     `json:"actions"`
	RollbackAvailable  bool                 `json:"rollback_available"`
	Metadata           map[string]string    `json:"metadata,omitempty"`
}

type FailoverEventType string

const (
	FailoverEventTypePlanned   FailoverEventType = "planned"
	FailoverEventTypeEmergency FailoverEventType = "emergency"
	FailoverEventTypeDrill     FailoverEventType = "drill"
)

type FailoverTrigger string

const (
	FailoverTriggerManual          FailoverTrigger = "manual"
	FailoverTriggerHealthCheck     FailoverTrigger = "health_check"
	FailoverTriggerLatency         FailoverTrigger = "latency_threshold"
	FailoverTriggerErrorRate       FailoverTrigger = "error_rate_threshold"
	FailoverTriggerRegionFailure   FailoverTrigger = "region_failure"
)

type FailoverSeverity string

const (
	FailoverSeverityInfo     FailoverSeverity = "info"
	FailoverSeverityWarning  FailoverSeverity = "warning"
	FailoverSeverityCritical FailoverSeverity = "critical"
)

type FailoverState string

const (
	FailoverStatePending    FailoverState = "pending"
	FailoverStateInProgress FailoverState = "in_progress"
	FailoverStateCompleted  FailoverState = "completed"
	FailoverStateFailed     FailoverState = "failed"
	FailoverStateRolledBack FailoverState = "rolled_back"
)

type FailoverAction struct {
	ActionType FailoverActionType `json:"action_type"`
	Target     string             `json:"target"`
	Parameters map[string]string  `json:"parameters,omitempty"`
	Status     string             `json:"status"`
	Error      string             `json:"error,omitempty"`
}

type FailoverActionType string

const (
	FailoverActionTypeUpdateDNS           FailoverActionType = "update_dns"
	FailoverActionTypeSwitchLoadBalancer  FailoverActionType = "switch_load_balancer"
	FailoverActionTypePromoteReplica      FailoverActionType = "promote_replica"
	FailoverActionTypeSyncState           FailoverActionType = "sync_state"
	FailoverActionTypeNotify              FailoverActionType = "notify"
)

type ClusterHealthSummary struct {
	ClusterID         string            `json:"cluster_id"`
	Region            string            `json:"region"`
	OverallStatus     HealthStatus      `json:"overall_status"`
	LastUpdated       time.Time         `json:"last_updated"`
	APIServerStatus   HealthStatus      `json:"api_server_status"`
	WorkerNodesStatus HealthStatus      `json:"worker_nodes_status"`
	StorageStatus     HealthStatus      `json:"storage_status"`
	NetworkStatus     HealthStatus      `json:"network_status"`
	IntegrationsStatus map[string]HealthStatus `json:"integrations_status"`
	AlertCount        int                `json:"alert_count"`
}

type RegionHealthSummary struct {
	Region            string                  `json:"region"`
	OverallStatus     HealthStatus            `json:"overall_status"`
	Clusters          []ClusterHealthSummary  `json:"clusters"`
	CrossRegionLatency time.Duration          `json:"cross_region_latency,omitempty"`
	FailoverReadiness FailoverReadiness       `json:"failover_readiness"`
}

type FailoverReadiness struct {
	ReadyForFailover bool      `json:"ready_for_failover"`
	LastCheck        time.Time `json:"last_check"`
	StateSyncStatus  string    `json:"state_sync_status"`
	SecretsReady     bool      `json:"secrets_ready"`
	RbacReady        bool      `json:"rbac_ready"`
	BlockingIssues   []string  `json:"blocking_issues,omitempty"`
}

type DNSRecordUpdate struct {
	HostedZoneID     string            `json:"hosted_zone_id"`
	RecordName       string            `json:"record_name"`
	RecordType       string            `json:"record_type"`
	Action           string            `json:"action"`
	TargetValue      string            `json:"target_value"`
	TargetType       string            `json:"target_type"`
	TTL              int64             `json:"ttl"`
	HealthCheckID    string            `json:"health_check_id,omitempty"`
	FailoverRecord   *Route53FailoverRecord `json:"failover_record,omitempty"`
}

type Route53FailoverRecord struct {
	FailoverRecordType string `json:"failover_record_type"`
	SetID             string `json:"set_id"`
	HealthCheckID     string `json:"health_check_id"`
}

type LoadBalancerTarget struct {
	LoadBalancerARN  string `json:"load_balancer_arn"`
	LoadBalancerName string `json:"load_balancer_name"`
	LoadBalancerType string `json:"load_balancer_type"`
	DNSName          string `json:"dns_name"`
	Scheme           string `json:"scheme"`
	VPCID            string `json:"vpc_id"`
	Regions          []string `json:"regions"`
	TargetGroups     []string `json:"target_groups"`
}

type StateSyncMessage struct {
	SyncID          string            `json:"sync_id"`
	SourceCluster   string            `json:"source_cluster"`
	TargetCluster   string            `json:"target_cluster"`
	Timestamp       time.Time         `json:"timestamp"`
	StateVersion    string            `json:"state_version"`
	Resources       []StateResource   `json:"resources"`
	SyncStatus      StateSyncStatus   `json:"sync_status"`
	ConflictResolution string         `json:"conflict_resolution,omitempty"`
}

type StateResource struct {
	ResourceType string `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	Namespace    string `json:namespace,omitempty"`
	Version      string `json:"version"`
	Checksum     string `json:"checksum"`
}

type StateSyncStatus string

const (
	StateSyncStatusPending   StateSyncStatus = "pending"
	StateSyncStatusInFlight StateSyncStatus = "in_flight"
	StateSyncStatusSynced   StateSyncStatus = "synced"
	StateSyncStatusFailed   StateSyncStatus = "failed"
	StateSyncStatusConflict StateSyncStatus = "conflict"
)


// === ARCHIVO: config/crd/bases/controlplane.example.com_controlplanes.yaml ===
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
    api-approved.kubernetes.io: "https://github.com/kubernetes-sigs/controller-runtime/pull/1909"
  creationTimestamp: null
  name: controlplanes.controlplane.example.com
spec:
  group: controlplane.example.com
  names:
    kind: ControlPlane
    listKind: ControlPlaneList
    plural: controlplanes
    singular: controlplane
  scope: Cluster
  versions:
    - name: v1alpha1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            apiVersion:
              type: string
            kind:
              type: string
            metadata:
              type: object
            spec:
              type: object
              properties:
                argocd:
                  type: object
                  properties:
                    namespace:
                      type: string
                    repository:
                      type: string
                    branch:
                      type: string
                    projectName:
                      type: string
                    serverUrl:
                      type: string
                    adminEnabled:
                      type: boolean
                    syncTimeout:
                      type: integer
                    retryLimit:
                      type: integer
                    controllerLogLevel:
                      type: string
                    applicationSetNamespace:
                      type: string
                  required:
                    - namespace
                    - repository
                crossplane:
                  type: object
                  properties:
                    providerName:
                      type: string
                    providerConfig:
                      type: string
                    compositionNamespace:
                      type: string
                    enableAutoComposition:
                      type: boolean
                    serviceAccountName:
                      type: string
                    controllerConfigRef:
                      type: string
                    runtimeNamespace:
                      type: string
                  required:
                    - providerName
                    - providerConfig
                vault:
                  type: object
                  properties:
                    address:
                      type: string
                    pathPrefix:
                      type: string
                    authMethod:
                      type: string
                      enum:
                        - kubernetes
                        - aws
                        - token
                    role:
                      type: string
                    k8sAuthPath:
                      type: string
                    secretEngine:
                      type: string
                    pkiRole:
                      type: string
                    caBundle:
                      type: string
                      format: byte
                    skipTLSVerify:
                      type: boolean
                  required:
                    - address
                    - pathPrefix
                    - authMethod
                aws:
                  type: object
                  properties:
                    region:
                      type: string
                    s3Bucket:
                      type: string
                    dynamoTable:
                      type: string
                    kmsKeyId:
                      type: string
                    enablePublicAccess:
                      type: boolean
                    vpcId:
                      type: string
                    subnetIds:
                      type: array
                      items:
                        type: string
                    securityGroupIds:
                      type: array
                      items:
                        type: string
                    iamRoleArn:
                      type: string
                    tags:
                      type: object
                      additionalProperties:
                        type: string
                  required:
                    - region
                    - s3Bucket
                regions:
                  type: array
                  items:
                    type: object
                    properties:
                      name:
                        type: string
                      priority:
                        type: integer
                      isPrimary:
                        type: boolean
                      clusterName:
                        type: string
                      clusterEndpoint:
                        type: string
                      clusterRegion:
                        type: string
                      vaultAddress:
                        type: string
                      enabledFeatures:
                        type: array
                        items:
                          type: string
                      healthCheckInterval:
                        type: integer
                      failoverWeight:
                        type: integer
                    required:
                      - name
                      - clusterName
                      - clusterEndpoint
                replication:
                  type: object
                  properties:
                    enabled:
                      type: boolean
                    strategy:
                      type: string
                      enum:
                        - sync
                        - async
                        - quorum
                    quorumSize:
                      type: integer
                    syncInterval:
                      type: integer
                    heartbeatInterval:
                      type: integer
                    consensusTimeout:
                      type: integer
                    conflictResolution:
                      type: string
                      enum:
                        - last-write-wins
                        - primary-wins
                        - manual
                    replicationNamespace:
                      type: string
                    etcdEndpoints:
                      type: array
                      items:
                        type: string
                  required:
                    - enabled
                    - strategy
                failover:
                  type: object
                  properties:
                    enabled:
                      type: boolean
                    autoFailover:
                      type: boolean
                    healthCheckInterval:
                      type: integer
                    failureThreshold:
                      type: integer
                    recoveryThreshold:
                      type: integer
                    preFailoverScript:
                      type: string
                    postFailoverScript:
                      type: string
                    failoverTimeout:
                      type: integer
                    rollbackEnabled:
                      type: boolean
                    maxFailoverDuration:
                      type: integer
                    healthCheckEndpoints:
                      type: array
                      items:
                        type: string
                rbac:
                  type: object
                  properties:
                    federated:
                      type: boolean
                    clusterRoles:
                      type: array
                      items:
                        type: string
                    namespaceRoles:
                      type: array
                      items:
                        type: string
                    syncInterval:
                      type: integer
                    defaultRole:
                      type: string
                    adminUsers:
                      type: array
                      items:
                        type: string
                monitoring:
                  type: object
                  properties:
                    enabled:
                      type: boolean
                    prometheusUrl:
                      type: string
                    grafanaUrl:
                      type: string
                    alertManagerUrl:
                      type: string
                    metricsPort:
                      type: integer
                    logLevel:
                      type: string
                    traceEndpoint:
                      type: string
                    retentionDays:
                      type: integer
                topology:
                  type: object
                  properties:
                    clusterCount:
                      type: integer
                    regionCount:
                      type: integer
                    meshMode:
                      type: string
                      enum:
                        - full
                        - star
                        - ring
                    serviceMesh:
                      type: string
                    ingressClass:
                      type: string
                    cniPlugin:
                      type: string
                    ccmProvider:
                      type: string
              required:
                - regions
                - argocd
                - crossplane
            status:
              type: object
              properties:
                phase:
                  type: string
                conditions:
                  type: array
                  items:
                    type: object
                    properties:
                      type:
                        type: string
                      status:
                        type: string
                      lastTransitionTime:
                        type: string
                        format: date-time
                      reason:
                        type: string
                      message:
                        type: string
                    required:
                      - type
                      - status
                regions:
                  type: array
                  items:
                    type: object
                    properties:
                      name:
                        type: string
                      status:
                        type: string
                      health:
                        type: string
                      lastSync:
                        type: string
                        format: date-time
                      error:
                        type: string
                activeFailover:
                  type: object
                  properties:
                    active:
                      type: boolean
                    sourceRegion:
                      type: string
                    targetRegion:
                      type: string
                    startTime:
                      type: string
                      format: date-time
                    completionTime:
                      type: string
                      format: date-time
                replicationStatus:
                  type: object
                  properties:
                    syncStatus:
                      type: string
                    lastSyncTime:
                      type: string
                      format: date-time
                    lag:
                      type: integer
                    healthyNodes:
                      type: array
                      items:
                        type: string
                observedGeneration:
                  type: integer
                lastReconcileTime:
                  type: string
                  format: date-time
      subresources:
        status: {}
      additionalPrinterColumns:
        - name: Status
          type: string
          JSONPath: .status.phase
        - name: Age
          type: date
          JSONPath: .metadata.creationTimestamp

---
// === ARCHIVO: config/rbac/role.yaml ===
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: controlplane-operator
  labels:
    app.kubernetes.io/name: controlplane-operator
    app.kubernetes.io/component: rbac
rules:
  - apiGroups:
      - ""
    resources:
      - configmaps
      - secrets
      - serviceaccounts
      - services
      - pods
      - pods/log
      - events
      - namespaces
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - ""
    resources:
      - configmaps/status
      - secrets/status
      - serviceaccounts/status
    verbs:
      - get
      - update
      - patch
  - apiGroups:
      - "apps"
    resources:
      - deployments
      - statefulsets
      - daemonsets
      - replicasets
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - "apps"
    resources:
      - deployments/status
      - statefulsets/status
      - daemonsets/status
    verbs:
      - get
      - update
      - patch
  - apiGroups:
      - "rbac.authorization.k8s.io"
    resources:
      - clusterrolebindings
      - clusterroles
      - rolebindings
      - roles
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
      - bind
      - escalate
  - apiGroups:
      - "rbac.authorization.k8s.io"
    resources:
      - clusterrolebindings/status
      - clusterroles/status
      - rolebindings/status
      - roles/status
    verbs:
      - get
      - update
      - patch
  - apiGroups:
      - "argoproj.io"
    resources:
      - applications
      - applicationsets
      - appprojects
      - applicationsets/finalizers
      - applications/finalizers
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - "argoproj.io"
    resources:
      - applications/status
      - applicationsets/status
      - appprojects/status
    verbs:
      - get
      - update
      - patch
  - apiGroups:
      - "pkg.crossplane.io"
    resources:
      - providers
      - providerconfigs
      - configurations
      - compositionrevisions
      - claims
      - environments
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - "pkg.crossplane.io"
    resources:
      - providers/status
      - providerconfigs/status
      - configurations/status
      - compositionrevisions/status
      - claims/status
      - environments/status
    verbs:
      - get
      - update
      - patch
  - apiGroups:
      - "apiextensions.k8s.io"
    resources:
      - customresourcedefinitions
      - customresourcedefinitions/status
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - "admissionregistration.k8s.io"
    resources:
      - validatingwebhookconfigurations
      - mutatingwebhookconfigurations
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - "coordination.k8s.io"
    resources:
      - leases
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - "controlplane.example.com"
    resources:
      - controlplanes
      - controlplanes/status
      - controlplanes/finalizers
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete

---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: controlplane-operator-leader-election
  namespace: controlplane-system
  labels:
    app.kubernetes.io/name: controlplane-operator
    app.kubernetes.io/component: rbac
rules:
  - apiGroups:
      - ""
      - coordination.k8s.io
    resources:
      - configmaps
      - leases
    verbs:
      - get
      - list
      - watch
      - create
      - update
      - patch
      - delete
  - apiGroups:
      - ""
    resources:
      - configmaps
    resourceNames:
      - controlplane-operator-leader-election
    verbs:
      - get
      - update
      - patch

---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: controlplane-operator-leader-election
  namespace: controlplane-system
  labels:
    app.kubernetes.io/name: controlplane-operator
    app.kubernetes.io/component: rbac
subjects:
  - kind: ServiceAccount
    name: controlplane-operator
    namespace: controlplane-system
roleRef:
  kind: Role
  name: controlplane-operator-leader-election
  apiGroup: rbac.authorization.k8s.io

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: controlplane-operator
  labels:
    app.kubernetes.io/name: controlplane-operator
    app.kubernetes.io/component: rbac
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: controlplane-operator
subjects:
  - kind: ServiceAccount
    name: controlplane-operator
    namespace: controlplane-system

---
// === ARCHIVO: config/samples/controlplane_v1alpha1_controlplane.yaml ===
apiVersion: controlplane.example.com/v1alpha1
kind: ControlPlane
metadata:
  name: multi-region-controlplane
  namespace: controlplane-system
  labels:
    app.kubernetes.io/name: controlplane-operator
    app.kubernetes.io/managed-by: controlplane-operator
spec:
  argocd:
    namespace: argocd
    repository: https://github.com/example/argocd-config
    branch: main
    projectName: multi-region
    serverUrl: https://argocd.example.com
    adminEnabled: true
    syncTimeout: 300
    retryLimit: 3
    controllerLogLevel: info
    applicationSetNamespace: argocd
  crossplane:
    providerName: aws
    providerConfig: aws-provider
    compositionNamespace: crossplane-system
    enableAutoComposition: true
    serviceAccountName: crossplane-provider
    runtimeNamespace: crossplane-system
  vault:
    address: https://vault.example.com:8200
    pathPrefix: secret/data/controlplane
    authMethod: kubernetes
    role: controlplane-operator
    k8sAuthPath: auth/kubernetes
    secretEngine: kv-v2
    pkiRole: controlplane-ca
    skipTLSVerify: false
  aws:
    region: us-east-1
    s3Bucket: controlplane-state-us-east-1
    dynamoTable: controlplane-replication
    kmsKeyId: alias/controlplane-master
    enablePublicAccess: false
    vpcId: vpc-0123456789abcdef0
    subnetIds:
      - subnet-0123456789abcdef0
      - subnet-0123456789abcdef1
    securityGroupIds:
      - sg-0123456789abcdef0
    iamRoleArn: arn:aws:iam::123456789012:role/ControlPlaneOperator
    tags:
      Environment: production
      ManagedBy: controlplane-operator
      Project: multi-region-platform
  regions:
    - name: us-east-1
      priority: 1
      isPrimary: true
      clusterName: us-east-1-cluster
      clusterEndpoint: https://kube-us-east-1.example.com:6443
      clusterRegion: us-east-1
      vaultAddress: https://vault-us-east-1.example.com:8200
      enabledFeatures:
        - gitops
        - provisioning
        - monitoring
        - rbac-federation
      healthCheckInterval: 30
      failoverWeight: 100
    - name: us-west-2
      priority: 2
      isPrimary: false
      clusterName: us-west-2-cluster
      clusterEndpoint: https://kube-us-west-2.example.com:6443
      clusterRegion: us-west-2
      vaultAddress: https://vault-us-west-2.example.com:8200
      enabledFeatures:
        - gitops
        - provisioning
        - monitoring
      healthCheckInterval: 30
      failoverWeight: 80
    - name: eu-west-1
      priority: 3
      isPrimary: false
      clusterName: eu-west-1-cluster
      clusterEndpoint: https://kube-eu-west-1.example.com:6443
      clusterRegion: eu-west-1
      vaultAddress: https://vault-eu-west-1.example.com:8200
      enabledFeatures:
        - gitops
        - monitoring
      healthCheckInterval: 30
      failoverWeight: 60
    - name: ap-southeast-1
      priority: 4
      isPrimary: false
      clusterName: ap-southeast-1-cluster
      clusterEndpoint: https://kube-ap-southeast-1.example.com:6443
      clusterRegion: ap-southeast-1
      vaultAddress: https://vault-ap-southeast-1.example.com:8200
      enabledFeatures:
        - gitops
        - monitoring
      healthCheckInterval: 30
      failoverWeight: 50
  replication:
    enabled: true
    strategy: quorum
    quorumSize: 3
    syncInterval: 5
    heartbeatInterval: 1
    consensusTimeout: 10
    conflictResolution: primary-wins
    replicationNamespace: crossplane-replication
    etcdEndpoints:
      - https://etcd-0.example.com:2379
      - https://etcd-1.example.com:2379
      - https://etcd-2.example.com:2379
  failover:
    enabled: true
    autoFailover: true
    healthCheckInterval: 15
    failureThreshold: 3
    recoveryThreshold: 2
    postFailoverScript: "/scripts/post-failover.sh"
    failoverTimeout: 300
    rollbackEnabled: true
    maxFailoverDuration: 600
    healthCheckEndpoints:
      - https://kube-us-east-1.example.com:6443/healthz
      - https://kube-us-west-2.example.com:6443/healthz
  rbac:
    federated: true
    clusterRoles:
      - admin
      - edit
      - view
      - custom-cluster-reader
    namespaceRoles:
      - admin
      - edit
      - view
      - custom-ns-deployer
    syncInterval: 60
    defaultRole: view
    adminUsers:
      - admin@example.com
      - platform-team@example.com
  monitoring:
    enabled: true
    prometheusUrl: https://prometheus.example.com
    grafanaUrl: https://grafana.example.com
    alertManagerUrl: https://alertmanager.example.com
    metricsPort: 8080
    logLevel: info
    traceEndpoint: https://tempo.example.com:4317
    retentionDays: 30
  topology:
    clusterCount: 12
    regionCount: 4
    meshMode: star
    serviceMesh: istio
    ingressClass: nginx
    cniPlugin: cilium
    ccmProvider: aws

// === ARCHIVO: test/e2e/controlplane_test.go ===
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

// === ARCHIVO: test/e2e/failover_test.go ===
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

// === ARCHIVO: test/e2e/replication_test.go ===
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
```
