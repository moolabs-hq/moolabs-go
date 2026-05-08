# RateCardFlatFee

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
**BillingCadence** | **string** | The billing cadence of the rate card. When null it means it is a one time fee. | 
**Price** | [**FlatPriceWithPaymentTerm**](FlatPriceWithPaymentTerm.md) | The price of the rate card. When null, the feature or service is free. | 
**Discounts** | Pointer to [**Discounts**](Discounts.md) | The discount of the rate card. For flat fee rate cards only percentage discounts are supported. Only available when price is set. | [optional] 

## Methods

### NewRateCardFlatFee

`func NewRateCardFlatFee(type_ string, key string, name string, billingCadence string, price FlatPriceWithPaymentTerm, ) *RateCardFlatFee`

NewRateCardFlatFee instantiates a new RateCardFlatFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateCardFlatFeeWithDefaults

`func NewRateCardFlatFeeWithDefaults() *RateCardFlatFee`

NewRateCardFlatFeeWithDefaults instantiates a new RateCardFlatFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RateCardFlatFee) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateCardFlatFee) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateCardFlatFee) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *RateCardFlatFee) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RateCardFlatFee) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RateCardFlatFee) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *RateCardFlatFee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateCardFlatFee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateCardFlatFee) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RateCardFlatFee) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RateCardFlatFee) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RateCardFlatFee) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RateCardFlatFee) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *RateCardFlatFee) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RateCardFlatFee) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RateCardFlatFee) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RateCardFlatFee) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFeatureKey

`func (o *RateCardFlatFee) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *RateCardFlatFee) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *RateCardFlatFee) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *RateCardFlatFee) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetEntitlementTemplate

`func (o *RateCardFlatFee) GetEntitlementTemplate() RateCardEntitlement`

GetEntitlementTemplate returns the EntitlementTemplate field if non-nil, zero value otherwise.

### GetEntitlementTemplateOk

`func (o *RateCardFlatFee) GetEntitlementTemplateOk() (*RateCardEntitlement, bool)`

GetEntitlementTemplateOk returns a tuple with the EntitlementTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTemplate

`func (o *RateCardFlatFee) SetEntitlementTemplate(v RateCardEntitlement)`

SetEntitlementTemplate sets EntitlementTemplate field to given value.

### HasEntitlementTemplate

`func (o *RateCardFlatFee) HasEntitlementTemplate() bool`

HasEntitlementTemplate returns a boolean if a field has been set.

### GetTaxConfig

`func (o *RateCardFlatFee) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *RateCardFlatFee) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *RateCardFlatFee) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *RateCardFlatFee) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.

### GetBillingCadence

`func (o *RateCardFlatFee) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *RateCardFlatFee) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *RateCardFlatFee) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetPrice

`func (o *RateCardFlatFee) GetPrice() FlatPriceWithPaymentTerm`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *RateCardFlatFee) GetPriceOk() (*FlatPriceWithPaymentTerm, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *RateCardFlatFee) SetPrice(v FlatPriceWithPaymentTerm)`

SetPrice sets Price field to given value.


### GetDiscounts

`func (o *RateCardFlatFee) GetDiscounts() Discounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *RateCardFlatFee) GetDiscountsOk() (*Discounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *RateCardFlatFee) SetDiscounts(v Discounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *RateCardFlatFee) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


