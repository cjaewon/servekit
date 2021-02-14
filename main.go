package main

import (
	"log"
	"net/http"
	"servekit/config"
	"servekit/fileserver"
)

func main() {
	config := config.LoadInConfig()
	fs := fileserver.StaticFileSystem{
		Fs:     http.Dir(config.Server.Path),
		Config: config,
	}

	http.Handle("/", http.FileServer(fs))
	log.Printf("Servekit is listening on %s.. \n", config.Server.Port)

	if err := http.ListenAndServe(config.Server.Port, nil); err != nil {
		log.Fatal(err)
	}
}
