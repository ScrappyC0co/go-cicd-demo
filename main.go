package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {                                        fmt.Fprintf(w, "Hello, CI/CD from Go!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server staring at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}


