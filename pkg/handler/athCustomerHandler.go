package handler

import (
	customerpb "backend/api/athCustomer/v1"
	"backend/pkg/io"
	"backend/pkg/utile"
	"context"
)

func (g *grpcServer) CreateCustomer(ctx context.Context, req *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerReply, error) {
	userRequest := io.AthUser{
		Email:      req.Email,
		UserSource: req.UserSource,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	storeErr := ""
	res := &customerpb.CreateCustomerReply{}

	createCustomerServiceRes := g.athCustomerService.CreateCustomer(ctx, userRequest)

	if createCustomerServiceRes.Error != nil {
		storeErr = createCustomerServiceRes.Error.Error()
		res.Status = createCustomerServiceRes.Success
		res.Message = createCustomerServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	customerData := createCustomerServiceRes.Data.(map[string]interface{})
	userID := customerData["user_id"].(int)

	profileRequest := io.AthUserProfile{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserID:    userID,
		ContactNo: req.Phone,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	profileServiceRes := g.athUserService.CreateUserProfile(ctx, profileRequest)
	if profileServiceRes.Error != nil {
		storeErr = profileServiceRes.Error.Error()
		res.Status = profileServiceRes.Success
		res.Message = profileServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	// trigger create token service
	var otpRequest io.AthUserOTP

	otpRequest.UserID = userID
	otpRequest.OTPNO = utile.RandomString(6)
	otpRequest.OTPExpiry = 3600
	otpRequest.OTPType = "phone"
	otpRequest.Contact = req.Phone

	otpServiceRes := g.athOtpService.CreateOTP(ctx, otpRequest)
	if otpServiceRes.Error != nil {
		storeErr = otpServiceRes.Error.Error()
		res.Status = otpServiceRes.Success
		res.Message = otpServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	res = &customerpb.CreateCustomerReply{
		Data: &customerpb.CreateCustomerReplyData{
			CustomerId: int32(userID),
		},
		Status:  createCustomerServiceRes.Success,
		Message: createCustomerServiceRes.Message,
		Error:   storeErr,
	}

	return res, nil
}

func (g *grpcServer) EditCustomer(ctx context.Context, req *customerpb.EditCustomerRequest) (*customerpb.GenericReply, error) {
	storeErr := ""
	res := &customerpb.GenericReply{}

	profileRequest := io.AthUserProfile{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		ContactNo:     req.Phone,
		UserProfileID: int(req.UserId),
		Models: io.Models{
			UpdatedBy: int(req.UpdatedBy),
		},
	}

	profileServiceRes := g.athUserService.EditUserProfile(ctx, profileRequest)
	if profileServiceRes.Error != nil {
		storeErr = profileServiceRes.Error.Error()
		res.Status = profileServiceRes.Success
		res.Message = profileServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	if req.Phone != "" {
		// trigger create token service
		var otpRequest io.AthUserOTP

		otpRequest.UserID = int(req.UserId)
		otpRequest.OTPNO = utile.RandomString(6)
		otpRequest.OTPExpiry = 3600
		otpRequest.OTPType = "phone"
		otpRequest.Contact = req.Phone

		otpServiceRes := g.athOtpService.CreateOTP(ctx, otpRequest)
		if otpServiceRes.Error != nil {
			storeErr = otpServiceRes.Error.Error()
			res.Status = otpServiceRes.Success
			res.Message = otpServiceRes.Message
			res.Error = storeErr
			return res, nil
		}
	}

	res = &customerpb.GenericReply{
		Status:  profileServiceRes.Success,
		Message: profileServiceRes.Message,
		Error:   storeErr,
	}

	return res, nil
}
