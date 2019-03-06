package handlers

import (
	"net/http"

	"github.com/upwork_go/otp"
)

func NewOTP() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/protected/json/otp/send", otp.RequestSMSValidationCode)
	return mux
}
