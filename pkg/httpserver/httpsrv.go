package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleFilesDownload(w http.ResponseWriter, r *http.Request) {

	var allowedDirs = map[string]string{
		"books": "../../go_books",
		"media": "../../media"}

	dir := strings.TrimPrefix(r.URL.Path, "/files/")
	path, ok := allowedDirs[dir]
	if !ok {
		http.NotFound(w, r)
		return
	}

	fs := http.FileServer(http.Dir(path))
	http.StripPrefix("/files/"+dir, fs).ServeHTTP(w, r)
}

func RunHttpSrv() {
	http.HandleFunc("/hw", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/files/", HandleFilesDownload)

	fmt.Println("SServer started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
