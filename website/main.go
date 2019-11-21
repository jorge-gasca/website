// Simple HTTP server

package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"website/handlers"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("jgasca.dev"),
		Cache:      autocert.DirCache("certs"),
	}
	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	http.HandleFunc("/", handlers.IndexHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//redirect HTTP	
	go func(){
			err := http.ListenAndServe(":http", certManager.HTTPHandler(nil))
			if err != nil {
				log.Fatal(err)
				return
			}
	}()

	log.Fatal(server.ListenAndServeTLS("", ""))
}
