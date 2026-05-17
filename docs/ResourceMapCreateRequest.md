# ResourceMapCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** |  | 
**ResourceId** | **string** |  | 
**ResourceType** | Pointer to **string** |  | [optional] 
**ServiceName** | **string** |  | 
**TeamName** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResourceMapCreateRequest

`func NewResourceMapCreateRequest(cloudProvider string, resourceId string, serviceName string, ) *ResourceMapCreateRequest`

NewResourceMapCreateRequest instantiates a new ResourceMapCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMapCreateRequestWithDefaults

`func NewResourceMapCreateRequestWithDefaults() *ResourceMapCreateRequest`

NewResourceMapCreateRequestWithDefaults instantiates a new ResourceMapCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ResourceMapCreateRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ResourceMapCreateRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ResourceMapCreateRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetResourceId

`func (o *ResourceMapCreateRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceMapCreateRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceMapCreateRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ResourceMapCreateRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceMapCreateRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceMapCreateRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceMapCreateRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetServiceName

`func (o *ResourceMapCreateRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceMapCreateRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceMapCreateRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetTeamName

`func (o *ResourceMapCreateRequest) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ResourceMapCreateRequest) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ResourceMapCreateRequest) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ResourceMapCreateRequest) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetEnvironment

`func (o *ResourceMapCreateRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceMapCreateRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceMapCreateRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ResourceMapCreateRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTags

`func (o *ResourceMapCreateRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceMapCreateRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceMapCreateRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourceMapCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


