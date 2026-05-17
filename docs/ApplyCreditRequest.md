# ApplyCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**AmountMicros** | **int32** |  | 
**CaseId** | Pointer to **string** | Case ID to disambiguate when invoice_id exists in multiple cases | [optional] 

## Methods

### NewApplyCreditRequest

`func NewApplyCreditRequest(invoiceId string, amountMicros int32, ) *ApplyCreditRequest`

NewApplyCreditRequest instantiates a new ApplyCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyCreditRequestWithDefaults

`func NewApplyCreditRequestWithDefaults() *ApplyCreditRequest`

NewApplyCreditRequestWithDefaults instantiates a new ApplyCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *ApplyCreditRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ApplyCreditRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ApplyCreditRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmountMicros

`func (o *ApplyCreditRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *ApplyCreditRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *ApplyCreditRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCaseId

`func (o *ApplyCreditRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ApplyCreditRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ApplyCreditRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *ApplyCreditRequest) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


