package validators

import "sort"

type ConfigMapManifest struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Data       Data     `json:"data"`
}

// The `deployment` validator needs access to these env vars so
// it can plug them into its template.  This needs to happen
// because the current Deployment has the env vars embedded but
// the new Deployment will access them via a `ConfigMap`.
var ConfigMapEnvVars []EnvVar

func (m ConfigMapManifest) Write() {
	for k, v := range m.Data {
		ConfigMapEnvVars = append(ConfigMapEnvVars, EnvVar{
			Name:  k,
			Value: v,
		})
	}
	sort.Slice(ConfigMapEnvVars, func(i, j int) bool {
		return ConfigMapEnvVars[i].Name < ConfigMapEnvVars[j].Name
	})
}
