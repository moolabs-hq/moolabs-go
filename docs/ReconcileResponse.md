# ReconcileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**TenantId** | **string** |  | 
**Status** | **string** |  | 
**CostEventIdsCount** | **int32** |  | 
**Message** | **string** |  | 

## Methods

### NewReconcileResponse

`func NewReconcileResponse(jobId string, tenantId string, status string, costEventIdsCount int32, message string, ) *ReconcileResponse`

NewReconcileResponse instantiates a new ReconcileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconcileResponseWithDefaults

`func NewReconcileResponseWithDefaults() *ReconcileResponse`

NewReconcileResponseWithDefaults instantiates a new ReconcileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ReconcileResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ReconcileResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ReconcileResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetTenantId

`func (o *ReconcileResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ReconcileResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ReconcileResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetStatus

`func (o *ReconcileResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReconcileResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReconcileResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCostEventIdsCount

`func (o *ReconcileResponse) GetCostEventIdsCount() int32`

GetCostEventIdsCount returns the CostEventIdsCount field if non-nil, zero value otherwise.

### GetCostEventIdsCountOk

`func (o *ReconcileResponse) GetCostEventIdsCountOk() (*int32, bool)`

GetCostEventIdsCountOk returns a tuple with the CostEventIdsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostEventIdsCount

`func (o *ReconcileResponse) SetCostEventIdsCount(v int32)`

SetCostEventIdsCount sets CostEventIdsCount field to given value.


### GetMessage

`func (o *ReconcileResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReconcileResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReconcileResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


