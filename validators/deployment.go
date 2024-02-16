package validators

import (
	"fmt"
)

type DeploymentManifest struct {
	APIVersion string         `json:"apiVersion"`
	Kind       string         `json:"kind"`
	Metadata   Metadata       `json:"metadata"`
	Spec       DeploymentSpec `json:"spec"`
}

type DeploymentSpec struct {
	Replicas int         `json:"replicas,omitempty"`
	Selector MatchLabels `json:"selector,omitempty"`
	Template PodSpec     `json:"template,omitempty"`
}

type MatchLabels struct {
	MatchLabel Labels `json:"matchLabels,omitempty"`
}

type PodSpec struct {
	Metadata Metadata   `json:"metadata,omitempty"`
	Spec     Containers `json:"spec,omitempty"`
}

type Containers struct {
	Containers []Container `json:"containers,omitempty"`
}

type Container struct {
	Name            string    `json:"name,omitempty"`
	Image           string    `json:"image,omitempty"`
	ImagePullPolicy string    `json:"imagePullPolicy,omitempty"`
	EnvVars         []EnvVar  `json:"env,omitempty"`
	EnvFrom         []EnvFrom `json:"envFrom,omitempty"`
	Ports           []Port    `json:"ports,omitempty"`
	Resources       Resources `json:"resources,omitempty"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnvFrom struct {
	ConfigMapRef Labels `json:"configMapRef,omitempty"`
}

type Port struct {
	ContainerPort int `json:"container_port,omitempty"`
}

type Resources struct {
	Limits   Resource `json:"limits,omitempty"`
	Requests Resource `json:"requests,omitempty"`
}

type Resource struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

func (m DeploymentManifest) PrintTopLevelManifest() {
	//	fmt.Println(m)
	// TODO: better use of formatters
	apiVersion := fmt.Sprintf("      APIVersion: %s\n", m.APIVersion)
	kind := fmt.Sprintf("           Kind: %s\n", m.Kind)
	metadata := fmt.Sprintf("       Metadata: %+v\n", m.Metadata)
	fmt.Println(apiVersion, kind, metadata)
}

func (m DeploymentManifest) PrintSpec() {
	// TODO: better use of formatters
	spec := m.Spec
	replicas := fmt.Sprintf("        Replicas: %d\n", spec.Replicas)
	selector := fmt.Sprintf("       Selector: %+v\n", spec.Selector)
	fmt.Println(replicas, selector)

	container := spec.Template.Spec.Containers[0]
	name := fmt.Sprintf("            Name: %s\n", container.Name)
	image := fmt.Sprintf("          Image: %s\n", container.Image)
	imagePullPolicy := fmt.Sprintf("ImagePullPolicy: %s\n", container.ImagePullPolicy)
	envVars := fmt.Sprintf("        EnvVars: %+v\n", container.EnvVars)
	envFrom := fmt.Sprintf("        EnvFrom: %+v\n", container.EnvFrom)
	ports := fmt.Sprintf("          Ports: %+v\n", container.Ports)
	resources := fmt.Sprintf("      Resources: %+v\n", container.Resources)
	fmt.Println(name, image, imagePullPolicy, envVars, envFrom, ports, resources)
}
