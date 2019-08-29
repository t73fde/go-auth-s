// Simple HTTP Basic Authentication web server
// Autor: Detlef Stern, <detlef.stern@hs-heilbronn.de>, (c) 2019
// License: Apache 2.0, see LICENSE

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func checkAuth(username, password string) bool {
	return username != "" && username[0] != 'x' && password != ""
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
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
	uPort := flag.Uint("p", 9876, "port number")
	flag.Parse()

	sPort := strconv.FormatUint(uint64(*uPort), 10)
	fmt.Println("Listening on port", sPort)

	http.HandleFunc("/", handleAuth)
	log.Fatal(http.ListenAndServe(":"+sPort, nil))
}
