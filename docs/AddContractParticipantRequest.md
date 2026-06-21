# AddContractParticipantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Side** | **string** |  | 
**Role** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewAddContractParticipantRequest

`func NewAddContractParticipantRequest(side string, role string, ) *AddContractParticipantRequest`

NewAddContractParticipantRequest instantiates a new AddContractParticipantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddContractParticipantRequestWithDefaults

`func NewAddContractParticipantRequestWithDefaults() *AddContractParticipantRequest`

NewAddContractParticipantRequestWithDefaults instantiates a new AddContractParticipantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSide

`func (o *AddContractParticipantRequest) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *AddContractParticipantRequest) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *AddContractParticipantRequest) SetSide(v string)`

SetSide sets Side field to given value.


### GetRole

`func (o *AddContractParticipantRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddContractParticipantRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddContractParticipantRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *AddContractParticipantRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddContractParticipantRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddContractParticipantRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddContractParticipantRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *AddContractParticipantRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddContractParticipantRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddContractParticipantRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddContractParticipantRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


