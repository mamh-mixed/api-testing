+++
title = "Install in K3D"
weight = -99
+++

```shell
k3d cluster create

k3d node edit k3d-k3s-default-serverlb \
    --port-add 30002:30002 \
    --port-add 30003:30003 \
    --port-add 30004:30004
```
