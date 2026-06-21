# ArcDunningTemplateAdminUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**OrgId** | **string** |  | 
**Permission** | **string** |  | 
**CanManage** | **bool** |  | 
**Users** | [**[]ArcDunningTemplateAdminUserDto**](ArcDunningTemplateAdminUserDto.md) |  | 

## Methods

### NewArcDunningTemplateAdminUsersResponse

`func NewArcDunningTemplateAdminUsersResponse(tenantId string, orgId string, permission string, canManage bool, users []ArcDunningTemplateAdminUserDto, ) *ArcDunningTemplateAdminUsersResponse`

NewArcDunningTemplateAdminUsersResponse instantiates a new ArcDunningTemplateAdminUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcDunningTemplateAdminUsersResponseWithDefaults

`func NewArcDunningTemplateAdminUsersResponseWithDefaults() *ArcDunningTemplateAdminUsersResponse`

NewArcDunningTemplateAdminUsersResponseWithDefaults instantiates a new ArcDunningTemplateAdminUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ArcDunningTemplateAdminUsersResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ArcDunningTemplateAdminUsersResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ArcDunningTemplateAdminUsersResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetOrgId

`func (o *ArcDunningTemplateAdminUsersResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ArcDunningTemplateAdminUsersResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ArcDunningTemplateAdminUsersResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetPermission

`func (o *ArcDunningTemplateAdminUsersResponse) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ArcDunningTemplateAdminUsersResponse) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ArcDunningTemplateAdminUsersResponse) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetCanManage

`func (o *ArcDunningTemplateAdminUsersResponse) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ArcDunningTemplateAdminUsersResponse) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ArcDunningTemplateAdminUsersResponse) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.


### GetUsers

`func (o *ArcDunningTemplateAdminUsersResponse) GetUsers() []ArcDunningTemplateAdminUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ArcDunningTemplateAdminUsersResponse) GetUsersOk() (*[]ArcDunningTemplateAdminUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ArcDunningTemplateAdminUsersResponse) SetUsers(v []ArcDunningTemplateAdminUserDto)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


