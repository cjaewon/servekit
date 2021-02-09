package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// SERVEKIT_PATH
	// SERVEKIT_PORT
	var (
		path = "./static"
		port = ":3000"
	)

	if os.Getenv("SERVEKIT_PATH") != "" {
		path = os.Getenv("SERVEKIT_PATH")
	}
	if os.Getenv("SERVEKIT_PORT") != "" {
		port = os.Getenv("SERVEKIT_PORT")
	}

	http.Handle("/", http.FileServer(http.Dir(path)))
	log.Printf("Servekit is listening on %s.. \n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
