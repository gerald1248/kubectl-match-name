package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func extractCurrentNamespace(byteArray []byte) (string, error) {
	err := preflightAsset(&byteArray)
	if err != nil {
		return "", fmt.Errorf("kubeconfig failed preflight check")
	}

	var kubeconfig KubeconfigObject
	if err = json.Unmarshal(byteArray, &kubeconfig); err != nil {
		return "", fmt.Errorf("can't unmarshal data: %v", err)
	}

	currentContext := kubeconfig.CurrentContext

	for _, context := range kubeconfig.Contexts {
		if context.Name == currentContext {
			return context.Context.Namespace, nil
		}
	}
	return "default", nil
}

func extractCurrentNamespaceFromFile(path string) (string, error) {
	byteArray, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("can't read %s: %v", path, err)
	}
	return extractCurrentNamespace(byteArray)
}
