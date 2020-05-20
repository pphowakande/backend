package io

/*

create table ath_admin_users (
  admin_id	integer(11)	auto_increment primary key not null,
  user_name	varchar(50)	default	NULL,
  password	varchar(50)	default	NULL,
  name	varchar(100)	default	NULL,
  email_address	varchar(100)	default	NULL,
  is_active	tinyint(1)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	integer(11)	references ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	integer(11)	references ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	integer(11)	references ath_admin_users(admin_id)
);

*/

type AthAdminUser struct {
	ID       int    `gorm:"column:admin_id;primary_key" json:"admin_id"`
	UserName string `gorm:"column:user_name" json:"user_name"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email_address" json:"email_address"`
	IsActive bool   `gorm:"column:is_active" json:"is_active"`
	Models
}
