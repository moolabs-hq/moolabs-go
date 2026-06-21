# ArcDunningTemplateAdminUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSubject** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**OrgRole** | Pointer to [**NullableOrgRole**](OrgRole.md) |  | [optional] 
**EligibleForGrant** | Pointer to **bool** |  | [optional] [default to false]
**GrantBlockedReason** | Pointer to **string** |  | [optional] 
**Grant** | Pointer to [**ArcDunningTemplateAdminUserGrantDto**](ArcDunningTemplateAdminUserGrantDto.md) |  | [optional] 

## Methods

### NewArcDunningTemplateAdminUserDto

`func NewArcDunningTemplateAdminUserDto(userSubject string, ) *ArcDunningTemplateAdminUserDto`

NewArcDunningTemplateAdminUserDto instantiates a new ArcDunningTemplateAdminUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcDunningTemplateAdminUserDtoWithDefaults

`func NewArcDunningTemplateAdminUserDtoWithDefaults() *ArcDunningTemplateAdminUserDto`

NewArcDunningTemplateAdminUserDtoWithDefaults instantiates a new ArcDunningTemplateAdminUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSubject

`func (o *ArcDunningTemplateAdminUserDto) GetUserSubject() string`

GetUserSubject returns the UserSubject field if non-nil, zero value otherwise.

### GetUserSubjectOk

`func (o *ArcDunningTemplateAdminUserDto) GetUserSubjectOk() (*string, bool)`

GetUserSubjectOk returns a tuple with the UserSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubject

`func (o *ArcDunningTemplateAdminUserDto) SetUserSubject(v string)`

SetUserSubject sets UserSubject field to given value.


### GetName

`func (o *ArcDunningTemplateAdminUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArcDunningTemplateAdminUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArcDunningTemplateAdminUserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArcDunningTemplateAdminUserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ArcDunningTemplateAdminUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ArcDunningTemplateAdminUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ArcDunningTemplateAdminUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ArcDunningTemplateAdminUserDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgRole

`func (o *ArcDunningTemplateAdminUserDto) GetOrgRole() OrgRole`

GetOrgRole returns the OrgRole field if non-nil, zero value otherwise.

### GetOrgRoleOk

`func (o *ArcDunningTemplateAdminUserDto) GetOrgRoleOk() (*OrgRole, bool)`

GetOrgRoleOk returns a tuple with the OrgRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRole

`func (o *ArcDunningTemplateAdminUserDto) SetOrgRole(v OrgRole)`

SetOrgRole sets OrgRole field to given value.

### HasOrgRole

`func (o *ArcDunningTemplateAdminUserDto) HasOrgRole() bool`

HasOrgRole returns a boolean if a field has been set.

### SetOrgRoleNil

`func (o *ArcDunningTemplateAdminUserDto) SetOrgRoleNil(b bool)`

 SetOrgRoleNil sets the value for OrgRole to be an explicit nil

### UnsetOrgRole
`func (o *ArcDunningTemplateAdminUserDto) UnsetOrgRole()`

UnsetOrgRole ensures that no value is present for OrgRole, not even an explicit nil
### GetEligibleForGrant

`func (o *ArcDunningTemplateAdminUserDto) GetEligibleForGrant() bool`

GetEligibleForGrant returns the EligibleForGrant field if non-nil, zero value otherwise.

### GetEligibleForGrantOk

`func (o *ArcDunningTemplateAdminUserDto) GetEligibleForGrantOk() (*bool, bool)`

GetEligibleForGrantOk returns a tuple with the EligibleForGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleForGrant

`func (o *ArcDunningTemplateAdminUserDto) SetEligibleForGrant(v bool)`

SetEligibleForGrant sets EligibleForGrant field to given value.

### HasEligibleForGrant

`func (o *ArcDunningTemplateAdminUserDto) HasEligibleForGrant() bool`

HasEligibleForGrant returns a boolean if a field has been set.

### GetGrantBlockedReason

`func (o *ArcDunningTemplateAdminUserDto) GetGrantBlockedReason() string`

GetGrantBlockedReason returns the GrantBlockedReason field if non-nil, zero value otherwise.

### GetGrantBlockedReasonOk

`func (o *ArcDunningTemplateAdminUserDto) GetGrantBlockedReasonOk() (*string, bool)`

GetGrantBlockedReasonOk returns a tuple with the GrantBlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantBlockedReason

`func (o *ArcDunningTemplateAdminUserDto) SetGrantBlockedReason(v string)`

SetGrantBlockedReason sets GrantBlockedReason field to given value.

### HasGrantBlockedReason

`func (o *ArcDunningTemplateAdminUserDto) HasGrantBlockedReason() bool`

HasGrantBlockedReason returns a boolean if a field has been set.

### GetGrant

`func (o *ArcDunningTemplateAdminUserDto) GetGrant() ArcDunningTemplateAdminUserGrantDto`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *ArcDunningTemplateAdminUserDto) GetGrantOk() (*ArcDunningTemplateAdminUserGrantDto, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *ArcDunningTemplateAdminUserDto) SetGrant(v ArcDunningTemplateAdminUserGrantDto)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *ArcDunningTemplateAdminUserDto) HasGrant() bool`

HasGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


