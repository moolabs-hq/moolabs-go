# PaymentTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of terms to be applied. | 
**Detail** | Pointer to **string** | Text detail of the chosen payment terms. | [optional] [readonly] 
**Notes** | Pointer to **string** | Description of the conditions for payment. | [optional] [readonly] 
**DueAt** | [**[]PaymentDueDate**](PaymentDueDate.md) | When the payment is due. | 

## Methods

### NewPaymentTerms

`func NewPaymentTerms(type_ string, dueAt []PaymentDueDate, ) *PaymentTerms`

NewPaymentTerms instantiates a new PaymentTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermsWithDefaults

`func NewPaymentTermsWithDefaults() *PaymentTerms`

NewPaymentTermsWithDefaults instantiates a new PaymentTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentTerms) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentTerms) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentTerms) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *PaymentTerms) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PaymentTerms) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PaymentTerms) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PaymentTerms) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetNotes

`func (o *PaymentTerms) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PaymentTerms) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PaymentTerms) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PaymentTerms) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDueAt

`func (o *PaymentTerms) GetDueAt() []PaymentDueDate`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *PaymentTerms) GetDueAtOk() (*[]PaymentDueDate, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *PaymentTerms) SetDueAt(v []PaymentDueDate)`

SetDueAt sets DueAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


