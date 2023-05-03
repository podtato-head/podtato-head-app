# The PodTatoHead Server
Podtato-head is a cloud-native application built to colorfully demonstrate delivery scenarios using different tools and services. This repository contains the PodTatoHead Application, it's manifest and a helm chart for easy deployment.

## Installing

### Kubernetes Manifest
The Kubernetes manifest is located in the `deploy` folder. It can be deployed using `kubectl` or `kustomize`.

<!---x-release-please-start-version-->
> ```kubectl apply -f https://github.com/podtato-head/podtato-head-app/releases/download/v0.3.3/manifest.yaml```
<!---x-release-please-end-->

### Helm Chart
The Helm chart is located in the `charts` folder. It can be deployed using `helm`.

> ```helm upgrade --install --wait podtato --namespace default oci://ghcr.io/podtato-head/charts/podtato-head```




