# PaymentTermInstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of terms to be applied. | 
**Detail** | Pointer to **string** | Text detail of the chosen payment terms. | [optional] [readonly] 
**Notes** | Pointer to **string** | Description of the conditions for payment. | [optional] [readonly] 

## Methods

### NewPaymentTermInstant

`func NewPaymentTermInstant(type_ string, ) *PaymentTermInstant`

NewPaymentTermInstant instantiates a new PaymentTermInstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermInstantWithDefaults

`func NewPaymentTermInstantWithDefaults() *PaymentTermInstant`

NewPaymentTermInstantWithDefaults instantiates a new PaymentTermInstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentTermInstant) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentTermInstant) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentTermInstant) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *PaymentTermInstant) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PaymentTermInstant) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PaymentTermInstant) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PaymentTermInstant) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetNotes

`func (o *PaymentTermInstant) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PaymentTermInstant) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PaymentTermInstant) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PaymentTermInstant) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


