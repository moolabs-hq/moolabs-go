# InvoiceLineAmountDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Id** | **string** | ID of the charge or discount. | [readonly] 
**Reason** | [**BillingDiscountReason**](BillingDiscountReason.md) | Reason code. | [readonly] 
**Description** | Pointer to **string** | Text description as to why the discount was applied. | [optional] [readonly] 
**ExternalIds** | Pointer to [**InvoiceLineAppExternalIds**](InvoiceLineAppExternalIds.md) | External IDs of the invoice in other apps such as Stripe. | [optional] [readonly] 
**Amount** | **string** | Fixed discount amount to apply (calculated if percent present). | [readonly] 

## Methods

### NewInvoiceLineAmountDiscount

`func NewInvoiceLineAmountDiscount(createdAt time.Time, updatedAt time.Time, id string, reason BillingDiscountReason, amount string, ) *InvoiceLineAmountDiscount`

NewInvoiceLineAmountDiscount instantiates a new InvoiceLineAmountDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineAmountDiscountWithDefaults

`func NewInvoiceLineAmountDiscountWithDefaults() *InvoiceLineAmountDiscount`

NewInvoiceLineAmountDiscountWithDefaults instantiates a new InvoiceLineAmountDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InvoiceLineAmountDiscount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceLineAmountDiscount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceLineAmountDiscount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceLineAmountDiscount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceLineAmountDiscount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceLineAmountDiscount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *InvoiceLineAmountDiscount) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InvoiceLineAmountDiscount) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InvoiceLineAmountDiscount) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InvoiceLineAmountDiscount) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *InvoiceLineAmountDiscount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineAmountDiscount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineAmountDiscount) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *InvoiceLineAmountDiscount) GetReason() BillingDiscountReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceLineAmountDiscount) GetReasonOk() (*BillingDiscountReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceLineAmountDiscount) SetReason(v BillingDiscountReason)`

SetReason sets Reason field to given value.


### GetDescription

`func (o *InvoiceLineAmountDiscount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineAmountDiscount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineAmountDiscount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLineAmountDiscount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalIds

`func (o *InvoiceLineAmountDiscount) GetExternalIds() InvoiceLineAppExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *InvoiceLineAmountDiscount) GetExternalIdsOk() (*InvoiceLineAppExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *InvoiceLineAmountDiscount) SetExternalIds(v InvoiceLineAppExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *InvoiceLineAmountDiscount) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetAmount

`func (o *InvoiceLineAmountDiscount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceLineAmountDiscount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceLineAmountDiscount) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


