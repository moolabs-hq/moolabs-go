# OpenFromScopeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CustomerId** | **string** |  | 
**BillingAccountId** | Pointer to **string** |  | [optional] 
**BillToEntityId** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**InvoiceIds** | Pointer to **[]string** | Invoice IDs to attach | [optional] 

## Methods

### NewOpenFromScopeRequest

`func NewOpenFromScopeRequest(accountId string, customerId string, ) *OpenFromScopeRequest`

NewOpenFromScopeRequest instantiates a new OpenFromScopeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenFromScopeRequestWithDefaults

`func NewOpenFromScopeRequestWithDefaults() *OpenFromScopeRequest`

NewOpenFromScopeRequestWithDefaults instantiates a new OpenFromScopeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *OpenFromScopeRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OpenFromScopeRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OpenFromScopeRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *OpenFromScopeRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OpenFromScopeRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OpenFromScopeRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBillingAccountId

`func (o *OpenFromScopeRequest) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *OpenFromScopeRequest) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *OpenFromScopeRequest) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *OpenFromScopeRequest) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetBillToEntityId

`func (o *OpenFromScopeRequest) GetBillToEntityId() string`

GetBillToEntityId returns the BillToEntityId field if non-nil, zero value otherwise.

### GetBillToEntityIdOk

`func (o *OpenFromScopeRequest) GetBillToEntityIdOk() (*string, bool)`

GetBillToEntityIdOk returns a tuple with the BillToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToEntityId

`func (o *OpenFromScopeRequest) SetBillToEntityId(v string)`

SetBillToEntityId sets BillToEntityId field to given value.

### HasBillToEntityId

`func (o *OpenFromScopeRequest) HasBillToEntityId() bool`

HasBillToEntityId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OpenFromScopeRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OpenFromScopeRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OpenFromScopeRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OpenFromScopeRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetInvoiceIds

`func (o *OpenFromScopeRequest) GetInvoiceIds() []string`

GetInvoiceIds returns the InvoiceIds field if non-nil, zero value otherwise.

### GetInvoiceIdsOk

`func (o *OpenFromScopeRequest) GetInvoiceIdsOk() (*[]string, bool)`

GetInvoiceIdsOk returns a tuple with the InvoiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceIds

`func (o *OpenFromScopeRequest) SetInvoiceIds(v []string)`

SetInvoiceIds sets InvoiceIds field to given value.

### HasInvoiceIds

`func (o *OpenFromScopeRequest) HasInvoiceIds() bool`

HasInvoiceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


