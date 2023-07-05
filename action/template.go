package main

import "fmt"

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type otp struct {
	iOtp IOtp
}

func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

type sms struct {
	otp
}

func (s *sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type emial struct {
	otp
}

func (e *emial) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (e *emial) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *emial) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (e *emial) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending sms: %s\n", message)
	return nil
}

func main() {
	smsOTP := &sms{}
	o := otp{smsOTP}
	err := o.genAndSendOTP(4)
	if err != nil {
		return
	}

	emailOTP := &emial{}
	o = otp{
		iOtp: emailOTP,
	}
	err2 := o.genAndSendOTP(4)
	if err2 != nil {
		return
	}
}
