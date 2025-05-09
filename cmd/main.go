package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MorseAppConverter: ", log.LstdFlags)

	servak := server.NewServer(logger)

	logger.Println("Start server...")

	err := servak.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Fatal()
	}
}
