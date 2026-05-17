# TenantProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantProvisionRequest

`func NewTenantProvisionRequest(orgId string, ) *TenantProvisionRequest`

NewTenantProvisionRequest instantiates a new TenantProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantProvisionRequestWithDefaults

`func NewTenantProvisionRequestWithDefaults() *TenantProvisionRequest`

NewTenantProvisionRequestWithDefaults instantiates a new TenantProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *TenantProvisionRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TenantProvisionRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TenantProvisionRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetName

`func (o *TenantProvisionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantProvisionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantProvisionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantProvisionRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


