# hcloud-k8s-floatingip

This tool assigns a [Hetzner Cloud](https://hetzner.cloud) floating ip to the
server it is being executed on. The idea is to run this tool as a service in
a Kubernetes cluster. This ensures that the floating ip is always assigned to
a healthy node.

**Note:** This assumes that you have named the Kubernetes nodes the same as
your servers on Hetzner Cloud.

## Usage

```
docker run -it --rm \
       -e HCLOUD_TOKEN= \
       -e FLOATING_IP_ID= \
       -e THIS_SERVER_NAME= \
       sh4rk/hcloud-k8s-floatingip
```

See [`k8s.yaml`](k8s.yaml) for a Kubernetes service description.
