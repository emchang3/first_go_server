package main

import (
	"fmt"
	"io/ioutil"
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

	err = loadContentPost(file, file, w, r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}

	file, err := getLatestFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = loadAbout("about", file, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contentPost(w http.ResponseWriter, r *http.Request) {
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

	err = loadContentPost(this, file, w, r, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func receiveContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	appToken := getToken()
	if appToken == "nah, bruh." {
		http.Error(w, "Invalid configuration.", http.StatusInternalServerError)
		return
	}

	submittedToken := r.Header["Application-Token"]
	if submittedToken == nil || submittedToken[0] != appToken {
		http.Error(w, "Invalid credentials.", http.StatusForbidden)
		return
	}

	filename := r.Header["Filename"]
	if filename == nil {
		http.Error(w, "Invalid filename.", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusInternalServerError)
		return
	}

	mb := fmt.Sprintf("%s", body)
	bytes := []byte(mb)

	file := fmt.Sprintf("content/%v.emc", filename[0])

	err = ioutil.WriteFile(file, bytes, 0644)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error writing request body.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
