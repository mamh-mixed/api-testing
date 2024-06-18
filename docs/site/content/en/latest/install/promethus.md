+++
title = "Install Prometheus"
weight = -99
+++

```shell
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

helm install prometheus prometheus-community/prometheus \
    --set server.service.type=NodePort \
    --set server.service.nodePort=30003

helm install prometheus-operator oci://registry-1.docker.io/bitnamicharts/kube-prometheus
```
