package handler

import (
	userpb "backend/api/athUser/v1"
	"backend/pkg/io"
	"context"
)

func (g *grpcServer) CreateOTP(ctx context.Context, req *userpb.OtpRequest) (*userpb.GenericReply, error) {
	otpRequest := io.AthUserOTP{
		OTPExpiry: 3600, // expired after 1 hour
		OTPNO:     req.OtpNo,
		OTPType:   req.OtpType,
		UserID:    int(req.UserId),
	}

	// trigger otp service
	otpServiceRes := g.athOtpService.CreateOTP(ctx, otpRequest)

	storeErr := ""
	if otpServiceRes.Error != nil {
		storeErr = otpServiceRes.Error.Error()
	}

	res := &userpb.GenericReply{
		Status:  otpServiceRes.Success,
		Message: otpServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}
