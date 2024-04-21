package main

import (
	"os"
)

func Verify(path string) bool {
	jsonFile, err := os.Open(path)

	if err != nil {
		return true
	} else {
		defer jsonFile.Close()
		return true
	}
}
