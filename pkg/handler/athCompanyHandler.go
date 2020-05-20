package handler

import (
	userpb "backend/api/athUser/v1"
	"backend/pkg/io"
	"context"
)

func (g *grpcServer) CreateCompany(ctx context.Context, req *userpb.CreateCompanyRequest) (*userpb.CreateCompanyReply, error) {

	request := io.AthCompany{}
	serviceRes := g.athCompanyService.Createcompany(ctx, request)
	storeErr := ""
	if serviceRes.Error != nil {
		storeErr = serviceRes.Error.Error()
	}

	res := &userpb.CreateCompanyReply{
		Data:    &userpb.CreateCompanyReplyData{},
		Status:  serviceRes.Success,
		Message: serviceRes.Message,
		Error:   storeErr,
	}
	return res, nil
}

func (g *grpcServer) CreateCompanyUser(ctx context.Context, req *userpb.CreateCompanyUserRequest) (*userpb.GenericReply, error) {

	request := io.AthCompanyUser{}
	serviceRes := g.athCompanyService.CreatecompanyUser(ctx, request)
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
