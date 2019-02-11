package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func getSecrets(client v1.CoreV1Interface, namespace string) ([]string, error) {
	objects, err := client.Secrets(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var names []string
	for _, object := range objects.Items {
		names = append(names, object.ObjectMeta.Name)
	}

	return names, nil
}
