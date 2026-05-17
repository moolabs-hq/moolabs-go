# PaymentFailedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**WalletId** | **string** |  | 
**FailureReason** | **string** |  | 
**OutboxEventId** | **string** |  | 

## Methods

### NewPaymentFailedResponse

`func NewPaymentFailedResponse(paymentId string, walletId string, failureReason string, outboxEventId string, ) *PaymentFailedResponse`

NewPaymentFailedResponse instantiates a new PaymentFailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentFailedResponseWithDefaults

`func NewPaymentFailedResponseWithDefaults() *PaymentFailedResponse`

NewPaymentFailedResponseWithDefaults instantiates a new PaymentFailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentFailedResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentFailedResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentFailedResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetWalletId

`func (o *PaymentFailedResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentFailedResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentFailedResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetFailureReason

`func (o *PaymentFailedResponse) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *PaymentFailedResponse) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *PaymentFailedResponse) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetOutboxEventId

`func (o *PaymentFailedResponse) GetOutboxEventId() string`

GetOutboxEventId returns the OutboxEventId field if non-nil, zero value otherwise.

### GetOutboxEventIdOk

`func (o *PaymentFailedResponse) GetOutboxEventIdOk() (*string, bool)`

GetOutboxEventIdOk returns a tuple with the OutboxEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboxEventId

`func (o *PaymentFailedResponse) SetOutboxEventId(v string)`

SetOutboxEventId sets OutboxEventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


