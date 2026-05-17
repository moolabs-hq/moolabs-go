# CreditMemoApplyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**AmountMicros** | **int32** |  | 

## Methods

### NewCreditMemoApplyRequest

`func NewCreditMemoApplyRequest(invoiceId string, amountMicros int32, ) *CreditMemoApplyRequest`

NewCreditMemoApplyRequest instantiates a new CreditMemoApplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditMemoApplyRequestWithDefaults

`func NewCreditMemoApplyRequestWithDefaults() *CreditMemoApplyRequest`

NewCreditMemoApplyRequestWithDefaults instantiates a new CreditMemoApplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreditMemoApplyRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditMemoApplyRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditMemoApplyRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmountMicros

`func (o *CreditMemoApplyRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *CreditMemoApplyRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *CreditMemoApplyRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


