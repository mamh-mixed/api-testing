+++
title = "Install Grafana"
weight = -99
+++

```shell
helm repo add grafana https://grafana.github.io/helm-charts

helm install grafana grafana/grafana \
    --set service.type=NodePort \
    --set service.nodePort=30004
```
