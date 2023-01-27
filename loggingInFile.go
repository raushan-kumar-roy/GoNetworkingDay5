package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetLevel(log.InfoLevel)
	http.HandleFunc("/", handleconnection)
	http.ListenAndServe(":8000", nil)
}

func handleconnection(writer http.ResponseWriter, request *http.Request) {
	log.WithFields(log.Fields{
		"method": request.Method,
		"path":   request.URL.Path,
		"ip":     request.RemoteAddr,
	}).Info("Received a request")
	writer.Write([]byte("Hello, Client!"))
}
