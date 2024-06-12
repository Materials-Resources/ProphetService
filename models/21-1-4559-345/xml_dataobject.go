package model

type XmlDataobject struct {
	bun.BaseModel      `bun:"table:xml_dataobject"`
	XmlDataobjectUid   int32     `bun:"xml_dataobject_uid,type:int,pk"`
	DataobjectName     string    `bun:"dataobject_name,type:varchar(40)"`
	BusinessObjectName string    `bun:"business_object_name,type:varchar(32)"`
	SetupEventName     string    `bun:"setup_event_name,type:varchar(32)"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`
}
