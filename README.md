# validator

`validator` will take one or more deployments and compare them against the same-named deployment in the cloud, which is considered the source of truth.  It does this by passing `validator` a kustomized-generated manifest(s) file of a local deployment and then comparing its values against a remote deployment running in `eks`, for example.

If there is any discrepancy between the two, `validator` will create a child directory in the `build` directory (the location of which is configurable) which is the name of the deployment.  Within this directory, the tool will create a directory for each type of Kubernetes resource that failed the equality comparison:

- `deployment`
- `ingress`
- `service`

There will be a `local` and a `remote` file within each of those child directories, and these are intended to be compared with a `diff` tool (`validator` is unopinionated about that).

For example, here is what the `build` directory looks like after a recent operation:

```bash
$ tree build
build/
├── aion-signup-outbox-relay/
│   └── deployment/
│       ├── local
│       └── remote
└── aion-subscription-outbox-relay/
    └── deployment/
        ├── local
        └── remote
```

> Deployments that are the same, i.e., the local deployment compares to the remote deployment (the source of truth), will not have a directory in `build`.

## Assumptions

- The `kubectl` tool has been installed.
- The `yq` tool has been installed.

## Cluster Context

This tool depends on the operator already having set the proper cluster context.  If you're not sure which is the current context, run the following:

```bash
$ kubectl config current-context
```

List all contexts:

```bash
$ kubectl config get-contexts
```

Set a particular context:

```bash
$ kubectl config use-context development-compliancepro-cacentra
```

## Examples

### Validate Single Deployment

To validate or compare a set of generated Kustomized manifests, simply point the `kustomize` build command at the `validator` tool:

```bash
$ ./validator --file <(kubectl kustomize ~/projects/veriforce/devops/gitops-test/applications/aion/aion-datapipeline-organizationfeature/overlays/beta/ | yq -o json)
```

That's it!  `validator` will range over the generated file of `json` blobs and parse each one in turn, sending its formatted output to `stdout`.

All that's left to do is use a `diff` tool to compare the files.

### Validate Multiple Deployments

This will validate all of the deployments in the cloud against the local deployment in the `gitops` repository:

```bash
$ make validate
```

## References

- [`kubectl config`](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_config/)
- [`yq` on GitHub](https://github.com/mikefarah/yq/)

