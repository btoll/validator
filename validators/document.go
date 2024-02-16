package validators

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/btoll/validator/lib"
)

var reKind = regexp.MustCompile(`"kind":\s"(?P<Kind>.*)",`)

type Manifest interface {
	ConfigMapManifest | DeploymentManifest | IngressManifest | ServiceManifest
	PrintTopLevelManifest()
	PrintSpec()
}

type Document[T Manifest] struct {
	Manifest T `json:"manifest"`
}

type Metadata struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Labels    Labels `json:"labels,omitempty"`
}

type Labels map[string]string

func (m *Document[T]) DecodeAndPrint(dec *json.Decoder) error {
	err := dec.Decode(&m.Manifest)
	if err != nil {
		return err
	}
	m.Manifest.PrintTopLevelManifest()
	m.Manifest.PrintSpec()
	return nil
}

func New(filename string) {
	b, _ := lib.GetFileContents(filename)
	v := reKind.FindAllSubmatch(b, -1)

	if len(v) > 0 {
		dec := json.NewDecoder(strings.NewReader(string(b)))
		var err error

		for _, kind := range v {
			switch string(kind[1]) {
			case "ConfigMap":
				var m Document[ConfigMapManifest]
				err = m.DecodeAndPrint(dec)
			case "Deployment":
				var m Document[DeploymentManifest]
				err = m.DecodeAndPrint(dec)
			case "Ingress":
				var m Document[IngressManifest]
				err = m.DecodeAndPrint(dec)
			case "Service":
				var m Document[ServiceManifest]
				err = m.DecodeAndPrint(dec)
			}

			fmt.Printf("----------------------------------------------\n\n")

			if err == io.EOF {
				// all done
				break
			}
		}
	}
}
