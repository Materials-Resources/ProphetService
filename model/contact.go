package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"table:contacts"`

	Id                        string          `bun:"id,type:varchar(16),pk"`
	AddressId                 float64         `bun:"address_id,type:decimal(19,0)"`
	Salutation                sql.NullString  `bun:"salutation,type:varchar(4)"`
	FirstName                 string          `bun:"first_name,type:varchar(15)"`
	Mi                        sql.NullString  `bun:"mi,type:varchar(2)"`
	LastName                  string          `bun:"last_name,type:varchar(24)"`
	Title                     sql.NullString  `bun:"title,type:varchar(40)"`
	DirectPhone               sql.NullString  `bun:"direct_phone,type:varchar(20)"`
	PhoneExt                  sql.NullString  `bun:"phone_ext,type:varchar(5)"`
	DirectFax                 sql.NullString  `bun:"direct_fax,type:varchar(20)"`
	FaxExt                    sql.NullString  `bun:"fax_ext,type:varchar(5)"`
	Beeper                    sql.NullString  `bun:"beeper,type:varchar(20)"`
	Cellular                  sql.NullString  `bun:"cellular,type:varchar(20)"`
	EmailAddress              sql.NullString  `bun:"email_address,type:varchar(255)"`
	DateCreated               time.Time       `bun:"date_created,type:datetime"`
	DateLastModified          time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string          `bun:"last_maintained_by,type:varchar(30)"`
	HomeAddress1              sql.NullString  `bun:"home_address1,type:varchar(40)"`
	HomeAddress2              sql.NullString  `bun:"home_address2,type:varchar(40)"`
	Birthday                  sql.NullTime    `bun:"birthday,type:datetime"`
	Anniversary               sql.NullTime    `bun:"anniversary,type:datetime"`
	HomePhone                 sql.NullString  `bun:"home_phone,type:varchar(20)"`
	EmailAddress2             sql.NullString  `bun:"email_address2,type:varchar(255)"`
	DeleteFlag                string          `bun:"delete_flag,type:char"`
	Comment1                  sql.NullString  `bun:"comment_1,type:varchar(30)"`
	Comment2                  sql.NullString  `bun:"comment_2,type:varchar(30)"`
	Comment3                  sql.NullString  `bun:"comment_3,type:varchar(30)"`
	Comment4                  sql.NullString  `bun:"comment_4,type:varchar(30)"`
	Salesrep                  sql.NullString  `bun:"salesrep,type:char"`
	Buyer                     string          `bun:"buyer,type:char"`
	Schedular                 sql.NullString  `bun:"schedular,type:char"`
	LoginId                   sql.NullString  `bun:"login_id,type:varchar(30)"`
	Dealer                    sql.NullString  `bun:"dealer,type:char"`
	Url                       sql.NullString  `bun:"url,type:varchar(255)"`
	Comments                  sql.NullString  `bun:"comments,type:text(2147483647)"`
	HomeEmailAddress          sql.NullString  `bun:"home_email_address,type:varchar(255)"`
	HomeFax                   sql.NullString  `bun:"home_fax,type:char(20)"`
	Class1id                  sql.NullString  `bun:"class_1id,type:varchar(255)"`
	Class2id                  sql.NullString  `bun:"class_2id,type:varchar(255)"`
	Class3id                  sql.NullString  `bun:"class_3id,type:varchar(255)"`
	Class4id                  sql.NullString  `bun:"class_4id,type:varchar(255)"`
	Class5id                  sql.NullString  `bun:"class_5id,type:varchar(255)"`
	EmployeeVendorId          sql.NullFloat64 `bun:"employee_vendor_id,type:decimal(19,0)"`
	DearField                 sql.NullString  `bun:"dear_field,type:varchar(50)"`
	OldContactId              sql.NullString  `bun:"old_contact_id,type:varchar(20)"`
	Employee                  sql.NullString  `bun:"employee,type:char"`
	UpperCombinedName         sql.NullString  `bun:"upper_combined_name,type:varchar(50)"`
	DescendingCombinedName    sql.NullString  `bun:"descending_combined_name,type:varchar(50)"`
	DirectWattsNumber         sql.NullString  `bun:"direct_watts_number,type:varchar(20)"`
	NoOfCycleDays             sql.NullInt32   `bun:"no_of_cycle_days,type:int"`
	Mailstop                  sql.NullString  `bun:"mailstop,type:varchar(30)"`
	CommissionClassId         sql.NullString  `bun:"commission_class_id,type:varchar(8)"`
	CompanyId                 sql.NullString  `bun:"company_id,type:varchar(8)"`
	VendorId                  sql.NullFloat64 `bun:"vendor_id,type:decimal(19,0)"`
	AddressName               sql.NullString  `bun:"address_name,type:varchar(255)"`
	Driver                    string          `bun:"driver,type:char"`
	InsideSalesrepFlag        string          `bun:"inside_salesrep_flag,type:char"`
	DefaultBranchId           sql.NullString  `bun:"default_branch_id,type:varchar(8)"`
	Technician                string          `bun:"technician,type:char"`
	LocationId                sql.NullFloat64 `bun:"location_id,type:decimal(19,0)"`
	TerritoryUid              sql.NullInt32   `bun:"territory_uid,type:int"`
	FuelSurchargePercentage   sql.NullFloat64 `bun:"fuel_surcharge_percentage,type:decimal(19,4)"`
	MaxFuelChargePerShip      sql.NullFloat64 `bun:"max_fuel_charge_per_ship,type:decimal(19,4)"`
	ContactTypeId             sql.NullInt32   `bun:"contact_type_id,type:int"`
	SalesManagerId            sql.NullString  `bun:"sales_manager_id,type:varchar(16)"`
	RoadnetDriverId           sql.NullString  `bun:"roadnet_driver_id,type:varchar(255)"`
	CellularExt               sql.NullString  `bun:"cellular_ext,type:varchar(5)"`
	ContactRoleUid            sql.NullInt32   `bun:"contact_role_uid,type:int"`
	SfdcAccountId             sql.NullString  `bun:"sfdc_account_id,type:varchar(255)"`
	SfdcCreateDate            sql.NullTime    `bun:"sfdc_create_date,type:datetime"`
	SfdcUpdateDate            sql.NullTime    `bun:"sfdc_update_date,type:datetime"`
	SfdcContactId             sql.NullString  `bun:"sfdc_contact_id,type:varchar(255)"`
	SalesAgencyFlag           sql.NullString  `bun:"sales_agency_flag,type:char"`
	SalesAgencyName           sql.NullString  `bun:"sales_agency_name,type:varchar(50)"`
	DeliveryOutputCd          int32           `bun:"delivery_output_cd,type:int"`
	DriverEnableGpsFlag       string          `bun:"driver_enable_gps_flag,type:char"`
	DateDirectFaxLastModified sql.NullTime    `bun:"date_direct_fax_last_modified,type:datetime"`
	AdsUser                   sql.NullString  `bun:"ads_user,type:varchar(50)"`
	SalesrepDefaultLocationId sql.NullFloat64 `bun:"salesrep_default_location_id,type:decimal(19,0)"`
	DatagateSubjectId         sql.NullInt32   `bun:"datagate_subject_id,type:int"`
	RentalId                  sql.NullString  `bun:"rental_id,type:varchar(255)"`
	EmailOptOut               string          `bun:"email_opt_out,type:char"`
	DriverRfcNo               sql.NullString  `bun:"driver_rfc_no,type:varchar(255)"`
	DriverLicenseNo           sql.NullString  `bun:"driver_license_no,type:varchar(255)"`
	DriverRegistrationId      sql.NullString  `bun:"driver_registration_id,type:varchar(255)"`
}
