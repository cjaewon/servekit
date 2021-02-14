package main

import (
	"net/http"
	"servekit/config"
	"servekit/fileserver"
	"servekit/logger"

	"go.uber.org/zap"
)

func main() {
	config := config.LoadInConfig()
	fs := fileserver.StaticFileSystem{
		Fs:     http.Dir(config.Server.Path),
		Config: config,
	}

	http.Handle("/", http.FileServer(fs))
	logger.Log().Info("Servekit is listening", zap.String("port", config.Server.Port))

	if err := http.ListenAndServe(config.Server.Port, nil); err != nil {
		logger.Log().Fatal("Failed to start http server", zap.Error(err))
	}
}
