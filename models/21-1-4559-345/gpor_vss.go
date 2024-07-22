package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type GporVss struct {
	bun.BaseModel                      `bun:"table:gpor_vss"`
	GporRunHdrUid                      int32      `bun:"gpor_run_hdr_uid,type:int,pk"`
	LocationId                         int32      `bun:"location_id,type:int,pk"`
	InvMastUid                         int32      `bun:"inv_mast_uid,type:int,pk"`
	CFuturereleasenonallocated         float64    `bun:"c_FutureReleaseNonallocated,type:decimal(19,4)"`
	CPendingbackorderamount            float64    `bun:"c_PendingBackorderAmount,type:decimal(19,4)"`
	CPendingbackorderamountprod        float64    `bun:"c_PendingBackorderAmountProd,type:decimal(19,4)"`
	CBlanketporeleaseqty               float64    `bun:"c_BlanketPOReleaseQty,type:decimal(19,4)"`
	CUnapprovedpoquantity              float64    `bun:"c_UnapprovedPOQuantity,type:decimal(19,4)"`
	CScheduleorderreleasesNonallocated float64    `bun:"c_ScheduleOrderReleases_nonallocated,type:decimal(19,4)"`
	CTbosum                            float64    `bun:"c_TBOSum,type:decimal(19,4)"`
	CProdorderqty                      float64    `bun:"c_ProdOrderQty,type:decimal(19,4)"`
	CExpeditedate                      *time.Time `bun:"c_ExpediteDate,type:datetime"`
	CComponentqty                      float64    `bun:"c_ComponentQty,type:decimal(19,4)"`
	CQtyavailableinotherrevisions      float64    `bun:"c_QtyAvailableInOtherRevisions,type:decimal(9,0),default:((0))"`  // Column to hold qty in other revisions that are not current.
	CSalesorderproductionorderqty      float64    `bun:"c_salesorderproductionorderqty,type:decimal(19,9),default:((0))"` // Open Quantity for Sales Orders Assemblies on 'P' Disposition
}
