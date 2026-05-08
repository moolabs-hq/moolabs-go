# InvoiceLineReplaceUpdate

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
**Id** | Pointer to **string** | The ID of the line. | [optional] 

## Methods

### NewInvoiceLineReplaceUpdate

`func NewInvoiceLineReplaceUpdate(name string, period Period, invoiceAt time.Time, ) *InvoiceLineReplaceUpdate`

NewInvoiceLineReplaceUpdate instantiates a new InvoiceLineReplaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineReplaceUpdateWithDefaults

`func NewInvoiceLineReplaceUpdateWithDefaults() *InvoiceLineReplaceUpdate`

NewInvoiceLineReplaceUpdateWithDefaults instantiates a new InvoiceLineReplaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceLineReplaceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceLineReplaceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceLineReplaceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InvoiceLineReplaceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineReplaceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineReplaceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLineReplaceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *InvoiceLineReplaceUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoiceLineReplaceUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoiceLineReplaceUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoiceLineReplaceUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoiceLineReplaceUpdate) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceLineReplaceUpdate) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceLineReplaceUpdate) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceLineReplaceUpdate) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetPeriod

`func (o *InvoiceLineReplaceUpdate) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InvoiceLineReplaceUpdate) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InvoiceLineReplaceUpdate) SetPeriod(v Period)`

SetPeriod sets Period field to given value.


### GetInvoiceAt

`func (o *InvoiceLineReplaceUpdate) GetInvoiceAt() time.Time`

GetInvoiceAt returns the InvoiceAt field if non-nil, zero value otherwise.

### GetInvoiceAtOk

`func (o *InvoiceLineReplaceUpdate) GetInvoiceAtOk() (*time.Time, bool)`

GetInvoiceAtOk returns a tuple with the InvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAt

`func (o *InvoiceLineReplaceUpdate) SetInvoiceAt(v time.Time)`

SetInvoiceAt sets InvoiceAt field to given value.


### GetPrice

`func (o *InvoiceLineReplaceUpdate) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceLineReplaceUpdate) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceLineReplaceUpdate) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoiceLineReplaceUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeatureKey

`func (o *InvoiceLineReplaceUpdate) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *InvoiceLineReplaceUpdate) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *InvoiceLineReplaceUpdate) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *InvoiceLineReplaceUpdate) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetRateCard

`func (o *InvoiceLineReplaceUpdate) GetRateCard() InvoiceUsageBasedRateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *InvoiceLineReplaceUpdate) GetRateCardOk() (*InvoiceUsageBasedRateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *InvoiceLineReplaceUpdate) SetRateCard(v InvoiceUsageBasedRateCard)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *InvoiceLineReplaceUpdate) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetId

`func (o *InvoiceLineReplaceUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineReplaceUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineReplaceUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceLineReplaceUpdate) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


