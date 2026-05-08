# CreditNoteOriginalInvoiceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the invoice. | 
**IssuedAt** | Pointer to **time.Time** | IssueAt reflects the time the document was issued. | [optional] [readonly] 
**Number** | Pointer to **string** | (Serial) Number of the referenced document. | [optional] [readonly] 
**Url** | **string** | Link to the source document. | [readonly] 
**Reason** | Pointer to **string** | Human readable description on why this reference is here or needs to be used. | [optional] [readonly] 
**Description** | Pointer to **string** | Additional details about the document. | [optional] [readonly] 

## Methods

### NewCreditNoteOriginalInvoiceRef

`func NewCreditNoteOriginalInvoiceRef(type_ string, url string, ) *CreditNoteOriginalInvoiceRef`

NewCreditNoteOriginalInvoiceRef instantiates a new CreditNoteOriginalInvoiceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteOriginalInvoiceRefWithDefaults

`func NewCreditNoteOriginalInvoiceRefWithDefaults() *CreditNoteOriginalInvoiceRef`

NewCreditNoteOriginalInvoiceRefWithDefaults instantiates a new CreditNoteOriginalInvoiceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreditNoteOriginalInvoiceRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditNoteOriginalInvoiceRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditNoteOriginalInvoiceRef) SetType(v string)`

SetType sets Type field to given value.


### GetIssuedAt

`func (o *CreditNoteOriginalInvoiceRef) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *CreditNoteOriginalInvoiceRef) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *CreditNoteOriginalInvoiceRef) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *CreditNoteOriginalInvoiceRef) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetNumber

`func (o *CreditNoteOriginalInvoiceRef) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreditNoteOriginalInvoiceRef) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreditNoteOriginalInvoiceRef) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreditNoteOriginalInvoiceRef) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUrl

`func (o *CreditNoteOriginalInvoiceRef) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreditNoteOriginalInvoiceRef) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreditNoteOriginalInvoiceRef) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReason

`func (o *CreditNoteOriginalInvoiceRef) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditNoteOriginalInvoiceRef) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditNoteOriginalInvoiceRef) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreditNoteOriginalInvoiceRef) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDescription

`func (o *CreditNoteOriginalInvoiceRef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditNoteOriginalInvoiceRef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditNoteOriginalInvoiceRef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditNoteOriginalInvoiceRef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


