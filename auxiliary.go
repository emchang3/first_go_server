package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

type Page struct {
  Title string
  Body []string
}

func loadPage(title string) (*Page, error) {
  filename := "content/" + title + ".txt"

  raw, err := ioutil.ReadFile(filename)
  if err != nil {
      return nil, err
  }

  myBody := fmt.Sprintf("%s", raw)
  bodySplit := strings.Split(myBody, "\n")

  body := make([]string, 0)
  for _, v := range bodySplit {
    if v != "" {
      body = append(body, v)
    }
  }

  return &Page{Title: title, Body: body}, nil
}
