## Install `yq`

From the [install section] on GitHub:

```bash
$ wget https://github.com/mikefarah/yq/releases/download/v4.40.5/yq_linux_amd64.tar.gz -O - \
    | tar xz \
    && sudo mv yq_linux_amd64 /usr/bin/yq
```

## Generate `json`

```bash
$ kubectl kustomize beta | yq -o=json -
```

## Example

Validate a `deployment`:

```bash
$ ./validator -file1 aion_nginx_deployment_beta.json -file2 <(kubectl get deployments.apps aion-nginx -ojson)
      APIVersion: apps/v1
            Kind: Deployment
        Metadata: {aion-nginx default map[app:aion-nginx]}

      APIVersion: apps/v1
            Kind: Deployment
        Metadata: {aion-nginx default map[app:aion-nginx]}

        Replicas: 6
        Selector: {map[app:aion-nginx]}

        Replicas: 6
        Selector: {map[app:aion-nginx]}

            Name: aion-nginx
           Image: 451310829282.dkr.ecr.us-east-1.amazonaws.com/aion/aion-nginx:development
 ImagePullPolicy:
         EnvVars: []
         EnvFrom: [{map[name:env-kf57km4hkd]}]
           Ports: [{0}]
       Resources: {{250m 1Gi} {100m 128Mi}}

            Name: aion-nginx
           Image: 451310829282.dkr.ecr.us-east-1.amazonaws.com/aion/aion-nginx:master-b1758a4
 ImagePullPolicy: Always
         EnvVars: [{Name:ENV Value:beta} {Name:DOMAIN Value:.default.svc.cluster.local} {Name:HTTPS_REDIRECT_FROM Value:www.app.beta-veriforceone.com} {Name:HTTPS_REDIRECT_TO Value:app.beta-veriforceone.com} {Name:HTTP_SCHEME Value:https} {Name:INSTRUCTORPORTAL_WEBAPPS Value:server instructorportal-web.beta-compliancepro-cacentral1.private;} {Name:LEGACY_WEBAPPS Value:server chronos-web.beta-compliancepro-cacentral1.private;} {Name:RESOLVER Value:kube-dns.kube-system.svc.cluster.local} {Name:SERVERNAME Value:app.beta-veriforceone.com www.app.beta-veriforceone.com} {Name:TRAINING_LEGACY_WEBAPPS Value:training-web.beta-compliancepro-cacentral1.private;} {Name:TRAINING_SERVER Value:beta-training.net} {Name:AwsRegion Value:ca-central-1} {Name:MAINTENANCE_MODE Value:off}]
         EnvFrom: []
           Ports: [{0}]
       Resources: {{ } { }}
```

Validate an `service`:

```bash
$ ./validator -file1 aion_nginx_ingress_beta.json -file2 <(kubectl get ingress aion-nginx -ojson)
```

```bash
$ ./validator -file1 aion_nginx_service_beta.json -file2 <(kubectl get svc aion-nginx -ojson)
```

Validate a `ingress`:

---

> Here is something fun:
>
> ```bash
> $ for doc in $(kubectl kustomize ~/projects/migrator/build/aion/aion-nginx/overlays/beta | yq -o=json | jq -c inputs)
>   do
>       ./validator -file1 <(echo "$doc") -file2 <(kubectl get deployments.app aion-nginx -ojson)
>   done
> ```

## References

[`yq` on GitHub]: https://github.com/mikefarah/yq/
[install section]: https://github.com/mikefarah/yq/#install

