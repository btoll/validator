package validators

import "fmt"

type ServiceManifest struct {
	APIVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   Metadata    `json:"metadata"`
	Spec       ServiceSpec `json:"spec"`
}

type ServiceSpec struct {
	Ports    []ServicePort
	Selector ServiceSelector
	Type     string
}

type ServicePort struct {
	Port       int
	TargetPort int
	Protocol   string
}

type ServiceSelector struct {
}

func (m ServiceManifest) PrintTopLevelManifest() {
	// TODO: better use of formatters
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", m.APIVersion)
	kind := fmt.Sprintf("           Kind: %s\n", m.Kind)
	metadata := fmt.Sprintf("       Metadata: %+v\n", m.Metadata)
	fmt.Println(apiVersion, kind, metadata)
}

func (m ServiceManifest) PrintSpec() {
	// TODO: better use of formatters
	spec := m.Spec
	ports := fmt.Sprintf("           Ports: %+v\n", spec.Ports)
	selector := fmt.Sprintf("       Selector: %s\n", spec.Selector)
	_type := fmt.Sprintf("           Type: %s\n", spec.Type)
	fmt.Println(ports, selector, _type)
}
