# EntitlementCreateInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**FeatureId** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**Type** | **string** |  | 
**IsSoftLimit** | Pointer to **bool** | If softLimit&#x3D;true the subject can use the feature even if the entitlement is exhausted, hasAccess will always be true. | [optional] [default to false]
**IsUnlimited** | Pointer to **bool** | Deprecated, ignored by the backend. Please use isSoftLimit instead; this field will be removed in the future. | [optional] [default to false]
**UsagePeriod** | [**RecurringPeriodCreateInput**](RecurringPeriodCreateInput.md) | The usage period associated with the entitlement. | 
**MeasureUsageFrom** | Pointer to [**MeasureUsageFrom**](MeasureUsageFrom.md) | Defines the time from which usage is measured. If not specified on creation, defaults to entitlement creation time. | [optional] 
**IssueAfterReset** | Pointer to **float64** | You can grant usage automatically alongside the entitlement, the example scenario would be creating a starting balance. If an amount is specified here, a grant will be created alongside the entitlement with the specified amount. That grant will have it&#39;s rollover settings configured in a way that after each reset operation, the balance will return the original amount specified here. Manually creating such a grant would mean having the \&quot;amount\&quot;, \&quot;minRolloverAmount\&quot;, and \&quot;maxRolloverAmount\&quot; fields all be the same. | [optional] 
**IssueAfterResetPriority** | Pointer to **int32** | Defines the grant priority for the default grant. | [optional] [default to 1]
**PreserveOverageAtReset** | Pointer to **bool** | If true, the overage is preserved at reset. If false, the usage is reset to 0. | [optional] [default to false]
**Config** | **string** | The JSON parsable config of the entitlement. This value is also returned when checking entitlement access and it is useful for configuring fine-grained access settings to the feature, implemented in your own system. Has to be an object. | 

## Methods

### NewEntitlementCreateInputs

`func NewEntitlementCreateInputs(type_ string, usagePeriod RecurringPeriodCreateInput, config string, ) *EntitlementCreateInputs`

NewEntitlementCreateInputs instantiates a new EntitlementCreateInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementCreateInputsWithDefaults

`func NewEntitlementCreateInputsWithDefaults() *EntitlementCreateInputs`

NewEntitlementCreateInputsWithDefaults instantiates a new EntitlementCreateInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *EntitlementCreateInputs) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementCreateInputs) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementCreateInputs) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *EntitlementCreateInputs) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetFeatureId

`func (o *EntitlementCreateInputs) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementCreateInputs) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementCreateInputs) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *EntitlementCreateInputs) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementCreateInputs) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementCreateInputs) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementCreateInputs) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementCreateInputs) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *EntitlementCreateInputs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementCreateInputs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementCreateInputs) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *EntitlementCreateInputs) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *EntitlementCreateInputs) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *EntitlementCreateInputs) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *EntitlementCreateInputs) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetIsUnlimited

`func (o *EntitlementCreateInputs) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *EntitlementCreateInputs) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *EntitlementCreateInputs) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.

### HasIsUnlimited

`func (o *EntitlementCreateInputs) HasIsUnlimited() bool`

HasIsUnlimited returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *EntitlementCreateInputs) GetUsagePeriod() RecurringPeriodCreateInput`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementCreateInputs) GetUsagePeriodOk() (*RecurringPeriodCreateInput, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementCreateInputs) SetUsagePeriod(v RecurringPeriodCreateInput)`

SetUsagePeriod sets UsagePeriod field to given value.


### GetMeasureUsageFrom

`func (o *EntitlementCreateInputs) GetMeasureUsageFrom() MeasureUsageFrom`

GetMeasureUsageFrom returns the MeasureUsageFrom field if non-nil, zero value otherwise.

### GetMeasureUsageFromOk

`func (o *EntitlementCreateInputs) GetMeasureUsageFromOk() (*MeasureUsageFrom, bool)`

GetMeasureUsageFromOk returns a tuple with the MeasureUsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsageFrom

`func (o *EntitlementCreateInputs) SetMeasureUsageFrom(v MeasureUsageFrom)`

SetMeasureUsageFrom sets MeasureUsageFrom field to given value.

### HasMeasureUsageFrom

`func (o *EntitlementCreateInputs) HasMeasureUsageFrom() bool`

HasMeasureUsageFrom returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *EntitlementCreateInputs) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *EntitlementCreateInputs) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *EntitlementCreateInputs) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *EntitlementCreateInputs) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *EntitlementCreateInputs) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *EntitlementCreateInputs) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *EntitlementCreateInputs) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *EntitlementCreateInputs) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *EntitlementCreateInputs) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *EntitlementCreateInputs) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *EntitlementCreateInputs) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *EntitlementCreateInputs) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.

### GetConfig

`func (o *EntitlementCreateInputs) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EntitlementCreateInputs) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EntitlementCreateInputs) SetConfig(v string)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


