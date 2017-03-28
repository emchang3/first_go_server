package server_test

import (
  "html/template"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  p, err := loadPage("test")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  t, err := template.ParseFiles("views/index.ego", "views/partials/something.ego")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  t.Execute(w, p)
}
