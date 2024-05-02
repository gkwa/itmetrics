package run

import (
	"fmt"
	"sort"

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

	packagesInterface, ok := r["packages"]
	if !ok {
		fmt.Println("Error: 'packages' field not found in KCL result")
		return
	}

	packages, ok := packagesInterface.([]interface{})
	if !ok {
		fmt.Println("Error: 'packages' field is not a slice of interfaces")
		return
	}

	sort.Slice(packages, func(i, j int) bool {
		pi := packages[i].(map[string]interface{})["check_installed"].(string)
		pj := packages[j].(map[string]interface{})["check_installed"].(string)
		return pi < pj
	})

	for _, pkg := range packages {
		name := pkg.(map[string]interface{})["name"].(string)
		checkInstalled := pkg.(map[string]interface{})["check_installed"].(string)
		fmt.Printf("Processing package: %s, check installed: %s\n", name, checkInstalled)
	}
}
