# RateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the RateCard. | 
**Key** | **string** | A semi-unique identifier for the resource. | 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**FeatureKey** | Pointer to **string** | The feature the customer is entitled to use. | [optional] 
**EntitlementTemplate** | Pointer to [**RateCardEntitlement**](RateCardEntitlement.md) | The entitlement of the rate card. Only available when featureKey is set. | [optional] 
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | The tax config of the rate card. When undefined, the tax config of the feature or the default tax config of the plan is used. | [optional] 
**BillingCadence** | **string** | The billing cadence of the rate card. | 
**Price** | [**RateCardUsageBasedPrice**](RateCardUsageBasedPrice.md) | The price of the rate card. When null, the feature or service is free. | 
**Discounts** | Pointer to [**Discounts**](Discounts.md) | The discounts of the rate card.  Flat fee rate cards only support percentage discounts. | [optional] 

## Methods

### NewRateCard

`func NewRateCard(type_ string, key string, name string, billingCadence string, price RateCardUsageBasedPrice, ) *RateCard`

NewRateCard instantiates a new RateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardWithDefaults

`func NewRateCardWithDefaults() *RateCard`

NewRateCardWithDefaults instantiates a new RateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RateCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCard) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *RateCard) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RateCard) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RateCard) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *RateCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateCard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RateCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RateCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RateCard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RateCard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *RateCard) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RateCard) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RateCard) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RateCard) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFeatureKey

`func (o *RateCard) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *RateCard) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *RateCard) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *RateCard) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetEntitlementTemplate

`func (o *RateCard) GetEntitlementTemplate() RateCardEntitlement`

GetEntitlementTemplate returns the EntitlementTemplate field if non-nil, zero value otherwise.

### GetEntitlementTemplateOk

`func (o *RateCard) GetEntitlementTemplateOk() (*RateCardEntitlement, bool)`

GetEntitlementTemplateOk returns a tuple with the EntitlementTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTemplate

`func (o *RateCard) SetEntitlementTemplate(v RateCardEntitlement)`

SetEntitlementTemplate sets EntitlementTemplate field to given value.

### HasEntitlementTemplate

`func (o *RateCard) HasEntitlementTemplate() bool`

HasEntitlementTemplate returns a boolean if a field has been set.

### GetTaxConfig

`func (o *RateCard) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *RateCard) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *RateCard) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *RateCard) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetBillingCadence

`func (o *RateCard) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *RateCard) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *RateCard) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetPrice

`func (o *RateCard) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RateCard) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RateCard) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.


### GetDiscounts

`func (o *RateCard) GetDiscounts() Discounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *RateCard) GetDiscountsOk() (*Discounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *RateCard) SetDiscounts(v Discounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *RateCard) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


