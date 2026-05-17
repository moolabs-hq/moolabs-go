# PTPCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromisedAmountMicros** | **int32** |  | 
**PromisedDate** | **string** |  | 
**CapturedFromChannel** | **string** |  | 
**CapturedBy** | **string** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**ContactId** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 

## Methods

### NewPTPCreateRequest

`func NewPTPCreateRequest(promisedAmountMicros int32, promisedDate string, capturedFromChannel string, capturedBy string, ) *PTPCreateRequest`

NewPTPCreateRequest instantiates a new PTPCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTPCreateRequestWithDefaults

`func NewPTPCreateRequestWithDefaults() *PTPCreateRequest`

NewPTPCreateRequestWithDefaults instantiates a new PTPCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromisedAmountMicros

`func (o *PTPCreateRequest) GetPromisedAmountMicros() int32`

GetPromisedAmountMicros returns the PromisedAmountMicros field if non-nil, zero value otherwise.

### GetPromisedAmountMicrosOk

`func (o *PTPCreateRequest) GetPromisedAmountMicrosOk() (*int32, bool)`

GetPromisedAmountMicrosOk returns a tuple with the PromisedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedAmountMicros

`func (o *PTPCreateRequest) SetPromisedAmountMicros(v int32)`

SetPromisedAmountMicros sets PromisedAmountMicros field to given value.


### GetPromisedDate

`func (o *PTPCreateRequest) GetPromisedDate() string`

GetPromisedDate returns the PromisedDate field if non-nil, zero value otherwise.

### GetPromisedDateOk

`func (o *PTPCreateRequest) GetPromisedDateOk() (*string, bool)`

GetPromisedDateOk returns a tuple with the PromisedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedDate

`func (o *PTPCreateRequest) SetPromisedDate(v string)`

SetPromisedDate sets PromisedDate field to given value.


### GetCapturedFromChannel

`func (o *PTPCreateRequest) GetCapturedFromChannel() string`

GetCapturedFromChannel returns the CapturedFromChannel field if non-nil, zero value otherwise.

### GetCapturedFromChannelOk

`func (o *PTPCreateRequest) GetCapturedFromChannelOk() (*string, bool)`

GetCapturedFromChannelOk returns a tuple with the CapturedFromChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedFromChannel

`func (o *PTPCreateRequest) SetCapturedFromChannel(v string)`

SetCapturedFromChannel sets CapturedFromChannel field to given value.


### GetCapturedBy

`func (o *PTPCreateRequest) GetCapturedBy() string`

GetCapturedBy returns the CapturedBy field if non-nil, zero value otherwise.

### GetCapturedByOk

`func (o *PTPCreateRequest) GetCapturedByOk() (*string, bool)`

GetCapturedByOk returns a tuple with the CapturedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedBy

`func (o *PTPCreateRequest) SetCapturedBy(v string)`

SetCapturedBy sets CapturedBy field to given value.


### GetCurrencyCode

`func (o *PTPCreateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PTPCreateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PTPCreateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PTPCreateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetContactId

`func (o *PTPCreateRequest) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *PTPCreateRequest) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *PTPCreateRequest) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *PTPCreateRequest) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PTPCreateRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PTPCreateRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PTPCreateRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PTPCreateRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetConfidence

`func (o *PTPCreateRequest) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PTPCreateRequest) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PTPCreateRequest) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PTPCreateRequest) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


