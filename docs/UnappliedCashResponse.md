# UnappliedCashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**RemittanceId** | **string** |  | 
**OriginalAmountMicros** | **int32** |  | 
**RemainingAmountMicros** | **int32** |  | 
**ReasonCode** | **string** |  | 
**Status** | **string** |  | 
**DispositionType** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUnappliedCashResponse

`func NewUnappliedCashResponse(id string, tenantId string, remittanceId string, originalAmountMicros int32, remainingAmountMicros int32, reasonCode string, status string, createdAt time.Time, updatedAt time.Time, ) *UnappliedCashResponse`

NewUnappliedCashResponse instantiates a new UnappliedCashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnappliedCashResponseWithDefaults

`func NewUnappliedCashResponseWithDefaults() *UnappliedCashResponse`

NewUnappliedCashResponseWithDefaults instantiates a new UnappliedCashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnappliedCashResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnappliedCashResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnappliedCashResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *UnappliedCashResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UnappliedCashResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UnappliedCashResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetRemittanceId

`func (o *UnappliedCashResponse) GetRemittanceId() string`

GetRemittanceId returns the RemittanceId field if non-nil, zero value otherwise.

### GetRemittanceIdOk

`func (o *UnappliedCashResponse) GetRemittanceIdOk() (*string, bool)`

GetRemittanceIdOk returns a tuple with the RemittanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemittanceId

`func (o *UnappliedCashResponse) SetRemittanceId(v string)`

SetRemittanceId sets RemittanceId field to given value.


### GetOriginalAmountMicros

`func (o *UnappliedCashResponse) GetOriginalAmountMicros() int32`

GetOriginalAmountMicros returns the OriginalAmountMicros field if non-nil, zero value otherwise.

### GetOriginalAmountMicrosOk

`func (o *UnappliedCashResponse) GetOriginalAmountMicrosOk() (*int32, bool)`

GetOriginalAmountMicrosOk returns a tuple with the OriginalAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmountMicros

`func (o *UnappliedCashResponse) SetOriginalAmountMicros(v int32)`

SetOriginalAmountMicros sets OriginalAmountMicros field to given value.


### GetRemainingAmountMicros

`func (o *UnappliedCashResponse) GetRemainingAmountMicros() int32`

GetRemainingAmountMicros returns the RemainingAmountMicros field if non-nil, zero value otherwise.

### GetRemainingAmountMicrosOk

`func (o *UnappliedCashResponse) GetRemainingAmountMicrosOk() (*int32, bool)`

GetRemainingAmountMicrosOk returns a tuple with the RemainingAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmountMicros

`func (o *UnappliedCashResponse) SetRemainingAmountMicros(v int32)`

SetRemainingAmountMicros sets RemainingAmountMicros field to given value.


### GetReasonCode

`func (o *UnappliedCashResponse) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *UnappliedCashResponse) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *UnappliedCashResponse) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetStatus

`func (o *UnappliedCashResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnappliedCashResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnappliedCashResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDispositionType

`func (o *UnappliedCashResponse) GetDispositionType() string`

GetDispositionType returns the DispositionType field if non-nil, zero value otherwise.

### GetDispositionTypeOk

`func (o *UnappliedCashResponse) GetDispositionTypeOk() (*string, bool)`

GetDispositionTypeOk returns a tuple with the DispositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionType

`func (o *UnappliedCashResponse) SetDispositionType(v string)`

SetDispositionType sets DispositionType field to given value.

### HasDispositionType

`func (o *UnappliedCashResponse) HasDispositionType() bool`

HasDispositionType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UnappliedCashResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnappliedCashResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnappliedCashResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UnappliedCashResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UnappliedCashResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UnappliedCashResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


