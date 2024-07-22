package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLinkEntityReq struct {
	bun.BaseModel               `bun:"table:document_link_entity_req"`
	DocumentLinkEntityReqUid    int32     `bun:"document_link_entity_req_uid,type:int,pk"`
	CompanyId                   *string   `bun:"company_id,type:varchar(8)"`
	EntityId                    int32     `bun:"entity_id,type:int"`
	EntityType                  int32     `bun:"entity_type,type:int"`
	RequireLotDocumentationFlag *string   `bun:"require_lot_documentation_flag,type:char(1)"`
	PrintFlag                   string    `bun:"print_flag,type:char(1),default:('N')"`
	FaxFlag                     string    `bun:"fax_flag,type:char(1),default:('N')"`
	EmailFlag                   string    `bun:"email_flag,type:char(1),default:('N')"`
	NumberOfCopies              int16     `bun:"number_of_copies,type:smallint,default:(0)"`
	RowStatusFlag               int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SendOutsideUseDocsFlag      string    `bun:"send_outside_use_docs_flag,type:char(1),default:('N')"`    // Send order/shipping document links created for
	SendOutsideUsePrintFlag     string    `bun:"send_outside_use_print_flag,type:char(1),default:('N')"`   // print outside use document links
	SendOutsideUseFaxFlag       string    `bun:"send_outside_use_fax_flag,type:char(1),default:('N')"`     // fax outside use document links
	SendOutsideUseEmailFlag     string    `bun:"send_outside_use_email_flag,type:char(1),default:('N')"`   // email outside use document links
	OutsideUseNumberOfCopies    int16     `bun:"outside_use_number_of_copies,type:smallint,default:((0))"` // number of copies to print for outside use document links
}
