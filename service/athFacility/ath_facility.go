package athUser

import (
	"backend/pkg/db"
	"backend/pkg/io"
	"context"
)

type AthFacilityService interface {
	CreateFacility(ctx context.Context, u io.AthFacilities) (res io.Response)
	EditFacility(ctx context.Context, u io.AthFacilities) (res io.Response)
	CreateFacilitySlots(ctx context.Context, u io.AthFacilitySlots) (res io.Response)
	EditFacilitySlots(ctx context.Context, u io.AthFacilitySlots) (res io.Response)
	BookFacility(ctx context.Context, u io.AthFacilityBookings) (res io.Response)
	BookFacilitySlots(ctx context.Context, u io.AthFacilityBookingSlots) (res io.Response)
	AddFacilityCustomRates(ctx context.Context, u io.AthFacilityCustomRates) (res io.Response)
}

type athFacilityService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthFacilityService(DbRepo db.Repository) AthFacilityService {
	return &athFacilityService{
		DbRepo: DbRepo,
	}
}

func (b *athFacilityService) CreateFacility(ctx context.Context, u io.AthFacilities) (res io.Response) {
	newFacility, err := b.DbRepo.CreateFacility(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new facility")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["facility_id"] = newFacility.FacilityID

	res.Data = data
	res = io.SuccessMessage(data, "Facility created")
	return
}

func (b *athFacilityService) EditFacility(ctx context.Context, u io.AthFacilities) (res io.Response) {
	newFacility, err := b.DbRepo.EditFacility(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error editing new facility")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["facility_id"] = newFacility.FacilityID

	res.Data = data
	res = io.SuccessMessage(data, "Facility edited")
	return
}

func (b *athFacilityService) CreateFacilitySlots(ctx context.Context, u io.AthFacilitySlots) (res io.Response) {
	newFacility, err := b.DbRepo.CreateFacilitySlots(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new facility slot")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["facility_slot_id"] = newFacility.FacilityID

	res.Data = data
	res = io.SuccessMessage(data, "Facility slot created")
	return
}

func (b *athFacilityService) EditFacilitySlots(ctx context.Context, u io.AthFacilitySlots) (res io.Response) {
	_, err := b.DbRepo.EditFacilitySlots(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error updating new facility slot")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "Facility slot updated")
	return
}

func (b *athFacilityService) BookFacility(ctx context.Context, u io.AthFacilityBookings) (res io.Response) {
	newFacilityBooking, err := b.DbRepo.BookFacility(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new facility booking")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["facility_booking_id"] = newFacilityBooking.BookingID
	data["facility_booking_no"] = newFacilityBooking.BookingNo

	res.Data = data
	res = io.SuccessMessage(data, "Facility booking created")
	return
}

func (b *athFacilityService) BookFacilitySlots(ctx context.Context, u io.AthFacilityBookingSlots) (res io.Response) {
	newFacilityBookingSlot, err := b.DbRepo.BookFacilitySlots(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new facility booking slot")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["facility_booking_slot_id"] = newFacilityBookingSlot.BookingSlotID

	res.Data = data
	res = io.SuccessMessage(data, "Facility booking slot created")
	return
}

func (b *athFacilityService) AddFacilityCustomRates(ctx context.Context, u io.AthFacilityCustomRates) (res io.Response) {
	_, err := b.DbRepo.AddFacilityCustomRates(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "err adding custom rates for facility")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "Facility custom rate added")
	return
}
