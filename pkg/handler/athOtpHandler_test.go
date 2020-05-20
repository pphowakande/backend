package handler

import (
	userpb "backend/api/athUser/v1"
	"backend/pkg/io"
	athOTP "backend/service/athOtp/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_CreateOTP(t *testing.T) {

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthOTPService := athOTP.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athOtpService: mockAthOTPService,
	}

	type args struct {
		ctx context.Context
		req *userpb.OtpRequest
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
				req: &userpb.OtpRequest{},
			},
			mock: func() {
				mockAthOTPService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"verify_token": "test_token",
					}})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "CreateOTP service error",
			args: args{
				ctx: ctx,
				req: &userpb.OtpRequest{},
			},
			mock: func() {
				mockAthOTPService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createotp"), Success: false})
			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to createotp",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.CreateOTP(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
