package io

import (
	"time"
)

/*

create table ath_users (
  user_id	integer(11)	auto_increment primary key not null,
  email_address	varchar(100)	not null,
  password	varchar(30)	default	NULL,
  reset_password_token	varchar(50)	default	NULL,
  reset_password_token_created_at	datetime	default	NULL,
  last_password_reset_at	datetime	default	NULL,
  user_source	tinyint(1)	default	NULL,
  last_login_at	datetime	default	NULL,
  last_login_ip	varchar(30)	not null	,
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11)	references ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	references ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	references ath_users(user_id)
);
*/
type AthUser struct {
	ID                          int       `gorm:"column:user_id;primary_key" json:"user_id"`
	Email                       string    `gorm:"column:email_address;unique;not null" json:"email_address"`
	Password                    string    `gorm:"column:password" json:"password"`
	ResetPasswordToken          string    `gorm:"column:reset_password_token" json:"reset_password_token"`
	ResetPasswordTokenCreatedAt time.Time `gorm:"column:reset_password_token_created_at" json:"reset_password_token_created_at"`
	LastPasswordResetAt         time.Time `gorm:"column:last_password_reset_at" json:"last_password_reset_at"`
	UserSource                  string    `gorm:"column:user_source" json:"user_source"`
	LastLoginAt                 time.Time `gorm:"column:last_login_at" json:"last_login_at"`
	LastLoginIp                 string    `gorm:"column:last_login_ip" json:"last_login_ip"`
	IsActive                    bool      `gorm:"column:is_active" json:"is_active"`
	Models
}

/*
create table ath_user_profiles (
  user_profile_id	integer(11)	auto_increment primary key not null,
  user_id	int(11)	,
  first_name	varchar(50)	not null	,
  last_name	varchar(50)	default	NULL,
  contact_no	varchar(30)	not null,
  gender	enum('m', 'f', 'o')	default	NULL,
  dob	date	default	NULL,
  profile_image	varchar(100)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);
*/

type AthUserProfile struct {
	UserProfileID int    `gorm:"column:user_profile_id;primary_key" json:"user_profile_id"`
	UserID        int    `gorm:"column:user_id" json:"user_id"`
	FirstName     string `gorm:"column:first_name" json:"first_name"`
	LastName      string `gorm:"column:last_name" json:"last_name"`
	ContactNo     string `gorm:"column:contact_no" json:"contact_no"`
	Gender        string `gorm:"column:gender" json:"gender"`
	DOB           string `gorm:"column:dob" json:"dob"`
	ProfileImage  string `gorm:"column:profile_image" json:"profile_image"`
	Models
}

type Verify struct {
	Code  string `gorm:"column:code" json:"code"`
	Type  string `gorm:"column:type"  json:"type"`
	Email string `gorm:"column:email" json:"email"`
}

type EmailVerify struct {
	Email string `gorm:"column:email" json:"email"`
}
