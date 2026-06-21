# PricingSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**QuoteId** | **string** |  | 
**QuoteVersion** | **int32** |  | 
**PricingSnapshotId** | **string** |  | 
**PricedOutput** | **map[string]interface{}** |  | 
**PricingTrace** | **map[string]interface{}** |  | 
**PricingTraceDigest** | **string** |  | 
**SourceVersions** | **map[string]interface{}** |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**SubtotalMicros** | **int32** |  | 
**DiscountMicros** | **int32** |  | 
**TotalMicros** | **int32** |  | 

## Methods

### NewPricingSnapshotResponse

`func NewPricingSnapshotResponse(id string, quoteId string, quoteVersion int32, pricingSnapshotId string, pricedOutput map[string]interface{}, pricingTrace map[string]interface{}, pricingTraceDigest string, sourceVersions map[string]interface{}, subtotalMicros int32, discountMicros int32, totalMicros int32, ) *PricingSnapshotResponse`

NewPricingSnapshotResponse instantiates a new PricingSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSnapshotResponseWithDefaults

`func NewPricingSnapshotResponseWithDefaults() *PricingSnapshotResponse`

NewPricingSnapshotResponseWithDefaults instantiates a new PricingSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingSnapshotResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingSnapshotResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingSnapshotResponse) SetId(v string)`

SetId sets Id field to given value.


### GetQuoteId

`func (o *PricingSnapshotResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PricingSnapshotResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PricingSnapshotResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *PricingSnapshotResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *PricingSnapshotResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *PricingSnapshotResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetPricingSnapshotId

`func (o *PricingSnapshotResponse) GetPricingSnapshotId() string`

GetPricingSnapshotId returns the PricingSnapshotId field if non-nil, zero value otherwise.

### GetPricingSnapshotIdOk

`func (o *PricingSnapshotResponse) GetPricingSnapshotIdOk() (*string, bool)`

GetPricingSnapshotIdOk returns a tuple with the PricingSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSnapshotId

`func (o *PricingSnapshotResponse) SetPricingSnapshotId(v string)`

SetPricingSnapshotId sets PricingSnapshotId field to given value.


### GetPricedOutput

`func (o *PricingSnapshotResponse) GetPricedOutput() map[string]interface{}`

GetPricedOutput returns the PricedOutput field if non-nil, zero value otherwise.

### GetPricedOutputOk

`func (o *PricingSnapshotResponse) GetPricedOutputOk() (*map[string]interface{}, bool)`

GetPricedOutputOk returns a tuple with the PricedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricedOutput

`func (o *PricingSnapshotResponse) SetPricedOutput(v map[string]interface{})`

SetPricedOutput sets PricedOutput field to given value.


### GetPricingTrace

`func (o *PricingSnapshotResponse) GetPricingTrace() map[string]interface{}`

GetPricingTrace returns the PricingTrace field if non-nil, zero value otherwise.

### GetPricingTraceOk

`func (o *PricingSnapshotResponse) GetPricingTraceOk() (*map[string]interface{}, bool)`

GetPricingTraceOk returns a tuple with the PricingTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTrace

`func (o *PricingSnapshotResponse) SetPricingTrace(v map[string]interface{})`

SetPricingTrace sets PricingTrace field to given value.


### GetPricingTraceDigest

`func (o *PricingSnapshotResponse) GetPricingTraceDigest() string`

GetPricingTraceDigest returns the PricingTraceDigest field if non-nil, zero value otherwise.

### GetPricingTraceDigestOk

`func (o *PricingSnapshotResponse) GetPricingTraceDigestOk() (*string, bool)`

GetPricingTraceDigestOk returns a tuple with the PricingTraceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTraceDigest

`func (o *PricingSnapshotResponse) SetPricingTraceDigest(v string)`

SetPricingTraceDigest sets PricingTraceDigest field to given value.


### GetSourceVersions

`func (o *PricingSnapshotResponse) GetSourceVersions() map[string]interface{}`

GetSourceVersions returns the SourceVersions field if non-nil, zero value otherwise.

### GetSourceVersionsOk

`func (o *PricingSnapshotResponse) GetSourceVersionsOk() (*map[string]interface{}, bool)`

GetSourceVersionsOk returns a tuple with the SourceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersions

`func (o *PricingSnapshotResponse) SetSourceVersions(v map[string]interface{})`

SetSourceVersions sets SourceVersions field to given value.


### GetCurrency

`func (o *PricingSnapshotResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PricingSnapshotResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PricingSnapshotResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PricingSnapshotResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSubtotalMicros

`func (o *PricingSnapshotResponse) GetSubtotalMicros() int32`

GetSubtotalMicros returns the SubtotalMicros field if non-nil, zero value otherwise.

### GetSubtotalMicrosOk

`func (o *PricingSnapshotResponse) GetSubtotalMicrosOk() (*int32, bool)`

GetSubtotalMicrosOk returns a tuple with the SubtotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalMicros

`func (o *PricingSnapshotResponse) SetSubtotalMicros(v int32)`

SetSubtotalMicros sets SubtotalMicros field to given value.


### GetDiscountMicros

`func (o *PricingSnapshotResponse) GetDiscountMicros() int32`

GetDiscountMicros returns the DiscountMicros field if non-nil, zero value otherwise.

### GetDiscountMicrosOk

`func (o *PricingSnapshotResponse) GetDiscountMicrosOk() (*int32, bool)`

GetDiscountMicrosOk returns a tuple with the DiscountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMicros

`func (o *PricingSnapshotResponse) SetDiscountMicros(v int32)`

SetDiscountMicros sets DiscountMicros field to given value.


### GetTotalMicros

`func (o *PricingSnapshotResponse) GetTotalMicros() int32`

GetTotalMicros returns the TotalMicros field if non-nil, zero value otherwise.

### GetTotalMicrosOk

`func (o *PricingSnapshotResponse) GetTotalMicrosOk() (*int32, bool)`

GetTotalMicrosOk returns a tuple with the TotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMicros

`func (o *PricingSnapshotResponse) SetTotalMicros(v int32)`

SetTotalMicros sets TotalMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


