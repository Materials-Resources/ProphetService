package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Language struct {
	bun.BaseModel        `bun:"table:language"`
	LanguageId           string    `bun:"language_id,type:varchar(8),pk"`                            // What is the unique identifier for this language?
	LanguageDescription  string    `bun:"language_description,type:varchar(30)"`                     // Description of the language.
	DeleteFlag           string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	LanguageUid          int32     `bun:"language_uid,type:int,autoincrement,identity"`
	CrystalReportsFolder string    `bun:"crystal_reports_folder,type:char(255),nullzero"` // Folder containing Crystal reports for this language
	IsoCode              string    `bun:"iso_code,type:varchar(5),nullzero"`              // ISO code for the language, consisting of a language code (from ISO-639) and an optional two-letter country code (from ISO-3166)
}
