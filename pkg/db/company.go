package db

import (
	"backend/pkg/io"
	"context"
	"log"
)

func (s Service) CreateCompany(ctx context.Context, data io.AthCompany) (newCompany io.AthCompany, err error) {
	err = s.DB.Save(&data).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return data, nil
}

func (s Service) CreateCompanyUser(ctx context.Context, data io.AthCompanyUser) (err error) {
	err = s.DB.Save(&data).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return
}

func (s Service) EditCompany(ctx context.Context, data io.AthCompany) (err error) {
	var u io.AthCompany
	err = s.DB.Where(io.AthCompany{ID: data.ID}).Find(&u).Error

	if data.Name != "" {
		u.Name = data.Name
	}

	if data.Contact != "" {
		u.Contact = data.Contact
	}

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
	return nil
}
