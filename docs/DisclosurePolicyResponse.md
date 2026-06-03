# DisclosurePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirement** | **string** |  | 
**DisclosureText** | Pointer to **string** |  | [optional] 
**PolicyHash** | **string** |  | 
**UpdatedByActorId** | Pointer to **string** |  | [optional] 

## Methods

### NewDisclosurePolicyResponse

`func NewDisclosurePolicyResponse(requirement string, policyHash string, ) *DisclosurePolicyResponse`

NewDisclosurePolicyResponse instantiates a new DisclosurePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosurePolicyResponseWithDefaults

`func NewDisclosurePolicyResponseWithDefaults() *DisclosurePolicyResponse`

NewDisclosurePolicyResponseWithDefaults instantiates a new DisclosurePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirement

`func (o *DisclosurePolicyResponse) GetRequirement() string`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *DisclosurePolicyResponse) GetRequirementOk() (*string, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *DisclosurePolicyResponse) SetRequirement(v string)`

SetRequirement sets Requirement field to given value.


### GetDisclosureText

`func (o *DisclosurePolicyResponse) GetDisclosureText() string`

GetDisclosureText returns the DisclosureText field if non-nil, zero value otherwise.

### GetDisclosureTextOk

`func (o *DisclosurePolicyResponse) GetDisclosureTextOk() (*string, bool)`

GetDisclosureTextOk returns a tuple with the DisclosureText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureText

`func (o *DisclosurePolicyResponse) SetDisclosureText(v string)`

SetDisclosureText sets DisclosureText field to given value.

### HasDisclosureText

`func (o *DisclosurePolicyResponse) HasDisclosureText() bool`

HasDisclosureText returns a boolean if a field has been set.

### GetPolicyHash

`func (o *DisclosurePolicyResponse) GetPolicyHash() string`

GetPolicyHash returns the PolicyHash field if non-nil, zero value otherwise.

### GetPolicyHashOk

`func (o *DisclosurePolicyResponse) GetPolicyHashOk() (*string, bool)`

GetPolicyHashOk returns a tuple with the PolicyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyHash

`func (o *DisclosurePolicyResponse) SetPolicyHash(v string)`

SetPolicyHash sets PolicyHash field to given value.


### GetUpdatedByActorId

`func (o *DisclosurePolicyResponse) GetUpdatedByActorId() string`

GetUpdatedByActorId returns the UpdatedByActorId field if non-nil, zero value otherwise.

### GetUpdatedByActorIdOk

`func (o *DisclosurePolicyResponse) GetUpdatedByActorIdOk() (*string, bool)`

GetUpdatedByActorIdOk returns a tuple with the UpdatedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByActorId

`func (o *DisclosurePolicyResponse) SetUpdatedByActorId(v string)`

SetUpdatedByActorId sets UpdatedByActorId field to given value.

### HasUpdatedByActorId

`func (o *DisclosurePolicyResponse) HasUpdatedByActorId() bool`

HasUpdatedByActorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


