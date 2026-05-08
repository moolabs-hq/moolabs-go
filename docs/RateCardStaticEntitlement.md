# RateCardStaticEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**Type** | **string** |  | 
**Config** | **string** | The JSON parsable config of the entitlement. This value is also returned when checking entitlement access and it is useful for configuring fine-grained access settings to the feature, implemented in your own system. Has to be an object. | 

## Methods

### NewRateCardStaticEntitlement

`func NewRateCardStaticEntitlement(type_ string, config string, ) *RateCardStaticEntitlement`

NewRateCardStaticEntitlement instantiates a new RateCardStaticEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardStaticEntitlementWithDefaults

`func NewRateCardStaticEntitlementWithDefaults() *RateCardStaticEntitlement`

NewRateCardStaticEntitlementWithDefaults instantiates a new RateCardStaticEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RateCardStaticEntitlement) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RateCardStaticEntitlement) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RateCardStaticEntitlement) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RateCardStaticEntitlement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *RateCardStaticEntitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCardStaticEntitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCardStaticEntitlement) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *RateCardStaticEntitlement) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RateCardStaticEntitlement) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RateCardStaticEntitlement) SetConfig(v string)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


