package validators

import (
	"fmt"
	"sort"

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
	Template Template      `json:"template,omitempty"`
}

type LabelSelector struct {
	MatchLabels Data `json:"matchLabels,omitempty"`
}

type Template struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Spec     PodSpec  `json:"spec,omitempty"`
}

type PodSpec struct {
	Containers   []Container `json:"containers,omitempty"`
	NodeSelector Data
}

type Container struct {
	Name            string    `json:"name,omitempty"`
	Image           string    `json:"image,omitempty"`
	ImagePullPolicy string    `json:"imagePullPolicy,omitempty"`
	Env             []EnvVar  `json:"env,omitempty"`
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

func WriteTemplate(to, from string, T any) error {
	f, err := lib.CreateBuildFile(to)
	if err != nil {
		return err
	}
	err = tpl.ExecuteTemplate(f, from, T)
	if err != nil {
		return err
	}
	return nil
}

func (m DeploymentManifest) Write() {
	dir := fmt.Sprintf("build/%s/deployment", m.Name)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	// Recall, `ConfigMapEnvVars` has already been sorted.
	// See the note in `validators/configmap` for more context.
	m.Spec.Template.Spec.Containers[0].Env = ConfigMapEnvVars
	WriteTemplate(fmt.Sprintf("%s/local", dir), "deployment.tpl", m)
	deployment, err := GetDeploymentClient(m.Name)
	if err != nil {
		fmt.Println("err", err)
	}
	// This values are empty in the returned struct.
	deployment.APIVersion = "apps/v1"
	deployment.Kind = "Deployment"
	// There should always be an `.Env` property here.  If there is not,
	// then this will fail and we can investigate.  This is a GOOD THING,
	// since by design the Deployment should have the env vars embedded in it.
	e := deployment.Spec.Template.Spec.Containers[0].Env
	sort.Slice(e, func(i, j int) bool {
		return e[i].Name < e[j].Name
	})
	WriteTemplate(fmt.Sprintf("%s/remote", dir), "deployment.tpl", deployment)
}
