basic/output.go: input.yml
	gonfique -in input.yml -out basic/output.go -pkg config

organized/output.go: input.yml
	gonfique -in input.yml -out organized/output.go -pkg config -organize

organized-used/output.go: input.yml
	gonfique -in input.yml -out organized-used/output.go -pkg config -organize -use organized-used/use.go

all: basic/output.go organized/output.go organized-used/output.go