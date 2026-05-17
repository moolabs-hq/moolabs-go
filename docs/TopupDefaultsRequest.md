# TopupDefaultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceMode** | Pointer to **string** |  | [optional] 
**CooldownHours** | Pointer to **float32** |  | [optional] 
**DailyCap** | Pointer to **int32** |  | [optional] 
**ThresholdTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewTopupDefaultsRequest

`func NewTopupDefaultsRequest() *TopupDefaultsRequest`

NewTopupDefaultsRequest instantiates a new TopupDefaultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopupDefaultsRequestWithDefaults

`func NewTopupDefaultsRequestWithDefaults() *TopupDefaultsRequest`

NewTopupDefaultsRequestWithDefaults instantiates a new TopupDefaultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceMode

`func (o *TopupDefaultsRequest) GetInvoiceMode() string`

GetInvoiceMode returns the InvoiceMode field if non-nil, zero value otherwise.

### GetInvoiceModeOk

`func (o *TopupDefaultsRequest) GetInvoiceModeOk() (*string, bool)`

GetInvoiceModeOk returns a tuple with the InvoiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMode

`func (o *TopupDefaultsRequest) SetInvoiceMode(v string)`

SetInvoiceMode sets InvoiceMode field to given value.

### HasInvoiceMode

`func (o *TopupDefaultsRequest) HasInvoiceMode() bool`

HasInvoiceMode returns a boolean if a field has been set.

### GetCooldownHours

`func (o *TopupDefaultsRequest) GetCooldownHours() float32`

GetCooldownHours returns the CooldownHours field if non-nil, zero value otherwise.

### GetCooldownHoursOk

`func (o *TopupDefaultsRequest) GetCooldownHoursOk() (*float32, bool)`

GetCooldownHoursOk returns a tuple with the CooldownHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownHours

`func (o *TopupDefaultsRequest) SetCooldownHours(v float32)`

SetCooldownHours sets CooldownHours field to given value.

### HasCooldownHours

`func (o *TopupDefaultsRequest) HasCooldownHours() bool`

HasCooldownHours returns a boolean if a field has been set.

### GetDailyCap

`func (o *TopupDefaultsRequest) GetDailyCap() int32`

GetDailyCap returns the DailyCap field if non-nil, zero value otherwise.

### GetDailyCapOk

`func (o *TopupDefaultsRequest) GetDailyCapOk() (*int32, bool)`

GetDailyCapOk returns a tuple with the DailyCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCap

`func (o *TopupDefaultsRequest) SetDailyCap(v int32)`

SetDailyCap sets DailyCap field to given value.

### HasDailyCap

`func (o *TopupDefaultsRequest) HasDailyCap() bool`

HasDailyCap returns a boolean if a field has been set.

### GetThresholdTemplate

`func (o *TopupDefaultsRequest) GetThresholdTemplate() string`

GetThresholdTemplate returns the ThresholdTemplate field if non-nil, zero value otherwise.

### GetThresholdTemplateOk

`func (o *TopupDefaultsRequest) GetThresholdTemplateOk() (*string, bool)`

GetThresholdTemplateOk returns a tuple with the ThresholdTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdTemplate

`func (o *TopupDefaultsRequest) SetThresholdTemplate(v string)`

SetThresholdTemplate sets ThresholdTemplate field to given value.

### HasThresholdTemplate

`func (o *TopupDefaultsRequest) HasThresholdTemplate() bool`

HasThresholdTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


