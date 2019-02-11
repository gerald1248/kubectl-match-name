package main

import (
	"testing"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func pod(namespace, image string) *v1.Pod {
	return &v1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: namespace}, Spec: v1.PodSpec{Containers: []v1.Container{{Image: image}}}}
}

func TestGetPods(t *testing.T) {
	objs := []runtime.Object{pod("default", "pod-a")}
	client := fake.NewSimpleClientset(objs...)
	names, err := getPods(client.CoreV1(), "default")
	expected := 1

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(names) != expected {
		t.Errorf("Expected exactly %d ConfigMaps", expected)
	}
}
