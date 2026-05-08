# InvoicePendingLineCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | Tax config specify the tax configuration for this line. | [optional] 
**Period** | [**Period**](Period.md) | Period of the line item applies to for revenue recognition pruposes.  Billing always treats periods as start being inclusive and end being exclusive. | 
**InvoiceAt** | **time.Time** | The time this line item should be invoiced. | 
**Price** | Pointer to [**RateCardUsageBasedPrice**](RateCardUsageBasedPrice.md) | Price of the usage-based item being sold. | [optional] 
**FeatureKey** | Pointer to **string** | The feature that the usage is based on. | [optional] 
**RateCard** | Pointer to [**InvoiceUsageBasedRateCard**](InvoiceUsageBasedRateCard.md) | The rate card that is used for this line.  The rate card captures the intent of the price and discounts for the usage-based item. | [optional] 

## Methods

### NewInvoicePendingLineCreate

`func NewInvoicePendingLineCreate(name string, period Period, invoiceAt time.Time, ) *InvoicePendingLineCreate`

NewInvoicePendingLineCreate instantiates a new InvoicePendingLineCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePendingLineCreateWithDefaults

`func NewInvoicePendingLineCreateWithDefaults() *InvoicePendingLineCreate`

NewInvoicePendingLineCreateWithDefaults instantiates a new InvoicePendingLineCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoicePendingLineCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoicePendingLineCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoicePendingLineCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvoicePendingLineCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoicePendingLineCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoicePendingLineCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoicePendingLineCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoicePendingLineCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoicePendingLineCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoicePendingLineCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoicePendingLineCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoicePendingLineCreate) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoicePendingLineCreate) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoicePendingLineCreate) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoicePendingLineCreate) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetPeriod

`func (o *InvoicePendingLineCreate) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvoicePendingLineCreate) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvoicePendingLineCreate) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetInvoiceAt

`func (o *InvoicePendingLineCreate) GetInvoiceAt() time.Time`

GetInvoiceAt returns the InvoiceAt field if non-nil, zero value otherwise.

### GetInvoiceAtOk

`func (o *InvoicePendingLineCreate) GetInvoiceAtOk() (*time.Time, bool)`

GetInvoiceAtOk returns a tuple with the InvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAt

`func (o *InvoicePendingLineCreate) SetInvoiceAt(v time.Time)`

SetInvoiceAt sets InvoiceAt field to given value.


### GetPrice

`func (o *InvoicePendingLineCreate) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoicePendingLineCreate) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoicePendingLineCreate) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoicePendingLineCreate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeatureKey

`func (o *InvoicePendingLineCreate) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *InvoicePendingLineCreate) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *InvoicePendingLineCreate) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *InvoicePendingLineCreate) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetRateCard

`func (o *InvoicePendingLineCreate) GetRateCard() InvoiceUsageBasedRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *InvoicePendingLineCreate) GetRateCardOk() (*InvoiceUsageBasedRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *InvoicePendingLineCreate) SetRateCard(v InvoiceUsageBasedRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *InvoicePendingLineCreate) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


