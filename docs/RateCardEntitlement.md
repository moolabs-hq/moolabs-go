# RateCardEntitlement

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
**Config** | **string** | The JSON parsable config of the entitlement. This value is also returned when checking entitlement access and it is useful for configuring fine-grained access settings to the feature, implemented in your own system. Has to be an object. | 

## Methods

### NewRateCardEntitlement

`func NewRateCardEntitlement(type_ string, config string, ) *RateCardEntitlement`

NewRateCardEntitlement instantiates a new RateCardEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardEntitlementWithDefaults

`func NewRateCardEntitlementWithDefaults() *RateCardEntitlement`

NewRateCardEntitlementWithDefaults instantiates a new RateCardEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *RateCardEntitlement) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RateCardEntitlement) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RateCardEntitlement) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RateCardEntitlement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *RateCardEntitlement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCardEntitlement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCardEntitlement) SetType(v string)`

SetType sets Type field to given value.


### GetIsSoftLimit

`func (o *RateCardEntitlement) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *RateCardEntitlement) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *RateCardEntitlement) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *RateCardEntitlement) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetIssueAfterReset

`func (o *RateCardEntitlement) GetIssueAfterReset() float64`

GetIssueAfterReset returns the IssueAfterReset field if non-nil, zero value otherwise.

### GetIssueAfterResetOk

`func (o *RateCardEntitlement) GetIssueAfterResetOk() (*float64, bool)`

GetIssueAfterResetOk returns a tuple with the IssueAfterReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterReset

`func (o *RateCardEntitlement) SetIssueAfterReset(v float64)`

SetIssueAfterReset sets IssueAfterReset field to given value.

### HasIssueAfterReset

`func (o *RateCardEntitlement) HasIssueAfterReset() bool`

HasIssueAfterReset returns a boolean if a field has been set.

### GetIssueAfterResetPriority

`func (o *RateCardEntitlement) GetIssueAfterResetPriority() int32`

GetIssueAfterResetPriority returns the IssueAfterResetPriority field if non-nil, zero value otherwise.

### GetIssueAfterResetPriorityOk

`func (o *RateCardEntitlement) GetIssueAfterResetPriorityOk() (*int32, bool)`

GetIssueAfterResetPriorityOk returns a tuple with the IssueAfterResetPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueAfterResetPriority

`func (o *RateCardEntitlement) SetIssueAfterResetPriority(v int32)`

SetIssueAfterResetPriority sets IssueAfterResetPriority field to given value.

### HasIssueAfterResetPriority

`func (o *RateCardEntitlement) HasIssueAfterResetPriority() bool`

HasIssueAfterResetPriority returns a boolean if a field has been set.

### GetPreserveOverageAtReset

`func (o *RateCardEntitlement) GetPreserveOverageAtReset() bool`

GetPreserveOverageAtReset returns the PreserveOverageAtReset field if non-nil, zero value otherwise.

### GetPreserveOverageAtResetOk

`func (o *RateCardEntitlement) GetPreserveOverageAtResetOk() (*bool, bool)`

GetPreserveOverageAtResetOk returns a tuple with the PreserveOverageAtReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverageAtReset

`func (o *RateCardEntitlement) SetPreserveOverageAtReset(v bool)`

SetPreserveOverageAtReset sets PreserveOverageAtReset field to given value.

### HasPreserveOverageAtReset

`func (o *RateCardEntitlement) HasPreserveOverageAtReset() bool`

HasPreserveOverageAtReset returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *RateCardEntitlement) GetUsagePeriod() string`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *RateCardEntitlement) GetUsagePeriodOk() (*string, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *RateCardEntitlement) SetUsagePeriod(v string)`

SetUsagePeriod sets UsagePeriod field to given value.

### HasUsagePeriod

`func (o *RateCardEntitlement) HasUsagePeriod() bool`

HasUsagePeriod returns a boolean if a field has been set.

### GetConfig

`func (o *RateCardEntitlement) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RateCardEntitlement) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RateCardEntitlement) SetConfig(v string)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


