package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type QuoteHdr struct {
	bun.BaseModel    `bun:"table:quote_hdr"`
	QuoteHdrUid      int32      `bun:"quote_hdr_uid,type:int,pk"`                                     // What is the unique -  system generated identifier for this quote?
	OeHdrUid         int32      `bun:"oe_hdr_uid,type:int"`                                           // Internal unique value to indicate an order.
	CompleteFlag     string     `bun:"complete_flag,type:char(1),default:('N')"`                      // Identifies whether the Quote has been marked as Complete by the user.  Should only contain values of Y & N.
	DateCreated      time.Time  `bun:"date_created,type:datetime,default:(getdate())"`                // Indicates the date/time this record was created.
	DateLastModified time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy string     `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	Date843Sent      *time.Time `bun:"date_843_sent,type:datetime"`                                   // Date when last EDI 843 was sent
	ExpirationDate   *time.Time `bun:"expiration_date,type:datetime"`                                 // Date quote expires
}
