# EffectivePermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**OrgId** | **string** |  | 
**ActorUserId** | **string** |  | 
**ActorDisplay** | Pointer to **string** |  | [optional] 
**ActorEmail** | Pointer to **string** |  | [optional] 
**OrgRole** | Pointer to **string** |  | [optional] 
**IsOrgAdmin** | **bool** |  | 
**Permissions** | **[]string** |  | 
**Grants** | [**[]EffectivePermissionGrantDto**](EffectivePermissionGrantDto.md) |  | 

## Methods

### NewEffectivePermissionsResponse

`func NewEffectivePermissionsResponse(tenantId string, orgId string, actorUserId string, isOrgAdmin bool, permissions []string, grants []EffectivePermissionGrantDto, ) *EffectivePermissionsResponse`

NewEffectivePermissionsResponse instantiates a new EffectivePermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectivePermissionsResponseWithDefaults

`func NewEffectivePermissionsResponseWithDefaults() *EffectivePermissionsResponse`

NewEffectivePermissionsResponseWithDefaults instantiates a new EffectivePermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *EffectivePermissionsResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EffectivePermissionsResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EffectivePermissionsResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetOrgId

`func (o *EffectivePermissionsResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EffectivePermissionsResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EffectivePermissionsResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetActorUserId

`func (o *EffectivePermissionsResponse) GetActorUserId() string`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *EffectivePermissionsResponse) GetActorUserIdOk() (*string, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *EffectivePermissionsResponse) SetActorUserId(v string)`

SetActorUserId sets ActorUserId field to given value.


### GetActorDisplay

`func (o *EffectivePermissionsResponse) GetActorDisplay() string`

GetActorDisplay returns the ActorDisplay field if non-nil, zero value otherwise.

### GetActorDisplayOk

`func (o *EffectivePermissionsResponse) GetActorDisplayOk() (*string, bool)`

GetActorDisplayOk returns a tuple with the ActorDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplay

`func (o *EffectivePermissionsResponse) SetActorDisplay(v string)`

SetActorDisplay sets ActorDisplay field to given value.

### HasActorDisplay

`func (o *EffectivePermissionsResponse) HasActorDisplay() bool`

HasActorDisplay returns a boolean if a field has been set.

### GetActorEmail

`func (o *EffectivePermissionsResponse) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *EffectivePermissionsResponse) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *EffectivePermissionsResponse) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *EffectivePermissionsResponse) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### GetOrgRole

`func (o *EffectivePermissionsResponse) GetOrgRole() string`

GetOrgRole returns the OrgRole field if non-nil, zero value otherwise.

### GetOrgRoleOk

`func (o *EffectivePermissionsResponse) GetOrgRoleOk() (*string, bool)`

GetOrgRoleOk returns a tuple with the OrgRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRole

`func (o *EffectivePermissionsResponse) SetOrgRole(v string)`

SetOrgRole sets OrgRole field to given value.

### HasOrgRole

`func (o *EffectivePermissionsResponse) HasOrgRole() bool`

HasOrgRole returns a boolean if a field has been set.

### GetIsOrgAdmin

`func (o *EffectivePermissionsResponse) GetIsOrgAdmin() bool`

GetIsOrgAdmin returns the IsOrgAdmin field if non-nil, zero value otherwise.

### GetIsOrgAdminOk

`func (o *EffectivePermissionsResponse) GetIsOrgAdminOk() (*bool, bool)`

GetIsOrgAdminOk returns a tuple with the IsOrgAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgAdmin

`func (o *EffectivePermissionsResponse) SetIsOrgAdmin(v bool)`

SetIsOrgAdmin sets IsOrgAdmin field to given value.


### GetPermissions

`func (o *EffectivePermissionsResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EffectivePermissionsResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EffectivePermissionsResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetGrants

`func (o *EffectivePermissionsResponse) GetGrants() []EffectivePermissionGrantDto`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *EffectivePermissionsResponse) GetGrantsOk() (*[]EffectivePermissionGrantDto, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *EffectivePermissionsResponse) SetGrants(v []EffectivePermissionGrantDto)`

SetGrants sets Grants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


