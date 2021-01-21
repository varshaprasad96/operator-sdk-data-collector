package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/varshaprasad96/operator-sdk-data-collector/pkg/collector"
	"github.com/varshaprasad96/operator-sdk-data-collector/pkg/fields"
	output "github.com/varshaprasad96/operator-sdk-data-collector/pkg/output/xlsx"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("path is a required argument\n")
		os.Exit(1)
	}

	inputList := make([]fields.Inputs, 0)

	for i := 0; i < len(args); i++ {
		v := strings.Split(args[i], ":")
		inputs := fields.Inputs{
			Path:    v[0],
			Source:  v[1],
			Version: v[2],
		}
		inputList = append(inputList, inputs)
	}

	// get operator data on all the 3 indexes
	operatorData := collector.CollectDump(inputList)

	filepath, err := os.Getwd()
	if err != nil {
		fmt.Printf("error  finding path to write report")
		os.Exit(1)
	}

	// TODO: Add csv format and make it more reusable
	err = output.GetOutput(operatorData, filepath+"/report/")
	if err != nil {
		fmt.Printf("something wrong while writing the output")
		os.Exit(1)
	}
}
