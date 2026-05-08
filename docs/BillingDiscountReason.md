# BillingDiscountReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Percentage** | **float64** | The percentage of the discount. | 
**CorrelationId** | Pointer to **string** | Correlation ID for the discount.  This is used to link discounts across different invoices (progressive billing use case).  If not provided, the invoicing engine will auto-generate one. When editing an invoice line, please make sure to keep the same correlation ID of the discount or in progressive billing setups the discount amounts might be incorrect. | [optional] 
**Quantity** | **string** | The quantity of the usage discount.  Must be positive. | 

## Methods

### NewBillingDiscountReason

`func NewBillingDiscountReason(type_ string, percentage float64, quantity string, ) *BillingDiscountReason`

NewBillingDiscountReason instantiates a new BillingDiscountReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDiscountReasonWithDefaults

`func NewBillingDiscountReasonWithDefaults() *BillingDiscountReason`

NewBillingDiscountReasonWithDefaults instantiates a new BillingDiscountReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingDiscountReason) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingDiscountReason) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingDiscountReason) SetType(v string)`

SetType sets Type field to given value.


### GetPercentage

`func (o *BillingDiscountReason) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *BillingDiscountReason) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *BillingDiscountReason) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.


### GetCorrelationId

`func (o *BillingDiscountReason) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *BillingDiscountReason) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *BillingDiscountReason) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *BillingDiscountReason) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingDiscountReason) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingDiscountReason) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingDiscountReason) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


