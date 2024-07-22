package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountHdr struct {
	bun.BaseModel                  `bun:"table:cycle_count_hdr"`
	CycleCountHdrUid               int32      `bun:"cycle_count_hdr_uid,type:int,pk"`                               // Unique identifier for a cycle_count_hdr record.
	CycleCountLocCriteriaUid       int32      `bun:"cycle_count_loc_criteria_uid,type:int"`                         // Unique id for the criteria used to generate the count.
	CycleCountNo                   int32      `bun:"cycle_count_no,type:int"`                                       // The cycle count number used to reference the count in the application.
	UnitOfMeasureCd                int32      `bun:"unit_of_measure_cd,type:int"`                                   // The unit of measure to show quantities on the report.
	ShowQtyAvailable               string     `bun:"show_qty_available,type:char(1)"`                               // Indicates whether to show the quantity available on the report.
	ShowQtyOnHand                  string     `bun:"show_qty_on_hand,type:char(1)"`                                 // Indicates whether to show the quantity on-hand on the report.
	ShowQtyAllocated               string     `bun:"show_qty_allocated,type:char(1)"`                               // Indicates whether to show the quantity allocated. on the report.
	DateCreated                    time.Time  `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified               time.Time  `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy               string     `bun:"last_maintained_by,type:varchar(30)"`                           // ID of the user who last maintained this record
	RowStatusFlag                  int32      `bun:"row_status_flag,type:int"`                                      // Indicates current record status.
	MonthlyCycleCountOnlyFlag      *string    `bun:"monthly_cycle_count_only_flag,type:char(1)"`                    // Indicates if only monthy items should be included in the cycle count report
	IncludeAllBinsFlag             *string    `bun:"include_all_bins_flag,type:char(1)"`                            // Custom flag to indicate including all bins when count by bins
	MaxNoItemsCounted              *int32     `bun:"max_no_items_counted,type:int"`                                 // Specifies number of items on the cycle count.
	MaxNoBinsCounted               *int32     `bun:"max_no_bins_counted,type:int"`                                  // Specifies number of bins on the cycle count.
	ShowSupplierPartNoFlag         *string    `bun:"show_supplier_part_no_flag,type:char(1)"`                       // Indicates whether to display the supplier part number on the cycle count.
	QtyAllocProcProdFlag           *string    `bun:"qty_alloc_proc_prod_flag,type:char(1)"`                         // Indicates whether to display how much is allocated to each Secondary Process or Production order..
	IncludeZeroQohItemFlag         *string    `bun:"include_zero_qoh_item_flag,type:char(1)"`                       // Indicates whether to include items with zero qoh in the cycle count..
	SortByLevel1                   *string    `bun:"sort_by_level1,type:varchar(255)"`                              // Column name for the first sort sequnce.
	SortModLevel1                  *string    `bun:"sort_mod_level1,type:char(1)"`                                  // Order of first sort sequence, D for Descending, A for Ascending.
	SortByLevel2                   *string    `bun:"sort_by_level2,type:varchar(255)"`                              // Column name for the second sort sequnce.
	SortModLevel2                  *string    `bun:"sort_mod_level2,type:char(1)"`                                  // Order of second sort sequence, D for Descending, A for Ascending.
	SortByLevel3                   *string    `bun:"sort_by_level3,type:varchar(255)"`                              // Column name for the third sort sequnce.
	SortModLevel3                  *string    `bun:"sort_mod_level3,type:char(1)"`                                  // Order of third sort sequence, D for Descending, A for Ascending.
	AbcClassList                   *string    `bun:"abc_class_list,type:varchar(8000)"`                             // A particular abc class list specified to produce a manual count list.
	PrimarySupplierList            *string    `bun:"primary_supplier_list,type:varchar(8000)"`                      // A particular primary supplier list specified to produce a manual count list.
	ItemList                       *string    `bun:"item_list,type:varchar(8000)"`                                  // A particular item list specified to produce a manual count list.
	ProductGroupList               *string    `bun:"product_group_list,type:varchar(8000)"`                         // A particular product group list specified to produce a manual count list.
	PutawayRankList                *string    `bun:"putaway_rank_list,type:varchar(8000)"`                          // A particular putaway rank list specified to produce a manual count list.
	PutawayAttributeList           *string    `bun:"putaway_attribute_list,type:varchar(8000)"`                     // A particular putaway attribute list specified to produce a manual count list.
	ItemClassList                  *string    `bun:"item_class_list,type:varchar(8000)"`                            // A particular item class list specified to produce a manual count list.
	BinSelectionOption             *string    `bun:"bin_selection_option,type:char(1)"`                             // Indicates whether the bin selection is a list(L) or a range(R).
	BinCdList                      *string    `bun:"bin_cd_list,type:varchar(8000)"`                                // A particular bin cd list specified to produce a manual count list. When bin_selection_option is set to L.
	BegBinCd                       *string    `bun:"beg_bin_cd,type:varchar(10)"`                                   // Begging bin cd for a particular bin range specified to produce a manual count list. When bin_selection_option is set to R.
	EndBinCd                       *string    `bun:"end_bin_cd,type:varchar(10)"`                                   // Ending bin cd for a particular bin range specified to produce a manual count list. When bin_selection_option is set to R.
	DateLastCountedBefore          *time.Time `bun:"date_last_counted_before,type:datetime"`                        // A particular date specified to produce a manual count list with only the items that have a date last counted before MM/DD/YYYY.
	AbcClassListExcludeFlag        string     `bun:"abc_class_list_exclude_flag,type:char(1),default:('N')"`        // Indicate where this cycle count excludes the ABC class list entered.
	PrimarySupplierListExcludeFlag string     `bun:"primary_supplier_list_exclude_flag,type:char(1),default:('N')"` // Indicate where this cycle count excludes the primary supplier list entered.
	ItemListExcludeFlag            string     `bun:"item_list_exclude_flag,type:char(1),default:('N')"`             // Indicate where this cycle count excludes the Item ID list entered.
	ProdGroupListExcludeFlag       string     `bun:"prod_group_list_exclude_flag,type:char(1),default:('N')"`       // Indicate where this cycle count excludes the product group ID list entered.
	PutawayRankListExcludeFlag     string     `bun:"putaway_rank_list_exclude_flag,type:char(1),default:('N')"`     // Indicate where this cycle count excludes the putaway rank list entered.
	PutawayAttribListExcludeFlag   string     `bun:"putaway_attrib_list_exclude_flag,type:char(1),default:('N')"`   // Indicate where this cycle count excludes the putaway attribute list entered.
	ItemClassListExcludeFlag       string     `bun:"item_class_list_exclude_flag,type:char(1),default:('N')"`       // Indicate where this cycle count excludes the item class list entered.
	BinCdListExcludeFlag           string     `bun:"bin_cd_list_exclude_flag,type:char(1),default:('N')"`           // Indicate where this cycle count excludes the bin list entered.
	DtlBinSortOrder                *string    `bun:"dtl_bin_sort_order,type:char(1)"`                               // Ascending(A) or Descending(D) order for detail bin appeared on the report.
}
