package main

import (
  "fmt"
  "github.com/joho/godotenv"
  "log"
  "net/http"
  // "net/url"
  "os"
)

func getPort() string {
  port := os.Getenv("PORT")

  if port == "" {
    return ":8080"
  }

  return fmt.Sprintf(":%v", port)
}

func routeHandler() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", index)
  http.HandleFunc("/post/", textPost)
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
