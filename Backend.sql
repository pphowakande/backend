DROP TABLE IF EXISTS ath_admin_users CASCADE;
create table ath_admin_users (
  admin_id	integer(11)	auto_increment primary key not null,	
  user_name	varchar(50)	default	NULL,
  password	varchar(50)	default	NULL,
  name	varchar(100)	default	NULL,
  email_address	varchar(100)	default	NULL,
  is_active	tinyint(1)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	integer(11)	REFERENCES ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	integer(11)	REFERENCES ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	integer(11) REFERENCES ath_admin_users(admin_id)
);

DROP TABLE IF EXISTS ath_amenities CASCADE;
create table ath_amenities (
  amenity_id integer(11)	auto_increment primary key not null,	
  amenity_name varchar(100) not null,	
  is_active	tinyint(1)	default	1,
  created_at datetime	default	NULL,
  created_by int(11) REFERENCES ath_admin_users(admin_id) ,
  updated_at datetime	default	NULL,
  updated_by int(11) REFERENCES ath_admin_users(admin_id),
  deleted_at datetime	default	NULL,
  deleted_by int(11) REFERENCES ath_admin_users(admin_id)

);

DROP TABLE IF EXISTS ath_companies CASCADE;
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
  created_by	int(11)	REFERENCES ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_admin_users(admin_id)

);

DROP TABLE IF EXISTS ath_users CASCADE;
create table ath_users (
  user_id	integer(11)	auto_increment primary key not null,
  email_address	varchar(100)	not null,
  password	varchar(100)	default	NULL,
  reset_password_token	varchar(50)	default	NULL,
  reset_password_token_created_at	datetime	default	NULL,
  last_password_reset_at	datetime	default	NULL,
  user_source	tinyint(1)	default	NULL,
  last_login_at	datetime	default	NULL,
  last_login_ip	varchar(30)	not null	,
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11) REFERENCES ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_admin_users(admin_id)
);

DROP TABLE IF EXISTS ath_company_users CASCADE;
create table ath_company_users (
  company_user_id integer(11)	auto_increment primary key not null,	
  company_id	int(11)	,
  user_id	int(11),
  created_at	datetime	default	NULL,
  created_by	int(11) REFERENCES ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	REFERENCES ath_admin_users(admin_id) ,

  FOREIGN KEY (user_id) REFERENCES ath_users(user_id),
  FOREIGN KEY (company_id) REFERENCES ath_companies(company_id)
);

DROP TABLE IF EXISTS ath_discounts CASCADE;
create table ath_discounts (
  discount_id integer(11)	auto_increment primary key not null,
  discount_code	varchar(30)		NOT NULL,
  discount_percent int(3)	default	NULL,
  discount_min	float	default	NULL,
  discount_max	float	default	NULL,
  discount_amount	float	default	NULL,
  discount_from_date	datetime	default	NULL,
  discount_to_date	datetime	default	NULL,
  discount_max_value	float	not null	,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_users(user_id)
);

