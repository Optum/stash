/*
Stash is a simple static file server in go
Usage:
	-p="8080": port to serve on
	-d=".":    the directory of static files to host

Navigating to http://localhost:8080 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"fmt"
)

var sharedDirectory = "SHARED_DIR"
var token = "TOKEN"
var logLevel = 1

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UP")
}

func main() {
  log.Printf("Starting stash...")
  log.Printf("Token for deleting files is: %s", os.Getenv(token))

	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", "./resources/persistent/", "the directory of static file to host")

	os.Setenv(sharedDirectory, *directory)
	log.Printf("Sharing files at: " + os.Getenv(sharedDirectory))

	flag.Parse()

	mux := http.NewServeMux()

	hostPort := net.JoinHostPort("0.0.0.0", *port)

  mux.HandleFunc("/", landingPage)
	mux.HandleFunc("/resources", resources)
	mux.HandleFunc("/resources/", resources)
	mux.HandleFunc("/health", health)
	log.Printf("Serving requests on HTTP port: %s\n", *port)

	server := http.Server{Handler: mux, Addr: hostPort}
	log.Fatal(server.ListenAndServe())

}
