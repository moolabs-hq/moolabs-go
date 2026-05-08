# CustomInvoicingSyncResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | Pointer to **string** | If set the invoice&#39;s number will be set to this value. | [optional] 
**ExternalId** | Pointer to **string** | If set the invoice&#39;s invoicing external ID will be set to this value. | [optional] 
**LineExternalIds** | Pointer to [**[]CustomInvoicingLineExternalIdMapping**](CustomInvoicingLineExternalIdMapping.md) | If set the invoice&#39;s line external IDs will be set to this value.  This can be used to reference the external system&#39;s entities in the invoice. | [optional] 
**LineDiscountExternalIds** | Pointer to [**[]CustomInvoicingLineDiscountExternalIdMapping**](CustomInvoicingLineDiscountExternalIdMapping.md) | If set the invoice&#39;s line discount external IDs will be set to this value.  This can be used to reference the external system&#39;s entities in the invoice. | [optional] 

## Methods

### NewCustomInvoicingSyncResult

`func NewCustomInvoicingSyncResult() *CustomInvoicingSyncResult`

NewCustomInvoicingSyncResult instantiates a new CustomInvoicingSyncResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingSyncResultWithDefaults

`func NewCustomInvoicingSyncResultWithDefaults() *CustomInvoicingSyncResult`

NewCustomInvoicingSyncResultWithDefaults instantiates a new CustomInvoicingSyncResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *CustomInvoicingSyncResult) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CustomInvoicingSyncResult) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CustomInvoicingSyncResult) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *CustomInvoicingSyncResult) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetExternalId

`func (o *CustomInvoicingSyncResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CustomInvoicingSyncResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CustomInvoicingSyncResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CustomInvoicingSyncResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLineExternalIds

`func (o *CustomInvoicingSyncResult) GetLineExternalIds() []CustomInvoicingLineExternalIdMapping`

GetLineExternalIds returns the LineExternalIds field if non-nil, zero value otherwise.

### GetLineExternalIdsOk

`func (o *CustomInvoicingSyncResult) GetLineExternalIdsOk() (*[]CustomInvoicingLineExternalIdMapping, bool)`

GetLineExternalIdsOk returns a tuple with the LineExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineExternalIds

`func (o *CustomInvoicingSyncResult) SetLineExternalIds(v []CustomInvoicingLineExternalIdMapping)`

SetLineExternalIds sets LineExternalIds field to given value.

### HasLineExternalIds

`func (o *CustomInvoicingSyncResult) HasLineExternalIds() bool`

HasLineExternalIds returns a boolean if a field has been set.

### GetLineDiscountExternalIds

`func (o *CustomInvoicingSyncResult) GetLineDiscountExternalIds() []CustomInvoicingLineDiscountExternalIdMapping`

GetLineDiscountExternalIds returns the LineDiscountExternalIds field if non-nil, zero value otherwise.

### GetLineDiscountExternalIdsOk

`func (o *CustomInvoicingSyncResult) GetLineDiscountExternalIdsOk() (*[]CustomInvoicingLineDiscountExternalIdMapping, bool)`

GetLineDiscountExternalIdsOk returns a tuple with the LineDiscountExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineDiscountExternalIds

`func (o *CustomInvoicingSyncResult) SetLineDiscountExternalIds(v []CustomInvoicingLineDiscountExternalIdMapping)`

SetLineDiscountExternalIds sets LineDiscountExternalIds field to given value.

### HasLineDiscountExternalIds

`func (o *CustomInvoicingSyncResult) HasLineDiscountExternalIds() bool`

HasLineDiscountExternalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


