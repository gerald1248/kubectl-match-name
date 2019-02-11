package main

import (
	"testing"
)

func TestGetObjects(t *testing.T) {
	sampleConfig := []byte(`
	apiVersion: v1
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

	var tests = []struct {
		description   string
		config        string
		namespace     string
		kind          string
		allNamespaces bool
	}{
		{"no_config", "{}", "default", "pod", false},
		{"unsupported_kind", string(sampleConfig), "default", "invalid", false},
	}

	//getObjects(kubeconfig string, namespace string, kind string, allNamespaces bool)
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := getObjects(test.config, test.namespace, test.kind, test.allNamespaces)
			if err == nil {
				t.Errorf("Function call with invalid parameters must fail")
			}
		})
	}
}

func TestHomeDir(t *testing.T) {
	homeDir := homeDir()
	if len(homeDir) == 0 {
		t.Errorf("home directory must not be empty")
	}
}
