# EntitlementPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementId** | **string** |  | 
**FeatureKey** | **string** |  | 
**EntitlementType** | **string** |  | 
**Amount** | Pointer to [**NullableAmount1**](Amount1.md) |  | [optional] 
**RateCard** | Pointer to [**RateCardPayload**](RateCardPayload.md) |  | [optional] 

## Methods

### NewEntitlementPayload

`func NewEntitlementPayload(entitlementId string, featureKey string, entitlementType string, ) *EntitlementPayload`

NewEntitlementPayload instantiates a new EntitlementPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementPayloadWithDefaults

`func NewEntitlementPayloadWithDefaults() *EntitlementPayload`

NewEntitlementPayloadWithDefaults instantiates a new EntitlementPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementId

`func (o *EntitlementPayload) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementPayload) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementPayload) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetFeatureKey

`func (o *EntitlementPayload) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementPayload) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementPayload) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetEntitlementType

`func (o *EntitlementPayload) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *EntitlementPayload) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *EntitlementPayload) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetAmount

`func (o *EntitlementPayload) GetAmount() Amount1`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EntitlementPayload) GetAmountOk() (*Amount1, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EntitlementPayload) SetAmount(v Amount1)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EntitlementPayload) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *EntitlementPayload) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *EntitlementPayload) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetRateCard

`func (o *EntitlementPayload) GetRateCard() RateCardPayload`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *EntitlementPayload) GetRateCardOk() (*RateCardPayload, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *EntitlementPayload) SetRateCard(v RateCardPayload)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *EntitlementPayload) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


