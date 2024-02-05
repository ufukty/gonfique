# `gonfique` - Type checked configs for Go programs

![Gonfique logo](assets/Gonfique@400w.png)

`gonfique` is a CLI tool for Go developers to automatically build exact struct definitions in Go that will match the provided YAML config. Designed to get all config accesses under type check. Makes breaking changes instant to notice when and where they happen. Since `gonfique` makes keeping the type definitions up-to-date easier, using more dynamic keys instead arrays is more practical. So, developers can access config data through field names instead error-prone array lookups.

Before `gonfique`

```yaml
Config:
  - Name: A
    Details:
      X:
      Y:
  - Name: B
    Details:
      X:
      Y:
```

```go
func main() {
  // ...
  fmt.Println(
    cfg.Find("A").Details.Get("X"),
    cfg[0].Details.Get("Y"),
  )
}
```

After `gonfique`

```yaml
A:
  X:
  Y:
B:
  X:
  Y:
```

```go
func main() {
  // ...
  fmt.Println(cfg.A.X, cfg.A.Y)
}
```

## Motivation

`gonfique` is necessary because of there are not many reliable and sustainable alternatives.

-   Accessing config data through hardcoded strings is risky. So, defining types to marshall into is necessary.
-   Manually defining types is also risky because they will get outdated eventually.
-   Config complexity is inevitable when there are multiple services/binaries that needs their config to stay in sync, eg. kubernetes config.

## Usage

### Install

```sh
go install github.com/ufukty/gonfique@v1.0.0
```

> [!IMPORTANT]
> Installing the latest version in `main` branch is not suggested as it is used for active development. Use [version flags](tags) as shown above for stable versions or download compiled binaries from [Releases](releases).

```sh
gonfique -h
Usage of gonfique:
  -in string
        input file path (yml or yaml)
  -out string
        output file path (go)
  -pkg string
        package name that will be inserted into the generated file
  -use string
        (optional) use type definitions found in <file>
```

### Generation

```sh
gonfique -in config.yml -out config.go -pkg main
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
> ```yaml
> - action: ""
>   detail: 0
> - action: ""
>   detail: ""
> ```


## Full example

-   Kubernetes example [Input YAML file](/examples/k8s/input.yml) > [Generated Go file](/examples/k8s/output.go)

## Limitations

-   Multidocument YAML files are not supported.
-   `gonfique` only creates inline type definitions. [See issue](issues/1) for discussion.


## Contribution

Issues are open for discussions and rest.

## License

Apache2. See LICENSE file.
