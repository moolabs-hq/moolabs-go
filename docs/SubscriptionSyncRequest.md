# SubscriptionSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Subscription ID from OpenMeter | 
**CustomerId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** | Subscription status | [optional] [default to "active"]
**BillingCadence** | Pointer to **string** | Billing cadence (ISO 8601 duration) | [optional] [default to "P1M"]
**BillingAnchor** | **string** | Billing anchor date | 
**PlanId** | **string** | Plan ID | 
**PlanVersion** | Pointer to **string** | Plan version | [optional] [default to "1"]
**ActiveFrom** | **string** | Active from date (ISO 8601) | 
**ActiveTo** | Pointer to **NullableString** |  | [optional] 
**EffectiveAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubscriptionSyncRequest

`func NewSubscriptionSyncRequest(subscriptionId string, billingAnchor string, planId string, activeFrom string, ) *SubscriptionSyncRequest`

NewSubscriptionSyncRequest instantiates a new SubscriptionSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionSyncRequestWithDefaults

`func NewSubscriptionSyncRequestWithDefaults() *SubscriptionSyncRequest`

NewSubscriptionSyncRequestWithDefaults instantiates a new SubscriptionSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SubscriptionSyncRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionSyncRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionSyncRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCustomerId

`func (o *SubscriptionSyncRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionSyncRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionSyncRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionSyncRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *SubscriptionSyncRequest) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SubscriptionSyncRequest) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetStatus

`func (o *SubscriptionSyncRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionSyncRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionSyncRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionSyncRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBillingCadence

`func (o *SubscriptionSyncRequest) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *SubscriptionSyncRequest) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *SubscriptionSyncRequest) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *SubscriptionSyncRequest) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *SubscriptionSyncRequest) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *SubscriptionSyncRequest) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *SubscriptionSyncRequest) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.


### GetPlanId

`func (o *SubscriptionSyncRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionSyncRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionSyncRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetPlanVersion

`func (o *SubscriptionSyncRequest) GetPlanVersion() string`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *SubscriptionSyncRequest) GetPlanVersionOk() (*string, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *SubscriptionSyncRequest) SetPlanVersion(v string)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *SubscriptionSyncRequest) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetActiveFrom

`func (o *SubscriptionSyncRequest) GetActiveFrom() string`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionSyncRequest) GetActiveFromOk() (*string, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionSyncRequest) SetActiveFrom(v string)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *SubscriptionSyncRequest) GetActiveTo() string`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionSyncRequest) GetActiveToOk() (*string, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionSyncRequest) SetActiveTo(v string)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionSyncRequest) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### SetActiveToNil

`func (o *SubscriptionSyncRequest) SetActiveToNil(b bool)`

 SetActiveToNil sets the value for ActiveTo to be an explicit nil

### UnsetActiveTo
`func (o *SubscriptionSyncRequest) UnsetActiveTo()`

UnsetActiveTo ensures that no value is present for ActiveTo, not even an explicit nil
### GetEffectiveAt

`func (o *SubscriptionSyncRequest) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *SubscriptionSyncRequest) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *SubscriptionSyncRequest) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *SubscriptionSyncRequest) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### SetEffectiveAtNil

`func (o *SubscriptionSyncRequest) SetEffectiveAtNil(b bool)`

 SetEffectiveAtNil sets the value for EffectiveAt to be an explicit nil

### UnsetEffectiveAt
`func (o *SubscriptionSyncRequest) UnsetEffectiveAt()`

UnsetEffectiveAt ensures that no value is present for EffectiveAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


