# CustomSubscriptionChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timing** | [**SubscriptionTiming**](SubscriptionTiming.md) | Timing configuration for the change, when the change should take effect. For changing a subscription, the accepted values depend on the subscription configuration. | 
**BillingAnchor** | Pointer to **time.Time** | The billing anchor of the subscription. The provided date will be normalized according to the billing cadence to the nearest recurrence before start time. If not provided, the previous subscription billing anchor will be used. | [optional] 
**CommercialOverrides** | Pointer to [**CommercialOverrides**](CommercialOverrides.md) | Commercial terms for this subscription change. | [optional] 
**CustomPlan** | [**CustomPlanInput**](CustomPlanInput.md) | The custom plan description which defines the Subscription. | 

## Methods

### NewCustomSubscriptionChange

`func NewCustomSubscriptionChange(timing SubscriptionTiming, customPlan CustomPlanInput, ) *CustomSubscriptionChange`

NewCustomSubscriptionChange instantiates a new CustomSubscriptionChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomSubscriptionChangeWithDefaults

`func NewCustomSubscriptionChangeWithDefaults() *CustomSubscriptionChange`

NewCustomSubscriptionChangeWithDefaults instantiates a new CustomSubscriptionChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiming

`func (o *CustomSubscriptionChange) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *CustomSubscriptionChange) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *CustomSubscriptionChange) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.


### GetBillingAnchor

`func (o *CustomSubscriptionChange) GetBillingAnchor() time.Time`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *CustomSubscriptionChange) GetBillingAnchorOk() (*time.Time, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *CustomSubscriptionChange) SetBillingAnchor(v time.Time)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *CustomSubscriptionChange) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetCommercialOverrides

`func (o *CustomSubscriptionChange) GetCommercialOverrides() CommercialOverrides`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *CustomSubscriptionChange) GetCommercialOverridesOk() (*CommercialOverrides, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *CustomSubscriptionChange) SetCommercialOverrides(v CommercialOverrides)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *CustomSubscriptionChange) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetCustomPlan

`func (o *CustomSubscriptionChange) GetCustomPlan() CustomPlanInput`

GetCustomPlan returns the CustomPlan field if non-nil, zero value otherwise.

### GetCustomPlanOk

`func (o *CustomSubscriptionChange) GetCustomPlanOk() (*CustomPlanInput, bool)`

GetCustomPlanOk returns a tuple with the CustomPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPlan

`func (o *CustomSubscriptionChange) SetCustomPlan(v CustomPlanInput)`

SetCustomPlan sets CustomPlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


