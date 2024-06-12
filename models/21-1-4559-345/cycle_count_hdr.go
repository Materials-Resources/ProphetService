package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CycleCountHdr struct {
	bun.BaseModel                  `bun:"table:cycle_count_hdr"`
	CycleCountHdrUid               int32     `bun:"cycle_count_hdr_uid,type:int,pk"`
	CycleCountLocCriteriaUid       int32     `bun:"cycle_count_loc_criteria_uid,type:int"`
	CycleCountNo                   int32     `bun:"cycle_count_no,type:int"`
	UnitOfMeasureCd                int32     `bun:"unit_of_measure_cd,type:int"`
	ShowQtyAvailable               string    `bun:"show_qty_available,type:char"`
	ShowQtyOnHand                  string    `bun:"show_qty_on_hand,type:char"`
	ShowQtyAllocated               string    `bun:"show_qty_allocated,type:char"`
	DateCreated                    time.Time `bun:"date_created,type:datetime"`
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag                  int32     `bun:"row_status_flag,type:int"`
	MonthlyCycleCountOnlyFlag      string    `bun:"monthly_cycle_count_only_flag,type:char,nullzero"`
	IncludeAllBinsFlag             string    `bun:"include_all_bins_flag,type:char,nullzero"`
	MaxNoItemsCounted              int32     `bun:"max_no_items_counted,type:int,nullzero"`
	MaxNoBinsCounted               int32     `bun:"max_no_bins_counted,type:int,nullzero"`
	ShowSupplierPartNoFlag         string    `bun:"show_supplier_part_no_flag,type:char,nullzero"`
	QtyAllocProcProdFlag           string    `bun:"qty_alloc_proc_prod_flag,type:char,nullzero"`
	IncludeZeroQohItemFlag         string    `bun:"include_zero_qoh_item_flag,type:char,nullzero"`
	SortByLevel1                   string    `bun:"sort_by_level1,type:varchar(255),nullzero"`
	SortModLevel1                  string    `bun:"sort_mod_level1,type:char,nullzero"`
	SortByLevel2                   string    `bun:"sort_by_level2,type:varchar(255),nullzero"`
	SortModLevel2                  string    `bun:"sort_mod_level2,type:char,nullzero"`
	SortByLevel3                   string    `bun:"sort_by_level3,type:varchar(255),nullzero"`
	SortModLevel3                  string    `bun:"sort_mod_level3,type:char,nullzero"`
	AbcClassList                   string    `bun:"abc_class_list,type:varchar(8000),nullzero"`
	PrimarySupplierList            string    `bun:"primary_supplier_list,type:varchar(8000),nullzero"`
	ItemList                       string    `bun:"item_list,type:varchar(8000),nullzero"`
	ProductGroupList               string    `bun:"product_group_list,type:varchar(8000),nullzero"`
	PutawayRankList                string    `bun:"putaway_rank_list,type:varchar(8000),nullzero"`
	PutawayAttributeList           string    `bun:"putaway_attribute_list,type:varchar(8000),nullzero"`
	ItemClassList                  string    `bun:"item_class_list,type:varchar(8000),nullzero"`
	BinSelectionOption             string    `bun:"bin_selection_option,type:char,nullzero"`
	BinCdList                      string    `bun:"bin_cd_list,type:varchar(8000),nullzero"`
	BegBinCd                       string    `bun:"beg_bin_cd,type:varchar(10),nullzero"`
	EndBinCd                       string    `bun:"end_bin_cd,type:varchar(10),nullzero"`
	DateLastCountedBefore          time.Time `bun:"date_last_counted_before,type:datetime,nullzero"`
	AbcClassListExcludeFlag        string    `bun:"abc_class_list_exclude_flag,type:char,default:('N')"`
	PrimarySupplierListExcludeFlag string    `bun:"primary_supplier_list_exclude_flag,type:char,default:('N')"`
	ItemListExcludeFlag            string    `bun:"item_list_exclude_flag,type:char,default:('N')"`
	ProdGroupListExcludeFlag       string    `bun:"prod_group_list_exclude_flag,type:char,default:('N')"`
	PutawayRankListExcludeFlag     string    `bun:"putaway_rank_list_exclude_flag,type:char,default:('N')"`
	PutawayAttribListExcludeFlag   string    `bun:"putaway_attrib_list_exclude_flag,type:char,default:('N')"`
	ItemClassListExcludeFlag       string    `bun:"item_class_list_exclude_flag,type:char,default:('N')"`
	BinCdListExcludeFlag           string    `bun:"bin_cd_list_exclude_flag,type:char,default:('N')"`
	DtlBinSortOrder                string    `bun:"dtl_bin_sort_order,type:char,nullzero"`
}
