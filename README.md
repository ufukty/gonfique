# Gonfique

<img src="assets/Gonfique.png" alt="Gonfique logo" height="300px">

Gonfique is a special kind of YAML-to-Go and JSON-to-Go that has the **customization options** developers need when they create mappings for config files. Gonfique also **works offline**. Unlike online services Gonfique is easier to integrate into build pipeline which makes effortless to keep mapping types always up-to-date.

Having Gonfique integrated into the build pipeline, developers can use extremely dynamic schemas like storing part of the config information in the keys. Dynamic keys are breeze to work with, as they make accessing particular entry a.dot.access.close. Before Gonfique, an update in the source file would need developer to open the online service and regenerate the mapping file. With Gonfique, _as the mapping file gets updated_, the LSP checks whole codebase at instant and IDE points to the files where a previously working config access went broken. So, the developer gets a chance to fix before prod.

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
  - [Download](#download)
  - [Install from source](#install-from-source)
- [Build](#build)
- [CLI usage](#cli-usage)
  - [Generation](#generation)
  - [Version](#version)
  - [Help](#help)
- [Features](#features)
- [Gonfique config puns&funs](#gonfique-config-punsfuns)
  - [Rules](#rules)
    - [Writing paths for value-targeting rules](#writing-paths-for-value-targeting-rules)
      - [Wildcards](#wildcards)
      - [Components](#components)
    - [Writing typenames for type-targeting rules](#writing-typenames-for-type-targeting-rules)
    - [Directives](#directives)
      - [Creating named separate type declarations with auto generated names with export](#creating-named-separate-type-declarations-with-auto-generated-names-with-export)
      - [Creating named separate type declarations with declare](#creating-named-separate-type-declarations-with-declare)
      - [Assigning types manually with replace](#assigning-types-manually-with-replace)
      - [Using Go maps for YAML dictionaries and JSON objects with dict](#using-go-maps-for-yaml-dictionaries-and-json-objects-with-dict)
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

Since Gonfique needs the version information to be embedded into the binary, you need to compile it with necessary flags. The Makefile in the root of project contains a recipe that conforms the criteria.

Version information is important because Gonfique stamps the generated file with it. Which is to allow your colleagues to reproduce the same results by getting the correct version of Gonfique in future.

### Download

Simply download the Gonfique binary compiled for your operating system and architecture. Then put it into some directory listed in the `$PATH` variable.

Make sure shell can find Gonfique:

```sh
which gonfique
```

If there is no output, then check if the directory is in `$PATH`. If the path is printed, you are ready to roll.

### Install from source

You need to clone repository and to run the Make recipe installs the binary after compiling with correct version information:

```sh
git clone github.com/ufukty/gonfique
cd gonfique
make install
```

## Build

You can build the Gonfique binaries for all architectures and operating systems at once. Clone the repository and run the Makefile recipe installs the binary after compiling with correct version information:

```sh
git clone github.com/ufukty/gonfique
cd gonfique
make build
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
    - Embedding other declared types
    - Parent refs
- Easy troubleshoot:
  - Path-directive kind mismatch
  - Directive value conflicts
  - Declaration conflicts
- Abstracts boring stuff:
  - Additive type creation for list-items and dict-values
- Attention to details:
  - Version stamping for reproducibility
  - Alphabetical sorting for minimal git diffs
- Supports JSON and YAML files

## Gonfique config (puns&funs)

Gonfique config is a YAML file which contains the customizations developer wants. Gonfique config is completely optional. If there is no need for any customization, this section is safe to skip.

Overall structure of a Gonfique config is very simple. They can contain 2 sections: `meta` and `rules`. Here is the syntax of the Gonfique config:

```yml
# Values wrapped with `<` and `>` are provided by user.

meta:
  package: <package-name>
  config-type: <typename>

rules:
  <path>:
    # customizations for value-targeting rules:
    declare: <typename>
    dict: [ struct | map ]
    export: <bool>
    replace: <typename> <import-path>

  <Typename>:
    # customizations for type-targeting rules:
    accessors: <keys...>
    embed: <typename>
    iterator: <bool>
    parent: <field-name>

.
```

### Rules

Rules section is where all the customizations are described in the Gonfique config. The whole section is a dict where the keys are individual targets and the values are sequence of one directive and its arguments.

As can be seen in the file above, directives can be set either on a type of a value that strictly corresponds to a certain point in the YAML/JSON file, or a typename that is declared by the former. You use paths to point Gonfique to a certain point in the YAML/JSON file. Then, Gonfique applies the directives in that rule to the type corresponds to that point of YAML/JSON file. At the other hand, you use typenames to point Go types previously declared by value targeting paths.

So, the rules target either **a value in the YAML/JSON file** or **a Go type** previously declared by former. Understanding the distinction between value and type targeting paths is crucial to write proper Gonfique configs.

### Paths

When you want to customize a value's type:

- to overwrite the type expression with a custom typename,
- to create a named type declaration and replace with the type expression,
- to convert the resolved type for a dict to `map` form `struct`,

those manipulations performed per target. Thus, those rules need to be defined with **value-targeting paths**.

Paths are how you describe a target you want to apply directives, when the directives are one of those that works on resolved or assigned types such as `dict`, `replace`, `declare` or `export`. Paths are consist of dot-separated terms. Order of terms follows the file hierarchy between dictionaries, lists and values. There are 4 kind of terms: `key`, `wildcard`, `component` and `type`.

| Term kind   | Description                                                                                                                             | Examples                  |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------- | ------------------------- |
| `key`       | A key of a dict that is<br>in YAML/JSON file                                                                                            | `alice`, `bob`, `charlie` |
| `wildcard`  | Matches one or multiple<br>number of nodes in<br>the key hierarchy in<br>YAML/JSON file                                                 | `*`, `**`                 |
| `component` | One of the 3 special<br>symbols to went from<br>a container's<br>(dict/list) type down<br>to its resolved element,<br>key or value type | `[]`, `[key]`, `[value]`  |
| `type`      | A typename in Go which<br>is previously declared<br>by another rule<br>with first 2<br>kind of path; wrapped with<br>angle brackets     | `<Student>`               |

Depending on 'what' you want to target (a type or a value) there are different writing styles of those paths. Value targeting paths let's you customize the resolved types of those values you give their addresses in the file. At the other hand you can use type-targeting paths to customize Go types you previously "declare"d with former. The need to separate between paths targeting 'values' and paths targeting 'types' arose by some directives working on 'values' and others on 'types'. When you write a rule with value targeting path, Gonfique actually applies the directives on the _resolved types_ of matching targets values in YAML/JSON file. This is different than type-targeting paths, which targets by a Go typename rather than a YAML/JSON file path and applied on an already declared Go type. Notice that type-targeting paths terminates with a typename, so they match a type.

| Path kind | Rule                                                                                                                                                            | Example / matching targets         |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| `value`   | Use `key`, `component` or `wildcard` terms.<br>Order follows file hierarchy.<br>Path should start with a `type` term<br>if the value fall into a declared type. | `students.[]`, `<Student>.classes` |
| `type`    | Use only one `type` term.                                                                                                                                       | `<Student>`                        |

Here is an example that demonstrates the usage of 3 kind of paths below. Notice that the second rule starts with a type term which contains the name of type declared by first rule. The 3rd rule only have one segment that contains the name of type that is requested to be implement accessors on one of its fields.

<table>
<thead>
<tr>
<td>Input</td>
<td>Config</td>
</tr>
</thead>
<tbody>
<tr>
<td>

```yml
students:
  alice:
    classes:
      - name: math
        scores: [100, 100]
```

</td>
<td>

```yml
rules:
  students.[]: { declare: Student }
  <Student>.classes.[]: { declare: Class }
  <Class>: { accessors: ["name"] }
```

</td>
</tr>
</tbody>
</table>

After completing this section, look at the [examples](#creating-named-separate-type-declarations-with-declare) in `declare` section.

###### Wildcards

Use wildcards to increase flexibility of paths against partial content shifts, changes in the input file's schema which are expected to happen over time as development continues. There are 2 wildcards: single `*` and double `**` wildcards. A wildcard containing path may match multiple targets. The directives on the rule will be applied to each target, individually. Gonfique will notify if a path doesn't get any match.

The single wildcard matches each key of the current dict. So, the term contains this wildcard doesn't filter the matches further after the previous term. It makes sense to use this wildcard when you want to select every key of a dict where the types match. For example, `students.*.name` makes sense when all keys in the students dictionary have assigned values sharing same schema, in this case, dicts at least sharing a single key named `name`.

The double wildcard matches multiple (0 or more) "depths" of the file. Which also trespasses **from list type to item type** or **dict type to value type**. Using double wildcard makes more sense when you want to define precise rules on far depths of the schema without the burden of writing the whole "ancestry" of the target, which would've need maintenance more frequently, as the overall shape of the schema expected to change more frequently. So, the double wildcard is for increasing flexibility/mobility against verbosity.

<table>
<thead>
<tr>
<td>Input</td>
<td>Path examples</td>
</tr>
</thead>
<tbody>
<tr>
<td>

```yaml
students:
  alice:
    fullname: Alice Gamma
    classes:
      math:
        scores: [100, 100]
  bob:
    fullname: Bob Delta
    classes:
      math:
        scores: [100, 100]
      biology:
        scores: [100, 100]
employees:
  charlie:
    fullname: Charlie Epsilon
    DoE: 01.01.2024
```

</td>
<td>

| Path                 | List of matches                                                                                                            |
| -------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `students.alice`     | `students.alice`                                                                                                           |
| `employees.charlie`  | `employees.charlie`                                                                                                        |
|                      |                                                                                                                            |
| `*`                  | `students`, `employees`                                                                                                    |
| `*.charlie`          | `employees.charlie`                                                                                                        |
| `students.*`         | `students.alice`, `students.bob`                                                                                           |
|                      |                                                                                                                            |
| `**`                 | `students`, `students.alice`, ..., <br>`students.alice.classes.math.scores.[]`,<br>... `employees`, ... (total 20 matches) |
| `**.alice`           | `students.alice`                                                                                                           |
| `students.**.scores` | `students.alice.classes.math.scores`,<br>`students.bob.classes.math.scores`                                                |
| `students.**`        | `students.alice`, ...,<br>`students.bob.classes.biology.scores.[]`<br>(total 15 matches)                                   |

</td>
</tr>
</tbody>
</table>

###### Components

There are 3 component selectors. One of them is for narrowing down from array type to its element type `[]`. The other two selectors are `[key]` and `[value]`. Those are to narrow down from a map type to either of key type or value type. If Gonfique sees any of those selectors, it expects the container to not be a struct but a list for element type selector or a dict for key/value type selector.

```go
type ArrayType []ItemType
```

Use `[key]` operator to target a key type and use `[value]` operator to target a value type.

```go
type MapType map[KeyType]ValueType
```

Paths contain `[key]` and `[value]` are only read when the container set `dict: map`. Gonfique will assign `string` as key type unless a rule on `[key]` dictates else.

> Input file:
>
> ```yml
> students:
>   alice: { /**/ }
>   bob: { /**/ }
> ```
>
> Gonfique config:
>
> ```yml
> rules:
>   students: { dict: map }
>   students.[key]: { replace: models.StudentName import/path/to/models }
>   students.[value]: { declare: Student }
> ```
>
> Output file:
>
> ```go
> import "import/path/to/models"
>
> type Student struct { /**/ }
>
> type Config struct {
>   Students map[StudentName]models.Student `yaml:"students"`
> }
> ```

#### Directives

There are 4 types of directives that can be used within `value` and `type-value` targeting rules:

- **export**  
  Generate a separate type declaration
- **declare**  
  Like `export` but you provide the typename
- **dict**  
  Use `map` instead of `struct` for a dictionary
- **replace**  
  Overwrites the resolved type definition with the provided typename

There are 4 types of directives that can be used within `type` targeting rules:

- **accessors**  
  Implement getter and setter method for fields
- **embed**  
  Embed another type
- **iterator**  
  Make structs iterable
- **parent**  
  Add a field to struct and assign the ref of parent

The directives `declare`, `explore`, `dict` and `replace` are applied on (resolved types of) dicts, lists and literal values in the input file. Other set of directives which consists by `accessors`, `embed`, `iterator` and `parent` target types. Those directives can further customize the types produced by `declare` directives. Since those directives target types rather than paths, they are expected in `types` section.

It might be helpful imagining `type` targeting rules are processed after than the `value` and `type-value` section. Order is because of both reasons: types needed to be declared before implementing methods on them, and one type can be used for more than one target.

##### Creating named (separate) type declarations with auto generated names with `export`

```yaml
rules:
  <path>:
    export: true
```

Exporting a path, will result every matching target's type to be declared as separate with an auto generated typename. The path match multiple target is completely fine and intended. If desired, the path could be `*` or `**` too. See also: [Automatic typename generation](#automatic-typename-generation). Note that auto generated typenames are dependent to each other because of collisions and readability. So, typenames' stability subject to schema stability. Thus, consecutive runs might produce different typename set. For the typenames stability matters prefer usage of `declare` directive.

##### Creating named (separate) type declarations with `declare`

```yaml
rules:
  <path>:
    declare: <typename>
```

Use `declare` directive to generate named type declaration(s) for matching targets. This directive merges the types of all matches, and requires them to share same schema.

> ![TIP]
>
> Examples with paths and `declare` directive:
>
> ```yaml
> rules:
>   apiVersion: { declare: ApiVersion }
>   spec.template.metadata: { declare: Metadata }
>   <Metadata>.labels.app: { declare: AppLabel }
> ```
>
> Don't confuse your mind with braces.
> You can still write them multi-line if you prefer:
>
> ```yaml
> spec.template.metadata.labels.app:
>   declare: AppLabel
> ```
>
> Wildcards lets users to write more flexible mappings. Single-level wildcards match with any key in a dictionary, and they can be used many times in a pathway:
>
> ```yaml
> rules:
>   spec.*.*.labels.app: { declare: AppLabel }
>   spec.template.*.labels.app: { declare: AppLabel }
> ```
>
> Multi-level wildcards passes many times from a dict to its keys and from an array to its item type. Below would match all of the `spec.app`, `spec.foo.app` and `spec.bar.[].app` same time:
>
> ```yaml
> rules:
>   spec.**.app: { declare: AppLabel }
> ```
>
> Square brackets can be used to pass from an array to its item type only once:
>
> ```yaml
> rules:
>   spec.template.spec.containers: { declare: Containers }
>   <Containers>.[]: { declare: Container }
>   <Container>.name: { declare: ContainerName }
> ```

> [!TIP]
>
> Multiple paths can declare same name on different targets. This is useful when the willing is reducing the number of types or making the relation between different part of input file explicit in mapping.
>
> ```yaml
> rules:
>   a.b: { declare: B }
>   c.b: { declare: B }
> ```
>
> Which might make possible to write functions that takes those part of file an argument:
>
> ```go
> func UtilityFunction(b config.B) { /* */ }
> ```

##### Assigning types manually with `replace`

```yaml
rules:
  <path>:
    replace: <typename> <import-path>
```

Assign specified type name instead resolving from source file. For example: `replace: int` or `replace: models.Employee acme/models`.

##### Using Go maps for YAML dictionaries and JSON objects with `dict`

```yaml
rules:
  <path>:
    dict: [ struct | map ]
.
```

Use a Go `map` instead of a `struct` to represent a dict. This might be useful when the set of keys for a dict is not fixed until the build time, but extends to runtime. Using `map` requires all values of the dict in path to have same or compatible schemas. Otherwise Gonfique will print an error on stderr and return with status other than 0.

| Value              | Paths made available with value  | Result                |
| ------------------ | -------------------------------- | --------------------- |
| `struct` (Default) | `<dict>.<key>`                   | `struct{ ... }`       |
| `map`              | `<dict>.[key]`, `<dict>.[value]` | `map[string]combined` |

> [!NOTE]  
> Anything between angle brackets just like `<key>` is to describe the value you provide. At the other hand `[key]` is a keyword. Don't use angle brackets in real file.

If you want to get a standard map type such as a `map[string]any`, you might as well use `replace: map[string]any` at the dict's path instead. `dict` directive suits better when only the keys are map and values are fixed in build time. The advantage of using `dict: map` over `replace: map[K]V` is when you also want to customize resolved key and value types through `[map]` and `[value]`.

Note that Gonfique will use `any` value type if any two of dict key's type conflict when the dict is requested to be defined as `map`. If your key types conflict but you need to keep the type safety without making the dict iterable; you might as well use [iterator](#making-structs-iterable-with-iterator) on a struct representation.

### Types

#### Writing typenames for type-targeting rules

When you want to manipulate a previously declared "Go type" to:

- add fields to assign parent refs,
- redefine it by embedding another type,
- implementing getters, setters on it,
- implementing iterator on it

those manipulations performed independently than the how many places the type you manipulate mentioned in the main mapping type. Because of multiple fields can use same type, editing a type declaration once effects all usages in the main mapping type. Thus, you need to write those rules with **type-targeting paths** rather than value-targeting paths.

#### Directives

##### Implementing getters and setters with `accessors`

```yaml
rules:
  <Typename>:
    accessors: [<key-1>, <key-2>, ...]
```

Accessors are getters and setters for fields. Gonfique can implement getters and setters on any field of a struct, any key of a dict. The code will contain input and output parameter types that is nicely matching the field type.

##### Making the hierarchy of types explicit with `embed`

```yaml
rules:
  <Typename>:
    embed: <typename>
```

Using `embed` directive will modify the generated type definition to make it look like it is derived from an embedded type. The resulting field list won't contain common fields that is also found in the embedded type. The embedded type should be amongst types generated with `declare` directive.

##### Making structs iterable with `iterator`

```yaml
rules:
  <Typename>:
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
rules:
  <Typename>:
    parent: <Fieldname>
```

Using `parent` adds a field to a declared type. The field name will be `<Fieldname>`. The `ReadConfig` function will be added necessary assignment statements to perform reference assignment after decoding completed and before function returns. Adding refs might be useful when the data defines an hierarchy where a traceback from a child to root is needed. The type of parent field will be assigned as `any` to eliminate type conflicts might raise from different users of the type have different type of parents.

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
        rules:
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

rules:
  apiVersion: { declare: ApiVersion }
  metadata.name: { declare: Name }
  spec.rules.[]: { declare: Rule }
  <Rule>.http.paths.[]: { declare: Path }
  spec.ports.[]: { declare: Port }
  spec.**.containers: { declare: SpecContainers }
  <SpecContainers>.[]: { declare: SpecContainer }
  <ContainerName>.name: { declare: ContainerName }
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

- Decide on language between JSON and YAML
- Decode the file into a `map[string]any` with the language's decoder
- Get the `reflect.Type` for `map`
- Transform `reflect.Type` -> `ast.TypeSpec`
- Traverse from top to children in DFS by keeping track of path:
  - Check if `replace` set on path. If so, apply and skip.
  - Check if `dict: dynamic-keys` set on path. If so, apply and skip.
  - Check if `declare` is
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

### Sorting declarations

Gonfique sorts the generated declarations before writing to file in order to minimize Git diffs and make the file reflect the part-whole relationship from top to bottom. This is just how it is like in the C language where the code can't refer to a symbol if it is not declared yet.

The sorting process first creates a dependency graph and places the symbols in reverse order (just to reverse it again at the end). It starts to perform DFS on root which is the type `Config`, if you didn't customize it. Places every declaration into a declaration array; with one exception. The visitor, postpones the placement of a declaration until its last mention placed down. Process ends with reversing the placement, since we started from root (dependent should be after its dependencies).

This process should leave minimum version control footprint and improve the readability of code for developers.

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

- Multidocument YAML files are not supported, caused by the decoder.
- Gonfique assigns `any` when sees `null` values. Consider use of `replace` directive in such targets.

## Contribution

Issues are open for discussions and rest.

## License

Apache2. See LICENSE file.
