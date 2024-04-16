package prophet_21_1_4559

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Contacts struct {
	bun.BaseModel             `bun:"table:contacts"`
	Id                        string          `bun:"id,type:varchar(16),pk"`
	AddressId                 float64         `bun:"address_id,type:decimal(19,0)"`
	Salutation                sql.NullString  `bun:"salutation,type:varchar(4),nullzero"`
	FirstName                 string          `bun:"first_name,type:varchar(15)"`
	Mi                        sql.NullString  `bun:"mi,type:varchar(2),nullzero"`
	LastName                  string          `bun:"last_name,type:varchar(24)"`
	Title                     sql.NullString  `bun:"title,type:varchar(40),nullzero"`
	DirectPhone               sql.NullString  `bun:"direct_phone,type:varchar(20),nullzero"`
	PhoneExt                  sql.NullString  `bun:"phone_ext,type:varchar(5),nullzero"`
	DirectFax                 sql.NullString  `bun:"direct_fax,type:varchar(20),nullzero"`
	FaxExt                    sql.NullString  `bun:"fax_ext,type:varchar(5),nullzero"`
	Beeper                    sql.NullString  `bun:"beeper,type:varchar(20),nullzero"`
	Cellular                  sql.NullString  `bun:"cellular,type:varchar(20),nullzero"`
	EmailAddress              sql.NullString  `bun:"email_address,type:varchar(255),nullzero"`
	DateCreated               time.Time       `bun:"date_created,type:datetime"`
	DateLastModified          time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string          `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	HomeAddress1              sql.NullString  `bun:"home_address1,type:varchar(40),nullzero"`
	HomeAddress2              sql.NullString  `bun:"home_address2,type:varchar(40),nullzero"`
	Birthday                  sql.NullTime    `bun:"birthday,type:datetime,nullzero"`
	Anniversary               sql.NullTime    `bun:"anniversary,type:datetime,nullzero"`
	HomePhone                 sql.NullString  `bun:"home_phone,type:varchar(20),nullzero"`
	EmailAddress2             sql.NullString  `bun:"email_address2,type:varchar(255),nullzero"`
	DeleteFlag                string          `bun:"delete_flag,type:char"`
	Comment1                  sql.NullString  `bun:"comment_1,type:varchar(30),nullzero"`
	Comment2                  sql.NullString  `bun:"comment_2,type:varchar(30),nullzero"`
	Comment3                  sql.NullString  `bun:"comment_3,type:varchar(30),nullzero"`
	Comment4                  sql.NullString  `bun:"comment_4,type:varchar(30),nullzero"`
	Salesrep                  sql.NullString  `bun:"salesrep,type:char,nullzero"`
	Buyer                     string          `bun:"buyer,type:char,default:('N')"`
	Schedular                 sql.NullString  `bun:"schedular,type:char,nullzero"`
	LoginId                   sql.NullString  `bun:"login_id,type:varchar(30),nullzero"`
	Dealer                    sql.NullString  `bun:"dealer,type:char,nullzero"`
	Url                       sql.NullString  `bun:"url,type:varchar(255),nullzero"`
	Comments                  sql.NullString  `bun:"comments,type:text(2147483647),nullzero"`
	HomeEmailAddress          sql.NullString  `bun:"home_email_address,type:varchar(255),nullzero"`
	HomeFax                   sql.NullString  `bun:"home_fax,type:char(20),nullzero"`
	Class1id                  sql.NullString  `bun:"class_1id,type:varchar(255),nullzero"`
	Class2id                  sql.NullString  `bun:"class_2id,type:varchar(255),nullzero"`
	Class3id                  sql.NullString  `bun:"class_3id,type:varchar(255),nullzero"`
	Class4id                  sql.NullString  `bun:"class_4id,type:varchar(255),nullzero"`
	Class5id                  sql.NullString  `bun:"class_5id,type:varchar(255),nullzero"`
	EmployeeVendorId          sql.NullFloat64 `bun:"employee_vendor_id,type:decimal(19,0),nullzero"`
	DearField                 sql.NullString  `bun:"dear_field,type:varchar(50),nullzero"`
	OldContactId              sql.NullString  `bun:"old_contact_id,type:varchar(20),nullzero"`
	Employee                  sql.NullString  `bun:"employee,type:char,nullzero"`
	UpperCombinedName         sql.NullString  `bun:"upper_combined_name,type:varchar(50),nullzero"`
	DescendingCombinedName    sql.NullString  `bun:"descending_combined_name,type:varchar(50),nullzero"`
	DirectWattsNumber         sql.NullString  `bun:"direct_watts_number,type:varchar(20),nullzero"`
	NoOfCycleDays             sql.NullInt32   `bun:"no_of_cycle_days,type:int,nullzero"`
	Mailstop                  sql.NullString  `bun:"mailstop,type:varchar(30),nullzero"`
	CommissionClassId         sql.NullString  `bun:"commission_class_id,type:varchar(8),nullzero"`
	CompanyId                 sql.NullString  `bun:"company_id,type:varchar(8),nullzero"`
	VendorId                  sql.NullFloat64 `bun:"vendor_id,type:decimal(19,0),nullzero"`
	AddressName               sql.NullString  `bun:"address_name,type:varchar(255),nullzero"`
	Driver                    string          `bun:"driver,type:char,default:('N')"`
	InsideSalesrepFlag        string          `bun:"inside_salesrep_flag,type:char,default:('N')"`
	DefaultBranchId           sql.NullString  `bun:"default_branch_id,type:varchar(8),nullzero"`
	Technician                string          `bun:"technician,type:char,default:('N')"`
	LocationId                sql.NullFloat64 `bun:"location_id,type:decimal(19,0),nullzero"`
	TerritoryUid              sql.NullInt32   `bun:"territory_uid,type:int,nullzero"`
	FuelSurchargePercentage   sql.NullFloat64 `bun:"fuel_surcharge_percentage,type:decimal(19,4),nullzero"`
	MaxFuelChargePerShip      sql.NullFloat64 `bun:"max_fuel_charge_per_ship,type:decimal(19,4),nullzero"`
	ContactTypeId             sql.NullInt32   `bun:"contact_type_id,type:int,nullzero"`
	SalesManagerId            sql.NullString  `bun:"sales_manager_id,type:varchar(16),nullzero"`
	RoadnetDriverId           sql.NullString  `bun:"roadnet_driver_id,type:varchar(255),nullzero"`
	CellularExt               sql.NullString  `bun:"cellular_ext,type:varchar(5),nullzero"`
	ContactRoleUid            sql.NullInt32   `bun:"contact_role_uid,type:int,nullzero"`
	SfdcAccountId             sql.NullString  `bun:"sfdc_account_id,type:varchar(255),nullzero"`
	SfdcCreateDate            sql.NullTime    `bun:"sfdc_create_date,type:datetime,nullzero"`
	SfdcUpdateDate            sql.NullTime    `bun:"sfdc_update_date,type:datetime,nullzero"`
	SfdcContactId             sql.NullString  `bun:"sfdc_contact_id,type:varchar(255),nullzero"`
	SalesAgencyFlag           sql.NullString  `bun:"sales_agency_flag,type:char,nullzero"`
	SalesAgencyName           sql.NullString  `bun:"sales_agency_name,type:varchar(50),nullzero"`
	DeliveryOutputCd          int32           `bun:"delivery_output_cd,type:int,default:((2844))"`
	DriverEnableGpsFlag       string          `bun:"driver_enable_gps_flag,type:char,default:('N')"`
	DateDirectFaxLastModified sql.NullTime    `bun:"date_direct_fax_last_modified,type:datetime,default:(getdate())"`
	AdsUser                   sql.NullString  `bun:"ads_user,type:varchar(50),nullzero"`
	SalesrepDefaultLocationId sql.NullFloat64 `bun:"salesrep_default_location_id,type:decimal(19,0),nullzero"`
	DatagateSubjectId         sql.NullInt32   `bun:"datagate_subject_id,type:int,nullzero"`
	RentalId                  sql.NullString  `bun:"rental_id,type:varchar(255),nullzero"`
	EmailOptOut               string          `bun:"email_opt_out,type:char,default:('N')"`
	DriverRfcNo               sql.NullString  `bun:"driver_rfc_no,type:varchar(255),nullzero"`
	DriverLicenseNo           sql.NullString  `bun:"driver_license_no,type:varchar(255),nullzero"`
	DriverRegistrationId      sql.NullString  `bun:"driver_registration_id,type:varchar(255),nullzero"`
}
