# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [catalog/v1alpha0/catalog.proto](#catalog_v1alpha0_catalog-proto)
    - [CreateGroupRequest](#catalog-v1alpha0-CreateGroupRequest)
    - [CreateGroupResponse](#catalog-v1alpha0-CreateGroupResponse)
    - [CreateProductRequest](#catalog-v1alpha0-CreateProductRequest)
    - [CreateProductResponse](#catalog-v1alpha0-CreateProductResponse)
    - [CreateSupplierRequest](#catalog-v1alpha0-CreateSupplierRequest)
    - [CreateSupplierResponse](#catalog-v1alpha0-CreateSupplierResponse)
    - [Cursor](#catalog-v1alpha0-Cursor)
    - [DeleteProductRequest](#catalog-v1alpha0-DeleteProductRequest)
    - [DeleteProductResponse](#catalog-v1alpha0-DeleteProductResponse)
    - [GetProductBySupplierRequest](#catalog-v1alpha0-GetProductBySupplierRequest)
    - [GetProductBySupplierResponse](#catalog-v1alpha0-GetProductBySupplierResponse)
    - [GetProductGroupRequest](#catalog-v1alpha0-GetProductGroupRequest)
    - [GetProductGroupResponse](#catalog-v1alpha0-GetProductGroupResponse)
    - [GetProductPriceRequest](#catalog-v1alpha0-GetProductPriceRequest)
    - [GetProductPriceResponse](#catalog-v1alpha0-GetProductPriceResponse)
    - [GetProductPriceResponse.ProductPrice](#catalog-v1alpha0-GetProductPriceResponse-ProductPrice)
    - [GetProductRequest](#catalog-v1alpha0-GetProductRequest)
    - [GetProductResponse](#catalog-v1alpha0-GetProductResponse)
    - [GetSupplierRequest](#catalog-v1alpha0-GetSupplierRequest)
    - [ListProductGroupsRequest](#catalog-v1alpha0-ListProductGroupsRequest)
    - [ListProductGroupsResponse](#catalog-v1alpha0-ListProductGroupsResponse)
    - [ListProductsRequest](#catalog-v1alpha0-ListProductsRequest)
    - [ListProductsResponse](#catalog-v1alpha0-ListProductsResponse)
    - [ListSuppliersRequest](#catalog-v1alpha0-ListSuppliersRequest)
    - [ListSuppliersResponse](#catalog-v1alpha0-ListSuppliersResponse)
    - [ProductDetail](#catalog-v1alpha0-ProductDetail)
    - [ProductFilter](#catalog-v1alpha0-ProductFilter)
    - [ProductGroup](#catalog-v1alpha0-ProductGroup)
    - [ProductSupplier](#catalog-v1alpha0-ProductSupplier)
    - [SetPrimarySupplierRequest](#catalog-v1alpha0-SetPrimarySupplierRequest)
    - [SetPrimarySupplierResponse](#catalog-v1alpha0-SetPrimarySupplierResponse)
    - [UpdateGroupRequest](#catalog-v1alpha0-UpdateGroupRequest)
    - [UpdateGroupResponse](#catalog-v1alpha0-UpdateGroupResponse)
    - [UpdateProductRequest](#catalog-v1alpha0-UpdateProductRequest)
    - [UpdateProductResponse](#catalog-v1alpha0-UpdateProductResponse)
    - [UpdateSupplierRequest](#catalog-v1alpha0-UpdateSupplierRequest)
    - [UpdateSupplierResponse](#catalog-v1alpha0-UpdateSupplierResponse)
    - [ValidationError](#catalog-v1alpha0-ValidationError)
  
    - [CatalogService](#catalog-v1alpha0-CatalogService)
  
- [customer/v1alpha0/customer.proto](#customer_v1alpha0_customer-proto)
    - [GetOrdersRequest](#customer-v1alpha0-GetOrdersRequest)
    - [GetOrdersResponse](#customer-v1alpha0-GetOrdersResponse)
    - [GetOrdersResponse.Order](#customer-v1alpha0-GetOrdersResponse-Order)
  
    - [CustomerService](#customer-v1alpha0-CustomerService)
  
- [inventory/v1alpha0/inventory.proto](#inventory_v1alpha0_inventory-proto)
    - [AddSupplierRequest](#inventory-v1alpha0-AddSupplierRequest)
    - [AddSupplierResponse](#inventory-v1alpha0-AddSupplierResponse)
    - [DeleteSupplierRequest](#inventory-v1alpha0-DeleteSupplierRequest)
    - [DeleteSupplierResponse](#inventory-v1alpha0-DeleteSupplierResponse)
    - [GetProductStockRequest](#inventory-v1alpha0-GetProductStockRequest)
    - [GetProductStockResponse](#inventory-v1alpha0-GetProductStockResponse)
    - [GetProductStockResponse.productStock](#inventory-v1alpha0-GetProductStockResponse-productStock)
    - [GetReceiptByIDRequest](#inventory-v1alpha0-GetReceiptByIDRequest)
    - [GetReceiptByIDResponse](#inventory-v1alpha0-GetReceiptByIDResponse)
    - [GetReceiptByIDResponse.Item](#inventory-v1alpha0-GetReceiptByIDResponse-Item)
    - [GetReceiptByIDResponse.Item.Order](#inventory-v1alpha0-GetReceiptByIDResponse-Item-Order)
  
    - [InventoryService](#inventory-v1alpha0-InventoryService)
  
- [order/v1alpha0/order.proto](#order_v1alpha0_order-proto)
    - [Address](#order-v1alpha0-Address)
    - [CreateOrderRequest](#order-v1alpha0-CreateOrderRequest)
    - [CreateOrderResponse](#order-v1alpha0-CreateOrderResponse)
    - [Customer](#order-v1alpha0-Customer)
    - [CustomerContact](#order-v1alpha0-CustomerContact)
    - [GetOrderRequest](#order-v1alpha0-GetOrderRequest)
    - [GetOrderResponse](#order-v1alpha0-GetOrderResponse)
    - [GetOrderResponse.OrderItem](#order-v1alpha0-GetOrderResponse-OrderItem)
    - [GetPickTicketByIdRequest](#order-v1alpha0-GetPickTicketByIdRequest)
    - [GetPickTicketByIdResponse](#order-v1alpha0-GetPickTicketByIdResponse)
  
    - [OrderService](#order-v1alpha0-OrderService)
  
- [supplier/v1alpha0/supplier.proto](#supplier_v1alpha0_supplier-proto)
    - [SupplierService](#supplier-v1alpha0-SupplierService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="catalog_v1alpha0_catalog-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## catalog/v1alpha0/catalog.proto



<a name="catalog-v1alpha0-CreateGroupRequest"></a>

### CreateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group | [ProductGroup](#catalog-v1alpha0-ProductGroup) |  | string sn = 1; //Unique identifier of Product Group string name = 2; // Description of group |






<a name="catalog-v1alpha0-CreateGroupResponse"></a>

### CreateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validation_errors | [ValidationError](#catalog-v1alpha0-ValidationError) | repeated |  |






<a name="catalog-v1alpha0-CreateProductRequest"></a>

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






<a name="catalog-v1alpha0-CreateProductResponse"></a>

### CreateProductResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validation_errors | [ValidationError](#catalog-v1alpha0-ValidationError) | repeated |  |






<a name="catalog-v1alpha0-CreateSupplierRequest"></a>

### CreateSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |
| supplier_product_sn | [string](#string) |  |  |
| list_price | [double](#double) |  |  |
| purchase_price | [double](#double) |  |  |






<a name="catalog-v1alpha0-CreateSupplierResponse"></a>

### CreateSupplierResponse







<a name="catalog-v1alpha0-Cursor"></a>

### Cursor



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| select | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |






<a name="catalog-v1alpha0-DeleteProductRequest"></a>

### DeleteProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |






<a name="catalog-v1alpha0-DeleteProductResponse"></a>

### DeleteProductResponse







<a name="catalog-v1alpha0-GetProductBySupplierRequest"></a>

### GetProductBySupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| supplier_id | [string](#string) |  |  |
| supplier_part_no | [string](#string) |  |  |






<a name="catalog-v1alpha0-GetProductBySupplierResponse"></a>

### GetProductBySupplierResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sn | [string](#string) |  |  |






<a name="catalog-v1alpha0-GetProductGroupRequest"></a>

### GetProductGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sn | [string](#string) |  |  |






<a name="catalog-v1alpha0-GetProductGroupResponse"></a>

### GetProductGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group | [ProductGroup](#catalog-v1alpha0-ProductGroup) |  |  |
| products | [ProductDetail](#catalog-v1alpha0-ProductDetail) | repeated |  |






<a name="catalog-v1alpha0-GetProductPriceRequest"></a>

### GetProductPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) | repeated |  |
| customer_id | [string](#string) |  |  |






<a name="catalog-v1alpha0-GetProductPriceResponse"></a>

### GetProductPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_prices | [GetProductPriceResponse.ProductPrice](#catalog-v1alpha0-GetProductPriceResponse-ProductPrice) | repeated |  |






<a name="catalog-v1alpha0-GetProductPriceResponse-ProductPrice"></a>

### GetProductPriceResponse.ProductPrice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| list_price | [double](#double) |  |  |
| customer_price | [double](#double) |  |  |






<a name="catalog-v1alpha0-GetProductRequest"></a>

### GetProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="catalog-v1alpha0-GetProductResponse"></a>

### GetProductResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [ProductDetail](#catalog-v1alpha0-ProductDetail) |  |  |






<a name="catalog-v1alpha0-GetSupplierRequest"></a>

### GetSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |






<a name="catalog-v1alpha0-ListProductGroupsRequest"></a>

### ListProductGroupsRequest







<a name="catalog-v1alpha0-ListProductGroupsResponse"></a>

### ListProductGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_groups | [ProductGroup](#catalog-v1alpha0-ProductGroup) | repeated |  |






<a name="catalog-v1alpha0-ListProductsRequest"></a>

### ListProductsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int32](#int32) |  |  |






<a name="catalog-v1alpha0-ListProductsResponse"></a>

### ListProductsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [ProductDetail](#catalog-v1alpha0-ProductDetail) | repeated |  |
| next_cursor | [int32](#int32) |  |  |






<a name="catalog-v1alpha0-ListSuppliersRequest"></a>

### ListSuppliersRequest







<a name="catalog-v1alpha0-ListSuppliersResponse"></a>

### ListSuppliersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_suppliers | [ProductSupplier](#catalog-v1alpha0-ProductSupplier) | repeated |  |






<a name="catalog-v1alpha0-ProductDetail"></a>

### ProductDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| stock_qty | [double](#double) |  |  |
| product_group_sn | [string](#string) |  |  |






<a name="catalog-v1alpha0-ProductFilter"></a>

### ProductFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group_id | [string](#string) |  |  |






<a name="catalog-v1alpha0-ProductGroup"></a>

### ProductGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uid | [int32](#int32) |  |  |






<a name="catalog-v1alpha0-ProductSupplier"></a>

### ProductSupplier



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [string](#string) |  |  |
| supplier_id | [string](#string) |  |  |
| supplier_sn | [string](#string) |  |  |
| list_price | [double](#double) |  |  |
| purchase_price | [double](#double) |  |  |
| delete | [bool](#bool) |  |  |






<a name="catalog-v1alpha0-SetPrimarySupplierRequest"></a>

### SetPrimarySupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| location_id | [double](#double) |  |  |
| supplier_id | [double](#double) |  |  |
| division_id | [double](#double) |  |  |






<a name="catalog-v1alpha0-SetPrimarySupplierResponse"></a>

### SetPrimarySupplierResponse







<a name="catalog-v1alpha0-UpdateGroupRequest"></a>

### UpdateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_group | [ProductGroup](#catalog-v1alpha0-ProductGroup) |  |  |






<a name="catalog-v1alpha0-UpdateGroupResponse"></a>

### UpdateGroupResponse







<a name="catalog-v1alpha0-UpdateProductRequest"></a>

### UpdateProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [ProductDetail](#catalog-v1alpha0-ProductDetail) |  |  |






<a name="catalog-v1alpha0-UpdateProductResponse"></a>

### UpdateProductResponse







<a name="catalog-v1alpha0-UpdateSupplierRequest"></a>

### UpdateSupplierRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_supplier | [ProductSupplier](#catalog-v1alpha0-ProductSupplier) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="catalog-v1alpha0-UpdateSupplierResponse"></a>

### UpdateSupplierResponse







<a name="catalog-v1alpha0-ValidationError"></a>

### ValidationError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [string](#string) |  |  |
| message | [string](#string) |  |  |





 

 

 


<a name="catalog-v1alpha0-CatalogService"></a>

### CatalogService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProducts | [ListProductsRequest](#catalog-v1alpha0-ListProductsRequest) | [ListProductsResponse](#catalog-v1alpha0-ListProductsResponse) | ListProducts returns a List of products |
| GetProduct | [GetProductRequest](#catalog-v1alpha0-GetProductRequest) | [GetProductResponse](#catalog-v1alpha0-GetProductResponse) | GetProduct returns a single product |
| CreateProduct | [CreateProductRequest](#catalog-v1alpha0-CreateProductRequest) | [CreateProductResponse](#catalog-v1alpha0-CreateProductResponse) | CreateProduct creates a product |
| UpdateProduct | [UpdateProductRequest](#catalog-v1alpha0-UpdateProductRequest) | [UpdateProductResponse](#catalog-v1alpha0-UpdateProductResponse) | UpdateProduct updates a product |
| DeleteProduct | [DeleteProductRequest](#catalog-v1alpha0-DeleteProductRequest) | [DeleteProductResponse](#catalog-v1alpha0-DeleteProductResponse) | DeleteProduct deletes a product |
| GetProductPrice | [GetProductPriceRequest](#catalog-v1alpha0-GetProductPriceRequest) | [GetProductPriceResponse](#catalog-v1alpha0-GetProductPriceResponse) | GetProductPrice returns the price of a product |
| ListSuppliers | [ListSuppliersRequest](#catalog-v1alpha0-ListSuppliersRequest) | [ListSuppliersResponse](#catalog-v1alpha0-ListSuppliersResponse) | ListSuppliers returns a List of suppliers for a product |
| GetSupplier | [GetSupplierRequest](#catalog-v1alpha0-GetSupplierRequest) | [ProductSupplier](#catalog-v1alpha0-ProductSupplier) | GetSupplier returns the supplier of a product |
| CreateSupplier | [CreateSupplierRequest](#catalog-v1alpha0-CreateSupplierRequest) | [CreateSupplierResponse](#catalog-v1alpha0-CreateSupplierResponse) | CreateSupplier creates a supplier for a product |
| UpdateSupplier | [UpdateSupplierRequest](#catalog-v1alpha0-UpdateSupplierRequest) | [UpdateSupplierResponse](#catalog-v1alpha0-UpdateSupplierResponse) | UpdateSupplier updates a supplier for a product |
| SetPrimarySupplier | [SetPrimarySupplierRequest](#catalog-v1alpha0-SetPrimarySupplierRequest) | [SetPrimarySupplierResponse](#catalog-v1alpha0-SetPrimarySupplierResponse) | SetPrimarySupplier sets the primary supplier for a product |
| ListProductGroups | [ListProductGroupsRequest](#catalog-v1alpha0-ListProductGroupsRequest) | [ListProductGroupsResponse](#catalog-v1alpha0-ListProductGroupsResponse) | ListProductGroups returns a List of product groups |
| GetProductGroup | [GetProductGroupRequest](#catalog-v1alpha0-GetProductGroupRequest) | [GetProductGroupResponse](#catalog-v1alpha0-GetProductGroupResponse) | GetProductGroup returns a single product group |
| CreateProductGroup | [CreateGroupRequest](#catalog-v1alpha0-CreateGroupRequest) | [CreateGroupResponse](#catalog-v1alpha0-CreateGroupResponse) | CreateProductGroup creates a product group |
| UpdateProductGroup | [UpdateGroupRequest](#catalog-v1alpha0-UpdateGroupRequest) | [UpdateGroupResponse](#catalog-v1alpha0-UpdateGroupResponse) | UpdateProductGroup updates a product group |
| GetProductBySupplier | [GetProductBySupplierRequest](#catalog-v1alpha0-GetProductBySupplierRequest) | [GetProductBySupplierResponse](#catalog-v1alpha0-GetProductBySupplierResponse) | GetProductBySupplier returns a product by supplier identifiers |

 



<a name="customer_v1alpha0_customer-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## customer/v1alpha0/customer.proto



<a name="customer-v1alpha0-GetOrdersRequest"></a>

### GetOrdersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| customer_id | [string](#string) |  |  |






<a name="customer-v1alpha0-GetOrdersResponse"></a>

### GetOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [GetOrdersResponse.Order](#customer-v1alpha0-GetOrdersResponse-Order) | repeated |  |






<a name="customer-v1alpha0-GetOrdersResponse-Order"></a>

### GetOrdersResponse.Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |





 

 

 


<a name="customer-v1alpha0-CustomerService"></a>

### CustomerService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOrders | [GetOrdersRequest](#customer-v1alpha0-GetOrdersRequest) | [GetOrdersResponse](#customer-v1alpha0-GetOrdersResponse) | rpc GetQuotes() returns (); rpc GetInvoices() returns (); |

 



<a name="inventory_v1alpha0_inventory-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## inventory/v1alpha0/inventory.proto



<a name="inventory-v1alpha0-AddSupplierRequest"></a>

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






<a name="inventory-v1alpha0-AddSupplierResponse"></a>

### AddSupplierResponse







<a name="inventory-v1alpha0-DeleteSupplierRequest"></a>

### DeleteSupplierRequest







<a name="inventory-v1alpha0-DeleteSupplierResponse"></a>

### DeleteSupplierResponse







<a name="inventory-v1alpha0-GetProductStockRequest"></a>

### GetProductStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) | repeated |  |






<a name="inventory-v1alpha0-GetProductStockResponse"></a>

### GetProductStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_stock | [GetProductStockResponse.productStock](#inventory-v1alpha0-GetProductStockResponse-productStock) | repeated |  |






<a name="inventory-v1alpha0-GetProductStockResponse-productStock"></a>

### GetProductStockResponse.productStock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  |  |
| quantity_available | [double](#double) |  |  |






<a name="inventory-v1alpha0-GetReceiptByIDRequest"></a>

### GetReceiptByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |






<a name="inventory-v1alpha0-GetReceiptByIDResponse"></a>

### GetReceiptByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  | Unique ID of receipt |
| items | [GetReceiptByIDResponse.Item](#inventory-v1alpha0-GetReceiptByIDResponse-Item) | repeated | Items that were received for this receipt |






<a name="inventory-v1alpha0-GetReceiptByIDResponse-Item"></a>

### GetReceiptByIDResponse.Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_uid | [int32](#int32) |  | Unique ID of product |
| product_primary_bin | [string](#string) |  | Primary bin where product is stored |
| received_quantity | [double](#double) |  | Quantity of product received |
| received_unit | [string](#string) |  | Unit associated with received_quantity |
| allocated_orders | [string](#string) | repeated | Orders that were allocated to |






<a name="inventory-v1alpha0-GetReceiptByIDResponse-Item-Order"></a>

### GetReceiptByIDResponse.Item.Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique ID of order |
| customer_name | [string](#string) |  | Customer Name |
| allocated_quantity | [double](#double) |  | Quantity of product allocated |
| allocated_unit | [string](#string) |  | Unit associated with allocated_quantity |





 

 

 


<a name="inventory-v1alpha0-InventoryService"></a>

### InventoryService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetProductStock | [GetProductStockRequest](#inventory-v1alpha0-GetProductStockRequest) | [GetProductStockResponse](#inventory-v1alpha0-GetProductStockResponse) | rpc AddSupplier(AddSupplierRequest) returns (AddSupplierResponse); rpc DeleteSupplier(DeleteSupplierRequest) returns (DeleteSupplierResponse); |
| GetReceiptByID | [GetReceiptByIDRequest](#inventory-v1alpha0-GetReceiptByIDRequest) | [GetReceiptByIDResponse](#inventory-v1alpha0-GetReceiptByIDResponse) | GetReceiptByID returns details for a inventory receipt given an identifier. |

 



<a name="order_v1alpha0_order-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order/v1alpha0/order.proto



<a name="order-v1alpha0-Address"></a>

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






<a name="order-v1alpha0-CreateOrderRequest"></a>

### CreateOrderRequest







<a name="order-v1alpha0-CreateOrderResponse"></a>

### CreateOrderResponse







<a name="order-v1alpha0-Customer"></a>

### Customer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |
| name | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="order-v1alpha0-CustomerContact"></a>

### CustomerContact



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |
| email | [string](#string) |  |  |
| title | [string](#string) |  |  |






<a name="order-v1alpha0-GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="order-v1alpha0-GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| order_items | [GetOrderResponse.OrderItem](#order-v1alpha0-GetOrderResponse-OrderItem) | repeated |  |
| status | [string](#string) |  |  |
| shipping_address | [Address](#order-v1alpha0-Address) |  |  |
| customer | [Customer](#order-v1alpha0-Customer) |  |  |
| contact | [CustomerContact](#order-v1alpha0-CustomerContact) |  |  |
| delivery_instructions | [string](#string) |  |  |
| taker | [string](#string) |  |  |
| purchase_order | [string](#string) |  |  |






<a name="order-v1alpha0-GetOrderResponse-OrderItem"></a>

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






<a name="order-v1alpha0-GetPickTicketByIdRequest"></a>

### GetPickTicketByIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |






<a name="order-v1alpha0-GetPickTicketByIdResponse"></a>

### GetPickTicketByIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [double](#double) |  |  |
| order_id | [string](#string) |  |  |
| order_purchase_order | [string](#string) |  |  |
| order_contact | [CustomerContact](#order-v1alpha0-CustomerContact) |  |  |
| shipping_address | [Address](#order-v1alpha0-Address) |  |  |
| delivery_instructions | [string](#string) |  |  |





 

 

 


<a name="order-v1alpha0-OrderService"></a>

### OrderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateOrder | [CreateOrderRequest](#order-v1alpha0-CreateOrderRequest) | [CreateOrderResponse](#order-v1alpha0-CreateOrderResponse) |  |
| GetOrder | [GetOrderRequest](#order-v1alpha0-GetOrderRequest) | [GetOrderResponse](#order-v1alpha0-GetOrderResponse) |  |
| GetPickTicketById | [GetPickTicketByIdRequest](#order-v1alpha0-GetPickTicketByIdRequest) | [GetPickTicketByIdResponse](#order-v1alpha0-GetPickTicketByIdResponse) |  |

 



<a name="supplier_v1alpha0_supplier-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## supplier/v1alpha0/supplier.proto


 

 

 


<a name="supplier-v1alpha0-SupplierService"></a>

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

