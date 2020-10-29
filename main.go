package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Matias-Barrios/GoUnitTestingExample/handlers"
)

var _routerprovider handlers.IHandlerProvider = handlers.HandlerProvider{}

func main() {
	r := _routerprovider.GetRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Running in port 8000!")
	srv.ListenAndServe()
}
