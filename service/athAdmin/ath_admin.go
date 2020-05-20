package athAdmin

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

type AthAdminUserService interface {
	LoginAdmin(ctx context.Context, u io.LoginRequest) (res io.Response)
	SignUpAdmin(ctx context.Context, u io.AthAdminUser) (res io.Response)
}

type athAdminUserService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthAdminService(DbRepo db.Repository) AthAdminUserService {
	return &athAdminUserService{
		DbRepo: DbRepo,
	}
}

func (b *athAdminUserService) LoginAdmin(ctx context.Context, u io.LoginRequest) (res io.Response) {
	fmt.Println("Inside login service")

	var userData io.AthAdminUser

	userData, res.Error = b.DbRepo.LoginAdmin(ctx, io.AthAdminUser{Email: u.Email, Password: u.Password})

	if res.Error != nil {
		res = io.FailureMessage(res.Error)
		return
	}
	if !userData.IsActive {
		log.Printf("Validation error: %v", "not verified")
		res = io.FailureMessage(fmt.Errorf(`you are not verified`))
		return
	}
	if !utile.ComparePasswords(userData.Password, []byte(u.Password)) {
		log.Printf("Password compare error: %v", "failed to compare password")
		res = io.FailureMessage(fmt.Errorf(`incorrect credentials`))
		return
	}
	data := make(map[string]string)
	data["user_id"] = strconv.Itoa(userData.ID)
	res = io.SuccessMessage(data)
	return
}

func (b *athAdminUserService) verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
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

func (b *athAdminUserService) SignUpAdmin(ctx context.Context, u io.AthAdminUser) (res io.Response) {
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

	u.CreatedBy = u.Models.CreatedBy
	fmt.Println("u : ", u)

	err := b.DbRepo.SignupAdmin(ctx, u)
	if err != nil {
		fmt.Println("if loop err")
		res = io.FailureMessage(res.Error, "email already exists")
		res.Success = false
		res.Error = err
		fmt.Println("res : ", res)
		fmt.Println("res.Error : ", res.Error)
		return
	}
	fmt.Println("user saved---")
	res = io.SuccessMessage(nil, "User has been saved")

	return
}
