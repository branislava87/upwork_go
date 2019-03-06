package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/upwork_go/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	server := &http.Server{}
	server.Handler = handlers.NewOTP()

	sslCertificate := os.Getenv("SSL_CERTIFICATE")
	sslKeystore := os.Getenv("SSL_KEYSTORE")

	if sslCertificate == "" || sslKeystore == "" {
		server.Addr = os.Getenv("SERVER_ADDR")
		fmt.Printf("Serving HTTP requests at %v\n", server.Addr)
		err = server.ListenAndServe()
	} else {
		server.Addr = os.Getenv("SERVER_ADDR_SECURE")
		fmt.Printf("Serving HTTPS requests at %v\n", server.Addr)
		err = server.ListenAndServeTLS(os.Getenv("SSL_CERTIFICATE"), os.Getenv("SSL_KEYSTORE"))
	}
}
