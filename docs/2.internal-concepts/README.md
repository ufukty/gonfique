# Internal concepts

## Pipeline

Reading input file

- Decide between the YAML/JSON decoder by looking to the file extension
- Decode the file into a `any` type variable.

The value would be in either type `map[string]any` or `[]any`.

Transforming to AST

- Get the `reflect.Type` for the variable.
- Transform `reflect.Type` into a `ast.Expr` with DFS.

This returns a type expression represented with a tree of `*ast.StructType`, `*ast.ArrayType`, and `*ast.Ident`s. Notice there is no map type mentioned in this step.

Processing Gonfique config

- Read the config
- Validate rules based only on the information present in the config file with `config/validate` package
- Apply value targeting rules:
  - Create a queue to perform BFS on type expression starting from root, with keeping track of if the visit is in forward movement or backward (backtracing).
  - In forward phase:
    - Match the value's path with rule paths (they contain wildcards)
    - If export enabled, reserve typename for path
    -
- Apply type targeting rules

Writing to file

- Collect information:
  - the type expression created from reflect type in AST,
  - named type declarations created by `declare` and `export`,
  - function declarations created by `accessors` and `iterator`
- Sort the type declarations in [dependency order](#Sorting-declarations).
- Insert function declarations in the order they will follow their receiver's type declaration.
- Use printer to print the file into memory for the first time
- Perform post process to adjust whitespaces
- Use formatter to print it to output file.

## Applying rules in BFS

Gonfique allows users to declare customizations on every node of example file. This results with some rules rely on others to be done its job in ancestor, or grandchildrens. Such as; choosing the map representation for a dictionary changes the path to children's type, and there might be occurrences where declaring multiple nodes with same typename only work without type conflicts amongst them, if the rules on one's grandchildren alter its type with another rule before declare directive has been started to process.

To address those needs, Gonfique finds the actual paths to nodes in BFS first then, applies directives in DFS backtracing. Finding actual paths with BFS starts on the root of type expression. Actual paths are what will be used for matching paths written in Gonfique config with actual nodes. This traversal also contains typename reservation step. In fact, the choice of BFS for finding paths is made because of the `export` directive. Applying of this directive requires generating typenames automatically, based on the value's path. Since choosing more-generic typenames for values closer to root aligns better with what developers would do when they write mapping type manually, the traversal performed in BFS order.

After the BFS travel lists the paths and nodes; Gonfique performs a DFS backtracing. This is a traversal in AST that starts from the leaves and progress toward the root. At each visit, the directives `declare`, `dict`, `export` and `replace` are applied.

As a design goal; a typical user should not be able to tell the existance of either BFS or forward/backward separation by observing Gonfique's behavior, unless they try.

## Automatic typename generation

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

## Combining element and value types

`gonfique` is able to assign one combined type for array element and map value types. It works best when all items of an collection possess the same schema in input file, but still works when they have combinable schemas. Type assignment occurs as below when items have not same but compatible schemas:

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

Gonfique's type combination works on multiple levels. Look at the example below. The global array's elements have different type of values assigned to `details` keys. This doesn't stop Gonfique to combine values assigned to `details` than combine the type of global array's elements. Thus, the array type assigned successfully.

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

```yaml:
- name: ""
  details:
    color: ""
- name: ""
  details:
    shape: ""
```

</td>
<td>

```go
type Details struct {
  Color string `yaml:"color"`
  Shape string `yaml:"shape"`
}

type Item struct {
  Name    string  `yaml:"name"`
  Details Details `yaml:"details"`
}

type Config []Item
```

</td>
</tr>
</tbody>
</table>

> [!IMPORTANT]
> Element and value types are assigned as `any`, if the types of example values are not combinable. Here is an example where the array elements have uncompatible types, due to the `b` is `int` in one, and `string` in another.
>
> ```yaml
> - a: ""
>   b: 0
> - a: ""
>   b: ""
> ```

## Sorting declarations

Gonfique sorts the generated declarations before writing to file in order to minimize Git diffs and make the file reflect the part-whole relationship from top to bottom. This is just how it is like in the C language where the code can't refer to a symbol if it is not declared yet.

The sorting process first creates a dependency graph and places the symbols in reverse order (just to reverse it again at the end). It starts to perform DFS on root which is the type `Config`, if you didn't customize it. Places every declaration into a declaration array; with one exception. The visitor, postpones the placement of a declaration until its last mention placed down. Process ends with reversing the placement, since we started from root (dependent should be after its dependencies).

This process should leave minimum version control footprint and improve the readability of code for developers.
