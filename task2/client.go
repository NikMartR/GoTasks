package main

import (
  "fmt"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	fmt.Println("Requesting /version")
	versionResp, err := client.Get("http://localhost:8080/version")
	handleResponse("GET /version", versionResp, err)

	fmt.Println("Запрос к /decode")
	reqBody, _ := json.Marshal(map[string]string{"inputString": "SGVsbG8gd29ybGQ="}) // "Hello world" in base64
	decodeReq, _ := http.NewRequestWithContext(ctx, "POST", "http://localhost:8080/decode", bytes.NewBuffer(reqBody))
	decodeReq.Header.Set("Content-Type", "application/json")
	decodeResp, err := client.Do(decodeReq)
	handleResponse("POST /decode", decodeResp, err)

	fmt.Println("Запрос к /hard-op")
	hardOpReq, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/hard-op", nil)
	hardOpResp, err := client.Do(hardOpReq)
	handleResponse("GET /hard-op", hardOpResp, err)
}

func handleResponse(name string, resp *http.Response, err error) {
	if err != nil {
		fmt.Printf("%s Ошибка: %v\n", name, err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s: Статус %d, Ответ: %s\n", name, resp.StatusCode, string(body))
}
