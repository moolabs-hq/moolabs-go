# InvoiceLineBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Id** | **string** | ID of the line. | 
**Type** | [**InvoiceLineTypes**](InvoiceLineTypes.md) | Type of the line.  A line&#39;s type cannot be changed after creation. | [readonly] 
**ManagedBy** | [**InvoiceLineManagedBy**](InvoiceLineManagedBy.md) | managedBy specifies if the line is manually added via the api or managed by OpenMeter. | [readonly] 
**Status** | [**InvoiceLineStatus**](InvoiceLineStatus.md) | Status of the line.  External calls always create valid lines, other line types are managed by the billing engine of OpenMeter. | [readonly] 
**Discounts** | Pointer to [**InvoiceLineDiscounts**](InvoiceLineDiscounts.md) | Discounts detailes applied to this line.  New discounts can be added via the invoice&#39;s discounts API, to facilitate discounts that are affecting multiple lines. | [optional] [readonly] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) | The invoice this item belongs to. | [optional] 
**Currency** | **string** | The currency of this line. | 
**Taxes** | Pointer to [**[]InvoiceLineTaxItem**](InvoiceLineTaxItem.md) | Taxes applied to the invoice totals. | [optional] 
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | Tax config specify the tax configuration for this line. | [optional] 
**Totals** | [**InvoiceTotals**](InvoiceTotals.md) | Totals for this line. | [readonly] 
**Period** | [**Period**](Period.md) | Period of the line item applies to for revenue recognition pruposes.  Billing always treats periods as start being inclusive and end being exclusive. | 
**InvoiceAt** | **time.Time** | The time this line item should be invoiced. | 
**ExternalIds** | Pointer to [**InvoiceLineAppExternalIds**](InvoiceLineAppExternalIds.md) | External IDs of the invoice in other apps such as Stripe. | [optional] [readonly] 
**Subscription** | Pointer to [**InvoiceLineSubscriptionReference**](InvoiceLineSubscriptionReference.md) | Subscription are the references to the subscritpions that this line is related to. | [optional] [readonly] 

## Methods

### NewInvoiceLineBase

`func NewInvoiceLineBase(name string, createdAt time.Time, updatedAt time.Time, id string, type_ InvoiceLineTypes, managedBy InvoiceLineManagedBy, status InvoiceLineStatus, currency string, totals InvoiceTotals, period Period, invoiceAt time.Time, ) *InvoiceLineBase`

NewInvoiceLineBase instantiates a new InvoiceLineBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineBaseWithDefaults

`func NewInvoiceLineBaseWithDefaults() *InvoiceLineBase`

NewInvoiceLineBaseWithDefaults instantiates a new InvoiceLineBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceLineBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceLineBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceLineBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvoiceLineBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLineBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceLineBase) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceLineBase) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceLineBase) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceLineBase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceLineBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceLineBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceLineBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceLineBase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceLineBase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceLineBase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *InvoiceLineBase) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InvoiceLineBase) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InvoiceLineBase) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InvoiceLineBase) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *InvoiceLineBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineBase) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InvoiceLineBase) GetType() InvoiceLineTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceLineBase) GetTypeOk() (*InvoiceLineTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceLineBase) SetType(v InvoiceLineTypes)`

SetType sets Type field to given value.


### GetManagedBy

`func (o *InvoiceLineBase) GetManagedBy() InvoiceLineManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *InvoiceLineBase) GetManagedByOk() (*InvoiceLineManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *InvoiceLineBase) SetManagedBy(v InvoiceLineManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### GetStatus

`func (o *InvoiceLineBase) GetStatus() InvoiceLineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceLineBase) GetStatusOk() (*InvoiceLineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceLineBase) SetStatus(v InvoiceLineStatus)`

SetStatus sets Status field to given value.


### GetDiscounts

`func (o *InvoiceLineBase) GetDiscounts() InvoiceLineDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceLineBase) GetDiscountsOk() (*InvoiceLineDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceLineBase) SetDiscounts(v InvoiceLineDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *InvoiceLineBase) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoiceLineBase) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceLineBase) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceLineBase) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceLineBase) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceLineBase) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLineBase) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLineBase) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTaxes

`func (o *InvoiceLineBase) GetTaxes() []InvoiceLineTaxItem`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *InvoiceLineBase) GetTaxesOk() (*[]InvoiceLineTaxItem, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *InvoiceLineBase) SetTaxes(v []InvoiceLineTaxItem)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *InvoiceLineBase) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoiceLineBase) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceLineBase) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceLineBase) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceLineBase) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetTotals

`func (o *InvoiceLineBase) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InvoiceLineBase) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InvoiceLineBase) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetPeriod

`func (o *InvoiceLineBase) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvoiceLineBase) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvoiceLineBase) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetInvoiceAt

`func (o *InvoiceLineBase) GetInvoiceAt() time.Time`

GetInvoiceAt returns the InvoiceAt field if non-nil, zero value otherwise.

### GetInvoiceAtOk

`func (o *InvoiceLineBase) GetInvoiceAtOk() (*time.Time, bool)`

GetInvoiceAtOk returns a tuple with the InvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAt

`func (o *InvoiceLineBase) SetInvoiceAt(v time.Time)`

SetInvoiceAt sets InvoiceAt field to given value.


### GetExternalIds

`func (o *InvoiceLineBase) GetExternalIds() InvoiceLineAppExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *InvoiceLineBase) GetExternalIdsOk() (*InvoiceLineAppExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *InvoiceLineBase) SetExternalIds(v InvoiceLineAppExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *InvoiceLineBase) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetSubscription

`func (o *InvoiceLineBase) GetSubscription() InvoiceLineSubscriptionReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *InvoiceLineBase) GetSubscriptionOk() (*InvoiceLineSubscriptionReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *InvoiceLineBase) SetSubscription(v InvoiceLineSubscriptionReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *InvoiceLineBase) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


