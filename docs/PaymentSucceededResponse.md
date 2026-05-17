# PaymentSucceededResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**GrantId** | **string** |  | 
**JournalEntryId** | **string** |  | 
**AmountMicros** | **int32** |  | 
**FxRateVersion** | **string** |  | 
**PaidAmountUsdMicros** | **int32** |  | 

## Methods

### NewPaymentSucceededResponse

`func NewPaymentSucceededResponse(paymentId string, grantId string, journalEntryId string, amountMicros int32, fxRateVersion string, paidAmountUsdMicros int32, ) *PaymentSucceededResponse`

NewPaymentSucceededResponse instantiates a new PaymentSucceededResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSucceededResponseWithDefaults

`func NewPaymentSucceededResponseWithDefaults() *PaymentSucceededResponse`

NewPaymentSucceededResponseWithDefaults instantiates a new PaymentSucceededResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentSucceededResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentSucceededResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentSucceededResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetGrantId

`func (o *PaymentSucceededResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *PaymentSucceededResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *PaymentSucceededResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.


### GetJournalEntryId

`func (o *PaymentSucceededResponse) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *PaymentSucceededResponse) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *PaymentSucceededResponse) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.


### GetAmountMicros

`func (o *PaymentSucceededResponse) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PaymentSucceededResponse) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PaymentSucceededResponse) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetFxRateVersion

`func (o *PaymentSucceededResponse) GetFxRateVersion() string`

GetFxRateVersion returns the FxRateVersion field if non-nil, zero value otherwise.

### GetFxRateVersionOk

`func (o *PaymentSucceededResponse) GetFxRateVersionOk() (*string, bool)`

GetFxRateVersionOk returns a tuple with the FxRateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateVersion

`func (o *PaymentSucceededResponse) SetFxRateVersion(v string)`

SetFxRateVersion sets FxRateVersion field to given value.


### GetPaidAmountUsdMicros

`func (o *PaymentSucceededResponse) GetPaidAmountUsdMicros() int32`

GetPaidAmountUsdMicros returns the PaidAmountUsdMicros field if non-nil, zero value otherwise.

### GetPaidAmountUsdMicrosOk

`func (o *PaymentSucceededResponse) GetPaidAmountUsdMicrosOk() (*int32, bool)`

GetPaidAmountUsdMicrosOk returns a tuple with the PaidAmountUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmountUsdMicros

`func (o *PaymentSucceededResponse) SetPaidAmountUsdMicros(v int32)`

SetPaidAmountUsdMicros sets PaidAmountUsdMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


