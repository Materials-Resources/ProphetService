package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeContactsCustomer struct {
	bun.BaseModel            `bun:"table:oe_contacts_customer"`
	CompanyId                string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	CustomerId               float64   `bun:"customer_id,type:decimal(19,0),pk"`                         // Unique code to identify a customer.
	ContactId                string    `bun:"contact_id,type:varchar(16),pk"`                            // Unique code to identify the contact for this record.
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ShopperUid               int32     `bun:"shopper_uid,type:int,default:(0)"`                          // Unique Identifier from shopper table
	StatementContact         string    `bun:"statement_contact,type:char(1),default:('N')"`              // Determines whether the contact is the [primary contact] for the company / customer combination when sending Customer Statements (specifically for emailing and faxing).
	CreditCardNo             string    `bun:"credit_card_no,type:varchar(255),nullzero"`                 // Numeric code that appears on the customers credit card.
	CreditCardType           string    `bun:"credit_card_type,type:varchar(255),nullzero"`               // The type of credit card the customer is using.
	CreditCardName           string    `bun:"credit_card_name,type:varchar(255),nullzero"`               // The name of a person or organization, as it appears on the credit card.
	CreditCardExpirationDate time.Time `bun:"credit_card_expiration_date,type:datetime,nullzero"`        // The month and year a credit card expires.
	PedigreeContact          string    `bun:"pedigree_contact,type:char(1),default:('N')"`               // Serves as pedigree relationship point of contact
}
