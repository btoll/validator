package validators

type ConfigMapManifest struct {
	APIVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Metadata   Metadata      `json:"metadata"`
	Spec       ConfigMapSpec `json:"spec"`
}

type ConfigMapSpec struct {
	Type string
}

func (c ConfigMapManifest) PrintTopLevelManifest() {
}

func (c ConfigMapManifest) PrintSpec() {
}
