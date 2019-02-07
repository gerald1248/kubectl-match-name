package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getServices(clientset *kubernetes.Clientset, namespace string) []string {
	objects, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []string
	for _, object := range objects.Items {
		names = append(names, object.ObjectMeta.Name)
	}

	return names
}
