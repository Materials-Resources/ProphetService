package model

type CreditcardTransresponse struct {
	bun.BaseModel              `bun:"table:creditcard_transresponse"`
	CreditcardTransresponseUid int32     `bun:"creditcard_transresponse_uid,type:int,pk"`
	ResponsetypeId             int32     `bun:"responsetype_id,type:int"`
	PaymentNumber              float64   `bun:"payment_number,type:decimal(19,0),nullzero"`
	RespfileLocation           string    `bun:"respfile_location,type:varchar(255),nullzero"`
	ResponseStatus             string    `bun:"response_status,type:varchar(10),nullzero"`
	ResponseDate               time.Time `bun:"response_date,type:datetime,nullzero"`
	PbResponseCode             string    `bun:"pb_response_code,type:varchar(50),nullzero"`
	PbResponseMessage          string    `bun:"pb_response_message,type:varchar(255),nullzero"`
	HostRefNumber              string    `bun:"host_ref_number,type:varchar(50),nullzero"`
	HostResponseCode           string    `bun:"host_response_code,type:varchar(50),nullzero"`
	HostResponseMessage        string    `bun:"host_response_message,type:varchar(255),nullzero"`
	AvsReturnCode              string    `bun:"avs_return_code,type:varchar(50),nullzero"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	CvvResponseCode            string    `bun:"cvv_response_code,type:varchar(50),nullzero"`
}
