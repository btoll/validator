APIVersion: {{ .APIVersion }}
Kind: {{ .Kind }}
Name: {{ .Name }}
Namespace: {{ .Namespace }}
Labels: {{ .Labels }}
{{- $spec := .Spec }}
Replicas: {{ $spec.Replicas }}
Selector: {{ $spec.Selector.MatchLabels }}
{{ $container := $spec.Template.Spec.Containers }}
{{- range $container -}}
Name: {{ .Name }}
Image: {{ .Image }}
ImagePullPolicy: {{ .ImagePullPolicy }}
EnvFrom: {{ .EnvFrom }}
Ports:
    {{- range .Ports -}}
    {{ .ContainerPort }}
    {{- end }}
Resources:
    Limits:{{ .Resources.Limits }}
    Requests: {{ .Resources.Requests }}
{{- end }}
