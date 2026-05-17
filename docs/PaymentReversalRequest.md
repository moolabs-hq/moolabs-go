# PaymentReversalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReversalType** | Pointer to **string** | chargeback, nsf, ach_return, refund, correction | [optional] [default to "correction"]
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentReversalRequest

`func NewPaymentReversalRequest() *PaymentReversalRequest`

NewPaymentReversalRequest instantiates a new PaymentReversalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReversalRequestWithDefaults

`func NewPaymentReversalRequestWithDefaults() *PaymentReversalRequest`

NewPaymentReversalRequestWithDefaults instantiates a new PaymentReversalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReversalType

`func (o *PaymentReversalRequest) GetReversalType() string`

GetReversalType returns the ReversalType field if non-nil, zero value otherwise.

### GetReversalTypeOk

`func (o *PaymentReversalRequest) GetReversalTypeOk() (*string, bool)`

GetReversalTypeOk returns a tuple with the ReversalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversalType

`func (o *PaymentReversalRequest) SetReversalType(v string)`

SetReversalType sets ReversalType field to given value.

### HasReversalType

`func (o *PaymentReversalRequest) HasReversalType() bool`

HasReversalType returns a boolean if a field has been set.

### GetReason

`func (o *PaymentReversalRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PaymentReversalRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PaymentReversalRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PaymentReversalRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


