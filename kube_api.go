package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strings"
)

func getObjects(kubeconfig string, namespace string, kind string, allNamespaces bool) ([]string, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	// fetch current namespace
	if len(namespace) == 0 {
		namespace, err = extractCurrentNamespaceFromFile(kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	// override: all namespaces
	if allNamespaces {
		namespace = ""
	}

	var names []string

	kind = strings.ToLower(kind)

	switch kind {
	case "cm", "configmap", "configmaps":
		names, err = getConfigMaps(clientset.CoreV1(), namespace)
	case "deploy", "deployment", "deployments":
		names, err = getDeployments(clientset.AppsV1beta1(), namespace)
	case "po", "pod", "pods":
		names, err = getPods(clientset.CoreV1(), namespace)
	case "secret", "secrets":
		names, err = getSecrets(clientset.CoreV1(), namespace)
	case "svc", "service", "services":
		names, err = getServices(clientset.CoreV1(), namespace)
	default:
		return nil, fmt.Errorf("unsupported object kind: %s", kind)
	}

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return names, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
