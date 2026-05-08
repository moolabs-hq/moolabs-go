# PlanSubscriptionChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timing** | [**SubscriptionTiming**](SubscriptionTiming.md) | Timing configuration for the change, when the change should take effect. For changing a subscription, the accepted values depend on the subscription configuration. | 
**Alignment** | Pointer to [**Alignment**](Alignment.md) | What alignment settings the subscription should have. | [optional] 
**Metadata** | Pointer to **map[string]string** | Arbitrary metadata associated with the subscription. | [optional] 
**CommercialOverrides** | Pointer to [**CommercialOverrides**](CommercialOverrides.md) | Commercial terms for this subscription (replaces prior overrides when set on change). | [optional] 
**Plan** | [**PlanReferenceInput**](PlanReferenceInput.md) | The plan reference to change to. | 
**StartingPhase** | Pointer to **string** | The key of the phase to start the subscription in. If not provided, the subscription will start in the first phase of the plan. | [optional] 
**Name** | Pointer to **string** | The name of the Subscription. If not provided the plan name is used. | [optional] 
**Description** | Pointer to **string** | Description for the Subscription. | [optional] 
**BillingAnchor** | Pointer to **time.Time** | The billing anchor of the subscription. The provided date will be normalized according to the billing cadence to the nearest recurrence before start time. If not provided, the previous subscription billing anchor will be used. | [optional] 

## Methods

### NewPlanSubscriptionChange

`func NewPlanSubscriptionChange(timing SubscriptionTiming, plan PlanReferenceInput, ) *PlanSubscriptionChange`

NewPlanSubscriptionChange instantiates a new PlanSubscriptionChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanSubscriptionChangeWithDefaults

`func NewPlanSubscriptionChangeWithDefaults() *PlanSubscriptionChange`

NewPlanSubscriptionChangeWithDefaults instantiates a new PlanSubscriptionChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiming

`func (o *PlanSubscriptionChange) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *PlanSubscriptionChange) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *PlanSubscriptionChange) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.


### GetAlignment

`func (o *PlanSubscriptionChange) GetAlignment() Alignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *PlanSubscriptionChange) GetAlignmentOk() (*Alignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *PlanSubscriptionChange) SetAlignment(v Alignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *PlanSubscriptionChange) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetMetadata

`func (o *PlanSubscriptionChange) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanSubscriptionChange) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanSubscriptionChange) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanSubscriptionChange) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCommercialOverrides

`func (o *PlanSubscriptionChange) GetCommercialOverrides() CommercialOverrides`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *PlanSubscriptionChange) GetCommercialOverridesOk() (*CommercialOverrides, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *PlanSubscriptionChange) SetCommercialOverrides(v CommercialOverrides)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *PlanSubscriptionChange) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetPlan

`func (o *PlanSubscriptionChange) GetPlan() PlanReferenceInput`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PlanSubscriptionChange) GetPlanOk() (*PlanReferenceInput, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PlanSubscriptionChange) SetPlan(v PlanReferenceInput)`

SetPlan sets Plan field to given value.


### GetStartingPhase

`func (o *PlanSubscriptionChange) GetStartingPhase() string`

GetStartingPhase returns the StartingPhase field if non-nil, zero value otherwise.

### GetStartingPhaseOk

`func (o *PlanSubscriptionChange) GetStartingPhaseOk() (*string, bool)`

GetStartingPhaseOk returns a tuple with the StartingPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPhase

`func (o *PlanSubscriptionChange) SetStartingPhase(v string)`

SetStartingPhase sets StartingPhase field to given value.

### HasStartingPhase

`func (o *PlanSubscriptionChange) HasStartingPhase() bool`

HasStartingPhase returns a boolean if a field has been set.

### GetName

`func (o *PlanSubscriptionChange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanSubscriptionChange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanSubscriptionChange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanSubscriptionChange) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PlanSubscriptionChange) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanSubscriptionChange) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanSubscriptionChange) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanSubscriptionChange) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *PlanSubscriptionChange) GetBillingAnchor() time.Time`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *PlanSubscriptionChange) GetBillingAnchorOk() (*time.Time, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *PlanSubscriptionChange) SetBillingAnchor(v time.Time)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *PlanSubscriptionChange) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


