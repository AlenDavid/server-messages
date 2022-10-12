# -*- mode: Python -*-

docker_build('alen/server-messages', '.', dockerfile='deployments/Dockerfile')
k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('server-messages-go', port_forwards=5050)
