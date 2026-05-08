# DiscountReasonRatecardPercentage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Percentage** | **float64** | The percentage of the discount. | 
**CorrelationId** | Pointer to **string** | Correlation ID for the discount.  This is used to link discounts across different invoices (progressive billing use case).  If not provided, the invoicing engine will auto-generate one. When editing an invoice line, please make sure to keep the same correlation ID of the discount or in progressive billing setups the discount amounts might be incorrect. | [optional] 

## Methods

### NewDiscountReasonRatecardPercentage

`func NewDiscountReasonRatecardPercentage(type_ string, percentage float64, ) *DiscountReasonRatecardPercentage`

NewDiscountReasonRatecardPercentage instantiates a new DiscountReasonRatecardPercentage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountReasonRatecardPercentageWithDefaults

`func NewDiscountReasonRatecardPercentageWithDefaults() *DiscountReasonRatecardPercentage`

NewDiscountReasonRatecardPercentageWithDefaults instantiates a new DiscountReasonRatecardPercentage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DiscountReasonRatecardPercentage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscountReasonRatecardPercentage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscountReasonRatecardPercentage) SetType(v string)`

SetType sets Type field to given value.


### GetPercentage

`func (o *DiscountReasonRatecardPercentage) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *DiscountReasonRatecardPercentage) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *DiscountReasonRatecardPercentage) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.


### GetCorrelationId

`func (o *DiscountReasonRatecardPercentage) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DiscountReasonRatecardPercentage) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DiscountReasonRatecardPercentage) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *DiscountReasonRatecardPercentage) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


