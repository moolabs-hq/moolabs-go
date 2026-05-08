# Entitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IsSoftLimit** | Pointer to **bool** | If softLimit&#x3D;true the subject can use the feature even if the entitlement is exhausted, hasAccess will always be true. | [optional] [default to false]
**IsUnlimited** | Pointer to **bool** | Deprecated, ignored by the backend. Please use isSoftLimit instead; this field will be removed in the future. | [optional] [default to false]
**IssueAfterReset** | Pointer to **float64** | You can grant usage automatically alongside the entitlement, the example scenario would be creating a starting balance. If an amount is specified here, a grant will be created alongside the entitlement with the specified amount. That grant will have it&#39;s rollover settings configured in a way that after each reset operation, the balance will return the original amount specified here. Manually creating such a grant would mean having the \&quot;amount\&quot;, \&quot;minRolloverAmount\&quot;, and \&quot;maxRolloverAmount\&quot; fields all be the same. | [optional] 
**IssueAfterResetPriority** | Pointer to **int32** | Defines the grant priority for the default grant. | [optional] [default to 1]
**PreserveOverageAtReset** | Pointer to **bool** | If true, the overage is preserved at reset. If false, the usage is reset to 0. | [optional] [default to false]
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**ActiveFrom** | **time.Time** | The cadence start of the resource. | 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | The annotations of the entitlement. | [optional] [readonly] 
**Id** | **string** | Readonly unique ULID identifier. | [readonly] 
**SubjectKey** | **string** | The identifier key unique to the subject. NOTE: Subjects are being deprecated, please use the new customer APIs. | 
**FeatureKey** | **string** | The feature the subject is entitled to use. | 
**FeatureId** | **string** | The feature the subject is entitled to use. | 
**LastReset** | **time.Time** | The time the last reset happened. | [readonly] 
**CurrentUsagePeriod** | [**Period**](Period.md) | The current usage period. | 
**MeasureUsageFrom** | **time.Time** | The time from which usage is measured. If not specified on creation, defaults to entitlement creation time. | [readonly] 
**UsagePeriod** | [**RecurringPeriod**](RecurringPeriod.md) | The defined usage period of the entitlement | 
**Config** | **string** | The JSON parsable config of the entitlement. This value is also returned when checking entitlement access and it is useful for configuring fine-grained access settings to the feature, implemented in your own system. Has to be an object. | 

## Methods

### NewEntitlement

`func NewEntitlement(type_ string, createdAt time.Time, updatedAt time.Time, activeFrom time.Time, id string, subjectKey string, featureKey string, featureId string, lastReset time.Time, currentUsagePeriod Period, measureUsageFrom time.Time, usagePeriod RecurringPeriod, config string, ) *Entitlement`

NewEntitlement instantiates a new Entitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementWithDefaults

`func NewEntitlementWithDefaults() *Entitlement`

NewEntitlementWithDefaults instantiates a new Entitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Entitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entitlement) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *Entitlement) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *Entitlement) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *Entitlement) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *Entitlement) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetIsUnlimited

`func (o *Entitlement) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *Entitlement) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *Entitlement) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.

### HasIsUnlimited

`func (o *Entitlement) HasIsUnlimited() bool`

HasIsUnlimited returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *Entitlement) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *Entitlement) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *Entitlement) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *Entitlement) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *Entitlement) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *Entitlement) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *Entitlement) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *Entitlement) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *Entitlement) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *Entitlement) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *Entitlement) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *Entitlement) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Entitlement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Entitlement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Entitlement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Entitlement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Entitlement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Entitlement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Entitlement) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Entitlement) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Entitlement) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Entitlement) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *Entitlement) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Entitlement) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Entitlement) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Entitlement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActiveFrom

`func (o *Entitlement) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *Entitlement) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *Entitlement) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *Entitlement) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *Entitlement) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *Entitlement) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *Entitlement) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetAnnotations

`func (o *Entitlement) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Entitlement) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Entitlement) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Entitlement) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetId

`func (o *Entitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlement) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectKey

`func (o *Entitlement) GetSubjectKey() string`

GetSubjectKey returns the SubjectKey field if non-nil, zero value otherwise.

### GetSubjectKeyOk

`func (o *Entitlement) GetSubjectKeyOk() (*string, bool)`

GetSubjectKeyOk returns a tuple with the SubjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKey

`func (o *Entitlement) SetSubjectKey(v string)`

SetSubjectKey sets SubjectKey field to given value.


### GetFeatureKey

`func (o *Entitlement) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *Entitlement) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *Entitlement) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetFeatureId

`func (o *Entitlement) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *Entitlement) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *Entitlement) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetLastReset

`func (o *Entitlement) GetLastReset() time.Time`

GetLastReset returns the LastReset field if non-nil, zero value otherwise.

### GetLastResetOk

`func (o *Entitlement) GetLastResetOk() (*time.Time, bool)`

GetLastResetOk returns a tuple with the LastReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReset

`func (o *Entitlement) SetLastReset(v time.Time)`

SetLastReset sets LastReset field to given value.


### GetCurrentUsagePeriod

`func (o *Entitlement) GetCurrentUsagePeriod() Period`

GetCurrentUsagePeriod returns the CurrentUsagePeriod field if non-nil, zero value otherwise.

### GetCurrentUsagePeriodOk

`func (o *Entitlement) GetCurrentUsagePeriodOk() (*Period, bool)`

GetCurrentUsagePeriodOk returns a tuple with the CurrentUsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsagePeriod

`func (o *Entitlement) SetCurrentUsagePeriod(v Period)`

SetCurrentUsagePeriod sets CurrentUsagePeriod field to given value.


### GetMeasureUsageFrom

`func (o *Entitlement) GetMeasureUsageFrom() time.Time`

GetMeasureUsageFrom returns the MeasureUsageFrom field if non-nil, zero value otherwise.

### GetMeasureUsageFromOk

`func (o *Entitlement) GetMeasureUsageFromOk() (*time.Time, bool)`

GetMeasureUsageFromOk returns a tuple with the MeasureUsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsageFrom

`func (o *Entitlement) SetMeasureUsageFrom(v time.Time)`

SetMeasureUsageFrom sets MeasureUsageFrom field to given value.


### GetUsagePeriod

`func (o *Entitlement) GetUsagePeriod() RecurringPeriod`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *Entitlement) GetUsagePeriodOk() (*RecurringPeriod, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *Entitlement) SetUsagePeriod(v RecurringPeriod)`

SetUsagePeriod sets UsagePeriod field to given value.


### GetConfig

`func (o *Entitlement) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Entitlement) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Entitlement) SetConfig(v string)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


