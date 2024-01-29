# `gonfic` Type checked, dynamic keyed configs for Go programs

`gonfic` is tool that creates a Go struct type definitions for given YAML files which has custom schemas. Since `gonfic` makes it easier to use custom schemas, developers can levarage dynamic-keyed configs to access information with autocomplete and without searching through items' values. Also existing Makefile users can get instantly notified when and where a config change breaks Go code (See [Serving suggestion](#serving-suggestions)).

Before `gonfic`

```yaml
Services:
    - Name: A
      Details:
           Port: -1
           Path: ""
    - Name: B
           Port: -1
           Path: ""
```

```go
func main() {
    cfg := ReadConfig()
    fmt.Println(
        cfg.Get("a").Get("port"),
        cfg.Get("a").Get("path"),
    )
}
```

After `gonfic`

```yaml
Services:
    A:
        Port: -1
        Path: ""
    B:
        Port: -1
        Path: ""
```

```go
func main() {
    cfg := ReadConfig()
    fmt.Println(cfg.Services.A.Port, cfg.Services.A.Path)
}
```

## Install

```sh
go get -u github.com/ufukty/gonfic
```

## Usage

```sh
gonfic -in config.yml -out gonfic.go -pkg main
```

## Serving suggestions

For existing Makefile users:

```Makefile
gonfic.go: config.yml
    gonfic -in config.yml -out gonfic.go -pkg main

all: gonfic.go
    ...
```

For existing Visual Studio Code users:

```json
{
    "runOnSave.commands": [
        {
            "match": "^config.yml$",
            "command": "cd '${fileDirname}' && gonfic"
        }
    ]
}
```

## Motivation

-   Using dynamic keys in config files is not problematic as using them in server-client communication. Because all keys are final before compilation starts.
-   Having to search to get a piece of config is error-prone, makes harder to notice breaking changes. Also creates ugly code.

## Known issues

-   Doesn't work for slices. (Fallbacks to `any` type) See `pkg/testdata/tc2/config.yml`
-   Doesn't escape keys of YAML file that collide with Go keywords (like `type`, start with capitalized letters)

## License

Apache2. See LICENSE file.
