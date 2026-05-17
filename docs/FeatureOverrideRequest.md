# FeatureOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** | Feature key | 
**Currency** | Pointer to **string** | Currency code | [optional] [default to "USD"]
**OverridePriceMicros** | Pointer to **int32** | Override price in micros | [optional] 
**OverrideDiscountBps** | Pointer to **int32** | Override discount in basis points | [optional] 

## Methods

### NewFeatureOverrideRequest

`func NewFeatureOverrideRequest(featureKey string, ) *FeatureOverrideRequest`

NewFeatureOverrideRequest instantiates a new FeatureOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureOverrideRequestWithDefaults

`func NewFeatureOverrideRequestWithDefaults() *FeatureOverrideRequest`

NewFeatureOverrideRequestWithDefaults instantiates a new FeatureOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *FeatureOverrideRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *FeatureOverrideRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *FeatureOverrideRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetCurrency

`func (o *FeatureOverrideRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeatureOverrideRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeatureOverrideRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeatureOverrideRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOverridePriceMicros

`func (o *FeatureOverrideRequest) GetOverridePriceMicros() int32`

GetOverridePriceMicros returns the OverridePriceMicros field if non-nil, zero value otherwise.

### GetOverridePriceMicrosOk

`func (o *FeatureOverrideRequest) GetOverridePriceMicrosOk() (*int32, bool)`

GetOverridePriceMicrosOk returns a tuple with the OverridePriceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePriceMicros

`func (o *FeatureOverrideRequest) SetOverridePriceMicros(v int32)`

SetOverridePriceMicros sets OverridePriceMicros field to given value.

### HasOverridePriceMicros

`func (o *FeatureOverrideRequest) HasOverridePriceMicros() bool`

HasOverridePriceMicros returns a boolean if a field has been set.

### GetOverrideDiscountBps

`func (o *FeatureOverrideRequest) GetOverrideDiscountBps() int32`

GetOverrideDiscountBps returns the OverrideDiscountBps field if non-nil, zero value otherwise.

### GetOverrideDiscountBpsOk

`func (o *FeatureOverrideRequest) GetOverrideDiscountBpsOk() (*int32, bool)`

GetOverrideDiscountBpsOk returns a tuple with the OverrideDiscountBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDiscountBps

`func (o *FeatureOverrideRequest) SetOverrideDiscountBps(v int32)`

SetOverrideDiscountBps sets OverrideDiscountBps field to given value.

### HasOverrideDiscountBps

`func (o *FeatureOverrideRequest) HasOverrideDiscountBps() bool`

HasOverrideDiscountBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


