# BillingImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Provider** | **string** |  | 
**BillingPeriodStart** | **string** |  | 
**BillingPeriodEnd** | **string** |  | 
**TotalCost** | **float32** |  | 
**Currency** | **string** |  | 
**ImportSource** | **string** |  | 
**ImportedAt** | **time.Time** |  | 

## Methods

### NewBillingImportResponse

`func NewBillingImportResponse(id string, tenantId string, provider string, billingPeriodStart string, billingPeriodEnd string, totalCost float32, currency string, importSource string, importedAt time.Time, ) *BillingImportResponse`

NewBillingImportResponse instantiates a new BillingImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingImportResponseWithDefaults

`func NewBillingImportResponseWithDefaults() *BillingImportResponse`

NewBillingImportResponseWithDefaults instantiates a new BillingImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingImportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingImportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingImportResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *BillingImportResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingImportResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingImportResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetProvider

`func (o *BillingImportResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BillingImportResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BillingImportResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetBillingPeriodStart

`func (o *BillingImportResponse) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *BillingImportResponse) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *BillingImportResponse) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.


### GetBillingPeriodEnd

`func (o *BillingImportResponse) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *BillingImportResponse) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *BillingImportResponse) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.


### GetTotalCost

`func (o *BillingImportResponse) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *BillingImportResponse) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *BillingImportResponse) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### GetCurrency

`func (o *BillingImportResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingImportResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingImportResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetImportSource

`func (o *BillingImportResponse) GetImportSource() string`

GetImportSource returns the ImportSource field if non-nil, zero value otherwise.

### GetImportSourceOk

`func (o *BillingImportResponse) GetImportSourceOk() (*string, bool)`

GetImportSourceOk returns a tuple with the ImportSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSource

`func (o *BillingImportResponse) SetImportSource(v string)`

SetImportSource sets ImportSource field to given value.


### GetImportedAt

`func (o *BillingImportResponse) GetImportedAt() time.Time`

GetImportedAt returns the ImportedAt field if non-nil, zero value otherwise.

### GetImportedAtOk

`func (o *BillingImportResponse) GetImportedAtOk() (*time.Time, bool)`

GetImportedAtOk returns a tuple with the ImportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedAt

`func (o *BillingImportResponse) SetImportedAt(v time.Time)`

SetImportedAt sets ImportedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


