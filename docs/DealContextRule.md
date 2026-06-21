# DealContextRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | **string** |  | 
**EscalateTo** | Pointer to **string** |  | [optional] 
**RequireApprover** | Pointer to **string** |  | [optional] 
**VerdictOverride** | Pointer to **string** |  | [optional] 
**Flag** | Pointer to **string** |  | [optional] 

## Methods

### NewDealContextRule

`func NewDealContextRule(when string, ) *DealContextRule`

NewDealContextRule instantiates a new DealContextRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealContextRuleWithDefaults

`func NewDealContextRuleWithDefaults() *DealContextRule`

NewDealContextRuleWithDefaults instantiates a new DealContextRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *DealContextRule) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *DealContextRule) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *DealContextRule) SetWhen(v string)`

SetWhen sets When field to given value.


### GetEscalateTo

`func (o *DealContextRule) GetEscalateTo() string`

GetEscalateTo returns the EscalateTo field if non-nil, zero value otherwise.

### GetEscalateToOk

`func (o *DealContextRule) GetEscalateToOk() (*string, bool)`

GetEscalateToOk returns a tuple with the EscalateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalateTo

`func (o *DealContextRule) SetEscalateTo(v string)`

SetEscalateTo sets EscalateTo field to given value.

### HasEscalateTo

`func (o *DealContextRule) HasEscalateTo() bool`

HasEscalateTo returns a boolean if a field has been set.

### GetRequireApprover

`func (o *DealContextRule) GetRequireApprover() string`

GetRequireApprover returns the RequireApprover field if non-nil, zero value otherwise.

### GetRequireApproverOk

`func (o *DealContextRule) GetRequireApproverOk() (*string, bool)`

GetRequireApproverOk returns a tuple with the RequireApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApprover

`func (o *DealContextRule) SetRequireApprover(v string)`

SetRequireApprover sets RequireApprover field to given value.

### HasRequireApprover

`func (o *DealContextRule) HasRequireApprover() bool`

HasRequireApprover returns a boolean if a field has been set.

### GetVerdictOverride

`func (o *DealContextRule) GetVerdictOverride() string`

GetVerdictOverride returns the VerdictOverride field if non-nil, zero value otherwise.

### GetVerdictOverrideOk

`func (o *DealContextRule) GetVerdictOverrideOk() (*string, bool)`

GetVerdictOverrideOk returns a tuple with the VerdictOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdictOverride

`func (o *DealContextRule) SetVerdictOverride(v string)`

SetVerdictOverride sets VerdictOverride field to given value.

### HasVerdictOverride

`func (o *DealContextRule) HasVerdictOverride() bool`

HasVerdictOverride returns a boolean if a field has been set.

### GetFlag

`func (o *DealContextRule) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *DealContextRule) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *DealContextRule) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *DealContextRule) HasFlag() bool`

HasFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


