# SnapshotComputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** | Deprecated. Tenant is derived from API key. If provided, must match authenticated tenant. | [optional] 
**PeriodType** | Pointer to **string** |  | [optional] [default to "daily"]
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 

## Methods

### NewSnapshotComputeRequest

`func NewSnapshotComputeRequest(periodStart time.Time, periodEnd time.Time, ) *SnapshotComputeRequest`

NewSnapshotComputeRequest instantiates a new SnapshotComputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotComputeRequestWithDefaults

`func NewSnapshotComputeRequestWithDefaults() *SnapshotComputeRequest`

NewSnapshotComputeRequestWithDefaults instantiates a new SnapshotComputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SnapshotComputeRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SnapshotComputeRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SnapshotComputeRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SnapshotComputeRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPeriodType

`func (o *SnapshotComputeRequest) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *SnapshotComputeRequest) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *SnapshotComputeRequest) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.

### HasPeriodType

`func (o *SnapshotComputeRequest) HasPeriodType() bool`

HasPeriodType returns a boolean if a field has been set.

### GetPeriodStart

`func (o *SnapshotComputeRequest) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *SnapshotComputeRequest) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *SnapshotComputeRequest) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *SnapshotComputeRequest) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *SnapshotComputeRequest) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *SnapshotComputeRequest) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


