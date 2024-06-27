package validators

import (
	"fmt"
	"sort"

	"github.com/btoll/validator/lib"
	v1 "k8s.io/api/core/v1"
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
	Volumes      []Volume    `json:"volumes,omitempty"`
	NodeSelector Data        `json:"nodeSelector,omitempty"`
}

type Volume struct {
	Name string           `json:"name,omitempty"`
	NFS  *NFSVolumeSource `json:"nfs,omitempty"`
}

type NFSVolumeSource struct {
	Server   string `json:"server,omitempty"`
	Path     string `json:"path,omitempty"`
	ReadOnly bool   `json:"read_only,omitempty"`
}

type Container struct {
	Name            string             `json:"name,omitempty"`
	Image           string             `json:"image,omitempty"`
	ImagePullPolicy string             `json:"imagePullPolicy,omitempty"`
	Env             []EnvVar           `json:"env,omitempty"`
	EnvFrom         []EnvFrom          `json:"envFrom,omitempty"`
	Ports           []v1.ContainerPort `json:"ports,omitempty"`
	Resources       Resources          `json:"resources,omitempty"`
	VolumeMounts    []VolumeMount      `json:"volumeMounts,omitempty"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnvFrom struct {
	ConfigMapRef Data `json:"configMapRef,omitempty"`
}

type Resources struct {
	Limits   ResourceList `json:"limits,omitempty"`
	Requests ResourceList `json:"requests,omitempty"`
}

type ResourceList map[string]string

type VolumeMount struct {
	Name      string `json:"name,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	SubPath   string `json:"subPath,omitempty"`
}

func sortPorts(s []v1.ContainerPort) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].ContainerPort < s[j].ContainerPort
	})
}

func (m DeploymentManifest) Write() {
	fmt.Println(m.Name)
	dir := fmt.Sprintf("build/%s/deployment", m.Name)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	// Recall, `ConfigMapEnvVars` has already been sorted.
	// See the note in `validators/configmap` for more context.
	m.Spec.Template.Spec.Containers[0].Env = ConfigMapEnvVars
	sortPorts(m.Spec.Template.Spec.Containers[0].Ports)
	localFile := fmt.Sprintf("%s/local", dir)
	WriteTemplate(localFile, "deployment.tpl", m)
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
	sort.SliceStable(e, func(i, j int) bool {
		return e[i].Name < e[j].Name
	})
	sortPorts(deployment.Spec.Template.Spec.Containers[0].Ports)
	remoteFile := fmt.Sprintf("%s/remote", dir)
	WriteTemplate(remoteFile, "deployment.tpl", deployment)
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
