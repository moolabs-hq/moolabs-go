# InvoiceLineUsageDiscount

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
**Quantity** | **string** | The usage to apply. | [readonly] 
**PreLinePeriodQuantity** | Pointer to **string** | The usage discount already applied to the previous split lines.  Only set if progressive billing is enabled and the line is a split line. | [optional] [readonly] 

## Methods

### NewInvoiceLineUsageDiscount

`func NewInvoiceLineUsageDiscount(createdAt time.Time, updatedAt time.Time, id string, reason BillingDiscountReason, quantity string, ) *InvoiceLineUsageDiscount`

NewInvoiceLineUsageDiscount instantiates a new InvoiceLineUsageDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUsageDiscountWithDefaults

`func NewInvoiceLineUsageDiscountWithDefaults() *InvoiceLineUsageDiscount`

NewInvoiceLineUsageDiscountWithDefaults instantiates a new InvoiceLineUsageDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InvoiceLineUsageDiscount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceLineUsageDiscount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceLineUsageDiscount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceLineUsageDiscount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceLineUsageDiscount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceLineUsageDiscount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *InvoiceLineUsageDiscount) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InvoiceLineUsageDiscount) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InvoiceLineUsageDiscount) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InvoiceLineUsageDiscount) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *InvoiceLineUsageDiscount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineUsageDiscount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineUsageDiscount) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *InvoiceLineUsageDiscount) GetReason() BillingDiscountReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceLineUsageDiscount) GetReasonOk() (*BillingDiscountReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceLineUsageDiscount) SetReason(v BillingDiscountReason)`

SetReason sets Reason field to given value.


### GetDescription

`func (o *InvoiceLineUsageDiscount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineUsageDiscount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineUsageDiscount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLineUsageDiscount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalIds

`func (o *InvoiceLineUsageDiscount) GetExternalIds() InvoiceLineAppExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *InvoiceLineUsageDiscount) GetExternalIdsOk() (*InvoiceLineAppExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *InvoiceLineUsageDiscount) SetExternalIds(v InvoiceLineAppExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *InvoiceLineUsageDiscount) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceLineUsageDiscount) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineUsageDiscount) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineUsageDiscount) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetPreLinePeriodQuantity

`func (o *InvoiceLineUsageDiscount) GetPreLinePeriodQuantity() string`

GetPreLinePeriodQuantity returns the PreLinePeriodQuantity field if non-nil, zero value otherwise.

### GetPreLinePeriodQuantityOk

`func (o *InvoiceLineUsageDiscount) GetPreLinePeriodQuantityOk() (*string, bool)`

GetPreLinePeriodQuantityOk returns a tuple with the PreLinePeriodQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLinePeriodQuantity

`func (o *InvoiceLineUsageDiscount) SetPreLinePeriodQuantity(v string)`

SetPreLinePeriodQuantity sets PreLinePeriodQuantity field to given value.

### HasPreLinePeriodQuantity

`func (o *InvoiceLineUsageDiscount) HasPreLinePeriodQuantity() bool`

HasPreLinePeriodQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


