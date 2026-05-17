# PortalPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **string** |  | 
**AmountMicros** | **int32** |  | 
**PaymentMethod** | Pointer to **string** |  | [optional] [default to "stripe"]

## Methods

### NewPortalPaymentRequest

`func NewPortalPaymentRequest(caseId string, amountMicros int32, ) *PortalPaymentRequest`

NewPortalPaymentRequest instantiates a new PortalPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalPaymentRequestWithDefaults

`func NewPortalPaymentRequestWithDefaults() *PortalPaymentRequest`

NewPortalPaymentRequestWithDefaults instantiates a new PortalPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *PortalPaymentRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PortalPaymentRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PortalPaymentRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetAmountMicros

`func (o *PortalPaymentRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PortalPaymentRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PortalPaymentRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetPaymentMethod

`func (o *PortalPaymentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PortalPaymentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PortalPaymentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PortalPaymentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


