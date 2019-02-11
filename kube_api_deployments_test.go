package main

import (
	"testing"

	"k8s.io/client-go/kubernetes/fake"
)

func TestGetDeployments(t *testing.T) {
	client := fake.NewSimpleClientset()
	names, err := getDeployments(client.AppsV1beta1(), "default")
	expected := 0

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(names) != expected {
		t.Errorf("Expected exactly %d ConfigMaps", expected)
	}
}
