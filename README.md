# Provider DataDog

`provider-datadog` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
DataDog API.

This provider is using the DataDog Terraform provider version `3.25.0` and has support for the following resources.

 - Dashboard
 - Dashboard JSON

### References
- DataDog Terraform provider <https://registry.terraform.io/providers/DataDog/datadog/latest/docs>
- DataDog API Documentation <https://docs.datadoghq.com/api/latest/>

## Installation

This provider is created using Upjet. The idea behind Upjet is to use it with Upbound Market, where you upload the package with their CLI tool.
At the time of writing, the Upbound Market is brand new and still a bit wonky. Their CLI isn't much different and there's limited documentation.
That is why I made this provider like the original Terrajet code generation tool. When running `make build` it will create two containers which you will have
to push to a container registry and reference in the `Provider` and `ControllerConfig`.

Prerequisites:

 - Go 1.19
 - Docker
 - Make

### Build the images
Clone this repository and checkout the desired version tag.

Install the submodules.
```
make submodules
```

Export some varibles to set the registry URL and image tag.
```
export BUILD_REGISTRY="some/registry-url"
export VERSION=v0.0.1
```

Start the build.
```
make build
```

Above build command will create two container images, push them to your registry.
```
REPOSITORY                                          TAG       IMAGE ID       SIZE
some/registry-url/provider-datadog-amd64           latest    092d173576b7   153MB
some/registry-url/provider-datadog-amd64           v0.0.1    092d173576b7   153MB
some/registry-url/provider-datadog-package-amd64   latest    f87030437c60   264kB
some/registry-url/provider-datadog-package-amd64   v0.0.1    f87030437c60   264kB
```

### Deploy to Kubernetes

Go to your DataDog Management page and create an API and an App key.

Create below secret with the DataDog URL, API and APP Key.

```
cat << EOF | kubectl apply -f - 
apiVersion: v1
kind: Secret
metadata:
  name: datadog-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<api-key-here>",
      "app_key": "<app-key-here>",
      "api_url": "https://api.datadoghq.eu"
    }
EOF
```

Apply below `Provider` and `ControllerConfig`. Replace the registry URL in both objects.
This will install the controller deployment and all the CRDs for your custom provider.

```
cat << EOF | kubectl apply -f - 
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-datadog
spec:
  package: some/registry-url/provider-datadog-package-amd64:v0.0.1
  controllerConfigRef:
    name: datadog-config
---
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: datadog-config
spec:
  image: some/registry-url/provider-datadog-amd64:v0.0.1
EOF
```

Apply a `ProviderConfig` that references the created secret with DataDog credentials. 
```
cat << EOF | kubectl apply -f - 
apiVersion: datadog.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: datadog-creds
      namespace: crossplane-system
      key: credentials
EOF
```

You can now start deploying DataDog resources.
You can find examples in the `examples` folder.

## Start Developing

Prerequisites:
 - Go 1.19
 - Docker
 - Kubectl
 - Make
 - [Just](https://github.com/casey/just)
 - [Kind](https://kind.sigs.k8s.io/)
 - [Helm](https://helm.sh/docs/intro/install/)

Setup the required prerequisites.

Create a secrets yaml file with below content and save it as `datadog-secret.yaml` in the root directory and change the appropriate values.
It will be added to the Kind cluster when you provision it. 

```
apiVersion: v1
kind: Secret
metadata:
  name: datadog-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<api-key-here>",
      "app_key": "<app-key-here>",
      "api_url": "https://api.datadoghq.eu"
    }
```

Export the Kind `kubeconfig` file location to use with `kubectl`.
```
export KUBECONFIG=./kubeconfig
```

Use `just` to bootstrap a Kind cluster.
Below command will:
- Create a Kind cluster.
- Create a `crossplane-system` namespace.
- Load the appropriate CRDs and ProviderConfig
- Install Crossplane using Helm

```
just cluster-up
```

Install the needed submodules.
```
make submodules
```

Generate the code.
```
make generate
```

Run the provider.
```
make run
```

You can now add DataDog resources to your Kind cluster and the provider should now create them in DataDog.

If you have created a new resources and you want to redeploy the CRDs to the cluster, run below command.
```
just reload-crds
```

When you're done, take down the Kind cluster with the following command.
```
just cluster-down
```

To test your newly added code run the `reviewable` command.
```
make reviewable
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/ok-amba/provider-datadog/issues).
