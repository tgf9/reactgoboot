package main

import (
	"io"
	"log"
	"net/http"

	"github.com/tgf9/reactgoboot/public"
)

func serveFile(name, contentType string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fd, err := public.FS.Open(name)
		if err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("serveFile: failed to open %s: %s", name, err)
			return
		}
		defer fd.Close()

		rw.Header().Set("Content-Type", contentType)

		if _, err := io.Copy(rw, fd); err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("serveFile: failed to copy %s: %s", name, err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/index.js", serveFile("index.js", "text/javascript; charset=utf-8"))
	http.HandleFunc("/index.css", serveFile("index.css", "text/css; charset=utf-8"))
	http.HandleFunc("/", serveFile("index.html", "text/html; charset=utf-8"))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
