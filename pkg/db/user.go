package db

import (
	"backend/pkg/io"
	"context"
	"fmt"
	"log"
	"time"
)

func (s Service) CreateUser(ctx context.Context, data io.AthUser) (newUser io.AthUser, err error) {
	var u io.AthUser
	err = s.DB.Where(io.AthUser{Email: data.Email}).Find(&u).Error
	if err == nil {
		err = fmt.Errorf(`email address already exist`)
		return
	}
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
	}
	return data, nil
}

func (s Service) CreateUserProfile(ctx context.Context, data io.AthUserProfile) (newUser io.AthUserProfile, err error) {
	var u io.AthUserProfile
	err = s.DB.Where(io.AthUserProfile{UserID: data.UserID}).Find(&u).Error
	if err == nil {
		err = fmt.Errorf(`User Profile already exist`)
		return
	}
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
	}
	return data, nil
}

func (s Service) EditUserProfile(ctx context.Context, data io.AthUserProfile) (newUser io.AthUserProfile, err error) {
	var u io.AthUserProfile
	err = s.DB.Where(io.AthUserProfile{UserProfileID: data.UserProfileID}).Find(&u).Error

	if data.FirstName != "" {
		u.FirstName = data.FirstName
	}

	if data.LastName != "" {
		u.LastName = data.LastName
	}

	if data.ContactNo != "" {
		u.ContactNo = data.ContactNo
	}

	if data.Models.UpdatedBy != 0 {
		u.UpdatedBy = data.Models.UpdatedBy
	}

	d := s.DB.Save(&u)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
	}
	return data, nil
}

func (s Service) EditUser(ctx context.Context, data io.AthUser) (newUser io.AthUser, err error) {
	var u io.AthUser
	err = s.DB.Where(io.AthUser{ID: data.ID}).Find(&u).Error

	if data.Email != "" {
		u.Email = data.Email
	}

	if data.Models.UpdatedBy != 0 {
		u.UpdatedBy = data.UpdatedBy
	}

	d := s.DB.Save(&u)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
	}
	return data, nil
}

func (s Service) VerifySignUpTokenUser(ctx context.Context, data io.Verify) (err error) {

	var u io.AthUser
	err = s.DB.Where(io.AthUser{}).Find(&u).Error
	if u.IsActive {
		err = fmt.Errorf(`you are already verified`)
		return
	}
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}
	if data.Type == "email" {
		u.IsActive = true
		err = s.DB.Save(&u).Error
	}
	return
}

func (s Service) VerifyEmail(ctx context.Context, data io.EmailVerify) (err error) {
	var u io.AthUser
	err = s.DB.Where(io.AthUser{}).Find(&u).Error
	if u.IsActive {
		err = fmt.Errorf(`you are already verified`)
		return
	}
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}

	u.IsActive = true
	err = s.DB.Save(&u).Error

	return
}

func (s Service) LoginUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error) {

	err = s.DB.Where(io.AthUser{Email: data.Email}).Find(&user).Error
	if err == nil {
		user.LastLoginAt = time.Now()
		user.LastLoginIp = ""
		s.DB.Save(&user)
	}
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}
	return
}

func (s Service) ForgotPasswordUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error) {

	err = s.DB.Where(io.AthUser{Email: data.Email}).Find(&user).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
		return data, err
	}
	user.LastPasswordResetAt = time.Now()
	user.ResetPasswordTokenCreatedAt = time.Now()
	user.ResetPasswordToken = data.ResetPasswordToken
	s.DB.Save(&user)
	return data, nil
}

func (s Service) ResetPasswordUser(ctx context.Context, data io.AthUser) (err error) {
	var user io.AthUser
	err = s.DB.Where(io.AthUser{Email: data.Email}).Find(&user).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
		return err
	}
	user.Password = data.Password
	s.DB.Save(&user)
	return nil
}

func (s Service) DeleteUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error) {

	err = s.DB.Where(io.AthUser{Email: data.Email}).Delete(&user).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}
	return
}
