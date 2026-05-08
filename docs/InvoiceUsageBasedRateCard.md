# InvoiceUsageBasedRateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | Pointer to **string** | The feature the customer is entitled to use. | [optional] 
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | The tax config of the rate card. When undefined, the tax config of the feature or the default tax config of the plan is used. | [optional] 
**Price** | [**RateCardUsageBasedPrice**](RateCardUsageBasedPrice.md) | The price of the rate card. When null, the feature or service is free. | 
**Discounts** | Pointer to [**BillingDiscounts**](BillingDiscounts.md) | The discounts that are applied to the line. | [optional] 

## Methods

### NewInvoiceUsageBasedRateCard

`func NewInvoiceUsageBasedRateCard(price RateCardUsageBasedPrice, ) *InvoiceUsageBasedRateCard`

NewInvoiceUsageBasedRateCard instantiates a new InvoiceUsageBasedRateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUsageBasedRateCardWithDefaults

`func NewInvoiceUsageBasedRateCardWithDefaults() *InvoiceUsageBasedRateCard`

NewInvoiceUsageBasedRateCardWithDefaults instantiates a new InvoiceUsageBasedRateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *InvoiceUsageBasedRateCard) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *InvoiceUsageBasedRateCard) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *InvoiceUsageBasedRateCard) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *InvoiceUsageBasedRateCard) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetTaxConfig

`func (o *InvoiceUsageBasedRateCard) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *InvoiceUsageBasedRateCard) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *InvoiceUsageBasedRateCard) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *InvoiceUsageBasedRateCard) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetPrice

`func (o *InvoiceUsageBasedRateCard) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceUsageBasedRateCard) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceUsageBasedRateCard) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.


### GetDiscounts

`func (o *InvoiceUsageBasedRateCard) GetDiscounts() BillingDiscounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *InvoiceUsageBasedRateCard) GetDiscountsOk() (*BillingDiscounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *InvoiceUsageBasedRateCard) SetDiscounts(v BillingDiscounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *InvoiceUsageBasedRateCard) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


