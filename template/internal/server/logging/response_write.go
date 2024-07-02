package logging

import (
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, b []byte) {
	_, err := w.Write(b)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}
