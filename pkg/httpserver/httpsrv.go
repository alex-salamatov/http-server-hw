package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

func RunFileServer(download_dir string) {
	fs := http.FileServer(http.Dir(download_dir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ext := path.Ext(r.URL.Path)
		if ext != "" && ext != "/" { //handle filename
			filename := path.Base(r.URL.Path)
			w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"") //force browser to download file
		}
		fs.ServeHTTP(w, r)
	})

	log.Printf("Serving directory %q on %s\n", download_dir, ":8080")
}

func RunHttpSrv(download_dir string) {
	http.HandleFunc("/hw", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	RunFileServer(download_dir)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
