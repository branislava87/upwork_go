package otp

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//send request to twilio for verification sms
var (
	SendTwilioSMS = sendTwilioSMS
)

func sendTwilioSMS(countryCode int, phoneNumber, locale, apiKey string) {

	urlAddress := os.Getenv("TWILIO_SEND_ADDRESS")
	client := &http.Client{}

	formValues := url.Values{}
	formValues.Add("via", "sms")
	formValues.Add("country_code", strconv.Itoa(countryCode))
	formValues.Add("phone_number", phoneNumber)
	formValues.Add("locale", locale)
	formValues.Add("code_length", os.Getenv("CODE_LENGTH"))

	req, err := http.NewRequest(http.MethodPost, urlAddress, strings.NewReader(formValues.Encode()))
	if err != nil {
		return
	}

	req.Header.Add("X-Authy-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpResponse, err := client.Do(req)
	if err != nil {
		return
	}
	defer httpResponse.Body.Close()
}
