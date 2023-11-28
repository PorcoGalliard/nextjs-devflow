package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	listenAddr := flag.String("listen Addr", ":5000", "using for opening the server")
	flag.Parse()
}