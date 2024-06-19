package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CrmRun struct {
	bun.BaseModel              `bun:"table:crm_run"`
	RunNumber                  int32     `bun:"run_number,type:int,pk"`                                       // Unique number of the call to the procedure p21_crm_master_inquiry.
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`                                // Customers company
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),pk"`                            // ID of the customer
	CustomerName               string    `bun:"customer_name,type:varchar(255),nullzero"`                     // customers name
	Last30DaysSales            float64   `bun:"last_30_days_sales,type:decimal(19,2),nullzero"`               // Sales amount invoiced in the last 30 days.
	Last60DaysSales            float64   `bun:"last_60_days_sales,type:decimal(19,2),nullzero"`               // Sales amount invoiced in the last 60 days.
	Last90DaysSales            float64   `bun:"last_90_days_sales,type:decimal(19,2),nullzero"`               // Sales amount invoiced in the last 90 days.
	Last365DaysSales           float64   `bun:"last_365_days_sales,type:decimal(19,2),nullzero"`              // Sales amount invoiced in the last 365 days.
	Last365DaysCommCost        float64   `bun:"last_365_days_comm_cost,type:decimal(19,9),nullzero"`          // Commission cost invoiced in the last 365 days.
	Last365DaysOtherCost       float64   `bun:"last_365_days_other_cost,type:decimal(19,9),nullzero"`         // Other cost invoiced in the last 365 days.
	Last365DaysFreightBilled   float64   `bun:"last_365_days_freight_billed,type:decimal(19,4),nullzero"`     // freigt billed on invoices in the last 365 days.
	Last365DaysFreightUnbilled float64   `bun:"last_365_days_freight_unbilled,type:decimal(19,4),nullzero"`   // Freight that wasnt billed on invoices in the last 365 days.
	PreviousYearSales          float64   `bun:"previous_year_sales,type:decimal(19,2),nullzero"`              // Sales amount invoiced in the previous year.
	DateAcctOpened             time.Time `bun:"date_acct_opened,type:datetime,nullzero"`                      // Date the customer account became active.
	LastHardTouchDate          time.Time `bun:"last_hard_touch_date,type:datetime,nullzero"`                  // The last hard-touch date for this customer.
	OpenOrderValue             float64   `bun:"open_order_value,type:decimal(19,9),nullzero"`                 // Total price of open order lines.
	OpenQuoteValue             float64   `bun:"open_quote_value,type:decimal(19,9),nullzero"`                 // Total price of open quote lines.
	OpenOpportunityValue       float64   `bun:"open_opportunity_value,type:decimal(19,9),nullzero"`           // Total size of open opportunities.
	NumberOfOpportunities      int32     `bun:"number_of_opportunities,type:int,nullzero"`                    // Number of open opportunities.
	AverageDso                 float64   `bun:"average_dso,type:decimal(19,9),nullzero"`                      // Average days sales outstanding.
	YtdPrctAmtWon              float64   `bun:"ytd_prct_amt_won,type:decimal(19,9),nullzero"`                 // Percent of quoted sales converted over the past year.
	YtdPrctLinesWon            float64   `bun:"ytd_prct_lines_won,type:decimal(19,9),nullzero"`               // Percent of quoted lines converted over the past year.
	RmaValueLast365            float64   `bun:"rma_value_last_365,type:decimal(19,9),nullzero"`               // Sales on open and completed rmas over the last 365 days.
	OrderValueLast365          float64   `bun:"order_value_last_365,type:decimal(19,9),nullzero"`             // Sales on open and completed orders over the last 365 days.
	Last365DaysCogs            float64   `bun:"last_365_days_cogs,type:decimal(19,9),nullzero"`               // COGS amount invoiced in the last 365 days.
	DaysSinceLastOrder         int32     `bun:"days_since_last_order,type:int,nullzero"`                      // Number of days since the customer last placed an order.
	DaysSinceFirstOrder        int32     `bun:"days_since_first_order,type:int,nullzero"`                     // Number of days since the customer entered their first order.
	NumberOfOrders             int32     `bun:"number_of_orders,type:int,nullzero"`                           // Number of orders for this customer.
	DaysSinceLastQuote         int32     `bun:"days_since_last_quote,type:int,nullzero"`                      // Number of days since the customer last placed a quote.
	DaysSinceFirstQuote        int32     `bun:"days_since_first_quote,type:int,nullzero"`                     // Number of days since the customer entered their first quote.
	NumberOfQuotes             int32     `bun:"number_of_quotes,type:int,nullzero"`                           // Number of quotes for this customer
	SicCode                    float64   `bun:"sic_code,type:decimal(6,0),nullzero"`                          // Unique SIC code.
	SicDescription             string    `bun:"sic_description,type:varchar(40),nullzero"`                    // Description of the SIC code.
	LeadSourceId               string    `bun:"lead_source_id,type:varchar(8),nullzero"`                      // Default lead source id.
	SourceDescription          string    `bun:"source_description,type:varchar(30),nullzero"`                 // Description of lead source.
	Class1Id                   string    `bun:"class_1id,type:varchar(255),nullzero"`                         // A user-defined code that identifies a group of customers.
	Class2Id                   string    `bun:"class_2id,type:varchar(255),nullzero"`                         // A user-defined code that identifies a group of customers.
	Class3Id                   string    `bun:"class_3id,type:varchar(255),nullzero"`                         // A user-defined code that identifies a group of customers.
	Class4Id                   string    `bun:"class_4id,type:varchar(255),nullzero"`                         // A user-defined code that identifies a group of customers.
	Class5Id                   string    `bun:"class_5id,type:varchar(255),nullzero"`                         // A user-defined code that identifies a group of customers.
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16),nullzero"`                        // Customers default salesrep.
	PhysAddress1               string    `bun:"phys_address1,type:varchar(50),nullzero"`                      // Customers address.
	PhysAddress2               string    `bun:"phys_address2,type:varchar(50),nullzero"`                      // Customers address.
	PhysCity                   string    `bun:"phys_city,type:varchar(50),nullzero"`                          // Customers city.
	PhysState                  string    `bun:"phys_state,type:varchar(50),nullzero"`                         // Customers state.
	PhysPostalCode             string    `bun:"phys_postal_code,type:varchar(10),nullzero"`                   // Customers postal code.
	PhysCountry                string    `bun:"phys_country,type:varchar(50),nullzero"`                       // Customers country
	CentralPhoneNumber         string    `bun:"central_phone_number,type:varchar(20),nullzero"`               // Customers phone number.
	NewCustomerFlag            string    `bun:"new_customer_flag,type:char(1),nullzero"`                      // Is this a new customer?
	SalesrepName               string    `bun:"salesrep_name,type:varchar(255),nullzero"`                     // The name of the customers default salesrep.
	InRewardsProgramFlag       string    `bun:"in_rewards_program_flag,type:char(1),nullzero"`                // Is the customer in a rewards program.
	NextTaskToComplete         string    `bun:"next_task_to_complete,type:varchar(16),nullzero"`              // ID of the next task to be completed.
	LastCompletedTask          string    `bun:"last_completed_task,type:varchar(16),nullzero"`                // ID of the last task that was completed.
	CustomerCategoryUid        int32     `bun:"customer_category_uid,type:int,nullzero"`                      // Customers category.
	WarehouseSizeCd            int32     `bun:"warehouse_size_cd,type:int,nullzero"`                          // Customers warehouse size.
	RetailSizeCd               int32     `bun:"retail_size_cd,type:int,nullzero"`                             // Customers retail size.
	Roa0To30                   float64   `bun:"roa_0_to_30,type:decimal(19,4),nullzero"`                      // Ratio of Attainment 0 to 30 day bucket.
	Roa31To60                  float64   `bun:"roa_31_to_60,type:decimal(19,4),nullzero"`                     // Ratio of Attainment 31 to 60 day bucket.
	Roa61To90                  float64   `bun:"roa_61_to_90,type:decimal(19,4),nullzero"`                     // Ratio of Attainment 61 to 90 day bucket.
	Gap0To30                   float64   `bun:"gap_0_to_30,type:decimal(19,4),nullzero"`                      // GAP 0 to 30 day bucket.
	Gap31To60                  float64   `bun:"gap_31_to_60,type:decimal(19,4),nullzero"`                     // GAP 31 to 60 day bucket.
	Gap61To90                  float64   `bun:"gap_61_to_90,type:decimal(19,4),nullzero"`                     // GAP 61 to 90 day bucket.
	MissedBuyFlag              string    `bun:"missed_buy_flag,type:char(1),nullzero"`                        // Does this customer have a missed buy?
	CreditStatus               string    `bun:"credit_status,type:char(8),nullzero"`                          // Customers credit status.
	AverageOrderSize           float64   `bun:"average_order_size,type:decimal(19,4),nullzero"`               // Average size of the customers orders.
	AvgDaysBetweenOrders       int32     `bun:"avg_days_between_orders,type:int,nullzero"`                    // The average days between the orders entered in the last 365 days
	DateLastInvoiced           time.Time `bun:"date_last_invoiced,type:datetime,nullzero"`                    // Date the customer was last invoiced
	TotalOrders                int32     `bun:"total_orders,type:int,nullzero"`                               // Total number of orders for a customer
	TotalQuotes                int32     `bun:"total_quotes,type:int,nullzero"`                               // Total number of quotes for a customer
	PhysAddress3               string    `bun:"phys_address3,type:varchar(50),nullzero"`                      // Physical address line 3
	CostToCarryLateInvoices    float64   `bun:"cost_to_carry_late_invoices,type:decimal(19,2),default:((0))"` // The annualized cost of borrowing to cover late invoices
}
