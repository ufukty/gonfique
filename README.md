# Gonfique

<img src="assets/Gonfique.png" alt="Gonfique logo" height="300px">

Gonfique is a special kind of YAML-to-Go and JSON-to-Go that has the **customization** options developers need when they create mappings for config files. Gonfique also works **offline** unlike online services which makes Gonfique easier to integrate into build process and always keep mapping types up-to-date.

Having Gonfique integrated into the build pipeline, developers can use extremely dynamic schemas like storing part of the config information in the keys. Dynamic keys are breeze to work with, as they make accessing to particular entry a.dot.access.close. Before Gonfique, an update in the source file would need developer to open the online service and regenerate the mapping file. With Gonfique, _as the mapping file gets updated_, the LSP checks whole codebase at instant and IDE points to the files where a previously working config access went broken. So, the developer gets a chance to fix before prod.

> ```go
> cfg.the.["road"].to.["panics"].is.["paved"].with.["hardcoded"]["strings"]
> ```
>
> — a Gopher

> ```go
> if r, ok := cfg.the.["road"]; ok {
>   if p, ok := r.to.["panics"]; ok {
>     if p2, ok := p.is.["paved"]; ok {
>       if h, ok := p2.with.["hardcoded"]; ok {
>         if s, ok := h.["strings"]; ok {
> ```
>
> — a wise Gopher

> ```go
> cfg.the.road.to.panics.is.paved.with.hardcoded.strings
> ```
>
> — A. Gopherstein

<details>
<summary>If you've never adopted a JSON-to-Go or YAML-to-Go; if this is first time you met one; if you are asking yourself why would not you write mapping types by hand...</summary>
...then I dare you to find mistakes in this mapping type:

<table>
<tr>
<td>

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

</td>
<td>

```go
type Endpoint struct {
  Name   string `yaml:"name"`
  Method string `yaml:"method"`
  Path   string `yaml:"path"`
}

type Tags struct {
  Path      string     `yaml:"path"`
  Endpoints []endpoint `yaml:"endpoint"`
}

type Service struct {
  Tags Tags `yaml:"tags"`
}

type Config struct {
  Github struct {
    Domain   int `yaml:"domain"`
    Path     int `yaml:"path"`
    Services []Service
  } `yaml:"github.com"`
}
```

</td>
</tr>
</table>

This one was an easy one. No one have enough time to deal with this in repeat. You should find at least 6 mistakes depending on how you count.

</details>

## TOC

