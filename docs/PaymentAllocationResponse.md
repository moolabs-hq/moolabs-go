# PaymentAllocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PaymentId** | **string** |  | 
**InvoiceId** | **string** |  | 
**AppliedAmountMicros** | **int32** |  | 
**InvoiceCurrencyCode** | **string** |  | 
**SettlementCurrencyCode** | **string** |  | 
**ApplicationSource** | Pointer to **string** |  | [optional] 
**ExternalApplicationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPaymentAllocationResponse

`func NewPaymentAllocationResponse(id string, paymentId string, invoiceId string, appliedAmountMicros int32, invoiceCurrencyCode string, settlementCurrencyCode string, createdAt time.Time, ) *PaymentAllocationResponse`

NewPaymentAllocationResponse instantiates a new PaymentAllocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAllocationResponseWithDefaults

`func NewPaymentAllocationResponseWithDefaults() *PaymentAllocationResponse`

NewPaymentAllocationResponseWithDefaults instantiates a new PaymentAllocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentAllocationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentAllocationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentAllocationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentId

`func (o *PaymentAllocationResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentAllocationResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentAllocationResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetInvoiceId

`func (o *PaymentAllocationResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentAllocationResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentAllocationResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAppliedAmountMicros

`func (o *PaymentAllocationResponse) GetAppliedAmountMicros() int32`

GetAppliedAmountMicros returns the AppliedAmountMicros field if non-nil, zero value otherwise.

### GetAppliedAmountMicrosOk

`func (o *PaymentAllocationResponse) GetAppliedAmountMicrosOk() (*int32, bool)`

GetAppliedAmountMicrosOk returns a tuple with the AppliedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmountMicros

`func (o *PaymentAllocationResponse) SetAppliedAmountMicros(v int32)`

SetAppliedAmountMicros sets AppliedAmountMicros field to given value.


### GetInvoiceCurrencyCode

`func (o *PaymentAllocationResponse) GetInvoiceCurrencyCode() string`

GetInvoiceCurrencyCode returns the InvoiceCurrencyCode field if non-nil, zero value otherwise.

### GetInvoiceCurrencyCodeOk

`func (o *PaymentAllocationResponse) GetInvoiceCurrencyCodeOk() (*string, bool)`

GetInvoiceCurrencyCodeOk returns a tuple with the InvoiceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCurrencyCode

`func (o *PaymentAllocationResponse) SetInvoiceCurrencyCode(v string)`

SetInvoiceCurrencyCode sets InvoiceCurrencyCode field to given value.


### GetSettlementCurrencyCode

`func (o *PaymentAllocationResponse) GetSettlementCurrencyCode() string`

GetSettlementCurrencyCode returns the SettlementCurrencyCode field if non-nil, zero value otherwise.

### GetSettlementCurrencyCodeOk

`func (o *PaymentAllocationResponse) GetSettlementCurrencyCodeOk() (*string, bool)`

GetSettlementCurrencyCodeOk returns a tuple with the SettlementCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrencyCode

`func (o *PaymentAllocationResponse) SetSettlementCurrencyCode(v string)`

SetSettlementCurrencyCode sets SettlementCurrencyCode field to given value.


### GetApplicationSource

`func (o *PaymentAllocationResponse) GetApplicationSource() string`

GetApplicationSource returns the ApplicationSource field if non-nil, zero value otherwise.

### GetApplicationSourceOk

`func (o *PaymentAllocationResponse) GetApplicationSourceOk() (*string, bool)`

GetApplicationSourceOk returns a tuple with the ApplicationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSource

`func (o *PaymentAllocationResponse) SetApplicationSource(v string)`

SetApplicationSource sets ApplicationSource field to given value.

### HasApplicationSource

`func (o *PaymentAllocationResponse) HasApplicationSource() bool`

HasApplicationSource returns a boolean if a field has been set.

### GetExternalApplicationId

`func (o *PaymentAllocationResponse) GetExternalApplicationId() string`

GetExternalApplicationId returns the ExternalApplicationId field if non-nil, zero value otherwise.

### GetExternalApplicationIdOk

`func (o *PaymentAllocationResponse) GetExternalApplicationIdOk() (*string, bool)`

GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApplicationId

`func (o *PaymentAllocationResponse) SetExternalApplicationId(v string)`

SetExternalApplicationId sets ExternalApplicationId field to given value.

### HasExternalApplicationId

`func (o *PaymentAllocationResponse) HasExternalApplicationId() bool`

HasExternalApplicationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentAllocationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentAllocationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentAllocationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


