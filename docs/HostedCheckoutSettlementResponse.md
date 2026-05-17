# HostedCheckoutSettlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | [**PaymentResponse**](PaymentResponse.md) |  | 
**Ptp** | [**HostedCheckoutSettlementPTPResponse**](HostedCheckoutSettlementPTPResponse.md) |  | 

## Methods

### NewHostedCheckoutSettlementResponse

`func NewHostedCheckoutSettlementResponse(payment PaymentResponse, ptp HostedCheckoutSettlementPTPResponse, ) *HostedCheckoutSettlementResponse`

NewHostedCheckoutSettlementResponse instantiates a new HostedCheckoutSettlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCheckoutSettlementResponseWithDefaults

`func NewHostedCheckoutSettlementResponseWithDefaults() *HostedCheckoutSettlementResponse`

NewHostedCheckoutSettlementResponseWithDefaults instantiates a new HostedCheckoutSettlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *HostedCheckoutSettlementResponse) GetPayment() PaymentResponse`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *HostedCheckoutSettlementResponse) GetPaymentOk() (*PaymentResponse, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *HostedCheckoutSettlementResponse) SetPayment(v PaymentResponse)`

SetPayment sets Payment field to given value.


### GetPtp

`func (o *HostedCheckoutSettlementResponse) GetPtp() HostedCheckoutSettlementPTPResponse`

GetPtp returns the Ptp field if non-nil, zero value otherwise.

### GetPtpOk

`func (o *HostedCheckoutSettlementResponse) GetPtpOk() (*HostedCheckoutSettlementPTPResponse, bool)`

GetPtpOk returns a tuple with the Ptp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtp

`func (o *HostedCheckoutSettlementResponse) SetPtp(v HostedCheckoutSettlementPTPResponse)`

SetPtp sets Ptp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


