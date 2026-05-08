# BillingWorkflowCollectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to [**BillingWorkflowCollectionAlignment**](BillingWorkflowCollectionAlignment.md) | The alignment for collecting the pending line items into an invoice. | [optional] 
**Interval** | Pointer to **string** | This grace period can be used to delay the collection of the pending line items specified in alignment.  This is useful, in case of multiple subscriptions having slightly different billing periods. | [optional] [default to "PT1H"]

## Methods

### NewBillingWorkflowCollectionSettings

`func NewBillingWorkflowCollectionSettings() *BillingWorkflowCollectionSettings`

NewBillingWorkflowCollectionSettings instantiates a new BillingWorkflowCollectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingWorkflowCollectionSettingsWithDefaults

`func NewBillingWorkflowCollectionSettingsWithDefaults() *BillingWorkflowCollectionSettings`

NewBillingWorkflowCollectionSettingsWithDefaults instantiates a new BillingWorkflowCollectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *BillingWorkflowCollectionSettings) GetAlignment() BillingWorkflowCollectionAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *BillingWorkflowCollectionSettings) GetAlignmentOk() (*BillingWorkflowCollectionAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *BillingWorkflowCollectionSettings) SetAlignment(v BillingWorkflowCollectionAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *BillingWorkflowCollectionSettings) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetInterval

`func (o *BillingWorkflowCollectionSettings) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BillingWorkflowCollectionSettings) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BillingWorkflowCollectionSettings) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BillingWorkflowCollectionSettings) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


