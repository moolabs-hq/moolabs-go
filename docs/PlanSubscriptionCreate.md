# PlanSubscriptionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to [**Alignment**](Alignment.md) | What alignment settings the subscription should have. | [optional] 
**Metadata** | Pointer to **map[string]string** | Arbitrary metadata associated with the subscription. | [optional] 
**CommercialOverrides** | Pointer to [**CommercialOverrides**](CommercialOverrides.md) | Commercial terms for this subscription (replaces prior overrides when set on change). | [optional] 
**Plan** | [**PlanReferenceInput**](PlanReferenceInput.md) | The plan reference to change to. | 
**StartingPhase** | Pointer to **string** | The key of the phase to start the subscription in. If not provided, the subscription will start in the first phase of the plan. | [optional] 
**Name** | Pointer to **string** | The name of the Subscription. If not provided the plan name is used. | [optional] 
**Description** | Pointer to **string** | Description for the Subscription. | [optional] 
**Timing** | Pointer to [**SubscriptionTiming**](SubscriptionTiming.md) | Timing configuration for the change, when the change should take effect. The default is immediate. | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer. Provide either the key or ID. Has presedence over the key. | [optional] 
**CustomerKey** | Pointer to **string** | The key of the customer. Provide either the key or ID. | [optional] 
**BillingAnchor** | Pointer to **time.Time** | The billing anchor of the subscription. The provided date will be normalized according to the billing cadence to the nearest recurrence before start time. If not provided, the subscription start time will be used. | [optional] 

## Methods

### NewPlanSubscriptionCreate

`func NewPlanSubscriptionCreate(plan PlanReferenceInput, ) *PlanSubscriptionCreate`

NewPlanSubscriptionCreate instantiates a new PlanSubscriptionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanSubscriptionCreateWithDefaults

`func NewPlanSubscriptionCreateWithDefaults() *PlanSubscriptionCreate`

NewPlanSubscriptionCreateWithDefaults instantiates a new PlanSubscriptionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *PlanSubscriptionCreate) GetAlignment() Alignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *PlanSubscriptionCreate) GetAlignmentOk() (*Alignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *PlanSubscriptionCreate) SetAlignment(v Alignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *PlanSubscriptionCreate) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetMetadata

`func (o *PlanSubscriptionCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanSubscriptionCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanSubscriptionCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanSubscriptionCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCommercialOverrides

`func (o *PlanSubscriptionCreate) GetCommercialOverrides() CommercialOverrides`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *PlanSubscriptionCreate) GetCommercialOverridesOk() (*CommercialOverrides, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *PlanSubscriptionCreate) SetCommercialOverrides(v CommercialOverrides)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *PlanSubscriptionCreate) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetPlan

`func (o *PlanSubscriptionCreate) GetPlan() PlanReferenceInput`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PlanSubscriptionCreate) GetPlanOk() (*PlanReferenceInput, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PlanSubscriptionCreate) SetPlan(v PlanReferenceInput)`

SetPlan sets Plan field to given value.


### GetStartingPhase

`func (o *PlanSubscriptionCreate) GetStartingPhase() string`

GetStartingPhase returns the StartingPhase field if non-nil, zero value otherwise.

### GetStartingPhaseOk

`func (o *PlanSubscriptionCreate) GetStartingPhaseOk() (*string, bool)`

GetStartingPhaseOk returns a tuple with the StartingPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPhase

`func (o *PlanSubscriptionCreate) SetStartingPhase(v string)`

SetStartingPhase sets StartingPhase field to given value.

### HasStartingPhase

`func (o *PlanSubscriptionCreate) HasStartingPhase() bool`

HasStartingPhase returns a boolean if a field has been set.

### GetName

`func (o *PlanSubscriptionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanSubscriptionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanSubscriptionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanSubscriptionCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PlanSubscriptionCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanSubscriptionCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanSubscriptionCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanSubscriptionCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTiming

`func (o *PlanSubscriptionCreate) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *PlanSubscriptionCreate) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *PlanSubscriptionCreate) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *PlanSubscriptionCreate) HasTiming() bool`

HasTiming returns a boolean if a field has been set.

### GetCustomerId

`func (o *PlanSubscriptionCreate) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PlanSubscriptionCreate) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PlanSubscriptionCreate) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PlanSubscriptionCreate) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerKey

`func (o *PlanSubscriptionCreate) GetCustomerKey() string`

GetCustomerKey returns the CustomerKey field if non-nil, zero value otherwise.

### GetCustomerKeyOk

`func (o *PlanSubscriptionCreate) GetCustomerKeyOk() (*string, bool)`

GetCustomerKeyOk returns a tuple with the CustomerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerKey

`func (o *PlanSubscriptionCreate) SetCustomerKey(v string)`

SetCustomerKey sets CustomerKey field to given value.

### HasCustomerKey

`func (o *PlanSubscriptionCreate) HasCustomerKey() bool`

HasCustomerKey returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *PlanSubscriptionCreate) GetBillingAnchor() time.Time`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *PlanSubscriptionCreate) GetBillingAnchorOk() (*time.Time, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *PlanSubscriptionCreate) SetBillingAnchor(v time.Time)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *PlanSubscriptionCreate) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


