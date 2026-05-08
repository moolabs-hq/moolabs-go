# Alignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillablesMustAlign** | Pointer to **bool** | Whether all Billable items and RateCards must align. Alignment means the Price&#39;s BillingCadence must align for both duration and anchor time. | [optional] 

## Methods

### NewAlignment

`func NewAlignment() *Alignment`

NewAlignment instantiates a new Alignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlignmentWithDefaults

`func NewAlignmentWithDefaults() *Alignment`

NewAlignmentWithDefaults instantiates a new Alignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillablesMustAlign

`func (o *Alignment) GetBillablesMustAlign() bool`

GetBillablesMustAlign returns the BillablesMustAlign field if non-nil, zero value otherwise.

### GetBillablesMustAlignOk

`func (o *Alignment) GetBillablesMustAlignOk() (*bool, bool)`

GetBillablesMustAlignOk returns a tuple with the BillablesMustAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillablesMustAlign

`func (o *Alignment) SetBillablesMustAlign(v bool)`

SetBillablesMustAlign sets BillablesMustAlign field to given value.

### HasBillablesMustAlign

`func (o *Alignment) HasBillablesMustAlign() bool`

HasBillablesMustAlign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


