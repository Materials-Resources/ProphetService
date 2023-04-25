# Protocol Documentation
<a name="top"></a>

## Table of Contents

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
    - [GetProductBySupplierPartIdRequest](#product-v1alpha0-GetProductBySupplierPartIdRequest)
    - [GetProductBySupplierPartIdResponse](#product-v1alpha0-GetProductBySupplierPartIdResponse)
    - [GetProductGroupsRequest](#product-v1alpha0-GetProductGroupsRequest)
    - [GetProductGroupsResponse](#product-v1alpha0-GetProductGroupsResponse)
    - [GetProductRequest](#product-v1alpha0-GetProductRequest)
    - [GetProductResponse](#product-v1alpha0-GetProductResponse)
    - [ProductGroup](#product-v1alpha0-ProductGroup)
  
    - [ProductService](#product-v1alpha0-ProductService)
  
- [warehouse/v1alpha0/warehouse.proto](#warehouse_v1alpha0_warehouse-proto)
    - [GetAllocatedOrdersForReceiptItemRequest](#warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemRequest)
    - [GetAllocatedOrdersForReceiptItemResponse](#warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemResponse)
    - [GetAllocatedOrdersForReceiptItemResponse.AllocatedOrder](#warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemResponse-AllocatedOrder)
    - [GetPickTicketRequest](#warehouse-v1alpha0-GetPickTicketRequest)
    - [GetPickTicketResponse](#warehouse-v1alpha0-GetPickTicketResponse)
    - [GetReceivingReportRequest](#warehouse-v1alpha0-GetReceivingReportRequest)
    - [GetReceivingReportResponse](#warehouse-v1alpha0-GetReceivingReportResponse)
    - [GetReceivingReportResponse.Product](#warehouse-v1alpha0-GetReceivingReportResponse-Product)
  
    - [WarehouseService](#warehouse-v1alpha0-WarehouseService)
  
- [Scalar Value Types](#scalar-value-types)



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
| po_no | [string](#string) |  |  |
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
| product_groups | [ProductGroup](#product-v1alpha0-ProductGroup) | repeated |  |






<a name="product-v1alpha0-GetProductRequest"></a>

### GetProductRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inv_mast_uid | [int32](#int32) |  | uid of product that we want to get |






<a name="product-v1alpha0-GetProductResponse"></a>

### GetProductResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inv_mast_uid | [int32](#int32) |  |  |
| item_id | [string](#string) |  | products item ID usually formatted as MRS-XXX00-123456 |
| item_desc | [string](#string) |  | short description of product such as a name of the item |
| extended_desc | [string](#string) |  | extended description of product |
| product_group_id | [string](#string) |  | Product group it belongs to |
| price | [double](#double) |  | Default Price of item |
| qty_available | [double](#double) |  |  |






<a name="product-v1alpha0-ProductGroup"></a>

### ProductGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |





 

 

 


<a name="product-v1alpha0-ProductService"></a>

### ProductService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetProduct | [GetProductRequest](#product-v1alpha0-GetProductRequest) | [GetProductResponse](#product-v1alpha0-GetProductResponse) |  |
| CreateProduct | [CreateProductRequest](#product-v1alpha0-CreateProductRequest) | [CreateProductResponse](#product-v1alpha0-CreateProductResponse) |  |
| GetProductBySupplierPartId | [GetProductBySupplierPartIdRequest](#product-v1alpha0-GetProductBySupplierPartIdRequest) | [GetProductBySupplierPartIdResponse](#product-v1alpha0-GetProductBySupplierPartIdResponse) |  |
| GetProductGroups | [GetProductGroupsRequest](#product-v1alpha0-GetProductGroupsRequest) | [GetProductGroupsResponse](#product-v1alpha0-GetProductGroupsResponse) |  |
| CreateProductGroup | [CreateProductGroupRequest](#product-v1alpha0-CreateProductGroupRequest) | [CreateProductGroupResponse](#product-v1alpha0-CreateProductGroupResponse) | Creates a new Product Group |

 



<a name="warehouse_v1alpha0_warehouse-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## warehouse/v1alpha0/warehouse.proto



<a name="warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemRequest"></a>

### GetAllocatedOrdersForReceiptItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| receipt_id | [int32](#int32) |  | Receipt id to search for |
| product_id | [int32](#int32) |  | Identifier of product |
| primary_bins | [string](#string) | repeated | List of primary bins to filter products by. Helpful for only returning orders for nonstock |






<a name="warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemResponse"></a>

### GetAllocatedOrdersForReceiptItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allocated_orders | [GetAllocatedOrdersForReceiptItemResponse.AllocatedOrder](#warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemResponse-AllocatedOrder) | repeated |  |






<a name="warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemResponse-AllocatedOrder"></a>

### GetAllocatedOrdersForReceiptItemResponse.AllocatedOrder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Identifier of allocated order |
| qty | [float](#float) |  | Quantity of product allocated to order |






<a name="warehouse-v1alpha0-GetPickTicketRequest"></a>

### GetPickTicketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Identifier of pick ticket |






<a name="warehouse-v1alpha0-GetPickTicketResponse"></a>

### GetPickTicketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  | Order Identifier belonging to pick ticket |






<a name="warehouse-v1alpha0-GetReceivingReportRequest"></a>

### GetReceivingReportRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Identifier of receipt |






<a name="warehouse-v1alpha0-GetReceivingReportResponse"></a>

### GetReceivingReportResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [GetReceivingReportResponse.Product](#warehouse-v1alpha0-GetReceivingReportResponse-Product) | repeated |  |






<a name="warehouse-v1alpha0-GetReceivingReportResponse-Product"></a>

### GetReceivingReportResponse.Product



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Identifier of product |
| qty | [float](#float) |  | Quantity of product received |
| unit_of_measure | [string](#string) |  |  |





 

 

 


<a name="warehouse-v1alpha0-WarehouseService"></a>

### WarehouseService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAllocatedOrdersForReceiptItem | [GetAllocatedOrdersForReceiptItemRequest](#warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemRequest) | [GetAllocatedOrdersForReceiptItemResponse](#warehouse-v1alpha0-GetAllocatedOrdersForReceiptItemResponse) |  |
| GetReceivingReport | [GetReceivingReportRequest](#warehouse-v1alpha0-GetReceivingReportRequest) | [GetReceivingReportResponse](#warehouse-v1alpha0-GetReceivingReportResponse) |  |
| GetPickTicket | [GetPickTicketRequest](#warehouse-v1alpha0-GetPickTicketRequest) | [GetPickTicketResponse](#warehouse-v1alpha0-GetPickTicketResponse) |  |

 



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

