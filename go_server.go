package main

import (
  "fmt"
  "log"
  "net/http"
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
}

func main() {
  routeHandler()

  port := getPort()
  fmt.Printf("\n--- Listening on: %v\n\n", port)

  log.Fatal(http.ListenAndServe(port, nil))
}
