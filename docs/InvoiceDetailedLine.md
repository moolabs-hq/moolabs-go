# InvoiceDetailedLine

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
**ManagedBy** | [**InvoiceLineManagedBy**](InvoiceLineManagedBy.md) | managedBy specifies if the line is manually added via the api or managed by OpenMeter. | [readonly] 
**Status** | [**InvoiceLineStatus**](InvoiceLineStatus.md) | Status of the line.  External calls always create valid lines, other line types are managed by the billing engine of OpenMeter. | [readonly] 
**Discounts** | Pointer to [**InvoiceLineDiscounts**](InvoiceLineDiscounts.md) | Discounts detailes applied to this line.  New discounts can be added via the invoice&#39;s discounts API, to facilitate discounts that are affecting multiple lines. | [optional] [readonly] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) | The invoice this item belongs to. | [optional] 
**Currency** | **string** | The currency of this line. | 
**Taxes** | Pointer to [**[]InvoiceLineTaxItem**](InvoiceLineTaxItem.md) | Taxes applied to the invoice totals. | [optional] 
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | Tax config specify the tax configuration for this line. | [optional] 
**Totals** | [**InvoiceTotals**](InvoiceTotals.md) | Totals for this line. | [readonly] 
**Period** | [**Period**](Period.md) | Period of the line item applies to for revenue recognition pruposes.  Billing always treats periods as start being inclusive and end being exclusive. | 
**ExternalIds** | Pointer to [**InvoiceLineAppExternalIds**](InvoiceLineAppExternalIds.md) | External IDs of the invoice in other apps such as Stripe. | [optional] [readonly] 
**Subscription** | Pointer to [**InvoiceLineSubscriptionReference**](InvoiceLineSubscriptionReference.md) | Subscription are the references to the subscritpions that this line is related to. | [optional] [readonly] 
**InvoiceAt** | **time.Time** | The time this line item should be invoiced. | 
**Type** | **string** | Type of the line. | [readonly] 
**PerUnitAmount** | Pointer to **string** | Price of the item being sold. | [optional] 
**PaymentTerm** | Pointer to [**PricePaymentTerm**](PricePaymentTerm.md) | Payment term of the line. | [optional] 
**Quantity** | Pointer to **string** | Quantity of the item being sold. | [optional] 
**RateCard** | Pointer to [**InvoiceDetailedLineRateCard**](InvoiceDetailedLineRateCard.md) | The rate card that is used for this line. | [optional] 
**Category** | Pointer to [**InvoiceDetailedLineCostCategory**](InvoiceDetailedLineCostCategory.md) | Category of the flat fee. | [optional] [readonly] 

## Methods

### NewInvoiceDetailedLine

`func NewInvoiceDetailedLine(name string, createdAt time.Time, updatedAt time.Time, id string, managedBy InvoiceLineManagedBy, status InvoiceLineStatus, currency string, totals InvoiceTotals, period Period, invoiceAt time.Time, type_ string, ) *InvoiceDetailedLine`

NewInvoiceDetailedLine instantiates a new InvoiceDetailedLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDetailedLineWithDefaults

`func NewInvoiceDetailedLineWithDefaults() *InvoiceDetailedLine`

NewInvoiceDetailedLineWithDefaults instantiates a new InvoiceDetailedLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceDetailedLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceDetailedLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceDetailedLine) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvoiceDetailedLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceDetailedLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceDetailedLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceDetailedLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceDetailedLine) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceDetailedLine) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceDetailedLine) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceDetailedLine) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceDetailedLine) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceDetailedLine) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceDetailedLine) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceDetailedLine) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceDetailedLine) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceDetailedLine) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *InvoiceDetailedLine) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InvoiceDetailedLine) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InvoiceDetailedLine) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InvoiceDetailedLine) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *InvoiceDetailedLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceDetailedLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceDetailedLine) SetId(v string)`

SetId sets Id field to given value.


### GetManagedBy

`func (o *InvoiceDetailedLine) GetManagedBy() InvoiceLineManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *InvoiceDetailedLine) GetManagedByOk() (*InvoiceLineManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *InvoiceDetailedLine) SetManagedBy(v InvoiceLineManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### GetStatus

`func (o *InvoiceDetailedLine) GetStatus() InvoiceLineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceDetailedLine) GetStatusOk() (*InvoiceLineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceDetailedLine) SetStatus(v InvoiceLineStatus)`

SetStatus sets Status field to given value.


### GetDiscounts

`func (o *InvoiceDetailedLine) GetDiscounts() InvoiceLineDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceDetailedLine) GetDiscountsOk() (*InvoiceLineDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceDetailedLine) SetDiscounts(v InvoiceLineDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *InvoiceDetailedLine) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoiceDetailedLine) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceDetailedLine) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceDetailedLine) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceDetailedLine) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceDetailedLine) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceDetailedLine) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceDetailedLine) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTaxes

`func (o *InvoiceDetailedLine) GetTaxes() []InvoiceLineTaxItem`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *InvoiceDetailedLine) GetTaxesOk() (*[]InvoiceLineTaxItem, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *InvoiceDetailedLine) SetTaxes(v []InvoiceLineTaxItem)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *InvoiceDetailedLine) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoiceDetailedLine) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceDetailedLine) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceDetailedLine) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceDetailedLine) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetTotals

`func (o *InvoiceDetailedLine) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InvoiceDetailedLine) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InvoiceDetailedLine) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetPeriod

