package db

import (
	"backend/pkg/io"
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) LoginAdmin(ctx context.Context, data io.AthAdminUser) (user io.AthAdminUser, err error) {

	err = s.DB.Where(io.AthAdminUser{}).Find(&user).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}
	return
}

func (s Service) SignupAdmin(ctx context.Context, data io.AthAdminUser) (err error) {
	var u io.AthAdminUser
	data.IsActive = true
	err = s.DB.Where(io.AthAdminUser{Email: data.Email}).Find(&u).Error
	if err == nil {
		err = fmt.Errorf(`email address already exist`)
		return
	}
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
	}
	return nil
}
