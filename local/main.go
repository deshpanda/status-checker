package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Response struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    targetURL := r.URL.Query().Get("url")
    if targetURL == "" {
        json.NewEncoder(w).Encode(Response{
            Status:  "Error",
            Message: "Usage: ?url=URL_TO_CHECK",
        })
        return
    }

    // Add https:// if no protocol specified
    if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
        targetURL = "https://" + targetURL
    }

    _, err := url.ParseRequestURI(targetURL)
    if err != nil {
        json.NewEncoder(w).Encode(Response{
            Status:  "Error",
            Message: "Invalid URL",
        })
        return
    }

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(targetURL)
    if err != nil {
        json.NewEncoder(w).Encode(Response{
            Status:  "Down",
            Message: err.Error(),
        })
        return
    }
    defer resp.Body.Close()

    status := "Up"
    if resp.StatusCode >= 400 {
        status = "Down"
    }

    json.NewEncoder(w).Encode(Response{
        Status:  status,
        Message: resp.Status,
    })
}

func main() {
    http.HandleFunc("/", checkHandler)
    log.Printf("Server starting on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}