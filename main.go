package main

import (
	"encoding/json"
	"io"
	"os"
)

func Verify(path string) bool {

	// Opening file at given path
	jsonFile, err := os.Open(path)

	// Checking if any errors occured at oppening
	if err != nil {
		return true
	} else {
		// Reading file as binary
		byteValue, _ := io.ReadAll(jsonFile)

		var policy Policy

		// Unmarshaling into Policy object
		json.Unmarshal(byteValue, &policy)

		// Checking for asterisk in all Statements
		for i := 0; i < len(policy.PolicyDocument.Statement); i++ {
			if policy.PolicyDocument.Statement[i].Resource == "*" {
				return false
			}
		}

		// Closing file after everything is finished
		defer jsonFile.Close()
		return true
	}
}

func main() {
	// Puting first argument as the file path
	argsWithoutProg := os.Args[1:]
	println(Verify(argsWithoutProg[0]))
}
