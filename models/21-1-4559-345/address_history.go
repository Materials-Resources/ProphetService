package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AddressHistory struct {
	bun.BaseModel     `bun:"table:address_history"`
	AddressHistoryUid int32     `bun:"address_history_uid,type:int,pk"`                               // Unique -  internal identifier for address_history information.
	CustomerName      string    `bun:"customer_name,type:varchar(50),nullzero"`                       // Indicates the customer name corresponding to the c
	ContactName       string    `bun:"contact_name,type:varchar(30),nullzero"`                        // What is the name of the contact -  at the time this address history row was written?
	Address1          string    `bun:"address1,type:varchar(50),nullzero"`                            // What is the first address line -  at the time this address history row was written?
	Address2          string    `bun:"address2,type:varchar(50),nullzero"`                            // This column is unused.
	State             string    `bun:"state,type:varchar(50),nullzero"`                               // What is the state -  at the time this address history row was written?
	PostalCode        string    `bun:"postal_code,type:char(10),nullzero"`                            // What is the postal  code -  at the time this address history row was written?
	Country           string    `bun:"country,type:varchar(50),nullzero"`                             // What is the country -  at the time this address history row was written?
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	City              string    `bun:"city,type:varchar(50),nullzero"`                                // What is the city -  at the time the address history row is created?
	Address3          string    `bun:"address3,type:varchar(50),nullzero"`                            // Address line 3
}
