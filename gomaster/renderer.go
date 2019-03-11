package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/thedevsaddam/renderer"
)

func main() {
	rnd := renderer.New()

	mux := http.NewServeMux()

	usr := struct {
		Name string
		Age  int
	}{"John Doe", 30}

	// serving String as text/plain
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rnd.String(w, http.StatusOK, "Welcome to renderer")
	})

	// serving success but no content
	mux.HandleFunc("/no-content", func(w http.ResponseWriter, r *http.Request) {
		rnd.NoContent(w)
	})

	// serving string as html
	mux.HandleFunc("/html-string", func(w http.ResponseWriter, r *http.Request) {
		rnd.HTMLString(w, http.StatusOK, "<h1>Hello Renderer!</h1>")
	})

	// serving JSON
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		rnd.JSON(w, http.StatusOK, usr)
	})

	// serving JSONP
	mux.HandleFunc("/jsonp", func(w http.ResponseWriter, r *http.Request) {
		rnd.JSONP(w, http.StatusOK, "callback", usr)
	})

	// serving XML
	mux.HandleFunc("/xml", func(w http.ResponseWriter, r *http.Request) {
		rnd.XML(w, http.StatusOK, usr)
	})

	// serving YAML
	mux.HandleFunc("/yaml", func(w http.ResponseWriter, r *http.Request) {
		rnd.YAML(w, http.StatusOK, usr)
	})

	// serving File as arbitary binary data
	mux.HandleFunc("/binary", func(w http.ResponseWriter, r *http.Request) {
		var reader io.Reader
		reader, _ = os.Open("../README.md")
		rnd.Binary(w, http.StatusOK, reader, "readme.md", true)
	})

	// serving File as inline
	mux.HandleFunc("/file-inline", func(w http.ResponseWriter, r *http.Request) {
		rnd.FileView(w, http.StatusOK, "../README.md", "readme.md")
	})

	// serving File as attachment
	mux.HandleFunc("/file-download", func(w http.ResponseWriter, r *http.Request) {
		rnd.FileDownload(w, http.StatusOK, "../README.md", "readme.md")
	})

	// serving File from reader as inline
	mux.HandleFunc("/file-reader", func(w http.ResponseWriter, r *http.Request) {
		var reader io.Reader
		reader, _ = os.Open("../README.md")
		rnd.File(w, http.StatusOK, reader, "readme.md", true)
	})

	// serving custom response using render and chaining methods
	mux.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(renderer.ContentType, renderer.ContentText)
		rnd.Render(w, http.StatusOK, []byte("Send the message as text response"))
	})

	port := ":9000"
	log.Println("Listening on port", port)
	http.ListenAndServe(port, mux)
}
