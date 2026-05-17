# CloudCostRowInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**ResourceId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 
**Cost** | [**Cost**](Cost.md) |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCloudCostRowInput

`func NewCloudCostRowInput(serviceName string, cost Cost, ) *CloudCostRowInput`

NewCloudCostRowInput instantiates a new CloudCostRowInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCostRowInputWithDefaults

`func NewCloudCostRowInputWithDefaults() *CloudCostRowInput`

NewCloudCostRowInputWithDefaults instantiates a new CloudCostRowInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *CloudCostRowInput) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CloudCostRowInput) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CloudCostRowInput) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetResourceId

`func (o *CloudCostRowInput) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CloudCostRowInput) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CloudCostRowInput) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CloudCostRowInput) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetRegion

`func (o *CloudCostRowInput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudCostRowInput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudCostRowInput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudCostRowInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUsageType

`func (o *CloudCostRowInput) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *CloudCostRowInput) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *CloudCostRowInput) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *CloudCostRowInput) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetCost

`func (o *CloudCostRowInput) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CloudCostRowInput) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CloudCostRowInput) SetCost(v Cost)`

SetCost sets Cost field to given value.


### GetCurrency

`func (o *CloudCostRowInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudCostRowInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudCostRowInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudCostRowInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTags

`func (o *CloudCostRowInput) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudCostRowInput) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudCostRowInput) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudCostRowInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


