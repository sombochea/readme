# Storage Class (Providers)

-   Ceph RBD

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
    name: fast
provisioner: kubernetes.io/rbd
parameters:
    monitors: 10.16.153.105:6789
    adminId: kube
    adminSecretName: ceph-secret
    adminSecretNamespace: kube-system
    pool: kube
    userId: kube
    userSecretName: ceph-secret-user
    userSecretNamespace: default
    fsType: ext4
    imageFormat: '2'
    imageFeatures: 'layering'
```
