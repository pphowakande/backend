package io

/*

create table ath_venues (
  venue_id	integer(11)	auto_increment primary key not null,
  company_id	int(11)	,
  venue_name	varchar(100)	default	NULL,
  email_address	varchar(100)	default	NULL,
  contact_no	varchar(30)	default	NULL,
  address	varchar(500)	default	NULL,
  get_lag	varchar(20)	default	NULL,
  geo_long	varchar(20)	default	NULL,
  description	varchar(300)	default	NULL,
  ammenities	varchar(100)	default	NULL,
  is_active	int(1)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	int(11) REFERENCES ath_users(user_id)	,
  updated_at	datetime	default	NULL,
  updated_by	int(11)	 REFERENCES ath_users(user_id) ,
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_users(user_id) ,

  FOREIGN KEY (company_id) REFERENCES ath_companies(company_id)
);

*/

type AthVenues struct {
	VenueID      int    `gorm:"column:venue_id;primary_key" json:"venue_id"`
	CompanyID    int    `gorm:"column:company_id;" json:"company_id"`
	VenueName    string `gorm:"column:venue_name" json:"venue_name"`
	EmailAddress string `gorm:"column:email_address" json:"email_address"`
	ContactNo    string `gorm:"column:contact_no" json:"contact_no"`
	Address      string `gorm:"column:address" json:"address"`
	GetLag       string `gorm:"column:get_lag" json:"get_lag"`
	GetLong      string `gorm:"column:geo_long" json:"geo_long"`
	Description  string `gorm:"column:description" json:"description"`
	Amenities    string `gorm:"column:ammenities" json:"ammenities"`
	IsActive     bool   `gorm:"column:is_active" json:"is_active"`
	Models
}

/*
create table ath_venue_images (
  venue_image_id	integer(11)	auto_increment primary key not null,
  venue_id	int(11)	not null,
  image_title	varchar(100)	default	NULL,
  image_url	TEXT	not null	,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	REFERENCES ath_users(user_id),

  FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id)
);
*/

type AthVenueImages struct {
	VenueImageID int    `gorm:"column:venue_image_id;primary_key" json:"venue_image_id"`
	VenueID      int    `gorm:"column:venue_id;" json:"venue_id"`
	ImageTitle   string `gorm:"column:image_title" json:"image_title"`
	ImageUrl     string `gorm:"column:image_url" json:"image_url"`
	Models
}

/*
create table ath_venue_hours (
 hour_id integer(11) auto_increment primary key not null,
 venue_id integer(11),
 day integer(10),
 opening_time varchar(20) not null,
 closing_time varchar(20) not null,
 is_active	tinyint(1)	default	1,
 created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_users(user_id)

 FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id)
);
*/

type AthVenueHours struct {
	HourID      int    `gorm:"column:hour_id;primary_key" json:"hour_id"`
	VenueID     int    `gorm:"column:venue_id;" json:"venue_id"`
	Day         int    `gorm:"column:day" json:"day"`
	OpeningTime string `gorm:"column:opening_time" json:"opening_time"`
	ClosingTime string `gorm:"column:closing_time" json:"closing_time"`
	IsActive    bool   `gorm:"column:is_active" json:"is_active"`
	Models
}

/*
create table ath_venue_holidays (
  holiday_id integer(11) auto_increment primary key not null,
  venue_id integer(11),
  title varchar(100),
  day integer(10),
  month varchar(20),
  year varchar(20),
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_users(user_id)
  FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id)

);
*/

type AthVenueHolidays struct {
	HolidayID int    `gorm:"column:holiday_id;primary_key" json:"holiday_id"`
	VenueID   int    `gorm:"column:venue_id;" json:"venue_id"`
	Title     string `gorm:"column:title" json:"title"`
	Day       int    `gorm:"column:day" json:"day"`
	Month     string `gorm:"column:month" json:"month"`
	Year      string `gorm:"column:year" json:"year"`
	IsActive  bool   `gorm:"column:is_active" json:"is_active"`
	Models
}
