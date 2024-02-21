package validators

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"regexp"
	"strings"

	"github.com/btoll/validator/lib"
)

var reKind = regexp.MustCompile(`"kind":\s"(?P<Kind>.*)",`)

type Manifest interface {
	ConfigMapManifest | DeploymentManifest | IngressManifest | ServiceManifest
	Write()
}

type Document[T Manifest] struct {
	Manifest T `json:"manifest"`
}

type Metadata struct {
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      Data              `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Data map[string]string

var tpl *template.Template

func (m *Document[T]) DecodeAndPrint(dec *json.Decoder) error {
	err := dec.Decode(&m.Manifest)
	if err != nil {
		return err
	}
	m.Manifest.Write()
	return nil
}

func New(filename string) {
	b, _ := lib.GetFileContents(filename)
	v := reKind.FindAllSubmatch(b, -1)

	if len(v) > 0 {
		var err error

		tpl, err = template.ParseGlob("tpl/*")
		if err != nil {
			fmt.Println(err)
		}

		dec := json.NewDecoder(strings.NewReader(string(b)))

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

			if err == io.EOF {
				// all done
				break
			}
		}
	}
}
