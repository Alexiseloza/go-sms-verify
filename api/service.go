package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("Code NOT Valid")
	}
	return nil
}
