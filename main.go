package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
)

func fatalExit (key string) {
	fmt.Printf("[ FATAL ] %s not set.\n" +
		"          Both API_URL and MESSAGE env vars must be defined!\n" +
		"          MESSAGE shall be a json like {\"channel\": \"devops\", \"username\": \"test\", \"text\": \"This is another test.\"}\n", key)
	os.Exit(1)
}

func logger (key string, val string) {
	if len(os.Getenv("DEBUG")) != 0 {
		fmt.Printf("[ DEBUG ] %s found\n          Value: %s\n", key, val)
	}
}

func getEnv (key string) {
	val, ok := os.LookupEnv(key)
	if !ok {
		fatalExit(key)
	} else {
		if len(val) == 0 {
			fatalExit(key)
		}
		logger(key, val)
	}
}

func main() {
	getEnv("API_URL")
	getEnv("MESSAGE")
	var url = os.Getenv("API_URL")
	var jsonStr = []byte(os.Getenv("MESSAGE"))

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ FATAL ERROR ] %s\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println("Result:", resp.Status)
}
