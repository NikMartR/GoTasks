package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const version = "v1.0.0"

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/version", handleVersion)
	http.HandleFunc("/decode", handleDecode)
	http.HandleFunc("/hard-op", handleHardOp)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Println("Сервер запущен на порту : 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Ошибка сервера:", err)
		}
	}()

	<-stop
	fmt.Println("Сервер выключается...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Ошибка выключения сервера:", err)
	}
	fmt.Println("Сервер выключен.")
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
		http.Error(w, "Некоректный запрос", http.StatusBadRequest)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(input.InputString)
	if err != nil {
		http.Error(w, "Ошибка декодирования base64", http.StatusBadRequest)
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
	sleepTime := rand.Intn(10) + 10
	time.Sleep(time.Duration(sleepTime) * time.Second)

	if rand.Intn(2) == 0 {
		http.Error(w, "Внутреняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Операция выполнена успешна")
}
