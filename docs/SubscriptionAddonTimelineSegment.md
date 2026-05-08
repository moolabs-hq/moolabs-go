# SubscriptionAddonTimelineSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The cadence start of the resource. | 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] 
**Quantity** | **int32** | The quantity of the add-on for the given period. | [readonly] 

## Methods

### NewSubscriptionAddonTimelineSegment

`func NewSubscriptionAddonTimelineSegment(activeFrom time.Time, quantity int32, ) *SubscriptionAddonTimelineSegment`

NewSubscriptionAddonTimelineSegment instantiates a new SubscriptionAddonTimelineSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAddonTimelineSegmentWithDefaults

`func NewSubscriptionAddonTimelineSegmentWithDefaults() *SubscriptionAddonTimelineSegment`

NewSubscriptionAddonTimelineSegmentWithDefaults instantiates a new SubscriptionAddonTimelineSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *SubscriptionAddonTimelineSegment) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionAddonTimelineSegment) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionAddonTimelineSegment) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *SubscriptionAddonTimelineSegment) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionAddonTimelineSegment) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionAddonTimelineSegment) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionAddonTimelineSegment) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionAddonTimelineSegment) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionAddonTimelineSegment) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionAddonTimelineSegment) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


