**Pipeline**

- Decode: `file` -> `map[string]any`
- Transform: `reflect.Type` -> `ast.TypeSpec`
- Substitude: replace types matching with user provided types
- Mapping: match user-provided paths and separate type expressions as type specs named as instructed by user
- Organize: separate the type definitions as standalone type specs and reuse them when definitions match
- Iterables: implements Range method on those dictionaries that all items are in same type