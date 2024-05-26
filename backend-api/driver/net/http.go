package net

import (
	"log"
	"net"
	"net/http"
)

func ServeHTTP(handler http.Handler, port string) error {
	server := &http.Server{
		Handler: handler,
	}
	// accept IPv4 also IPv6
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	log.Printf("Listening HTTP connection on %s", port)
	if err := server.Serve(listener); err != nil {
		return err
	}
	return nil
}
