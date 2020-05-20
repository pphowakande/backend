package io

/*
create table ath_companies (
  company_id integer(11)	auto_increment primary key not null,
  company_name	varchar(250)	default	NULL,
  address	varchar(500)	not null	,
  contact_no	varchar(30)	default	NULL,
  email_address	varchar(100)	default	NULL,
  gst_no	varchar(50)	not null	,
  gst_no_file	varchar(100)	not null	,
  pan_no	varchar(20)	not null	,
  pan_no_file	varchar(100)	not null	,
  bank_account_no	varchar(20)	not null	,
  is_active	tinyint(1)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	int(11)	references ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	references ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	references ath_admin_users(admin_id)
);

create table ath_company_users (
  company_user_id integer(11)	auto_increment primary key not null,
  company_id	int(11)	references ath_companies(company_id),
  user_id	int(11)	references ath_users(user_id),
  created_at	datetime	default	NULL,
  created_by	int(11)	references ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	references ath_admin_users(admin_id)
);

*/
type AthCompany struct {
	ID        int    `gorm:"column:company_id;primary_key" json:"company_id"`
	Name      string `gorm:"column:company_name" json:"company_name"`
	Address   string `gorm:"column:address" json:"address"`
	Contact   string `gorm:"column:contact_no" json:"contact_no"`
	Email     string `gorm:"column:email_address" json:"email_address"`
	GstNo     string `gorm:"column:gst_no" json:"gst_no"`
	GstNoFile string `gorm:"column:gst_no_file" json:"gst_no_file"`
	PanNo     string `gorm:"column:pan_no" json:"pan_no"`
	PanNoFile string `gorm:"column:pan_no_file" json:"pan_no_file"`
	BankAccNo string `gorm:"column:bank_account_no" json:"bank_account_no"`
	IsActive  bool   `gorm:"column:is_active" json:"is_active"`
	Models
}

type AthCompanyUser struct {
	ID        int `gorm:"column:company_user_id;primary_key" json:"company_user_id"`
	CompanyID int `gorm:"column:company_id" json:"company_id"`
	UserID    int `gorm:"column:user_id" json:"user_id"`
	Models
}
