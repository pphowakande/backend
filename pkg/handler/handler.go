package handler

import (
	adminpb "backend/api/athAdmin/v1"
	customerpb "backend/api/athCustomer/v1"
	facilitypb "backend/api/athFacility/v1"
	userpb "backend/api/athUser/v1"
	venuepb "backend/api/athVenue/v1"

	adminService "backend/service/athAdmin"
	companyService "backend/service/athCompany"
	userService "backend/service/athUser"

	otpService "backend/service/athOtp"

	customerService "backend/service/athCustomer"
	facilityService "backend/service/athFacility"
	venueService "backend/service/athVenue"

	"context"
)

type GrpcService interface {
	SignupUser(ctx context.Context, in *userpb.SignupRequest) (*userpb.SignupReply, error)
	PhoneVerifyUser(ctx context.Context, in *userpb.PhoneVerifyRequest) (*userpb.GenericReply, error)
	ResendCode(ctx context.Context, in *userpb.ResendCodeRequest) (*userpb.SignupReply, error)
	EmailVerifyUser(ctx context.Context, in *userpb.EmailVerifyRequest) (*userpb.GenericReply, error)
	LoginUser(ctx context.Context, in *userpb.LoginRequest) (*userpb.LoginReply, error)
	ForgotPasswordUser(ctx context.Context, in *userpb.ForgotPasswordRequest) (*userpb.ForgotPasswordReply, error)
	ResetPasswordUser(ctx context.Context, in *userpb.ResetPasswordRequest) (*userpb.GenericReply, error)
	EditUser(ctx context.Context, in *userpb.EditUserRequest) (*userpb.GenericReply, error)

	LoginAdmin(ctx context.Context, in *adminpb.LoginAdminRequest) (*adminpb.LoginAdminReply, error)
	SignupAdmin(ctx context.Context, in *adminpb.SignupAdminRequest) (*adminpb.GenericReply, error)

	CreateCompany(ctx context.Context, in *userpb.CreateCompanyRequest) (*userpb.CreateCompanyReply, error)
	CreateCompanyUser(ctx context.Context, in *userpb.CreateCompanyUserRequest) (*userpb.GenericReply, error)

	CreateOTP(ctx context.Context, in *userpb.OtpRequest) (*userpb.GenericReply, error)

	CreateVenue(ctx context.Context, in *venuepb.CreateVenueRequest) (*venuepb.CreateVenueReply, error)
	EditVenue(ctx context.Context, in *venuepb.EditVenueRequest) (*venuepb.GenericReply, error)
	CreateVenueHoliday(ctx context.Context, in *venuepb.CreateVenueHolidayRequest) (*venuepb.CreateVenueHolidayReply, error)
	DeleteVenueHoliday(ctx context.Context, in *venuepb.DeleteVenueHolidayRequest) (*venuepb.GenericReply, error)
	CreateFacility(ctx context.Context, in *facilitypb.CreateFacilityRequest) (*facilitypb.CreateFacilityReply, error)
	EditFacility(ctx context.Context, in *facilitypb.EditFacilityRequest) (*facilitypb.GenericReply, error)
	BookFacility(ctx context.Context, in *facilitypb.BookFacilityRequest) (*facilitypb.BookFacilityReply, error)

	CreateCustomer(ctx context.Context, in *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerReply, error)
	EditCustomer(ctx context.Context, in *customerpb.EditCustomerRequest) (*customerpb.GenericReply, error)
}

type grpcServer struct {
	athUserService     userService.AthUserService
	athCompanyService  companyService.AthCompanyService
	athAdminUserSevice adminService.AthAdminUserService
	athOtpService      otpService.AthOtpService
	athVenueService    venueService.AthVenueService
	athFacilityService facilityService.AthFacilityService
	athCustomerService customerService.AthCustomerService
}

func NewGrpcService(athCompanyService companyService.AthCompanyService, athUserService userService.AthUserService, athAdminUserService adminService.AthAdminUserService, athOtpService otpService.AthOtpService, athVenueService venueService.AthVenueService, athFacilityService facilityService.AthFacilityService, athCustomerService customerService.AthCustomerService) GrpcService {
	return &grpcServer{
		athUserService:     athUserService,
		athAdminUserSevice: athAdminUserService,
		athCompanyService:  athCompanyService,
		athOtpService:      athOtpService,
		athVenueService:    athVenueService,
		athFacilityService: athFacilityService,
		athCustomerService: athCustomerService,
	}
}
