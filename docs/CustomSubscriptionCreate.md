# CustomSubscriptionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommercialOverrides** | Pointer to [**CommercialOverrides**](CommercialOverrides.md) | Commercial terms for this subscription change. | [optional] 
**CustomPlan** | [**CustomPlanInput**](CustomPlanInput.md) | The custom plan description which defines the Subscription. | 
**Timing** | Pointer to [**SubscriptionTiming**](SubscriptionTiming.md) | Timing configuration for the change, when the change should take effect. The default is immediate. | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer. Provide either the key or ID. Has presedence over the key. | [optional] 
**CustomerKey** | Pointer to **string** | The key of the customer. Provide either the key or ID. | [optional] 
**BillingAnchor** | Pointer to **time.Time** | The billing anchor of the subscription. The provided date will be normalized according to the billing cadence to the nearest recurrence before start time. If not provided, the subscription start time will be used. | [optional] 

## Methods

### NewCustomSubscriptionCreate

`func NewCustomSubscriptionCreate(customPlan CustomPlanInput, ) *CustomSubscriptionCreate`

NewCustomSubscriptionCreate instantiates a new CustomSubscriptionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomSubscriptionCreateWithDefaults

`func NewCustomSubscriptionCreateWithDefaults() *CustomSubscriptionCreate`

NewCustomSubscriptionCreateWithDefaults instantiates a new CustomSubscriptionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommercialOverrides

`func (o *CustomSubscriptionCreate) GetCommercialOverrides() CommercialOverrides`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *CustomSubscriptionCreate) GetCommercialOverridesOk() (*CommercialOverrides, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *CustomSubscriptionCreate) SetCommercialOverrides(v CommercialOverrides)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *CustomSubscriptionCreate) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetCustomPlan

`func (o *CustomSubscriptionCreate) GetCustomPlan() CustomPlanInput`

GetCustomPlan returns the CustomPlan field if non-nil, zero value otherwise.

### GetCustomPlanOk

`func (o *CustomSubscriptionCreate) GetCustomPlanOk() (*CustomPlanInput, bool)`

GetCustomPlanOk returns a tuple with the CustomPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPlan

`func (o *CustomSubscriptionCreate) SetCustomPlan(v CustomPlanInput)`

SetCustomPlan sets CustomPlan field to given value.


### GetTiming

`func (o *CustomSubscriptionCreate) GetTiming() SubscriptionTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *CustomSubscriptionCreate) GetTimingOk() (*SubscriptionTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *CustomSubscriptionCreate) SetTiming(v SubscriptionTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *CustomSubscriptionCreate) HasTiming() bool`

HasTiming returns a boolean if a field has been set.

### GetCustomerId

`func (o *CustomSubscriptionCreate) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomSubscriptionCreate) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomSubscriptionCreate) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomSubscriptionCreate) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerKey

`func (o *CustomSubscriptionCreate) GetCustomerKey() string`

GetCustomerKey returns the CustomerKey field if non-nil, zero value otherwise.

### GetCustomerKeyOk

`func (o *CustomSubscriptionCreate) GetCustomerKeyOk() (*string, bool)`

GetCustomerKeyOk returns a tuple with the CustomerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerKey

`func (o *CustomSubscriptionCreate) SetCustomerKey(v string)`

SetCustomerKey sets CustomerKey field to given value.

### HasCustomerKey

`func (o *CustomSubscriptionCreate) HasCustomerKey() bool`

HasCustomerKey returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *CustomSubscriptionCreate) GetBillingAnchor() time.Time`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *CustomSubscriptionCreate) GetBillingAnchorOk() (*time.Time, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *CustomSubscriptionCreate) SetBillingAnchor(v time.Time)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *CustomSubscriptionCreate) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


