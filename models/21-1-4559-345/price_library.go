package model

type PriceLibrary struct {
	bun.BaseModel                 `bun:"table:price_library"`
	PriceLibraryUid               int32     `bun:"price_library_uid,type:int,pk"`
	PriceLibraryId                string    `bun:"price_library_id,type:varchar(20)"`
	Description                   string    `bun:"description,type:varchar(255)"`
	CategoryCd                    int16     `bun:"category_cd,type:smallint"`
	TypeCd                        int16     `bun:"type_cd,type:smallint"`
	SourcePriceCd                 int16     `bun:"source_price_cd,type:smallint,nullzero"`
	Multiplier                    float64   `bun:"multiplier,type:decimal(19,9),nullzero"`
	RowStatusFlag                 int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated                   time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(30)"`
	CreatedBy                     string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TerminalId                    float64   `bun:"terminal_id,type:decimal(19,0),nullzero"`
	CompanyId                     string    `bun:"company_id,type:varchar(255),nullzero"`
	LibraryOnContract             string    `bun:"library_on_contract,type:char,default:('N')"`
	StrategicPriceLibraryFlag     string    `bun:"strategic_price_library_flag,type:char,default:('N')"`
	CustomerSizeCd                int32     `bun:"customer_size_cd,type:int,nullzero"`
	CustomerCategoryUid           int32     `bun:"customer_category_uid,type:int,nullzero"`
	ExcludeFromCustMultiplierFlag string    `bun:"exclude_from_cust_multiplier_flag,type:char,nullzero"`
}
