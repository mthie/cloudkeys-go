package main

import (
	"mime"
	"net/http"
	"path/filepath"
)

func serveAssets(res http.ResponseWriter, r *http.Request) {
	data, err := Asset(r.RequestURI[1:])
	if err != nil {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(r.RequestURI)))
	res.Write(data)
}
