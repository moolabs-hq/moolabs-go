# EntitlementMeteredV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**IsSoftLimit** | Pointer to **bool** | If softLimit&#x3D;true the subject can use the feature even if the entitlement is exhausted, hasAccess will always be true. | [optional] [default to false]
**PreserveOverageAtReset** | Pointer to **bool** | If true, the overage is preserved at reset. If false, the usage is reset to 0. | [optional] [default to false]
**IssueAfterReset** | Pointer to **float64** | You can grant usage automatically alongside the entitlement, the example scenario would be creating a starting balance. If an amount is specified here, a grant will be created alongside the entitlement with the specified amount. That grant will have it&#39;s rollover settings configured in a way that after each reset operation, the balance will return the original amount specified here. Manually creating such a grant would mean having the \&quot;amount\&quot;, \&quot;minRolloverAmount\&quot;, and \&quot;maxRolloverAmount\&quot; fields all be the same. | [optional] 
**IssueAfterResetPriority** | Pointer to **int32** | Defines the grant priority for the default grant. | [optional] [default to 1]
**Issue** | Pointer to [**IssueAfterReset**](IssueAfterReset.md) | Issue after reset | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**ActiveFrom** | **time.Time** | The cadence start of the resource. | 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | The annotations of the entitlement. | [optional] [readonly] 
**Id** | **string** | Readonly unique ULID identifier. | [readonly] 
**FeatureKey** | **string** | The feature the subject is entitled to use. | 
**FeatureId** | **string** | The feature the subject is entitled to use. | 
**LastReset** | **time.Time** | The time the last reset happened. | [readonly] 
**CurrentUsagePeriod** | [**Period**](Period.md) | The current usage period. | [readonly] 
**MeasureUsageFrom** | **time.Time** | The time from which usage is measured. If not specified on creation, defaults to entitlement creation time. | [readonly] 
**UsagePeriod** | [**RecurringPeriod**](RecurringPeriod.md) | THe usage period of the entitlement. | [readonly] 
**CustomerKey** | Pointer to **string** | The identifier key unique to the customer | [optional] 
**CustomerId** | **string** | The identifier unique to the customer | 

## Methods

### NewEntitlementMeteredV2

`func NewEntitlementMeteredV2(type_ string, createdAt time.Time, updatedAt time.Time, activeFrom time.Time, id string, featureKey string, featureId string, lastReset time.Time, currentUsagePeriod Period, measureUsageFrom time.Time, usagePeriod RecurringPeriod, customerId string, ) *EntitlementMeteredV2`

NewEntitlementMeteredV2 instantiates a new EntitlementMeteredV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementMeteredV2WithDefaults

`func NewEntitlementMeteredV2WithDefaults() *EntitlementMeteredV2`

NewEntitlementMeteredV2WithDefaults instantiates a new EntitlementMeteredV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntitlementMeteredV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementMeteredV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementMeteredV2) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *EntitlementMeteredV2) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *EntitlementMeteredV2) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *EntitlementMeteredV2) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *EntitlementMeteredV2) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *EntitlementMeteredV2) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *EntitlementMeteredV2) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *EntitlementMeteredV2) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *EntitlementMeteredV2) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *EntitlementMeteredV2) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *EntitlementMeteredV2) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *EntitlementMeteredV2) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *EntitlementMeteredV2) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *EntitlementMeteredV2) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *EntitlementMeteredV2) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *EntitlementMeteredV2) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *EntitlementMeteredV2) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetIssue

`func (o *EntitlementMeteredV2) GetIssue() IssueAfterReset`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *EntitlementMeteredV2) GetIssueOk() (*IssueAfterReset, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *EntitlementMeteredV2) SetIssue(v IssueAfterReset)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *EntitlementMeteredV2) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntitlementMeteredV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntitlementMeteredV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntitlementMeteredV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EntitlementMeteredV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntitlementMeteredV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntitlementMeteredV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EntitlementMeteredV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntitlementMeteredV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntitlementMeteredV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntitlementMeteredV2) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementMeteredV2) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementMeteredV2) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementMeteredV2) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementMeteredV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActiveFrom

