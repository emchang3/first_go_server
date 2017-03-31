package main

import (
  "fmt"
  "html/template"
  "io/ioutil"
  "net/http"
  "strings"
)

type Page struct {
  Title string
  Body []string
}

func loadTextPost(file string, title string, w http.ResponseWriter, r *http.Request) error {
  filename := "content/" + file + ".txt"

  raw, err := ioutil.ReadFile(filename)
  if err != nil {
    return err
  }

  myBody := fmt.Sprintf("%s", raw)
  bodySplit := strings.Split(myBody, "\n")

  body := make([]string, 0)
  for _, v := range bodySplit {
    if v != "" {
      body = append(body, v)
    }
  }

  p := &Page{Title: title, Body: body}

  t, err := template.ParseFiles("views/index.gohtml", "views/partials/content.gohtml")
  if err != nil {
    return err
  }

  t.Execute(w, p)
  return nil
}
