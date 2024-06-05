# Directive File

> ![WARNING]  
> Directive file feature is currently experimental. During the experiment; its usage, behavior or existence can change or get removed without warnings.

A directive file is a YAML file that contains a dictionary of keypaths and directives. Example file contains one keypath and 4 directives for it:

```yaml
infra.servers.*:
  named: Vps
  embed: Basic
  parent: Servers
  accessors: [Cores, Ram, Disk]
```

## Keypath

Keypaths is the path of any value in a config file defined by keys to be followed in order to access that value separated by dots. Such as, the keypath of `Hello world` is considered as `a.b.c` in this config file:

```yaml
a:
  b:
    c: Hello world
```

### Wildcards

Use wildcards to increase flexibility of directives against partial content shifts, changes in the config files expected to happen over time.

There are 3 wildcards: `*`, `[]`, `**` which respectively means:

- to match **any key of the dictionary** in the current depth,
- to match **item type of the array** in the current depth,
- to match any key of the dictionary or the item type of array type in the every depth.

Details are in the [mapping](mapping.md) docs.

A wildcard-containing-keypath can result with multiple matches. In case of multiple matches, the directives will be applied to each match.

Gonfique will notify if a keypath doesn't get any match.

| Keypath            | Example matches                                         |
| ------------------ | ------------------------------------------------------- |
| `**.alice.bob`     | `alice.bob`, `x.alice.bob`, `x.x.alice.bob`             |
| `*.bob.charlie`    | `x.bob.charlie`, `y.bob.charlie`                        |
| `alice.**.charlie` | `alice.charlie`, `alice.x.charlie`, `alice.x.x.charlie` |
| `dave.[]`          | item type of `dave` array                               |
| `dave.[].erin`     | `erin` key in the item type of `dave` array             |
| `dave.[].*`        | Every key in the item type of `dave` array              |

Availability of each key in item types is subject to [array type defining behavior](arrays.md).

## Directives

| Directive                        | Function                                                                                                                                                                |
| -------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <nobr>`named: TypeName`</nobr>   | Create a named type instead inline type with the type definition resolved from config file. Eg. `named: Employee`                                                       |
| <nobr>`type: TypeName`</nobr>    | Assign specified type name instead resolving from YAML file. Eg. `type: http.Method`                                                                                    |
| <nobr>`parent: FieldName`</nobr> | Assign a field which is named as `FieldName` that it's value will get set as the pointer of parent (containing) struct in the ReadConfig function. Eg. `parent: Parent` |
| <nobr>`embed: TypeName`</nobr>   | Defined type will contain `TypeName` as embedded struct. Eg. `embed: Base`                                                                                              |

Note that combining `type` directive with many other directives is not possible as explained in the next section.

### `type` directive

Types are either "assigned" manually or "resolved" by looking to the value in config file. Default behavior is automatic resolution. So, gonfique inspects the value in config file.

To change the behavior, to force assigning a type of your choice, use `type: TypeName` directive. For example: `type: int` or `type: http.Method`.

Limitations:

- Combining `type` directive with either of `parent`, `embed` or `accessors` is not allowed.

- When a dict/list type is "assigned", no other directives are allowed on keys/item of that dict/list.

### `accessors` directive

Accessors are Getters and Setters for fields. To make gonfique to implement accessors on desired keys of a dict, use `accessors` directive as such:

```yaml
organization.employees.*:
  accessors:
```

### `parent` directive

Use `parent: FieldName` directive to add a reference of parent to child. This will be useful when the data defines a hierarchy that a traceback from a child to root is needed.

Considering the config file...

```yaml
eve:
  frank: ...
```

...this directive file...

```yaml
eve:
  named: Eve

eve.frank:
  named: Frank
  parent: MyParent
```

...will result with the type `Frank` getting a field named `MyParent` and in type of `*Eve`. Also, `ReadConfig` function will get a statement added, which will assign the reference of parent to child:

```go
type Eve struct {
  Frank Frank
}

type Frank struct {
  MyParent *Eve // notice
}

func ReadConfig() (Config, error) {
  ...
  cfg.A.Eve.Frank.MyParent = &cfg.A.Eve // notice
  ...
}
```

## Full example

Given the config file:

```yaml
a:
  b:
    c: 2.0
    d: 0.5
  e:
    f:
      g: dolor
      h: consectetur
i:
  - j: vusce1
    k: nam
  - j: vusce2
    l: ac
```

and directive file:

```yaml
a.b:
  named: Blowfish
a.b.c:
  named: C
a.b.d:
  named: C

a.e:
  named: Eve

a.e.f:
  named: Frank
  parent: Parent
```

```

```
