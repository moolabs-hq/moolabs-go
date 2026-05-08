# GetFxRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**AsOf** | **time.Time** | Timestamp to get rate for | 

## Methods

### NewGetFxRateRequest

`func NewGetFxRateRequest(tenantId string, poolId string, asOf time.Time, ) *GetFxRateRequest`

NewGetFxRateRequest instantiates a new GetFxRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFxRateRequestWithDefaults

`func NewGetFxRateRequestWithDefaults() *GetFxRateRequest`

NewGetFxRateRequestWithDefaults instantiates a new GetFxRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *GetFxRateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetFxRateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetFxRateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *GetFxRateRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GetFxRateRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GetFxRateRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAsOf

`func (o *GetFxRateRequest) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GetFxRateRequest) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GetFxRateRequest) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


