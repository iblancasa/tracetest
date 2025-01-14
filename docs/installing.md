# Installation

During the setup, we'll deploy Tracetest and Postgres with Helm.

For the architectural overview of the components, please check the [Architecture](architecture.md) page.

## **Prerequsities**

### **Installation Requirements**

Tools needed for the installation:

- [Helm v3](https://helm.sh/docs/intro/install/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)

## **Installation**

### Install script

We provide a simple install script that can install all required components:

```
curl -L https://raw.githubusercontent.com/kubeshop/tracetest/main/setup.sh | bash -s
```

This command will install Tracetest using the default settings. You can configure the following options:

| Option                   | description                                  | Default value      |
| ------------------------ | -------------------------------------------- | ------------------ |
| --help                   | show help message                            | n/a                |
| --namespace              | target installation k8s namespace            | tracetest          |
| --trace-backend          | trace backend (jaeger or tempo)              | jaeger             |
| --trace-backend-endpoint | trace backend endpoint                       | jaeger-query:16685 |
| --skip-collector         | if set, don't install the otel-collector     | n/a                |
| --skip-pma               | if set, don't install the sample application | n/a                |
| --skip-backend           | if set, don't install the jaeger backend     | n/a                |

Example with custom options:

```
curl -L https://raw.githubusercontent.com/kubeshop/tracetest/main/setup.sh | bash -s -- --skip-pma --namespace my-custom-namespace
```

### Using Helm

Container images are hosted on the Docker Hub [Tracetest repository](https://hub.docker.com/r/kubeshop/tracetest).

Tracetest currently supports two traces backend: Jaeger and Grafana Tempo.

#### **Jaeger**

Tracetest uses [Jaeger Query Service `16685` port](https://www.jaegertracing.io/docs/1.32/deployment/#query-service--ui) to find Traces using gRPC protocol.

The commands below will install Tracetest connecting to the Jaeger tracing backend on `jaeger-query:16685`.

```sh
# Install Kubeshop Helm repo and update it
helm repo add kubeshop https://kubeshop.github.io/helm-charts
helm repo update

helm install tracetest kubeshop/tracetest \
  --set telemetry.dataStores.jaeger.jaeger.endpoint="jaeger-query:16685" \ # update this value to point to your jaeger install
  --set telemetry.exporters.collector.exporter.collector.endpoint="otel-collector:4317" \ # update this value to point to your collector install
  --set server.telemetry.dataStore="jaeger"
```

#### **Grafana Tempo**

Tracetest uses [Grafana Tempo's Server's `9095` port](https://grafana.com/docs/tempo/latest/configuration/#server) to find Traces using gRPC protocol.

The commands below will install the Tracetest application connecting to the Grafana Tempo tracing backend on `grafana-tempo:9095`:

```sh
# Install Kubeshop Helm repo and update it
helm repo add kubeshop https://kubeshop.github.io/helm-charts
helm repo update

helm install tracetest kubeshop/tracetest \
  --set telemetry.dataStores.tempo.tempo.endpoint="grafana-tempo:9095" \ # update this value to point to your tempo install
  --set telemetry.exporters.collector.exporter.collector.endpoint="otel-collector:4317" \ # update this value to point to your collector install
  --set server.telemetry.dataStore="tempo"
```

### **Have a different backend trace data store?**

[Tell us](https://github.com/kubeshop/tracetest/issues/new?assignees=&labels=&template=feature_request.md&title=) which one you have and we will see if we can add support for it!

## **Uninstallation**

The following command will uninstall Tracetest with Postgres:

```sh
helm delete tracetest
```

## CLI Instalation
Every time we release a new version of Tracetest, we generate binaries for Linux, MacOS, and Windows. Supporting both amd64, and ARM64 architectures. You can find the latest version [here](https://github.com/kubeshop/tracetest/releases/latest).

### Linux

```sh
curl -L https://raw.githubusercontent.com/kubeshop/tracetest/main/install-cli.sh | bash
```

### MacOS

```sh
curl -L https://raw.githubusercontent.com/kubeshop/tracetest/main/install-cli.sh | sh
```

### Windows
Download one of the files from the latest tag, extract to your machine, and then [add the tracetest binary to your PATH variable](https://stackoverflow.com/a/41895179)
