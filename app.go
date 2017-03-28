package server_test

import (
  // "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

type Page struct {
  Title string
  Body []byte
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"

  body, err := ioutil.ReadFile(filename)
  if err != nil {
      return nil, err
  }

  return &Page{Title: title, Body: body}, nil
}

func routes()  {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/", index)
}

func main() {
  routes()

  log.Fatal(http.ListenAndServe(":8080", nil))
}
