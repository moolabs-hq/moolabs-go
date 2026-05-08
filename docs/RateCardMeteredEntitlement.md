# RateCardMeteredEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**Type** | **string** |  | 
**IsSoftLimit** | Pointer to **bool** | If softLimit&#x3D;true the subject can use the feature even if the entitlement is exhausted, hasAccess will always be true. | [optional] [default to false]
**IssueAfterReset** | Pointer to **float64** | You can grant usage automatically alongside the entitlement, the example scenario would be creating a starting balance. If an amount is specified here, a grant will be created alongside the entitlement with the specified amount. That grant will have it&#39;s rollover settings configured in a way that after each reset operation, the balance will return the original amount specified here. Manually creating such a grant would mean having the \&quot;amount\&quot;, \&quot;minRolloverAmount\&quot;, and \&quot;maxRolloverAmount\&quot; fields all be the same. | [optional] 
**IssueAfterResetPriority** | Pointer to **int32** | Defines the grant priority for the default grant. | [optional] [default to 1]
**PreserveOverageAtReset** | Pointer to **bool** | If true, the overage is preserved at reset. If false, the usage is reset to 0. | [optional] [default to false]
**UsagePeriod** | Pointer to **string** | The interval of the metered entitlement. Defaults to the billing cadence of the rate card. | [optional] 

## Methods

### NewRateCardMeteredEntitlement

`func NewRateCardMeteredEntitlement(type_ string, ) *RateCardMeteredEntitlement`

NewRateCardMeteredEntitlement instantiates a new RateCardMeteredEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardMeteredEntitlementWithDefaults

`func NewRateCardMeteredEntitlementWithDefaults() *RateCardMeteredEntitlement`

NewRateCardMeteredEntitlementWithDefaults instantiates a new RateCardMeteredEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RateCardMeteredEntitlement) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RateCardMeteredEntitlement) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RateCardMeteredEntitlement) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RateCardMeteredEntitlement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *RateCardMeteredEntitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCardMeteredEntitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCardMeteredEntitlement) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *RateCardMeteredEntitlement) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *RateCardMeteredEntitlement) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *RateCardMeteredEntitlement) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *RateCardMeteredEntitlement) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *RateCardMeteredEntitlement) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *RateCardMeteredEntitlement) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *RateCardMeteredEntitlement) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *RateCardMeteredEntitlement) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *RateCardMeteredEntitlement) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *RateCardMeteredEntitlement) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *RateCardMeteredEntitlement) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *RateCardMeteredEntitlement) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *RateCardMeteredEntitlement) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *RateCardMeteredEntitlement) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *RateCardMeteredEntitlement) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *RateCardMeteredEntitlement) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *RateCardMeteredEntitlement) GetUsagePeriod() string`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *RateCardMeteredEntitlement) GetUsagePeriodOk() (*string, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *RateCardMeteredEntitlement) SetUsagePeriod(v string)`

SetUsagePeriod sets UsagePeriod field to given value.

### HasUsagePeriod

`func (o *RateCardMeteredEntitlement) HasUsagePeriod() bool`

HasUsagePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


