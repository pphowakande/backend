package io

import (
	"time"
)

/*

create table ath_user_otp (
  otp_id	integer(11)	auto_increment primary key not null,
  user_id	int(11) references ath_users(user_id),
	otp_no	varchar(100)	not	NULL,
	otp_type	varchar(30)	not	NULL,
  otp_expiry	int(11)	not	NULL	,
  created_at	datetime	default	NULL,
  updated_at	datetime	default	NULL,
  is_active	tinyint(1)	default	1
);
*/

type AthUserOTP struct {
	ID        int       `gorm:"column:otp_id;primary_key" json:"otp_id"`
	UserID    int       `gorm:"column:user_id;not null" json:"user_id"`
	Contact   string    `gorm:"column:contact_no" json:"contact_no"`
	OTPNO     string    `gorm:"column:otp_no" json:"otp_no"`
	OTPType   string    `gorm:"column:otp_type" json:"otp_type"`
	OTPExpiry int       `gorm:"column:otp_expiry" json:"otp_expiry"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
}
