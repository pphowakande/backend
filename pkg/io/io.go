package io

import (
	"time"
)

type Models struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	//DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	CreatedBy int `gorm:"column:created_by" json:"created_by"`
	UpdatedBy int `gorm:"column:updated_by" json:"updated_by"`
	//DeletedBy int `gorm:"column:deleted_by" json:"deleted_by"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success" `
	Error   error       `json:"error"`
}

type LoginRequest struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

type ForgotPasswordRequest struct {
	Email string `gorm:"column:email" json:"email"`
}

type DeleteVenueHolidays struct {
	HolidayID int `gorm:"column:holiday_id" json:"holiday_id"`
	VenueID   int `gorm:"column:venue_id" json:"venue_id"`
}

type ResetPasswordRequest struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

type OtpRequest struct {
	OtpNo     string `gorm:"column:otp_no" json:"otp_no"`
	OtpType   string `gorm:"column:otp_type" json:"otp_type"`
	OtpExpiry string `gorm:"column:otp_expiry" json:"otp_expiry"`
	UserID    string `gorm:"column:user_id" json:"user_id"`
}

func SuccessMessage(data interface{}, msg ...string) Response {
	newMessage := "Success"
	if len(msg) > 0 {
		newMessage = msg[0]
	}
	return Response{
		Data:    data,
		Success: true,
		Message: newMessage,
	}
}
func FailureMessage(err error, msg ...string) Response {
	newMessage := "Failure"
	if len(msg) > 0 {
		newMessage = msg[0]
	}
	return Response{
		Success: false,
		Message: newMessage,
		Error:   err,
	}
}
