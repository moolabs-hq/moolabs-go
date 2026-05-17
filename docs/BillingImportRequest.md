# BillingImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**TotalCost** | [**TotalCost**](TotalCost.md) |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**LineItems** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RawInvoiceRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingImportRequest

`func NewBillingImportRequest(provider string, periodStart string, periodEnd string, totalCost TotalCost, ) *BillingImportRequest`

NewBillingImportRequest instantiates a new BillingImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingImportRequestWithDefaults

`func NewBillingImportRequestWithDefaults() *BillingImportRequest`

NewBillingImportRequestWithDefaults instantiates a new BillingImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *BillingImportRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BillingImportRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BillingImportRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPeriodStart

`func (o *BillingImportRequest) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BillingImportRequest) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BillingImportRequest) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *BillingImportRequest) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BillingImportRequest) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BillingImportRequest) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetTotalCost

`func (o *BillingImportRequest) GetTotalCost() TotalCost`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *BillingImportRequest) GetTotalCostOk() (*TotalCost, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *BillingImportRequest) SetTotalCost(v TotalCost)`

SetTotalCost sets TotalCost field to given value.


### GetCurrency

`func (o *BillingImportRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingImportRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingImportRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingImportRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLineItems

`func (o *BillingImportRequest) GetLineItems() []map[string]interface{}`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BillingImportRequest) GetLineItemsOk() (*[]map[string]interface{}, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BillingImportRequest) SetLineItems(v []map[string]interface{})`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BillingImportRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetRawInvoiceRef

`func (o *BillingImportRequest) GetRawInvoiceRef() string`

GetRawInvoiceRef returns the RawInvoiceRef field if non-nil, zero value otherwise.

### GetRawInvoiceRefOk

`func (o *BillingImportRequest) GetRawInvoiceRefOk() (*string, bool)`

GetRawInvoiceRefOk returns a tuple with the RawInvoiceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawInvoiceRef

`func (o *BillingImportRequest) SetRawInvoiceRef(v string)`

SetRawInvoiceRef sets RawInvoiceRef field to given value.

### HasRawInvoiceRef

`func (o *BillingImportRequest) HasRawInvoiceRef() bool`

HasRawInvoiceRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


