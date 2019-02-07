package main

import (
	"testing"
)

func TestExtractCurrentNamespace(t *testing.T) {
	kubeconfig := []byte(`apiVersion: v1
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
	namespace, err := extractCurrentNamespace(kubeconfig)

	if err != nil {
		t.Errorf("Must accept valid kubeconfig: %v", err)
		return
	}

	expected := "kube-system"
	if namespace != "kube-system" {
		t.Errorf("Expected namespace %s, got %s", expected, namespace)
	}
}
