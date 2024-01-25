package data

type OTPData struct {
	Phonenumber string `json:"phonenumber,omitempty" validate:"required"`
}

type VerifyData struct {
	User *OTP `json:"user,omitempty validate:"required"`
	Code string `json:"code,omitempty validate:"required"`
}