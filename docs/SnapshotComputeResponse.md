# SnapshotComputeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**TenantId** | **string** |  | 
**PeriodType** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**Status** | **string** |  | 
**Snapshot** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSnapshotComputeResponse

`func NewSnapshotComputeResponse(jobId string, tenantId string, periodType string, periodStart string, periodEnd string, status string, ) *SnapshotComputeResponse`

NewSnapshotComputeResponse instantiates a new SnapshotComputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotComputeResponseWithDefaults

`func NewSnapshotComputeResponseWithDefaults() *SnapshotComputeResponse`

NewSnapshotComputeResponseWithDefaults instantiates a new SnapshotComputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *SnapshotComputeResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *SnapshotComputeResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *SnapshotComputeResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetTenantId

`func (o *SnapshotComputeResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SnapshotComputeResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SnapshotComputeResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPeriodType

`func (o *SnapshotComputeResponse) GetPeriodType() string`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *SnapshotComputeResponse) GetPeriodTypeOk() (*string, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *SnapshotComputeResponse) SetPeriodType(v string)`

SetPeriodType sets PeriodType field to given value.


### GetPeriodStart

`func (o *SnapshotComputeResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *SnapshotComputeResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *SnapshotComputeResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *SnapshotComputeResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *SnapshotComputeResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *SnapshotComputeResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetStatus

`func (o *SnapshotComputeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotComputeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotComputeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSnapshot

`func (o *SnapshotComputeResponse) GetSnapshot() map[string]interface{}`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SnapshotComputeResponse) GetSnapshotOk() (*map[string]interface{}, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SnapshotComputeResponse) SetSnapshot(v map[string]interface{})`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SnapshotComputeResponse) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


