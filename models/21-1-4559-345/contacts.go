package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Contacts struct {
	bun.BaseModel             `bun:"table:contacts"`
	Id                        string    `bun:"id,type:varchar(16),pk"`
	AddressId                 float64   `bun:"address_id,type:decimal(19,0)"`
	Salutation                string    `bun:"salutation,type:varchar(4),nullzero"`
	FirstName                 string    `bun:"first_name,type:varchar(15)"`
	Mi                        string    `bun:"mi,type:varchar(2),nullzero"`
	LastName                  string    `bun:"last_name,type:varchar(24)"`
	Title                     string    `bun:"title,type:varchar(40),nullzero"`
	DirectPhone               string    `bun:"direct_phone,type:varchar(20),nullzero"`
	PhoneExt                  string    `bun:"phone_ext,type:varchar(5),nullzero"`
	DirectFax                 string    `bun:"direct_fax,type:varchar(20),nullzero"`
	FaxExt                    string    `bun:"fax_ext,type:varchar(5),nullzero"`
	Beeper                    string    `bun:"beeper,type:varchar(20),nullzero"`
	Cellular                  string    `bun:"cellular,type:varchar(20),nullzero"`
	EmailAddress              string    `bun:"email_address,type:varchar(255),nullzero"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	HomeAddress1              string    `bun:"home_address1,type:varchar(40),nullzero"`
	HomeAddress2              string    `bun:"home_address2,type:varchar(40),nullzero"`
	Birthday                  time.Time `bun:"birthday,type:datetime,nullzero"`
	Anniversary               time.Time `bun:"anniversary,type:datetime,nullzero"`
	HomePhone                 string    `bun:"home_phone,type:varchar(20),nullzero"`
	EmailAddress2             string    `bun:"email_address2,type:varchar(255),nullzero"`
	DeleteFlag                string    `bun:"delete_flag,type:char"`
	Comment1                  string    `bun:"comment_1,type:varchar(30),nullzero"`
	Comment2                  string    `bun:"comment_2,type:varchar(30),nullzero"`
	Comment3                  string    `bun:"comment_3,type:varchar(30),nullzero"`
	Comment4                  string    `bun:"comment_4,type:varchar(30),nullzero"`
	Salesrep                  string    `bun:"salesrep,type:char,nullzero"`
	Buyer                     string    `bun:"buyer,type:char,default:('N')"`
	Schedular                 string    `bun:"schedular,type:char,nullzero"`
	LoginId                   string    `bun:"login_id,type:varchar(30),nullzero"`
	Dealer                    string    `bun:"dealer,type:char,nullzero"`
	Url                       string    `bun:"url,type:varchar(255),nullzero"`
	Comments                  string    `bun:"comments,type:text(2147483647),nullzero"`
	HomeEmailAddress          string    `bun:"home_email_address,type:varchar(255),nullzero"`
	HomeFax                   string    `bun:"home_fax,type:char(20),nullzero"`
	Class1id                  string    `bun:"class_1id,type:varchar(255),nullzero"`
	Class2id                  string    `bun:"class_2id,type:varchar(255),nullzero"`
	Class3id                  string    `bun:"class_3id,type:varchar(255),nullzero"`
	Class4id                  string    `bun:"class_4id,type:varchar(255),nullzero"`
	Class5id                  string    `bun:"class_5id,type:varchar(255),nullzero"`
	EmployeeVendorId          float64   `bun:"employee_vendor_id,type:decimal(19,0),nullzero"`
	DearField                 string    `bun:"dear_field,type:varchar(50),nullzero"`
	OldContactId              string    `bun:"old_contact_id,type:varchar(20),nullzero"`
	Employee                  string    `bun:"employee,type:char,nullzero"`
	UpperCombinedName         string    `bun:"upper_combined_name,type:varchar(50),nullzero"`
	DescendingCombinedName    string    `bun:"descending_combined_name,type:varchar(50),nullzero"`
	DirectWattsNumber         string    `bun:"direct_watts_number,type:varchar(20),nullzero"`
	NoOfCycleDays             int32     `bun:"no_of_cycle_days,type:int,nullzero"`
	Mailstop                  string    `bun:"mailstop,type:varchar(30),nullzero"`
	CommissionClassId         string    `bun:"commission_class_id,type:varchar(8),nullzero"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),nullzero"`
	VendorId                  float64   `bun:"vendor_id,type:decimal(19,0),nullzero"`
	AddressName               string    `bun:"address_name,type:varchar(255),nullzero"`
	Driver                    string    `bun:"driver,type:char,default:('N')"`
	InsideSalesrepFlag        string    `bun:"inside_salesrep_flag,type:char,default:('N')"`
	DefaultBranchId           string    `bun:"default_branch_id,type:varchar(8),nullzero"`
	Technician                string    `bun:"technician,type:char,default:('N')"`
	LocationId                float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	TerritoryUid              int32     `bun:"territory_uid,type:int,nullzero"`
	FuelSurchargePercentage   float64   `bun:"fuel_surcharge_percentage,type:decimal(19,4),nullzero"`
	MaxFuelChargePerShip      float64   `bun:"max_fuel_charge_per_ship,type:decimal(19,4),nullzero"`
	ContactTypeId             int32     `bun:"contact_type_id,type:int,nullzero"`
	SalesManagerId            string    `bun:"sales_manager_id,type:varchar(16),nullzero"`
	RoadnetDriverId           string    `bun:"roadnet_driver_id,type:varchar(255),nullzero"`
	CellularExt               string    `bun:"cellular_ext,type:varchar(5),nullzero"`
	ContactRoleUid            int32     `bun:"contact_role_uid,type:int,nullzero"`
	SfdcAccountId             string    `bun:"sfdc_account_id,type:varchar(255),nullzero"`
	SfdcCreateDate            time.Time `bun:"sfdc_create_date,type:datetime,nullzero"`
	SfdcUpdateDate            time.Time `bun:"sfdc_update_date,type:datetime,nullzero"`
	SfdcContactId             string    `bun:"sfdc_contact_id,type:varchar(255),nullzero"`
	SalesAgencyFlag           string    `bun:"sales_agency_flag,type:char,nullzero"`
	SalesAgencyName           string    `bun:"sales_agency_name,type:varchar(50),nullzero"`
	DeliveryOutputCd          int32     `bun:"delivery_output_cd,type:int,default:((2844))"`
	DriverEnableGpsFlag       string    `bun:"driver_enable_gps_flag,type:char,default:('N')"`
	DateDirectFaxLastModified time.Time `bun:"date_direct_fax_last_modified,type:datetime,default:(getdate())"`
	AdsUser                   string    `bun:"ads_user,type:varchar(50),nullzero"`
	SalesrepDefaultLocationId float64   `bun:"salesrep_default_location_id,type:decimal(19,0),nullzero"`
	DatagateSubjectId         int32     `bun:"datagate_subject_id,type:int,nullzero"`
	RentalId                  string    `bun:"rental_id,type:varchar(255),nullzero"`
	EmailOptOut               string    `bun:"email_opt_out,type:char,default:('N')"`
	DriverRfcNo               string    `bun:"driver_rfc_no,type:varchar(255),nullzero"`
	DriverLicenseNo           string    `bun:"driver_license_no,type:varchar(255),nullzero"`
	DriverRegistrationId      string    `bun:"driver_registration_id,type:varchar(255),nullzero"`
}
