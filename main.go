package main

import (
	"net/http"

	"github.com/bispoman/go-APIs/routes"

	log "github.com/sirupsen/logrus"
)

func main() {

	PORT := 8080

	log.Info("Starting application in port: ", PORT)

	appRoutes := routes.GetRoutes()

	http.Handle("/", appRoutes)

	log.Fatal(http.ListenAndServe(":"+striconv.ParseInt(PORT), nil))
}
