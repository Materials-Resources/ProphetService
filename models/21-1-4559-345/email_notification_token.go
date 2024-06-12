package model

import (
	"github.com/uptrace/bun"
	"time"
)

type EmailNotificationToken struct {
	bun.BaseModel             `bun:"table:email_notification_token"`
	EmailNotificationTokenUid int32     `bun:"email_notification_token_uid,type:int,identity"`
	TokenName                 string    `bun:"token_name,type:varchar(255)"`
	TokenDescription          string    `bun:"token_description,type:varchar(255)"`
	TokenArea                 int32     `bun:"token_area,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DataTypeCd                int32     `bun:"data_type_cd,type:int,default:((1255))"`
}
