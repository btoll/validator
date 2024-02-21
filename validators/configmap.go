package validators

import "fmt"

type ConfigMapManifest struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Data       Data     `json:"data"`
}

func (m ConfigMapManifest) Print() {
	// TODO: better use of formatters
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", m.APIVersion)
	kind := fmt.Sprintf("           Kind: %s\n", m.Kind)
	name := fmt.Sprintf("           Name: %+v\n", m.Metadata.Name)
	namespace := fmt.Sprintf("      Namespace: %+v\n", m.Metadata.Namespace)
	labels := fmt.Sprintf("         Labels:")
	fmt.Println(apiVersion, kind, name, namespace, labels)
	for k, v := range m.Metadata.Labels {
		fmt.Printf("                  %s: %s\n", k, v)
	}
	data := fmt.Sprintf("            Data:")
	fmt.Println(data)
	for k, v := range m.Data {
		fmt.Printf("                  %s: %s\n", k, v)
	}
}
