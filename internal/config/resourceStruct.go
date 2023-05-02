package config

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"
)

type Pod struct {
	Name            string
	Status          string
	CreatedAt       string
	UniqueID        string
	NodeName        string
	IP              string
	ContainersCount int
	ContainersInfo  []Container
	Labels          map[string]string
}

type Container struct {
	Name            string
	Image           string
	ImagePullPolicy string
	Container       int
	Port            []v1.ContainerPort
}

type Deployment struct {
	Name      string
	Status    string
	CreatedAt string
	UniqueID  string
	Labels    map[string]string
}

type Service struct {
	Name      string
	CreatedAt string
	UniqueID  string
	Labels    map[string]string
	Spec      interface{}
}

type Namespace struct {
	Name      string
	CreatedAt string
	UniqueID  string
}

type Configmap struct {
	Name string
}

type Daemonset struct {
	Name      string
	CreatedAt string
	UniqueID  string
	Labels    map[string]string
}

type Event struct {
	Name       string
	Type       string
	ObjectName string
	CreatedAt  string
	UniqueID   string
}

type Replicationcontroller struct {
	Name      string
	CreatedAt string
	UniqueID  string
	Labels    map[string]string
}

type Replicaset struct {
	Name           string
	CreatedAt      string
	UniqueID       string
	ReplicasetSpec interface{}
	Labels         map[string]string
}

type Secret struct {
	Name      string
	SecretMap map[string]string
	Type      string
	CreatedAt string
	UniqueID  string
}

type WideResult struct {
	ApiVersion string
	Items      interface{}
	Kind       interface{}
	Metadata   interface{}
}

type Objmap struct {
	Objectkind string
	Objectname string
}

type Time struct {
	Time *metav1.Time
}
