package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Message struct {
    Text string `json:"message"`
}

func main() {
    http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
        message := Message{Text: "Hello from Go backend!"}
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        json.NewEncoder(w).Encode(message)
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}