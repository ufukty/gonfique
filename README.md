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

-   Using dynamic keys in config files is not problematic as using them in server-client communication. Because all keys are final before compilation starts.
-   Having to search to get a piece of config is error-prone, makes harder to notice breaking changes. Also creates ugly code.

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

```sh
# Run
gonfique -in config.yml -out config.go -pkg main
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
							List struct {
								Method string `yaml:"Method"`
								Path   string `yaml:"Path"`
							} `yaml:"List"`
						} `yaml:"Endpoints"`
					} `yaml:"Document"`
					Objectives struct {
						Path      string `yaml:"Path"`
						Endpoints struct {
							Create struct {
								Method string `yaml:"Method"`
								Path   string `yaml:"Path"`
							} `yaml:"Create"`
						} `yaml:"Endpoints"`
					} `yaml:"Objectives"`
					Tags struct {
						Path      string `yaml:"Path"`
						Endpoints struct {
							Creation struct {
								Path   string `yaml:"Path"`
								Method string `yaml:"Method"`
							} `yaml:"Creation"`
							Assign struct {
								Method string `yaml:"Method"`
								Path   string `yaml:"Path"`
							} `yaml:"Assign"`
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
// main.go
package main

import "fmt"

func main() {
    cfg, err := ReadConfig("config.yml")
    if err != nil {
        //
    }
    // fields are suggested as typing by IDE
    fmt.Println(cfg.Github.
        Gateways.Public.
        Services.Document.
        Endpoints.List.Method,
    )
}
```

## Limitations

-   Slices are not supported as using dynamic keys makes accessin information easier.
-   YAML file can not have Go keywords as keys, eg. `type: user`

## Todo

-   [ ] Use user provided types for pieces of YAML when schemas match. (`-existing <file>` flag)

## Contribution

Issues are open for discussions and rest.

## License

Apache2. See LICENSE file.
