package httpserver

import (
	"fmt"
	"net/http"
)

func RunHttpSrv() {
	http.HandleFunc("/hw", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
