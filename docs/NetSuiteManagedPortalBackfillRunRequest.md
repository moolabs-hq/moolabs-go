# NetSuiteManagedPortalBackfillRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**MaxJobs** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewNetSuiteManagedPortalBackfillRunRequest

`func NewNetSuiteManagedPortalBackfillRunRequest() *NetSuiteManagedPortalBackfillRunRequest`

NewNetSuiteManagedPortalBackfillRunRequest instantiates a new NetSuiteManagedPortalBackfillRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetSuiteManagedPortalBackfillRunRequestWithDefaults

`func NewNetSuiteManagedPortalBackfillRunRequestWithDefaults() *NetSuiteManagedPortalBackfillRunRequest`

NewNetSuiteManagedPortalBackfillRunRequestWithDefaults instantiates a new NetSuiteManagedPortalBackfillRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *NetSuiteManagedPortalBackfillRunRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetSuiteManagedPortalBackfillRunRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetSuiteManagedPortalBackfillRunRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *NetSuiteManagedPortalBackfillRunRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetMaxJobs

`func (o *NetSuiteManagedPortalBackfillRunRequest) GetMaxJobs() int32`

GetMaxJobs returns the MaxJobs field if non-nil, zero value otherwise.

### GetMaxJobsOk

`func (o *NetSuiteManagedPortalBackfillRunRequest) GetMaxJobsOk() (*int32, bool)`

GetMaxJobsOk returns a tuple with the MaxJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobs

`func (o *NetSuiteManagedPortalBackfillRunRequest) SetMaxJobs(v int32)`

SetMaxJobs sets MaxJobs field to given value.

### HasMaxJobs

`func (o *NetSuiteManagedPortalBackfillRunRequest) HasMaxJobs() bool`

HasMaxJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


