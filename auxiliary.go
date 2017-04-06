package main

import (
  "fmt"
  "html/template"
  "io/ioutil"
  "net/http"
  "strconv"
  "strings"
)

type Page struct {
  Title string
  Body []string
  Split func(string, string) []string
}

func getLatestFile() (int, []int, error) {
  var files []int
  dirname := "content"
  contents, err := ioutil.ReadDir(dirname)
  if err != nil {
    return -1, files, err
  }

  latest := 0
  for _, file := range contents {
    prefixString := strings.Split(file.Name(), ".")[0]
    prefixInt, err := strconv.Atoi(prefixString)
    if err != nil {
      return -1, files, err
    }
    files = append(files, prefixInt)

		if prefixInt > latest {
		  latest = prefixInt
		}
	}

  // fmt.Println(files)

  return latest, files, nil
}

func loadTextPost(file string, w http.ResponseWriter, r *http.Request) error {
  filename := "content/" + file + ".txt"

  raw, err := ioutil.ReadFile(filename)
  if err != nil {
    return err
  }

  getLatestFile()

  myBody := fmt.Sprintf("%s", raw)
  bodySplit := strings.Split(myBody, "\n")

  body := make([]string, 0)
  title := ""
  for _, v := range bodySplit {
    if strings.Split(v, " ")[0] == "//" {
      title = strings.Join(strings.Split(v, " ")[1:], " ")
      continue
    }
    if v != "" {
      body = append(body, v)
    }
  }

  p := &Page{Title: title, Body: body, Split: strings.Split}

  t, err := template.ParseFiles("views/index.gohtml", "views/partials/content.gohtml", "views/partials/menuButton.gohtml")
  if err != nil {
    return err
  }

  t.Execute(w, p)
  return nil
}
