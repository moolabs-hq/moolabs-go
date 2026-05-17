# CostAdjustmentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**OriginalCostEventId** | Pointer to **string** |  | [optional] 
**OriginalCostEventTimestamp** | Pointer to **time.Time** |  | [optional] 
**ExternalAdjustmentId** | Pointer to **string** | Optional client-owned idempotency key for adjustments. When present, server dedups on (tenant_id, external_adjustment_id); re-POST returns the original 201 row. Backward-compatible: existing callers that omit this field get append-only behavior. | [optional] 
**AdjustmentType** | **string** |  | 
**AdjustmentScope** | Pointer to **string** |  | [optional] [default to "event"]
**OriginalTotal** | Pointer to [**NullableOriginalTotal**](OriginalTotal.md) |  | [optional] 
**AdjustedTotal** | Pointer to [**NullableAdjustedTotal**](AdjustedTotal.md) |  | [optional] 
**AdjustmentAmount** | [**AdjustmentAmount**](AdjustmentAmount.md) |  | 
**AdjustmentPeriodStart** | Pointer to **time.Time** |  | [optional] 
**AdjustmentPeriodEnd** | Pointer to **time.Time** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**ApprovedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCostAdjustmentCreate

`func NewCostAdjustmentCreate(tenantId string, adjustmentType string, adjustmentAmount AdjustmentAmount, reason string, ) *CostAdjustmentCreate`

NewCostAdjustmentCreate instantiates a new CostAdjustmentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostAdjustmentCreateWithDefaults

`func NewCostAdjustmentCreateWithDefaults() *CostAdjustmentCreate`

NewCostAdjustmentCreateWithDefaults instantiates a new CostAdjustmentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CostAdjustmentCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CostAdjustmentCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CostAdjustmentCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetOriginalCostEventId

`func (o *CostAdjustmentCreate) GetOriginalCostEventId() string`

GetOriginalCostEventId returns the OriginalCostEventId field if non-nil, zero value otherwise.

### GetOriginalCostEventIdOk

`func (o *CostAdjustmentCreate) GetOriginalCostEventIdOk() (*string, bool)`

GetOriginalCostEventIdOk returns a tuple with the OriginalCostEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCostEventId

`func (o *CostAdjustmentCreate) SetOriginalCostEventId(v string)`

SetOriginalCostEventId sets OriginalCostEventId field to given value.

### HasOriginalCostEventId

`func (o *CostAdjustmentCreate) HasOriginalCostEventId() bool`

HasOriginalCostEventId returns a boolean if a field has been set.

### GetOriginalCostEventTimestamp

`func (o *CostAdjustmentCreate) GetOriginalCostEventTimestamp() time.Time`

GetOriginalCostEventTimestamp returns the OriginalCostEventTimestamp field if non-nil, zero value otherwise.

### GetOriginalCostEventTimestampOk

`func (o *CostAdjustmentCreate) GetOriginalCostEventTimestampOk() (*time.Time, bool)`

GetOriginalCostEventTimestampOk returns a tuple with the OriginalCostEventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCostEventTimestamp

`func (o *CostAdjustmentCreate) SetOriginalCostEventTimestamp(v time.Time)`

SetOriginalCostEventTimestamp sets OriginalCostEventTimestamp field to given value.

### HasOriginalCostEventTimestamp

`func (o *CostAdjustmentCreate) HasOriginalCostEventTimestamp() bool`

HasOriginalCostEventTimestamp returns a boolean if a field has been set.

### GetExternalAdjustmentId

`func (o *CostAdjustmentCreate) GetExternalAdjustmentId() string`

GetExternalAdjustmentId returns the ExternalAdjustmentId field if non-nil, zero value otherwise.

### GetExternalAdjustmentIdOk

`func (o *CostAdjustmentCreate) GetExternalAdjustmentIdOk() (*string, bool)`

GetExternalAdjustmentIdOk returns a tuple with the ExternalAdjustmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAdjustmentId

`func (o *CostAdjustmentCreate) SetExternalAdjustmentId(v string)`

SetExternalAdjustmentId sets ExternalAdjustmentId field to given value.

### HasExternalAdjustmentId

`func (o *CostAdjustmentCreate) HasExternalAdjustmentId() bool`

HasExternalAdjustmentId returns a boolean if a field has been set.

### GetAdjustmentType

`func (o *CostAdjustmentCreate) GetAdjustmentType() string`

GetAdjustmentType returns the AdjustmentType field if non-nil, zero value otherwise.

### GetAdjustmentTypeOk

`func (o *CostAdjustmentCreate) GetAdjustmentTypeOk() (*string, bool)`

GetAdjustmentTypeOk returns a tuple with the AdjustmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentType

`func (o *CostAdjustmentCreate) SetAdjustmentType(v string)`

SetAdjustmentType sets AdjustmentType field to given value.


### GetAdjustmentScope

`func (o *CostAdjustmentCreate) GetAdjustmentScope() string`

GetAdjustmentScope returns the AdjustmentScope field if non-nil, zero value otherwise.

### GetAdjustmentScopeOk

`func (o *CostAdjustmentCreate) GetAdjustmentScopeOk() (*string, bool)`

GetAdjustmentScopeOk returns a tuple with the AdjustmentScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentScope

`func (o *CostAdjustmentCreate) SetAdjustmentScope(v string)`

SetAdjustmentScope sets AdjustmentScope field to given value.

### HasAdjustmentScope

`func (o *CostAdjustmentCreate) HasAdjustmentScope() bool`

