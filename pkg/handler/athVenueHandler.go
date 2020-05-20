package handler

import (
	venuepb "backend/api/athVenue/v1"
	"backend/pkg/io"
	"context"
	"strings"
)

func (g *grpcServer) CreateVenue(ctx context.Context, req *venuepb.CreateVenueRequest) (*venuepb.CreateVenueReply, error) {

	venueRequest := io.AthVenues{
		CompanyID:    int(req.CompanyId),
		VenueName:    req.Name,
		Description:  req.Description,
		EmailAddress: req.Email,
		ContactNo:    req.Phone,
		Address:      req.Address,
		GetLag:       req.GetLat,
		GetLong:      req.GetLong,
		Amenities:    req.Amenities,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	// get venue id from venue service
	venueServiceRes := g.athVenueService.CreateVenue(ctx, venueRequest)

	storeErr := ""
	res := &venuepb.CreateVenueReply{}

	if venueServiceRes.Error != nil {
		storeErr = venueServiceRes.Error.Error()
		res.Status = venueServiceRes.Success
		res.Message = venueServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	venueData := venueServiceRes.Data.(map[string]interface{})
	venueID := venueData["venue_id"].(int)

	// save hours of operation
	for key, val := range req.Hours {
		splittedTime := strings.Split(val, "-")

		hoursOfOpeRequest := io.AthVenueHours{
			Day:         int(key),
			VenueID:     venueID,
			OpeningTime: splittedTime[0],
			ClosingTime: splittedTime[1],
			Models: io.Models{
				CreatedBy: int(req.CreatedBy),
			},
		}
		hoursOfOperationRes := g.athVenueService.SaveHoursOfOperation(ctx, hoursOfOpeRequest)

		if hoursOfOperationRes.Error != nil {
			storeErr = hoursOfOperationRes.Error.Error()
			res.Status = hoursOfOperationRes.Success
			res.Message = hoursOfOperationRes.Message
			res.Error = storeErr
			return res, nil
		}
	}

	res = &venuepb.CreateVenueReply{
		Data: &venuepb.CreateVenuepReplyData{
			VenueId: int32(venueID),
		},
		Status:  venueServiceRes.Success,
		Message: venueServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) EditVenue(ctx context.Context, req *venuepb.EditVenueRequest) (*venuepb.GenericReply, error) {
	venueRequest := io.AthVenues{
		VenueID:      int(req.VenueId),
		VenueName:    req.Name,
		Description:  req.Description,
		EmailAddress: req.Email,
		ContactNo:    req.Phone,
		Address:      req.Address,
		GetLag:       req.GetLat,
		GetLong:      req.GetLong,
		Amenities:    req.Amenities,
		Models: io.Models{
			UpdatedBy: int(req.UpdatedBy),
		}}

	venueServiceRes := g.athVenueService.EditVenue(ctx, venueRequest)

	storeErr := ""
	res := &venuepb.GenericReply{}

	if venueServiceRes.Error != nil {
		storeErr = venueServiceRes.Error.Error()
		res.Status = venueServiceRes.Success
		res.Message = venueServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	// save hours of operation
	for key, val := range req.Hours {
		splittedTime := strings.Split(val, "-")

		hoursOfOpeRequest := io.AthVenueHours{
			HourID:      int(key),
			OpeningTime: splittedTime[0],
			ClosingTime: splittedTime[1],
			Models: io.Models{
				UpdatedBy: int(req.UpdatedBy),
			},
		}
		hoursOfOperationRes := g.athVenueService.EditHoursOfOperation(ctx, hoursOfOpeRequest)

		if hoursOfOperationRes.Error != nil {
			storeErr = hoursOfOperationRes.Error.Error()
			res.Status = hoursOfOperationRes.Success
			res.Message = hoursOfOperationRes.Message
			res.Error = storeErr
			return res, nil
		}
	}

	if req.DeleteHours != nil {
		for _, hour := range req.DeleteHours {
			hoursOfOpeRequest := io.AthVenueHours{
				HourID: int(hour),
			}

			hoursOfOperationRes := g.athVenueService.DeleteHoursOfOperation(ctx, hoursOfOpeRequest)

			if hoursOfOperationRes.Error != nil {
				storeErr = hoursOfOperationRes.Error.Error()
				res.Status = hoursOfOperationRes.Success
				res.Message = hoursOfOperationRes.Message
				res.Error = storeErr
			}
		}
	}

	res = &venuepb.GenericReply{
		Status:  venueServiceRes.Success,
		Message: venueServiceRes.Message,
		Error:   storeErr,
	}

	return res, nil
}

func (g *grpcServer) CreateVenueHoliday(ctx context.Context, req *venuepb.CreateVenueHolidayRequest) (*venuepb.CreateVenueHolidayReply, error) {
	venueholidayRequest := io.AthVenueHolidays{
		Title:   req.Title,
		VenueID: int(req.VenueId),
		Month:   req.Month,
		Year:    req.Year,
		Day:     int(req.Day),

		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	// get venue id from venue service
	venueServiceRes := g.athVenueService.CreateVenueHoliday(ctx, venueholidayRequest)

	storeErr := ""
	res := &venuepb.CreateVenueHolidayReply{}

	if venueServiceRes.Error != nil {
		storeErr = venueServiceRes.Error.Error()
		res.Status = venueServiceRes.Success
		res.Message = venueServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	venueData := venueServiceRes.Data.(map[string]interface{})
	holidayID := venueData["venue_holiday_id"].(int)

	res = &venuepb.CreateVenueHolidayReply{
		Data: &venuepb.CreateVenueHolidayReplyData{
			HolidayId: int32(holidayID),
		},
		Status:  venueServiceRes.Success,
		Message: venueServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) DeleteVenueHoliday(ctx context.Context, req *venuepb.DeleteVenueHolidayRequest) (*venuepb.GenericReply, error) {
	venueholidayRequest := io.DeleteVenueHolidays{
		HolidayID: int(req.HolidayId),
		VenueID:   int(req.VenueId),
	}

	// get venue id from venue service
	venueServiceRes := g.athVenueService.DeleteVenueHoliday(ctx, venueholidayRequest)

	storeErr := ""
	res := &venuepb.GenericReply{}

	if venueServiceRes.Error != nil {
		storeErr = venueServiceRes.Error.Error()
		res.Status = venueServiceRes.Success
		res.Message = venueServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	res = &venuepb.GenericReply{
		Status:  venueServiceRes.Success,
		Message: venueServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}
