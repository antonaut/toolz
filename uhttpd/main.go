package uhttpd

import (
	"log"
	"net/http"
)

const (
	PORT = ":8080"
	DIR  = "."
)

func Serve() {
	log.Printf("Serving '%s' on port %s", DIR, PORT)
	log.Fatal(http.ListenAndServe(PORT, http.FileServer(http.Dir(DIR))))
}
