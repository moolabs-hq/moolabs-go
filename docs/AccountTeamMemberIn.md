# AccountTeamMemberIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**AccountTeamRole**](AccountTeamRole.md) |  | 
**Name** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] [default to false]
**NotifyOnEscalation** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewAccountTeamMemberIn

`func NewAccountTeamMemberIn(role AccountTeamRole, name string, ) *AccountTeamMemberIn`

NewAccountTeamMemberIn instantiates a new AccountTeamMemberIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTeamMemberInWithDefaults

`func NewAccountTeamMemberInWithDefaults() *AccountTeamMemberIn`

NewAccountTeamMemberInWithDefaults instantiates a new AccountTeamMemberIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AccountTeamMemberIn) GetRole() AccountTeamRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccountTeamMemberIn) GetRoleOk() (*AccountTeamRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccountTeamMemberIn) SetRole(v AccountTeamRole)`

SetRole sets Role field to given value.


### GetName

`func (o *AccountTeamMemberIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountTeamMemberIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountTeamMemberIn) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *AccountTeamMemberIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountTeamMemberIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountTeamMemberIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountTeamMemberIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsPrimary

`func (o *AccountTeamMemberIn) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *AccountTeamMemberIn) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *AccountTeamMemberIn) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *AccountTeamMemberIn) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetNotifyOnEscalation

`func (o *AccountTeamMemberIn) GetNotifyOnEscalation() bool`

GetNotifyOnEscalation returns the NotifyOnEscalation field if non-nil, zero value otherwise.

### GetNotifyOnEscalationOk

`func (o *AccountTeamMemberIn) GetNotifyOnEscalationOk() (*bool, bool)`

GetNotifyOnEscalationOk returns a tuple with the NotifyOnEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnEscalation

`func (o *AccountTeamMemberIn) SetNotifyOnEscalation(v bool)`

SetNotifyOnEscalation sets NotifyOnEscalation field to given value.

### HasNotifyOnEscalation

`func (o *AccountTeamMemberIn) HasNotifyOnEscalation() bool`

HasNotifyOnEscalation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


