package main

import (
	"flag"
	"fmt"

	"github.com/alex-salamatov/http_server_hw/pkg/httpserver"
)

func main() {
	fmt.Println("HTTP server Hello World with files download")

	// Command‑line flags
	dir := flag.String("dir", ".", "directory for downloading")
	flag.Parse()

	httpserver.RunHttpSrv(*dir)
}
