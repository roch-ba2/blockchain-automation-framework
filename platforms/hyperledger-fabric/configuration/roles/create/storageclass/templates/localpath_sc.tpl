apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: {{ sc_name }}
provisioner: rancher.io/local-path
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer