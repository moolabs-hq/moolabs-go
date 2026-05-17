# PaymentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **string** |  | 
**RemittanceId** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**BaseCurrencyCode** | Pointer to **string** |  | [optional] 
**AmountBaseMicros** | Pointer to **int32** |  | [optional] 
**PaymentMethod** | **string** | stripe, ach, wire, check, credit_adjustment, other | 
**ExternalPaymentId** | Pointer to **string** |  | [optional] 
**ReceivedAt** | **time.Time** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPaymentCreate

`func NewPaymentCreate(caseId string, amountMicros int32, paymentMethod string, receivedAt time.Time, ) *PaymentCreate`

NewPaymentCreate instantiates a new PaymentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCreateWithDefaults

`func NewPaymentCreateWithDefaults() *PaymentCreate`

NewPaymentCreateWithDefaults instantiates a new PaymentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *PaymentCreate) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PaymentCreate) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PaymentCreate) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetRemittanceId

`func (o *PaymentCreate) GetRemittanceId() string`

GetRemittanceId returns the RemittanceId field if non-nil, zero value otherwise.

### GetRemittanceIdOk

`func (o *PaymentCreate) GetRemittanceIdOk() (*string, bool)`

GetRemittanceIdOk returns a tuple with the RemittanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemittanceId

`func (o *PaymentCreate) SetRemittanceId(v string)`

SetRemittanceId sets RemittanceId field to given value.

### HasRemittanceId

`func (o *PaymentCreate) HasRemittanceId() bool`

HasRemittanceId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PaymentCreate) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentCreate) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentCreate) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaymentCreate) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetAmountMicros

`func (o *PaymentCreate) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PaymentCreate) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PaymentCreate) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *PaymentCreate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentCreate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentCreate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PaymentCreate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBaseCurrencyCode

`func (o *PaymentCreate) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *PaymentCreate) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *PaymentCreate) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.

### HasBaseCurrencyCode

`func (o *PaymentCreate) HasBaseCurrencyCode() bool`

HasBaseCurrencyCode returns a boolean if a field has been set.

### GetAmountBaseMicros

`func (o *PaymentCreate) GetAmountBaseMicros() int32`

GetAmountBaseMicros returns the AmountBaseMicros field if non-nil, zero value otherwise.

### GetAmountBaseMicrosOk

`func (o *PaymentCreate) GetAmountBaseMicrosOk() (*int32, bool)`

GetAmountBaseMicrosOk returns a tuple with the AmountBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBaseMicros

`func (o *PaymentCreate) SetAmountBaseMicros(v int32)`

SetAmountBaseMicros sets AmountBaseMicros field to given value.

### HasAmountBaseMicros

`func (o *PaymentCreate) HasAmountBaseMicros() bool`

HasAmountBaseMicros returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentCreate) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentCreate) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentCreate) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetExternalPaymentId

`func (o *PaymentCreate) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *PaymentCreate) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *PaymentCreate) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *PaymentCreate) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetReceivedAt

`func (o *PaymentCreate) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *PaymentCreate) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *PaymentCreate) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetMetadata

`func (o *PaymentCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


