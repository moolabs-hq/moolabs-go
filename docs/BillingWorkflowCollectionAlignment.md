# BillingWorkflowCollectionAlignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of alignment. | 
**RecurringPeriod** | [**RecurringPeriodV2**](RecurringPeriodV2.md) | The recurring period for the alignment. | 

## Methods

### NewBillingWorkflowCollectionAlignment

`func NewBillingWorkflowCollectionAlignment(type_ string, recurringPeriod RecurringPeriodV2, ) *BillingWorkflowCollectionAlignment`

NewBillingWorkflowCollectionAlignment instantiates a new BillingWorkflowCollectionAlignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWorkflowCollectionAlignmentWithDefaults

`func NewBillingWorkflowCollectionAlignmentWithDefaults() *BillingWorkflowCollectionAlignment`

NewBillingWorkflowCollectionAlignmentWithDefaults instantiates a new BillingWorkflowCollectionAlignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillingWorkflowCollectionAlignment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingWorkflowCollectionAlignment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingWorkflowCollectionAlignment) SetType(v string)`

SetType sets Type field to given value.


### GetRecurringPeriod

`func (o *BillingWorkflowCollectionAlignment) GetRecurringPeriod() RecurringPeriodV2`

GetRecurringPeriod returns the RecurringPeriod field if non-nil, zero value otherwise.

### GetRecurringPeriodOk

`func (o *BillingWorkflowCollectionAlignment) GetRecurringPeriodOk() (*RecurringPeriodV2, bool)`

GetRecurringPeriodOk returns a tuple with the RecurringPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPeriod

`func (o *BillingWorkflowCollectionAlignment) SetRecurringPeriod(v RecurringPeriodV2)`

SetRecurringPeriod sets RecurringPeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


