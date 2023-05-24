cluster_name := "datadog-provider"

[private]
default:
  @just -l

# Bring up the Kind cluster
cluster-up: && load-crds install-crossplane
  kind create cluster --name {{cluster_name}} --image kindest/node:v1.26.2
  sleep 5
  kubectl wait --namespace kube-system --for=condition=ready pod --selector="tier=control-plane" --timeout=180s
  kubectl create namespace crossplane-system
  kubectl --namespace crossplane-system apply -f ./datadog-secret.yaml

# Bring down the Kind cluster
cluster-down:
  kind delete cluster --name {{cluster_name}}
  -rm ./kubeconfig

# Load cluster CRDs
load-crds:
  kubectl create -f ./package/crds
  sleep 3
  kubectl apply -f ./examples/providerconfig/providerconfig.yaml

# Remove cluster CRDs
rm-crds:
  -kubectl patch providerconfig default -p '{"metadata":{"finalizers": []}}' --type=merge
  -kubectl delete -f ./examples/providerconfig/providerconfig.yaml
  -kubectl delete -f ./package/crds

# Reload cluster CRDs
reload-crds: rm-crds && load-crds

# Clean build cache
clean-files:
  rm -rvf .cache .work _output

# Install crossplane using Helm
install-crossplane:
  helm repo add crossplane-stable https://charts.crossplane.io/stable
  helm repo update
  helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
