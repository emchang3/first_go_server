package main

import (
  "fmt"
  "log"
  "net/http"
)

func routeHandler()  {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", index)
}

func main() {
  routeHandler()

  port := ":8080"

  log.Fatal(http.ListenAndServe(port, nil))
  fmt.Printf("--- Listening on %v.\n", port)
}
