# DiscountReasonRatecardUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Quantity** | **string** | The quantity of the usage discount.  Must be positive. | 
**CorrelationId** | Pointer to **string** | Correlation ID for the discount.  This is used to link discounts across different invoices (progressive billing use case).  If not provided, the invoicing engine will auto-generate one. When editing an invoice line, please make sure to keep the same correlation ID of the discount or in progressive billing setups the discount amounts might be incorrect. | [optional] 

## Methods

### NewDiscountReasonRatecardUsage

`func NewDiscountReasonRatecardUsage(type_ string, quantity string, ) *DiscountReasonRatecardUsage`

NewDiscountReasonRatecardUsage instantiates a new DiscountReasonRatecardUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountReasonRatecardUsageWithDefaults

`func NewDiscountReasonRatecardUsageWithDefaults() *DiscountReasonRatecardUsage`

NewDiscountReasonRatecardUsageWithDefaults instantiates a new DiscountReasonRatecardUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscountReasonRatecardUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscountReasonRatecardUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscountReasonRatecardUsage) SetType(v string)`

SetType sets Type field to given value.


### GetQuantity

`func (o *DiscountReasonRatecardUsage) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DiscountReasonRatecardUsage) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DiscountReasonRatecardUsage) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetCorrelationId

`func (o *DiscountReasonRatecardUsage) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DiscountReasonRatecardUsage) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DiscountReasonRatecardUsage) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *DiscountReasonRatecardUsage) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


