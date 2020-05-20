package db

import (
	config "backend/cmd/config"
	"backend/pkg/io"
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
	//logger log.Logger
}

func Connect() *gorm.DB {

	configuration := config.GetConfig()
	arg := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", configuration.DbUser, configuration.DbPass, configuration.DbHost, configuration.DbName)
	db, err := gorm.Open("mysql", arg)
	if err != nil {
		panic(err.Error())
	}
	//Migrate the schema
	db.AutoMigrate(&io.AthUser{})
	db.AutoMigrate(&io.AthAdminUser{})
	db.AutoMigrate(&io.AthCompanyUser{})
	db.AutoMigrate(&io.AthCompany{})
	db.AutoMigrate(&io.AthUserOTP{})
	db.LogMode(true)
	return db
}

func New(db *gorm.DB) Repository {
	// return  repository
	return &Service{
		DB: db,
	}
}

type Repository interface {
	CreateUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error)
	VerifySignUpTokenUser(ctx context.Context, data io.Verify) (err error)
	LoginUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error)
	ForgotPasswordUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error)
	ResetPasswordUser(ctx context.Context, data io.AthUser) (err error)
	DeleteUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error)

	CreateCompany(ctx context.Context, data io.AthCompany) (company io.AthCompany, err error)
	EditCompany(ctx context.Context, data io.AthCompany) (err error)
	CreateCompanyUser(ctx context.Context, data io.AthCompanyUser) error

	LoginAdmin(ctx context.Context, data io.AthAdminUser) (user io.AthAdminUser, err error)
	SignupAdmin(ctx context.Context, data io.AthAdminUser) (err error)

	CreateOTP(ctx context.Context, data io.AthUserOTP) (err error)
	VerifyOTP(ctx context.Context, data io.Verify) (err error)

	VerifyEmail(ctx context.Context, data io.EmailVerify) (err error)
	CreateUserProfile(ctx context.Context, data io.AthUserProfile) (user io.AthUserProfile, err error)

	CreateVenue(ctx context.Context, data io.AthVenues) (venue io.AthVenues, err error)
	EditVenue(ctx context.Context, data io.AthVenues) (venue io.AthVenues, err error)
	CreateVenueHoliday(ctx context.Context, data io.AthVenueHolidays) (venueHolidays io.AthVenueHolidays, err error)
	DeleteVenueHoliday(ctx context.Context, data io.DeleteVenueHolidays) (err error)

	SaveHoursOfOperation(ctx context.Context, data io.AthVenueHours) (hour io.AthVenueHours, err error)
	EditHoursOfOperation(ctx context.Context, data io.AthVenueHours) (hour io.AthVenueHours, err error)
	DeleteHoursOfOperation(ctx context.Context, data io.AthVenueHours) (err error)
	CreateFacility(ctx context.Context, data io.AthFacilities) (facility io.AthFacilities, err error)
	EditFacility(ctx context.Context, data io.AthFacilities) (facility io.AthFacilities, err error)
	CreateFacilitySlots(ctx context.Context, data io.AthFacilitySlots) (facility io.AthFacilitySlots, err error)
	EditFacilitySlots(ctx context.Context, data io.AthFacilitySlots) (facility []io.AthFacilitySlots, err error)
	BookFacility(ctx context.Context, data io.AthFacilityBookings) (facilityBooking io.AthFacilityBookings, err error)
	BookFacilitySlots(ctx context.Context, data io.AthFacilityBookingSlots) (facilityBookingslot io.AthFacilityBookingSlots, err error)
	GetSportCategories(ctx context.Context) (sportCategories io.AthSportCategories, err error)

	AddFacilityCustomRates(ctx context.Context, data io.AthFacilityCustomRates) (customRates io.AthFacilityCustomRates, err error)

	EditUserProfile(ctx context.Context, data io.AthUserProfile) (user io.AthUserProfile, err error)
	EditUser(ctx context.Context, data io.AthUser) (user io.AthUser, err error)
}
