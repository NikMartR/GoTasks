package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const version = "v1.0.0"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/version", handleVersion)
	mux.HandleFunc("/decode", handleDecode)
	mux.HandleFunc("/hard-op", handleHardOp)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Сервер запущен на порту 8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, version)
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InputString string `json:"inputString"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(input.InputString)
	if err != nil {
		http.Error(w, "Invalid base64 string", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"outputString": string(decoded),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleHardOp(w http.ResponseWriter, r *http.Request) {
	// Имитация долгой операции
	sleepTime := rand.Intn(10) + 10 // 10–20 секунд
	time.Sleep(time.Duration(sleepTime) * time.Second)
	if rand.Intn(2) == 0 {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Operation completed successfully")
}
