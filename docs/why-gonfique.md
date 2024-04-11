# Why gonfique?

Since `gonfique` makes **keeping the type definitions up-to-date** easier, using more dynamic keys instead arrays is more practical. So, developers can access config data through field names instead error-prone array lookups.

`gonfique` is necessary because of there are not many reliable and sustainable alternatives.

- Accessing config data through **hardcoded strings are risky**. So, defining types to marshall into is necessary.
- Manually defining types is also risky because they will **get outdated** eventually.
- Config **complexity is inevitable** when there are multiple services/binaries that needs their config to stay in sync, eg. kubernetes config.

## Your config file

### Before `gonfique`

Currently, you are storing multiple items **in arrays** like this

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

Then, you define `Config` type:

```go
type Config struct {
  Github struct {
    Domain   int
    Path     int
    Services []Service
  } `yaml:"github.com"`
}
```

And check again & again, if you made a mistake...

Lastly, access items **by lookups** with **hardcoded strings**. Compiler won't catch if the information gets outdated. Which leads you notice problems in **runtime**.

```go
func main() {
  // ...
  list, found := cfg.Github.Services.Tags.Endpoints.Lookup("list")
  if found {
    fmt.Println(list.Path)
  }
}
```

### After `gonfique`

You don't have to worry anymore about storing part of the config information in schema, since type definitions are **automatically generated**, **instantly updated** and accesses are **re-verified by compiler** as config changed.

```yaml
github:
  domain: github.com
  path: /api/v1.0.0
  services:
    tags:
      path: tags
      endpoints:
        list: { method: "GET", path: "tags" }
        create: { method: "POST", path: "tag" }
        assign: { method: "POST", path: "assign" }
        delete: { method: "DELETE", path: "tag" }
gitlab:
  domain: gitlab.com
```

Access items through fields instead hardcoded strings. Always under type-check, and easily updated as config changes.

```go
func main() {
  // ...
  fmt.Println(cfg.Github.Services.Tags.Endpoints.List.Path)
}
```

Arrays and dictionaries are still **iteratable** via `.Range` method (if the `-organize` flag provided):

```go
func main() {
  // ...
  for key, ep := range cfg.Github.Services.Tags.Endpoints.Range() {
    fmt.Println(key, ep.Path)
  }
}
```
