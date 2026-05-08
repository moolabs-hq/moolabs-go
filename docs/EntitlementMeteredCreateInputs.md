# EntitlementMeteredCreateInputs

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

## Methods

### NewEntitlementMeteredCreateInputs

`func NewEntitlementMeteredCreateInputs(type_ string, usagePeriod RecurringPeriodCreateInput, ) *EntitlementMeteredCreateInputs`

NewEntitlementMeteredCreateInputs instantiates a new EntitlementMeteredCreateInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementMeteredCreateInputsWithDefaults

`func NewEntitlementMeteredCreateInputsWithDefaults() *EntitlementMeteredCreateInputs`

NewEntitlementMeteredCreateInputsWithDefaults instantiates a new EntitlementMeteredCreateInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *EntitlementMeteredCreateInputs) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementMeteredCreateInputs) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementMeteredCreateInputs) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *EntitlementMeteredCreateInputs) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetFeatureId

`func (o *EntitlementMeteredCreateInputs) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementMeteredCreateInputs) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementMeteredCreateInputs) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *EntitlementMeteredCreateInputs) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementMeteredCreateInputs) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementMeteredCreateInputs) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementMeteredCreateInputs) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementMeteredCreateInputs) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *EntitlementMeteredCreateInputs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementMeteredCreateInputs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementMeteredCreateInputs) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *EntitlementMeteredCreateInputs) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *EntitlementMeteredCreateInputs) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *EntitlementMeteredCreateInputs) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *EntitlementMeteredCreateInputs) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetIsUnlimited

`func (o *EntitlementMeteredCreateInputs) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *EntitlementMeteredCreateInputs) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *EntitlementMeteredCreateInputs) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.

### HasIsUnlimited

`func (o *EntitlementMeteredCreateInputs) HasIsUnlimited() bool`

HasIsUnlimited returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *EntitlementMeteredCreateInputs) GetUsagePeriod() RecurringPeriodCreateInput`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementMeteredCreateInputs) GetUsagePeriodOk() (*RecurringPeriodCreateInput, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementMeteredCreateInputs) SetUsagePeriod(v RecurringPeriodCreateInput)`

SetUsagePeriod sets UsagePeriod field to given value.


### GetMeasureUsageFrom

`func (o *EntitlementMeteredCreateInputs) GetMeasureUsageFrom() MeasureUsageFrom`

GetMeasureUsageFrom returns the MeasureUsageFrom field if non-nil, zero value otherwise.

### GetMeasureUsageFromOk

`func (o *EntitlementMeteredCreateInputs) GetMeasureUsageFromOk() (*MeasureUsageFrom, bool)`

GetMeasureUsageFromOk returns a tuple with the MeasureUsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsageFrom

`func (o *EntitlementMeteredCreateInputs) SetMeasureUsageFrom(v MeasureUsageFrom)`

SetMeasureUsageFrom sets MeasureUsageFrom field to given value.

### HasMeasureUsageFrom

`func (o *EntitlementMeteredCreateInputs) HasMeasureUsageFrom() bool`

HasMeasureUsageFrom returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *EntitlementMeteredCreateInputs) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *EntitlementMeteredCreateInputs) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *EntitlementMeteredCreateInputs) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *EntitlementMeteredCreateInputs) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *EntitlementMeteredCreateInputs) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *EntitlementMeteredCreateInputs) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *EntitlementMeteredCreateInputs) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *EntitlementMeteredCreateInputs) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *EntitlementMeteredCreateInputs) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *EntitlementMeteredCreateInputs) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *EntitlementMeteredCreateInputs) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *EntitlementMeteredCreateInputs) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


