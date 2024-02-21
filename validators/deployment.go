package validators

import (
	"fmt"

	"github.com/btoll/validator/lib"
)

type DeploymentManifest struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   `json:"metadata"`
	Spec       DeploymentSpec `json:"spec"`
}

type DeploymentSpec struct {
	Replicas int           `json:"replicas,omitempty"`
	Selector LabelSelector `json:"selector,omitempty"`
	Template PodSpec       `json:"template,omitempty"`
}

type LabelSelector struct {
	MatchLabels Data `json:"matchLabels,omitempty"`
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
	ConfigMapRef Data `json:"configMapRef,omitempty"`
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

func (m DeploymentManifest) Write() {
	properServiceName := lib.GetProperServiceName(m.Name)
	m.Name = properServiceName

	dir := fmt.Sprintf("build/%s/deployment", properServiceName)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	f, err := lib.CreateBuildFile(fmt.Sprintf("%s/local", dir))
	if err != nil {
		fmt.Println("err", err)
	}
	err = tpl.ExecuteTemplate(f, "deployment.tpl", m)
	if err != nil {
		fmt.Println("err", err)
	}

	PrintDeployment(properServiceName)
}
