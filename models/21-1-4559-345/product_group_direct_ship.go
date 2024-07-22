package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ProductGroupDirectShip struct {
	bun.BaseModel             `bun:"table:product_group_direct_ship"`
	ProductGroupDirectShipUid int32     `bun:"product_group_direct_ship_uid,type:int,pk"`                    // Unique Identifier for the table
	CompanyId                 string    `bun:"company_id,type:varchar(8)"`                                   // Identifies the company
	ProductGroupId            string    `bun:"product_group_id,type:varchar(8)"`                             // Identifies the product group
	DirectShipStatus          string    `bun:"direct_ship_status,type:char(1),default:('S')"`                // Contains the direct ship status. The only valid values are A(always), S(system) and Q(quantity).
	DirectShipQty             float64   `bun:"direct_ship_qty,type:decimal(19,9),default:(0)"`               // When direct_ship_status is Q, this will hold the cutoff quantity for a direct shipment line.
	DeleteFlag                string    `bun:"delete_flag,type:char(1),default:('N')"`                       // Indicates if the row is deleted or not.
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	OrderValue                float64   `bun:"order_value,type:decimal(19,9),default:((0))"`                 // When direct_ship_status is V value, this holds the cutoff order value for a direct ship line.
}
