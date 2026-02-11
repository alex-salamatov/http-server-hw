package main

import (
	"fmt"

	"github.com/alex-salamatov/http_server_hw/pkg/httpserver"
)

func main() {
	fmt.Println("HTTP server Hello World")

	httpserver.RunHttpSrv()
}
