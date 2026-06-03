# DisclosurePolicyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirement** | Pointer to **string** |  | [optional] [default to "not_required"]
**DisclosureText** | Pointer to **string** |  | [optional] 
**ChangeReason** | **string** |  | 

## Methods

### NewDisclosurePolicyUpdateRequest

`func NewDisclosurePolicyUpdateRequest(changeReason string, ) *DisclosurePolicyUpdateRequest`

NewDisclosurePolicyUpdateRequest instantiates a new DisclosurePolicyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosurePolicyUpdateRequestWithDefaults

`func NewDisclosurePolicyUpdateRequestWithDefaults() *DisclosurePolicyUpdateRequest`

NewDisclosurePolicyUpdateRequestWithDefaults instantiates a new DisclosurePolicyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirement

`func (o *DisclosurePolicyUpdateRequest) GetRequirement() string`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *DisclosurePolicyUpdateRequest) GetRequirementOk() (*string, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *DisclosurePolicyUpdateRequest) SetRequirement(v string)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *DisclosurePolicyUpdateRequest) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.

### GetDisclosureText

`func (o *DisclosurePolicyUpdateRequest) GetDisclosureText() string`

GetDisclosureText returns the DisclosureText field if non-nil, zero value otherwise.

### GetDisclosureTextOk

`func (o *DisclosurePolicyUpdateRequest) GetDisclosureTextOk() (*string, bool)`

GetDisclosureTextOk returns a tuple with the DisclosureText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureText

`func (o *DisclosurePolicyUpdateRequest) SetDisclosureText(v string)`

SetDisclosureText sets DisclosureText field to given value.

### HasDisclosureText

`func (o *DisclosurePolicyUpdateRequest) HasDisclosureText() bool`

HasDisclosureText returns a boolean if a field has been set.

### GetChangeReason

`func (o *DisclosurePolicyUpdateRequest) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *DisclosurePolicyUpdateRequest) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *DisclosurePolicyUpdateRequest) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


