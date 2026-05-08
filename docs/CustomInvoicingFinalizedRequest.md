# CustomInvoicingFinalizedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoicing** | Pointer to [**CustomInvoicingFinalizedInvoicingRequest**](CustomInvoicingFinalizedInvoicingRequest.md) | The result of the synchronization. | [optional] 
**Payment** | Pointer to [**CustomInvoicingFinalizedPaymentRequest**](CustomInvoicingFinalizedPaymentRequest.md) | The result of the payment synchronization. | [optional] 

## Methods

### NewCustomInvoicingFinalizedRequest

`func NewCustomInvoicingFinalizedRequest() *CustomInvoicingFinalizedRequest`

NewCustomInvoicingFinalizedRequest instantiates a new CustomInvoicingFinalizedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingFinalizedRequestWithDefaults

`func NewCustomInvoicingFinalizedRequestWithDefaults() *CustomInvoicingFinalizedRequest`

NewCustomInvoicingFinalizedRequestWithDefaults instantiates a new CustomInvoicingFinalizedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoicing

`func (o *CustomInvoicingFinalizedRequest) GetInvoicing() CustomInvoicingFinalizedInvoicingRequest`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *CustomInvoicingFinalizedRequest) GetInvoicingOk() (*CustomInvoicingFinalizedInvoicingRequest, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *CustomInvoicingFinalizedRequest) SetInvoicing(v CustomInvoicingFinalizedInvoicingRequest)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *CustomInvoicingFinalizedRequest) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.

### GetPayment

`func (o *CustomInvoicingFinalizedRequest) GetPayment() CustomInvoicingFinalizedPaymentRequest`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *CustomInvoicingFinalizedRequest) GetPaymentOk() (*CustomInvoicingFinalizedPaymentRequest, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *CustomInvoicingFinalizedRequest) SetPayment(v CustomInvoicingFinalizedPaymentRequest)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *CustomInvoicingFinalizedRequest) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


