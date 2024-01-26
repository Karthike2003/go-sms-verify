package api

import (
	"errors"	
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWtihParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	resp, err := client.Verify.V2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CheckVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)
	resp, err := client.Verify.V2.CheckVerification(envSERVICESID(), params)
	if err != nil {
		return err
	}
}

if *resp.Status != "approved" {
	return errors.New("invalid code")
}