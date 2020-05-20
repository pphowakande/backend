package athCompany

import (
	"backend/pkg/db"
	"backend/pkg/io"
	"context"
)

type AthCompanyService interface {
	Createcompany(ctx context.Context, u io.AthCompany) (res io.Response)
	EditCompany(ctx context.Context, u io.AthCompany) (res io.Response)
	CreatecompanyUser(ctx context.Context, u io.AthCompanyUser) (res io.Response)
}

type athCompanyService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

func NewBasicAthCompanyService(DbRepo db.Repository) AthCompanyService {
	return &athCompanyService{
		DbRepo: DbRepo,
	}
}

func (b *athCompanyService) Createcompany(ctx context.Context, u io.AthCompany) (res io.Response) {

	newCompany, err := b.DbRepo.CreateCompany(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "email already exists")
		res.Error = err
		return
	}
	data := make(map[string]interface{})
	data["company_id"] = newCompany.ID

	res.Data = data
	res = io.SuccessMessage(data, "company created")
	return
}

func (b *athCompanyService) EditCompany(ctx context.Context, u io.AthCompany) (res io.Response) {

	err := b.DbRepo.EditCompany(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "Error updating company")
		res.Error = err
		return
	}
	res = io.SuccessMessage(nil, "company created")
	return
}

func (b *athCompanyService) CreatecompanyUser(ctx context.Context, u io.AthCompanyUser) (res io.Response) {
	// u.CompanyID = u.CompanyID
	// u.UserID = u.UserID
	// u.Models.CreatedBy = u.Models.CreatedBy

	err := b.DbRepo.CreateCompanyUser(ctx, u)
	if err != nil {
		res = io.FailureMessage(res.Error, "email already exists")
		res.Error = err
		return
	}

	data := make(map[string]interface{})
	res.Data = data
	res = io.SuccessMessage(data, "Company User has been added successfully")

	return
}
