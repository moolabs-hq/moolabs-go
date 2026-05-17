# ReplayDeadLetterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** | Optional tenant override | [optional] 
**PoolId** | Pointer to **string** | Optional pool override | [optional] 

## Methods

### NewReplayDeadLetterRequest

`func NewReplayDeadLetterRequest() *ReplayDeadLetterRequest`

NewReplayDeadLetterRequest instantiates a new ReplayDeadLetterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayDeadLetterRequestWithDefaults

`func NewReplayDeadLetterRequestWithDefaults() *ReplayDeadLetterRequest`

NewReplayDeadLetterRequestWithDefaults instantiates a new ReplayDeadLetterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ReplayDeadLetterRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ReplayDeadLetterRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ReplayDeadLetterRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ReplayDeadLetterRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPoolId

`func (o *ReplayDeadLetterRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ReplayDeadLetterRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ReplayDeadLetterRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ReplayDeadLetterRequest) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


