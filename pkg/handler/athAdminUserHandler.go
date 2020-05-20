package handler

import (
	adminpb "backend/api/athAdmin/v1"
	"backend/pkg/io"
	"context"
)

func (g *grpcServer) LoginAdmin(ctx context.Context, req *adminpb.LoginAdminRequest) (*adminpb.LoginAdminReply, error) {
	request := io.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	serviceRes := g.athAdminUserSevice.LoginAdmin(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}
	uid := ""
	if serviceRes.Error == nil {
		userID := serviceRes.Data.(map[string]string)
		uid = userID["user_id"]
	}
	res := &adminpb.LoginAdminReply{
		Data: &adminpb.LoginAdminReplyData{
			AdminId: uid,
		},
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) SignupAdmin(ctx context.Context, req *adminpb.SignupAdminRequest) (*adminpb.GenericReply, error) {
	userRequest := io.AthAdminUser{
		Email:    req.Email,
		Password: req.Password,
		UserName: req.UserName,
		Name:     req.Name,
		Models: io.Models{
			CreatedBy: int(req.CreatedBy),
		},
	}

	// get user id from user service
	userServiceRes := g.athAdminUserSevice.SignUpAdmin(ctx, userRequest)
	storeErr := ""
	res := &adminpb.GenericReply{}

	if userServiceRes.Error != nil {
		storeErr = userServiceRes.Error.Error()
		res.Status = userServiceRes.Success
		res.Message = userServiceRes.Message
		res.Error = storeErr
		return res, nil
	}

	res = &adminpb.GenericReply{
		Status:  userServiceRes.Success,
		Message: userServiceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}
