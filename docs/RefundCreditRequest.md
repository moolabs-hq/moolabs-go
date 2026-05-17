# RefundCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMicros** | **int32** |  | 
**RefundReference** | Pointer to **string** |  | [optional] 

## Methods

### NewRefundCreditRequest

`func NewRefundCreditRequest(amountMicros int32, ) *RefundCreditRequest`

NewRefundCreditRequest instantiates a new RefundCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundCreditRequestWithDefaults

`func NewRefundCreditRequestWithDefaults() *RefundCreditRequest`

NewRefundCreditRequestWithDefaults instantiates a new RefundCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountMicros

`func (o *RefundCreditRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *RefundCreditRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *RefundCreditRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetRefundReference

`func (o *RefundCreditRequest) GetRefundReference() string`

GetRefundReference returns the RefundReference field if non-nil, zero value otherwise.

### GetRefundReferenceOk

`func (o *RefundCreditRequest) GetRefundReferenceOk() (*string, bool)`

GetRefundReferenceOk returns a tuple with the RefundReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundReference

`func (o *RefundCreditRequest) SetRefundReference(v string)`

SetRefundReference sets RefundReference field to given value.

### HasRefundReference

`func (o *RefundCreditRequest) HasRefundReference() bool`

HasRefundReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


