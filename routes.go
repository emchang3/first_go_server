package main

import (
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  err := loadTextPost("1", w, r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
