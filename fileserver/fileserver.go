package fileserver

import (
	"net/http"
	"os"
	"path/filepath"
	"servekit/config"
)

// StaticFileSystem is wrapper of http FileSystem
type StaticFileSystem struct {
	Fs     http.FileSystem
	Config *config.Config
}

// Open iswrapper of http FileSystem Open
func (fs StaticFileSystem) Open(path string) (http.File, error) {
	f, err := fs.Fs.Open(path)

	if os.IsNotExist(err) && fs.Config.Server.Mode == "client-side-rendering" {
		return fs.Fs.Open("index.html")
	} else if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() && !fs.Config.Server.Overview {
		index := filepath.Join(path, "index.html")
		indexFile, err := fs.Fs.Open(index)
		if err != nil {
			f.Close()
			return nil, os.ErrNotExist
		}

		defer indexFile.Close()
	}

	return f, nil
}
