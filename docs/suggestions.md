# Serving suggestions

For existing Makefile users:

```Makefile
config.go: config.yml
    gonfique -in config.yml -out config.go -pkg main

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
