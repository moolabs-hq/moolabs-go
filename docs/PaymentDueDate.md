# PaymentDueDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueAt** | **time.Time** | When the payment is due. | [readonly] 
**Notes** | Pointer to **string** | Other details to take into account for the due date. | [optional] [readonly] 
**Amount** | **string** | How much needs to be paid by the date. | [readonly] 
**Percent** | Pointer to **float64** | Percentage of the total that should be paid by the date. | [optional] [readonly] 
**Currency** | Pointer to **string** | If different from the parent document&#39;s base currency. | [optional] [readonly] 

## Methods

### NewPaymentDueDate

`func NewPaymentDueDate(dueAt time.Time, amount string, ) *PaymentDueDate`

NewPaymentDueDate instantiates a new PaymentDueDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDueDateWithDefaults

`func NewPaymentDueDateWithDefaults() *PaymentDueDate`

NewPaymentDueDateWithDefaults instantiates a new PaymentDueDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueAt

`func (o *PaymentDueDate) GetDueAt() time.Time`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *PaymentDueDate) GetDueAtOk() (*time.Time, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *PaymentDueDate) SetDueAt(v time.Time)`

SetDueAt sets DueAt field to given value.


### GetNotes

`func (o *PaymentDueDate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PaymentDueDate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PaymentDueDate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PaymentDueDate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentDueDate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentDueDate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentDueDate) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetPercent

`func (o *PaymentDueDate) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *PaymentDueDate) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *PaymentDueDate) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *PaymentDueDate) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentDueDate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentDueDate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentDueDate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentDueDate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


