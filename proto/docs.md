# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [inventory/v1alpha0/inventory.proto](#inventory_v1alpha0_inventory-proto)
    - [AddSupplierRequest](#inventory-v1alpha0-AddSupplierRequest)
    - [AddSupplierResponse](#inventory-v1alpha0-AddSupplierResponse)
    - [DeleteSupplierRequest](#inventory-v1alpha0-DeleteSupplierRequest)
    - [DeleteSupplierResponse](#inventory-v1alpha0-DeleteSupplierResponse)
  
    - [InventoryService](#inventory-v1alpha0-InventoryService)
  
- [order/v1alpha0/order.proto](#order_v1alpha0_order-proto)
    - [GetOrderItemRequest](#order-v1alpha0-GetOrderItemRequest)
    - [GetOrderItemResponse](#order-v1alpha0-GetOrderItemResponse)
    - [GetOrderItemResponse.Item](#order-v1alpha0-GetOrderItemResponse-Item)
    - [GetOrderRequest](#order-v1alpha0-GetOrderRequest)
    - [GetOrderResponse](#order-v1alpha0-GetOrderResponse)
    - [GetOrderResponse.ShippingAddress](#order-v1alpha0-GetOrderResponse-ShippingAddress)
  
    - [OrderService](#order-v1alpha0-OrderService)
  
- [product/v1alpha0/product.proto](#product_v1alpha0_product-proto)
    - [CreateProductGroupRequest](#product-v1alpha0-CreateProductGroupRequest)
    - [CreateProductGroupResponse](#product-v1alpha0-CreateProductGroupResponse)
    - [CreateProductRequest](#product-v1alpha0-CreateProductRequest)
    - [CreateProductResponse](#product-v1alpha0-CreateProductResponse)
    - [GetGroupProductsRequest](#product-v1alpha0-GetGroupProductsRequest)
    - [GetGroupProductsResponse](#product-v1alpha0-GetGroupProductsResponse)
    - [GetProductBySupplierPartIdRequest](#product-v1alpha0-GetProductBySupplierPartIdRequest)
    - [GetProductBySupplierPartIdResponse](#product-v1alpha0-GetProductBySupplierPartIdResponse)
    - [GetProductGroupsRequest](#product-v1alpha0-GetProductGroupsRequest)
    - [GetProductGroupsResponse](#product-v1alpha0-GetProductGroupsResponse)
    - [GetProductGroupsResponse.ProductGroup](#product-v1alpha0-GetProductGroupsResponse-ProductGroup)
    - [GetProductRequest](#product-v1alpha0-GetProductRequest)
    - [GetProductResponse](#product-v1alpha0-GetProductResponse)
    - [GetProductResponse.Product](#product-v1alpha0-GetProductResponse-Product)
    - [Product](#product-v1alpha0-Product)
    - [ProductGroup](#product-v1alpha0-ProductGroup)
  
    - [ProductService](#product-v1alpha0-ProductService)
  
- [receiving/v1alpha0/receiving.proto](#receiving_v1alpha0_receiving-proto)
    - [GetReceiptRequest](#receiving-v1alpha0-GetReceiptRequest)
    - [GetReceiptResponse](#receiving-v1alpha0-GetReceiptResponse)
    - [GetReceiptResponse.Item](#receiving-v1alpha0-GetReceiptResponse-Item)
    - [GetReceiptResponse.Item.Order](#receiving-v1alpha0-GetReceiptResponse-Item-Order)
  
    - [ReceivingService](#receiving-v1alpha0-ReceivingService)
  
- [shipping/v1alpha0/shipping.proto](#shipping_v1alpha0_shipping-proto)
    - [GetPickTicketRequest](#shipping-v1alpha0-GetPickTicketRequest)
    - [GetPickTicketResponse](#shipping-v1alpha0-GetPickTicketResponse)
  
    - [ShippingService](#shipping-v1alpha0-ShippingService)
  
- [supplier/v1alpha0/supplier.proto](#supplier_v1alpha0_supplier-proto)
    - [SupplierService](#supplier-v1alpha0-SupplierService)
  
- [Scalar Value Types](#scalar-value-types)



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






 

 

 


<a name="inventory-v1alpha0-InventoryService"></a>

### InventoryService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddSupplier | [AddSupplierRequest](#inventory-v1alpha0-AddSupplierRequest) | [AddSupplierResponse](#inventory-v1alpha0-AddSupplierResponse) |  |
| DeleteSupplier | [DeleteSupplierRequest](#inventory-v1alpha0-DeleteSupplierRequest) | [DeleteSupplierResponse](#inventory-v1alpha0-DeleteSupplierResponse) |  |

 



<a name="order_v1alpha0_order-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order/v1alpha0/order.proto



<a name="order-v1alpha0-GetOrderItemRequest"></a>

### GetOrderItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  |  |
| product_id | [int32](#int32) | repeated |  |






<a name="order-v1alpha0-GetOrderItemResponse"></a>

### GetOrderItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [GetOrderItemResponse.Item](#order-v1alpha0-GetOrderItemResponse-Item) | repeated |  |






<a name="order-v1alpha0-GetOrderItemResponse-Item"></a>

### GetOrderItemResponse.Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [int32](#int32) |  |  |
| quantity_ordered | [float](#float) |  |  |
| unit_price | [float](#float) |  |  |
| unit_measurement | [string](#string) |  |  |
| total_price | [float](#float) |  |  |
| customer_product_id | [string](#string) |  |  |






<a name="order-v1alpha0-GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="order-v1alpha0-GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delivery_instructions | [string](#string) |  |  |
| shipping_address | [GetOrderResponse.ShippingAddress](#order-v1alpha0-GetOrderResponse-ShippingAddress) |  |  |
| purchase_order_id | [string](#string) |  |  |
| contact_id | [string](#string) |  |  |
| taker | [string](#string) |  |  |






<a name="order-v1alpha0-GetOrderResponse-ShippingAddress"></a>

### GetOrderResponse.ShippingAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| line_one | [string](#string) |  |  |
| line_two | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| postal_code | [string](#string) |  |  |
| country | [string](#string) |  |  |





 

 

 


<a name="order-v1alpha0-OrderService"></a>

### OrderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOrder | [GetOrderRequest](#order-v1alpha0-GetOrderRequest) | [GetOrderResponse](#order-v1alpha0-GetOrderResponse) |  |
| GetOrderItem | [GetOrderItemRequest](#order-v1alpha0-GetOrderItemRequest) | [GetOrderItemResponse](#order-v1alpha0-GetOrderItemResponse) |  |

 



<a name="product_v1alpha0_product-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## product/v1alpha0/product.proto



<a name="product-v1alpha0-CreateProductGroupRequest"></a>

### CreateProductGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique identifier of Product Group |
| description | [string](#string) |  | Description of group |
| asset_account | [string](#string) |  | Default Asset Account |
| revenue_account | [string](#string) |  | Default Revenue Account |
| cos_account | [string](#string) |  | Default COS Account |






<a name="product-v1alpha0-CreateProductGroupResponse"></a>

### CreateProductGroupResponse







<a name="product-v1alpha0-CreateProductRequest"></a>

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






<a name="product-v1alpha0-CreateProductResponse"></a>

### CreateProductResponse







<a name="product-v1alpha0-GetGroupProductsRequest"></a>

### GetGroupProductsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Product Group ID |






<a name="product-v1alpha0-GetGroupProductsResponse"></a>

### GetGroupProductsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_id | [int32](#int32) | repeated | List of products |






<a name="product-v1alpha0-GetProductBySupplierPartIdRequest"></a>

### GetProductBySupplierPartIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| supplier_id | [double](#double) |  |  |
| supplier_part_id | [string](#string) |  |  |






<a name="product-v1alpha0-GetProductBySupplierPartIdResponse"></a>

### GetProductBySupplierPartIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inv_mast_id | [int32](#int32) |  |  |






<a name="product-v1alpha0-GetProductGroupsRequest"></a>

### GetProductGroupsRequest







<a name="product-v1alpha0-GetProductGroupsResponse"></a>

### GetProductGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product_groups | [GetProductGroupsResponse.ProductGroup](#product-v1alpha0-GetProductGroupsResponse-ProductGroup) | repeated |  |






<a name="product-v1alpha0-GetProductGroupsResponse-ProductGroup"></a>

### GetProductGroupsResponse.ProductGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="product-v1alpha0-GetProductRequest"></a>

### GetProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [int32](#int32) | repeated | list of ProductId&#39;s to fetch |






<a name="product-v1alpha0-GetProductResponse"></a>

### GetProductResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [GetProductResponse.Product](#product-v1alpha0-GetProductResponse-Product) | repeated |  |






<a name="product-v1alpha0-GetProductResponse-Product"></a>

### GetProductResponse.Product



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| product_sn | [string](#string) |  | products item ID usually formatted as MRS-XXX00-123456 |
| name | [string](#string) |  | short description of product such as a name of the item |
| description | [string](#string) |  | extended description of product |
| product_group_id | [string](#string) |  | Product group it belongs to |
| price | [double](#double) |  | Default Price of item |
| qty_available | [double](#double) |  |  |






<a name="product-v1alpha0-Product"></a>

### Product







<a name="product-v1alpha0-ProductGroup"></a>

### ProductGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |





 

 

 


<a name="product-v1alpha0-ProductService"></a>

### ProductService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetProduct | [GetProductRequest](#product-v1alpha0-GetProductRequest) | [GetProductResponse](#product-v1alpha0-GetProductResponse) |  |
| CreateProduct | [CreateProductRequest](#product-v1alpha0-CreateProductRequest) | [CreateProductResponse](#product-v1alpha0-CreateProductResponse) |  |
| GetProductBySupplierPartId | [GetProductBySupplierPartIdRequest](#product-v1alpha0-GetProductBySupplierPartIdRequest) | [GetProductBySupplierPartIdResponse](#product-v1alpha0-GetProductBySupplierPartIdResponse) |  |
| GetProductGroups | [GetProductGroupsRequest](#product-v1alpha0-GetProductGroupsRequest) | [GetProductGroupsResponse](#product-v1alpha0-GetProductGroupsResponse) |  |
| GetGroupProducts | [GetGroupProductsRequest](#product-v1alpha0-GetGroupProductsRequest) | [GetGroupProductsResponse](#product-v1alpha0-GetGroupProductsResponse) |  |
| CreateProductGroup | [CreateProductGroupRequest](#product-v1alpha0-CreateProductGroupRequest) | [CreateProductGroupResponse](#product-v1alpha0-CreateProductGroupResponse) | Creates a new Product Group |

 



<a name="receiving_v1alpha0_receiving-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## receiving/v1alpha0/receiving.proto



<a name="receiving-v1alpha0-GetReceiptRequest"></a>

### GetReceiptRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="receiving-v1alpha0-GetReceiptResponse"></a>

### GetReceiptResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| items | [GetReceiptResponse.Item](#receiving-v1alpha0-GetReceiptResponse-Item) | repeated |  |






<a name="receiving-v1alpha0-GetReceiptResponse-Item"></a>

### GetReceiptResponse.Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| sn | [string](#string) |  |  |
| name | [string](#string) |  |  |
| qty_received | [float](#float) |  |  |
| primary_bin | [string](#string) |  |  |
| allocated_orders | [GetReceiptResponse.Item.Order](#receiving-v1alpha0-GetReceiptResponse-Item-Order) | repeated |  |






<a name="receiving-v1alpha0-GetReceiptResponse-Item-Order"></a>

### GetReceiptResponse.Item.Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Order ID |
| name | [string](#string) |  | Order Name |
| qty | [float](#float) |  | Quantity Allocated |





 

 

 


<a name="receiving-v1alpha0-ReceivingService"></a>

### ReceivingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetReceipt | [GetReceiptRequest](#receiving-v1alpha0-GetReceiptRequest) | [GetReceiptResponse](#receiving-v1alpha0-GetReceiptResponse) |  |

 



<a name="shipping_v1alpha0_shipping-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## shipping/v1alpha0/shipping.proto



<a name="shipping-v1alpha0-GetPickTicketRequest"></a>

### GetPickTicketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="shipping-v1alpha0-GetPickTicketResponse"></a>

### GetPickTicketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ship_name | [string](#string) |  |  |
| ship_line_one | [string](#string) |  |  |
| ship_line_two | [string](#string) |  |  |
| ship_city | [string](#string) |  |  |
| ship_state | [string](#string) |  |  |
| ship_postal_code | [string](#string) |  |  |
| ship_country | [string](#string) |  |  |
| order_id | [string](#string) |  |  |
| order_po | [string](#string) |  |  |
| delivery_instructions | [string](#string) |  |  |
| contact_name | [string](#string) |  |  |
| pick_ticket_id | [int32](#int32) |  |  |





 

 

 


<a name="shipping-v1alpha0-ShippingService"></a>

### ShippingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPickTicket | [GetPickTicketRequest](#shipping-v1alpha0-GetPickTicketRequest) | [GetPickTicketResponse](#shipping-v1alpha0-GetPickTicketResponse) |  |

 



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

