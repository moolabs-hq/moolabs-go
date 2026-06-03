# ArcDunningPermissionGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**OrgId** | **string** |  | 
**UserSubject** | **string** |  | 
**GrantReason** | **string** |  | 
**IdempotencyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewArcDunningPermissionGrantRequest

`func NewArcDunningPermissionGrantRequest(tenantId string, orgId string, userSubject string, grantReason string, ) *ArcDunningPermissionGrantRequest`

NewArcDunningPermissionGrantRequest instantiates a new ArcDunningPermissionGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcDunningPermissionGrantRequestWithDefaults

`func NewArcDunningPermissionGrantRequestWithDefaults() *ArcDunningPermissionGrantRequest`

NewArcDunningPermissionGrantRequestWithDefaults instantiates a new ArcDunningPermissionGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ArcDunningPermissionGrantRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ArcDunningPermissionGrantRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ArcDunningPermissionGrantRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetOrgId

`func (o *ArcDunningPermissionGrantRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ArcDunningPermissionGrantRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ArcDunningPermissionGrantRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUserSubject

`func (o *ArcDunningPermissionGrantRequest) GetUserSubject() string`

GetUserSubject returns the UserSubject field if non-nil, zero value otherwise.

### GetUserSubjectOk

`func (o *ArcDunningPermissionGrantRequest) GetUserSubjectOk() (*string, bool)`

GetUserSubjectOk returns a tuple with the UserSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubject

`func (o *ArcDunningPermissionGrantRequest) SetUserSubject(v string)`

SetUserSubject sets UserSubject field to given value.


### GetGrantReason

`func (o *ArcDunningPermissionGrantRequest) GetGrantReason() string`

GetGrantReason returns the GrantReason field if non-nil, zero value otherwise.

### GetGrantReasonOk

`func (o *ArcDunningPermissionGrantRequest) GetGrantReasonOk() (*string, bool)`

GetGrantReasonOk returns a tuple with the GrantReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantReason

`func (o *ArcDunningPermissionGrantRequest) SetGrantReason(v string)`

SetGrantReason sets GrantReason field to given value.


### GetIdempotencyKey

`func (o *ArcDunningPermissionGrantRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ArcDunningPermissionGrantRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ArcDunningPermissionGrantRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ArcDunningPermissionGrantRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


