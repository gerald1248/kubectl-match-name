package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta1"
)

func getDeployments(client v1beta1.AppsV1beta1Interface, namespace string) ([]string, error) {
	objects, err := client.Deployments(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var names []string
	for _, object := range objects.Items {
		names = append(names, object.ObjectMeta.Name)
	}

	return names, nil
}
