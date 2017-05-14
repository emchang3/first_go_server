package main

import (
	"crypto/tls"
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

	http.Handle("/", indexGz)
	http.Handle("/post/", postGz)
	http.HandleFunc("/submit", receiveContent)
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

	if port != ":80" {
		// log.Fatal(http.ListenAndServe(port, nil))
		log.Fatal(http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil))
	} else {
		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
		srv := &http.Server{
			Addr:      port,
			Handler:   nil,
			TLSConfig: cfg}
		log.Fatal(srv.ListenAndServeTLS("jnsq-bundle.pem", "jnsq.ninja.pem"))
	}
}
