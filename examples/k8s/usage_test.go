package main

import (
	"fmt"
	"os"

	example1 "github.com/ufukty/gonfique/examples/k8s/basic"
	example2 "github.com/ufukty/gonfique/examples/k8s/organized"
	example3 "github.com/ufukty/gonfique/examples/k8s/organized-used"
)

func acceptPort(port example3.Port) {
	fmt.Println(port.TargetPort, port.Protocol, port.Port)
}

// With basic usage, you can get all config accesses under type-check.
// So, when your config changed, you only need to rerun gonfique to see
// where the broken accesses are.
func ExampleBasicUsage() {
	cfg, err := example1.ReadConfig("input.yml")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(cfg.Metadata.Name) // Output: my-deployment
}

// When -organize flag used, gonfique also define ".Range" methods on the
// dictionaries which their all items share same schema.
// So, you can iterate over them just like;
func ExampleIteratingOverDictionaries() {
	cfg2, err := example2.ReadConfig("input.yml")
	if err != nil {
		os.Exit(1)
	}
	for yamlkey, metadata := range cfg2.Spec.Template.Metadata.Range() {
		fmt.Println(yamlkey, "=>", metadata.App)
	}
	// Output:
	// labels => my-app
}

func ExampleUsingExistingTypeDefinitionsToReferComponents() {
	cfg3, err := example3.ReadConfig("input.yml")
	if err != nil {
		os.Exit(1)
	}
	for _, port := range cfg3.Spec.Ports {
		acceptPort(port)
	}
	// Output: 80 TCP 80
}
