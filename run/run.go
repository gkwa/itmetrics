package run

import (
	"fmt"
	"strings"

	kcl "kcl-lang.io/kcl-go"
)

func Run(manifestPath string) {
	result, err := kcl.Run(manifestPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := result.First().ToMap()
	if err != nil {
		fmt.Println(err)
		return
	}

	examplesInterface, ok := r["examples"]
	if !ok {
		fmt.Println("Error: 'examples' field not found in KCL result")
		return
	}

	examples, ok := examplesInterface.([]interface{})
	if !ok {
		fmt.Println("Error: 'examples' field is not a slice of interfaces")
		return
	}

	for _, example := range examples {
		exampleMap, ok := example.(map[string]interface{})
		if !ok {
			fmt.Println("Error: example is not a map")
			continue
		}

		txtarInterface, ok := exampleMap["txtar"]
		if !ok {
			fmt.Println("Error: 'txtar' field not found in example")
			continue
		}

		txtar, ok := txtarInterface.(string)
		if !ok {
			fmt.Println("Error: 'txtar' field is not a string")
			continue
		}

		txtarLines := strings.Split(txtar, "\n")

		for _, line := range txtarLines {
			fmt.Println(line)
		}

		fmt.Println("---")
	}
}
