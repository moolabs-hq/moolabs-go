# FeaturePriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** | Feature key | 
**Currency** | Pointer to **string** | Currency code | [optional] [default to "USD"]
**ListPriceMicros** | **int32** | List price in micros | 
**NetPriceMicros** | **int32** | Net price in micros (after discounts) | 

## Methods

### NewFeaturePriceRequest

`func NewFeaturePriceRequest(featureKey string, listPriceMicros int32, netPriceMicros int32, ) *FeaturePriceRequest`

NewFeaturePriceRequest instantiates a new FeaturePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturePriceRequestWithDefaults

`func NewFeaturePriceRequestWithDefaults() *FeaturePriceRequest`

NewFeaturePriceRequestWithDefaults instantiates a new FeaturePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *FeaturePriceRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *FeaturePriceRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *FeaturePriceRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetCurrency

`func (o *FeaturePriceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeaturePriceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeaturePriceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeaturePriceRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetListPriceMicros

`func (o *FeaturePriceRequest) GetListPriceMicros() int32`

GetListPriceMicros returns the ListPriceMicros field if non-nil, zero value otherwise.

### GetListPriceMicrosOk

`func (o *FeaturePriceRequest) GetListPriceMicrosOk() (*int32, bool)`

GetListPriceMicrosOk returns a tuple with the ListPriceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPriceMicros

`func (o *FeaturePriceRequest) SetListPriceMicros(v int32)`

SetListPriceMicros sets ListPriceMicros field to given value.


### GetNetPriceMicros

`func (o *FeaturePriceRequest) GetNetPriceMicros() int32`

GetNetPriceMicros returns the NetPriceMicros field if non-nil, zero value otherwise.

### GetNetPriceMicrosOk

`func (o *FeaturePriceRequest) GetNetPriceMicrosOk() (*int32, bool)`

GetNetPriceMicrosOk returns a tuple with the NetPriceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPriceMicros

`func (o *FeaturePriceRequest) SetNetPriceMicros(v int32)`

SetNetPriceMicros sets NetPriceMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


