package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	response, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	logWriter := logWriter{}

	io.Copy(logWriter, response.Body)
}

func (logWriter) Write(bytes []byte) (int, error) {
	fmt.Println(string(bytes))
	fmt.Println("Just wrote this many bytes:", len(bytes))
	return len(bytes), nil
}
