## Install `yq`

From the [install section] on GitHub:

```bash
$ wget https://github.com/mikefarah/yq/releases/download/v4.40.5/yq_linux_amd64.tar.gz -O - \
    | tar xz \
    && sudo mv yq_linux_amd64 /usr/bin/yq
```

## Example

```bash
$ ./validator -file1 aion_nginx_kustomized_beta.json -file2 <(kubectl get deployments.apps aion-nginx -ojson)
            Name: aion-nginx
           Image: 451310829282.dkr.ecr.us-east-1.amazonaws.com/aion/aion-nginx:development
 ImagePullPolicy:
         EnvFrom: [{map[name:env-kf57km4hkd]}]
           Ports: [{0}]
       Resources: {{250m 1Gi} {100m 128Mi}}

            Name: aion-nginx
           Image: 451310829282.dkr.ecr.us-east-1.amazonaws.com/aion/aion-nginx:master-d90e721
 ImagePullPolicy: Always
         EnvFrom: []
           Ports: [{0}]
       Resources: {{ } { }}

```

## Generate `json`

```bash
$ kubectl kustomize beta | yq -o=json -
```

## References

[`yq` on GitHub]: https://github.com/mikefarah/yq/
[install section]: https://github.com/mikefarah/yq/#install

