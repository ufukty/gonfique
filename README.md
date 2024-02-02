# `gonfique` - Type checked configs for Go programs

![Gonfique logo](assets/Gonfique@400w.png)

`gonfique` is a tool that creates Go struct type definitions for given YAML file which has custom schemas. Since `gonfique` makes it easier to use custom schemas, developers can levarage dynamic-keyed configs to access information with autocomplete and without searching through items' values. Also existing Makefile users can get instantly notified when and where a config change breaks Go code (See [Serving suggestion](#serving-suggestions)).

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

## Install

```sh
go install github.com/ufukty/gonfique
```

## Usage

```sh
gonfique -in config.yml -out config.go -pkg main
```

```
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

## Serving suggestions

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

## Motivation

gonfique is necessary because of there are not many reliable and sustainable alternatives.

-   Accessing config data through hardcoded strings is risky. So, defining types to marshall into is necessary.
-   Manually defining types is also risky because they will get outdated eventually.
-   Config complexity is inevitable when there are multiple services/binaries that needs their config to stay in sync, eg. kubernetes config.

## Full example

```yml
# config.yml
Github:
  Domain: github.com
  Gateways:
    Public:
      Path: /api/v1.0.0
      Services:
        Document:
          Path: document
          Endpoints:
            List: { Method: "GET", Path: "list/{root}" }
        Objectives:
          Path: tasks
          Endpoints:
            Create: { Method: "POST", Path: "task" }
        Tags:
          Path: tags
          Endpoints:
            Creation: { Method: "POST", Path: "" }
            Assign: { Method: "POST", Path: "assign" }
Gitlab:
  Domain: gitlab.com
```

```go
// (optional) models.go
package main

type Endpoint struct {
  Method string
  Path   string
}
```

```sh
gonfique -in config.yml -out config.go -pkg main -use models.go
```

```go
// Produced file "config.go"
package main

import (
  "fmt"
  "os"
  "gopkg.in/yaml.v3"
)

type Config struct {
  Github struct {
    Domain   string `yaml:"Domain"`
    Gateways struct {
      Public struct {
        Path     string `yaml:"Path"`
        Services struct {
          Document struct {
            Path      string `yaml:"Path"`
            Endpoints struct {
              List Endpoint `yaml:"List"` // here
            } `yaml:"Endpoints"`
          } `yaml:"Document"`
          Objectives struct {
            Path      string `yaml:"Path"`
            Endpoints struct {
              Create Endpoint `yaml:"Create"` // here
            } `yaml:"Endpoints"`
          } `yaml:"Objectives"`
          Tags struct {
            Path      string `yaml:"Path"`
            Endpoints struct {
              Creation Endpoint `yaml:"Creation"` // here
              Assign   Endpoint `yaml:"Assign"` // here
            } `yaml:"Endpoints"`
          } `yaml:"Tags"`
        } `yaml:"Services"`
      } `yaml:"Public"`
    } `yaml:"Gateways"`
  } `yaml:"Github"`
  Gitlab struct {
    Domain string `yaml:"Domain"`
  } `yaml:"Gitlab"`
}

func ReadConfig(path string) (Config, error) {
  f, err := os.Open(path)
  if err != nil {
    return Config{}, fmt.Errorf("opening config file: %w", err)
  }
  cfg := Config{}
  err = yaml.NewDecoder(f).Decode(&cfg)
  if err != nil {
    return Config{}, fmt.Errorf("decoding config file: %w", err)
  }
  return cfg, nil
}
```

```go
// usage "main.go"
package main

import "fmt"

func RegisterEndpoints(eps []Endpoint) {
  for _, ep := range eps {
    fmt.Println(ep.Method, ep.Path)
  }
}

func main() {
  cfg, err := ReadConfig("config.yml")
  if err != nil {
    //
  }
  // fields are suggested as typing by IDE

  endpoints := Github.Gateways.Public.Services.Tags.Endpoints
  RegisterEndpoints([]Endpoint{
    endpoints.Creation, endpoints.Assign,
  })
}
```

## Limitations

-   Slices are not supported as using dynamic keys makes accessin information easier.
-   YAML file can not have Go keywords as keys, eg. `type: user`

## Todo

-   [x] Use user provided types for pieces of YAML when schemas match. (`-use <file>` flag)
-   [x] Stable field order in produced type spec

## Contribution

Issues are open for discussions and rest.

## License

Apache2. See LICENSE file.
