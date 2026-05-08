# EditSubscriptionRemovePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**PhaseKey** | **string** |  | 
**Shift** | [**RemovePhaseShifting**](RemovePhaseShifting.md) |  | 

## Methods

### NewEditSubscriptionRemovePhase

`func NewEditSubscriptionRemovePhase(op string, phaseKey string, shift RemovePhaseShifting, ) *EditSubscriptionRemovePhase`

NewEditSubscriptionRemovePhase instantiates a new EditSubscriptionRemovePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSubscriptionRemovePhaseWithDefaults

`func NewEditSubscriptionRemovePhaseWithDefaults() *EditSubscriptionRemovePhase`

NewEditSubscriptionRemovePhaseWithDefaults instantiates a new EditSubscriptionRemovePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *EditSubscriptionRemovePhase) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EditSubscriptionRemovePhase) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EditSubscriptionRemovePhase) SetOp(v string)`

SetOp sets Op field to given value.


### GetPhaseKey

`func (o *EditSubscriptionRemovePhase) GetPhaseKey() string`

GetPhaseKey returns the PhaseKey field if non-nil, zero value otherwise.

### GetPhaseKeyOk

`func (o *EditSubscriptionRemovePhase) GetPhaseKeyOk() (*string, bool)`

GetPhaseKeyOk returns a tuple with the PhaseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseKey

`func (o *EditSubscriptionRemovePhase) SetPhaseKey(v string)`

SetPhaseKey sets PhaseKey field to given value.


### GetShift

`func (o *EditSubscriptionRemovePhase) GetShift() RemovePhaseShifting`

GetShift returns the Shift field if non-nil, zero value otherwise.

### GetShiftOk

`func (o *EditSubscriptionRemovePhase) GetShiftOk() (*RemovePhaseShifting, bool)`

GetShiftOk returns a tuple with the Shift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShift

`func (o *EditSubscriptionRemovePhase) SetShift(v RemovePhaseShifting)`

SetShift sets Shift field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


