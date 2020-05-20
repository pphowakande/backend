package handler

import (
	userpb "backend/api/athUser/v1"
	"backend/pkg/io"
	athCompany "backend/service/athCompany/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_CreateCompany(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthCompanyService := athCompany.NewMockAthCompanyService(ctrl)

	services := grpcServer{
		athCompanyService: mockAthCompanyService,
	}

	type args struct {
		ctx context.Context
		req *userpb.CreateCompanyRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.CreateCompanyReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.CreateCompanyRequest{},
			},
			mock: func() {
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{Success: true})
			},
			want: &userpb.CreateCompanyReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "Createcompany service error",
			args: args{
				ctx: ctx,
				req: &userpb.CreateCompanyRequest{},
			},
			mock: func() {
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createcompany"), Success: false})
			},
			want: &userpb.CreateCompanyReply{
				Status: false,
				Error:  "failed to createcompany",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateCompany(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_CreateCompanyUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthCompanyService := athCompany.NewMockAthCompanyService(ctrl)
	services := grpcServer{
		athCompanyService: mockAthCompanyService,
	}

	type args struct {
		ctx context.Context
		req *userpb.CreateCompanyUserRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.GenericReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.CreateCompanyUserRequest{},
			},
			mock: func() {
				mockAthCompanyService.
					EXPECT().
					CreatecompanyUser(ctx, gomock.Any()).
					Return(io.Response{Success: true})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "CreatecompanyUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.CreateCompanyUserRequest{},
			},
			mock: func() {
				mockAthCompanyService.
					EXPECT().
					CreatecompanyUser(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createcompanyUser"), Success: false})
			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to createcompanyUser",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateCompanyUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
