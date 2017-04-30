package main

import (
  "fmt"
  "github.com/joho/godotenv"
  "log"
  "net/http"
  // "net/url"
)

func routeHandler() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", index)
  http.HandleFunc("/post/", contentPost)
  http.HandleFunc("/submit", receiveContent)
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  routeHandler()

  port := getPort()
  fmt.Printf("\n--- Listening on: %v\n\n", port)

  log.Fatal(http.ListenAndServe(port, nil))
}
