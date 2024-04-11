# Mapping

Users can specify how the generated structs will be named. The way to describe mappings is writing a list of `pathway: typename` pairs into a file and pass it to `gonfique` with `-mapping <file>`. Syntax of the mapping file is shown below:

```yaml
apiVersion: ApiVersion
metadata.name: Name
spec.template.metadata.labels.app: AppLabel
```

Wildcards lets users to write more flexible mappings.

Single-level wildcards match with any key in a dictionary, and they can be used many times in a pathway. The specified type name will be

```yaml
spec.template.*.labels.app: AppLabel
spec.*.*.labels.app: AppLabel
```

Multi-level wildcards match zero-to-many depth of dictionaries:

```yaml
spec.**.app: AppLabel
```

That would match all of the `spec.app`, `spec.foo.app` and `spec.bar.[].app` same time.

Array item type:

```yaml
spec.template.spec.containers.[]: Container
```

A key's type in any item:

```yaml
spec.template.spec.containers.[].Name: ContainerName
```

If the array type also needs to be given a name:

```yaml
spec.template.spec.containers: Containers
```
