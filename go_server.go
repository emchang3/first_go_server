package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/joho/godotenv"
)

func routeHandler() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	indexGz := gziphandler.GzipHandler(http.HandlerFunc(index))
	postGz := gziphandler.GzipHandler(http.HandlerFunc(contentPost))

	http.Handle("/", indexGz)
	http.Handle("/post/", postGz)
	http.HandleFunc("/submit", receiveContent)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routeHandler()

	port := getPort()
	fmt.Printf("\n--- Listening on: %v\n\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
