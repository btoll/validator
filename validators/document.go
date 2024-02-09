package validators

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/btoll/validator/lib"
)

// TODO
var reDeployment = regexp.MustCompile(`"kind": "Deployment"`)
var reIngress = regexp.MustCompile(`"kind": "Ingress"`)
var reService = regexp.MustCompile(`"kind": "Service"`)

func InterleaveDocuments(d []Document) {
	// This is janky.
	m0 := d[0].Manifest
	m1 := d[1].Manifest

	m0.PrintTopLevelManifest()
	m1.PrintTopLevelManifest()

	m0.PrintSpec()
	m1.PrintSpec()
	//	PrintDeploymentSpec(m1.Spec)

	//	PrintContainer(m0)
	//
	// PrintContainer(&s1.Template.Spec.Containers[0])
}

type Document struct {
	Name     string   `json:"filename"`
	Manifest Manifest `json:"manifest"`
}

type Metadata struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Labels    Labels `json:"labels,omitempty"`
}

type Labels map[string]string

func NewDocument(filename string) *Document {
	contents, err := lib.GetFileContents(filename)
	// TODO I'm not crazy about the current error checking.
	lib.CheckError(err)
	var d Document
	if reDeployment.Match(contents) {
		d = Document{
			Name:     filename,
			Manifest: &DeploymentManifest{},
		}
	} else if reIngress.Match(contents) {
		d = Document{
			Name:     filename,
			Manifest: &IngressManifest{},
		}
	} else {
		d = Document{
			Name:     filename,
			Manifest: &ServiceManifest{},
		}
	}
	err = json.Unmarshal(contents, &d.Manifest)
	lib.CheckError(err)
	return &d
}

func (d *Document) Print() {
	d.Manifest.PrintTopLevelManifest()
	d.Manifest.PrintSpec()
	// PrintContainer(&d.Manifest.Spec.Template.Spec.Containers[0])
}

func (d *Document) String() string {
	return fmt.Sprintf("%#v", d)
}
