APIVersion: {{ .APIVersion }}
Kind: {{ .Kind }}
Name: {{ .Name }}
Namespace: {{ .Namespace }}
Annotations:
{{- range $key, $value := .Annotations }}
    {{ $key }}: {{ $value }}
{{- end }}
Rules:
{{- range .Spec.Rules }}
    {{- range .HTTP.Paths }}
        Path: {{ .Path }}
        PathType: {{ .PathType }}
        Backend:
	    Service:
                Name: {{ .Backend.Service.Name }}
                Port:
                    Name: {{ .Backend.Service.Port.Name }}
                    Number: {{ .Backend.Service.Port.Number }}
    {{ end }}
{{- end }}

