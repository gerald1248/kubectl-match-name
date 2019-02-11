package main

import (
	"testing"

	"k8s.io/client-go/kubernetes/fake"
)

func TestGetServices(t *testing.T) {
	client := fake.NewSimpleClientset()
	names, err := getServices(client.CoreV1(), "default")
	expected := 0

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(names) != expected {
		t.Errorf("Expected exactly %d Services", expected)
	}
}
