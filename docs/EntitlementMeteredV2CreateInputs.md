# EntitlementMeteredV2CreateInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**FeatureId** | Pointer to **string** | The feature the subject is entitled to use. Either featureKey or featureId is required. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**Type** | **string** |  | 
**IsSoftLimit** | Pointer to **bool** | If softLimit&#x3D;true the subject can use the feature even if the entitlement is exhausted, hasAccess will always be true. | [optional] [default to false]
**UsagePeriod** | [**RecurringPeriodCreateInput**](RecurringPeriodCreateInput.md) | The usage period associated with the entitlement. | 
**MeasureUsageFrom** | Pointer to [**MeasureUsageFrom**](MeasureUsageFrom.md) | Defines the time from which usage is measured. If not specified on creation, defaults to entitlement creation time. | [optional] 
**PreserveOverageAtReset** | Pointer to **bool** | If true, the overage is preserved at reset. If false, the usage is reset to 0. | [optional] [default to false]
**IssueAfterReset** | Pointer to **float64** | You can grant usage automatically alongside the entitlement, the example scenario would be creating a starting balance. If an amount is specified here, a grant will be created alongside the entitlement with the specified amount. That grant will have it&#39;s rollover settings configured in a way that after each reset operation, the balance will return the original amount specified here. Manually creating such a grant would mean having the \&quot;amount\&quot;, \&quot;minRolloverAmount\&quot;, and \&quot;maxRolloverAmount\&quot; fields all be the same. | [optional] 
**IssueAfterResetPriority** | Pointer to **int32** | Defines the grant priority for the default grant. | [optional] [default to 1]
**Issue** | Pointer to [**IssueAfterReset**](IssueAfterReset.md) | Issue after reset | [optional] 
**Grants** | Pointer to [**[]EntitlementGrantCreateInputV2**](EntitlementGrantCreateInputV2.md) | Grants | [optional] 

## Methods

### NewEntitlementMeteredV2CreateInputs

`func NewEntitlementMeteredV2CreateInputs(type_ string, usagePeriod RecurringPeriodCreateInput, ) *EntitlementMeteredV2CreateInputs`

NewEntitlementMeteredV2CreateInputs instantiates a new EntitlementMeteredV2CreateInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementMeteredV2CreateInputsWithDefaults

`func NewEntitlementMeteredV2CreateInputsWithDefaults() *EntitlementMeteredV2CreateInputs`

NewEntitlementMeteredV2CreateInputsWithDefaults instantiates a new EntitlementMeteredV2CreateInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *EntitlementMeteredV2CreateInputs) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementMeteredV2CreateInputs) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementMeteredV2CreateInputs) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *EntitlementMeteredV2CreateInputs) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetFeatureId

`func (o *EntitlementMeteredV2CreateInputs) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementMeteredV2CreateInputs) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementMeteredV2CreateInputs) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *EntitlementMeteredV2CreateInputs) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementMeteredV2CreateInputs) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementMeteredV2CreateInputs) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementMeteredV2CreateInputs) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementMeteredV2CreateInputs) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *EntitlementMeteredV2CreateInputs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementMeteredV2CreateInputs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementMeteredV2CreateInputs) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *EntitlementMeteredV2CreateInputs) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *EntitlementMeteredV2CreateInputs) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *EntitlementMeteredV2CreateInputs) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *EntitlementMeteredV2CreateInputs) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *EntitlementMeteredV2CreateInputs) GetUsagePeriod() RecurringPeriodCreateInput`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementMeteredV2CreateInputs) GetUsagePeriodOk() (*RecurringPeriodCreateInput, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementMeteredV2CreateInputs) SetUsagePeriod(v RecurringPeriodCreateInput)`

SetUsagePeriod sets UsagePeriod field to given value.


### GetMeasureUsageFrom

`func (o *EntitlementMeteredV2CreateInputs) GetMeasureUsageFrom() MeasureUsageFrom`

GetMeasureUsageFrom returns the MeasureUsageFrom field if non-nil, zero value otherwise.

### GetMeasureUsageFromOk

`func (o *EntitlementMeteredV2CreateInputs) GetMeasureUsageFromOk() (*MeasureUsageFrom, bool)`

GetMeasureUsageFromOk returns a tuple with the MeasureUsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsageFrom

`func (o *EntitlementMeteredV2CreateInputs) SetMeasureUsageFrom(v MeasureUsageFrom)`

SetMeasureUsageFrom sets MeasureUsageFrom field to given value.

### HasMeasureUsageFrom

`func (o *EntitlementMeteredV2CreateInputs) HasMeasureUsageFrom() bool`

HasMeasureUsageFrom returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *EntitlementMeteredV2CreateInputs) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *EntitlementMeteredV2CreateInputs) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *EntitlementMeteredV2CreateInputs) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *EntitlementMeteredV2CreateInputs) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *EntitlementMeteredV2CreateInputs) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *EntitlementMeteredV2CreateInputs) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *EntitlementMeteredV2CreateInputs) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *EntitlementMeteredV2CreateInputs) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *EntitlementMeteredV2CreateInputs) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *EntitlementMeteredV2CreateInputs) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *EntitlementMeteredV2CreateInputs) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *EntitlementMeteredV2CreateInputs) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetIssue

`func (o *EntitlementMeteredV2CreateInputs) GetIssue() IssueAfterReset`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *EntitlementMeteredV2CreateInputs) GetIssueOk() (*IssueAfterReset, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *EntitlementMeteredV2CreateInputs) SetIssue(v IssueAfterReset)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *EntitlementMeteredV2CreateInputs) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetGrants

`func (o *EntitlementMeteredV2CreateInputs) GetGrants() []EntitlementGrantCreateInputV2`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *EntitlementMeteredV2CreateInputs) GetGrantsOk() (*[]EntitlementGrantCreateInputV2, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *EntitlementMeteredV2CreateInputs) SetGrants(v []EntitlementGrantCreateInputV2)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *EntitlementMeteredV2CreateInputs) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


