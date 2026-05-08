# InvoiceGenericDocumentRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InvoiceDocumentRefType**](InvoiceDocumentRefType.md) | Type of the document referenced. | [readonly] 
**Reason** | Pointer to **string** | Human readable description on why this reference is here or needs to be used. | [optional] [readonly] 
**Description** | Pointer to **string** | Additional details about the document. | [optional] [readonly] 

## Methods

### NewInvoiceGenericDocumentRef

`func NewInvoiceGenericDocumentRef(type_ InvoiceDocumentRefType, ) *InvoiceGenericDocumentRef`

NewInvoiceGenericDocumentRef instantiates a new InvoiceGenericDocumentRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceGenericDocumentRefWithDefaults

`func NewInvoiceGenericDocumentRefWithDefaults() *InvoiceGenericDocumentRef`

NewInvoiceGenericDocumentRefWithDefaults instantiates a new InvoiceGenericDocumentRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InvoiceGenericDocumentRef) GetType() InvoiceDocumentRefType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceGenericDocumentRef) GetTypeOk() (*InvoiceDocumentRefType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceGenericDocumentRef) SetType(v InvoiceDocumentRefType)`

SetType sets Type field to given value.


### GetReason

`func (o *InvoiceGenericDocumentRef) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceGenericDocumentRef) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceGenericDocumentRef) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoiceGenericDocumentRef) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceGenericDocumentRef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceGenericDocumentRef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceGenericDocumentRef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceGenericDocumentRef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


