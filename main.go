package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		name = strings.ToLower(name)
		if name != "user-agent" {
			continue
		}
		for _, h := range headers {
			h = strings.ToLower(h)
			if strings.Contains(h, "mobile") {
				fmt.Fprintf(w, "<h1 style=\"font-size: 48px;\">Mobile</h1>")
				return
			}
		}
		fmt.Fprintf(w, "<h1 style=\"font-size: 48px;\">Desktop</h1>")
	}
}

func main() {
	http.HandleFunc("/pineocleen/connected", headers)
	err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/scio.dev/fullchain.pem", "/etc/letsencrypt/live/scio.dev/privkey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	http.ListenAndServe(":80", nil)
}
