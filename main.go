// Simple HTTP Basic Authentication web server
// Autor: Detlef Stern, <detlef.stern@hs-heilbronn.de>, (c) 2019
// License: Apache 2.0, see LICENSE

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func checkAuth(username, password string) bool {
	if username == "" || username[0] == 'x' || password == "" {
		return false
	}
	if username[0] == 'q' && password != username {
		return false
	}
	return true
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "text/plain")
	h.Set("Content-Length", "0")

	var code int

	if username, password, ok := r.BasicAuth(); ok {
		if checkAuth(username, password) {
			code = http.StatusOK
		} else {
			code = http.StatusForbidden
		}
	} else {
		h.Set("WWW-Authenticate", `Basic realm="Default"`)
		code = http.StatusUnauthorized
	}
	w.WriteHeader(code)
}

func main() {
	port := flag.Uint("p", 9876, "port number")
	flag.Parse()

	var srv *server
	fmt.Printf("Listening on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), srv))
}
