# SubscriptionAlignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillablesMustAlign** | Pointer to **bool** | Whether all Billable items and RateCards must align. Alignment means the Price&#39;s BillingCadence must align for both duration and anchor time. | [optional] 
**CurrentAlignedBillingPeriod** | Pointer to [**Period**](Period.md) | The current billing period. Only has value if the subscription is aligned and active. | [optional] 

## Methods

### NewSubscriptionAlignment

`func NewSubscriptionAlignment() *SubscriptionAlignment`

NewSubscriptionAlignment instantiates a new SubscriptionAlignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAlignmentWithDefaults

`func NewSubscriptionAlignmentWithDefaults() *SubscriptionAlignment`

NewSubscriptionAlignmentWithDefaults instantiates a new SubscriptionAlignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillablesMustAlign

`func (o *SubscriptionAlignment) GetBillablesMustAlign() bool`

GetBillablesMustAlign returns the BillablesMustAlign field if non-nil, zero value otherwise.

### GetBillablesMustAlignOk

`func (o *SubscriptionAlignment) GetBillablesMustAlignOk() (*bool, bool)`

GetBillablesMustAlignOk returns a tuple with the BillablesMustAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillablesMustAlign

`func (o *SubscriptionAlignment) SetBillablesMustAlign(v bool)`

SetBillablesMustAlign sets BillablesMustAlign field to given value.

### HasBillablesMustAlign

`func (o *SubscriptionAlignment) HasBillablesMustAlign() bool`

HasBillablesMustAlign returns a boolean if a field has been set.

### GetCurrentAlignedBillingPeriod

`func (o *SubscriptionAlignment) GetCurrentAlignedBillingPeriod() Period`

GetCurrentAlignedBillingPeriod returns the CurrentAlignedBillingPeriod field if non-nil, zero value otherwise.

### GetCurrentAlignedBillingPeriodOk

`func (o *SubscriptionAlignment) GetCurrentAlignedBillingPeriodOk() (*Period, bool)`

GetCurrentAlignedBillingPeriodOk returns a tuple with the CurrentAlignedBillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAlignedBillingPeriod

`func (o *SubscriptionAlignment) SetCurrentAlignedBillingPeriod(v Period)`

SetCurrentAlignedBillingPeriod sets CurrentAlignedBillingPeriod field to given value.

### HasCurrentAlignedBillingPeriod

`func (o *SubscriptionAlignment) HasCurrentAlignedBillingPeriod() bool`

HasCurrentAlignedBillingPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


