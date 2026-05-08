# InvoiceLineSubscriptionReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**IDResource**](IDResource.md) | The subscription. | [readonly] 
**Phase** | [**IDResource**](IDResource.md) | The phase of the subscription. | [readonly] 
**Item** | [**IDResource**](IDResource.md) | The item this line is related to. | [readonly] 
**BillingPeriod** | [**Period**](Period.md) | The billing period of the subscription. In case the subscription item&#39;s billing period is different from the subscription&#39;s billing period, this field will contain the billing period of the subscription itself.  For example, in case of: - A monthly billed subscription anchored to 2025-01-01 - A subscription item billed daily  An example line would have the period of 2025-01-02 to 2025-01-03 as the item is billed daily, but the subscription&#39;s billing period will be 2025-01-01 to 2025-01-31. | [readonly] 

## Methods

### NewInvoiceLineSubscriptionReference

`func NewInvoiceLineSubscriptionReference(subscription IDResource, phase IDResource, item IDResource, billingPeriod Period, ) *InvoiceLineSubscriptionReference`

NewInvoiceLineSubscriptionReference instantiates a new InvoiceLineSubscriptionReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineSubscriptionReferenceWithDefaults

`func NewInvoiceLineSubscriptionReferenceWithDefaults() *InvoiceLineSubscriptionReference`

NewInvoiceLineSubscriptionReferenceWithDefaults instantiates a new InvoiceLineSubscriptionReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *InvoiceLineSubscriptionReference) GetSubscription() IDResource`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *InvoiceLineSubscriptionReference) GetSubscriptionOk() (*IDResource, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *InvoiceLineSubscriptionReference) SetSubscription(v IDResource)`

SetSubscription sets Subscription field to given value.


### GetPhase

`func (o *InvoiceLineSubscriptionReference) GetPhase() IDResource`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *InvoiceLineSubscriptionReference) GetPhaseOk() (*IDResource, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *InvoiceLineSubscriptionReference) SetPhase(v IDResource)`

SetPhase sets Phase field to given value.


### GetItem

`func (o *InvoiceLineSubscriptionReference) GetItem() IDResource`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *InvoiceLineSubscriptionReference) GetItemOk() (*IDResource, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *InvoiceLineSubscriptionReference) SetItem(v IDResource)`

SetItem sets Item field to given value.


### GetBillingPeriod

`func (o *InvoiceLineSubscriptionReference) GetBillingPeriod() Period`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *InvoiceLineSubscriptionReference) GetBillingPeriodOk() (*Period, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *InvoiceLineSubscriptionReference) SetBillingPeriod(v Period)`

SetBillingPeriod sets BillingPeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


