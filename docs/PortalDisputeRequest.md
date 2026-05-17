# PortalDisputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**DisputedAmountMicros** | **int32** |  | 
**Description** | **string** |  | 
**CaseId** | Pointer to **string** | Case ID to disambiguate when invoice_id exists in multiple cases | [optional] 

## Methods

### NewPortalDisputeRequest

`func NewPortalDisputeRequest(invoiceId string, disputedAmountMicros int32, description string, ) *PortalDisputeRequest`

NewPortalDisputeRequest instantiates a new PortalDisputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalDisputeRequestWithDefaults

`func NewPortalDisputeRequestWithDefaults() *PortalDisputeRequest`

NewPortalDisputeRequestWithDefaults instantiates a new PortalDisputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *PortalDisputeRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PortalDisputeRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PortalDisputeRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetDisputedAmountMicros

`func (o *PortalDisputeRequest) GetDisputedAmountMicros() int32`

GetDisputedAmountMicros returns the DisputedAmountMicros field if non-nil, zero value otherwise.

### GetDisputedAmountMicrosOk

`func (o *PortalDisputeRequest) GetDisputedAmountMicrosOk() (*int32, bool)`

GetDisputedAmountMicrosOk returns a tuple with the DisputedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmountMicros

`func (o *PortalDisputeRequest) SetDisputedAmountMicros(v int32)`

SetDisputedAmountMicros sets DisputedAmountMicros field to given value.


### GetDescription

`func (o *PortalDisputeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortalDisputeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortalDisputeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCaseId

`func (o *PortalDisputeRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PortalDisputeRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PortalDisputeRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *PortalDisputeRequest) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


