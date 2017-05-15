package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type link struct {
	Title string
	File  string
}

type page struct {
	Title    string
	Body     []string
	Current  int
	Latest   int
	Pages    []link
	Split    func(string, string) []string
	Add      func(int, int) int
	Subtract func(int, int) int
}

type special struct {
	Title string
	Body  []string
	Pages []link
	Split func(string, string) []string
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func getLatestFile() (int, error) {
	dirname := "content"
	contents, err := ioutil.ReadDir(dirname)
	if err != nil {
		return -1, err
	}

	latest := 0
	for _, file := range contents {
		prefixString := strings.Split(file.Name(), ".")[0]
		prefixInt, err := strconv.Atoi(prefixString)
		if err != nil {
			return -1, err
		}

		if prefixInt > latest {
			latest = prefixInt
		}
	}

	return latest, nil
}

func getPost(file int) (string, []string, error) {
	myFile := strconv.Itoa(file)
	filename := "content/" + myFile + ".emc"

	body := make([]string, 0)

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", body, err
	}

	myBody := fmt.Sprintf("%s", raw)
	bodySplit := strings.Split(myBody, "\n")

	title := ""
	for _, v := range bodySplit {
		if strings.Split(v, " ")[0] == "^t" {
			title = strings.Join(strings.Split(v, " ")[1:], " ")
			continue
		}
		if v != "" {
			body = append(body, v)
		}
	}

	return title, body, err
}

func getSpecial(file string) (string, []string, error) {
	filename := "special/" + file + ".emc"

	body := make([]string, 0)

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", body, err
	}

	myBody := fmt.Sprintf("%s", raw)
	bodySplit := strings.Split(myBody, "\n")

	title := ""
	for _, v := range bodySplit {
		if strings.Split(v, " ")[0] == "^t" {
			title = strings.Join(strings.Split(v, " ")[1:], " ")
			continue
		}
		if v != "" {
			body = append(body, v)
		}
	}

	return title, body, err
}

func loadContentPost(file int, latest int, w http.ResponseWriter, r *http.Request, isIndex bool) error {
	// fmt.Println(file, latest)

	title, body, err := getPost(file)
	if err != nil {
		return err
	}

	pages := make([]link, 0)
	current := latest
	for current > 0 {
		currentTitle, _, err := getPost(current)
		if err != nil {
			return err
		}
		// fmt.Println(currentTitle)

		currentLink := link{Title: currentTitle, File: strconv.Itoa(current)}
		pages = append(pages, currentLink)
		current--
	}
	// fmt.Println(pages)

	p := &page{Title: title, Body: body, Current: file, Latest: latest, Pages: pages, Split: strings.Split, Add: add, Subtract: subtract}

	tmplt := ""
	if isIndex {
		tmplt = "views/index.gohtml"
	} else {
		tmplt = "views/post.gohtml"
	}

	t, err := template.ParseFiles(tmplt, "views/partials/content.gohtml", "views/partials/menuButton.gohtml", "views/partials/navMenu.gohtml", "views/partials/navArrows.gohtml")
	if err != nil {
		return err
	}

	t.Execute(w, p)
	return nil
}

func loadAbout(file string, latest int, w http.ResponseWriter, r *http.Request) error {
	// fmt.Println(file, latest)

	title, body, err := getSpecial(file)
	if err != nil {
		return err
	}

	pages := make([]link, 0)
	current := latest
	for current > 0 {
		currentTitle, _, err := getPost(current)
		if err != nil {
			return err
		}
		// fmt.Println(currentTitle)

		currentLink := link{Title: currentTitle, File: strconv.Itoa(current)}
		pages = append(pages, currentLink)
		current--
	}
	// fmt.Println(pages)

	p := &special{Title: title, Body: body, Pages: pages, Split: strings.Split}

	tmplt := "views/about.gohtml"

	t, err := template.ParseFiles(tmplt, "views/partials/content.gohtml", "views/partials/menuButton.gohtml", "views/partials/navMenu.gohtml")
	if err != nil {
		return err
	}

	t.Execute(w, p)
	return nil
}
