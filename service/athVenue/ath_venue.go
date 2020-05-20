package athUser

import (
	"backend/pkg/db"
	"backend/pkg/io"
	"context"
)

type AthVenueService interface {
	CreateVenue(ctx context.Context, u io.AthVenues) (res io.Response)
	EditVenue(ctx context.Context, u io.AthVenues) (res io.Response)
	CreateVenueHoliday(ctx context.Context, u io.AthVenueHolidays) (res io.Response)
	SaveHoursOfOperation(ctx context.Context, u io.AthVenueHours) (res io.Response)
	EditHoursOfOperation(ctx context.Context, u io.AthVenueHours) (res io.Response)
	DeleteVenueHoliday(ctx context.Context, u io.DeleteVenueHolidays) (res io.Response)
	DeleteHoursOfOperation(ctx context.Context, u io.AthVenueHours) (res io.Response)
}

type athVenueService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthVenueService(DbRepo db.Repository) AthVenueService {
	return &athVenueService{
		DbRepo: DbRepo,
	}
}

func (b *athVenueService) CreateVenue(ctx context.Context, u io.AthVenues) (res io.Response) {
	newVenue, err := b.DbRepo.CreateVenue(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new venue profile")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["venue_id"] = newVenue.VenueID

	res.Data = data
	res = io.SuccessMessage(data, "Venue Profile created")
	return
}

func (b *athVenueService) CreateVenueHoliday(ctx context.Context, u io.AthVenueHolidays) (res io.Response) {
	newVenueHoliday, err := b.DbRepo.CreateVenueHoliday(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new venue holiday")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["venue_holiday_id"] = newVenueHoliday.HolidayID

	res.Data = data
	res = io.SuccessMessage(data, "Venue Holiday created")
	return
}

func (b *athVenueService) SaveHoursOfOperation(ctx context.Context, u io.AthVenueHours) (res io.Response) {

	newVenueHours, err := b.DbRepo.SaveHoursOfOperation(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new venue hour")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["hour_id"] = newVenueHours.HourID

	res.Data = data
	res = io.SuccessMessage(data, "Venue Hour Added")
	return
}

func (b *athVenueService) EditHoursOfOperation(ctx context.Context, u io.AthVenueHours) (res io.Response) {

	_, err := b.DbRepo.EditHoursOfOperation(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error updating venue hour")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "Venue Hour updated")
	return
}

func (b *athVenueService) DeleteVenueHoliday(ctx context.Context, u io.DeleteVenueHolidays) (res io.Response) {
	err := b.DbRepo.DeleteVenueHoliday(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error deleting venue holiday")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "Venue holiday deleted")
	return
}

func (b *athVenueService) EditVenue(ctx context.Context, u io.AthVenues) (res io.Response) {
	_, err := b.DbRepo.EditVenue(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error updating venue")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "Venue updated")
	return
}

func (b *athVenueService) DeleteHoursOfOperation(ctx context.Context, u io.AthVenueHours) (res io.Response) {
	err := b.DbRepo.DeleteHoursOfOperation(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error deleting venue hours")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "Venue hour deleted")
	return
}
