package main

import (
  "fmt"
  "bytes"
  "context"
  "encoding/json"
  "io"
  "net/http"
  "time"
)

func main() {
  client := &http.Client{}
  ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
  defer cancel()

  baseURL := "http://localhost:8080"

  versionResp, err := doGet(ctx, client, baseURL+"/version")
  if err != nil {
  fmt.Println("Ошибка запроса /version:", err)
 } else {
  fmt.Println("Version:", versionResp)
 }

  encoded := base64Encoder("Hello, World!")
  body := map[string]string{"inputString": encoded}
  decodedResp, err := doPost(ctx, client, baseURL+"/decode", body)
  if err != nil {
    fmt.Println("Ошибка запроса /decode:", err)
 } else {
    fmt.Println("Decoded:", decodedResp)
 }

 hardOpResp, err := doGet(ctx, client, baseURL+"/hard-op")
 if err != nil {
   fmt.Println("Ошибка запроса /hard-op:", err)
 } else {
   fmt.Println("Hard-op response:", hardOpResp)
 }
}

func doGet(ctx context.Context, client *http.Client, url string) (string, error) {
  req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
  if err != nil {
    return "", err
 }

 resp, err := client.Do(req)
 if err != nil {
   return "", err
 }
 defer resp.Body.Close()

 body, _ := io.ReadAll(resp.Body)
 return string(body), nil
}

func doPost(ctx context.Context, client *http.Client, url string, payload interface{}) (string, error) {
 data, err := json.Marshal(payload)
 if err != nil {
  return "", err
 }

 req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
 if err != nil {
  return "", err
 }
 req.Header.Set("Content-Type", "application/json")

 resp, err := client.Do(req)
 if err != nil {
  return "", err
 }
 defer resp.Body.Close()

 body, _ := io.ReadAll(resp.Body)
 return string(body), nil
}

func base64Encoder(input string) string {
 return base64.StdEncoding.EncodeToString([]byte(input))
}
