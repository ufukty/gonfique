# Gonfique config

Gonfique config is a YAML file which contains the customizations developer wants. Gonfique config is completely optional. If there is no need for any customization, this section is safe to skip.

Overall structure of a Gonfique config is very simple. They can contain 2 sections: `meta` and `rules`. Here is the syntax of the Gonfique config:

```yaml
# Values wrapped with `<` and `>` are provided by user.

meta:
  package: <package-name>
  config-type: <typename>

rules:
  <path-to-resolved-type>:
    # customizations for value-targeting rules:
    declare: <typename>
    dict: [struct | map]
    export: <bool>
    replace: <typename> <import-path>

  <declared-type>:
    # customizations for type-targeting rules:
    accessors: <keys...>
    embed: <typename>
    iterator: <bool>
    parent: <field-name>
```

## Rules

The `rules` dictionary is where you put all the customizations you want to apply on resolved types of values in the file, and other types you declare in the process. The each key-value pair in `rules` dictionary is processed as target(s) and directive(s). Targets can either be a value described by its path, or a type declared by former.

**Directives** can be set either on a value or type. In the case of value; the directives are applied on the resolved type of the value you describe its path. In case of type, the path is expected to describe a type that is declared by the former, during the generation process. There is one custom case, where the value you want to customize does directly or eventually belong to a previously declared type. In such case, the path is still categorized as value targeting, although its path expected to start with the name of containing type according to path syntax. Next section explains the terms and the syntax of paths.

So, the rules **target** either a value in the YAML/JSON file or a Go type previously declared by former. Understanding the distinction between value and type targeting paths is crucial to write proper Gonfique configs.

There are 4 types of **directives** that can be applied on resolved/assigned types of the input file values:

- **export**  
  Generate a separate type declaration with the resolved type (or the assigned type, if present)
- **declare**  
  Like `export` but you provide the typename
- **dict**  
  Use `map` instead of `struct` for a dictionary
- **replace**  
  Overwrites the resolved type definition with the provided typename

There are also 4 types of directives that can be applied on types declared during the generation process to allow you further customize them:

- **accessors**  
  Implement getter and setter method for fields
- **embed**  
  Embed another type
- **iterator**  
  Make structs iterable
- **parent**  
  Add a field to struct and assign the ref of parent

## Paths

When you want to customize a value's resolved type to apply value targeting directives such as `declare`, `dict`, `export` or `replace`; those manipulations performed per target. Thus, those rules need to be defined with **value-targeting paths**. Paths are written as a sequence of dot-separated **terms**. Order of terms follows the file hierarchy between dictionaries, lists and values. There are 4 kind of terms: `key`, `wildcard`, `component` and `type` that needs to be combined in special way to create a value or type targeting path.

### Terms

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

```yaml
students:
  alice:
    classes:
      - name: math
        scores: [100, 100]
```

</td>
<td>

```yaml
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

### Targeting resolved types

#### Wildcards

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

#### Array and map types

In Go, defining an arrays type need element type to be known, and defining a map type need the key and value types to be known. Gonfique, when it faces with a list, compares the all elements that present in the input file, to see if they have any conflicting component. If there is no conflift amongst all values, then Gonfique finds the element type by combining types of all items. This also applied on dictionaries when they are directed to be represented with a Go map instead of a Go struct (look `dict` directive). Gonfique checks all key's values for their types. If there is no conflicts, then the value type assigned as the combined type of all values. The key type is originally declared as `string`, although it can be customizable.

Customizing element, key and value types are possible with special terms called component selectors. There are 3 component selectors. One of them is for narrowing down from array type to its element type `[]`.

```go
type ArrayType []ItemType
```

The other two selectors are `[key]` and `[value]`. They are to narrow down from a map type to either of key type or value type. If Gonfique sees any of those selectors, it expects the container to not be a struct but a list for element type selector or a dict for key/value type selector.

Use `[key]` operator to target a key type and use `[value]` operator to target a value type.

```go
type MapType map[KeyType]ValueType
```

Paths contain `[key]` and `[value]` are only read when the container set `dict: map`. Gonfique will assign `string` as key type unless a rule on `[key]` dictates else.

> Input file:
>
> ```yaml
> students:
>   alice: { /**/ }
>   bob: { /**/ }
> ```
>
> Gonfique config:
>
> ```yaml
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

Component terms can be used in termination of a path to select the component type; or anywhere after a collection type to narrow down the scope of rule.

### Targeting declared types

When you want to manipulate a previously declared "Go type" to:

- add fields to assign parent refs,
- redefine it by embedding another type,
- implementing getters, setters on it,
- implementing iterator on it

those manipulations performed independently than the how many places the type you manipulate mentioned in the main mapping type. Because of multiple fields can use same type, editing a type declaration once effects all usages in the main mapping type. Thus, you need to write those rules with **type-targeting paths** rather than value-targeting paths.

## Directives

