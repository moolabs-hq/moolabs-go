# InvoiceLine

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
**InvoiceAt** | **time.Time** | The time this line item should be invoiced. | 
**ExternalIds** | Pointer to [**InvoiceLineAppExternalIds**](InvoiceLineAppExternalIds.md) | External IDs of the invoice in other apps such as Stripe. | [optional] [readonly] 
**Subscription** | Pointer to [**InvoiceLineSubscriptionReference**](InvoiceLineSubscriptionReference.md) | Subscription are the references to the subscritpions that this line is related to. | [optional] [readonly] 
**Type** | **string** | Type of the line. | [readonly] 
**Price** | Pointer to [**RateCardUsageBasedPrice**](RateCardUsageBasedPrice.md) | Price of the usage-based item being sold. | [optional] 
**FeatureKey** | Pointer to **string** | The feature that the usage is based on. | [optional] 
**Children** | Pointer to [**[]InvoiceDetailedLine**](InvoiceDetailedLine.md) | The lines detailing the item or service sold. | [optional] 
**RateCard** | Pointer to [**InvoiceUsageBasedRateCard**](InvoiceUsageBasedRateCard.md) | The rate card that is used for this line.  The rate card captures the intent of the price and discounts for the usage-based item. | [optional] 
**Quantity** | Pointer to **string** | The quantity of the item being sold.  Any usage discounts applied previously are deducted from this quantity. | [optional] [readonly] 
**MeteredQuantity** | Pointer to **string** | The quantity of the item that has been metered for the period before any discounts were applied. | [optional] [readonly] 
**PreLinePeriodQuantity** | Pointer to **string** | The quantity of the item used before this line&#39;s period.  It is non-zero in case of progressive billing, when this shows how much of the usage was already billed.  Any usage discounts applied previously are deducted from this quantity. | [optional] [readonly] 
**MeteredPreLinePeriodQuantity** | Pointer to **string** | The metered quantity of the item used in before this line&#39;s period without any discounts applied.  It is non-zero in case of progressive billing, when this shows how much of the usage was already billed. | [optional] [readonly] 

## Methods

### NewInvoiceLine

`func NewInvoiceLine(name string, createdAt time.Time, updatedAt time.Time, id string, managedBy InvoiceLineManagedBy, status InvoiceLineStatus, currency string, totals InvoiceTotals, period Period, invoiceAt time.Time, type_ string, ) *InvoiceLine`

NewInvoiceLine instantiates a new InvoiceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineWithDefaults

`func NewInvoiceLineWithDefaults() *InvoiceLine`

NewInvoiceLineWithDefaults instantiates a new InvoiceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceLine) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvoiceLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceLine) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceLine) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceLine) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceLine) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceLine) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceLine) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceLine) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceLine) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceLine) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceLine) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *InvoiceLine) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *InvoiceLine) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *InvoiceLine) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *InvoiceLine) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *InvoiceLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLine) SetId(v string)`

SetId sets Id field to given value.


### GetManagedBy

`func (o *InvoiceLine) GetManagedBy() InvoiceLineManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *InvoiceLine) GetManagedByOk() (*InvoiceLineManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *InvoiceLine) SetManagedBy(v InvoiceLineManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### GetStatus

`func (o *InvoiceLine) GetStatus() InvoiceLineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceLine) GetStatusOk() (*InvoiceLineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceLine) SetStatus(v InvoiceLineStatus)`

SetStatus sets Status field to given value.


### GetDiscounts

`func (o *InvoiceLine) GetDiscounts() InvoiceLineDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceLine) GetDiscountsOk() (*InvoiceLineDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceLine) SetDiscounts(v InvoiceLineDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *InvoiceLine) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoiceLine) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceLine) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceLine) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceLine) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceLine) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLine) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLine) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTaxes

`func (o *InvoiceLine) GetTaxes() []InvoiceLineTaxItem`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *InvoiceLine) GetTaxesOk() (*[]InvoiceLineTaxItem, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *InvoiceLine) SetTaxes(v []InvoiceLineTaxItem)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *InvoiceLine) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoiceLine) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceLine) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceLine) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceLine) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetTotals

`func (o *InvoiceLine) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *InvoiceLine) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *InvoiceLine) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetPeriod

