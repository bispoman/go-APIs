package controller

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Info("hello func hit,responding.")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Oh, hi :)")
}
