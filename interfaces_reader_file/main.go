package main

import (
	"fmt"
	"os"
)

func main() {

	var fileName string

	if existsArgument() {
		fileName = os.Args[1]
	}

	fmt.Println(fileName)
}

func existsArgument() bool {
	return len(os.Args) > 1
}
