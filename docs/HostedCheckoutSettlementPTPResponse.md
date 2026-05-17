# HostedCheckoutSettlementPTPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CaseId** | **string** |  | 
**InvoiceId** | Pointer to **string** |  | [optional] 
**PromisedAmountMicros** | **int32** |  | 
**FulfilledAmountMicros** | **int32** |  | 
**Status** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewHostedCheckoutSettlementPTPResponse

`func NewHostedCheckoutSettlementPTPResponse(id string, caseId string, promisedAmountMicros int32, fulfilledAmountMicros int32, status string, updatedAt time.Time, ) *HostedCheckoutSettlementPTPResponse`

NewHostedCheckoutSettlementPTPResponse instantiates a new HostedCheckoutSettlementPTPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCheckoutSettlementPTPResponseWithDefaults

`func NewHostedCheckoutSettlementPTPResponseWithDefaults() *HostedCheckoutSettlementPTPResponse`

NewHostedCheckoutSettlementPTPResponseWithDefaults instantiates a new HostedCheckoutSettlementPTPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HostedCheckoutSettlementPTPResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostedCheckoutSettlementPTPResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostedCheckoutSettlementPTPResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCaseId

`func (o *HostedCheckoutSettlementPTPResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *HostedCheckoutSettlementPTPResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *HostedCheckoutSettlementPTPResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetInvoiceId

`func (o *HostedCheckoutSettlementPTPResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *HostedCheckoutSettlementPTPResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *HostedCheckoutSettlementPTPResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *HostedCheckoutSettlementPTPResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPromisedAmountMicros

`func (o *HostedCheckoutSettlementPTPResponse) GetPromisedAmountMicros() int32`

GetPromisedAmountMicros returns the PromisedAmountMicros field if non-nil, zero value otherwise.

### GetPromisedAmountMicrosOk

`func (o *HostedCheckoutSettlementPTPResponse) GetPromisedAmountMicrosOk() (*int32, bool)`

GetPromisedAmountMicrosOk returns a tuple with the PromisedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedAmountMicros

`func (o *HostedCheckoutSettlementPTPResponse) SetPromisedAmountMicros(v int32)`

SetPromisedAmountMicros sets PromisedAmountMicros field to given value.


### GetFulfilledAmountMicros

`func (o *HostedCheckoutSettlementPTPResponse) GetFulfilledAmountMicros() int32`

GetFulfilledAmountMicros returns the FulfilledAmountMicros field if non-nil, zero value otherwise.

### GetFulfilledAmountMicrosOk

`func (o *HostedCheckoutSettlementPTPResponse) GetFulfilledAmountMicrosOk() (*int32, bool)`

GetFulfilledAmountMicrosOk returns a tuple with the FulfilledAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAmountMicros

`func (o *HostedCheckoutSettlementPTPResponse) SetFulfilledAmountMicros(v int32)`

SetFulfilledAmountMicros sets FulfilledAmountMicros field to given value.


### GetStatus

`func (o *HostedCheckoutSettlementPTPResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostedCheckoutSettlementPTPResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostedCheckoutSettlementPTPResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *HostedCheckoutSettlementPTPResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HostedCheckoutSettlementPTPResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HostedCheckoutSettlementPTPResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


