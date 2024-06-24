package validators

import (
	"fmt"

	"github.com/btoll/validator/lib"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ServiceManifest struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   `json:"metadata"`
	Spec       ServiceSpec `json:"spec"`
}

type ServiceSpec struct {
	Ports    []ServicePort
	Selector Data
	Type     string
}

type ServicePort struct {
	Port       int
	TargetPort intstr.IntOrString
	Protocol   string
}

func (m ServiceManifest) Write() {
	dir := fmt.Sprintf("build/%s/service", m.Name)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	localFile := fmt.Sprintf("%s/local", dir)
	WriteTemplate(localFile, "service.tpl", m)

	service, err := GetServiceClient(m.Name)
	if err != nil {
		fmt.Println("err", err)
	}
	// This values are empty in the returned struct.
	service.APIVersion = "v1"
	service.Kind = "Service"
	remoteFile := fmt.Sprintf("%s/remote", dir)
	WriteTemplate(remoteFile, "service.tpl", service)
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
