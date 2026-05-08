# RetryUnpricedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**UsageEventId** | **string** | Usage event ID to retry | 

## Methods

### NewRetryUnpricedRequest

`func NewRetryUnpricedRequest(tenantId string, poolId string, usageEventId string, ) *RetryUnpricedRequest`

NewRetryUnpricedRequest instantiates a new RetryUnpricedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryUnpricedRequestWithDefaults

`func NewRetryUnpricedRequestWithDefaults() *RetryUnpricedRequest`

NewRetryUnpricedRequestWithDefaults instantiates a new RetryUnpricedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RetryUnpricedRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RetryUnpricedRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RetryUnpricedRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *RetryUnpricedRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *RetryUnpricedRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *RetryUnpricedRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetUsageEventId

`func (o *RetryUnpricedRequest) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *RetryUnpricedRequest) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *RetryUnpricedRequest) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