`func (o *InvoiceLine) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvoiceLine) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvoiceLine) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetInvoiceAt

`func (o *InvoiceLine) GetInvoiceAt() time.Time`

GetInvoiceAt returns the InvoiceAt field if non-nil, zero value otherwise.

### GetInvoiceAtOk

`func (o *InvoiceLine) GetInvoiceAtOk() (*time.Time, bool)`

GetInvoiceAtOk returns a tuple with the InvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAt

`func (o *InvoiceLine) SetInvoiceAt(v time.Time)`

SetInvoiceAt sets InvoiceAt field to given value.


### GetExternalIds

`func (o *InvoiceLine) GetExternalIds() InvoiceLineAppExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *InvoiceLine) GetExternalIdsOk() (*InvoiceLineAppExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *InvoiceLine) SetExternalIds(v InvoiceLineAppExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *InvoiceLine) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetSubscription

`func (o *InvoiceLine) GetSubscription() InvoiceLineSubscriptionReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *InvoiceLine) GetSubscriptionOk() (*InvoiceLineSubscriptionReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *InvoiceLine) SetSubscription(v InvoiceLineSubscriptionReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *InvoiceLine) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetType

`func (o *InvoiceLine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceLine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceLine) SetType(v string)`

SetType sets Type field to given value.


### GetPrice

`func (o *InvoiceLine) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceLine) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceLine) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoiceLine) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeatureKey

`func (o *InvoiceLine) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *InvoiceLine) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *InvoiceLine) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *InvoiceLine) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetChildren

`func (o *InvoiceLine) GetChildren() []InvoiceDetailedLine`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *InvoiceLine) GetChildrenOk() (*[]InvoiceDetailedLine, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *InvoiceLine) SetChildren(v []InvoiceDetailedLine)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *InvoiceLine) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRateCard

`func (o *InvoiceLine) GetRateCard() InvoiceUsageBasedRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *InvoiceLine) GetRateCardOk() (*InvoiceUsageBasedRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *InvoiceLine) SetRateCard(v InvoiceUsageBasedRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *InvoiceLine) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceLine) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLine) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLine) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetMeteredQuantity

`func (o *InvoiceLine) GetMeteredQuantity() string`

GetMeteredQuantity returns the MeteredQuantity field if non-nil, zero value otherwise.

### GetMeteredQuantityOk

`func (o *InvoiceLine) GetMeteredQuantityOk() (*string, bool)`

GetMeteredQuantityOk returns a tuple with the MeteredQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteredQuantity

`func (o *InvoiceLine) SetMeteredQuantity(v string)`

SetMeteredQuantity sets MeteredQuantity field to given value.

### HasMeteredQuantity

`func (o *InvoiceLine) HasMeteredQuantity() bool`

HasMeteredQuantity returns a boolean if a field has been set.

### GetPreLinePeriodQuantity

`func (o *InvoiceLine) GetPreLinePeriodQuantity() string`

GetPreLinePeriodQuantity returns the PreLinePeriodQuantity field if non-nil, zero value otherwise.

### GetPreLinePeriodQuantityOk

`func (o *InvoiceLine) GetPreLinePeriodQuantityOk() (*string, bool)`

GetPreLinePeriodQuantityOk returns a tuple with the PreLinePeriodQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLinePeriodQuantity

`func (o *InvoiceLine) SetPreLinePeriodQuantity(v string)`

SetPreLinePeriodQuantity sets PreLinePeriodQuantity field to given value.

### HasPreLinePeriodQuantity

`func (o *InvoiceLine) HasPreLinePeriodQuantity() bool`

HasPreLinePeriodQuantity returns a boolean if a field has been set.

### GetMeteredPreLinePeriodQuantity

`func (o *InvoiceLine) GetMeteredPreLinePeriodQuantity() string`

GetMeteredPreLinePeriodQuantity returns the MeteredPreLinePeriodQuantity field if non-nil, zero value otherwise.

### GetMeteredPreLinePeriodQuantityOk

`func (o *InvoiceLine) GetMeteredPreLinePeriodQuantityOk() (*string, bool)`

GetMeteredPreLinePeriodQuantityOk returns a tuple with the MeteredPreLinePeriodQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteredPreLinePeriodQuantity

`func (o *InvoiceLine) SetMeteredPreLinePeriodQuantity(v string)`

SetMeteredPreLinePeriodQuantity sets MeteredPreLinePeriodQuantity field to given value.

### HasMeteredPreLinePeriodQuantity

`func (o *InvoiceLine) HasMeteredPreLinePeriodQuantity() bool`

HasMeteredPreLinePeriodQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


