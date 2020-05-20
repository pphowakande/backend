package handler

import (
	facilitypb "backend/api/athFacility/v1"
	"backend/pkg/io"
	"context"
	"strings"
	"time"

	"github.com/thanhpk/randstr"
)

func (g *grpcServer) CreateFacility(ctx context.Context, req *facilitypb.CreateFacilityRequest) (*facilitypb.CreateFacilityReply, error) {

	facilityRequest := io.AthFacilities{
		VenueID:                 int(req.VenueId),
		FacilityName:            req.Name,
		FacilityBasePrice:       req.BasePrice,
		TimeSlot:                int(req.TimeSlot),
		FacilitySportCategories: req.SportCategories,

		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}
	// get facility id from facility service
	facilityServiceRes := g.athFacilityService.CreateFacility(ctx, facilityRequest)

	storeErr := ""
	res := &facilitypb.CreateFacilityReply{}

	if facilityServiceRes.Error != nil {
		storeErr = facilityServiceRes.Error.Error()
		res.Status = facilityServiceRes.Success
		res.Message = facilityServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	faciityData := facilityServiceRes.Data.(map[string]interface{})
	facilityID := faciityData["facility_id"].(int)

	for weekKey, weekVal := range req.WeekData.Weekdays {

		if weekVal != nil {

			ContainFlag := strings.Contains(weekVal.GetSlot(), "-") // true

			if ContainFlag == true {

				splittedTime := strings.Split(weekVal.GetSlot(), "-")

				facilitySlotRequest := io.AthFacilitySlots{
					FacilityID:   facilityID,
					UserID:       int(req.CreatedBy),
					SlotDays:     "weekdays",
					SlotType:     weekKey,
					SlotPrice:    weekVal.GetPrice(),
					SlotFromTime: splittedTime[0],
					SlotToTime:   splittedTime[1],

					Models: io.Models{
						CreatedBy: int(req.CreatedBy),
					},
				}

				// add facility slots
				FacilitySlotsServiceRes := g.athFacilityService.CreateFacilitySlots(ctx, facilitySlotRequest)
				if FacilitySlotsServiceRes.Error != nil {
					storeErr = FacilitySlotsServiceRes.Error.Error()
					res.Status = FacilitySlotsServiceRes.Success
					res.Message = FacilitySlotsServiceRes.Message
					res.Error = storeErr
					return res, nil
				}
			}
		} else {
			storeErr = "Wrong value for Weekdays"
			res.Status = false
			res.Message = "Wrong value for Weekdays"
			res.Error = storeErr
			return res, nil
		}
	}

	for weekKey, weekVal := range req.WeekData.Weekends {

		if weekVal != nil {

			ContainFlag := strings.Contains(weekVal.GetSlot(), "-") // true

			if ContainFlag == true {

				splittedTime := strings.Split(weekVal.GetSlot(), "-")

				facilitySlotRequest := io.AthFacilitySlots{
					FacilityID:   facilityID,
					UserID:       int(req.CreatedBy),
					SlotDays:     "weekends",
					SlotType:     weekKey,
					SlotPrice:    weekVal.GetPrice(),
					SlotFromTime: splittedTime[0],
					SlotToTime:   splittedTime[1],

					Models: io.Models{
						CreatedBy: int(req.CreatedBy),
					},
				}

				// add facility slots
				FacilitySlotsServiceRes := g.athFacilityService.CreateFacilitySlots(ctx, facilitySlotRequest)
				if FacilitySlotsServiceRes.Error != nil {
					storeErr = FacilitySlotsServiceRes.Error.Error()
					res.Status = FacilitySlotsServiceRes.Success
					res.Message = FacilitySlotsServiceRes.Message
					res.Error = storeErr
					return res, nil
				}
			}
		} else {
			storeErr = "Wrong value for Weekends"
			res.Status = false
			res.Message = "Wrong value for Weekends"
			res.Error = storeErr
			return res, nil
		}
	}

	res = &facilitypb.CreateFacilityReply{
		Data: &facilitypb.CreateFacilityReplyData{
			FacilityId: int32(facilityID),
		},
		Status:  facilityServiceRes.Success,
		Message: facilityServiceRes.Message,
		Error:   storeErr,
	}

	return res, nil
}

func (g *grpcServer) EditFacility(ctx context.Context, req *facilitypb.EditFacilityRequest) (*facilitypb.GenericReply, error) {
	facilityRequest := io.AthFacilities{
		FacilityName:            req.Name,
		FacilityBasePrice:       req.BasePrice,
		FacilitySportCategories: req.SportCategories,
		FacilityID:              int(req.FacilityId),

		Models: io.Models{
			UpdatedBy: int(req.UpdatedBy),
		},
	}

	// get facility id from facility service
	facilityServiceRes := g.athFacilityService.EditFacility(ctx, facilityRequest)

	storeErr := ""
	res := &facilitypb.GenericReply{}

	if facilityServiceRes.Error != nil {
		storeErr = facilityServiceRes.Error.Error()
		res.Status = facilityServiceRes.Success
		res.Message = facilityServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	for weekkey, weekval := range req.WeekData.Weekdays {

		facilitySlotRequest := io.AthFacilitySlots{
			FacilityID: int(req.FacilityId),
			UserID:     int(req.UpdatedBy),
			SlotDays:   "weekdays",
			SlotType:   weekkey,
			SlotPrice:  weekval,

			Models: io.Models{
				UpdatedBy: int(req.UpdatedBy),
			},
		}
		// add facility slots
		FacilitySlotsServiceRes := g.athFacilityService.EditFacilitySlots(ctx, facilitySlotRequest)
		if FacilitySlotsServiceRes.Error != nil {
			storeErr = FacilitySlotsServiceRes.Error.Error()
			res.Status = FacilitySlotsServiceRes.Success
			res.Message = FacilitySlotsServiceRes.Message
			res.Error = storeErr
			return res, nil
		}
	}
	for weekkey, weekval := range req.WeekData.Weekends {

		facilitySlotRequest := io.AthFacilitySlots{
			FacilityID: int(req.FacilityId),
			UserID:     int(req.UpdatedBy),
			SlotDays:   "weekends",
			SlotType:   weekkey,
			SlotPrice:  weekval,

			Models: io.Models{
				UpdatedBy: int(req.UpdatedBy),
			},
		}
		// add facility slots
		FacilitySlotsServiceRes := g.athFacilityService.EditFacilitySlots(ctx, facilitySlotRequest)
		if FacilitySlotsServiceRes.Error != nil {
			storeErr = FacilitySlotsServiceRes.Error.Error()
			res.Status = FacilitySlotsServiceRes.Success
			res.Message = FacilitySlotsServiceRes.Message
			res.Error = storeErr
			return res, nil
		}
	}

	// if custom rates are provided

	if req.CustomRates != nil {

		for _, eachRate := range req.CustomRates {
			date, _ := time.Parse("2006-01-02", eachRate.Date)

			facilityCustomRateRequest := io.AthFacilityCustomRates{
				FacilityID:     int(eachRate.FacilityId),
				FacilitySlotID: int(eachRate.FacilitySlotId),
				UserID:         int(eachRate.UserId),
				Date:           date,
				SlotPrice:      eachRate.SlotPrice,
				IsActive:       eachRate.Available,
				Models: io.Models{
					CreatedBy: int(eachRate.UserId),
				},
			}

			// store facility custom rates
			facilityCustomRatesServiceRes := g.athFacilityService.AddFacilityCustomRates(ctx, facilityCustomRateRequest)

			if facilityCustomRatesServiceRes.Error != nil {
				storeErr = facilityCustomRatesServiceRes.Error.Error()
				res.Status = facilityCustomRatesServiceRes.Success
				res.Message = facilityCustomRatesServiceRes.Message
				res.Error = storeErr
				return res, nil
			}
		}
	}

	res = &facilitypb.GenericReply{
		Status:  facilityServiceRes.Success,
		Message: facilityServiceRes.Message,
		Error:   storeErr,
	}

	return res, nil
}

func (g *grpcServer) BookFacility(ctx context.Context, req *facilitypb.BookFacilityRequest) (*facilitypb.BookFacilityReply, error) {

	bookingNo := randstr.Hex(16)
	layout := "2006-01-02T15:04:05Z"
	bookingDate, _ := time.Parse(layout, req.BookingDate)
	bookingDateAfter := bookingDate.Add(-330 * time.Minute)

	facilityBookingsRequest := io.AthFacilityBookings{
		UserID:          int(req.CreatedBy),
		BookingNo:       bookingNo,
		BookingDate:     bookingDateAfter,
		BaseTotalAmount: req.BaseTotalAmount,
		DiscountAmount:  req.DiscountAmount,
		BookingAmount:   req.BookingAmount,
		BookingFee:      req.BookingFee,

		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	// get facility booking id from facility service
	facilityBookingServiceRes := g.athFacilityService.BookFacility(ctx, facilityBookingsRequest)

	storeErr := ""
	res := &facilitypb.BookFacilityReply{}

	if facilityBookingServiceRes.Error != nil {
		storeErr = facilityBookingServiceRes.Error.Error()
		res.Status = facilityBookingServiceRes.Success
		res.Message = facilityBookingServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	faciityBookingData := facilityBookingServiceRes.Data.(map[string]interface{})
	facilityBookingID := faciityBookingData["facility_booking_id"].(int)

	slotFromDate, _ := time.Parse(layout, req.SlotFromDate)
	slotFromDateAfter := slotFromDate.Add(-330 * time.Minute)
	slotToDate, _ := time.Parse(layout, req.SlotToDate)
	slotToDateAfter := slotToDate.Add(-330 * time.Minute)

	// save facility booking slots
	facilityBookingSlotRequest := io.AthFacilityBookingSlots{
		BookingID:        facilityBookingID,
		FaciliySlotID:    int(req.FacilitySlotId),
		SlotDays:         req.SlotDays,
		SlotFromDate:     slotFromDateAfter,
		SlotToDate:       slotToDateAfter,
		SlotBookingPrice: req.BookingAmount,
	}

	// add facility slots
	FacilityBookingSlotsServiceRes := g.athFacilityService.BookFacilitySlots(ctx, facilityBookingSlotRequest)
	if FacilityBookingSlotsServiceRes.Error != nil {
		storeErr = FacilityBookingSlotsServiceRes.Error.Error()
		res.Status = FacilityBookingSlotsServiceRes.Success
		res.Message = FacilityBookingSlotsServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	res = &facilitypb.BookFacilityReply{
		Data: &facilitypb.BookFacilityReplyData{
			BookingNo: faciityBookingData["facility_booking_no"].(string),
			BookingId: int32(facilityBookingID),
		},
		Status:  facilityBookingServiceRes.Success,
		Message: facilityBookingServiceRes.Message,
		Error:   storeErr,
	}

	return res, nil
}
