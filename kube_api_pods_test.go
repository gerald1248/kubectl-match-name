package main

import (
	"testing"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func pod(namespace, image string) *v1.Pod {
	return &v1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: namespace}, Spec: v1.PodSpec{Containers: []v1.Container{{Image: image}}}}
}

func TestGetPods(t *testing.T) {
	var tests = []struct {
		description string
		namespace   string
		expected    int
		objs        []runtime.Object
	}{
		{"no_pods", "default", 0, nil},
		{"one_pod", "default", 1, []runtime.Object{pod("default", "image")}},
		{"wrong_namespace", "default", 0, []runtime.Object{pod("wrong", "image")}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := fake.NewSimpleClientset(test.objs...)
			names, err := getPods(client.CoreV1(), test.namespace)

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if len(names) != test.expected {
				t.Errorf("Expected exactly %d pods", test.expected)
			}
		})
	}
}
