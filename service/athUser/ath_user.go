package athUser

import (
	"backend/pkg/db"
	"backend/pkg/io"
	"backend/pkg/utile"
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"
)

type AthUserService interface {
	LoginUser(ctx context.Context, u io.LoginRequest) (res io.Response)
	SignUpUser(ctx context.Context, u io.AthUser) (res io.Response)
	EditUser(ctx context.Context, u io.AthUser) (res io.Response)
	EmailVerify(ctx context.Context, u io.EmailVerify) (res io.Response)
	PhoneVerifyUser(ctx context.Context, u io.Verify) (res io.Response)
	ResendCode(ctx context.Context, u io.Verify) (res io.Response)
	ForgotPasswordUser(ctx context.Context, u io.ForgotPasswordRequest) (res io.Response)
	ResetPasswordUser(ctx context.Context, u io.ResetPasswordRequest) (res io.Response)
	CreateUserProfile(ctx context.Context, u io.AthUserProfile) (res io.Response)
	EditUserProfile(ctx context.Context, u io.AthUserProfile) (res io.Response)
}

type athUserService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthUserService(DbRepo db.Repository) AthUserService {
	return &athUserService{
		DbRepo: DbRepo,
	}
}

func (b *athUserService) LoginUser(ctx context.Context, u io.LoginRequest) (res io.Response) {

	var userData io.AthUser

	userData, res.Error = b.DbRepo.LoginUser(ctx, io.AthUser{Email: u.Email, Password: u.Password})
	if res.Error != nil {
		res = io.FailureMessage(res.Error)
		res.Error = res.Error
		return
	}
	if !userData.IsActive {
		log.Printf("Validation error: %v", "not verified")
		res = io.FailureMessage(fmt.Errorf(`you are not verified`))
		res.Error = res.Error
		return
	}
	if !utile.ComparePasswords(userData.Password, []byte(u.Password)) {
		log.Printf("Password compare error: %v", "failed to compare password")
		res = io.FailureMessage(fmt.Errorf(`incorrect credentials`))
		res.Error = res.Error
		return
	}
	data := make(map[string]string)
	data["user_id"] = strconv.Itoa(userData.ID)
	res = io.SuccessMessage(data)
	return
}

func (b *athUserService) verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	//sevenOrMore = letters >= 7
	sevenOrMore = len(s) >= 6
	return
}

func (b *athUserService) SignUpUser(ctx context.Context, u io.AthUser) (res io.Response) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(u.Email) == false {
		log.Printf("Validation error: %v", "incorrect email format")
		res.Error = fmt.Errorf("invalid email")
		res.Message = "invalid email address"
		return
	}
	if len(u.Password) == 0 || len(u.Password) < 8 {
		log.Printf("Validation error: %v", "password length must be greater or equal to 8")
		res.Error = fmt.Errorf("password length must be greater or equal to 8")
		res.Message = "invalid password"
		return

	}
	if len(u.Password) > 0 {
		sevenOrMore, number, upper, special := b.verifyPassword(u.Password)
		if !sevenOrMore || !number || !upper || !special {
			log.Printf("Validation error: %v", "password must contain atleast a number,uppercase alphabet and special character")
			res.Error = fmt.Errorf("password must contain atleast a number,uppercase alphabet and special character")
			res.Message = "invalid password"
			res = io.FailureMessage(res.Error)
			return
		}
	}
	if u.Password, res.Error = utile.HashAndSalt([]byte(u.Password)); res.Error != nil {
		log.Printf("hash password generate error: %v", res.Error)
		res.Message = "invalid password"
		res = io.FailureMessage(res.Error)
		return
	}

	u.UserSource = u.UserSource
	u.CreatedBy = u.Models.CreatedBy

	newUser, err := b.DbRepo.CreateUser(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "email already exists")
		res.Success = false
		res.Error = err
		return
	}

	data := make(map[string]interface{})
	data["user_id"] = newUser.ID

	/*
		emailMsg := "Email verification code:" + PhoneVerifyToken
		res.Error = email.NewEmailSend(u.Email, "Verify Email", emailMsg)
		if res.Error != nil {
			b.DbRepo.Delete(ctx, io.AthUser{Email: u.Email})
			res = io.FailureMessage(res.Error, "Failed to send email")
			return
		}
	*/

	res.Data = data
	res = io.SuccessMessage(data, "User has been saved")

	return
}

