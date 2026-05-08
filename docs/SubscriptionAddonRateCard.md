# SubscriptionAddonRateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RateCard** | [**RateCard**](RateCard.md) | The rate card. | 
**AffectedSubscriptionItemIds** | **[]string** | The IDs of the subscription items that this rate card belongs to. | 

## Methods

### NewSubscriptionAddonRateCard

`func NewSubscriptionAddonRateCard(rateCard RateCard, affectedSubscriptionItemIds []string, ) *SubscriptionAddonRateCard`

NewSubscriptionAddonRateCard instantiates a new SubscriptionAddonRateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAddonRateCardWithDefaults

`func NewSubscriptionAddonRateCardWithDefaults() *SubscriptionAddonRateCard`

NewSubscriptionAddonRateCardWithDefaults instantiates a new SubscriptionAddonRateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRateCard

`func (o *SubscriptionAddonRateCard) GetRateCard() RateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *SubscriptionAddonRateCard) GetRateCardOk() (*RateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *SubscriptionAddonRateCard) SetRateCard(v RateCard)`

SetRateCard sets RateCard field to given value.


### GetAffectedSubscriptionItemIds

`func (o *SubscriptionAddonRateCard) GetAffectedSubscriptionItemIds() []string`

GetAffectedSubscriptionItemIds returns the AffectedSubscriptionItemIds field if non-nil, zero value otherwise.

### GetAffectedSubscriptionItemIdsOk

`func (o *SubscriptionAddonRateCard) GetAffectedSubscriptionItemIdsOk() (*[]string, bool)`

GetAffectedSubscriptionItemIdsOk returns a tuple with the AffectedSubscriptionItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSubscriptionItemIds

`func (o *SubscriptionAddonRateCard) SetAffectedSubscriptionItemIds(v []string)`

SetAffectedSubscriptionItemIds sets AffectedSubscriptionItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


