package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ammar-alaa/common"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR",":3000")
)

func main() {


	mux := http.NewServeMux()
	handler :=NewHander()
	handler.registerRoutes(mux)

	log.Printf("Starting http server at %s",httpAddr)

	// Start the server listening
	if err:= http.ListenAndServe(httpAddr,mux); err != nil {
		log.Fatal("failed to start the http server")
	}
}