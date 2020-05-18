package mock

import v1 "k8s.io/api/core/v1"

type Mock struct {
	ServiceDefinition ServiceDefinition `json:"serviceDefinition"`
	Watch             bool              `json:"watch"`
	Selector          IngressSelector   `json:"selector"`
	Host              string            `json:"host"`
	Definition        string            `json:"definition"`
	Name              string            `json:"name"`
}

type ServiceDefinition struct {
	Port        int            `json:"port"`
	ServiceType v1.ServiceType `json:"serviceType"`
}

type IngressSelector struct {
	Ingress   string `json:"ingress"`
	Namespace string `json:"namespace"`
}
