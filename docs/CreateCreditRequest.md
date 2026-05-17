# CreateCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**CustomerId** | **string** |  | 
**Amount** | **float32** |  | 
**Reason** | **string** |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**Allocation** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateCreditRequest

`func NewCreateCreditRequest(invoiceId string, customerId string, amount float32, reason string, ) *CreateCreditRequest`

NewCreateCreditRequest instantiates a new CreateCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreditRequestWithDefaults

`func NewCreateCreditRequestWithDefaults() *CreateCreditRequest`

NewCreateCreditRequestWithDefaults instantiates a new CreateCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreateCreditRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateCreditRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateCreditRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetCustomerId

`func (o *CreateCreditRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateCreditRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateCreditRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetAmount

`func (o *CreateCreditRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCreditRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCreditRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *CreateCreditRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCreditRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCreditRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCurrency

`func (o *CreateCreditRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCreditRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCreditRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCreditRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAllocation

`func (o *CreateCreditRequest) GetAllocation() map[string]interface{}`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *CreateCreditRequest) GetAllocationOk() (*map[string]interface{}, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *CreateCreditRequest) SetAllocation(v map[string]interface{})`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *CreateCreditRequest) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


