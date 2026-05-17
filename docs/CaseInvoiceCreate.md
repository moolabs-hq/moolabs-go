# CaseInvoiceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**BaseCurrencyCode** | Pointer to **string** |  | [optional] 
**ExchangeRate** | Pointer to [**NullableExchangeRate**](ExchangeRate.md) |  | [optional] 
**AmountBaseMicros** | Pointer to **int32** |  | [optional] 
**PaidBaseMicros** | Pointer to **int32** |  | [optional] 
**DisputedBaseMicros** | Pointer to **int32** |  | [optional] 
**CreditedBaseMicros** | Pointer to **int32** |  | [optional] 
**WrittenOffBaseMicros** | Pointer to **int32** |  | [optional] 
**DueDate** | **string** |  | 
**IssueDate** | Pointer to **string** |  | [optional] 
**PaidMicros** | Pointer to **int32** |  | [optional] [default to 0]
**DisputedMicros** | Pointer to **int32** |  | [optional] [default to 0]
**CreditedMicros** | Pointer to **int32** |  | [optional] [default to 0]
**WrittenOffMicros** | Pointer to **int32** |  | [optional] [default to 0]
**InvoiceType** | Pointer to **string** |  | [optional] [default to "usage"]

## Methods

### NewCaseInvoiceCreate

`func NewCaseInvoiceCreate(invoiceId string, amountMicros int32, dueDate string, ) *CaseInvoiceCreate`

NewCaseInvoiceCreate instantiates a new CaseInvoiceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseInvoiceCreateWithDefaults

`func NewCaseInvoiceCreateWithDefaults() *CaseInvoiceCreate`

NewCaseInvoiceCreateWithDefaults instantiates a new CaseInvoiceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CaseInvoiceCreate) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CaseInvoiceCreate) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CaseInvoiceCreate) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceNumber

`func (o *CaseInvoiceCreate) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CaseInvoiceCreate) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CaseInvoiceCreate) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *CaseInvoiceCreate) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetAmountMicros

`func (o *CaseInvoiceCreate) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *CaseInvoiceCreate) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *CaseInvoiceCreate) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *CaseInvoiceCreate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CaseInvoiceCreate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CaseInvoiceCreate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CaseInvoiceCreate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBaseCurrencyCode

