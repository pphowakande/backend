package handler

import (
	userpb "backend/api/athUser/v1"
	"backend/pkg/io"
	"backend/pkg/utile"
	"context"
	"strings"
)

func (g *grpcServer) SignupUser(ctx context.Context, req *userpb.SignupRequest) (*userpb.SignupReply, error) {
	userRequest := io.AthUser{
		Email:      req.Email,
		Password:   req.Password,
		UserSource: req.UserSource,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	companyRequest := io.AthCompany{
		Name:    req.CompanyName,
		Contact: req.ContactNo,
		Email:   req.Email,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	// get user id from user service
	userServiceRes := g.athUserService.SignUpUser(ctx, userRequest)

	storeErr := ""
	res := &userpb.SignupReply{}

	if userServiceRes.Error != nil {
		storeErr = userServiceRes.Error.Error()
		res.Status = userServiceRes.Success
		res.Message = userServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	// get company id from company service
	companyServiceRes := g.athCompanyService.Createcompany(ctx, companyRequest)
	if companyServiceRes.Error != nil {
		storeErr = companyServiceRes.Error.Error()
		res.Status = companyServiceRes.Success
		res.Message = companyServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	companyData := companyServiceRes.Data.(map[string]interface{})
	companyID := companyData["company_id"].(int)

	userData := userServiceRes.Data.(map[string]interface{})
	userID := userData["user_id"].(int)

	// add details to company_user table
	companyUserReq := io.AthCompanyUser{
		CompanyID: companyID,
		UserID:    userID,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	companyUserServiceRes := g.athCompanyService.CreatecompanyUser(ctx, companyUserReq)
	if companyUserServiceRes.Error != nil {
		storeErr = companyUserServiceRes.Error.Error()
		res.Status = companyUserServiceRes.Success
		res.Message = companyUserServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	// add user profile details
	FirstName := ""
	LastName := ""

	if strings.Contains(req.FullName, " ") {
		splittedName := strings.Split(req.FullName, " ")
		FirstName = splittedName[0]
		LastName = splittedName[1]
	} else {
		FirstName = req.FullName
		LastName = ""
	}

	profileRequest := io.AthUserProfile{
		FirstName: FirstName,
		LastName:  LastName,
		UserID:    userID,
		ContactNo: req.ContactNo,
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
	otpRequest.Contact = req.ContactNo

	otpServiceRes := g.athOtpService.CreateOTP(ctx, otpRequest)
	if otpServiceRes.Error != nil {
		storeErr = otpServiceRes.Error.Error()
		res.Status = otpServiceRes.Success
		res.Message = otpServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	verifyPhoneToken := ""
	verifyPhoneTokenData := otpServiceRes.Data.(map[string]interface{})
	verifyPhoneToken = verifyPhoneTokenData["verify_token"].(string)

	res = &userpb.SignupReply{
		Data: &userpb.SignUpReplyData{
			VerifyPhoneToken: verifyPhoneToken,
			CompanyId:        int32(companyID),
			UserId:           int32(userID),
		},
		Status:  otpServiceRes.Success,
		Message: otpServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) PhoneVerifyUser(ctx context.Context, req *userpb.PhoneVerifyRequest) (*userpb.GenericReply, error) {

	request := io.Verify{
		Code:  req.Code,
		Type:  req.Type,
		Email: req.Email,
	}
	serviceRes := g.athOtpService.VerifyOTP(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}
	res := &userpb.GenericReply{
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) ResendCode(ctx context.Context, req *userpb.ResendCodeRequest) (*userpb.SignupReply, error) {
	storeErr := ""
	res := &userpb.SignupReply{}

	// trigger create token service
	var otpRequest io.AthUserOTP

	otpRequest.UserID = int(req.UserId)
	otpRequest.OTPNO = utile.RandomString(6)
	otpRequest.OTPExpiry = 3600
	otpRequest.OTPType = "phone"
	otpRequest.Contact = req.ContactNo

	otpServiceRes := g.athOtpService.CreateOTP(ctx, otpRequest)
	if otpServiceRes.Error != nil {
		storeErr = otpServiceRes.Error.Error()
		res.Status = otpServiceRes.Success
		res.Message = otpServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	verifyPhoneToken := ""
	verifyPhoneTokenData := otpServiceRes.Data.(map[string]interface{})
	verifyPhoneToken = verifyPhoneTokenData["verify_token"].(string)

	res = &userpb.SignupReply{
		Data: &userpb.SignUpReplyData{
			VerifyPhoneToken: verifyPhoneToken,
		},
		Status:  otpServiceRes.Success,
		Message: otpServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) EmailVerifyUser(ctx context.Context, req *userpb.EmailVerifyRequest) (*userpb.GenericReply, error) {

	request := io.EmailVerify{
		Email: req.Email,
	}
	serviceRes := g.athUserService.EmailVerify(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}
	res := &userpb.GenericReply{
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) LoginUser(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginReply, error) {
	request := io.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	serviceRes := g.athUserService.LoginUser(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}
	uid := ""
	if serviceRes.Error == nil {
		userId := serviceRes.Data.(map[string]string)
		uid = userId["user_id"]
	}
	res := &userpb.LoginReply{
		Data: &userpb.LoginReplyData{
			UserId: uid,
		},
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) ForgotPasswordUser(ctx context.Context, req *userpb.ForgotPasswordRequest) (*userpb.ForgotPasswordReply, error) {

	request := io.ForgotPasswordRequest{
		Email: req.Email,
	}
	res := &userpb.ForgotPasswordReply{}
	serviceRes := g.athUserService.ForgotPasswordUser(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}
	res = &userpb.ForgotPasswordReply{
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	if serviceRes.Error != nil {
		return res, nil
	}

	ForgotPasswordToken := ""
	ForgotPasswordTokenData := serviceRes.Data.(map[string]interface{})
	ForgotPasswordToken = ForgotPasswordTokenData["reset_password_token"].(string)

	res = &userpb.ForgotPasswordReply{
		Data: &userpb.ForgotReplyData{
			ResetPasswordToken: ForgotPasswordToken,
		},
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) ResetPasswordUser(ctx context.Context, req *userpb.ResetPasswordRequest) (*userpb.GenericReply, error) {
	request := io.ResetPasswordRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	serviceRes := g.athUserService.ResetPasswordUser(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}
	res := &userpb.GenericReply{
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) EditUser(ctx context.Context, req *userpb.EditUserRequest) (*userpb.GenericReply, error) {

	storeErr := ""
	res := &userpb.GenericReply{}

	// edit user data
	if req.Email != "" {

		userRequest := io.AthUser{
			ID:    int(req.UserId),
			Email: req.Email,
			Models: io.Models{
				UpdatedBy: int(req.UpdatedBy),
			},
		}

		// update user email
		userServiceRes := g.athUserService.EditUser(ctx, userRequest)
		if userServiceRes.Error != nil {
			storeErr = userServiceRes.Error.Error()
			res.Status = userServiceRes.Success
			res.Message = userServiceRes.Message
			res.Error = storeErr
			return res, nil
		}
	}

	// edit company data
	if (req.CompanyName != "") || (req.ContactNo != "") || (req.Email != "") {
		companyRequest := io.AthCompany{
			ID:      int(req.CompanyId),
			Name:    req.CompanyName,
			Contact: req.ContactNo,
			Email:   req.Email,
			Models: io.Models{
				UpdatedBy: int(req.UpdatedBy),
			},
		}

		// update company data
		companyServiceRes := g.athCompanyService.EditCompany(ctx, companyRequest)
		if companyServiceRes.Error != nil {
			storeErr = companyServiceRes.Error.Error()
			res.Status = companyServiceRes.Success
			res.Message = companyServiceRes.Message
			res.Error = storeErr
			return res, nil
		}
	}

	// edit user profile data
	if req.FullName != "" {
		// edit  user profile details
		FirstName := ""
		LastName := ""

		if strings.Contains(req.FullName, " ") {
			splittedName := strings.Split(req.FullName, " ")
			FirstName = splittedName[0]
			LastName = splittedName[1]
		} else {
			FirstName = req.FullName
			LastName = ""
		}

		profileRequest := io.AthUserProfile{
			UserProfileID: int(req.ProfileId),
			FirstName:     FirstName,
			LastName:      LastName,
			UserID:        int(req.UserId),
			ContactNo:     req.ContactNo,
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
	}

	// check if phone has been updated. If yes, do phone verification

	if req.ContactNo != "" {
		// trigger create token service
		var otpRequest io.AthUserOTP

		otpRequest.UserID = int(req.UserId)
		otpRequest.OTPNO = utile.RandomString(6)
		otpRequest.OTPExpiry = 3600
		otpRequest.OTPType = "phone"
		otpRequest.Contact = req.ContactNo

		otpServiceRes := g.athOtpService.CreateOTP(ctx, otpRequest)
		if otpServiceRes.Error != nil {
			storeErr = otpServiceRes.Error.Error()
			res.Status = otpServiceRes.Success
			res.Message = otpServiceRes.Message
			res.Error = storeErr
			return res, nil
		}

		res = &userpb.GenericReply{
			Status:  otpServiceRes.Success,
			Message: otpServiceRes.Message,
			Error:   storeErr,
		}
	}

	res = &userpb.GenericReply{
		Status:  true,
		Message: "Success",
		Error:   storeErr,
	}

	return res, nil
}
