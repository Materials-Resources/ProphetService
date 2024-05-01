# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [catalog/v1/catalog.proto](#catalog_v1_catalog-proto)
    - [CreateProductGroupRequest](#catalog-v1-CreateProductGroupRequest)
    - [CreateProductGroupResponse](#catalog-v1-CreateProductGroupResponse)
    - [CreateProductRequest](#catalog-v1-CreateProductRequest)
    - [CreateProductResponse](#catalog-v1-CreateProductResponse)
    - [CreateSupplierRequest](#catalog-v1-CreateSupplierRequest)
    - [CreateSupplierResponse](#catalog-v1-CreateSupplierResponse)
    - [Cursor](#catalog-v1-Cursor)
    - [DeleteProductRequest](#catalog-v1-DeleteProductRequest)
    - [DeleteProductResponse](#catalog-v1-DeleteProductResponse)
    - [GetBasicProductDetailsRequest](#catalog-v1-GetBasicProductDetailsRequest)
    - [GetBasicProductDetailsResponse](#catalog-v1-GetBasicProductDetailsResponse)
    - [GetBasicProductDetailsResponse.BasicProductDetail](#catalog-v1-GetBasicProductDetailsResponse-BasicProductDetail)
    - [GetProductBySupplierRequest](#catalog-v1-GetProductBySupplierRequest)
    - [GetProductBySupplierResponse](#catalog-v1-GetProductBySupplierResponse)
    - [GetProductGroupRequest](#catalog-v1-GetProductGroupRequest)
    - [GetProductGroupResponse](#catalog-v1-GetProductGroupResponse)
    - [GetProductPriceRequest](#catalog-v1-GetProductPriceRequest)
    - [GetProductPriceResponse](#catalog-v1-GetProductPriceResponse)
    - [GetProductPriceResponse.ProductPrice](#catalog-v1-GetProductPriceResponse-ProductPrice)
    - [GetProductRequest](#catalog-v1-GetProductRequest)
    - [GetProductResponse](#catalog-v1-GetProductResponse)
    - [GetSupplierRequest](#catalog-v1-GetSupplierRequest)
    - [GetSupplierResponse](#catalog-v1-GetSupplierResponse)
    - [ListProductGroupsRequest](#catalog-v1-ListProductGroupsRequest)
    - [ListProductGroupsResponse](#catalog-v1-ListProductGroupsResponse)
    - [ListProductsRequest](#catalog-v1-ListProductsRequest)
    - [ListProductsResponse](#catalog-v1-ListProductsResponse)
    - [ListSuppliersRequest](#catalog-v1-ListSuppliersRequest)
    - [ListSuppliersResponse](#catalog-v1-ListSuppliersResponse)
    - [ProductDetail](#catalog-v1-ProductDetail)
    - [ProductFilter](#catalog-v1-ProductFilter)
    - [ProductGroup](#catalog-v1-ProductGroup)
    - [ProductSupplier](#catalog-v1-ProductSupplier)
    - [SetPrimarySupplierRequest](#catalog-v1-SetPrimarySupplierRequest)
    - [SetPrimarySupplierResponse](#catalog-v1-SetPrimarySupplierResponse)
    - [UpdateProductGroupRequest](#catalog-v1-UpdateProductGroupRequest)
    - [UpdateProductGroupResponse](#catalog-v1-UpdateProductGroupResponse)
    - [UpdateProductRequest](#catalog-v1-UpdateProductRequest)
    - [UpdateProductResponse](#catalog-v1-UpdateProductResponse)
    - [UpdateSupplierRequest](#catalog-v1-UpdateSupplierRequest)
    - [UpdateSupplierResponse](#catalog-v1-UpdateSupplierResponse)
    - [ValidationError](#catalog-v1-ValidationError)
  
    - [CatalogService](#catalog-v1-CatalogService)
  
- [customer/v1/customer.proto](#customer_v1_customer-proto)
    - [CustomerBranch](#customer-v1-CustomerBranch)
    - [GetContactByIdRequest](#customer-v1-GetContactByIdRequest)
    - [GetContactByIdResponse](#customer-v1-GetContactByIdResponse)
    - [GetContactByIdResponse.Contact](#customer-v1-GetContactByIdResponse-Contact)
    - [GetCustomerRequest](#customer-v1-GetCustomerRequest)
    - [GetCustomerResponse](#customer-v1-GetCustomerResponse)
    - [GetOrdersRequest](#customer-v1-GetOrdersRequest)
    - [GetOrdersResponse](#customer-v1-GetOrdersResponse)
    - [GetOrdersResponse.Order](#customer-v1-GetOrdersResponse-Order)
  
    - [CustomerService](#customer-v1-CustomerService)
  
- [inventory/v1/inventory.proto](#inventory_v1_inventory-proto)
    - [AddSupplierRequest](#inventory-v1-AddSupplierRequest)
    - [AddSupplierResponse](#inventory-v1-AddSupplierResponse)
    - [DeleteSupplierRequest](#inventory-v1-DeleteSupplierRequest)
    - [DeleteSupplierResponse](#inventory-v1-DeleteSupplierResponse)
    - [GetProductStockRequest](#inventory-v1-GetProductStockRequest)
    - [GetProductStockResponse](#inventory-v1-GetProductStockResponse)
    - [GetProductStockResponse.ProductStock](#inventory-v1-GetProductStockResponse-ProductStock)
    - [GetReceiptByIDRequest](#inventory-v1-GetReceiptByIDRequest)
    - [GetReceiptByIDResponse](#inventory-v1-GetReceiptByIDResponse)
    - [GetReceiptByIDResponse.Item](#inventory-v1-GetReceiptByIDResponse-Item)
    - [GetReceiptByIDResponse.Item.Order](#inventory-v1-GetReceiptByIDResponse-Item-Order)
  
    - [InventoryService](#inventory-v1-InventoryService)
  
- [order/v1/order.proto](#order_v1_order-proto)
    - [Address](#order-v1-Address)
    - [BasicOrder](#order-v1-BasicOrder)
    - [CreateOrderRequest](#order-v1-CreateOrderRequest)
    - [CreateOrderResponse](#order-v1-CreateOrderResponse)
    - [CreateQuoteRequest](#order-v1-CreateQuoteRequest)
    - [CreateQuoteRequest.OrderItem](#order-v1-CreateQuoteRequest-OrderItem)
    - [CreateQuoteResponse](#order-v1-CreateQuoteResponse)
    - [Customer](#order-v1-Customer)
    - [CustomerContact](#order-v1-CustomerContact)
    - [Filters](#order-v1-Filters)
    - [GetOrderRequest](#order-v1-GetOrderRequest)
    - [GetOrderResponse](#order-v1-GetOrderResponse)
    - [GetOrderResponse.OrderItem](#order-v1-GetOrderResponse-OrderItem)
    - [GetPickTicketByIdRequest](#order-v1-GetPickTicketByIdRequest)
    - [GetPickTicketByIdResponse](#order-v1-GetPickTicketByIdResponse)
    - [ListOrdersByCustomerBranchRequest](#order-v1-ListOrdersByCustomerBranchRequest)
    - [ListOrdersByCustomerBranchResponse](#order-v1-ListOrdersByCustomerBranchResponse)
    - [ListOrdersByCustomerRequest](#order-v1-ListOrdersByCustomerRequest)
    - [ListOrdersByCustomerResponse](#order-v1-ListOrdersByCustomerResponse)
    - [ListOrdersByTakerRequest](#order-v1-ListOrdersByTakerRequest)
    - [ListOrdersByTakerResponse](#order-v1-ListOrdersByTakerResponse)
    - [OrderDetails](#order-v1-OrderDetails)
    - [OrderItem](#order-v1-OrderItem)
    - [PageMetadata](#order-v1-PageMetadata)
  
    - [PageDirection](#order-v1-PageDirection)
  
    - [OrderService](#order-v1-OrderService)
  
- [supplier/v1/supplier.proto](#supplier_v1_supplier-proto)
    - [SupplierService](#supplier-v1-SupplierService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="catalog_v1_catalog-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## catalog/v1/catalog.proto



<a name="catalog-v1-CreateProductGroupRequest"></a>

### CreateProductGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group | [ProductGroup](#catalog-v1-ProductGroup) |  | string sn = 1; //Unique identifier of Product Group string name = 2; // Description of group |






<a name="catalog-v1-CreateProductGroupResponse"></a>

### CreateProductGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validation_errors | [ValidationError](#catalog-v1-ValidationError) | repeated |  |






<a name="catalog-v1-CreateProductRequest"></a>

### CreateProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| extended_description | [string](#string) |  |  |
| list_price | [float](#float) |  |  |
| product_group_id | [string](#string) |  |  |
| upc | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |
| supplier_part_no | [string](#string) |  |  |
| cost | [float](#float) |  |  |






<a name="catalog-v1-CreateProductResponse"></a>

### CreateProductResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validation_errors | [ValidationError](#catalog-v1-ValidationError) | repeated |  |






<a name="catalog-v1-CreateSupplierRequest"></a>

### CreateSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |
| supplier_product_sn | [string](#string) |  |  |
| list_price | [double](#double) |  |  |
| purchase_price | [double](#double) |  |  |






<a name="catalog-v1-CreateSupplierResponse"></a>

### CreateSupplierResponse







<a name="catalog-v1-Cursor"></a>

### Cursor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| select | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |






<a name="catalog-v1-DeleteProductRequest"></a>

### DeleteProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |






<a name="catalog-v1-DeleteProductResponse"></a>

### DeleteProductResponse







<a name="catalog-v1-GetBasicProductDetailsRequest"></a>

### GetBasicProductDetailsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) | repeated |  |






<a name="catalog-v1-GetBasicProductDetailsResponse"></a>

### GetBasicProductDetailsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| basic_product_details | [GetBasicProductDetailsResponse.BasicProductDetail](#catalog-v1-GetBasicProductDetailsResponse-BasicProductDetail) | repeated |  |






<a name="catalog-v1-GetBasicProductDetailsResponse-BasicProductDetail"></a>

### GetBasicProductDetailsResponse.BasicProductDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="catalog-v1-GetProductBySupplierRequest"></a>

### GetProductBySupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| supplier_id | [double](#double) |  |  |
| supplier_part_no | [string](#string) |  |  |






<a name="catalog-v1-GetProductBySupplierResponse"></a>

### GetProductBySupplierResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sn | [string](#string) |  |  |






<a name="catalog-v1-GetProductGroupRequest"></a>

### GetProductGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sn | [string](#string) |  |  |






<a name="catalog-v1-GetProductGroupResponse"></a>

### GetProductGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group | [ProductGroup](#catalog-v1-ProductGroup) |  |  |
| products | [ProductDetail](#catalog-v1-ProductDetail) | repeated | **Deprecated.**  |
| product_uids | [int32](#int32) | repeated |  |






<a name="catalog-v1-GetProductPriceRequest"></a>

### GetProductPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) | repeated |  |
| customer_id | [string](#string) |  |  |






<a name="catalog-v1-GetProductPriceResponse"></a>

### GetProductPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_prices | [GetProductPriceResponse.ProductPrice](#catalog-v1-GetProductPriceResponse-ProductPrice) | repeated |  |






<a name="catalog-v1-GetProductPriceResponse-ProductPrice"></a>

### GetProductPriceResponse.ProductPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| list_price | [double](#double) |  |  |
| customer_price | [double](#double) |  |  |






<a name="catalog-v1-GetProductRequest"></a>

### GetProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="catalog-v1-GetProductResponse"></a>

### GetProductResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [ProductDetail](#catalog-v1-ProductDetail) |  |  |






<a name="catalog-v1-GetSupplierRequest"></a>

### GetSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |






<a name="catalog-v1-GetSupplierResponse"></a>

### GetSupplierResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_supplier | [ProductSupplier](#catalog-v1-ProductSupplier) |  |  |






<a name="catalog-v1-ListProductGroupsRequest"></a>

### ListProductGroupsRequest







<a name="catalog-v1-ListProductGroupsResponse"></a>

### ListProductGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_groups | [ProductGroup](#catalog-v1-ProductGroup) | repeated |  |






<a name="catalog-v1-ListProductsRequest"></a>

### ListProductsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int32](#int32) |  |  |






<a name="catalog-v1-ListProductsResponse"></a>

### ListProductsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [ProductDetail](#catalog-v1-ProductDetail) | repeated |  |
| next_cursor | [int32](#int32) |  |  |






<a name="catalog-v1-ListSuppliersRequest"></a>

### ListSuppliersRequest







<a name="catalog-v1-ListSuppliersResponse"></a>

### ListSuppliersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_suppliers | [ProductSupplier](#catalog-v1-ProductSupplier) | repeated |  |






<a name="catalog-v1-ProductDetail"></a>

### ProductDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| stock_qty | [double](#double) |  |  |
| product_group_sn | [string](#string) |  |  |






<a name="catalog-v1-ProductFilter"></a>

### ProductFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group_id | [string](#string) |  |  |






<a name="catalog-v1-ProductGroup"></a>

### ProductGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uid | [int32](#int32) |  |  |






<a name="catalog-v1-ProductSupplier"></a>

### ProductSupplier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |
| supplier_sn | [string](#string) |  |  |
| list_price | [double](#double) |  |  |
| purchase_price | [double](#double) |  |  |
| delete | [bool](#bool) |  |  |






<a name="catalog-v1-SetPrimarySupplierRequest"></a>

### SetPrimarySupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| location_id | [double](#double) |  |  |
| supplier_id | [double](#double) |  |  |
| division_id | [double](#double) |  |  |






<a name="catalog-v1-SetPrimarySupplierResponse"></a>

### SetPrimarySupplierResponse







<a name="catalog-v1-UpdateProductGroupRequest"></a>

### UpdateProductGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group | [ProductGroup](#catalog-v1-ProductGroup) |  |  |






<a name="catalog-v1-UpdateProductGroupResponse"></a>

### UpdateProductGroupResponse







<a name="catalog-v1-UpdateProductRequest"></a>

### UpdateProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [ProductDetail](#catalog-v1-ProductDetail) |  |  |






<a name="catalog-v1-UpdateProductResponse"></a>

### UpdateProductResponse







<a name="catalog-v1-UpdateSupplierRequest"></a>

### UpdateSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_supplier | [ProductSupplier](#catalog-v1-ProductSupplier) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="catalog-v1-UpdateSupplierResponse"></a>

### UpdateSupplierResponse







<a name="catalog-v1-ValidationError"></a>

### ValidationError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [string](#string) |  |  |
| message | [string](#string) |  |  |





 

 

 


<a name="catalog-v1-CatalogService"></a>

### CatalogService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProducts | [ListProductsRequest](#catalog-v1-ListProductsRequest) | [ListProductsResponse](#catalog-v1-ListProductsResponse) | ListProducts returns a List of products |
| GetProduct | [GetProductRequest](#catalog-v1-GetProductRequest) | [GetProductResponse](#catalog-v1-GetProductResponse) | GetProduct returns a single product |
| CreateProduct | [CreateProductRequest](#catalog-v1-CreateProductRequest) | [CreateProductResponse](#catalog-v1-CreateProductResponse) | CreateProduct creates a product |
| UpdateProduct | [UpdateProductRequest](#catalog-v1-UpdateProductRequest) | [UpdateProductResponse](#catalog-v1-UpdateProductResponse) | UpdateProduct updates a product |
| DeleteProduct | [DeleteProductRequest](#catalog-v1-DeleteProductRequest) | [DeleteProductResponse](#catalog-v1-DeleteProductResponse) | DeleteProduct deletes a product |
| GetBasicProductDetails | [GetBasicProductDetailsRequest](#catalog-v1-GetBasicProductDetailsRequest) | [GetBasicProductDetailsResponse](#catalog-v1-GetBasicProductDetailsResponse) | GetBasicProductDetails returns a List of basic product details |
| GetProductPrice | [GetProductPriceRequest](#catalog-v1-GetProductPriceRequest) | [GetProductPriceResponse](#catalog-v1-GetProductPriceResponse) | GetProductPrice returns the price of a product |
| ListSuppliers | [ListSuppliersRequest](#catalog-v1-ListSuppliersRequest) | [ListSuppliersResponse](#catalog-v1-ListSuppliersResponse) | ListSuppliers returns a List of suppliers for a product |
| GetSupplier | [GetSupplierRequest](#catalog-v1-GetSupplierRequest) | [GetSupplierResponse](#catalog-v1-GetSupplierResponse) | GetSupplier returns the supplier of a product |
| CreateSupplier | [CreateSupplierRequest](#catalog-v1-CreateSupplierRequest) | [CreateSupplierResponse](#catalog-v1-CreateSupplierResponse) | CreateSupplier creates a supplier for a product |
| UpdateSupplier | [UpdateSupplierRequest](#catalog-v1-UpdateSupplierRequest) | [UpdateSupplierResponse](#catalog-v1-UpdateSupplierResponse) | UpdateSupplier updates a supplier for a product |
| SetPrimarySupplier | [SetPrimarySupplierRequest](#catalog-v1-SetPrimarySupplierRequest) | [SetPrimarySupplierResponse](#catalog-v1-SetPrimarySupplierResponse) | SetPrimarySupplier sets the primary supplier for a product |
| ListProductGroups | [ListProductGroupsRequest](#catalog-v1-ListProductGroupsRequest) | [ListProductGroupsResponse](#catalog-v1-ListProductGroupsResponse) | ListProductGroups returns a List of product groups |
| GetProductGroup | [GetProductGroupRequest](#catalog-v1-GetProductGroupRequest) | [GetProductGroupResponse](#catalog-v1-GetProductGroupResponse) | GetProductGroup returns a single product group |
| CreateProductGroup | [CreateProductGroupRequest](#catalog-v1-CreateProductGroupRequest) | [CreateProductGroupResponse](#catalog-v1-CreateProductGroupResponse) | CreateProductGroup creates a product group |
| UpdateProductGroup | [UpdateProductGroupRequest](#catalog-v1-UpdateProductGroupRequest) | [UpdateProductGroupResponse](#catalog-v1-UpdateProductGroupResponse) | UpdateProductGroup updates a product group |
| GetProductBySupplier | [GetProductBySupplierRequest](#catalog-v1-GetProductBySupplierRequest) | [GetProductBySupplierResponse](#catalog-v1-GetProductBySupplierResponse) | GetProductBySupplier returns a product by supplier identifiers |

 



<a name="customer_v1_customer-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## customer/v1/customer.proto



<a name="customer-v1-CustomerBranch"></a>

### CustomerBranch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |
| name | [string](#string) |  |  |






<a name="customer-v1-GetContactByIdRequest"></a>

### GetContactByIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="customer-v1-GetContactByIdResponse"></a>

### GetContactByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contact | [GetContactByIdResponse.Contact](#customer-v1-GetContactByIdResponse-Contact) |  |  |
| branches | [CustomerBranch](#customer-v1-CustomerBranch) | repeated |  |






<a name="customer-v1-GetContactByIdResponse-Contact"></a>

### GetContactByIdResponse.Contact



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| first_name | [string](#string) |  |  |
| last_name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phone | [string](#string) |  |  |






<a name="customer-v1-GetCustomerRequest"></a>

### GetCustomerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_id | [double](#double) |  |  |






<a name="customer-v1-GetCustomerResponse"></a>

### GetCustomerResponse







<a name="customer-v1-GetOrdersRequest"></a>

### GetOrdersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_id | [string](#string) |  |  |






<a name="customer-v1-GetOrdersResponse"></a>

### GetOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [GetOrdersResponse.Order](#customer-v1-GetOrdersResponse-Order) | repeated |  |






<a name="customer-v1-GetOrdersResponse-Order"></a>

### GetOrdersResponse.Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |





 

 

 


<a name="customer-v1-CustomerService"></a>

### CustomerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOrders | [GetOrdersRequest](#customer-v1-GetOrdersRequest) | [GetOrdersResponse](#customer-v1-GetOrdersResponse) |  |
| GetCustomer | [GetCustomerRequest](#customer-v1-GetCustomerRequest) | [GetCustomerResponse](#customer-v1-GetCustomerResponse) | GetCustomer returns the customer information |
| GetContactById | [GetContactByIdRequest](#customer-v1-GetContactByIdRequest) | [GetContactByIdResponse](#customer-v1-GetContactByIdResponse) | rpc GetQuotes() returns (); rpc GetInvoices() returns (); |

 



<a name="inventory_v1_inventory-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## inventory/v1/inventory.proto



<a name="inventory-v1-AddSupplierRequest"></a>

### AddSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serial_id | [int32](#int32) |  |  |
| supplier_id | [int32](#int32) |  |  |
| division_id | [int32](#int32) |  |  |
| lead_time | [int32](#int32) |  |  |
| list_price | [float](#float) |  |  |
| cost | [float](#float) |  |  |
| supplier_part_no | [string](#string) |  |  |






<a name="inventory-v1-AddSupplierResponse"></a>

### AddSupplierResponse







<a name="inventory-v1-DeleteSupplierRequest"></a>

### DeleteSupplierRequest







<a name="inventory-v1-DeleteSupplierResponse"></a>

### DeleteSupplierResponse







<a name="inventory-v1-GetProductStockRequest"></a>

### GetProductStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) | repeated |  |






<a name="inventory-v1-GetProductStockResponse"></a>

### GetProductStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_stock | [GetProductStockResponse.ProductStock](#inventory-v1-GetProductStockResponse-ProductStock) | repeated |  |






<a name="inventory-v1-GetProductStockResponse-ProductStock"></a>

### GetProductStockResponse.ProductStock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| quantity_available | [double](#double) |  |  |






<a name="inventory-v1-GetReceiptByIDRequest"></a>

### GetReceiptByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |






<a name="inventory-v1-GetReceiptByIDResponse"></a>

### GetReceiptByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  | Unique ID of receipt |
| items | [GetReceiptByIDResponse.Item](#inventory-v1-GetReceiptByIDResponse-Item) | repeated | Items that were received for this receipt |






<a name="inventory-v1-GetReceiptByIDResponse-Item"></a>

### GetReceiptByIDResponse.Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  | Unique ID of product |
| product_primary_bin | [string](#string) |  | Primary bin where product is stored |
| received_quantity | [double](#double) |  | Quantity of product received |
| received_unit | [string](#string) |  | Unit associated with received_quantity |
| allocated_orders | [string](#string) | repeated | Orders that were allocated to |






<a name="inventory-v1-GetReceiptByIDResponse-Item-Order"></a>

### GetReceiptByIDResponse.Item.Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique ID of order |
| customer_name | [string](#string) |  | Customer Name |
| allocated_quantity | [double](#double) |  | Quantity of product allocated |
| allocated_unit | [string](#string) |  | Unit associated with allocated_quantity |





 

 

 


<a name="inventory-v1-InventoryService"></a>

### InventoryService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetProductStock | [GetProductStockRequest](#inventory-v1-GetProductStockRequest) | [GetProductStockResponse](#inventory-v1-GetProductStockResponse) | rpc AddSupplier(AddSupplierRequest) returns (AddSupplierResponse); rpc DeleteSupplier(DeleteSupplierRequest) returns (DeleteSupplierResponse); |
| GetReceiptByID | [GetReceiptByIDRequest](#inventory-v1-GetReceiptByIDRequest) | [GetReceiptByIDResponse](#inventory-v1-GetReceiptByIDResponse) | GetReceiptByID returns details for a inventory receipt given an identifier. |

 



<a name="order_v1_order-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order/v1/order.proto



<a name="order-v1-Address"></a>

### Address



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |
| name | [string](#string) |  |  |
| line_one | [string](#string) |  |  |
| line_two | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| postal_code | [string](#string) |  |  |
| country | [string](#string) |  |  |






<a name="order-v1-BasicOrder"></a>

### BasicOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| taker | [string](#string) |  |  |
| purchase_order | [string](#string) |  |  |
| order_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| requested_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| completed | [bool](#bool) |  |  |






<a name="order-v1-CreateOrderRequest"></a>

### CreateOrderRequest







<a name="order-v1-CreateOrderResponse"></a>

### CreateOrderResponse







<a name="order-v1-CreateQuoteRequest"></a>

### CreateQuoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_items | [CreateQuoteRequest.OrderItem](#order-v1-CreateQuoteRequest-OrderItem) | repeated |  |
| customer_id | [double](#double) |  | **Deprecated.** Deprecated: customer_id is pulled from the customer_branch_id |
| shipping_address_id | [double](#double) |  | **Deprecated.** Deprecated: use company_branch_id |
| contact_id | [string](#string) |  | customer accounts contact_id |
| requested_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | date the customer requested the quote to be completed |
| customer_branch_id | [double](#double) |  | represents ship_to in prophet and is used as a sub location of a customer record |
| delivery_instructions | [string](#string) |  | instructions pertaining to the delivery of the order |
| purchase_order | [string](#string) |  | customer generated purchase order number for this quote |






<a name="order-v1-CreateQuoteRequest-OrderItem"></a>

### CreateQuoteRequest.OrderItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  | unique identifier for the product |
| quantity_ordered | [double](#double) |  | quantity of the product to be quoted |
| required_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | **Deprecated.**  |






<a name="order-v1-CreateQuoteResponse"></a>

### CreateQuoteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id of the created quote |






<a name="order-v1-Customer"></a>

### Customer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |
| name | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="order-v1-CustomerContact"></a>

### CustomerContact



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |
| email | [string](#string) |  |  |
| title | [string](#string) |  |  |






<a name="order-v1-Filters"></a>

### Filters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int32](#int32) |  |  |
| direction | [PageDirection](#order-v1-PageDirection) |  |  |






<a name="order-v1-GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="order-v1-GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| order_items | [GetOrderResponse.OrderItem](#order-v1-GetOrderResponse-OrderItem) | repeated |  |
| status | [string](#string) |  |  |
| shipping_address | [Address](#order-v1-Address) |  |  |
| customer | [Customer](#order-v1-Customer) |  |  |
| contact | [CustomerContact](#order-v1-CustomerContact) |  |  |
| delivery_instructions | [string](#string) |  |  |
| taker | [string](#string) |  |  |
| purchase_order | [string](#string) |  |  |






<a name="order-v1-GetOrderResponse-OrderItem"></a>

### GetOrderResponse.OrderItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| quantity_ordered | [double](#double) |  |  |
| quantity_unit | [string](#string) |  |  |
| cost_per_unit | [double](#double) |  |  |
| total_price | [double](#double) |  |  |
| product_uid | [int32](#int32) |  |  |






<a name="order-v1-GetPickTicketByIdRequest"></a>

### GetPickTicketByIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |






<a name="order-v1-GetPickTicketByIdResponse"></a>

### GetPickTicketByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |
| order_id | [string](#string) |  |  |
| order_purchase_order | [string](#string) |  |  |
| order_contact | [CustomerContact](#order-v1-CustomerContact) |  |  |
| shipping_address | [Address](#order-v1-Address) |  |  |
| delivery_instructions | [string](#string) |  |  |






<a name="order-v1-ListOrdersByCustomerBranchRequest"></a>

### ListOrdersByCustomerBranchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_branch_id | [double](#double) |  |  |
| filters | [Filters](#order-v1-Filters) |  |  |






<a name="order-v1-ListOrdersByCustomerBranchResponse"></a>

### ListOrdersByCustomerBranchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [BasicOrder](#order-v1-BasicOrder) | repeated |  |
| metadata | [PageMetadata](#order-v1-PageMetadata) |  |  |






<a name="order-v1-ListOrdersByCustomerRequest"></a>

### ListOrdersByCustomerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_id | [double](#double) |  |  |
| filters | [Filters](#order-v1-Filters) |  |  |






<a name="order-v1-ListOrdersByCustomerResponse"></a>

### ListOrdersByCustomerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [BasicOrder](#order-v1-BasicOrder) | repeated |  |
| metadata | [PageMetadata](#order-v1-PageMetadata) |  |  |






<a name="order-v1-ListOrdersByTakerRequest"></a>

### ListOrdersByTakerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taker | [string](#string) |  |  |






<a name="order-v1-ListOrdersByTakerResponse"></a>

### ListOrdersByTakerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [BasicOrder](#order-v1-BasicOrder) | repeated |  |






<a name="order-v1-OrderDetails"></a>

### OrderDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| taker | [string](#string) |  |  |
| purchase_order | [string](#string) |  |  |
| shipping_address | [Address](#order-v1-Address) |  |  |
| customer | [Customer](#order-v1-Customer) |  |  |
| contact | [CustomerContact](#order-v1-CustomerContact) |  |  |
| delivery_instructions | [string](#string) |  |  |






<a name="order-v1-OrderItem"></a>

### OrderItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| quantity_ordered | [double](#double) |  |  |
| required_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | **Deprecated.**  |






<a name="order-v1-PageMetadata"></a>

### PageMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prev_cursor | [int32](#int32) |  |  |
| next_cursor | [int32](#int32) |  |  |





 


<a name="order-v1-PageDirection"></a>

### PageDirection


| Name | Number | Description |
| ---- | ------ | ----------- |
| PAGE_DIRECTION_UNSPECIFIED | 0 |  |
| PAGE_DIRECTION_NEXT | 1 |  |
| PAGE_DIRECTION_PREVIOUS | 2 |  |


 

 


<a name="order-v1-OrderService"></a>

### OrderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListOrdersByCustomerBranch | [ListOrdersByCustomerBranchRequest](#order-v1-ListOrdersByCustomerBranchRequest) | [ListOrdersByCustomerBranchResponse](#order-v1-ListOrdersByCustomerBranchResponse) |  |
| ListOrdersByCustomer | [ListOrdersByCustomerRequest](#order-v1-ListOrdersByCustomerRequest) | [ListOrdersByCustomerResponse](#order-v1-ListOrdersByCustomerResponse) |  |
| ListOrdersByTaker | [ListOrdersByTakerRequest](#order-v1-ListOrdersByTakerRequest) | [ListOrdersByTakerResponse](#order-v1-ListOrdersByTakerResponse) |  |
| GetOrder | [GetOrderRequest](#order-v1-GetOrderRequest) | [GetOrderResponse](#order-v1-GetOrderResponse) | GetOrder returns the order details for a given order id |
| CreateOrder | [CreateOrderRequest](#order-v1-CreateOrderRequest) | [CreateOrderResponse](#order-v1-CreateOrderResponse) |  |
| CreateQuote | [CreateQuoteRequest](#order-v1-CreateQuoteRequest) | [CreateQuoteResponse](#order-v1-CreateQuoteResponse) |  |
| GetPickTicketById | [GetPickTicketByIdRequest](#order-v1-GetPickTicketByIdRequest) | [GetPickTicketByIdResponse](#order-v1-GetPickTicketByIdResponse) |  |

 



<a name="supplier_v1_supplier-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## supplier/v1/supplier.proto


 

 

 


<a name="supplier-v1-SupplierService"></a>

### SupplierService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

