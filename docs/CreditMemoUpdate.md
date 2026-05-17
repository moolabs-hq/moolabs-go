# CreditMemoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**CreditAmountMicros** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreditMemoUpdate

`func NewCreditMemoUpdate() *CreditMemoUpdate`

NewCreditMemoUpdate instantiates a new CreditMemoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditMemoUpdateWithDefaults

`func NewCreditMemoUpdateWithDefaults() *CreditMemoUpdate`

NewCreditMemoUpdateWithDefaults instantiates a new CreditMemoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceNumber

`func (o *CreditMemoUpdate) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *CreditMemoUpdate) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *CreditMemoUpdate) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *CreditMemoUpdate) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetReason

`func (o *CreditMemoUpdate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditMemoUpdate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditMemoUpdate) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreditMemoUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCreditAmountMicros

`func (o *CreditMemoUpdate) GetCreditAmountMicros() int32`

GetCreditAmountMicros returns the CreditAmountMicros field if non-nil, zero value otherwise.

### GetCreditAmountMicrosOk

`func (o *CreditMemoUpdate) GetCreditAmountMicrosOk() (*int32, bool)`

GetCreditAmountMicrosOk returns a tuple with the CreditAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountMicros

`func (o *CreditMemoUpdate) SetCreditAmountMicros(v int32)`

SetCreditAmountMicros sets CreditAmountMicros field to given value.

### HasCreditAmountMicros

`func (o *CreditMemoUpdate) HasCreditAmountMicros() bool`

HasCreditAmountMicros returns a boolean if a field has been set.

### GetDescription

`func (o *CreditMemoUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditMemoUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditMemoUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditMemoUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


