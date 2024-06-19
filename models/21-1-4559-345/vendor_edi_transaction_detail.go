package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorEdiTransactionDetail struct {
	bun.BaseModel           `bun:"table:vendor_edi_transaction_detail"`
	VendorEdiTransDetailUid int32     `bun:"vendor_edi_trans_detail_uid,type:int,pk"`        // uid of the table, primary key
	VendorEdiTransactionUid int32     `bun:"vendor_edi_transaction_uid,type:int"`            // pointer to the master record that this attribute is for.
	Name                    string    `bun:"name,type:varchar(255)"`                         // generic column used to define an attribute for the given EDI transaction type
	Value                   string    `bun:"value,type:varchar(255)"`                        // generic column used to hold attribute value
	DataTypeCd              int32     `bun:"data_type_cd,type:int"`                          // attribute values data  type. For example int
	DataTypeLength          int32     `bun:"data_type_length,type:int"`                      // size of attribute values data type
	DataTypeScale           int32     `bun:"data_type_scale,type:int"`                       // scale of attribute value
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"` // date created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // datetime last modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // CommerCenter user id
}
