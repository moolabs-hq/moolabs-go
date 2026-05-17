# ResourceMapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CloudProvider** | **string** |  | 
**ResourceId** | **string** |  | 
**ResourceType** | **string** |  | 
**ServiceName** | **string** |  | 
**TeamName** | **string** |  | 
**Environment** | **string** |  | 
**Tags** | **map[string]interface{}** |  | 
**IsActive** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewResourceMapResponse

`func NewResourceMapResponse(id string, tenantId string, cloudProvider string, resourceId string, resourceType string, serviceName string, teamName string, environment string, tags map[string]interface{}, isActive bool, createdAt string, updatedAt string, ) *ResourceMapResponse`

NewResourceMapResponse instantiates a new ResourceMapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMapResponseWithDefaults

`func NewResourceMapResponseWithDefaults() *ResourceMapResponse`

NewResourceMapResponseWithDefaults instantiates a new ResourceMapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceMapResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceMapResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceMapResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ResourceMapResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ResourceMapResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ResourceMapResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCloudProvider

`func (o *ResourceMapResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ResourceMapResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ResourceMapResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetResourceId

`func (o *ResourceMapResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceMapResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceMapResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ResourceMapResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceMapResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceMapResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetServiceName

`func (o *ResourceMapResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceMapResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceMapResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetTeamName

`func (o *ResourceMapResponse) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ResourceMapResponse) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ResourceMapResponse) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetEnvironment

`func (o *ResourceMapResponse) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceMapResponse) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceMapResponse) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetTags

`func (o *ResourceMapResponse) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceMapResponse) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceMapResponse) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.


### GetIsActive

`func (o *ResourceMapResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ResourceMapResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ResourceMapResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedAt

`func (o *ResourceMapResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceMapResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceMapResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourceMapResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceMapResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceMapResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


