package validators

import (
	"fmt"

	"github.com/btoll/validator/lib"
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
	TargetPort int
	Protocol   string
}

func (m ServiceManifest) Write() {
	properServiceName := lib.GetProperServiceName(m.Name)
	m.Name = properServiceName

	dir := fmt.Sprintf("build/%s/service", properServiceName)
	err := lib.CreateBuildDir(dir)
	if err != nil {
		fmt.Println("err", err)
	}
	f, err := lib.CreateBuildFile(fmt.Sprintf("%s/local", dir))
	if err != nil {
		fmt.Println("err", err)
	}
	err = tpl.ExecuteTemplate(f, "service.tpl", m)
	if err != nil {
		fmt.Println("err", err)
	}

	PrintService(properServiceName)
}