`func (o *EntitlementMeteredV2) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *EntitlementMeteredV2) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *EntitlementMeteredV2) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *EntitlementMeteredV2) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *EntitlementMeteredV2) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *EntitlementMeteredV2) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *EntitlementMeteredV2) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetAnnotations

`func (o *EntitlementMeteredV2) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EntitlementMeteredV2) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EntitlementMeteredV2) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EntitlementMeteredV2) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetId

`func (o *EntitlementMeteredV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementMeteredV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementMeteredV2) SetId(v string)`

SetId sets Id field to given value.


### GetFeatureKey

`func (o *EntitlementMeteredV2) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementMeteredV2) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementMeteredV2) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetFeatureId

`func (o *EntitlementMeteredV2) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementMeteredV2) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementMeteredV2) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetLastReset

`func (o *EntitlementMeteredV2) GetLastReset() time.Time`

GetLastReset returns the LastReset field if non-nil, zero value otherwise.

### GetLastResetOk

`func (o *EntitlementMeteredV2) GetLastResetOk() (*time.Time, bool)`

GetLastResetOk returns a tuple with the LastReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReset

`func (o *EntitlementMeteredV2) SetLastReset(v time.Time)`

SetLastReset sets LastReset field to given value.


### GetCurrentUsagePeriod

`func (o *EntitlementMeteredV2) GetCurrentUsagePeriod() Period`

GetCurrentUsagePeriod returns the CurrentUsagePeriod field if non-nil, zero value otherwise.

### GetCurrentUsagePeriodOk

`func (o *EntitlementMeteredV2) GetCurrentUsagePeriodOk() (*Period, bool)`

GetCurrentUsagePeriodOk returns a tuple with the CurrentUsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsagePeriod

`func (o *EntitlementMeteredV2) SetCurrentUsagePeriod(v Period)`

SetCurrentUsagePeriod sets CurrentUsagePeriod field to given value.


### GetMeasureUsageFrom

`func (o *EntitlementMeteredV2) GetMeasureUsageFrom() time.Time`

GetMeasureUsageFrom returns the MeasureUsageFrom field if non-nil, zero value otherwise.

### GetMeasureUsageFromOk

`func (o *EntitlementMeteredV2) GetMeasureUsageFromOk() (*time.Time, bool)`

GetMeasureUsageFromOk returns a tuple with the MeasureUsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasureUsageFrom

`func (o *EntitlementMeteredV2) SetMeasureUsageFrom(v time.Time)`

SetMeasureUsageFrom sets MeasureUsageFrom field to given value.


### GetUsagePeriod

`func (o *EntitlementMeteredV2) GetUsagePeriod() RecurringPeriod`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementMeteredV2) GetUsagePeriodOk() (*RecurringPeriod, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementMeteredV2) SetUsagePeriod(v RecurringPeriod)`

SetUsagePeriod sets UsagePeriod field to given value.


### GetCustomerKey

`func (o *EntitlementMeteredV2) GetCustomerKey() string`

GetCustomerKey returns the CustomerKey field if non-nil, zero value otherwise.

### GetCustomerKeyOk

`func (o *EntitlementMeteredV2) GetCustomerKeyOk() (*string, bool)`

GetCustomerKeyOk returns a tuple with the CustomerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerKey

`func (o *EntitlementMeteredV2) SetCustomerKey(v string)`

SetCustomerKey sets CustomerKey field to given value.

### HasCustomerKey

`func (o *EntitlementMeteredV2) HasCustomerKey() bool`

HasCustomerKey returns a boolean if a field has been set.

### GetCustomerId

`func (o *EntitlementMeteredV2) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EntitlementMeteredV2) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EntitlementMeteredV2) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


