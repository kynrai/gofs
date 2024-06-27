package handlers

import (
	"embed"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
)

func HandleAssets(assets embed.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hasher := crc32.NewIEEE()
		file, err := assets.Open(r.URL.Path)
		if err != nil {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		if _, err := io.Copy(hasher, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		etag := fmt.Sprintf("%d-%v", len(r.URL.Path), hasher.Sum32())
		ifNoneMatch := r.Header.Get("If-None-Match")

		if ifNoneMatch == etag {
			w.WriteHeader(http.StatusNotModified)
		} else {
			w.Header().Add("ETag", etag)
			http.FileServerFS(assets).ServeHTTP(w, r)
		}
	}
}