`func (o *CaseInvoiceCreate) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *CaseInvoiceCreate) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *CaseInvoiceCreate) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.

### HasBaseCurrencyCode

`func (o *CaseInvoiceCreate) HasBaseCurrencyCode() bool`

HasBaseCurrencyCode returns a boolean if a field has been set.

### GetExchangeRate

`func (o *CaseInvoiceCreate) GetExchangeRate() ExchangeRate`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *CaseInvoiceCreate) GetExchangeRateOk() (*ExchangeRate, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *CaseInvoiceCreate) SetExchangeRate(v ExchangeRate)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *CaseInvoiceCreate) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *CaseInvoiceCreate) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *CaseInvoiceCreate) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetAmountBaseMicros

`func (o *CaseInvoiceCreate) GetAmountBaseMicros() int32`

GetAmountBaseMicros returns the AmountBaseMicros field if non-nil, zero value otherwise.

### GetAmountBaseMicrosOk

`func (o *CaseInvoiceCreate) GetAmountBaseMicrosOk() (*int32, bool)`

GetAmountBaseMicrosOk returns a tuple with the AmountBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBaseMicros

`func (o *CaseInvoiceCreate) SetAmountBaseMicros(v int32)`

SetAmountBaseMicros sets AmountBaseMicros field to given value.

### HasAmountBaseMicros

`func (o *CaseInvoiceCreate) HasAmountBaseMicros() bool`

HasAmountBaseMicros returns a boolean if a field has been set.

### GetPaidBaseMicros

`func (o *CaseInvoiceCreate) GetPaidBaseMicros() int32`

GetPaidBaseMicros returns the PaidBaseMicros field if non-nil, zero value otherwise.

### GetPaidBaseMicrosOk

`func (o *CaseInvoiceCreate) GetPaidBaseMicrosOk() (*int32, bool)`

GetPaidBaseMicrosOk returns a tuple with the PaidBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidBaseMicros

`func (o *CaseInvoiceCreate) SetPaidBaseMicros(v int32)`

SetPaidBaseMicros sets PaidBaseMicros field to given value.

### HasPaidBaseMicros

`func (o *CaseInvoiceCreate) HasPaidBaseMicros() bool`

HasPaidBaseMicros returns a boolean if a field has been set.

### GetDisputedBaseMicros

`func (o *CaseInvoiceCreate) GetDisputedBaseMicros() int32`

GetDisputedBaseMicros returns the DisputedBaseMicros field if non-nil, zero value otherwise.

### GetDisputedBaseMicrosOk

`func (o *CaseInvoiceCreate) GetDisputedBaseMicrosOk() (*int32, bool)`

GetDisputedBaseMicrosOk returns a tuple with the DisputedBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedBaseMicros

`func (o *CaseInvoiceCreate) SetDisputedBaseMicros(v int32)`

SetDisputedBaseMicros sets DisputedBaseMicros field to given value.

### HasDisputedBaseMicros

`func (o *CaseInvoiceCreate) HasDisputedBaseMicros() bool`

HasDisputedBaseMicros returns a boolean if a field has been set.

### GetCreditedBaseMicros

`func (o *CaseInvoiceCreate) GetCreditedBaseMicros() int32`

GetCreditedBaseMicros returns the CreditedBaseMicros field if non-nil, zero value otherwise.

### GetCreditedBaseMicrosOk

`func (o *CaseInvoiceCreate) GetCreditedBaseMicrosOk() (*int32, bool)`

GetCreditedBaseMicrosOk returns a tuple with the CreditedBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedBaseMicros

`func (o *CaseInvoiceCreate) SetCreditedBaseMicros(v int32)`

SetCreditedBaseMicros sets CreditedBaseMicros field to given value.

### HasCreditedBaseMicros

`func (o *CaseInvoiceCreate) HasCreditedBaseMicros() bool`

HasCreditedBaseMicros returns a boolean if a field has been set.

### GetWrittenOffBaseMicros

`func (o *CaseInvoiceCreate) GetWrittenOffBaseMicros() int32`

GetWrittenOffBaseMicros returns the WrittenOffBaseMicros field if non-nil, zero value otherwise.

### GetWrittenOffBaseMicrosOk

`func (o *CaseInvoiceCreate) GetWrittenOffBaseMicrosOk() (*int32, bool)`

GetWrittenOffBaseMicrosOk returns a tuple with the WrittenOffBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenOffBaseMicros

`func (o *CaseInvoiceCreate) SetWrittenOffBaseMicros(v int32)`

SetWrittenOffBaseMicros sets WrittenOffBaseMicros field to given value.

### HasWrittenOffBaseMicros

`func (o *CaseInvoiceCreate) HasWrittenOffBaseMicros() bool`

HasWrittenOffBaseMicros returns a boolean if a field has been set.

### GetDueDate

`func (o *CaseInvoiceCreate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CaseInvoiceCreate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CaseInvoiceCreate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetIssueDate

`func (o *CaseInvoiceCreate) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaseInvoiceCreate) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaseInvoiceCreate) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaseInvoiceCreate) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPaidMicros

`func (o *CaseInvoiceCreate) GetPaidMicros() int32`

GetPaidMicros returns the PaidMicros field if non-nil, zero value otherwise.

### GetPaidMicrosOk

`func (o *CaseInvoiceCreate) GetPaidMicrosOk() (*int32, bool)`

GetPaidMicrosOk returns a tuple with the PaidMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMicros

`func (o *CaseInvoiceCreate) SetPaidMicros(v int32)`

SetPaidMicros sets PaidMicros field to given value.

### HasPaidMicros

`func (o *CaseInvoiceCreate) HasPaidMicros() bool`

HasPaidMicros returns a boolean if a field has been set.

### GetDisputedMicros

`func (o *CaseInvoiceCreate) GetDisputedMicros() int32`

GetDisputedMicros returns the DisputedMicros field if non-nil, zero value otherwise.

### GetDisputedMicrosOk

`func (o *CaseInvoiceCreate) GetDisputedMicrosOk() (*int32, bool)`

GetDisputedMicrosOk returns a tuple with the DisputedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedMicros

`func (o *CaseInvoiceCreate) SetDisputedMicros(v int32)`

SetDisputedMicros sets DisputedMicros field to given value.

### HasDisputedMicros

`func (o *CaseInvoiceCreate) HasDisputedMicros() bool`

HasDisputedMicros returns a boolean if a field has been set.

### GetCreditedMicros

`func (o *CaseInvoiceCreate) GetCreditedMicros() int32`

GetCreditedMicros returns the CreditedMicros field if non-nil, zero value otherwise.

### GetCreditedMicrosOk

`func (o *CaseInvoiceCreate) GetCreditedMicrosOk() (*int32, bool)`

GetCreditedMicrosOk returns a tuple with the CreditedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedMicros

`func (o *CaseInvoiceCreate) SetCreditedMicros(v int32)`

SetCreditedMicros sets CreditedMicros field to given value.

### HasCreditedMicros

`func (o *CaseInvoiceCreate) HasCreditedMicros() bool`

HasCreditedMicros returns a boolean if a field has been set.

### GetWrittenOffMicros

`func (o *CaseInvoiceCreate) GetWrittenOffMicros() int32`

GetWrittenOffMicros returns the WrittenOffMicros field if non-nil, zero value otherwise.

### GetWrittenOffMicrosOk

`func (o *CaseInvoiceCreate) GetWrittenOffMicrosOk() (*int32, bool)`

GetWrittenOffMicrosOk returns a tuple with the WrittenOffMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenOffMicros

`func (o *CaseInvoiceCreate) SetWrittenOffMicros(v int32)`

SetWrittenOffMicros sets WrittenOffMicros field to given value.

### HasWrittenOffMicros

`func (o *CaseInvoiceCreate) HasWrittenOffMicros() bool`

HasWrittenOffMicros returns a boolean if a field has been set.

### GetInvoiceType

`func (o *CaseInvoiceCreate) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *CaseInvoiceCreate) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *CaseInvoiceCreate) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *CaseInvoiceCreate) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


