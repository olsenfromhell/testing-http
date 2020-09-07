package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"minitank/src"
)

func main() {
	sendRequest()
}

func sendRequest() {

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(src.Body)
	req, _ := http.NewRequest("POST", "https://example.com/master/auth/login", buf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)

	fmt.Println("Response Status:", res.Status)

}