The directives `declare`, `explore`, `dict` and `replace` are applied on resolved types of values of the input file. Other set of directives which consists by `accessors`, `embed`, `iterator` and `parent` target types. Those directives can further customize the types declared by `declare` directives.

It might be helpful imagining rules that target declared types are processed after than the rules target resolved/assigned type of values in the input file. This ordering is because of both reasons:

- Types needed to be declared before implementing methods on them,
- One type can be used for more than one target.

### Directives applied on resolved types

#### Creating auto-named type declarations with `export`

```yaml
rules:
  <path-to-resolved-type>:
    export: true
```

Exporting a path, will result every matching target's type to be declared as separate with an auto generated typename. The path match multiple target is completely fine and intended. If desired, the path could be `*` or `**` too. See also: [Automatic typename generation](#automatic-typename-generation). Note that auto generated typenames are dependent to each other because of collisions and readability. So, typenames' stability subject to schema stability. Thus, consecutive runs might produce different typename set. For the typenames stability matters prefer usage of `declare` directive.

#### Creating named type declarations with `declare`

```yaml
rules:
  <path-to-resolved-type>:
    declare: <typename>
```

Use `declare` directive to generate named type declaration(s) for matching targets. This directive merges the types of all matches, and requires them to share same schema.

> [!TIP]
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

#### Overwriting resolved types with `replace`

```yaml
rules:
  <path-to-resolved-type>:
    replace: <typename> <import-path>
```

Assign specified type name instead resolving from source file. For example: `replace: int` or `replace: models.Employee acme/models`.

#### Using Go maps for dictionaries with `dict`

```yaml
rules:
  <path-to-resolved-type>:
    dict: [ struct | map
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

### Directives applied on declared types

#### Implementing getters and setters with `accessors`

```yaml
rules:
  <declared-type>:
    accessors: [<key-1>, <key-2>, ...]
```

Accessors are getters and setters for fields. Gonfique can implement getters and setters on any field of a struct, any key of a dict. The code will contain input and output parameter types that is nicely matching the field type.

<table>
<thead>
<tr>
<td>Input</td>
<td>Output</td>
</tr>
</thead>
<tbody>
<tr>
<td>

```yaml
rules:
  "**.endpoints.*":
    declare: Endpoint
  "<Endpoint>.method":
    replace: http.Method module/http
  "<Endpoint>":
    accessors: ["method"]
```

</td>
<td>

```go
import "module/http"

type Endpoint struct {
  Method http.Method `yaml:"method"`
  Path   string      `yaml:"path"`
}

func (e Endpoint) GetMethod() http.Method {
  return e.Method
}

func (e *Endpoint) SetMethod(v http.Method) {
  e.Method = v
}
```

</td>
</tr>
</tbody>
</table>

#### Making the hierarchy of types explicit with `embed`

```yaml
rules:
  <declared-type>:
    embed: <typename>
```

Using `embed` directive will modify the generated type definition to make it look like it is derived from an embedded type. The resulting field list won't contain common fields that is also found in the embedded type. The embedded type should be amongst types generated with `declare` directive.

#### Making structs iterable with `iterator`

```yaml
rules:
  <declared-type>:
    iterator: <bool>
```

Since the corresponding type for a struct-represented section of the input file is actually a dict of string keys and values; Gonfique can include additional data in the generated file to allow you access the "keys" as strings.

Combined with `iterator` directive, Gonfique let's you use your 'structs' in a previously unimagined way:

```go
for name, details := range cfg.employees.Fields() { /* */ }
```

where the employees were originally a dict and represented with a struct in Go. With iterator support on structs, you can keep your way to access values through fields like it is a `struct` and also have another way to iterate over them like it is a `map`.

The declarations of both `Objectives` type and its `Fields` method can be seen in example below. Keys returned by iterator are actual values from YAML/JSON file that correspond to a field. Value type is coming from the all fields sharing same type.

```go
type Objectives struct {
  Create Endpoint `yaml:"create"`
  Delete Endpoint `yaml:"delete"`
  Get    Endpoint `yaml:"get"`
  Patch  Endpoint `yaml:"patch"`
  Put    Endpoint `yaml:"put"`
}

func (o Objectives) Fields() iter.Seq2[string, Endpoint] {
  return func(yield func(string, Endpoint) bool) {
    mp := map[string]Endpoint{
      "create": o.Create,
      "delete": o.Delete,
      "get":    o.Get,
      "patch":  o.Patch,
      "put":    o.Put,
    }
    for k, v := range mp {
      if !yield(k, v) {
        return
      }
    }
  }
}
```

#### Adding a field for parent access with `parent`

```yaml
rules:
  <declared-type>:
    parent: <Fieldname>
```

Using `parent` adds a field to a declared type. The field name will be `<Fieldname>`. The `ReadConfig` function will be added necessary assignment statements to perform reference assignment after decoding completed and before function returns. Adding refs might be useful when the data defines an hierarchy where a traceback from a child to root is needed. The type of parent field will be assigned as `any` to eliminate type conflicts might raise from different users of the type have different type of parents.
