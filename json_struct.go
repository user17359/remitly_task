package main

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
