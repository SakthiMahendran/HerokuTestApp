package main

import (
	"net/http"
	"os"
)

type TestServer struct {
}

func (ts *TestServer) StartServing() {
	http.HandleFunc("/", ts.requestHandler)

	http.ListenAndServe(":80", nil)
}

func (ts *TestServer) requestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if name := r.Form.Get("Name"); name != "" {
		ts.respondHtml(w, r, "webpage/thankyou.html")
		return
	}

	if r.URL.Path == "/" {
		ts.respondHtml(w, r, "webpage/index.html")
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (ts *TestServer) respondHtml(w http.ResponseWriter, r *http.Request, htmlPath string) {
	w.Write(ts.readIndexHtml(htmlPath))
}

func (ts *TestServer) readIndexHtml(path string) []byte {
	fileContent, _ := os.ReadFile(path)
	return fileContent
}
