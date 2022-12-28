package testserver

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type TestServer struct {
	LogFile *os.File
}

func (ts *TestServer) StartServing() {
	http.HandleFunc("/", ts.requestHandler)

	http.ListenAndServe(":80", nil)
}

func (ts *TestServer) requestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if name := r.Form.Get("Name"); name != "" {
		templ := template.Must(template.ParseFiles("webpage/thankyou.html"))

		userinfo := userInfo{Ip: r.Header.Get("X-FORWARDED-FOR"), Device: r.Header.Get("User-Agent"), Name: name}

		err := templ.Execute(w, userinfo)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		ts.logAndSave(userinfo)
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

func (ts *TestServer) logAndSave(u userInfo) {
	fmt.Println("u.Name", u.Name, "u.Ip: ", u.Ip, "u.Device", u.Device)
	fmt.Fprintln(ts.LogFile, "u.Name", u.Name, "u.Ip: ", u.Ip, "u.Device", u.Device)
}
