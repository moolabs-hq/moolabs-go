# AppApiV1SubscriptionsRouterSubscriptionSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Subscription ID from OpenMeter | 
**CustomerId** | Pointer to **string** | Customer ID | [optional] 
**Status** | Pointer to **string** | Subscription status | [optional] [default to "active"]
**BillingCadence** | Pointer to **string** | Billing cadence (ISO 8601 duration) | [optional] [default to "P1M"]
**BillingAnchor** | **string** | Billing anchor date | 
**PlanId** | **string** | Plan ID | 
**PlanVersion** | Pointer to **string** | Plan version | [optional] [default to "1"]
**ActiveFrom** | **string** | Active from date (ISO 8601) | 
**ActiveTo** | Pointer to **string** | Active to date (ISO 8601) | [optional] 
**EffectiveAt** | Pointer to **string** | When this change takes effect (ISO 8601) | [optional] 
**FeatureOverrides** | Pointer to [**[]FeatureOverrideRequest**](FeatureOverrideRequest.md) | Per-feature price overrides | [optional] 
**FeaturePrices** | Pointer to [**[]FeaturePriceRequest**](FeaturePriceRequest.md) | Per-feature pricing | [optional] 

## Methods

### NewAppApiV1SubscriptionsRouterSubscriptionSyncRequest

`func NewAppApiV1SubscriptionsRouterSubscriptionSyncRequest(subscriptionId string, billingAnchor string, planId string, activeFrom string, ) *AppApiV1SubscriptionsRouterSubscriptionSyncRequest`

NewAppApiV1SubscriptionsRouterSubscriptionSyncRequest instantiates a new AppApiV1SubscriptionsRouterSubscriptionSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiV1SubscriptionsRouterSubscriptionSyncRequestWithDefaults

`func NewAppApiV1SubscriptionsRouterSubscriptionSyncRequestWithDefaults() *AppApiV1SubscriptionsRouterSubscriptionSyncRequest`

NewAppApiV1SubscriptionsRouterSubscriptionSyncRequestWithDefaults instantiates a new AppApiV1SubscriptionsRouterSubscriptionSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetStatus

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBillingCadence

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.


### GetPlanId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlanVersion

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetPlanVersion() string`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetPlanVersionOk() (*string, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetPlanVersion(v string)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetActiveFrom

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetActiveFrom() string`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetActiveFromOk() (*string, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetActiveFrom(v string)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetActiveTo() string`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetActiveToOk() (*string, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetActiveTo(v string)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetFeatureOverrides

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetFeatureOverrides() []FeatureOverrideRequest`

GetFeatureOverrides returns the FeatureOverrides field if non-nil, zero value otherwise.

### GetFeatureOverridesOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetFeatureOverridesOk() (*[]FeatureOverrideRequest, bool)`

GetFeatureOverridesOk returns a tuple with the FeatureOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureOverrides

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetFeatureOverrides(v []FeatureOverrideRequest)`

SetFeatureOverrides sets FeatureOverrides field to given value.

### HasFeatureOverrides

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasFeatureOverrides() bool`

HasFeatureOverrides returns a boolean if a field has been set.

### GetFeaturePrices

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetFeaturePrices() []FeaturePriceRequest`

GetFeaturePrices returns the FeaturePrices field if non-nil, zero value otherwise.

### GetFeaturePricesOk

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) GetFeaturePricesOk() (*[]FeaturePriceRequest, bool)`

GetFeaturePricesOk returns a tuple with the FeaturePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePrices

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) SetFeaturePrices(v []FeaturePriceRequest)`

SetFeaturePrices sets FeaturePrices field to given value.

### HasFeaturePrices

`func (o *AppApiV1SubscriptionsRouterSubscriptionSyncRequest) HasFeaturePrices() bool`

HasFeaturePrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