HasAdjustmentScope returns a boolean if a field has been set.

### GetOriginalTotal

`func (o *CostAdjustmentCreate) GetOriginalTotal() OriginalTotal`

GetOriginalTotal returns the OriginalTotal field if non-nil, zero value otherwise.

### GetOriginalTotalOk

`func (o *CostAdjustmentCreate) GetOriginalTotalOk() (*OriginalTotal, bool)`

GetOriginalTotalOk returns a tuple with the OriginalTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTotal

`func (o *CostAdjustmentCreate) SetOriginalTotal(v OriginalTotal)`

SetOriginalTotal sets OriginalTotal field to given value.

### HasOriginalTotal

`func (o *CostAdjustmentCreate) HasOriginalTotal() bool`

HasOriginalTotal returns a boolean if a field has been set.

### SetOriginalTotalNil

`func (o *CostAdjustmentCreate) SetOriginalTotalNil(b bool)`

 SetOriginalTotalNil sets the value for OriginalTotal to be an explicit nil

### UnsetOriginalTotal
`func (o *CostAdjustmentCreate) UnsetOriginalTotal()`

UnsetOriginalTotal ensures that no value is present for OriginalTotal, not even an explicit nil
### GetAdjustedTotal

`func (o *CostAdjustmentCreate) GetAdjustedTotal() AdjustedTotal`

GetAdjustedTotal returns the AdjustedTotal field if non-nil, zero value otherwise.

### GetAdjustedTotalOk

`func (o *CostAdjustmentCreate) GetAdjustedTotalOk() (*AdjustedTotal, bool)`

GetAdjustedTotalOk returns a tuple with the AdjustedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTotal

`func (o *CostAdjustmentCreate) SetAdjustedTotal(v AdjustedTotal)`

SetAdjustedTotal sets AdjustedTotal field to given value.

### HasAdjustedTotal

`func (o *CostAdjustmentCreate) HasAdjustedTotal() bool`

HasAdjustedTotal returns a boolean if a field has been set.

### SetAdjustedTotalNil

`func (o *CostAdjustmentCreate) SetAdjustedTotalNil(b bool)`

 SetAdjustedTotalNil sets the value for AdjustedTotal to be an explicit nil

### UnsetAdjustedTotal
`func (o *CostAdjustmentCreate) UnsetAdjustedTotal()`

UnsetAdjustedTotal ensures that no value is present for AdjustedTotal, not even an explicit nil
### GetAdjustmentAmount

`func (o *CostAdjustmentCreate) GetAdjustmentAmount() AdjustmentAmount`

GetAdjustmentAmount returns the AdjustmentAmount field if non-nil, zero value otherwise.

### GetAdjustmentAmountOk

`func (o *CostAdjustmentCreate) GetAdjustmentAmountOk() (*AdjustmentAmount, bool)`

GetAdjustmentAmountOk returns a tuple with the AdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmount

`func (o *CostAdjustmentCreate) SetAdjustmentAmount(v AdjustmentAmount)`

SetAdjustmentAmount sets AdjustmentAmount field to given value.


### GetAdjustmentPeriodStart

`func (o *CostAdjustmentCreate) GetAdjustmentPeriodStart() time.Time`

GetAdjustmentPeriodStart returns the AdjustmentPeriodStart field if non-nil, zero value otherwise.

### GetAdjustmentPeriodStartOk

`func (o *CostAdjustmentCreate) GetAdjustmentPeriodStartOk() (*time.Time, bool)`

GetAdjustmentPeriodStartOk returns a tuple with the AdjustmentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentPeriodStart

`func (o *CostAdjustmentCreate) SetAdjustmentPeriodStart(v time.Time)`

SetAdjustmentPeriodStart sets AdjustmentPeriodStart field to given value.

### HasAdjustmentPeriodStart

`func (o *CostAdjustmentCreate) HasAdjustmentPeriodStart() bool`

HasAdjustmentPeriodStart returns a boolean if a field has been set.

### GetAdjustmentPeriodEnd

`func (o *CostAdjustmentCreate) GetAdjustmentPeriodEnd() time.Time`

GetAdjustmentPeriodEnd returns the AdjustmentPeriodEnd field if non-nil, zero value otherwise.

### GetAdjustmentPeriodEndOk

`func (o *CostAdjustmentCreate) GetAdjustmentPeriodEndOk() (*time.Time, bool)`

GetAdjustmentPeriodEndOk returns a tuple with the AdjustmentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentPeriodEnd

`func (o *CostAdjustmentCreate) SetAdjustmentPeriodEnd(v time.Time)`

SetAdjustmentPeriodEnd sets AdjustmentPeriodEnd field to given value.

### HasAdjustmentPeriodEnd

`func (o *CostAdjustmentCreate) HasAdjustmentPeriodEnd() bool`

HasAdjustmentPeriodEnd returns a boolean if a field has been set.

### GetProvider

`func (o *CostAdjustmentCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CostAdjustmentCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CostAdjustmentCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CostAdjustmentCreate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetModel

`func (o *CostAdjustmentCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CostAdjustmentCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CostAdjustmentCreate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CostAdjustmentCreate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetReason

`func (o *CostAdjustmentCreate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CostAdjustmentCreate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CostAdjustmentCreate) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSource

`func (o *CostAdjustmentCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CostAdjustmentCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CostAdjustmentCreate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CostAdjustmentCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetApprovedBy

`func (o *CostAdjustmentCreate) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *CostAdjustmentCreate) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *CostAdjustmentCreate) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *CostAdjustmentCreate) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


