package handler

import (
	userpb "backend/api/athUser/v1"
	"backend/pkg/io"
	athCompany "backend/service/athCompany/mock"
	athOtp "backend/service/athOtp/mock"
	athUser "backend/service/athUser/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGrpcServer_SignupUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)
	mockAthCompanyService := athCompany.NewMockAthCompanyService(ctrl)
	mockAthOtpService := athOtp.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athUserService:     mockAthUserService,
		athCompanyService:  mockAthCompanyService,
		athAdminUserSevice: nil,
		athOtpService:      mockAthOtpService,
	}

	type args struct {
		ctx context.Context
		req *userpb.SignupRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.SignupReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.SignupRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					SignUpUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"user_id": 2}})
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"company_id": 1}})
				mockAthCompanyService.
					EXPECT().
					CreatecompanyUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"company_id": 1}})
				mockAthUserService.
					EXPECT().
					CreateUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOtpService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Data: map[string]interface{}{"verify_token": "test_token"}, Success: true})
			},
			want: &userpb.SignupReply{
				Data: &userpb.SignUpReplyData{
					VerifyPhoneToken: "test_token",
					CompanyId:        int32(1),
					UserId:           int32(1),
				},
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "SignUpUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.SignupRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					SignUpUser(ctx, gomock.Any()).
					Return(io.Response{
						Error:   errors.New("failed to signUpUser"),
						Success: false,
					})
			},
			want: &userpb.SignupReply{
				Status: false,
				Error:  "failed to signUpUser",
			},
			wantErr: false,
		},
		{
			name: "Createcompany service error",
			args: args{
				ctx: ctx,
				req: &userpb.SignupRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					SignUpUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"user_id": 1}})
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{
						Error:   errors.New("failed to createcompany"),
						Success: false,
					})
			},
			want: &userpb.SignupReply{
				Status: false,
				Error:  "failed to createcompany",
			},
			wantErr: false,
		},
		{
			name: "CreatecompanyUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.SignupRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					SignUpUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"user_id": 1}})
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"company_id": 1}})
				mockAthCompanyService.
					EXPECT().
					CreatecompanyUser(ctx, gomock.Any()).
					Return(io.Response{
						Error:   errors.New("failed to createcompanyuser"),
						Success: false,
					})
			},
			want: &userpb.SignupReply{
				Status: false,
				Error:  "failed to createcompanyuser",
			},
			wantErr: false,
		},
		{
			name: "CreateUserProfile service error",
			args: args{
				ctx: ctx,
				req: &userpb.SignupRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					SignUpUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"user_id": 1}})
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"company_id": 1}})
				mockAthCompanyService.
					EXPECT().
					CreatecompanyUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthUserService.
					EXPECT().
					CreateUserProfile(ctx, gomock.Any()).
					Return(io.Response{
						Error:   errors.New("failed to createuserprofile"),
						Success: false,
					})
			},
			want: &userpb.SignupReply{
				Status: false,
				Error:  "failed to createuserprofile",
			},
			wantErr: false,
		},
		{
			name: "CreateOTP service error",
			args: args{
				ctx: ctx,
				req: &userpb.SignupRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					SignUpUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"user_id": 1}})
				mockAthCompanyService.
					EXPECT().
					Createcompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{"company_id": 1}})
				mockAthCompanyService.
					EXPECT().
					CreatecompanyUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthUserService.
					EXPECT().
					CreateUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOtpService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{
						Error:   errors.New("failed to createOTP"),
						Success: false,
					})
			},
			want: &userpb.SignupReply{
				Status: false,
				Error:  "failed to createOTP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.SignupUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_PhoneVerifyUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthOtpService := athOtp.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athOtpService: mockAthOtpService,
	}

	type args struct {
		ctx context.Context
		req *userpb.PhoneVerifyRequest
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
				req: &userpb.PhoneVerifyRequest{},
			},
			mock: func() {
				mockAthOtpService.
					EXPECT().
					VerifyOTP(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "VerifyOTP service error",
			args: args{
				ctx: ctx,
				req: &userpb.PhoneVerifyRequest{},
			},
			mock: func() {
				mockAthOtpService.
					EXPECT().
					VerifyOTP(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to verifyOTP"), Success: false})
			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to verifyOTP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.PhoneVerifyUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_LoginUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)

	services := grpcServer{
		athUserService: mockAthUserService,
	}

	type args struct {
		ctx context.Context
		req *userpb.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.LoginReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.LoginRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					LoginUser(ctx, gomock.Any()).
					Return(io.Response{Data: map[string]string{"user_id": "test_uid"}, Success: true})
			},
			want: &userpb.LoginReply{
				Data: &userpb.LoginReplyData{
					UserId: "test_uid",
				},
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "LoginUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.LoginRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					LoginUser(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to loginUser"), Success: false})
			},
			want: &userpb.LoginReply{
				Status: false,
				Error:  "failed to loginUser",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.LoginUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_ForgotPasswordUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)

	services := grpcServer{
		athUserService: mockAthUserService,
	}

	type args struct {
		ctx context.Context
		req *userpb.ForgotPasswordRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.ForgotPasswordReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.ForgotPasswordRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					ForgotPasswordUser(ctx, gomock.Any()).
					Return(io.Response{Data: map[string]interface{}{"reset_password_token": "test_password_token"}, Success: true})
			},
			want: &userpb.ForgotPasswordReply{
				Data: &userpb.ForgotReplyData{
					ResetPasswordToken: "test_password_token",
				},
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "ForgotPasswordUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.ForgotPasswordRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					ForgotPasswordUser(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to forgotPasswordUser"), Success: false})
			},
			want: &userpb.ForgotPasswordReply{
				Status: false,
				Error:  "failed to forgotPasswordUser",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.ForgotPasswordUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_ResetPasswordUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)

	services := grpcServer{
		athUserService: mockAthUserService,
	}

	type args struct {
		ctx context.Context
		req *userpb.ResetPasswordRequest
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
				req: &userpb.ResetPasswordRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					ResetPasswordUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "ResetPasswordUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.ResetPasswordRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					ResetPasswordUser(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to resetPasswordUser"), Success: false})
			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to resetPasswordUser",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.ResetPasswordUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_ResendCode(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthOtpService := athOtp.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athOtpService: mockAthOtpService,
	}

	type args struct {
		ctx context.Context
		req *userpb.ResendCodeRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    *userpb.SignupReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: ctx,
				req: &userpb.ResendCodeRequest{},
			},
			mock: func() {
				mockAthOtpService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{
						"verify_token": "test_token",
					}})
			},
			want: &userpb.SignupReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "ResendCode service error",
			args: args{
				ctx: ctx,
				req: &userpb.ResendCodeRequest{},
			},
			mock: func() {
				mockAthOtpService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to resendcode"), Success: false})
			},
			want: &userpb.SignupReply{
				Status: false,
				Error:  "failed to resendcode",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.ResendCode(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_EmailVerifyUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)

	services := grpcServer{
		athUserService: mockAthUserService,
	}

	type args struct {
		ctx context.Context
		req *userpb.EmailVerifyRequest
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
				req: &userpb.EmailVerifyRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EmailVerify(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "ResendCode service error",
			args: args{
				ctx: ctx,
				req: &userpb.EmailVerifyRequest{},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EmailVerify(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to resendcode"), Success: false})
			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to resendcode",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.EmailVerifyUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}

func TestGrpcServer_EditUser(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAthUserService := athUser.NewMockAthUserService(ctrl)
	mockAthCompanyService := athCompany.NewMockAthCompanyService(ctrl)
	mockAthOtpService := athOtp.NewMockAthOtpService(ctrl)

	services := grpcServer{
		athUserService:    mockAthUserService,
		athCompanyService: mockAthCompanyService,
		athOtpService:     mockAthOtpService,
	}

	type args struct {
		ctx context.Context
		req *userpb.EditUserRequest
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
				req: &userpb.EditUserRequest{
					Email:     "test@test.com",
					FullName:  "Christiano Ronaldo",
					ContactNo: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthCompanyService.
					EXPECT().
					EditCompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOtpService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
			},
			want: &userpb.GenericReply{
				Status: true,
			},
			wantErr: false,
		},
		{
			name: "EditUser service error",
			args: args{
				ctx: ctx,
				req: &userpb.EditUserRequest{
					Email:     "test@test.com",
					FullName:  "Christiano Ronaldo",
					ContactNo: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUser(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to edituser"), Success: false})

			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to edituser",
			},
			wantErr: false,
		},
		{
			name: "EditCompany service error",
			args: args{
				ctx: ctx,
				req: &userpb.EditUserRequest{
					Email:     "test@test.com",
					FullName:  "Christiano Ronaldo",
					ContactNo: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthCompanyService.
					EXPECT().
					EditCompany(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to editcompany"), Success: false})

			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to editcompany",
			},
			wantErr: false,
		},
		{
			name: "EditUserProfile service error",
			args: args{
				ctx: ctx,
				req: &userpb.EditUserRequest{
					Email:     "test@test.com",
					FullName:  "Christiano Ronaldo",
					ContactNo: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthCompanyService.
					EXPECT().
					EditCompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to edituserprofile"), Success: false})

			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to edituserprofile",
			},
			wantErr: false,
		},
		{
			name: "CreateOTP service error",
			args: args{
				ctx: ctx,
				req: &userpb.EditUserRequest{
					Email:     "test@test.com",
					FullName:  "Christiano Ronaldo",
					ContactNo: "0123456789",
				},
			},
			mock: func() {
				mockAthUserService.
					EXPECT().
					EditUser(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthCompanyService.
					EXPECT().
					EditCompany(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthUserService.
					EXPECT().
					EditUserProfile(ctx, gomock.Any()).
					Return(io.Response{Success: true, Data: map[string]interface{}{}})
				mockAthOtpService.
					EXPECT().
					CreateOTP(ctx, gomock.Any()).
					Return(io.Response{Error: errors.New("failed to createOTP"), Success: false})

			},
			want: &userpb.GenericReply{
				Status: false,
				Error:  "failed to createOTP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := services.EditUser(tt.args.ctx, tt.args.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.want, got)
			}

			assert.NoError(t, err)
		})
	}
}
