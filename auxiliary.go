package main

import (
  "io/ioutil"
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
