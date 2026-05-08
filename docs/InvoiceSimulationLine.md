# InvoiceSimulationLine

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
**Quantity** | **string** | The quantity of the item being sold. | 
**PreLinePeriodQuantity** | Pointer to **string** | The quantity of the item used before this line&#39;s period, if the line is billed progressively. | [optional] 
**Id** | Pointer to **string** | ID of the line. If not specified it will be auto-generated.  When discounts are specified, this must be provided, so that the discount can reference it. | [optional] 

## Methods

### NewInvoiceSimulationLine

`func NewInvoiceSimulationLine(name string, period Period, invoiceAt time.Time, quantity string, ) *InvoiceSimulationLine`

NewInvoiceSimulationLine instantiates a new InvoiceSimulationLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSimulationLineWithDefaults

`func NewInvoiceSimulationLineWithDefaults() *InvoiceSimulationLine`

NewInvoiceSimulationLineWithDefaults instantiates a new InvoiceSimulationLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceSimulationLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceSimulationLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceSimulationLine) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvoiceSimulationLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceSimulationLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceSimulationLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceSimulationLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceSimulationLine) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceSimulationLine) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceSimulationLine) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceSimulationLine) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoiceSimulationLine) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceSimulationLine) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceSimulationLine) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceSimulationLine) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetPeriod

`func (o *InvoiceSimulationLine) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvoiceSimulationLine) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvoiceSimulationLine) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetInvoiceAt

`func (o *InvoiceSimulationLine) GetInvoiceAt() time.Time`

GetInvoiceAt returns the InvoiceAt field if non-nil, zero value otherwise.

### GetInvoiceAtOk

`func (o *InvoiceSimulationLine) GetInvoiceAtOk() (*time.Time, bool)`

GetInvoiceAtOk returns a tuple with the InvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAt

`func (o *InvoiceSimulationLine) SetInvoiceAt(v time.Time)`

SetInvoiceAt sets InvoiceAt field to given value.


### GetPrice

`func (o *InvoiceSimulationLine) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceSimulationLine) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceSimulationLine) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoiceSimulationLine) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeatureKey

`func (o *InvoiceSimulationLine) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *InvoiceSimulationLine) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *InvoiceSimulationLine) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *InvoiceSimulationLine) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetRateCard

`func (o *InvoiceSimulationLine) GetRateCard() InvoiceUsageBasedRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *InvoiceSimulationLine) GetRateCardOk() (*InvoiceUsageBasedRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *InvoiceSimulationLine) SetRateCard(v InvoiceUsageBasedRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *InvoiceSimulationLine) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceSimulationLine) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceSimulationLine) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceSimulationLine) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetPreLinePeriodQuantity

`func (o *InvoiceSimulationLine) GetPreLinePeriodQuantity() string`

GetPreLinePeriodQuantity returns the PreLinePeriodQuantity field if non-nil, zero value otherwise.

### GetPreLinePeriodQuantityOk

`func (o *InvoiceSimulationLine) GetPreLinePeriodQuantityOk() (*string, bool)`

GetPreLinePeriodQuantityOk returns a tuple with the PreLinePeriodQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLinePeriodQuantity

`func (o *InvoiceSimulationLine) SetPreLinePeriodQuantity(v string)`

SetPreLinePeriodQuantity sets PreLinePeriodQuantity field to given value.

### HasPreLinePeriodQuantity

`func (o *InvoiceSimulationLine) HasPreLinePeriodQuantity() bool`

HasPreLinePeriodQuantity returns a boolean if a field has been set.

### GetId

`func (o *InvoiceSimulationLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceSimulationLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceSimulationLine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceSimulationLine) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


