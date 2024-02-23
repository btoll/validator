package validators

import (
	"fmt"

	"github.com/btoll/validator/lib"
)

type IngressManifest struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   `json:"metadata"`
	Spec       IngressSpec `json:"spec"`
}

type IngressSpec struct {
	Rules []IngressRule `json:"rules,omitempty"`
}

type IngressRule struct {
	IngressRuleValue `json:",inline,omitempty"`
}

type IngressRuleValue struct {
	HTTP HTTPIngressRuleValue `json:"http,omitempty"`
}

type HTTPIngressRuleValue struct {
	Paths []HTTPIngressPath `json:"paths,omitempty"`
}

type HTTPIngressPath struct {
	Path     string         `json:"path,omitempty"`
	PathType string         `json:"pathType,omitempty"`
	Backend  IngressBackend `json:"backend,omitempty"`
}

type IngressBackend struct {
	Service IngressServiceBackend `json:"service,omitempty"`
}

type IngressServiceBackend struct {
	Name string             `json:"name,omitempty"`
	Port ServiceBackendPort `json:"port,omitempty"`
}

type ServiceBackendPort struct {
	Name   string `json:"name,omitempty"`
	Number int    `json:"number,omitempty"`
}

func (m IngressManifest) Write() {
	properServiceName := lib.GetProperServiceName(m.Name)
	m.Name = properServiceName

	dir := fmt.Sprintf("build/%s/ingress", properServiceName)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	WriteTemplate(fmt.Sprintf("%s/local", dir), "ingress.tpl", m)

	ingress, err := GetIngressClient(properServiceName)
	if err != nil {
		fmt.Println("err", err)
	}
	// This values are empty in the returned struct.
	ingress.APIVersion = "networking.k8s.io/v1"
	ingress.Kind = "Ingress"
	WriteTemplate(fmt.Sprintf("%s/remote", dir), "ingress.tpl", ingress)
}
