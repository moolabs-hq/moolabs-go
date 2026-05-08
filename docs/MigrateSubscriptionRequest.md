# MigrateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timing** | Pointer to [**SubscriptionTiming**](SubscriptionTiming.md) | Timing configuration for the migration, when the migration should take effect. If not supported by the subscription, 400 will be returned. | [optional] 
**TargetVersion** | Pointer to **int32** | The version of the plan to migrate to. If not provided, the subscription will migrate to the latest version of the current plan. | [optional] 
**StartingPhase** | Pointer to **string** | The key of the phase to start the subscription in. If not provided, the subscription will start in the first phase of the plan. | [optional] 
**BillingAnchor** | Pointer to **time.Time** | The billing anchor of the subscription. The provided date will be normalized according to the billing cadence to the nearest recurrence before start time. If not provided, the previous subscription billing anchor will be used. | [optional] 
**CommercialOverrides** | Pointer to [**CommercialOverrides**](CommercialOverrides.md) | Commercial terms for the migrated subscription. Omitted preserves the current subscription&#39;s overrides. | [optional] 

## Methods

### NewMigrateSubscriptionRequest

`func NewMigrateSubscriptionRequest() *MigrateSubscriptionRequest`

NewMigrateSubscriptionRequest instantiates a new MigrateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateSubscriptionRequestWithDefaults

`func NewMigrateSubscriptionRequestWithDefaults() *MigrateSubscriptionRequest`

NewMigrateSubscriptionRequestWithDefaults instantiates a new MigrateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiming

`func (o *MigrateSubscriptionRequest) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *MigrateSubscriptionRequest) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *MigrateSubscriptionRequest) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *MigrateSubscriptionRequest) HasTiming() bool`

HasTiming returns a boolean if a field has been set.

### GetTargetVersion

`func (o *MigrateSubscriptionRequest) GetTargetVersion() int32`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *MigrateSubscriptionRequest) GetTargetVersionOk() (*int32, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *MigrateSubscriptionRequest) SetTargetVersion(v int32)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *MigrateSubscriptionRequest) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetStartingPhase

`func (o *MigrateSubscriptionRequest) GetStartingPhase() string`

GetStartingPhase returns the StartingPhase field if non-nil, zero value otherwise.

### GetStartingPhaseOk

`func (o *MigrateSubscriptionRequest) GetStartingPhaseOk() (*string, bool)`

GetStartingPhaseOk returns a tuple with the StartingPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPhase

`func (o *MigrateSubscriptionRequest) SetStartingPhase(v string)`

SetStartingPhase sets StartingPhase field to given value.

### HasStartingPhase

`func (o *MigrateSubscriptionRequest) HasStartingPhase() bool`

HasStartingPhase returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *MigrateSubscriptionRequest) GetBillingAnchor() time.Time`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *MigrateSubscriptionRequest) GetBillingAnchorOk() (*time.Time, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *MigrateSubscriptionRequest) SetBillingAnchor(v time.Time)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *MigrateSubscriptionRequest) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetCommercialOverrides

`func (o *MigrateSubscriptionRequest) GetCommercialOverrides() CommercialOverrides`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *MigrateSubscriptionRequest) GetCommercialOverridesOk() (*CommercialOverrides, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *MigrateSubscriptionRequest) SetCommercialOverrides(v CommercialOverrides)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *MigrateSubscriptionRequest) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


