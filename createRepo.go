package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type createRepo struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: createRepo <repoName>")
		os.Exit(2)
	}
	repoName := os.Args[1]
	token, found := os.LookupEnv("API_TOKEN")
	if !found {
		fmt.Println("Must set API_TOKEN environment variable")
		os.Exit(2)
	}
	url := "https://gitea.pfdev.de/api/v1/user/repos"
	method := "POST"

	payload := &createRepo{
		Name:    repoName,
		Private: true,
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 201 {
		fmt.Fprintf(os.Stderr, "Code: %d Error: %v", res.StatusCode, err)
		os.Exit(1)
	}
}
