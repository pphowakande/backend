package db

import (
	"backend/pkg/io"
	"context"
	"fmt"
	"log"
)

func (s Service) CreateVenue(ctx context.Context, data io.AthVenues) (newVenue io.AthVenues, err error) {
	var u io.AthVenues
	err = s.DB.Where(io.AthVenues{EmailAddress: data.EmailAddress}).Find(&u).Error
	if err == nil {
		err = fmt.Errorf(`email address already exist`)
		return
	}
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return data, nil
}

func (s Service) EditVenue(ctx context.Context, data io.AthVenues) (newVenue io.AthVenues, err error) {
	var u io.AthVenues
	err = s.DB.Where(io.AthVenues{VenueID: data.VenueID}).Find(&u).Error

	if data.VenueName != "" {
		u.VenueName = data.VenueName
	}

	if data.Description != "" {
		u.Description = data.Description
	}

	if data.ContactNo != "" {
		u.ContactNo = data.ContactNo
	}

	if data.EmailAddress != "" {
		u.EmailAddress = data.EmailAddress
	}

	if data.Address != "" {
		u.Address = data.Address
	}

	if data.GetLag != "" {
		u.GetLag = data.GetLag
	}

	if data.GetLong != "" {
		u.GetLong = data.GetLong
	}

	if data.Amenities != "" {
		u.Amenities = data.Amenities
	}

	u.UpdatedBy = data.UpdatedBy

	d := s.DB.Save(&u)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return data, nil
}

func (s Service) CreateVenueHoliday(ctx context.Context, data io.AthVenueHolidays) (newVenueHoliday io.AthVenueHolidays, err error) {
	//var u io.AthVenueHolidays
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return data, nil
}

func (s Service) DeleteVenueHoliday(ctx context.Context, data io.DeleteVenueHolidays) (err error) {
	var u io.AthVenueHolidays
	err = s.DB.Where(io.AthVenueHolidays{HolidayID: data.HolidayID, VenueID: data.VenueID}).Find(&u).Error
	u.IsActive = true
	err = s.DB.Save(&u).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return err
}

func (s Service) SaveHoursOfOperation(ctx context.Context, data io.AthVenueHours) (newVenueHhour io.AthVenueHours, err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return data, nil
}

func (s Service) EditHoursOfOperation(ctx context.Context, data io.AthVenueHours) (newVenueHhour io.AthVenueHours, err error) {
	var u io.AthVenueHours
	err = s.DB.Where(io.AthVenueHours{HourID: data.HourID}).Find(&u).Error

	if data.OpeningTime != "" {
		u.OpeningTime = data.OpeningTime
	}

	if data.ClosingTime != "" {
		u.ClosingTime = data.ClosingTime
	}

	u.UpdatedBy = data.UpdatedBy

	d := s.DB.Save(&u)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return data, nil
}

func (s Service) DeleteHoursOfOperation(ctx context.Context, data io.AthVenueHours) (err error) {
	var u io.AthVenueHours
	err = s.DB.Where(io.AthVenueHours{HourID: data.HourID}).Find(&u).Error
	u.IsActive = true
	err = s.DB.Save(&u).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return err
}
