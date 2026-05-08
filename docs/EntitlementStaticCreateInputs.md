# EntitlementStaticCreateInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**FeatureId** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**UsagePeriod** | Pointer to [**RecurringPeriodCreateInput**](RecurringPeriodCreateInput.md) | The usage period associated with the entitlement. | [optional] 
**Type** | **string** |  | 
**Config** | **string** | The JSON parsable config of the entitlement. This value is also returned when checking entitlement access and it is useful for configuring fine-grained access settings to the feature, implemented in your own system. Has to be an object. | 

## Methods

### NewEntitlementStaticCreateInputs

`func NewEntitlementStaticCreateInputs(type_ string, config string, ) *EntitlementStaticCreateInputs`

NewEntitlementStaticCreateInputs instantiates a new EntitlementStaticCreateInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementStaticCreateInputsWithDefaults

`func NewEntitlementStaticCreateInputsWithDefaults() *EntitlementStaticCreateInputs`

NewEntitlementStaticCreateInputsWithDefaults instantiates a new EntitlementStaticCreateInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *EntitlementStaticCreateInputs) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementStaticCreateInputs) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementStaticCreateInputs) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *EntitlementStaticCreateInputs) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetFeatureId

`func (o *EntitlementStaticCreateInputs) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementStaticCreateInputs) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementStaticCreateInputs) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *EntitlementStaticCreateInputs) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementStaticCreateInputs) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementStaticCreateInputs) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementStaticCreateInputs) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementStaticCreateInputs) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *EntitlementStaticCreateInputs) GetUsagePeriod() RecurringPeriodCreateInput`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementStaticCreateInputs) GetUsagePeriodOk() (*RecurringPeriodCreateInput, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementStaticCreateInputs) SetUsagePeriod(v RecurringPeriodCreateInput)`

SetUsagePeriod sets UsagePeriod field to given value.

### HasUsagePeriod

`func (o *EntitlementStaticCreateInputs) HasUsagePeriod() bool`

HasUsagePeriod returns a boolean if a field has been set.

### GetType

`func (o *EntitlementStaticCreateInputs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementStaticCreateInputs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementStaticCreateInputs) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *EntitlementStaticCreateInputs) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EntitlementStaticCreateInputs) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EntitlementStaticCreateInputs) SetConfig(v string)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


