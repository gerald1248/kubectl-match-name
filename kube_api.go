package main

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getObjects(kubeconfig string, namespace string, kind string, allNamespaces bool) []string {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// fetch current namespace
	if len(namespace) == 0 {
		namespace, err = extractCurrentNamespaceFromFile(kubeconfig)
		if err != nil {
			panic(err.Error())
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
	}

	return names
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
