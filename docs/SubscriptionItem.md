# SubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**ActiveFrom** | **time.Time** | The cadence start of the resource. | 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] 
**Key** | **string** | The identifier of the RateCard. SubscriptionItem/RateCard can be identified, it has a reference:  1. If a Feature is associated with the SubscriptionItem, it is identified by the Feature 1.1 It can be an ID reference, for an exact version of the Feature (Features can change across versions) 1.2 It can be a Key reference, which always refers to the latest (active or inactive) version of a Feature  2. If a Feature is not associated with the SubscriptionItem, it is referenced by the Price  We say \&quot;referenced by the Price\&quot; regardless of how a price itself is referenced, it colloquially makes sense to say \&quot;paying the same price for the same thing\&quot;. In practice this should be derived from what&#39;s printed on the invoice line-item. | 
**FeatureKey** | Pointer to **string** | The feature&#39;s key (if present). | [optional] 
**BillingCadence** | **string** | The billing cadence of the rate card. When null, the rate card is a one-time purchase. | 
**Price** | [**RateCardUsageBasedPrice**](RateCardUsageBasedPrice.md) | The price of the rate card. When null, the feature or service is free. | 
**Discounts** | Pointer to [**Discounts**](Discounts.md) | The discounts applied to the rate card. | [optional] 
**Included** | Pointer to [**SubscriptionItemIncluded**](SubscriptionItemIncluded.md) | Describes what access is gained via the SubscriptionItem | [optional] 
**TaxConfig** | Pointer to [**TaxConfig**](TaxConfig.md) | The tax config of the Subscription Item. When undefined, the tax config of the feature or the default tax config of the plan is used. | [optional] 

## Methods

### NewSubscriptionItem

`func NewSubscriptionItem(id string, name string, createdAt time.Time, updatedAt time.Time, activeFrom time.Time, key string, billingCadence string, price RateCardUsageBasedPrice, ) *SubscriptionItem`

NewSubscriptionItem instantiates a new SubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemWithDefaults

`func NewSubscriptionItemWithDefaults() *SubscriptionItem`

NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SubscriptionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionItem) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionItem) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionItem) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *SubscriptionItem) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubscriptionItem) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubscriptionItem) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubscriptionItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetActiveFrom

`func (o *SubscriptionItem) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionItem) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionItem) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *SubscriptionItem) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionItem) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionItem) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionItem) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetKey

`func (o *SubscriptionItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SubscriptionItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SubscriptionItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetFeatureKey

`func (o *SubscriptionItem) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *SubscriptionItem) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *SubscriptionItem) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *SubscriptionItem) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetBillingCadence

`func (o *SubscriptionItem) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *SubscriptionItem) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *SubscriptionItem) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetPrice

`func (o *SubscriptionItem) GetPrice() RateCardUsageBasedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionItem) GetPriceOk() (*RateCardUsageBasedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionItem) SetPrice(v RateCardUsageBasedPrice)`

SetPrice sets Price field to given value.


### GetDiscounts

`func (o *SubscriptionItem) GetDiscounts() Discounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *SubscriptionItem) GetDiscountsOk() (*Discounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *SubscriptionItem) SetDiscounts(v Discounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *SubscriptionItem) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetIncluded

`func (o *SubscriptionItem) GetIncluded() SubscriptionItemIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SubscriptionItem) GetIncludedOk() (*SubscriptionItemIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SubscriptionItem) SetIncluded(v SubscriptionItemIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *SubscriptionItem) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetTaxConfig

`func (o *SubscriptionItem) GetTaxConfig() TaxConfig`

GetTaxConfig returns the TaxConfig field if non-nil, zero value otherwise.

### GetTaxConfigOk

`func (o *SubscriptionItem) GetTaxConfigOk() (*TaxConfig, bool)`

GetTaxConfigOk returns a tuple with the TaxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfig

`func (o *SubscriptionItem) SetTaxConfig(v TaxConfig)`

SetTaxConfig sets TaxConfig field to given value.

### HasTaxConfig

`func (o *SubscriptionItem) HasTaxConfig() bool`

HasTaxConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


