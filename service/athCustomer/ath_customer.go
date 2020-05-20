package athCustomer

import (
	"backend/pkg/db"
	"backend/pkg/io"
	"context"
	"fmt"
	"log"
	"regexp"
)

type AthCustomerService interface {
	CreateCustomer(ctx context.Context, u io.AthUser) (res io.Response)
}

type athCustomerService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthCustomerService(DbRepo db.Repository) AthCustomerService {
	return &athCustomerService{
		DbRepo: DbRepo,
	}
}

func (b *athCustomerService) CreateCustomer(ctx context.Context, u io.AthUser) (res io.Response) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(u.Email) == false {
		log.Printf("Validation error: %v", "incorrect email format")
		res.Error = fmt.Errorf("invalid email")
		res.Message = "invalid email address"
		return
	}
	newUser, err := b.DbRepo.CreateUser(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "email already exists")
		res.Success = false
		res.Error = err
		return
	}

	data := make(map[string]interface{})
	data["user_id"] = newUser.ID

	res.Data = data
	res = io.SuccessMessage(data, "Customer has been saved")

	return
}
