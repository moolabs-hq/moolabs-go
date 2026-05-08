# MappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**MeterSlug** | **string** |  | 
**FeatureKey** | **string** |  | 
**EffectiveFrom** | **string** |  | 
**EffectiveTo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMappingResponse

`func NewMappingResponse(id string, tenantId string, meterSlug string, featureKey string, effectiveFrom string, ) *MappingResponse`

NewMappingResponse instantiates a new MappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingResponseWithDefaults

`func NewMappingResponseWithDefaults() *MappingResponse`

NewMappingResponseWithDefaults instantiates a new MappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MappingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MappingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MappingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *MappingResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MappingResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MappingResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetMeterSlug

`func (o *MappingResponse) GetMeterSlug() string`

GetMeterSlug returns the MeterSlug field if non-nil, zero value otherwise.

### GetMeterSlugOk

`func (o *MappingResponse) GetMeterSlugOk() (*string, bool)`

GetMeterSlugOk returns a tuple with the MeterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSlug

`func (o *MappingResponse) SetMeterSlug(v string)`

SetMeterSlug sets MeterSlug field to given value.


### GetFeatureKey

`func (o *MappingResponse) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *MappingResponse) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *MappingResponse) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetEffectiveFrom

`func (o *MappingResponse) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *MappingResponse) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *MappingResponse) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.


### GetEffectiveTo

`func (o *MappingResponse) GetEffectiveTo() string`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *MappingResponse) GetEffectiveToOk() (*string, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *MappingResponse) SetEffectiveTo(v string)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *MappingResponse) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### SetEffectiveToNil

`func (o *MappingResponse) SetEffectiveToNil(b bool)`

 SetEffectiveToNil sets the value for EffectiveTo to be an explicit nil

### UnsetEffectiveTo
`func (o *MappingResponse) UnsetEffectiveTo()`

UnsetEffectiveTo ensures that no value is present for EffectiveTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


