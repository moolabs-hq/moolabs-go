# ResetEntitlementUsageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveAt** | Pointer to **time.Time** | The time at which the reset takes effect, defaults to now. The reset cannot be in the future. The provided value is truncated to the minute due to how historical meter data is stored. | [optional] 
**RetainAnchor** | Pointer to **bool** | Determines whether the usage period anchor is retained or reset to the effectiveAt time. - If true, the usage period anchor is retained. - If false, the usage period anchor is reset to the effectiveAt time. | [optional] 
**PreserveOverage** | Pointer to **bool** | Determines whether the overage is preserved or forgiven, overriding the entitlement&#39;s default behavior. - If true, the overage is preserved. - If false, the overage is forgiven. | [optional] 

## Methods

### NewResetEntitlementUsageInput

`func NewResetEntitlementUsageInput() *ResetEntitlementUsageInput`

NewResetEntitlementUsageInput instantiates a new ResetEntitlementUsageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetEntitlementUsageInputWithDefaults

`func NewResetEntitlementUsageInputWithDefaults() *ResetEntitlementUsageInput`

NewResetEntitlementUsageInputWithDefaults instantiates a new ResetEntitlementUsageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveAt

`func (o *ResetEntitlementUsageInput) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ResetEntitlementUsageInput) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ResetEntitlementUsageInput) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *ResetEntitlementUsageInput) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetRetainAnchor

`func (o *ResetEntitlementUsageInput) GetRetainAnchor() bool`

GetRetainAnchor returns the RetainAnchor field if non-nil, zero value otherwise.

### GetRetainAnchorOk

`func (o *ResetEntitlementUsageInput) GetRetainAnchorOk() (*bool, bool)`

GetRetainAnchorOk returns a tuple with the RetainAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAnchor

`func (o *ResetEntitlementUsageInput) SetRetainAnchor(v bool)`

SetRetainAnchor sets RetainAnchor field to given value.

### HasRetainAnchor

`func (o *ResetEntitlementUsageInput) HasRetainAnchor() bool`

HasRetainAnchor returns a boolean if a field has been set.

### GetPreserveOverage

`func (o *ResetEntitlementUsageInput) GetPreserveOverage() bool`

GetPreserveOverage returns the PreserveOverage field if non-nil, zero value otherwise.

### GetPreserveOverageOk

`func (o *ResetEntitlementUsageInput) GetPreserveOverageOk() (*bool, bool)`

GetPreserveOverageOk returns a tuple with the PreserveOverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveOverage

`func (o *ResetEntitlementUsageInput) SetPreserveOverage(v bool)`

SetPreserveOverage sets PreserveOverage field to given value.

### HasPreserveOverage

`func (o *ResetEntitlementUsageInput) HasPreserveOverage() bool`

HasPreserveOverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


