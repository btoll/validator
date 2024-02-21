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
	Selector Data
	Type     string
}

type ServicePort struct {
	Port       int
	TargetPort int
	Protocol   string
}

func (m ServiceManifest) Print() {
	// TODO: better use of formatters
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", m.APIVersion)
	kind := fmt.Sprintf("           Kind: %s\n", m.Kind)
	//	metadata := fmt.Sprintf("       Metadata: %+v\n", m.Metadata)
	name := fmt.Sprintf("           Name: %+v\n", m.Metadata.Name)
	namespace := fmt.Sprintf("      Namespace: %+v\n", m.Metadata.Namespace)
	labels := fmt.Sprintf("         Labels:")
	fmt.Println(apiVersion, kind, name, namespace, labels)
	for k, v := range m.Metadata.Labels {
		fmt.Printf("                  %s: %s\n", k, v)
	}

	spec := m.Spec
	ports := fmt.Sprintf("           Ports: %+v\n", spec.Ports)
	selector := fmt.Sprintf("       Selector: %+v\n", spec.Selector)
	_type := fmt.Sprintf("           Type: %s\n", spec.Type)
	fmt.Println(ports, selector, _type)

	PrintService(m.Metadata.Name)
}
