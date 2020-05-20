package io

import (
	"time"
)

/*
create table ath_facilities (
  facility_id	integer(11)	auto_increment primary key not null,
  venue_id	int(11)	not null ,
  facility_name	varchar(250)	not null	,
  facility_base_price	float	default	NULL,
  time_slot	int(6)	default	NULL,
  facility_sport_categories	int(11)	,
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	REFERENCES ath_users(user_id),

  FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id)
);
*/

type AthFacilities struct {
	FacilityID              int     `gorm:"column:facility_id;primary_key" json:"facility_id"`
	VenueID                 int     `gorm:"column:venue_id;" json:"venue_id"`
	FacilityName            string  `gorm:"column:facility_name" json:"facility_name"`
	FacilityBasePrice       float32 `gorm:"column:facility_base_price" json:"facility_base_price"`
	TimeSlot                int     `gorm:"column:time_slot" json:"time_slot"`
	FacilitySportCategories string  `gorm:"column:facility_sport_categories" json:"facility_sport_categories"`
	IsActive                bool    `gorm:"column:is_active" json:"is_active"`
	Models
}

/*
create table ath_facility_slots (
  facility_slot_id	integer(11)	auto_increment primary key not null,
  facility_id	int(11)	,
  user_id	int(11)	,
  slot_days	varchar(10)	default	NULL,
  slot_type varchar(10)	default	NULL,
  slot_from_time varchar(10)	default	NULL,
  slot_to_time varchar(10)	default	NULL,
  slot_price	float	default	NULL,
  is_active	tinyint(1)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	REFERENCES ath_users(user_id),

  FOREIGN KEY (facility_id) REFERENCES ath_facilities(facility_id),
  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);
*/

type AthFacilitySlots struct {
	FacilitySlotID int     `gorm:"column:facility_slot_id;primary_key" json:"facility_slot_id"`
	FacilityID     int     `gorm:"column:facility_id;" json:"facility_id"`
	UserID         int     `gorm:"column:user_id" json:"user_id"`
	SlotDays       string  `gorm:"column:slot_days" json:"slot_days"`
	SlotType       string  `gorm:"column:slot_type" json:"slot_type"`
	SlotFromTime   string  `gorm:"column:slot_from_time" json:"slot_from_time"`
	SlotToTime     string  `gorm:"column:slot_to_time" json:"slot_to_time"`
	SlotPrice      float32 `gorm:"column:slot_price" json:"slot_price"`
	IsActive       bool    `gorm:"column:is_active" json:"is_active"`
	Models
}

/*
create table ath_sport_categories (
  sport_category_id	integer(11)	auto_increment primary key not null,
  category_name	varchar(100)	not null	,
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_admin_users(admin_id)

);
*/

type AthSportCategories struct {
	SportCategoryID int  `gorm:"column:sport_category_id;primary_key" json:"sport_category_id"`
	CategoryName    int  `gorm:"column:category_name;" json:"category_name"`
	IsActive        bool `gorm:"column:is_active" json:"is_active"`
	Models
}

/*

create table ath_facility_bookings (
  booking_id integer(11)	auto_increment primary key not null,
  user_id	int(11)	,
  booking_no	varchar(30)	not null	,
  booking_date	datetime	not null,
  base_total_amount	float	not null	,
  discount_amount	float	not null	,
  booking_amount	float	not null,
  booking_fee	float	not null,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_users(user_id)
);

*/

type AthFacilityBookings struct {
	BookingID       int       `gorm:"column:booking_id;primary_key" json:"booking_id"`
	UserID          int       `gorm:"column:user_id;" json:"user_id"`
	BookingNo       string    `gorm:"column:booking_no" json:"booking_no"`
	BookingDate     time.Time `gorm:"column:booking_date" json:"booking_date"`
	BaseTotalAmount float32   `gorm:"column:base_total_amount" json:"base_total_amount"`
	DiscountAmount  float32   `gorm:"column:discount_amount" json:"discount_amount"`
	BookingAmount   float32   `gorm:"column:booking_amount" json:"booking_amount"`
	BookingFee      float32   `gorm:"column:booking_fee" json:"booking_fee"`
	Models
}

/*

create table ath_facility_booking_slots (
  booking_slot_id	integer(11)	auto_increment primary key not null,
  booking_id	int(11)	default	NULL  ,
  facility_slot_id	int(11)	default	NULL ,
  slot_days	varchar(10)	default	NULL,
  slot_from_date	datetime	not null,
  slot_to_date	datetime	not null	,
  slot_booking_price	float	default	NULL,

  FOREIGN KEY (booking_id) REFERENCES ath_facility_bookings(booking_id),
  FOREIGN KEY (facility_slot_id) REFERENCES ath_facility_slots(facility_slot_id)
);

*/

type AthFacilityBookingSlots struct {
	BookingSlotID    int       `gorm:"column:booking_slot_id;primary_key" json:"booking_slot_id"`
	BookingID        int       `gorm:"column:booking_id;" json:"booking_id"`
	FaciliySlotID    int       `gorm:"column:facility_slot_id" json:"facility_slot_id"`
	SlotDays         string    `gorm:"column:slot_days" json:"slot_days"`
	SlotFromDate     time.Time `gorm:"column:slot_from_date" json:"slot_from_date"`
	SlotToDate       time.Time `gorm:"column:slot_to_date" json:"slot_to_date"`
	SlotBookingPrice float32   `gorm:"column:slot_booking_price" json:"slot_booking_price"`
}

/*
create table ath_facility_custom_rates (
  rate_id integer(11) auto_increment primary key not null,
  facility_id integer(11),
  user_id integer(11),
  facility_slot_id integer(11),
  date date,
  slot_price float,
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_users(user_id),

  FOREIGN KEY (facility_id) REFERENCES ath_facilities(facility_id),
  FOREIGN KEY (user_id) REFERENCES ath_users(user_id),
  FOREIGN KEY (facility_slot_id) REFERENCES ath_facility_slots(facility_slot_id)

);
*/

type AthFacilityCustomRates struct {
	RateID         int       `gorm:"column:rate_id;primary_key" json:"rate_id"`
	FacilityID     int       `gorm:"column:facility_id;" json:"facility_id"`
	UserID         int       `gorm:"column:user_id" json:"user_id"`
	FacilitySlotID int       `gorm:"column:facility_slot_id" json:"facility_slot_id"`
	Date           time.Time `gorm:"column:date" json:"date"`
	SlotPrice      float32   `gorm:"column:slot_price" json:"slot_price"`
	IsActive       bool      `gorm:"column:is_active" json:"is_active"`
	Models
}
