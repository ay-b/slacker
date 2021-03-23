package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func main() {
	getEnv := func(key string) {
		_, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("[ FATAL ] %s not set.\n[ FATAL ] SLACK_API_URL and MESSAGE env vars must be defined!\n[ FATAL ] MESSAGE shall be a json like {\"channel\": \"#devops\", \"username\": \"test\", \"text\": \"This is another test.\"}\n", key)
			os.Exit(1)
		} else {
			fmt.Printf("%s found\n", key)
		}
	}
	getEnv("SLACK_API_URL")
	getEnv("MESSAGE")
	var url = os.Getenv("SLACK_API_URL")
	var jsonStr = []byte(os.Getenv("MESSAGE"))

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Result:", resp.Status)
}