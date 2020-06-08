package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-http-utils/logger"

	"github.com/tj/spa"
)

func main() {
	addr := flag.String("address", ":3000", "Server bind address.")

	// usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s: [dir]\n", os.Args[0])
		flag.PrintDefaults()
	}

	// parse
	flag.Parse()
	dir := flag.Arg(0)

	if dir == "" {
		dir = "."
	}

	// server
	log.Printf("Serving files from %q", dir)
	server := spa.Server{
		Dir: dir,
	}

	// logging
	h := logger.Handler(server, os.Stdout, logger.CommonLoggerType)

	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, h))
}
