package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Content struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func main() {
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []Content{
				{ID: 1, Title: "Облачный Go API", Author: "Render", Description: "Работает через HTTPS!", Image: "https://picsum.photos/200"},
			},
		})
	})

	// Сервис сам передаст нужный порт
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Сервер запущен на порту " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}