package main

import (
	"io"
	"log"
	"net/http"

	"github.com/tgf9/reactgoboot/dist"
)

func handleMainJS(rw http.ResponseWriter, r *http.Request) {
	log.Println("handleMainJS", r.Method, r.URL.Path)

	fd, err := dist.FS.Open("main.js")
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer fd.Close()

	rw.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	if _, err := io.Copy(rw, fd); err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func handleIndex(rw http.ResponseWriter, r *http.Request) {
	log.Println("handleIndex", r.Method, r.URL.Path)

	fd, err := dist.FS.Open("index.html")
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer fd.Close()

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := io.Copy(rw, fd); err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/main.js", handleMainJS)
	http.HandleFunc("/", handleIndex)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
