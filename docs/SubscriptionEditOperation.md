# SubscriptionEditOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**PhaseKey** | **string** |  | 
**RateCard** | [**RateCard**](RateCard.md) |  | 
**ItemKey** | **string** |  | 
**Phase** | [**SubscriptionPhaseCreate**](SubscriptionPhaseCreate.md) |  | 
**Shift** | [**RemovePhaseShifting**](RemovePhaseShifting.md) |  | 
**ExtendBy** | **string** |  | 

## Methods

### NewSubscriptionEditOperation

`func NewSubscriptionEditOperation(op string, phaseKey string, rateCard RateCard, itemKey string, phase SubscriptionPhaseCreate, shift RemovePhaseShifting, extendBy string, ) *SubscriptionEditOperation`

NewSubscriptionEditOperation instantiates a new SubscriptionEditOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEditOperationWithDefaults

`func NewSubscriptionEditOperationWithDefaults() *SubscriptionEditOperation`

NewSubscriptionEditOperationWithDefaults instantiates a new SubscriptionEditOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *SubscriptionEditOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SubscriptionEditOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SubscriptionEditOperation) SetOp(v string)`

SetOp sets Op field to given value.


### GetPhaseKey

`func (o *SubscriptionEditOperation) GetPhaseKey() string`

GetPhaseKey returns the PhaseKey field if non-nil, zero value otherwise.

### GetPhaseKeyOk

`func (o *SubscriptionEditOperation) GetPhaseKeyOk() (*string, bool)`

GetPhaseKeyOk returns a tuple with the PhaseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseKey

`func (o *SubscriptionEditOperation) SetPhaseKey(v string)`

SetPhaseKey sets PhaseKey field to given value.


### GetRateCard

`func (o *SubscriptionEditOperation) GetRateCard() RateCard`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *SubscriptionEditOperation) GetRateCardOk() (*RateCard, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *SubscriptionEditOperation) SetRateCard(v RateCard)`

SetRateCard sets RateCard field to given value.


### GetItemKey

`func (o *SubscriptionEditOperation) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *SubscriptionEditOperation) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *SubscriptionEditOperation) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.


### GetPhase

`func (o *SubscriptionEditOperation) GetPhase() SubscriptionPhaseCreate`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SubscriptionEditOperation) GetPhaseOk() (*SubscriptionPhaseCreate, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SubscriptionEditOperation) SetPhase(v SubscriptionPhaseCreate)`

SetPhase sets Phase field to given value.


### GetShift

`func (o *SubscriptionEditOperation) GetShift() RemovePhaseShifting`

GetShift returns the Shift field if non-nil, zero value otherwise.

### GetShiftOk

`func (o *SubscriptionEditOperation) GetShiftOk() (*RemovePhaseShifting, bool)`

GetShiftOk returns a tuple with the Shift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShift

`func (o *SubscriptionEditOperation) SetShift(v RemovePhaseShifting)`

SetShift sets Shift field to given value.


### GetExtendBy

`func (o *SubscriptionEditOperation) GetExtendBy() string`

GetExtendBy returns the ExtendBy field if non-nil, zero value otherwise.

### GetExtendByOk

`func (o *SubscriptionEditOperation) GetExtendByOk() (*string, bool)`

GetExtendByOk returns a tuple with the ExtendBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendBy

`func (o *SubscriptionEditOperation) SetExtendBy(v string)`

SetExtendBy sets ExtendBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