func (b *athUserService) EditUser(ctx context.Context, u io.AthUser) (res io.Response) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(u.Email) == false {
		log.Printf("Validation error: %v", "incorrect email format")
		res.Error = fmt.Errorf("invalid email")
		res.Message = "invalid email address"
		return
	}

	_, err := b.DbRepo.EditUser(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "error updating user")
		res.Success = false
		res.Error = err
		return
	}

	res = io.SuccessMessage(nil, "User has been updated")

	return
}

func (b *athUserService) EmailVerify(ctx context.Context, u io.EmailVerify) (res io.Response) {
	err := b.DbRepo.VerifyEmail(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error updating isactive flag")
		res.Success = false
		res.Error = err
		return
	}

	res = io.SuccessMessage(nil, "Email verified successfully")

	return
}

func (b *athUserService) PhoneVerifyUser(ctx context.Context, u io.Verify) (res io.Response) {
	if res.Error = b.DbRepo.VerifySignUpTokenUser(ctx, u); res.Error != nil {
		res = io.FailureMessage(res.Error)
		res.Error = res.Error
		return
	}
	res = io.SuccessMessage(nil, "Your account has been verified")
	return
}

func (b *athUserService) ResendCode(ctx context.Context, u io.Verify) (res io.Response) {
	if res.Error = b.DbRepo.VerifySignUpTokenUser(ctx, u); res.Error != nil {
		res = io.FailureMessage(res.Error)
		res.Error = res.Error
		return
	}
	res = io.SuccessMessage(nil, "Your account has been verified")
	return
}

func (b *athUserService) ForgotPasswordUser(ctx context.Context, u io.ForgotPasswordRequest) (res io.Response) {
	ForgotToken := utile.RandomString(6)
	Forgotdata, err := b.DbRepo.ForgotPasswordUser(ctx, io.AthUser{Email: u.Email, ResetPasswordToken: ForgotToken})

	if err != nil {
		res = io.FailureMessage(res.Error)
		res.Error = fmt.Errorf("Record not found")
		return
	}

	data := make(map[string]interface{})
	data["forgot_password_token"] = Forgotdata.ResetPasswordToken

	/*
		if res.Error = email.NewEmailSend(u.Email, "Forgot Password Email", "Email Forgot password code:"+ForgotToken); res.Error != nil {
			res = io.FailureMessage(res.Error, "Failed to send email")
			return
		}*/
	res = io.SuccessMessage(data, "An email has been sent to Forgot password")
	return
}

func (b *athUserService) ResetPasswordUser(ctx context.Context, u io.ResetPasswordRequest) (res io.Response) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(u.Email) == false {
		log.Printf("Validation error: %v", "incorrect email format")
		res.Error = fmt.Errorf("invalid email")
		res.Message = "invalid email address"
		return
	}
	if len(u.Password) == 0 || len(u.Password) < 8 {
		log.Printf("Validation error: %v", "password length must be greater or equal to 8")
		res.Error = fmt.Errorf("password length must be greater or equal to 8")
		res.Message = "invalid password"
		return

	}
	if len(u.Password) > 0 {
		sevenOrMore, number, upper, special := b.verifyPassword(u.Password)
		if !sevenOrMore || !number || !upper || !special {
			log.Printf("Validation error: %v", "password must contain atleast a number,uppercase alphabet and special character")
			res.Error = fmt.Errorf("password must contain atleast a number,uppercase alphabet and special character")
			res.Message = "invalid password"
			res = io.FailureMessage(res.Error)
			return
		}
	}
	if u.Password, res.Error = utile.HashAndSalt([]byte(u.Password)); res.Error != nil {
		log.Printf("hash password generate error: %v", res.Error)
		res.Message = "invalid password"
		res = io.FailureMessage(res.Error)
		return
	}

	u.Password = u.Password

	err := b.DbRepo.ResetPasswordUser(ctx, io.AthUser{Email: u.Email, Password: u.Password})
	if err != nil {
		res = io.FailureMessage(res.Error, "email already exists")
		return
	}

	res = io.SuccessMessage(nil, "Password has been reset successfully")
	return
}

func (b *athUserService) CreateUserProfile(ctx context.Context, u io.AthUserProfile) (res io.Response) {
	newProfile, err := b.DbRepo.CreateUserProfile(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error creating new user profile")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["user_profile_id"] = newProfile.UserProfileID

	res.Data = data
	res = io.SuccessMessage(data, "User Profile created")
	return
}

func (b *athUserService) EditUserProfile(ctx context.Context, u io.AthUserProfile) (res io.Response) {
	_, err := b.DbRepo.EditUserProfile(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error editing user profile")
		res.Error = err
		return
	}

	res = io.SuccessMessage(nil, "User Profile updated")
	return
}