- [Install](#install)
- [CLI usage](#cli-usage)
  - [Generation](#generation)
  - [Version](#version)
  - [Help](#help)
- [Features](#features)
- [Gonfique config](#gonfique-config-punsfuns)
  - [Paths section](#paths-section)
    - [Writing paths](#writing-paths)
      - [Wildcards](#wildcards)
      - [Arrays](#arrays)
    - [Path directives](#path-directives)
      - [Creating named separate type declarations with auto generated names with export](#creating-named-separate-type-declarations-with-auto-generated-names-with-export)
      - [Creating named separate type declarations with declare](#creating-named-separate-type-declarations-with-declare)
      - [Assigning types manually with replace](#assigning-types-manually-with-replace)
  - [Types section](#types-section)
    - [Type directives](#type-directives)
      - [Implementing getters and setters with accessors](#implementing-getters-and-setters-with-accessors)
      - [Making the hierarchy of types explicit with embed](#making-the-hierarchy-of-types-explicit-with-embed)
      - [Making structs iterable with iterator](#making-structs-iterable-with-iterator)
      - [Adding a field for parent access with parent](#adding-a-field-for-parent-access-with-parent)
        - [Notes](#notes)
- [Full examples](#full-examples)
  - [With customization](#with-customization)
  - [Without customization](#without-customization)
- [Internal Concepts](#internal-concepts)
  - [Pipeline](#pipeline)
  - [Automatic type resolution vs. manual type assignment](#automatic-type-resolution-vs-manual-type-assignment)
  - [Decision process to generate type declarations](#decision-process-to-generate-type-declarations)
  - [Automatic typename generation](#automatic-typename-generation)
  - [Decision process on array type](#decision-process-on-array-type)
- [Serving suggestions](#serving-suggestions)
- [Troubleshoot](#troubleshoot)
  - [Combining parent and declare on a group of matches](#combining-parent-and-declare-on-a-group-of-matches)
- [Considerations](#considerations)
- [Contribution](#contribution)
- [License](#license)

## Install

Use version tags to avoid development versions as such:

```sh
go install github.com/ufukty/gonfique@v1.5.3
```

## CLI usage

### Generation

```sh
gonfique generate -in <path> -out <path> [ -config <path> ]
```

For basic usage without any customization config flag is not needed.

### Version

```sh
gonfique version
```

### Help

```sh
gonfique help
```

Better keep reading

## Features

- Works offline:
  - Private
  - Nicely integrates to build pipeline
- Rich customization:
  - Named or inline types
  - Auto generated or user provided typenames
  - Type replacement
  - Map/struct option for dicts
  - Implements methods on declared types:
    - Iterators
    - Accessors (getter/setter)
  - Enriches declared types:
    - Embedding with other declared types
    - Parent refs
- Abstracts boring stuff:
  - Smart array support that concludes on one common or combined item type
  - Recognizes `time.Duration` values
- Supports JSON and YAML files

## Gonfique config (puns&funs)

Gonfique config is a YAML file which contains the customizations developer wants. Gonfique config is completely optional. If there is no need for any customization, this section is safe to skip.

Gonfique accepts most of the customizations on resolved types and derived types. Beside that there is also couple small things we gonna talk in `meta` section.

Overall structure of a Gonfique config is very simple. They can contain 3 sections: `meta`, `paths` and `types`.

```yml
# Values wrapped with `<` and `>` are provided by user.

meta:
  package: <package-name>
  config-type: <typename>

paths:
  <path>: [ export | declare | replace | map-values ] <args...>

types:
  <typename>:
    accessors: <keys...>
    embed: <typename>
    iterator: <bool>
    parent: <field-name>
```

The section `paths` is processed earlier than the `types` in the generation process. The types which are declared by the `declare` directives in the `paths` section can be further customized by directives in the `types` section. Keep reading for the explanation of `declare`. If none of the entries in the `path` section contain `declare` directive on any path, than there is no need for `types` section.

### Paths section

Which is a dict where the keys are individual paths and the values are sequence of one directive and its arguments. Paths are strings that matches one or more target in the input file.

#### Writing paths

Paths are written in a special yet simple syntax in the form of dot-separated sequence of keys and square brackets. Keys are for selecting a key from the dictionary, brackets are for passing from an array to the item (resolved item type). Each path expected to match one or more sections in the input file. Gonfique checks matches and warns the developer if a path matches no target in the input file.

In the example below, the path of `alpha.beta.charlie` resolves to the string `Hello world`.

```yml
alpha:
  beta:
    charlie: Hello world
```

After completing this section, look at the [examples](#creating-named-separate-type-declarations-with-declare) in `declare` section.

##### Wildcards

Use wildcards to increase flexibility of directives against partial content shifts, changes in the config files which are expected to happen over time.

There are 2 wildcards: `*`, `**`. The first matches every key of the current dict. Second one matches multiple (0+) levels of keys. Second one also passes from arrays to item types. A wildcard containing path may result with multiple matches. The directives on path will be applied to each match of path.

Gonfique will notify if a path doesn't get any match.

| Path     | Example matches                |
| -------- | ------------------------------ |
| `a.b`    | `a.b`                          |
| `*.a.b`  | `x.a.b`                        |
| `**.a.b` | `a.b`, `x.a.b`, `x.x.a.b`, ... |
| `a.**.b` | `a.b`, `a.x.b`, `a.x.x.b`, ... |
| `*`      | `x`, `y`, `z`, ...             |
| `**`     | `x`, `x.x`, `x.x.x`, ...       |

##### Arrays

Arrays can be given directives too. But there is a separation between an array's type and its element type.

```go
type ArrayType []ItemType
```

If there is a pair of square brackets like `[]`, then Gonfique expects to see an array in the target in input file.

| Path     | Matches                   | Consequences                                                                              |
| -------- | ------------------------- | ----------------------------------------------------------------------------------------- |
| `a.[]`   | the `a` array's item type | `a` must be an array                                                                      |
| `a.[].*` | the item type's every key | `a` must be an array, item type of `a` must be a dict                                     |
| `a.[].b` | the item type's `b` key   | `a` must be an array, item type of `a` must be a dict, the dict must have a key named `b` |

#### Path directives

There are 4 alternative directives. the first 3 directives also can be used as an argument to the 4th

- **export**  
  Generates a separate type declaration for the resolved type with the shortest name based on the path
- **declare**  
  Like `export` but the user specify the name for the declaration
- **replace**  
  Overwrites the resolved type definition with the provided typename
- **map**  
  Implements the dict with a map instead of a struct. Accepts the argument as a directive on value type

```yml
paths:
  <path>: export

  <path>: declare <typename>

  <path>: replace <existing-typename> <import-path>

  <path>: map export
  <path>: map declare <typename>
  <path>: map replace <existing-typename> <import-path>
```

##### Creating named (separate) type declarations with auto generated names with `export`

```yaml
paths:
  <path>: export
```

Exporting a path, will result every matching target's type to be declared as separate with an auto generated typename. The path match multiple target is completely fine and intended. If desired, the path could be `*` or `**` too. See also: [Automatic typename generation](#automatic-typename-generation). Note that auto generated typenames are dependent to each other because of collisions and readability. So, typenames' stability subject to schema stability. Thus, consecutive runs might produce different typename set. For the typenames stability matters prefer usage of `declare` directive.

##### Creating named (separate) type declarations with `declare`

```yaml
paths:
  <path>: declare <typename>
```

Use `declare` directive to generate named type declaration(s) for matching targets. This directive merges the types of all matches, and requires them to share same schema. There can be multiple rules mentioning same typename in `declare` directive.

> ![TIP]
>
> Examples:
>
> ```yaml
> paths:
>   apiVersion: declare ApiVersion
>   metadata.name: declare Name
>   spec.template.metadata.labels.app: declare AppLabel
> ```
>
> Wildcards lets users to write more flexible mappings. Single-level wildcards match with any key in a dictionary, and they can be used many times in a pathway:
>
> ```yaml
> paths:
>   spec.*.*.labels.app: declare AppLabel
>   spec.template.*.labels.app: declare AppLabel
> ```
>
> Multi-level wildcards passes many times from a dict to its keys and from an array to its item type. Below would match all of the `spec.app`, `spec.foo.app` and `spec.bar.[].app` same time:
>
> ```yaml
> paths:
>   spec.**.app: declare AppLabel
> ```
>
> Square brackets can be used to pass from an array to its item type only once:
>
> ```yaml
> paths:
>   spec.template.spec.containers: declare Containers
>   spec.template.spec.containers.[]: declare Container
>   spec.template.spec.containers.[].Name: declare ContainerName
> ```

##### Assigning types manually with `replace`

```yaml
paths:
  <path>: replace <typename> <import-path>
```

Assign specified type name instead resolving from source file. For example: `replace int` or `replace Employee acme/models`.

### Types section

Developer can customize types declared by Gonfique. Customization list includes implementing methods on types and mutating the field list of generated struct.

Types section is a dictionary where keys are individual typenames and values are a dict of directives and their parameters.

#### Type directives

There are 4 types of type directives: `accessors`, `embed`, `iterator`, `parent`.

##### Implementing getters and setters with `accessors`

```yaml
types:
  <typename>:
    accessors: [<key-1>, <key-2>, ...]
```

Accessors are getters and setters for fields. Gonfique can implement getters and setters on any field of a struct, any key of a dict. The code will contain input and output parameter types that is nicely matching the field type.

##### Making the hierarchy of types explicit with `embed`

```yaml
types:
  <typename>:
    embed: <typename>
```

Using `embed` directive will modify the generated type definition to make it look like it is derived from an embedded type. The resulting field list won't contain common fields that is also found in the embedded type. The embedded type should be amongst types generated with `declare` directive.

##### Making structs iterable with `iterator`

```yaml
types:
  <typename>:
    iterator: <bool>
```

Since the corresponding type for a struct-represented section of the input file is actually a dict of string keys and values; Gonfique can include additional data in the generated file to allow you access the "keys" as strings.

Combined with `iterator` directive, Gonfique let's you use your 'structs' in a previously unimagined way:

```go
for name, details := range cfg.employees { /* */ }
```

where the employees were originally a dict and represented with a struct in Go. With iterator support on structs, you can keep your way to access values through fields like it is a `struct` and also have another way to iterate over them like it is a `map`.

##### Adding a field for parent access with `parent`

```yaml
<typename>:
  parent: Fieldname
```

Using `parent` adds a field to generated type. The field name will be `fieldname` and its value will be the reference of its `level`th level of parent. Adding refs may be useful when the data defines an hierarchy a traceback from a child to root is needed.

###### Notes

- Adding parent refs to structs as fields requires the type of parent to be mentioned in type definition; so type's reusability gets limited to targets with same type parents.
- Combining `parent` and `declare` may result with failure when parent types differ.
- Adding parent refs alters the body of ReadConfig function, as the refs need to be assigned after initialization.

## Full examples

Example input file is like this:

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

### With customization

For this run, Gonfique has been provided a Gonfique config:

```yml
meta:
  package: config
  config-type: Kubernetes

paths:
  apiVersion: declare ApiVersion
  metadata.name: declare Name
  spec.rules.[]: declare Rule
  spec.rules.[].http.paths.[]: declare Path
  spec.ports.[]: declare Port
  spec.**.containers: declare SpecContainers
  spec.**.containers.[]: declare SpecContainer
  spec.**.containers.[].name: declare ContainerName
```

Output:

```go
package config

import (
  "fmt"
  "os"

  "gopkg.in/yaml.v3"
)

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

type Kubernetes struct {
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

func ReadKubernetes(path string) (Kubernetes, error) {
  file, err := os.Open(path)
  if err != nil {
    return Kubernetes{}, fmt.Errorf("opening config file: %w", err)
  }
  defer file.Close()
  c := Kubernetes{}
  err = yaml.NewDecoder(file).Decode(&c)
  if err != nil {
    return Kubernetes{}, fmt.Errorf("decoding config file: %w", err)
  }
  return c, nil
}
```

### Without customization

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

## Internal Concepts

### Pipeline

- Decode: `file` -> `map[string]any`
- Transform: `reflect.Type` -> `ast.TypeSpec`
- Substitude: replace types matching with user provided types
- Mapping: match user-provided paths and separate type expressions as type specs named as instructed by user
- Organize: separate the type definitions as standalone type specs and reuse them when definitions match
- Iterables: implements Range method on those dictionaries that all items are in same type

### Automatic type resolution vs. manual type assignment

Gonfique can resolve any key/list/value's type by simply looking to it. While this behaviour is the default, Gonfique users can choose to opt-out automatic type resolution for any dict/list/value in the config file.

When type resolution disabled by using `replace` directive on any dict/list, Gonfique won't apply any directives for their "children" (that is all dicts, lists and values eventually belong to that object, subtree).

### Decision process to generate type declarations

Gonfique needs to generate named type declarations in order to implement methods on them, or refer to them in other contexts in general.

Structs matching any criteria below will get its type declared automatically, if not already requested with `declare`:

- Contains a field directed to implement accessors on,
- Contains a field which needs a `parent` ref in its type definition.

### Automatic typename generation

Gonfique will generate arbitrary typenames as needed. The name will be based on the path, the minimum number of last segments that won't collide with other typenames. As the generated typenames are depending to each other, they may change next time the config file gets a key with same name. So, they are instable for schema changes. For example:

```yaml
lorem:
  dolor: ...
  ipsum:
    dolor: ...
    sit: ...
  sit: ...
```

| Target path         | Generated typename |
| ------------------- | ------------------ |
| `lorem.dolor`       | `dolor`            |
| `lorem.ipsum.dolor` | `ipsumDolor`       |
| `lorem.ipsum.sit`   | `ipsumSit`         |
| `lorem.ipsum`       | `ipsum`            |
| `lorem.sit`         | `sit`              |

### Decision process on array type

`gonfique` assignes the necessary slice type to arrays. It works best when all items of an array possess the same schema or at least compatible schemas. Type assignment occurs as below when items have not same but compatible schemas:

<table>
<thead>
<td>Input</td>
<td>Output</td>
<thead>
<tr>
<td>

```yaml
- a: ""
  b: ""
- a: ""
  c: 0
```

</td>
<td>

```go
[]struct {
  A string
  B string
  C int
}
```

</td>
<tr>
<table>

> [!IMPORTANT]
> Slice type gets defined as `[]any` if shared keys have different type values. Like `detail` has given `int` and `string` values below:
>
> ```yaml
> - a: ""
>   b: 0
> - a: ""
>   b: ""
> ```

## Serving suggestions

For existing Makefile users:

```Makefile
config.go: config.yml gonfique.yml
    gonfique -in config.yml -out config.go -config gonfique.yml

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

## Troubleshoot

### Combining `parent` and `declare` on a group of matches

It might not be obvious to everyone at first thought; but when parent and named is set together on a group of target, parents of those targets need to be in the same type. Otherwise, you want Gonfique to produce invalid Go code. Because adding parent fields to struct definitions alter their types in a way they end-up being exclusive to one parent type.

When both directives set together on a group of matches, make sure parents of matches are in same type. If they are not; either use separate rules to define different names for conflicting matches. Or, let Gonfique to generate unique typenames by **not using** `declare` directive. See `exported` directive if all is needed is to access type name from outside package and the typename itself is arbitrary.

## Considerations

- Multidocument YAML files are not supported.
- Gonfique assigns `any` when sees `null` values.

## Contribution

Issues are open for discussions and rest.

- [How it works?](docs/how-it-works.md)

## License

Apache2. See LICENSE file.
