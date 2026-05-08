# EntitlementValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasAccess** | **bool** | Whether the subject has access to the feature. Shared accross all entitlement types. | [readonly] 
**Balance** | Pointer to **float64** | Only available for metered entitlements. Metered entitlements are built around a balance calculation where feature usage is deducted from the issued grants. Balance represents the remaining balance of the entitlement, it&#39;s value never turns negative. | [optional] [readonly] 
**Usage** | Pointer to **float64** | Only available for metered entitlements. Returns the total feature usage in the current period. | [optional] [readonly] 
**Overage** | Pointer to **float64** | Only available for metered entitlements. Overage represents the usage that wasn&#39;t covered by grants, e.g. if the subject had a total feature usage of 100 in the period but they were only granted 80, there would be 20 overage. | [optional] [readonly] 
**TotalAvailableGrantAmount** | Pointer to **float64** | The summed value of all grant amounts that are active at the time of the query. | [optional] [readonly] 
**Config** | Pointer to **string** | Only available for static entitlements. The JSON parsable config of the entitlement. | [optional] [readonly] 

## Methods

### NewEntitlementValue

`func NewEntitlementValue(hasAccess bool, ) *EntitlementValue`

NewEntitlementValue instantiates a new EntitlementValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValueWithDefaults

`func NewEntitlementValueWithDefaults() *EntitlementValue`

NewEntitlementValueWithDefaults instantiates a new EntitlementValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasAccess

`func (o *EntitlementValue) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *EntitlementValue) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *EntitlementValue) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.


### GetBalance

`func (o *EntitlementValue) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EntitlementValue) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EntitlementValue) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *EntitlementValue) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetUsage

`func (o *EntitlementValue) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EntitlementValue) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EntitlementValue) SetUsage(v float64)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EntitlementValue) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetOverage

`func (o *EntitlementValue) GetOverage() float64`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *EntitlementValue) GetOverageOk() (*float64, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *EntitlementValue) SetOverage(v float64)`

SetOverage sets Overage field to given value.

### HasOverage

`func (o *EntitlementValue) HasOverage() bool`

HasOverage returns a boolean if a field has been set.

### GetTotalAvailableGrantAmount

`func (o *EntitlementValue) GetTotalAvailableGrantAmount() float64`

GetTotalAvailableGrantAmount returns the TotalAvailableGrantAmount field if non-nil, zero value otherwise.

### GetTotalAvailableGrantAmountOk

`func (o *EntitlementValue) GetTotalAvailableGrantAmountOk() (*float64, bool)`

GetTotalAvailableGrantAmountOk returns a tuple with the TotalAvailableGrantAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailableGrantAmount

`func (o *EntitlementValue) SetTotalAvailableGrantAmount(v float64)`

SetTotalAvailableGrantAmount sets TotalAvailableGrantAmount field to given value.

### HasTotalAvailableGrantAmount

`func (o *EntitlementValue) HasTotalAvailableGrantAmount() bool`

HasTotalAvailableGrantAmount returns a boolean if a field has been set.

### GetConfig

`func (o *EntitlementValue) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EntitlementValue) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EntitlementValue) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EntitlementValue) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


