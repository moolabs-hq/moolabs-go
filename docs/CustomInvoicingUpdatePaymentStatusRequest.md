# CustomInvoicingUpdatePaymentStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**CustomInvoicingPaymentTrigger**](CustomInvoicingPaymentTrigger.md) | The trigger to be executed on the invoice. | 
**ExternalPaymentId** | Pointer to **string** | If set the invoice&#39;s payment external ID will be set to this value. | [optional] 
**PaymentProvider** | Pointer to **string** | If set the invoice&#39;s payment provider metadata will be set to this value. | [optional] 

## Methods

### NewCustomInvoicingUpdatePaymentStatusRequest

`func NewCustomInvoicingUpdatePaymentStatusRequest(trigger CustomInvoicingPaymentTrigger, ) *CustomInvoicingUpdatePaymentStatusRequest`

NewCustomInvoicingUpdatePaymentStatusRequest instantiates a new CustomInvoicingUpdatePaymentStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingUpdatePaymentStatusRequestWithDefaults

`func NewCustomInvoicingUpdatePaymentStatusRequestWithDefaults() *CustomInvoicingUpdatePaymentStatusRequest`

NewCustomInvoicingUpdatePaymentStatusRequestWithDefaults instantiates a new CustomInvoicingUpdatePaymentStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *CustomInvoicingUpdatePaymentStatusRequest) GetTrigger() CustomInvoicingPaymentTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CustomInvoicingUpdatePaymentStatusRequest) GetTriggerOk() (*CustomInvoicingPaymentTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CustomInvoicingUpdatePaymentStatusRequest) SetTrigger(v CustomInvoicingPaymentTrigger)`

SetTrigger sets Trigger field to given value.


### GetExternalPaymentId

`func (o *CustomInvoicingUpdatePaymentStatusRequest) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *CustomInvoicingUpdatePaymentStatusRequest) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *CustomInvoicingUpdatePaymentStatusRequest) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *CustomInvoicingUpdatePaymentStatusRequest) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetPaymentProvider

`func (o *CustomInvoicingUpdatePaymentStatusRequest) GetPaymentProvider() string`

GetPaymentProvider returns the PaymentProvider field if non-nil, zero value otherwise.

### GetPaymentProviderOk

`func (o *CustomInvoicingUpdatePaymentStatusRequest) GetPaymentProviderOk() (*string, bool)`

GetPaymentProviderOk returns a tuple with the PaymentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProvider

`func (o *CustomInvoicingUpdatePaymentStatusRequest) SetPaymentProvider(v string)`

SetPaymentProvider sets PaymentProvider field to given value.

### HasPaymentProvider

`func (o *CustomInvoicingUpdatePaymentStatusRequest) HasPaymentProvider() bool`

HasPaymentProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


