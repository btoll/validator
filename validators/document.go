package validators

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
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

func RemoveDir(buildDir, serviceDir string) error {
	err := os.RemoveAll(buildDir)
	if err != nil {
		return err
	}
	f, err := os.Open(serviceDir)
	_, err = f.ReadDir(1)
	if errors.Is(err, io.EOF) {
		err := os.RemoveAll(serviceDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func Validate(localFile, remoteFile string) (bool, error) {
	lb, err := os.ReadFile(localFile)
	if err != nil {
		return false, err
	}
	rb, err := os.ReadFile(remoteFile)
	if err != nil {
		return false, err
	}
	return string(lb) == string(rb), nil
}

func WriteTemplate(to, from string, T any) error {
	f, err := os.Create(to)
	if err != nil {
		return err
	}
	err = tpl.ExecuteTemplate(f, from, T)
	if err != nil {
		return err
	}
	return nil
}
