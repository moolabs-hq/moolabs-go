# SubscriptionItemIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | [**Feature**](Feature.md) | The feature the customer is entitled to use. | 
**Entitlement** | Pointer to [**Entitlement**](Entitlement.md) | The entitlement of the Subscription Item. | [optional] 

## Methods

### NewSubscriptionItemIncluded

`func NewSubscriptionItemIncluded(feature Feature, ) *SubscriptionItemIncluded`

NewSubscriptionItemIncluded instantiates a new SubscriptionItemIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemIncludedWithDefaults

`func NewSubscriptionItemIncludedWithDefaults() *SubscriptionItemIncluded`

NewSubscriptionItemIncludedWithDefaults instantiates a new SubscriptionItemIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *SubscriptionItemIncluded) GetFeature() Feature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *SubscriptionItemIncluded) GetFeatureOk() (*Feature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *SubscriptionItemIncluded) SetFeature(v Feature)`

SetFeature sets Feature field to given value.


### GetEntitlement

`func (o *SubscriptionItemIncluded) GetEntitlement() Entitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *SubscriptionItemIncluded) GetEntitlementOk() (*Entitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *SubscriptionItemIncluded) SetEntitlement(v Entitlement)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *SubscriptionItemIncluded) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


