package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	multipartFileKey = "file"

	csvExtension = ".csv"

	// FUTURE CONSIDERATION: read host and port from env variables
	defaultPort = "8080"

	recordsKey = "records"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@./testData/matrix.csv' "localhost:8080/echo"

func main() {
	handler := NewHandler()
	log.Println("server is running on port " + defaultPort)
	err := http.ListenAndServe(net.JoinHostPort("", defaultPort), handler)
	if err != nil {
		log.Println(fmt.Sprintf("error during http listening: %s", err.Error()))
		return
	}
}
