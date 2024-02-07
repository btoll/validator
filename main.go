package main

// TODO:
// - use something like go-flags
// 		+ https://github.com/jessevdk/go-flags

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type Validator struct {
	Documents []Document
}

func NewValidator(d1, d2 *Document) *Validator {
	return &Validator{
		Documents: []Document{*d1, *d2},
	}
}

func printContainer(c Container) {
	// TODO: better use of formatters
	name := fmt.Sprintf("            Name: %s\n", c.Name)
	image := fmt.Sprintf("          Image: %s\n", c.Image)
	imagePullPolicy := fmt.Sprintf("ImagePullPolicy: %s\n", c.ImagePullPolicy)
	envFrom := fmt.Sprintf("        EnvFrom: %s\n", c.EnvFrom)
	ports := fmt.Sprintf("          Ports: %d\n", c.Ports)
	resources := fmt.Sprintf("      Resources: %s\n", c.Resources)
	fmt.Println(name, image, imagePullPolicy, envFrom, ports, resources)
}

func (v *Validator) Validate() {
	d1 := v.Documents[0]
	d2 := v.Documents[1]

	c1 := d1.Spec.Template.Spec.Containers[0]
	printContainer(c1)
	c2 := d2.Spec.Template.Spec.Containers[0]
	printContainer(c2)
}

//type Document struct {
//	Name string
//	Data []byte
//}

func NewDocument(filename string) *Document {
	// TODO
	var err error
	var contents1 []byte
	if checkFileExists(filename) {
		contents1, err = os.ReadFile(filename)
		checkError(err)
	}
	m := Document{}
	err = json.Unmarshal(contents1, &m)
	checkError(err)
	return &m
}

func (m *Document) String() string {
	return fmt.Sprintf("%#v", m)
}

type Document struct {
	APIVersion string       `json:"apiVersion,omitempty"`
	Kind       string       `json:"kind,omitempty"`
	Metadata   Metadata     `json:"metadata,omitempty"`
	Spec       ManifestSpec `json:"spec,omitempty"`
}

type Metadata struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Labels    Labels `json:"labels,omitempty"`
}

type Labels map[string]string

type ManifestSpec struct {
	Replicas int         `json:"replicas,omitempty"`
	Selector MatchLabels `json:"selector,omitempty"`
	Template Template    `json:"template,omitempty"`
}

type MatchLabels struct {
	MatchLabel Labels `json:"matchLabels,omitempty"`
}

type Template struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Spec     PodSpec  `json:"spec,omitempty"`
}

type PodSpec struct {
	Containers []Container `json:"containers,omitempty"`
}

type Container struct {
	Name            string    `json:"name,omitempty"`
	Image           string    `json:"image,omitempty"`
	ImagePullPolicy string    `json:"imagePullPolicy,omitempty"`
	EnvFrom         []EnvFrom `json:"envFrom,omitempty"`
	Ports           []Port    `json:"ports,omitempty"`
	Resources       Resources `json:"resources,omitempty"`
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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func main() {
	file1 := flag.String("file1", "deployment.json", "The name of the file to validate")
	file2 := flag.String("file2", "deployment.json", "The name of the other file to validate")
	//	raw := flag.Bool("raw", false, "Print the raw unmarshaled JSON")
	flag.Parse()

	v := NewValidator(
		NewDocument(*file1),
		NewDocument(*file2))
	v.Validate()
}
