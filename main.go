package main

import (
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("asterisk.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("import sucessful")
	}

	defer jsonFile.Close()
}
