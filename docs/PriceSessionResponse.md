# PriceSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**QuoteVersion** | **int32** |  | 
**PricingSnapshotId** | **string** |  | 
**NormalizedInput** | **map[string]interface{}** |  | 
**PricedOutput** | **map[string]interface{}** |  | 
**PricingTrace** | **map[string]interface{}** |  | 
**PricingTraceDigest** | **string** |  | 
**SourceVersions** | **map[string]interface{}** |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**SubtotalMicros** | **int32** |  | 
**DiscountMicros** | **int32** |  | 
**TotalMicros** | **int32** |  | 

## Methods

### NewPriceSessionResponse

`func NewPriceSessionResponse(quoteId string, quoteVersion int32, pricingSnapshotId string, normalizedInput map[string]interface{}, pricedOutput map[string]interface{}, pricingTrace map[string]interface{}, pricingTraceDigest string, sourceVersions map[string]interface{}, subtotalMicros int32, discountMicros int32, totalMicros int32, ) *PriceSessionResponse`

NewPriceSessionResponse instantiates a new PriceSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceSessionResponseWithDefaults

`func NewPriceSessionResponseWithDefaults() *PriceSessionResponse`

NewPriceSessionResponseWithDefaults instantiates a new PriceSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *PriceSessionResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *PriceSessionResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *PriceSessionResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *PriceSessionResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *PriceSessionResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *PriceSessionResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetPricingSnapshotId

`func (o *PriceSessionResponse) GetPricingSnapshotId() string`

GetPricingSnapshotId returns the PricingSnapshotId field if non-nil, zero value otherwise.

### GetPricingSnapshotIdOk

`func (o *PriceSessionResponse) GetPricingSnapshotIdOk() (*string, bool)`

GetPricingSnapshotIdOk returns a tuple with the PricingSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSnapshotId

`func (o *PriceSessionResponse) SetPricingSnapshotId(v string)`

SetPricingSnapshotId sets PricingSnapshotId field to given value.


### GetNormalizedInput

`func (o *PriceSessionResponse) GetNormalizedInput() map[string]interface{}`

GetNormalizedInput returns the NormalizedInput field if non-nil, zero value otherwise.

### GetNormalizedInputOk

`func (o *PriceSessionResponse) GetNormalizedInputOk() (*map[string]interface{}, bool)`

GetNormalizedInputOk returns a tuple with the NormalizedInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedInput

`func (o *PriceSessionResponse) SetNormalizedInput(v map[string]interface{})`

SetNormalizedInput sets NormalizedInput field to given value.


### GetPricedOutput

`func (o *PriceSessionResponse) GetPricedOutput() map[string]interface{}`

GetPricedOutput returns the PricedOutput field if non-nil, zero value otherwise.

### GetPricedOutputOk

`func (o *PriceSessionResponse) GetPricedOutputOk() (*map[string]interface{}, bool)`

GetPricedOutputOk returns a tuple with the PricedOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricedOutput

`func (o *PriceSessionResponse) SetPricedOutput(v map[string]interface{})`

SetPricedOutput sets PricedOutput field to given value.


### GetPricingTrace

`func (o *PriceSessionResponse) GetPricingTrace() map[string]interface{}`

GetPricingTrace returns the PricingTrace field if non-nil, zero value otherwise.

### GetPricingTraceOk

`func (o *PriceSessionResponse) GetPricingTraceOk() (*map[string]interface{}, bool)`

GetPricingTraceOk returns a tuple with the PricingTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTrace

`func (o *PriceSessionResponse) SetPricingTrace(v map[string]interface{})`

SetPricingTrace sets PricingTrace field to given value.


### GetPricingTraceDigest

`func (o *PriceSessionResponse) GetPricingTraceDigest() string`

GetPricingTraceDigest returns the PricingTraceDigest field if non-nil, zero value otherwise.

### GetPricingTraceDigestOk

`func (o *PriceSessionResponse) GetPricingTraceDigestOk() (*string, bool)`

GetPricingTraceDigestOk returns a tuple with the PricingTraceDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingTraceDigest

`func (o *PriceSessionResponse) SetPricingTraceDigest(v string)`

SetPricingTraceDigest sets PricingTraceDigest field to given value.


### GetSourceVersions

`func (o *PriceSessionResponse) GetSourceVersions() map[string]interface{}`

GetSourceVersions returns the SourceVersions field if non-nil, zero value otherwise.

### GetSourceVersionsOk

`func (o *PriceSessionResponse) GetSourceVersionsOk() (*map[string]interface{}, bool)`

GetSourceVersionsOk returns a tuple with the SourceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersions

`func (o *PriceSessionResponse) SetSourceVersions(v map[string]interface{})`

SetSourceVersions sets SourceVersions field to given value.


### GetCurrency

`func (o *PriceSessionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PriceSessionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PriceSessionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PriceSessionResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSubtotalMicros

`func (o *PriceSessionResponse) GetSubtotalMicros() int32`

GetSubtotalMicros returns the SubtotalMicros field if non-nil, zero value otherwise.

### GetSubtotalMicrosOk

`func (o *PriceSessionResponse) GetSubtotalMicrosOk() (*int32, bool)`

GetSubtotalMicrosOk returns a tuple with the SubtotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalMicros

`func (o *PriceSessionResponse) SetSubtotalMicros(v int32)`

SetSubtotalMicros sets SubtotalMicros field to given value.


### GetDiscountMicros

`func (o *PriceSessionResponse) GetDiscountMicros() int32`

GetDiscountMicros returns the DiscountMicros field if non-nil, zero value otherwise.

### GetDiscountMicrosOk

`func (o *PriceSessionResponse) GetDiscountMicrosOk() (*int32, bool)`

GetDiscountMicrosOk returns a tuple with the DiscountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMicros

`func (o *PriceSessionResponse) SetDiscountMicros(v int32)`

SetDiscountMicros sets DiscountMicros field to given value.


### GetTotalMicros

`func (o *PriceSessionResponse) GetTotalMicros() int32`

GetTotalMicros returns the TotalMicros field if non-nil, zero value otherwise.

### GetTotalMicrosOk

`func (o *PriceSessionResponse) GetTotalMicrosOk() (*int32, bool)`

GetTotalMicrosOk returns a tuple with the TotalMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMicros

`func (o *PriceSessionResponse) SetTotalMicros(v int32)`

SetTotalMicros sets TotalMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


