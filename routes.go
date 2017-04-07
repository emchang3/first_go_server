package main

import (
  "net/http"
  "strconv"
  "strings"
)

func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  file, err := getLatestFile()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = loadTextPost(file, file, w, r, true)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func textPost(w http.ResponseWriter, r *http.Request) {
  secondary := strings.Split(r.URL.Path, "/")[2]
  this, err := strconv.Atoi(secondary)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  file, err := getLatestFile()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = loadTextPost(this, file, w, r, false)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
