# Directive File

> ![WARNING]
> Directive file feature is currently experimental. During the experiment; its usage, behavior or existence can change or get removed without warnings.

A directive file is a YAML file that contains a dictionary of keypaths and directives. Example file contains one keypath and 4 directives for it:

```yaml
infra.servers:
  export: True

infra.servers.*:
  named: Vps
  embed: Basic
  accessors: [Cores, Ram, Disk]

"*.**":
  parent: Parent
```

## Path

Path is the path of any dict-key/list-item/value in a config file defined by keys to be followed in order to access that value separated by dots. Such as, the keypath of `Hello world` is considered as `a.b.c` in this config file:

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

A wildcard-containing-keypath can result with multiple matches. In case of multiple matches, the directives will be applied to each match.

Gonfique will notify if a keypath doesn't get any match.

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

There are 6 different directive that can be set on a keypath. See explanations for conflicting directives.

```yaml
a.key.path:
  named: TypeName
  accessors: [FieldName, FieldName, ...]
  parent: FieldName
  embed: TypeName
  import: PackageName
  type: TypeName
```

### `named`

```yaml
a.key.path:
  named: TypeName
```

Create a named type instead inline type with the type definition resolved from config file. Eg. `named: Employee`

### `accessors`

```yaml
a.key.path:
  accessors: [FieldName, FieldName, ...]
```

Accessors are getters and setters for fields. Gonfique can implement getters and setters on any field of a struct. The code will contain input and output parameter types that is nicely matching the field type.

Note; accessors will be defined on all types the rule matches. Multiple rules matching same target containing conflicting directives is illegal.

### `embed`

```yaml
a.key.path:
  embed: TypeName
```

Use embed directive to make gonfique to define keys for the struct without the fields in embedded struct. Embedded type should be a struct, not an interface. TypeName could be either declared inside or [outside](#import) the package.

### `parent`

```yaml
a.key.path:
  parent: FieldName
```

Add a field `FieldName` to the generated type which will be assigned the pointer of parent, `a.key`. This will also change the body of ReadConfig function. This will be useful when the data defines a hierarchy that a traceback from a child to root is needed.

Note: If a `parent` containing wildcard keypath gets multiple matches; different type parents will needed to be declared with different typenames. So, combining `parent` with `named` may result with failure.

### `export`

```yaml
a.key.path:
  export: True/False
```

Directs [automatic type name generation](#automatically-generated-type-names) to generate exported (capitalized) type names for the match of path or each match of wildcard path. This has no effect when `named` or `type` is also set.

### `import`

```yaml
a.key.path:
  import: PackageName
```

Adds the package name (or path) into import list that will be in the top of generated file. The package will be imported only if the rule gets any match in the config file. Usefull when combined with `type`, `embed` to refer to types outside of package.

### `type`

```yaml
a.key.path:
  type: TypeName
```

Assign specified type name instead resolving from YAML file. For example: `type: int` or `type: http.Method`. Note that, `type` directive can only be combined with `import`.

## Internal Concepts

### Automatic type resolution vs. manual type assignment

Gonfique can resolve any key/list/value's type by simply looking to it. While this behaviour is the default, Gonfique users can choose to opt-out automatic type resolution for any dict/list/value in the config file.

When type resolution disabled by using `type` directive on any dict/list, Gonfique won't apply any directives for their "children" (that is all dicts, lists and values eventually belong to that object).

### Automatically generated type names

Gonfique needs to move inline type definitions of field/item types to named type definitions in order to implement methods on them (or refer to them in other contexts in general).

Some of the cases that gonfique automatically create a type name for a field/item type:

- Implementing accessors needs Gonfique to refer type names of struct and field. So `accessor` enables automatic type name generation on the field and struct type.

- Using `parent` (on child) without `named` (on the parent) will result with gonfique assign an automatically generated type name for the parent type.

The name will be based on the keypath, the minimum number of last segments that won't collide with other typenames. As the choosen name is bound to context, it can change next time the config file gets a key with same name. Thus, the generated type name is unexported.

### Resolving Conflicts

- **When multiple keypaths match with same target:**
  - Returns error, if two rules contains conflicting directives.
  - Combines directives, if none conflicts.
- **When one keypath match with multiple targets**
  - Returns error, if not all of the targets are not sharing same type
- **When one keypath match a target in addition to another which its type or type name is provided by user at another rule**

## Full example

Config file:

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

Directive file:

```yaml
# export each auto generated type name by default
"**":
  exported: true

# add a '.Parent' field to each struct type that is under 'spec' key
spec.**:
  parent: Parent

# make the item type of ports list named 'Port'
spec.ports.[]:
  named: Port

# assign the type 'Protocol' to 'protocol' field of item type
spec.ports.[].protocol:
  type: Protocol
```

Output:

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

## Troubleshoot

### Combining `parent` and `named` on a group of matches

It might not be obvious to everyone at first thought; but when parent and named is set together on a group of target, parents of those targets need to be in the same type. Otherwise, you want Gonfique to produce invalid Go code. Because adding parent fields to struct definitions alter their types in a way they end-up being exclusive to one parent type.

When both directives set together on a group of matches, make sure parents of matches are in same type. If they are not; either use separate rules to define different names for conflicting matches. Or, let Gonfique to generate unique typenames by **not using** `named` directive. See `exported` directive if all is needed is to access type name from outside package and the typename itself is arbitrary.
