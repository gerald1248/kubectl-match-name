package main

// KubeconfigObject captures the user's ~/.kube/config file
type KubeconfigObject struct {
	APIVersion     string          `json:"apiVersion"`
	Kind           string          `json:"kind"`
	CurrentContext string          `json:"current-context"`
	Contexts       []ContextObject `json:"contexts"`
	Users          []interface{}   `json:"users"`
	Clusters       []interface{}   `json:"clusters"`
}

// ContextObject represents a single context item
type ContextObject struct {
	Name    string              `json:"name"`
	Context NestedContextObject `json:"context"`
}

// NestedContextObject contains the namespace
type NestedContextObject struct {
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
	User      string `json:"user"`
}

// Pod is a minimal type for a pod manifest
type Pod struct {
	APIVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
}

// Metadata represents the container for the name
type Metadata struct {
	Name string `json:"name"`
}
