# RemittanceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**ExternalReference** | Pointer to **string** |  | [optional] 
**PayerName** | Pointer to **string** |  | [optional] 
**PayerAccountRef** | Pointer to **string** |  | [optional] 
**RawReferenceText** | Pointer to **string** |  | [optional] 
**ReceivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRemittanceCreateRequest

`func NewRemittanceCreateRequest(source string, amountMicros int32, ) *RemittanceCreateRequest`

NewRemittanceCreateRequest instantiates a new RemittanceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemittanceCreateRequestWithDefaults

`func NewRemittanceCreateRequestWithDefaults() *RemittanceCreateRequest`

NewRemittanceCreateRequestWithDefaults instantiates a new RemittanceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RemittanceCreateRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RemittanceCreateRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RemittanceCreateRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetAmountMicros

`func (o *RemittanceCreateRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *RemittanceCreateRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *RemittanceCreateRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *RemittanceCreateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RemittanceCreateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RemittanceCreateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RemittanceCreateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetExternalReference

`func (o *RemittanceCreateRequest) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *RemittanceCreateRequest) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *RemittanceCreateRequest) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *RemittanceCreateRequest) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPayerName

`func (o *RemittanceCreateRequest) GetPayerName() string`

GetPayerName returns the PayerName field if non-nil, zero value otherwise.

### GetPayerNameOk

`func (o *RemittanceCreateRequest) GetPayerNameOk() (*string, bool)`

GetPayerNameOk returns a tuple with the PayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerName

`func (o *RemittanceCreateRequest) SetPayerName(v string)`

SetPayerName sets PayerName field to given value.

### HasPayerName

`func (o *RemittanceCreateRequest) HasPayerName() bool`

HasPayerName returns a boolean if a field has been set.

### GetPayerAccountRef

`func (o *RemittanceCreateRequest) GetPayerAccountRef() string`

GetPayerAccountRef returns the PayerAccountRef field if non-nil, zero value otherwise.

### GetPayerAccountRefOk

`func (o *RemittanceCreateRequest) GetPayerAccountRefOk() (*string, bool)`

GetPayerAccountRefOk returns a tuple with the PayerAccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAccountRef

`func (o *RemittanceCreateRequest) SetPayerAccountRef(v string)`

SetPayerAccountRef sets PayerAccountRef field to given value.

### HasPayerAccountRef

`func (o *RemittanceCreateRequest) HasPayerAccountRef() bool`

HasPayerAccountRef returns a boolean if a field has been set.

### GetRawReferenceText

`func (o *RemittanceCreateRequest) GetRawReferenceText() string`

GetRawReferenceText returns the RawReferenceText field if non-nil, zero value otherwise.

### GetRawReferenceTextOk

`func (o *RemittanceCreateRequest) GetRawReferenceTextOk() (*string, bool)`

GetRawReferenceTextOk returns a tuple with the RawReferenceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawReferenceText

`func (o *RemittanceCreateRequest) SetRawReferenceText(v string)`

SetRawReferenceText sets RawReferenceText field to given value.

### HasRawReferenceText

`func (o *RemittanceCreateRequest) HasRawReferenceText() bool`

HasRawReferenceText returns a boolean if a field has been set.

### GetReceivedAt

`func (o *RemittanceCreateRequest) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *RemittanceCreateRequest) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *RemittanceCreateRequest) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *RemittanceCreateRequest) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


