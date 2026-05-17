# PTPFulfillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**AmountMicros** | **int32** |  | 

## Methods

### NewPTPFulfillRequest

`func NewPTPFulfillRequest(paymentId string, amountMicros int32, ) *PTPFulfillRequest`

NewPTPFulfillRequest instantiates a new PTPFulfillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTPFulfillRequestWithDefaults

`func NewPTPFulfillRequestWithDefaults() *PTPFulfillRequest`

NewPTPFulfillRequestWithDefaults instantiates a new PTPFulfillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PTPFulfillRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PTPFulfillRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PTPFulfillRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetAmountMicros

`func (o *PTPFulfillRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PTPFulfillRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PTPFulfillRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