`func (o *InvoiceDetailedLine) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvoiceDetailedLine) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvoiceDetailedLine) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetExternalIds

`func (o *InvoiceDetailedLine) GetExternalIds() InvoiceLineAppExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *InvoiceDetailedLine) GetExternalIdsOk() (*InvoiceLineAppExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *InvoiceDetailedLine) SetExternalIds(v InvoiceLineAppExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *InvoiceDetailedLine) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetSubscription

`func (o *InvoiceDetailedLine) GetSubscription() InvoiceLineSubscriptionReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *InvoiceDetailedLine) GetSubscriptionOk() (*InvoiceLineSubscriptionReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *InvoiceDetailedLine) SetSubscription(v InvoiceLineSubscriptionReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *InvoiceDetailedLine) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetInvoiceAt

`func (o *InvoiceDetailedLine) GetInvoiceAt() time.Time`

GetInvoiceAt returns the InvoiceAt field if non-nil, zero value otherwise.

### GetInvoiceAtOk

`func (o *InvoiceDetailedLine) GetInvoiceAtOk() (*time.Time, bool)`

GetInvoiceAtOk returns a tuple with the InvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAt

`func (o *InvoiceDetailedLine) SetInvoiceAt(v time.Time)`

SetInvoiceAt sets InvoiceAt field to given value.


### GetType

`func (o *InvoiceDetailedLine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceDetailedLine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceDetailedLine) SetType(v string)`

SetType sets Type field to given value.


### GetPerUnitAmount

`func (o *InvoiceDetailedLine) GetPerUnitAmount() string`

GetPerUnitAmount returns the PerUnitAmount field if non-nil, zero value otherwise.

### GetPerUnitAmountOk

`func (o *InvoiceDetailedLine) GetPerUnitAmountOk() (*string, bool)`

GetPerUnitAmountOk returns a tuple with the PerUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnitAmount

`func (o *InvoiceDetailedLine) SetPerUnitAmount(v string)`

SetPerUnitAmount sets PerUnitAmount field to given value.

### HasPerUnitAmount

`func (o *InvoiceDetailedLine) HasPerUnitAmount() bool`

HasPerUnitAmount returns a boolean if a field has been set.

### GetPaymentTerm

`func (o *InvoiceDetailedLine) GetPaymentTerm() PricePaymentTerm`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *InvoiceDetailedLine) GetPaymentTermOk() (*PricePaymentTerm, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *InvoiceDetailedLine) SetPaymentTerm(v PricePaymentTerm)`

SetPaymentTerm sets PaymentTerm field to given value.

### HasPaymentTerm

`func (o *InvoiceDetailedLine) HasPaymentTerm() bool`

HasPaymentTerm returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceDetailedLine) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceDetailedLine) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceDetailedLine) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceDetailedLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRateCard

`func (o *InvoiceDetailedLine) GetRateCard() InvoiceDetailedLineRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *InvoiceDetailedLine) GetRateCardOk() (*InvoiceDetailedLineRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *InvoiceDetailedLine) SetRateCard(v InvoiceDetailedLineRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *InvoiceDetailedLine) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetCategory

`func (o *InvoiceDetailedLine) GetCategory() InvoiceDetailedLineCostCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InvoiceDetailedLine) GetCategoryOk() (*InvoiceDetailedLineCostCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InvoiceDetailedLine) SetCategory(v InvoiceDetailedLineCostCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InvoiceDetailedLine) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


