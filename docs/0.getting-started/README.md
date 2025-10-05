# Getting started

## Install

Since Gonfique needs the version information to be embedded into the binary, you need to compile it with necessary flags. The Makefile in the root of project contains a recipe that conforms the criteria.

Version information is important because Gonfique stamps the generated file with it. Which is to allow your colleagues to reproduce the same results by getting the correct version of Gonfique in future.

### Download

Simply download the Gonfique binary compiled for your operating system and architecture. Then put it into some directory listed in the `$PATH` variable.

Make sure shell can find Gonfique:

```sh
which gonfique
```

If there is no output, then check if the directory is in `$PATH`. If the path is printed, you are ready to roll.

### Install from source

You need to clone repository and to run the Make recipe installs the binary after compiling with correct version information:

```sh
git clone go.ufukty.com/gonfique
cd gonfique
make install
```

## Build

You can build the Gonfique binaries for all architectures and operating systems at once. Clone the repository and run the Makefile recipe installs the binary after compiling with correct version information:

```sh
git clone go.ufukty.com/gonfique
cd gonfique
make build
```

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
