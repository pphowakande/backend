package handler

import (
	userpb "backend/api/athAdmin/v1"
	"backend/pkg/io"
	athAdmin "backend/service/athAdmin/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_SignupAdmin(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthAdminUserService := athAdmin.NewMockAthAdminUserService(ctrl)

	services := grpcServer{
		athAdminUserSevice: mockAthAdminUserService,
	}

	type args struct {
		ctx context.Context
		req *userpb.SignupAdminRequest
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
				req: &userpb.SignupAdminRequest{},
			},
			mock: func() {
				mockAthAdminUserService.
					EXPECT().
					SignUpAdmin(ctx, gomock.Any()).
					Return(io.Response{Success: true})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "SignUpAdmin service error",
			args: args{
				ctx: ctx,
				req: &userpb.SignupAdminRequest{},
			},
			mock: func() {
				mockAthAdminUserService.
					EXPECT().
					SignUpAdmin(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to signUpAdmin"), Success: false})
			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to signUpAdmin",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.SignupAdmin(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_LoginAdmin(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthAdminUserService := athAdmin.NewMockAthAdminUserService(ctrl)

	services := grpcServer{
		athAdminUserSevice: mockAthAdminUserService,
	}

	type args struct {
		ctx context.Context
		req *userpb.LoginAdminRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.LoginAdminReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.LoginAdminRequest{},
			},
			mock: func() {
				mockAthAdminUserService.
					EXPECT().
					LoginAdmin(ctx, gomock.Any()).
					Return(io.Response{Data: map[string]string{"user_id": "test_uid"}, Success: true})
			},
			want: &userpb.LoginAdminReply{
				Data: &userpb.LoginAdminReplyData{
					AdminId: "test_uid",
				},
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "LoginAdmin service error",
			args: args{
				ctx: ctx,
				req: &userpb.LoginAdminRequest{},
			},
			mock: func() {
				mockAthAdminUserService.
					EXPECT().
					LoginAdmin(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to loginAdmin"), Success: false})
			},
			want: &userpb.LoginAdminReply{
				Status: false,
				Error:  "failed to loginAdmin",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.LoginAdmin(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
