# PaymentInstructionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | **bool** |  | 
**PaymentInstructionsText** | Pointer to **string** |  | [optional] 
**PaymentInstructionsHash** | **string** |  | 
**UpdatedByActorId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentInstructionsResponse

`func NewPaymentInstructionsResponse(ready bool, paymentInstructionsHash string, ) *PaymentInstructionsResponse`

NewPaymentInstructionsResponse instantiates a new PaymentInstructionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstructionsResponseWithDefaults

`func NewPaymentInstructionsResponseWithDefaults() *PaymentInstructionsResponse`

NewPaymentInstructionsResponseWithDefaults instantiates a new PaymentInstructionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *PaymentInstructionsResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *PaymentInstructionsResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *PaymentInstructionsResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetPaymentInstructionsText

`func (o *PaymentInstructionsResponse) GetPaymentInstructionsText() string`

GetPaymentInstructionsText returns the PaymentInstructionsText field if non-nil, zero value otherwise.

### GetPaymentInstructionsTextOk

`func (o *PaymentInstructionsResponse) GetPaymentInstructionsTextOk() (*string, bool)`

GetPaymentInstructionsTextOk returns a tuple with the PaymentInstructionsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructionsText

`func (o *PaymentInstructionsResponse) SetPaymentInstructionsText(v string)`

SetPaymentInstructionsText sets PaymentInstructionsText field to given value.

### HasPaymentInstructionsText

`func (o *PaymentInstructionsResponse) HasPaymentInstructionsText() bool`

HasPaymentInstructionsText returns a boolean if a field has been set.

### GetPaymentInstructionsHash

`func (o *PaymentInstructionsResponse) GetPaymentInstructionsHash() string`

GetPaymentInstructionsHash returns the PaymentInstructionsHash field if non-nil, zero value otherwise.

### GetPaymentInstructionsHashOk

`func (o *PaymentInstructionsResponse) GetPaymentInstructionsHashOk() (*string, bool)`

GetPaymentInstructionsHashOk returns a tuple with the PaymentInstructionsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructionsHash

`func (o *PaymentInstructionsResponse) SetPaymentInstructionsHash(v string)`

SetPaymentInstructionsHash sets PaymentInstructionsHash field to given value.


### GetUpdatedByActorId

`func (o *PaymentInstructionsResponse) GetUpdatedByActorId() string`

GetUpdatedByActorId returns the UpdatedByActorId field if non-nil, zero value otherwise.

### GetUpdatedByActorIdOk

`func (o *PaymentInstructionsResponse) GetUpdatedByActorIdOk() (*string, bool)`

GetUpdatedByActorIdOk returns a tuple with the UpdatedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByActorId

`func (o *PaymentInstructionsResponse) SetUpdatedByActorId(v string)`

SetUpdatedByActorId sets UpdatedByActorId field to given value.

### HasUpdatedByActorId

`func (o *PaymentInstructionsResponse) HasUpdatedByActorId() bool`

HasUpdatedByActorId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentInstructionsResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentInstructionsResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentInstructionsResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentInstructionsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


