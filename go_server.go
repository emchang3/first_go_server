package main

import (
  // "fmt"
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

  log.Fatal(http.ListenAndServe(":8080", nil))
}
