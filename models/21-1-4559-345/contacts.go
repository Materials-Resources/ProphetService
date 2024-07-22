package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Contacts struct {
	bun.BaseModel             `bun:"table:contacts"`
	Id                        string     `bun:"id,type:varchar(16),pk"`                                          // This column is no longer being used and is scheduled for removal in later version.
	AddressId                 float64    `bun:"address_id,type:decimal(19,0)"`                                   // What is the address of this contact?
	Salutation                *string    `bun:"salutation,type:varchar(4)"`                                      // What is the salutation for this contact?
	FirstName                 string     `bun:"first_name,type:varchar(15)"`                                     // What is the first name of the contact?
	Mi                        *string    `bun:"mi,type:varchar(2)"`                                              // Middle Initial
	LastName                  string     `bun:"last_name,type:varchar(24)"`                                      // Last name of the contact.
	Title                     *string    `bun:"title,type:varchar(40)"`                                          // What is the title of this contact?
	DirectPhone               *string    `bun:"direct_phone,type:varchar(20)"`                                   // What is the direct telephone number for this conta
	PhoneExt                  *string    `bun:"phone_ext,type:varchar(5)"`                                       // What is the phone extension for this contact?
	DirectFax                 *string    `bun:"direct_fax,type:varchar(20)"`                                     // What is the direct fax number for this contact?
	FaxExt                    *string    `bun:"fax_ext,type:varchar(5)"`                                         // What is the fax extension for this contact?
	Beeper                    *string    `bun:"beeper,type:varchar(20)"`                                         // What is the beeper number for this contact?
	Cellular                  *string    `bun:"cellular,type:varchar(20)"`                                       // What is the mobile telephone number for this contact?
	EmailAddress              *string    `bun:"email_address,type:varchar(255)"`                                 // What is the email address for this contact?
	DateCreated               time.Time  `bun:"date_created,type:datetime"`                                      // Indicates the date/time this record was created.
	DateLastModified          time.Time  `bun:"date_last_modified,type:datetime"`                                // Indicates the date/time this record was last modified.
	LastMaintainedBy          string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`       // ID of the user who last maintained this record
	HomeAddress1              *string    `bun:"home_address1,type:varchar(40)"`                                  // What is the first line of the home address for this contact?
	HomeAddress2              *string    `bun:"home_address2,type:varchar(40)"`                                  // What is the second line of the home address for this contact?
	Birthday                  *time.Time `bun:"birthday,type:datetime"`                                          // What is the birthday of this contact?
	Anniversary               *time.Time `bun:"anniversary,type:datetime"`                                       // What is the anniversary date of this contact?
	HomePhone                 *string    `bun:"home_phone,type:varchar(20)"`                                     // Home phone number of contact.
	EmailAddress2             *string    `bun:"email_address2,type:varchar(255)"`                                // What is the Compuserve ID of this contact?
	DeleteFlag                string     `bun:"delete_flag,type:char(1)"`                                        // Indicates whether this record is logically deleted
	Comment1                  *string    `bun:"comment_1,type:varchar(30)"`                                      // This column is no longer being used and is scheduled for removal in later version.
	Comment2                  *string    `bun:"comment_2,type:varchar(30)"`                                      // This column is no longer being used and is scheduled for removal in later version.
	Comment3                  *string    `bun:"comment_3,type:varchar(30)"`                                      // This column is no longer being used and is scheduled for removal in later version.
	Comment4                  *string    `bun:"comment_4,type:varchar(30)"`                                      // This column is no longer being used and is scheduled for removal in later version.
	Salesrep                  *string    `bun:"salesrep,type:char(1)"`                                           // This column is no longer being used and is scheduled for removal in later version.
	Buyer                     string     `bun:"buyer,type:char(1),default:('N')"`                                // Indicates that the contact is a buyer.
	Schedular                 *string    `bun:"schedular,type:char(1)"`                                          // This column is no longer being used and is scheduled for removal in later version.
	LoginId                   *string    `bun:"login_id,type:varchar(30)"`                                       // Login ID for an external application. Display only.
	Dealer                    *string    `bun:"dealer,type:char(1)"`                                             // This column is no longer being used and is scheduled for removal in later version.
	Url                       *string    `bun:"url,type:varchar(255)"`                                           // The address of a contact web site on the Internet.  It is an acronym for Uniform Resource Locator.
	Comments                  *string    `bun:"comments,type:text"`                                              // User defined comments about the activity
	HomeEmailAddress          *string    `bun:"home_email_address,type:varchar(255)"`                            // What is the home Compuserve ID for this contact?
	HomeFax                   *string    `bun:"home_fax,type:char(20)"`                                          // What is the home fax number of this contact?
	Class1Id                  *string    `bun:"class_1id,type:varchar(255)"`                                     // What is the customer class?
	Class2Id                  *string    `bun:"class_2id,type:varchar(255)"`                                     // What is the class for this ship to vendor?
	Class3Id                  *string    `bun:"class_3id,type:varchar(255)"`                                     // What is the contac class?
	Class4Id                  *string    `bun:"class_4id,type:varchar(255)"`                                     // What is the class for this ship to vendor?
	Class5Id                  *string    `bun:"class_5id,type:varchar(255)"`                                     // What is the contac class?
	EmployeeVendorId          *float64   `bun:"employee_vendor_id,type:decimal(19,0)"`                           // This column is no longer being used and is scheduled for removal in later version.
	DearField                 *string    `bun:"dear_field,type:varchar(50)"`                                     // What salutation should be used for this contact?
	OldContactId              *string    `bun:"old_contact_id,type:varchar(20)"`                                 // This will be used when sending address info into W
	Employee                  *string    `bun:"employee,type:char(1)"`                                           // This column is no longer being used and is scheduled for removal in later version.
	UpperCombinedName         *string    `bun:"upper_combined_name,type:varchar(50)"`                            // This is similar to descending_combined_name.
	DescendingCombinedName    *string    `bun:"descending_combined_name,type:varchar(50)"`                       // This is the combination of the last name -  first name and middle initial columns
	DirectWattsNumber         *string    `bun:"direct_watts_number,type:varchar(20)"`                            // Direct Watts Phone Number
	NoOfCycleDays             *int32     `bun:"no_of_cycle_days,type:int"`                                       // Number of days to wait before a follow up with the
	Mailstop                  *string    `bun:"mailstop,type:varchar(30)"`                                       // An internal address for a specific employee/contact in a company.
	CommissionClassId         *string    `bun:"commission_class_id,type:varchar(8)"`                             // What is the unique identifier for this item commission class?
	CompanyId                 *string    `bun:"company_id,type:varchar(8)"`                                      // Unique code that identifies a company.
	VendorId                  *float64   `bun:"vendor_id,type:decimal(19,0)"`                                    // What is the unique vendor identifier for this row?
	AddressName               *string    `bun:"address_name,type:varchar(255)"`                                  // Indicates the address name corresponding to addres
	Driver                    string     `bun:"driver,type:char(1),default:('N')"`                               // Indicates that the contact is a driver. Used with PDA/delivery feature.
	InsideSalesrepFlag        string     `bun:"inside_salesrep_flag,type:char(1),default:('N')"`                 // Indicate whether the salesrep is an inside salesrep
	DefaultBranchId           *string    `bun:"default_branch_id,type:varchar(8)"`                               // Branch with which the contact is associated
	Technician                string     `bun:"technician,type:char(1),default:('N')"`                           // Set to Y if this is a service technician
	LocationId                *float64   `bun:"location_id,type:decimal(19,0)"`                                  // Service technician location id
	TerritoryUid              *int32     `bun:"territory_uid,type:int"`                                          // Custom feature 26039.  This column will hold the territory used for commission calcs.
	FuelSurchargePercentage   *float64   `bun:"fuel_surcharge_percentage,type:decimal(19,4)"`                    // A custom column which indicates the percentage that will be applied against a shipment when a shipment is confirmed.
	MaxFuelChargePerShip      *float64   `bun:"max_fuel_charge_per_ship,type:decimal(19,4)"`                     // A custom column which indicates the maximum fuel surcharge that will be applied against a shipment when a shipment is confirmed.
	ContactTypeId             *int32     `bun:"contact_type_id,type:int"`                                        // The contact type id from the contact type table
	SalesManagerId            *string    `bun:"sales_manager_id,type:varchar(16)"`                               // This column will contain the Sales Manger's ID that the Sales Rep reports to.
	RoadnetDriverId           *string    `bun:"roadnet_driver_id,type:varchar(255)"`                             // ID that UPS Roadnet has assigned to this driver.
	CellularExt               *string    `bun:"cellular_ext,type:varchar(5)"`                                    // Extension for cellular number.  Can be used with cellular field for alternate phone number.
	ContactRoleUid            *int32     `bun:"contact_role_uid,type:int"`                                       // Contact role for this contact
	SfdcAccountId             *string    `bun:"sfdc_account_id,type:varchar(255)"`                               // Salesforce.com account id - added by the import
	SfdcCreateDate            *time.Time `bun:"sfdc_create_date,type:datetime"`                                  // Date the record was created in Salesforce.com - added by import.
	SfdcUpdateDate            *time.Time `bun:"sfdc_update_date,type:datetime"`                                  // Date the record was updated in Salesforce.com - added by import.
	SfdcContactId             *string    `bun:"sfdc_contact_id,type:varchar(255)"`                               // Salesforce.com contact id - added by the import
	SalesAgencyFlag           *string    `bun:"sales_agency_flag,type:char(1)"`                                  // Custom column to indicate this is a sales agency
	SalesAgencyName           *string    `bun:"sales_agency_name,type:varchar(50)"`                              // Custom column to store sales agency name
	DeliveryOutputCd          int32      `bun:"delivery_output_cd,type:int,default:((2844))"`                    // Default delivery list download format code.
	DriverEnableGpsFlag       string     `bun:"driver_enable_gps_flag,type:char(1),default:('N')"`               // Flag for enabling GPS tracking of a drivers route
	DateDirectFaxLastModified *time.Time `bun:"date_direct_fax_last_modified,type:datetime,default:(getdate())"` // Customer (F45503): date the direct_fax column was last modified.  Populated via a trigger.
	AdsUser                   *string    `bun:"ads_user,type:varchar(50)"`                                       // ADS Subscriber
	SalesrepDefaultLocationId *float64   `bun:"salesrep_default_location_id,type:decimal(19,0)"`                 // Location used to determine gl postings when cost and revenue are split based on salesrep.
	DatagateSubjectId         *int32     `bun:"datagate_subject_id,type:int"`                                    // Custom (F68765): for the Datagate sales extract (export), used as the subject ID in the 107 and 109 sections
	RentalId                  *string    `bun:"rental_id,type:varchar(255)"`                                     // Id to link the contact with the Rental Essentials customer
	EmailOptOut               string     `bun:"email_opt_out,type:char(1),default:('N')"`                        // column for making a contact a nocontact for email blast
	DriverRfcNo               *string    `bun:"driver_rfc_no,type:varchar(255)"`                                 // MX: RFC number associated with associated driver.
	DriverLicenseNo           *string    `bun:"driver_license_no,type:varchar(255)"`                             // MX: License number associated with associated driver.
	DriverRegistrationId      *string    `bun:"driver_registration_id,type:varchar(255)"`                        // Tax id number of the driver if country of residence not MX.
}
