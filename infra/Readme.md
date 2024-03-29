# infra

## cluster scalability

1. start from 1 node
2. setup control-plane with complete default config from kubeKey
3. add 2 more nodes(control-plane,worker,etcd) in yaml
4. `kk add nodes -f ./ai-server/infra/dev-cluster.yaml`
5. attention: kubeKey forbids etcd nodes to be deleted

## image registry

1. use tencent cloud personal plan for early development
2. later: [self-hosted harbor ha](https://github.com/kubesphere/kubekey/blob/master/docs/harbor-ha.md)

## user traffic scalability

1. reduce cloud module using like CLB/CVM(dedicated LB server)
2. expose 443 ports on each node ready for ingress
3. access any one of servers via nodePubIP:443
4. use dns loadbalancer to route traffic to any one of servers

## gpu

nvidia gpu-operator chart

```bash
helm install --wait --generate-name \
     -n nvidia-gpu-operator --create-namespace \
     nvidia/gpu-operator
```

## storage

1. dev: localPath provisioner
   ```bash
   kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/v0.0.26/deploy/local-path-storage.yaml
   ```
2. prod: openebs?

## middleware

1. mariadb
  ```bash
  helm upgrade -i -n mariadb --create-namespace mariadb oci://registry-1.docker.io/bitnamicharts/mariadb
  ```
