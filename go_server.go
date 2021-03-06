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
	nopref := http.StripPrefix("/public/", fs)
	filesGz := gziphandler.GzipHandler(nopref)
	http.Handle("/public/", filesGz)

	// activatorGz := gziphandler.GzipHandler(http.HandlerFunc(fs2))
	// http.Handle("/432FB6766878ED13CC007C095B54B76A.txt", activatorGz)

	indexGz := gziphandler.GzipHandler(http.HandlerFunc(index))
	postGz := gziphandler.GzipHandler(http.HandlerFunc(contentPost))
	aboutGz := gziphandler.GzipHandler(http.HandlerFunc(about))

	http.Handle("/", indexGz)
	http.Handle("/post/", postGz)
	http.HandleFunc("/submit", receiveContent)
	http.Handle("/site-and-author", aboutGz)
	http.Handle("/contact", aboutGz)
}

// func fs2(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "432FB6766878ED13CC007C095B54B76A.txt")
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routeHandler()

	port := getPort()
	fmt.Printf("\n--- Listening:%v\n\n", port)

	cert, key := getCreds()

	log.Fatal(http.ListenAndServeTLS(port, cert, key, nil))
}
