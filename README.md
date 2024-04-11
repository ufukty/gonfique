# `gonfique` - Type checked configs for Go programs

<img src="assets/Gonfique.png" alt="Gonfique logo" height="300px">

`gonfique` is a CLI tool for Go developers to automatically build exact **struct definitions** in Go that will match the provided YAML config. Designed to get all config accesses under **type check**. Makes breaking changes instant to notice when and where they happen.

Since `gonfique` makes **keeping the type definitions up-to-date** easier, using more dynamic keys instead arrays is more practical. So, developers can access config data through field names instead error-prone array lookups.

`gonfique` is necessary because of there are not many reliable and sustainable alternatives.

- Accessing config data through **hardcoded strings are risky**. So, defining types to marshall into is necessary.
- Manually defining types is also risky because they will **get outdated** eventually.
- Config **complexity is inevitable** when there are multiple services/binaries that needs their config to stay in sync, eg. kubernetes config.

## Your config file

### Before `gonfique`

Currently, you are storing multiple items **in arrays** like this

```yaml
github:
  domain: github.com
  path: /api/v1.0.0
  services:
    tags:
      path: tags
      endpoints:
        - name: list
          method: GET
          path: "list/{root}"
        - name: create
          method: POST
          path: "task"
        - name: assign
          method: POST
          path: "assign"
        - name: delete
          method: DELETE
          path: ""
gitlab:
  domain: gitlab.com
```

Then, you define `Config` type:

```go
type Config struct {
  Github struct {
    Domain   int
    Path     int
    Services []Service
  } `yaml:"github.com"`
}
```

And check again & again, if you made a mistake...

Lastly, access items **by lookups** with **hardcoded strings**. Compiler won't catch if the information gets outdated. Which leads you notice problems in **runtime**.

```go
func main() {
  // ...
  list, found := cfg.Github.Services.Tags.Endpoints.Lookup("list")
  if found {
    fmt.Println(list.Path)
  }
}
```

### After `gonfique`

You don't have to worry anymore about storing part of the config information in schema, since type definitions are **automatically generated**, **instantly updated** and accesses are **re-verified by compiler** as config changed.

```yaml
github:
  domain: github.com
  path: /api/v1.0.0
  services:
    tags:
      path: tags
      endpoints:
        list: { method: "GET", path: "tags" }
        create: { method: "POST", path: "tag" }
        assign: { method: "POST", path: "assign" }
        delete: { method: "DELETE", path: "tag" }
gitlab:
  domain: gitlab.com
```

Access items through fields instead hardcoded strings. Always under type-check, and easily updated as config changes.

```go
func main() {
  // ...
  fmt.Println(cfg.Github.Services.Tags.Endpoints.List.Path)
}
```

Arrays and dictionaries are still **iteratable** via `.Range` method (if the `-organize` flag provided):

```go
func main() {
  // ...
  for key, ep := range cfg.Github.Services.Tags.Endpoints.Range() {
    fmt.Println(key, ep.Path)
  }
}
```

## Full example

### Kubernetes

- [Input config for all](/examples/k8s/input.yml)
- Different flag combinations:
  - [Generated Go file](/examples/k8s/basic/output.go) when only `-in`, `-out` and `-pkg` flags are set
  - [Generated Go file](/examples/k8s/organized/output.go) when also `-organize` flag is set
  - [Generated Go file](/examples/k8s/organized-used/output.go) when both `-organize` and `-use <file>` flag are set
- [Usage for each](/examples/k8s/usage_test.go)

## Usage

### Install

```sh
go install github.com/ufukty/gonfique@v1.1.0
```

> [!IMPORTANT]
> Do not install code from `main` branch. Use [Version tags](https://github.com/ufukty/gonfique/tags) as shown above for stable versions or download from [Releases](https://github.com/ufukty/gonfique/releases).

### Generation

```sh
gonfique -in config.yml -out config.go -pkg main [-use <file>] [-organize]
```

```sh
$ gonfique -h
Usage of gonfique:
  -in string
        input file path (yml or yaml)
  -organize
        (optional) defines the types of struct fields that are also structs separately instead inline, with auto generated UNSTABLE names.
  -out string
        output file path (go)
  -pkg string
        package name that will be inserted into the generated file
  -use string
        (optional) use type definitions found in <file>
```

### Serving suggestions

For existing Makefile users:

```Makefile
config.go: config.yml
    gonfique -in config.yml -out config.go -pkg main

all: config.go
    ...
```

For existing Visual Studio Code users:

```json
{
  "runOnSave.commands": [
    {
      "match": "^config.yml$",
      "command": "cd '${fileDirname}' && make config.go"
    }
  ]
}
```

## Features

### Using existing type definitions

```go
package main

type Endpoint struct {
  Path, Method string
}
```

```sh
gonfique -in config.yml -out config.go -pkg main -use use.go
```

### User Provided Type Names

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

## Considerations

### Arrays

`gonfique` assignes the necessary slice type to arrays. It works best when all items of an array possess the same schema or at least compatible schemas. Type assignment occurs as below when items have not same but compatible schemas:

```yaml
# input
- action: foo
  foo-details: ""
- action: bar
  bar-details: ""
```

```go
// output
[]struct {
  Action     string
  FooDetails string
  BarDetails string
}
```

> [!IMPORTANT]
> Slice type gets defined as `[]any` if shared keys have different type values. Like `detail` has given `int` and `string` values below:
>
> ```yaml
> - action: ""
>   detail: 0
> - action: ""
>   detail: ""
> ```

## Limitations

- Multidocument YAML files are not supported.
- `gonfique` only creates inline type definitions. [See issue](issues/1) for discussion.

## Contribution

Issues are open for discussions and rest.

## License

Apache2. See LICENSE file.
