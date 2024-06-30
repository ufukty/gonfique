# Directive File

> ![WARNING]
> Directive file feature is currently experimental. During the experiment; its usage, behavior or existence can change or get removed without warnings.

> ![WARNING]
> Documentation might not be up to date with implementation at the moment.

A directive file is a YAML file that contains a dictionary of paths and directives. Example file contains 3 path and 4 directives for it:

```yaml
infra.servers:
  export: True

infra.servers.*:
  declare: Vps
  accessors: [Cores, Ram, Disk]
```

## Path

Path is the path of any dict-key/list-item/value in a config file defined by keys to be followed in order to access that value separated by dots. Such as, the path of `Hello world` is considered as `a.b.c` in this config file:

```yaml
a:
  b:
    c: Hello world
```

Paths can be used for various purposes when combined with directives such as: creating a named type declaration for resolved type, assigning a type to matching part, manipulate the resolved type etc.

### Wildcards

Use wildcards to increase flexibility of directives against partial content shifts, changes in the config files expected to happen over time.

There are 3 wildcards: `*`, `[]`, `**` which respectively means:

- to match **any key of the dictionary** in the current depth,
- to match **item type of the array** in the current depth,
- to match any key of the dictionary or the item type of array type in the every depth.

Details are in the [mapping](mapping.md) docs.

A wildcard-containing-path can result with multiple matches. In case of multiple matches, the directives will be applied to each match.

Gonfique will notify if a path doesn't get any match.

| Path               | Example matches                                         |
| ------------------ | ------------------------------------------------------- |
| `**.alice.bob`     | `alice.bob`, `x.alice.bob`, `x.x.alice.bob`             |
| `*.bob.charlie`    | `x.bob.charlie`, `y.bob.charlie`                        |
| `alice.**.charlie` | `alice.charlie`, `alice.x.charlie`, `alice.x.x.charlie` |
| `dave.[]`          | item type of `dave` array                               |
| `dave.[].erin`     | `erin` key in the item type of `dave` array             |
| `dave.[].*`        | Every key in the item type of `dave` array              |

Availability of each key in item types is subject to [array type defining behavior](arrays.md).

## Directives

There are 5 different directive that can be set on a path. See explanations for conflicting directives.

```yaml
a.key.path:
  accessors: ...
  declare: ...
  embed: ... # planned
  export: ...
  parent: ...
  replace: ... # planned
```

### `declare`

```yaml
a.key.path:
  declare: Typename
```

Use `declare` directive to generate named type declaration(s) for matching targets. This directive merges the types of all matches, and requires them to share same schema. There can be multiple rules mentioning same typename in `declare` directive.

**Notes**

- `declare` can be combined with `replace`
- `declare` overrides `export`

### `export`

```yaml
a.key.path:
  export: true # default is false
```

Setting `export` to true will result with automatically generated typenames to be exported, meanining is that they'll start with a capitalized letter. This will only have effect when [automatic type declaration](#automatically-deciding-to-generate-type-declarations) gets triggered.

**Notes**

- Typenames are dependent to each other because of collisions and readability. So, typenames' stability subject to schema stability.
- Exporting doesn't need to merge the types of all matches. So, a rule can match targets in different schemas and set exporting to true, such as `**`.

### `accessors`

```yaml
a.key.path:
  accessors: [key-name-1, key-name-2, ...]
```

Accessors are getters and setters for fields. Gonfique can implement getters and setters on any field of a struct, any key of a dict. The code will contain input and output parameter types that is nicely matching the field type.

**Notes**

- Accessors will be defined on all types the rule matches. Define paths that will only match same type targets.
- Multiple rules matching same target containing conflicting directives is illegal, as well as, one rule match with different type targets.

### `embed`

> [!NOTE]
> This directive is currently here for preview and unavailable for use.

```yaml
a.key.path:
  embed:
    typename: Typename
    import-path: path/to/package/to/import
    import-as: packageAlias
```

Using `embed` directive will modify the generated type definition to make it look like it is derived from an embedded type. The resulting field list won't contain common fields with embedded struct.

Use `import-path` if the embedded type is outside of package specified with CLI flag. Also use `import-as` when an alias is desired for imported package.

**Notes**

- Embedded type should be a struct, not an interface.

### `parent`

```yaml
a.key.path:
  parent: Fieldname
```

Using `parent` adds a field to generated type. The field name will be `fieldname` and its value will be the reference of its `level`th level of parent. Adding refs may be useful when the data defines an hierarchy a traceback from a child to root is needed.

**Notes**

- Adding parent refs to structs as fields requires the type of parent to be mentioned in type definition; so type's reusability gets limited to targets with same type parents.
- Combining `parent` and `declare` may result with failure when parent types differ.
- Adding parent refs alters the body of ReadConfig function, as the refs need to be assigned after initialization.

### `replace`

> [!NOTE]
> This directive is currently here for preview and unavailable for use.

```yaml
a.key.path:
  typename: Typename
  import-path: path/to/package/to/import
  import-as: packageAlias
```

Assign specified type name instead resolving from source file. For example: `type: int`

**Notes**

- Use `import-path` if the embedded type is outside of package specified with CLI flag. Also use `import-as` when an alias is desired for imported package.
- When `declare` and `replace` is combined, generated file will contain a type declaration with underlying type is the specified type.

## Internal Concepts

### Automatic type resolution vs. manual type assignment

Gonfique can resolve any key/list/value's type by simply looking to it. While this behaviour is the default, Gonfique users can choose to opt-out automatic type resolution for any dict/list/value in the config file.

When type resolution disabled by using `replace` directive on any dict/list, Gonfique won't apply any directives for their "children" (that is all dicts, lists and values eventually belong to that object, subtree).

### Deciding to generate type declarations

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

## Troubleshoot

### Combining `parent` and `declare` on a group of matches

It might not be obvious to everyone at first thought; but when parent and named is set together on a group of target, parents of those targets need to be in the same type. Otherwise, you want Gonfique to produce invalid Go code. Because adding parent fields to struct definitions alter their types in a way they end-up being exclusive to one parent type.

When both directives set together on a group of matches, make sure parents of matches are in same type. If they are not; either use separate rules to define different names for conflicting matches. Or, let Gonfique to generate unique typenames by **not using** `declare` directive. See `exported` directive if all is needed is to access type name from outside package and the typename itself is arbitrary.
