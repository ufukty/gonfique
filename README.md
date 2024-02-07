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

`gonfique` is necessary because of there are not many reliable and sustainable alternatives.

-   Accessing config data through hardcoded strings is risky. So, defining types to marshall into is necessary.
-   Manually defining types is also risky because they will get outdated eventually.
-   Config complexity is inevitable when there are multiple services/binaries that needs their config to stay in sync, eg. kubernetes config.

## Full example

-   Kubernetes example [Input config for all](/examples/k8s/input.yml) [Usage for each](/examples/k8s/usage_test.go)
    -   [Generated Go file](/examples/k8s/basic/output.go) when only `-in`, `-out` and `-pkg` flags are set
    -   [Generated Go file](/examples/k8s/organized/output.go) when also `-organize` flag is set
    -   [Generated Go file](/examples/k8s/organized-used/output.go) when both `-organize` and `-use <file>` flag are set

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


## Limitations

-   Multidocument YAML files are not supported.
-   `gonfique` only creates inline type definitions. [See issue](issues/1) for discussion.


## Contribution

Issues are open for discussions and rest.

## License

Apache2. See LICENSE file.
