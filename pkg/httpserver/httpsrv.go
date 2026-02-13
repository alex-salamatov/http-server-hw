package httpserver

import (
	"fmt"
	"net/http"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {

	download_folder := r.URL.Query().Get("id")
	if download_folder == "" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Download files from ")
	fmt.Fprintln(w, download_folder)
}

func RunHttpSrv() {
	http.HandleFunc("/hw", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/download/", HandleDownload)

	fmt.Println("SServer started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
