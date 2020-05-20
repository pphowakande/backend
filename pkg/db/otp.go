package db

import (
	"backend/pkg/io"
	"context"
	"log"
	"time"
)

func (s Service) CreateOTP(ctx context.Context, data io.AthUserOTP) (err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save otp error: %v", d.Error)
		return d.Error
	}
	return d.Error
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (s Service) VerifyOTP(ctx context.Context, data io.Verify) (err error) {
	// check if otp exists
	var u io.AthUserOTP
	err = s.DB.Where(io.AthUserOTP{OTPNO: data.Code, OTPType: data.Type, IsActive: false}).Find(&u).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
		return err
	}
	// if exists , check its expiration time
	expiry_date := u.CreatedAt.Add(time.Second * time.Duration(u.OTPExpiry))
	dt := time.Now()

	otpActive := false

	// check if expiry date is greater than current date
	if inTimeSpan(u.CreatedAt, expiry_date, dt) {
		otpActive = true
		//fmt.Println(dt, "is between", u.CreatedAt, "and", expiry_date, ".")
	} else {
		otpActive = false
		//fmt.Println("otp is expired")
		return err
	}

	if otpActive == true {
		u.IsActive = true
		err = s.DB.Save(&u).Error

		if err != nil {
			log.Printf("Query failed error: %v", err)
			return err
		}

		//fmt.Println("isactive from otp table is updated----------------------")
		if data.Type == "Email" {
			// set user to isactive inside table "ath_users"

			var user io.AthUser
			user.ID = u.UserID

			err = s.DB.Model(user).Update("is_active", true).Error
			if err != nil {
				log.Printf("Query failed error: %v", err)
				return err
			}
		}
		return
	}
	return
}
