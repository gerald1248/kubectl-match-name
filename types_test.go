package main

import (
	"encoding/json"
	"testing"
)

func TestKubeConfigObject(t *testing.T) {
	bytes := []byte(`apiVersion: v1
clusters:
- cluster:
    certificate-authority: /Users/gerald/.minikube/ca.crt
    server: https://192.168.99.100:8443
  name: 192-168-99-100:8443
- cluster:
    certificate-authority: /Users/gerald/.minikube/ca.crt
    server: https://192.168.99.100:8443
  name: minikube
contexts:
- context:
    cluster: 192-168-99-100:8443
    namespace: default
    user: /192-168-99-100:8443
  name: default/192-168-99-100:8443/
- context:
    cluster: minikube
    namespace: kube-system
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: /192-168-99-100:8443
  user:
    client-certificate: /Users/gerald/.minikube/client.crt
    client-key: /Users/gerald/.minikube/client.key
- name: minikube
  user:
    client-certificate: /Users/gerald/.minikube/client.crt
    client-key: /Users/gerald/.minikube/client.key
`)
	var kubeconfig KubeconfigObject
	if err := json.Unmarshal(bytes, &kubeconfig); err == nil {
		t.Errorf("KubeConfig object must be compatible with sample YAML input: %s", err)
	}
}
