> [!IMPORTANT]  
> Check docs for upcoming Gonfique 2 in [dev branch](https://github.com/ufukty/gonfique/tree/dev) and go to [Gonfique Playground](https://gonfique-playground.pages.dev/) to try pre-alpha in your browser

# Gonfique

<img src="assets/Gonfique.png" alt="Gonfique logo" height="300px">

Gonfique is a CLI tool for Go developers to automatically build exact **struct definitions** in Go that will match the provided YAML or JSON config. Makes instant to notice **when and where a breaking change** occurs. Since compiler warns whenever it happens by type-checking, and source control shows where the change exactly is.

## TOC

- [Full example](#full-example)
  - [Output with nested declarations](#output-with-nested-declarations)
  - [Creating separate (named) type declarations](#creating-separate-named-type-declarations)
- [Install](#install)
- [Usage](#usage)
- [Features](#features)
- [Docs](#docs)
- [Limitations](#limitations)
- [Contribution](#contribution)
- [Stargazers-over-time](#stargazers-over-time)
- [License](#license)

## Full example

Say this file is your input:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  namespace: my-namespace
type: Opaque
data:
  my-key: my-value
  password: cGFzc3dvcmQ=
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
  rules:
    - host: myapp.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: my-service
                port:
                  number: 80
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
        - name: my-container
          image: my-image
          ports:
            - containerPort: 80
          envFrom:
            - configMapRef:
                name: my-config
            - secretRef:
                name: my-secret
```

### Output with nested declarations

Gonfique creates the most compact type declaration for the provided file when it's not supplied a mapping file or the flag `--organize`. The file will contain only one declaration with all sub-structs defined inlined. Such as:

```go
package config

import (
  "fmt"
  "os"

  "gopkg.in/yaml.v3"
)

type Config struct {
  ApiVersion string `yaml:"apiVersion"`
  Data       struct {
    MyKey    string `yaml:"my-key"`
    Password string `yaml:"password"`
  } `yaml:"data"`
  Kind     string `yaml:"kind"`
  Metadata struct {
    Name      string `yaml:"name"`
    Namespace string `yaml:"namespace"`
  } `yaml:"metadata"`
  Spec struct {
    Ports []struct {
      Port       int    `yaml:"port"`
      Protocol   string `yaml:"protocol"`
      TargetPort int    `yaml:"targetPort"`
    } `yaml:"ports"`
    Replicas int `yaml:"replicas"`
    Rules    []struct {
      Host string `yaml:"host"`
      Http struct {
        Paths []struct {
          Backend struct {
            Service struct {
              Name string `yaml:"name"`
              Port struct {
                Number int `yaml:"number"`
              } `yaml:"port"`
            } `yaml:"service"`
          } `yaml:"backend"`
          Path     string `yaml:"path"`
          PathType string `yaml:"pathType"`
        } `yaml:"paths"`
      } `yaml:"http"`
    } `yaml:"rules"`
    Selector struct {
      MatchLabels struct {
        App string `yaml:"app"`
      } `yaml:"matchLabels"`
    } `yaml:"selector"`
    Template struct {
      Metadata struct {
        Labels struct {
          App string `yaml:"app"`
        } `yaml:"labels"`
      } `yaml:"metadata"`
      Spec struct {
        Containers []struct {
          EnvFrom []struct {
            ConfigMapRef struct {
              Name string `yaml:"name"`
            } `yaml:"configMapRef"`
            SecretRef struct {
              Name string `yaml:"name"`
            } `yaml:"secretRef"`
          } `yaml:"envFrom"`
          Image string `yaml:"image"`
          Name  string `yaml:"name"`
          Ports []struct {
            ContainerPort int `yaml:"containerPort"`
          } `yaml:"ports"`
        } `yaml:"containers"`
      } `yaml:"spec"`
    } `yaml:"template"`
  } `yaml:"spec"`
  Type string `yaml:"type"`
}

func ReadConfig(path string) (Config, error) {
  file, err := os.Open(path)
  if err != nil {
    return Config{}, fmt.Errorf("opening config file: %w", err)
  }
  defer file.Close()
  c := Config{}
  err = yaml.NewDecoder(file).Decode(&c)
  if err != nil {
    return Config{}, fmt.Errorf("decoding config file: %w", err)
  }
  return c, nil
}
```

### Creating separate (named) type declarations

You can create named declarations for types you want. Just provide a mapping file that contains the "paths" of targets and the typename you pick. Pass the path of this file as an argument (via `--mappings <path>`) to gonfique. See docs: [Mapping file](docs/mapping.md)

```yml
apiVersion: ApiVersion
metadata.name: Name
spec.rules.[]: Rule
spec.rules.[].http.paths.[]: Path
spec.ports.[]: Port
spec.**.containers: SpecContainers
spec.**.containers.[]: SpecContainer
spec.**.containers.[].name: ContainerName
```

This time the generated file would contain additional types as the mappings file directs. The main type declaration refers to those type in related substructures.

```go
// ...

type ApiVersion string
type ContainerName string
type Name string
type Path struct {
  Backend struct {
    Service struct {
      Name string `yaml:"name"`
      Port struct {
        Number int `yaml:"number"`
      } `yaml:"port"`
    } `yaml:"service"`
  } `yaml:"backend"`
  Path     string `yaml:"path"`
  PathType string `yaml:"pathType"`
}
type Port struct {
  Port       int    `yaml:"port"`
  Protocol   string `yaml:"protocol"`
  TargetPort int    `yaml:"targetPort"`
}
type Rule struct {
  Host string `yaml:"host"`
  Http struct {
    Paths []Path `yaml:"paths"`
  } `yaml:"http"`
}
type SpecContainer struct {
  EnvFrom []struct {
    ConfigMapRef struct {
      Name string `yaml:"name"`
    } `yaml:"configMapRef"`
    SecretRef struct {
      Name string `yaml:"name"`
    } `yaml:"secretRef"`
  } `yaml:"envFrom"`
  Image string        `yaml:"image"`
  Name  ContainerName `yaml:"name"`
  Ports []struct {
    ContainerPort int `yaml:"containerPort"`
  } `yaml:"ports"`
}
type SpecContainers []SpecContainer
type Config struct {
  ApiVersion ApiVersion `yaml:"apiVersion"`
  Data       struct {
    MyKey    string `yaml:"my-key"`
    Password string `yaml:"password"`
  } `yaml:"data"`
  Kind     string `yaml:"kind"`
  Metadata struct {
    Name      Name   `yaml:"name"`
    Namespace string `yaml:"namespace"`
  } `yaml:"metadata"`
  Spec struct {
    Ports    []Port `yaml:"ports"`
    Replicas int    `yaml:"replicas"`
    Rules    []Rule `yaml:"rules"`
    Selector struct {
      MatchLabels struct {
        App string `yaml:"app"`
      } `yaml:"matchLabels"`
    } `yaml:"selector"`
    Template struct {
      Metadata struct {
        Labels struct {
          App string `yaml:"app"`
        } `yaml:"labels"`
      } `yaml:"metadata"`
      Spec struct {
        Containers SpecContainers `yaml:"containers"`
      } `yaml:"spec"`
    } `yaml:"template"`
  } `yaml:"spec"`
  Type string `yaml:"type"`
}

// ...
```

See outputs for different flag combinations:

- [`-organize`](/examples/k8s/organized/output.go)
- [`-organize`, `-use`](/examples/k8s/organized-used/output.go)

## Install

Use version tags to avoid development versions as such:

```sh
go install github.com/ufukty/gonfique@v1.5.3
```

## Usage

```sh
gonfique -in config.yml -out config.go -pkg main [-type-name "Config"] [-use "use.go"] [-organize] [-mappings "mappings.yml"]
```

Run `gonfique --help` for [parameter details](docs/parameter-details.txt).

## Features

- Works offline
- Specify names for detected types via `-mapping` flag to export safely. [More](docs/mapping.md)
- Defines named or inline types depending on user choice (`-organize` flag).
- Reuses user-defined types provided in a file when schemas match via `-use` flag
- Supports arrays:
  - When all array items share same schema; array types are generated as `[]<item type>`.
  - When all array items don't share same schema, but all items are compatible; array types are generated as `[]<combined item types>`.
- Supports dictionaries:
  - Implements `.Range()` method on those dictionaries that all values share same schema, so they can be iterable via for loops. Such as `for service, details := range cfg.Services.Range() { ... }`
- Supports `time.Duration` values such as `200ms` or `1µs`.
- Supports JSON and YAML config files.

## Docs

- [Why Gonfique?](docs/why-gonfique.md)
- [Arrays](docs/arrays.md)
- [Mapping file](docs/mapping.md)
- [Suggestions](docs/suggestions.md)

## Limitations

- Multidocument YAML files are not supported.
- Mapping file is for generating an exported type definition with the given name and detected type; not to specify type.
- Type of keys with value of `null` gets assigned as `any`.

## Contribution

Issues are open for discussions and rest.

- [How it works?](docs/how-it-works.md)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/ufukty/gonfique.svg?variant=adaptive)](https://starchart.cc/ufukty/gonfique)

## License

Apache2. See LICENSE file.
