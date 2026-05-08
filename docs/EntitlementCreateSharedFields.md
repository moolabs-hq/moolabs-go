# EntitlementCreateSharedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**FeatureId** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**UsagePeriod** | Pointer to [**RecurringPeriodCreateInput**](RecurringPeriodCreateInput.md) | The usage period associated with the entitlement. | [optional] 

## Methods

### NewEntitlementCreateSharedFields

`func NewEntitlementCreateSharedFields() *EntitlementCreateSharedFields`

NewEntitlementCreateSharedFields instantiates a new EntitlementCreateSharedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementCreateSharedFieldsWithDefaults

`func NewEntitlementCreateSharedFieldsWithDefaults() *EntitlementCreateSharedFields`

NewEntitlementCreateSharedFieldsWithDefaults instantiates a new EntitlementCreateSharedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *EntitlementCreateSharedFields) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementCreateSharedFields) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementCreateSharedFields) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *EntitlementCreateSharedFields) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetFeatureId

`func (o *EntitlementCreateSharedFields) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementCreateSharedFields) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementCreateSharedFields) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *EntitlementCreateSharedFields) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementCreateSharedFields) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementCreateSharedFields) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementCreateSharedFields) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementCreateSharedFields) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *EntitlementCreateSharedFields) GetUsagePeriod() RecurringPeriodCreateInput`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementCreateSharedFields) GetUsagePeriodOk() (*RecurringPeriodCreateInput, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementCreateSharedFields) SetUsagePeriod(v RecurringPeriodCreateInput)`

SetUsagePeriod sets UsagePeriod field to given value.

### HasUsagePeriod

`func (o *EntitlementCreateSharedFields) HasUsagePeriod() bool`

HasUsagePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


