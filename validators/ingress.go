package validators

import "fmt"

type IngressManifest struct {
	APIVersion string          `json:"apiVersion"`
	Kind       string          `json:"kind"`
	Metadata   IngressMetadata `json:"metadata"`
	Spec       IngressSpec     `json:"spec"`
}

type IngressMetadata struct {
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type IngressSpec struct {
	Rules []Rule `json:"rules,omitempty"`
}

type Rule struct {
	HTTP PathMap `json:"http,omitempty"`
}

type PathMap struct {
	Paths []Path `json:"paths,omitempty"`
}

type Path struct {
	Name     string     `json:"path,omitempty"`
	PathType string     `json:"pathType,omitempty"`
	Backend  BackendMap `json:"backend,omitempty"`
}

type BackendMap struct {
	Service ServiceMap `json:"service,omitempty"`
}

type ServiceMap struct {
	Name string  `json:"name,omitempty"`
	Port PortMap `json:"port,omitempty"`
}

type PortMap struct {
	Name   string `json:"name,omitempty"`
	Number int    `json:"number,omitempty"`
}

func (m *IngressManifest) PrintTopLevelManifest() {
	// TODO: better use of formatters
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", m.APIVersion)
	kind := fmt.Sprintf("           Kind: %s\n", m.Kind)
	metadata := fmt.Sprintf("       Metadata: %+v\n", m.Metadata)
	fmt.Println(apiVersion, kind, metadata)
}

func (m *IngressManifest) PrintSpec() {
	// TODO: better use of formatters
	spec := m.Spec
	for _, rule := range spec.Rules {
		for _, path := range rule.HTTP.Paths {
			fmt.Println("            Rule:")
			fmt.Printf("                     Path: %s\n", path.Name)
			fmt.Printf("                 PathType: %s\n", path.PathType)
			fmt.Printf("                  Backend: %+v\n\n", path.Backend)
		}
	}
}
