package db

import (
	"backend/pkg/io"
	"context"
	"log"
)

func (s Service) CreateFacility(ctx context.Context, data io.AthFacilities) (newFacility io.AthFacilities, err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
		return data, d.Error
	}
	return data, nil
}

func (s Service) EditFacility(ctx context.Context, data io.AthFacilities) (newFacility io.AthFacilities, err error) {
	var u io.AthFacilities
	err = s.DB.Where(io.AthFacilities{FacilityID: data.FacilityID}).Find(&u).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}

	if data.FacilityName != "" {
		u.FacilityName = data.FacilityName
	}
	if data.FacilitySportCategories != "" {
		u.FacilitySportCategories = data.FacilitySportCategories
	}

	if data.FacilityBasePrice != 0 {
		u.FacilityBasePrice = data.FacilityBasePrice
	}

	if data.UpdatedBy != 0 {
		u.UpdatedBy = data.UpdatedBy
	}

	err = s.DB.Save(&u).Error
	return u, err
}

func (s Service) CreateFacilitySlots(ctx context.Context, data io.AthFacilitySlots) (newFacilitySlots io.AthFacilitySlots, err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
		return data, d.Error
	}
	return data, nil
}

func (s Service) EditFacilitySlots(ctx context.Context, data io.AthFacilitySlots) (newFacilitySlots []io.AthFacilitySlots, err error) {
	var u []io.AthFacilitySlots
	err = s.DB.Where(io.AthFacilitySlots{FacilityID: data.FacilityID, SlotDays: data.SlotDays, SlotType: data.SlotType}).Find(&u).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
		return u, err
	}

	for _, eachSlot := range u {
		eachSlot.SlotPrice = data.SlotPrice
		err = s.DB.Save(&eachSlot).Error
	}

	return u, err
}

func (s Service) BookFacility(ctx context.Context, data io.AthFacilityBookings) (newFacilityBookings io.AthFacilityBookings, err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
		return data, d.Error
	}
	return data, nil
}

func (s Service) BookFacilitySlots(ctx context.Context, data io.AthFacilityBookingSlots) (newFacilityBookingSlot io.AthFacilityBookingSlots, err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
		return data, d.Error
	}
	return data, nil
}

func (s Service) GetSportCategories(ctx context.Context) (sportCategories io.AthSportCategories, err error) {
	var u io.AthSportCategories
	err = s.DB.Where(io.AthSportCategories{}).Find(&u).Error
	if err != nil {
		log.Printf("Query failed error: %v", err)
	}
	return u, err
}

func (s Service) AddFacilityCustomRates(ctx context.Context, data io.AthFacilityCustomRates) (customRates io.AthFacilityCustomRates, err error) {
	d := s.DB.Save(&data)
	if d.Error != nil {
		log.Printf("Failed to save error: %v", d.Error)
		return data, d.Error
	}
	return data, nil
}
