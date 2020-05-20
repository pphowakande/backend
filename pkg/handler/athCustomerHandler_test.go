package handler

import (
	customerpb "backend/api/athCustomer/v1"
	"backend/pkg/io"
	athCustomer "backend/service/athCustomer/mock"
	athOTP "backend/service/athOtp/mock"
	athUser "backend/service/athUser/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_CreateCustomer(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthCustomerService := athCustomer.NewMockAthCustomerService(ctrl)
	mockAthUserService := athUser.NewMockAthUserService(ctrl)
	mockAthOTPService := athOTP.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athCustomerService: mockAthCustomerService,
		athUserService:     mockAthUserService,
		athOtpService:      mockAthOTPService,
	}

	type args struct {
		ctx context.Context
		req *customerpb.CreateCustomerRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *customerpb.CreateCustomerReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &customerpb.CreateCustomerRequest{},
			},
			mock: func() {
				mockAthCustomerService.
					EXPECT().
					CreateCustomer(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"user_id": 1,
					}})
				mockAthUserService.
					EXPECT().
					CreateUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOTPService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Data: map[string]interface{}{"verify_token": "test_token"}, Success: true})
			},
			want: &customerpb.CreateCustomerReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "CreateCustomer service error",
			args: args{
				ctx: ctx,
				req: &customerpb.CreateCustomerRequest{},
			},
			mock: func() {
				mockAthCustomerService.
					EXPECT().
					CreateCustomer(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createcustomer"), Success: false})
			},
			want: &customerpb.CreateCustomerReply{
				Status: false,
				Error:  "failed to createcustomer",
			},
			wantErr: false,
		},
		{
			name: "CreateUserProfile service error",
			args: args{
				ctx: ctx,
				req: &customerpb.CreateCustomerRequest{},
			},
			mock: func() {
				mockAthCustomerService.
					EXPECT().
					CreateCustomer(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"user_id": 1,
					}})
				mockAthUserService.
					EXPECT().
					CreateUserProfile(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createuserprofile"), Success: false})
			},
			want: &customerpb.CreateCustomerReply{
				Status: false,
				Error:  "failed to createuserprofile",
			},
			wantErr: false,
		},
		{
			name: "CreateOTP service error",
			args: args{
				ctx: ctx,
				req: &customerpb.CreateCustomerRequest{},
			},
			mock: func() {
				mockAthCustomerService.
					EXPECT().
					CreateCustomer(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"user_id": 1,
					}})
				mockAthUserService.
					EXPECT().
					CreateUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOTPService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createOTP"), Success: false})
			},
			want: &customerpb.CreateCustomerReply{
				Status: false,
				Error:  "failed to createOTP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateCustomer(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_EditCustomer(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)
	mockAthOTPService := athOTP.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athUserService: mockAthUserService,
		athOtpService:  mockAthOTPService,
	}

	type args struct {
		ctx context.Context
		req *customerpb.EditCustomerRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *customerpb.GenericReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &customerpb.EditCustomerRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &customerpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "OK With Phone Req",
			args: args{
				ctx: ctx,
				req: &customerpb.EditCustomerRequest{
					Phone: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOTPService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"verify_token": "test_token",
					}})
			},
			want: &customerpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "EditCustomer service error",
			args: args{
				ctx: ctx,
				req: &customerpb.EditCustomerRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to editcustomer"), Success: false})
			},
			want: &customerpb.GenericReply{
				Status: false,
				Error:  "failed to editcustomer",
			},
			wantErr: false,
		},
		{
			name: "CreateOTP service error",
			args: args{
				ctx: ctx,
				req: &customerpb.EditCustomerRequest{
					Phone: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOTPService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createOTP"), Success: false})
			},
			want: &customerpb.GenericReply{
				Status: false,
				Error:  "failed to createOTP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := services.EditCustomer(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
