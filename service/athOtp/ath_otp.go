package athUser

import (
	"backend/pkg/db"
	"backend/pkg/io"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AthOtpService interface {
	CreateOTP(ctx context.Context, u io.AthUserOTP) (res io.Response)
	VerifyOTP(ctx context.Context, u io.Verify) (res io.Response)
}

type athOtpService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthOtpService(DbRepo db.Repository) AthOtpService {
	return &athOtpService{
		DbRepo: DbRepo,
	}
}

func (b *athOtpService) CreateOTP(ctx context.Context, u io.AthUserOTP) (res io.Response) {
	err := b.DbRepo.CreateOTP(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error saving OTP in database")
		res.Error = err
		return
	}

	uri, err := url.ParseRequestURI("http://in-cs-app.sit.n3b.bookmyshow.org/sendsms")
	fmt.Println("uri : ", uri)
	fmt.Println("err : ", err)
	if err != nil {
		fmt.Println("err here : ", err)
		res = io.FailureMessage(nil, "Error Sending OTP to given mobile number")
		res.Error = err
		return
		// handle err
	}

	//var testMessage string
	testMessage := "OTP to verify your mobile number is " + u.OTPNO

	values := uri.Query()
	values.Set("to", u.Contact)
	values.Set("type", "Confirmation")
	values.Set("message", testMessage)
	values.Set("country", "91")
	values.Set("medium", "test")
	values.Set("refcode", "test")

	fmt.Println("values : ", values)

	uri.RawQuery = values.Encode()

	fmt.Println("uri after : ", uri)

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	fmt.Println("req : ", req)
	if err != nil {
		fmt.Println("err here 2 : ", err)
		// handle err
		res = io.FailureMessage(nil, "Error Sending OTP to given mobile number")
		res.Error = err
		return
	}

	resp, err := http.DefaultClient.Do(req)
	fmt.Println("resp : ", resp)
	if err != nil {
		fmt.Println("err here 3 : ", err)
		// handle err
		res = io.FailureMessage(nil, "Error Sending OTP to given mobile number")
		res.Error = err
		return
	}

	defer resp.Body.Close()

	fmt.Println("resp.StatusCode : ", resp.StatusCode)

	if resp.StatusCode != 200 {
		res = io.FailureMessage(nil, "Error Sending OTP to given mobile number")
		res.Error = err
	}

	/*
		emailMsg := "Email verification code:" + PhoneVerifyToken
		res.Error = email.NewEmailSend(u.Email, "Verify Email", emailMsg)
		if res.Error != nil {
			b.DbRepo.Delete(ctx, io.AthUser{Email: u.Email})
			res = io.FailureMessage(res.Error, "Failed to send email")
			return
		}
	*/

	data := make(map[string]interface{})
	data["verify_token"] = u.OTPNO

	res = io.SuccessMessage(data, "OTP has been sent to given mobile number")

	return
}

func (b *athOtpService) VerifyOTP(ctx context.Context, u io.Verify) (res io.Response) {
	err := b.DbRepo.VerifyOTP(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, err.Error())
		res.Error = err
		return
	}

	if u.Email != "" {

		requestBody, err := json.Marshal(map[string]string{
			"email":   u.Email,
			"ac":      "WEBIN",
			"subject": "test",
			"body":    "test",
			"eticket": "N",
			"tid":     "0",
		})

		resp, err := http.Post("http://in-cs-app.sit.n3b.bookmyshow.org/send", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("err here 2 : ", err)
			// handle err
			res = io.FailureMessage(nil, "Error Sending email to given email address")
			res.Error = err
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			res = io.SuccessMessage(nil, "OTP has been verified successfully")
			return
		}
		res = io.FailureMessage(nil, "Error Sending email to given email address")
		res.Error = err
	}
	res = io.SuccessMessage(nil, "OTP has been verified successfully")
	return
}
