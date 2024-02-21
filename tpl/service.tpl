APIVersion: {{ .APIVersion }}
Kind: {{ .Kind }}
Name: {{ .Name }}
Namespace: {{ .Namespace }}
Labels: {{ .Labels }}
{{- $spec := .Spec }}
Selector: {{ $spec.Selector }}
Ports:
    {{ range $spec.Ports -}}
    Port: {{ .Port }}
    Protocol: {{ .Protocol }}
    TargetPort: {{ .TargetPort }}
    {{- end }}
Ports: {{ $spec.Type }}
