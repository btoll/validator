## Install `yq`

From the [install section] on GitHub:

```bash
$ wget https://github.com/mikefarah/yq/releases/download/v4.40.5/yq_linux_amd64.tar.gz -O - \
    | tar xz \
    && sudo mv yq_linux_amd64 /usr/bin/yq
```

## Example

To validate or compare a set of generated Kustomized manifests, simply point the `kustomize` build command at the `validator` tool:

```bash
$ ./validator -file <(kubectl kustomize ~/projects/migrator/build/project_z/debug/overlays/beta | yq -o=json)
```

That's it!  `validator` will range over the generated file of `json` blobs and parse each one in turn, sending its formatted output to `stdout`.

## References

[`yq` on GitHub]: https://github.com/mikefarah/yq/
[install section]: https://github.com/mikefarah/yq/#install

