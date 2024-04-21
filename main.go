package main

import (
	"encoding/json"
	"io"
	"os"
)

type Policy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Sid       string   `json:"Sid"`
	Effect    string   `json:"Allow"`
	Principal string   `json:"Principal"`
	Action    []string `json:"Action"`
	Resource  string   `json:"Resource"`
	Condition string   `json:"Condition"`
}

func Verify(path string) bool {
	jsonFile, err := os.Open(path)

	if err != nil {
		return true
	} else {
		byteValue, _ := io.ReadAll(jsonFile)

		var policy Policy

		json.Unmarshal(byteValue, &policy)

		for i := 0; i < len(policy.PolicyDocument.Statement); i++ {
			if policy.PolicyDocument.Statement[i].Resource == "*" {
				return false
			}
		}

		defer jsonFile.Close()
		return true
	}
}
