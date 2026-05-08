# CreateMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeterSlug** | **string** | Meter slug from OpenMeter | 
**FeatureKey** | **string** | Feature key to map to | 
**EffectiveAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateMappingRequest

`func NewCreateMappingRequest(meterSlug string, featureKey string, ) *CreateMappingRequest`

NewCreateMappingRequest instantiates a new CreateMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMappingRequestWithDefaults

`func NewCreateMappingRequestWithDefaults() *CreateMappingRequest`

NewCreateMappingRequestWithDefaults instantiates a new CreateMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterSlug

`func (o *CreateMappingRequest) GetMeterSlug() string`

GetMeterSlug returns the MeterSlug field if non-nil, zero value otherwise.

### GetMeterSlugOk

`func (o *CreateMappingRequest) GetMeterSlugOk() (*string, bool)`

GetMeterSlugOk returns a tuple with the MeterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSlug

`func (o *CreateMappingRequest) SetMeterSlug(v string)`

SetMeterSlug sets MeterSlug field to given value.


### GetFeatureKey

`func (o *CreateMappingRequest) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *CreateMappingRequest) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *CreateMappingRequest) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetEffectiveAt

`func (o *CreateMappingRequest) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *CreateMappingRequest) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *CreateMappingRequest) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *CreateMappingRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### SetEffectiveAtNil

`func (o *CreateMappingRequest) SetEffectiveAtNil(b bool)`

 SetEffectiveAtNil sets the value for EffectiveAt to be an explicit nil

### UnsetEffectiveAt
`func (o *CreateMappingRequest) UnsetEffectiveAt()`

UnsetEffectiveAt ensures that no value is present for EffectiveAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


