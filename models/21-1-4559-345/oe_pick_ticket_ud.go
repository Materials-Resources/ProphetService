package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OePickTicketUd struct {
	bun.BaseModel     `bun:"table:oe_pick_ticket_ud"`
	OePickTicketUdUid int32     `bun:"oe_pick_ticket_ud_uid,type:int,autoincrement,scanonly,pk"`
	PickTicketNo      float64   `bun:"pick_ticket_no,type:decimal(19,0),unique"`
	Picker            string    `bun:"picker,type:varchar(40),nullzero"` // Picker
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
