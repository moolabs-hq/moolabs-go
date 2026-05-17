# AccountTeamMemberPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**AccountTeamRole**](AccountTeamRole.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**NotifyOnEscalation** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountTeamMemberPatch

`func NewAccountTeamMemberPatch() *AccountTeamMemberPatch`

NewAccountTeamMemberPatch instantiates a new AccountTeamMemberPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTeamMemberPatchWithDefaults

`func NewAccountTeamMemberPatchWithDefaults() *AccountTeamMemberPatch`

NewAccountTeamMemberPatchWithDefaults instantiates a new AccountTeamMemberPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AccountTeamMemberPatch) GetRole() AccountTeamRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccountTeamMemberPatch) GetRoleOk() (*AccountTeamRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccountTeamMemberPatch) SetRole(v AccountTeamRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AccountTeamMemberPatch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetName

`func (o *AccountTeamMemberPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTeamMemberPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTeamMemberPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountTeamMemberPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *AccountTeamMemberPatch) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountTeamMemberPatch) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountTeamMemberPatch) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountTeamMemberPatch) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsPrimary

`func (o *AccountTeamMemberPatch) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *AccountTeamMemberPatch) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *AccountTeamMemberPatch) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *AccountTeamMemberPatch) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetNotifyOnEscalation

`func (o *AccountTeamMemberPatch) GetNotifyOnEscalation() bool`

GetNotifyOnEscalation returns the NotifyOnEscalation field if non-nil, zero value otherwise.

### GetNotifyOnEscalationOk

`func (o *AccountTeamMemberPatch) GetNotifyOnEscalationOk() (*bool, bool)`

GetNotifyOnEscalationOk returns a tuple with the NotifyOnEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnEscalation

`func (o *AccountTeamMemberPatch) SetNotifyOnEscalation(v bool)`

SetNotifyOnEscalation sets NotifyOnEscalation field to given value.

### HasNotifyOnEscalation

`func (o *AccountTeamMemberPatch) HasNotifyOnEscalation() bool`

HasNotifyOnEscalation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


