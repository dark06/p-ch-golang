package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	respone, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := io.ReadAll(respone.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if respone.StatusCode == 200 && string(body) == "Hello World" {
		fmt.Println("Ok")
		os.Exit(0)
	} else {
		fmt.Println("Unexpected response")
		os.Exit(1)
	}
}
