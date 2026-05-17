# InvoiceUpsertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetsuiteInvoiceId** | **string** |  | 
**MoometerInvoiceId** | **string** |  | 
**NetsuiteCustomerId** | **string** |  | 
**MoometerCustomerId** | **string** |  | 
**InvoiceNumber** | **string** |  | 
**InvoiceDate** | **string** |  | 
**DueDate** | **string** |  | 
**CurrencyCode** | **string** |  | 
**AmountMicros** | **int32** |  | 
**PaidMicros** | Pointer to **int32** |  | [optional] [default to 0]
**OpenMicros** | **int32** |  | 
**Status** | Pointer to **string** |  | [optional] [default to "open"]
**SourceUpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvoiceUpsertRequest

`func NewInvoiceUpsertRequest(netsuiteInvoiceId string, moometerInvoiceId string, netsuiteCustomerId string, moometerCustomerId string, invoiceNumber string, invoiceDate string, dueDate string, currencyCode string, amountMicros int32, openMicros int32, ) *InvoiceUpsertRequest`

NewInvoiceUpsertRequest instantiates a new InvoiceUpsertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpsertRequestWithDefaults

`func NewInvoiceUpsertRequestWithDefaults() *InvoiceUpsertRequest`

NewInvoiceUpsertRequestWithDefaults instantiates a new InvoiceUpsertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetsuiteInvoiceId

`func (o *InvoiceUpsertRequest) GetNetsuiteInvoiceId() string`

GetNetsuiteInvoiceId returns the NetsuiteInvoiceId field if non-nil, zero value otherwise.

### GetNetsuiteInvoiceIdOk

`func (o *InvoiceUpsertRequest) GetNetsuiteInvoiceIdOk() (*string, bool)`

GetNetsuiteInvoiceIdOk returns a tuple with the NetsuiteInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteInvoiceId

`func (o *InvoiceUpsertRequest) SetNetsuiteInvoiceId(v string)`

SetNetsuiteInvoiceId sets NetsuiteInvoiceId field to given value.


### GetMoometerInvoiceId

`func (o *InvoiceUpsertRequest) GetMoometerInvoiceId() string`

GetMoometerInvoiceId returns the MoometerInvoiceId field if non-nil, zero value otherwise.

### GetMoometerInvoiceIdOk

`func (o *InvoiceUpsertRequest) GetMoometerInvoiceIdOk() (*string, bool)`

GetMoometerInvoiceIdOk returns a tuple with the MoometerInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoometerInvoiceId

`func (o *InvoiceUpsertRequest) SetMoometerInvoiceId(v string)`

SetMoometerInvoiceId sets MoometerInvoiceId field to given value.


### GetNetsuiteCustomerId

`func (o *InvoiceUpsertRequest) GetNetsuiteCustomerId() string`

GetNetsuiteCustomerId returns the NetsuiteCustomerId field if non-nil, zero value otherwise.

### GetNetsuiteCustomerIdOk

`func (o *InvoiceUpsertRequest) GetNetsuiteCustomerIdOk() (*string, bool)`

GetNetsuiteCustomerIdOk returns a tuple with the NetsuiteCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteCustomerId

`func (o *InvoiceUpsertRequest) SetNetsuiteCustomerId(v string)`

SetNetsuiteCustomerId sets NetsuiteCustomerId field to given value.


### GetMoometerCustomerId

`func (o *InvoiceUpsertRequest) GetMoometerCustomerId() string`

GetMoometerCustomerId returns the MoometerCustomerId field if non-nil, zero value otherwise.

### GetMoometerCustomerIdOk

`func (o *InvoiceUpsertRequest) GetMoometerCustomerIdOk() (*string, bool)`

GetMoometerCustomerIdOk returns a tuple with the MoometerCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoometerCustomerId

`func (o *InvoiceUpsertRequest) SetMoometerCustomerId(v string)`

SetMoometerCustomerId sets MoometerCustomerId field to given value.


### GetInvoiceNumber

`func (o *InvoiceUpsertRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceUpsertRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceUpsertRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetInvoiceDate

`func (o *InvoiceUpsertRequest) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoiceUpsertRequest) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoiceUpsertRequest) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetDueDate

`func (o *InvoiceUpsertRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceUpsertRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceUpsertRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetCurrencyCode

`func (o *InvoiceUpsertRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceUpsertRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceUpsertRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetAmountMicros

`func (o *InvoiceUpsertRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *InvoiceUpsertRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *InvoiceUpsertRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetPaidMicros

`func (o *InvoiceUpsertRequest) GetPaidMicros() int32`

GetPaidMicros returns the PaidMicros field if non-nil, zero value otherwise.

### GetPaidMicrosOk

`func (o *InvoiceUpsertRequest) GetPaidMicrosOk() (*int32, bool)`

GetPaidMicrosOk returns a tuple with the PaidMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMicros

`func (o *InvoiceUpsertRequest) SetPaidMicros(v int32)`

SetPaidMicros sets PaidMicros field to given value.

### HasPaidMicros

`func (o *InvoiceUpsertRequest) HasPaidMicros() bool`

HasPaidMicros returns a boolean if a field has been set.

### GetOpenMicros

`func (o *InvoiceUpsertRequest) GetOpenMicros() int32`

GetOpenMicros returns the OpenMicros field if non-nil, zero value otherwise.

### GetOpenMicrosOk

`func (o *InvoiceUpsertRequest) GetOpenMicrosOk() (*int32, bool)`

GetOpenMicrosOk returns a tuple with the OpenMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMicros

`func (o *InvoiceUpsertRequest) SetOpenMicros(v int32)`

SetOpenMicros sets OpenMicros field to given value.


### GetStatus

`func (o *InvoiceUpsertRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceUpsertRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceUpsertRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceUpsertRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSourceUpdatedAt

`func (o *InvoiceUpsertRequest) GetSourceUpdatedAt() time.Time`

GetSourceUpdatedAt returns the SourceUpdatedAt field if non-nil, zero value otherwise.

### GetSourceUpdatedAtOk

`func (o *InvoiceUpsertRequest) GetSourceUpdatedAtOk() (*time.Time, bool)`

GetSourceUpdatedAtOk returns a tuple with the SourceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUpdatedAt

`func (o *InvoiceUpsertRequest) SetSourceUpdatedAt(v time.Time)`

SetSourceUpdatedAt sets SourceUpdatedAt field to given value.

### HasSourceUpdatedAt

`func (o *InvoiceUpsertRequest) HasSourceUpdatedAt() bool`

HasSourceUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


