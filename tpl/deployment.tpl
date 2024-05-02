APIVersion: {{ .APIVersion }}
Kind: {{ .Kind }}
Name: {{ .Name }}
Namespace: {{ .Namespace }}
Labels: {{ .Labels }}
{{- $spec := .Spec }}
Replicas: {{ $spec.Replicas }}
Selector: {{ $spec.Selector.MatchLabels }}

Volumes:
    {{ $volumes := $spec.Template.Spec.Volumes }}
    {{- range $volumes -}}
    Name: {{ .Name }}
    VolumeSource: {{ .VolumeSource }}
    {{- end }}

NodeSelector: {{ $spec.Template.Spec.NodeSelector }}

Container:
    {{ $containers := $spec.Template.Spec.Containers }}
    {{- range $containers -}}
    Name: {{ .Name }}
    Image: {{ .Image }}
    ImagePullPolicy: {{ .ImagePullPolicy }}
    Env:
        {{- range .Env }}
        {{ .Name }}={{ .Value }}
        {{- end }}
    Ports:
        {{- range .Ports -}}
        {{ .ContainerPort }}
        {{- end }}
    Resources:
        Limits:{{ .Resources.Limits }}
        Requests: {{ .Resources.Requests }}
    VolumeMounts:
        {{- range .VolumeMounts }}
        Name: {{ .Name }}
        MountPath: {{ .MountPath }}
        SubPath: {{ .SubPath }}
        {{- end }}
    {{- end }}
