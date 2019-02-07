package main

import (
	"fmt"
	"os"
	"errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

	switch kind {
	case "po", "pod", "pods":
		names = getPods(clientset, namespace)
	case "cm", "configmap", "configmaps" :
		names = getConfigMaps(clientset, namespace)
	case "svc", "service":
		names = getServices(clientset, namespace)
	default:
		return nil, errors.New(fmt.Sprintf("unsupported object kind: %s", kind))
	}

	return names, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
