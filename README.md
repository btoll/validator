## Example

To validate or compare a set of generated Kustomized manifests, simply point the `kustomize` build command at the `validator` tool:

```bash
$ ./validator -file <(kubectl kustomize ~/projects/migrator/build/project_z/debug/overlays/beta | yq -ojson)
```

That's it!  `validator` will range over the generated file of `json` blobs and parse each one in turn, sending its formatted output to `stdout`.

All that's left to do is use a `diff` tool to compare the files.

## References

[`yq` on GitHub]: https://github.com/mikefarah/yq/