DROP TABLE IF EXISTS ath_payment_gateway CASCADE;
create table ath_payment_gateway (
  payment_gateway_id	integer(11)	auto_increment primary key not null,
  name	varchar(50)	default	NULL,
  app_id	varchar(50)	default	NULL,
  token_key	varchar(100)	default	NULL,
  is_active	tinyint(1)	default	1,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_admin_users(admin_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_admin_users(admin_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11) REFERENCES ath_admin_users(admin_id)
);

DROP TABLE IF EXISTS ath_facility_bookings CASCADE;
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

DROP TABLE IF EXISTS ath_payments CASCADE;
create table ath_payments (
  payment_id	integer(11)	auto_increment primary key not null,
  booking_id	int(11)	,
  user_id	int(11)	,
  payment_amount	float	default	NULL,
  payment_collected	float	default	NULL,
  payment_mode	tinyint(1)	default	NULL,
  cheque_no	varchar(20)	default	NULL,
  bank_name	varchar(100)	default	NULL,
  payment_gateway_id	int(11)	,
  txt_id	varchar(50)	default	NULL,
  payment_status	varchar(50)	default	NULL,
  payment_message	text	default	NULL,
  payment_date	datetime	default	NULL,
  payment_ip_address	varchar(30)	NOT NULL,

  FOREIGN KEY (booking_id) REFERENCES ath_facility_bookings(booking_id),
  FOREIGN KEY (user_id) REFERENCES ath_users(user_id),
  FOREIGN KEY (payment_gateway_id) REFERENCES ath_payment_gateway(payment_gateway_id)
);


DROP TABLE IF EXISTS ath_sport_categories CASCADE;
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

DROP TABLE IF EXISTS ath_venues CASCADE;
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

DROP TABLE IF EXISTS ath_facilities CASCADE;
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

DROP TABLE IF EXISTS ath_facility_booking_discounts CASCADE;
create table ath_facility_booking_discounts (
  booking_discount_id	integer(11)	auto_increment primary key not null,
  booking_id	int(11)	,
  discount_id	int(11)	,
  created_at	datetime	default	NULL,

  FOREIGN KEY (booking_id) REFERENCES ath_facility_bookings(booking_id),
  FOREIGN KEY (discount_id) REFERENCES ath_discounts(discount_id)
);

DROP TABLE IF EXISTS ath_facility_slots CASCADE;
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

DROP TABLE IF EXISTS ath_facility_booking_slots CASCADE;
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

DROP TABLE IF EXISTS ath_user_profiles CASCADE;
create table ath_user_profiles (
  user_profile_id	integer(11)	auto_increment primary key not null,	
  user_id	int(11)	,
  first_name	varchar(50)	not null	,
  last_name	varchar(50)	default	NULL,
  contact_no	varchar(30)	not null,
  gender	enum('m', 'f', 'o')	default	NULL,
  dob	date	default	NULL,
  profile_image	varchar(100)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  updated_at	datetime	default	NULL,
  updated_by	int(11)	REFERENCES ath_users(user_id),

  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);



DROP TABLE IF EXISTS ath_venue_images CASCADE;
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

DROP TABLE IF EXISTS ath_venue_review_ratings CASCADE;
create table ath_venue_review_ratings (
  review_rating_id	integer(11)	auto_increment primary key not null,	
  venue_id	int(11)	default	NULL,
  user_id	int(11)	,
  review_comment	varchar(250)	default	NULL,
  review_rating	tinyint(1)	default	NULL,
  created_at	datetime,
  created_ip_address	varchar(30)	default	NULL,
  is_active	tinyint(1)	default	1,

  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);

DROP TABLE IF EXISTS ath_user_social CASCADE;
create table ath_user_social (
  social_id	integer(11)	auto_increment primary key not null,	
  user_id	int(11) ,
  fb_auth_token	varchar(255)	default	NULL,
  fb_auth_token_expiry	int(11)		,
  google_auth_token	varchar(255)	default	NULL,
  google_auth_token_expiry	int(11)	,	
  created_at	datetime	default	NULL,
  updated_at	datetime	default	NULL,

  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);

DROP TABLE IF EXISTS ath_venue_users CASCADE;
create table ath_venue_users (
  venue_user_id	integer(11)	auto_increment primary key not null,	
  venue_id	int(11)	,
  user_id	int(11) ,
  user_designation	varchar(100)	default	NULL,
  created_at	datetime	default	NULL,
  created_by	int(11)	REFERENCES ath_users(user_id),
  deleted_at	datetime	default	NULL,
  deleted_by	int(11)	REFERENCES ath_users(user_id),

  FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id),
  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);

DROP TABLE IF EXISTS ath_user_otps CASCADE;
create table ath_user_otps (
  otp_id	integer(11)	auto_increment primary key not null,		
  user_id	int(11) ,
  otp_no	varchar(100)	not	NULL,
  otp_type	varchar(100)	not	NULL,
  otp_expiry	int(11)	not	NULL	,
  created_at	datetime	default	NULL,
  updated_at	datetime	default	NULL,
  is_active	tinyint(1)	default	1,

  FOREIGN KEY (user_id) REFERENCES ath_users(user_id)
);

DROP TABLE IF EXISTS ath_venue_hours CASCADE;
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
  deleted_by	int(11) REFERENCES ath_users(user_id),

 FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id)
);

DROP TABLE IF EXISTS ath_venue_holidays CASCADE;
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
  deleted_by	int(11) REFERENCES ath_users(user_id),
  FOREIGN KEY (venue_id) REFERENCES ath_venues(venue_id)

);

DROP TABLE IF EXISTS ath_facility_custom_rates CASCADE;
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
