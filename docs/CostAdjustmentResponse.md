# CostAdjustmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**OriginalCostEventId** | **string** |  | 
**OriginalCostEventTimestamp** | **time.Time** |  | 
**AdjustmentType** | **string** |  | 
**AdjustmentAmount** | **string** |  | 
**Reason** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewCostAdjustmentResponse

`func NewCostAdjustmentResponse(id string, tenantId string, originalCostEventId string, originalCostEventTimestamp time.Time, adjustmentType string, adjustmentAmount string, reason string, createdAt time.Time, ) *CostAdjustmentResponse`

NewCostAdjustmentResponse instantiates a new CostAdjustmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostAdjustmentResponseWithDefaults

`func NewCostAdjustmentResponseWithDefaults() *CostAdjustmentResponse`

NewCostAdjustmentResponseWithDefaults instantiates a new CostAdjustmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CostAdjustmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CostAdjustmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CostAdjustmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CostAdjustmentResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostAdjustmentResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostAdjustmentResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetOriginalCostEventId

`func (o *CostAdjustmentResponse) GetOriginalCostEventId() string`

GetOriginalCostEventId returns the OriginalCostEventId field if non-nil, zero value otherwise.

### GetOriginalCostEventIdOk

`func (o *CostAdjustmentResponse) GetOriginalCostEventIdOk() (*string, bool)`

GetOriginalCostEventIdOk returns a tuple with the OriginalCostEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCostEventId

`func (o *CostAdjustmentResponse) SetOriginalCostEventId(v string)`

SetOriginalCostEventId sets OriginalCostEventId field to given value.


### GetOriginalCostEventTimestamp

`func (o *CostAdjustmentResponse) GetOriginalCostEventTimestamp() time.Time`

GetOriginalCostEventTimestamp returns the OriginalCostEventTimestamp field if non-nil, zero value otherwise.

### GetOriginalCostEventTimestampOk

`func (o *CostAdjustmentResponse) GetOriginalCostEventTimestampOk() (*time.Time, bool)`

GetOriginalCostEventTimestampOk returns a tuple with the OriginalCostEventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCostEventTimestamp

`func (o *CostAdjustmentResponse) SetOriginalCostEventTimestamp(v time.Time)`

SetOriginalCostEventTimestamp sets OriginalCostEventTimestamp field to given value.


### GetAdjustmentType

`func (o *CostAdjustmentResponse) GetAdjustmentType() string`

GetAdjustmentType returns the AdjustmentType field if non-nil, zero value otherwise.

### GetAdjustmentTypeOk

`func (o *CostAdjustmentResponse) GetAdjustmentTypeOk() (*string, bool)`

GetAdjustmentTypeOk returns a tuple with the AdjustmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentType

`func (o *CostAdjustmentResponse) SetAdjustmentType(v string)`

SetAdjustmentType sets AdjustmentType field to given value.


### GetAdjustmentAmount

`func (o *CostAdjustmentResponse) GetAdjustmentAmount() string`

GetAdjustmentAmount returns the AdjustmentAmount field if non-nil, zero value otherwise.

### GetAdjustmentAmountOk

`func (o *CostAdjustmentResponse) GetAdjustmentAmountOk() (*string, bool)`

GetAdjustmentAmountOk returns a tuple with the AdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmount

`func (o *CostAdjustmentResponse) SetAdjustmentAmount(v string)`

SetAdjustmentAmount sets AdjustmentAmount field to given value.


### GetReason

`func (o *CostAdjustmentResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CostAdjustmentResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CostAdjustmentResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreatedAt

`func (o *CostAdjustmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CostAdjustmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CostAdjustmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


