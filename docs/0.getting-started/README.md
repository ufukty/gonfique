# Getting started

## Install

Download, compile and install the latest release of Gonfique:

```sh
go install go.ufukty.com/gonfique@v2.0.0-alpha.2
```

Validate installation:

```sh
gonfique version
```

If shell reports it can't find the command then check your `PATH` contain `GOBIN` and `GOPATH` or use a familiar `GOBIN` with the install command.

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

## Serving suggestions

### With Makefile

For existing Makefile users:

```Makefile
config.go: config.yml gonfique.yml
    gonfique -in config.yml -out config.go -config gonfique.yml

all: config.go
    ...
```

### With Visual Studio Code

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
