package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type HTMLParser struct {
	xml.Name
	BodyText string `xml:"body"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	response, err := client.Get("http://172.31.28.67")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	extractedText := &HTMLParser{}
	xml.Unmarshal([]byte(body), extractedText)
	if response.StatusCode == 200 && extractedText.BodyText == "Hello World!" {
		fmt.Println("Ok")
		os.Exit(0)
	} else {
		fmt.Println("Unexpected response")
		os.Exit(1)
	}
}
