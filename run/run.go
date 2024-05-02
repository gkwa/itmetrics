package run

import (
	"fmt"
	"os"
	"path/filepath"
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

	for i, example := range examples {
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

		ordinal := i + 1
		dirName := fmt.Sprintf("example-%03d", ordinal)

		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dirName, err)
			continue
		}

		filePath := filepath.Join(dirName, "manifest.txtar")

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filePath, err)
			continue
		}
		defer file.Close()

		_, err = file.WriteString(strings.TrimSpace(txtar))
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			continue
		}

		fmt.Printf("Processed example %d\n", ordinal)
	}
}
