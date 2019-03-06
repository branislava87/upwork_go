package otp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	RequestSMSValidationCode = requestSMSValidationCode
)

func requestSMSValidationCode(w http.ResponseWriter, r *http.Request) {
	type validationCodeRequest struct {
		PhoneNumber string `json:"phone_number"`
		CountryCode int    `json:"country_code"`
		Locale      string `json:"locale"`
		APIKey      string `json:"api_key"`
	}

	var err error
	var request = &validationCodeRequest{}

	err = json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		fmt.Println("Request error!")
	}

	var envAPIKey = os.Getenv("API_KEY")
	if request.APIKey == envAPIKey {
		sendTwilioSMS(request.CountryCode, request.PhoneNumber, request.Locale, request.APIKey)
	}

}
