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
	dir := fmt.Sprintf("build/%s/ingress", m.Name)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	localFile := fmt.Sprintf("%s/local", dir)
	WriteTemplate(localFile, "ingress.tpl", m)

	ingress, err := GetIngressClient(m.Name)
	if err != nil {
		fmt.Println("err", err)
	}
	// This values are empty in the returned struct.
	ingress.APIVersion = "networking.k8s.io/v1"
	ingress.Kind = "Ingress"
	remoteFile := fmt.Sprintf("%s/remote", dir)
	WriteTemplate(remoteFile, "ingress.tpl", ingress)
	b, err := Validate(localFile, remoteFile)
	if err != nil {
		fmt.Println("err", err)
	}
	if b {
		err = RemoveDir(dir, fmt.Sprintf("build/%s", m.Name))
		if err != nil {
			fmt.Println(err)
		}
	}
}
