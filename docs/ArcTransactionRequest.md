# ArcTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArcTxnId** | **string** |  | 
**TxnType** | **string** |  | 
**MoometerCustomerId** | **string** |  | 
**MoometerInvoiceId** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | **string** |  | 
**EffectiveAt** | **time.Time** |  | 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewArcTransactionRequest

`func NewArcTransactionRequest(arcTxnId string, txnType string, moometerCustomerId string, amountMicros int32, currencyCode string, effectiveAt time.Time, ) *ArcTransactionRequest`

NewArcTransactionRequest instantiates a new ArcTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcTransactionRequestWithDefaults

`func NewArcTransactionRequestWithDefaults() *ArcTransactionRequest`

NewArcTransactionRequestWithDefaults instantiates a new ArcTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArcTxnId

`func (o *ArcTransactionRequest) GetArcTxnId() string`

GetArcTxnId returns the ArcTxnId field if non-nil, zero value otherwise.

### GetArcTxnIdOk

`func (o *ArcTransactionRequest) GetArcTxnIdOk() (*string, bool)`

GetArcTxnIdOk returns a tuple with the ArcTxnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArcTxnId

`func (o *ArcTransactionRequest) SetArcTxnId(v string)`

SetArcTxnId sets ArcTxnId field to given value.


### GetTxnType

`func (o *ArcTransactionRequest) GetTxnType() string`

GetTxnType returns the TxnType field if non-nil, zero value otherwise.

### GetTxnTypeOk

`func (o *ArcTransactionRequest) GetTxnTypeOk() (*string, bool)`

GetTxnTypeOk returns a tuple with the TxnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnType

`func (o *ArcTransactionRequest) SetTxnType(v string)`

SetTxnType sets TxnType field to given value.


### GetMoometerCustomerId

`func (o *ArcTransactionRequest) GetMoometerCustomerId() string`

GetMoometerCustomerId returns the MoometerCustomerId field if non-nil, zero value otherwise.

### GetMoometerCustomerIdOk

`func (o *ArcTransactionRequest) GetMoometerCustomerIdOk() (*string, bool)`

GetMoometerCustomerIdOk returns a tuple with the MoometerCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoometerCustomerId

`func (o *ArcTransactionRequest) SetMoometerCustomerId(v string)`

SetMoometerCustomerId sets MoometerCustomerId field to given value.


### GetMoometerInvoiceId

`func (o *ArcTransactionRequest) GetMoometerInvoiceId() string`

GetMoometerInvoiceId returns the MoometerInvoiceId field if non-nil, zero value otherwise.

### GetMoometerInvoiceIdOk

`func (o *ArcTransactionRequest) GetMoometerInvoiceIdOk() (*string, bool)`

GetMoometerInvoiceIdOk returns a tuple with the MoometerInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoometerInvoiceId

`func (o *ArcTransactionRequest) SetMoometerInvoiceId(v string)`

SetMoometerInvoiceId sets MoometerInvoiceId field to given value.

### HasMoometerInvoiceId

`func (o *ArcTransactionRequest) HasMoometerInvoiceId() bool`

HasMoometerInvoiceId returns a boolean if a field has been set.

### GetAmountMicros

`func (o *ArcTransactionRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *ArcTransactionRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *ArcTransactionRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *ArcTransactionRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ArcTransactionRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ArcTransactionRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetEffectiveAt

`func (o *ArcTransactionRequest) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *ArcTransactionRequest) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *ArcTransactionRequest) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


### GetPayload

`func (o *ArcTransactionRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ArcTransactionRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ArcTransactionRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ArcTransactionRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


