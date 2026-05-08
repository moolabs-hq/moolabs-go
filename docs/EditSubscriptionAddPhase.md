# EditSubscriptionAddPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Phase** | [**SubscriptionPhaseCreate**](SubscriptionPhaseCreate.md) |  | 

## Methods

### NewEditSubscriptionAddPhase

`func NewEditSubscriptionAddPhase(op string, phase SubscriptionPhaseCreate, ) *EditSubscriptionAddPhase`

NewEditSubscriptionAddPhase instantiates a new EditSubscriptionAddPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSubscriptionAddPhaseWithDefaults

`func NewEditSubscriptionAddPhaseWithDefaults() *EditSubscriptionAddPhase`

NewEditSubscriptionAddPhaseWithDefaults instantiates a new EditSubscriptionAddPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *EditSubscriptionAddPhase) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *EditSubscriptionAddPhase) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *EditSubscriptionAddPhase) SetOp(v string)`

SetOp sets Op field to given value.


### GetPhase

`func (o *EditSubscriptionAddPhase) GetPhase() SubscriptionPhaseCreate`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *EditSubscriptionAddPhase) GetPhaseOk() (*SubscriptionPhaseCreate, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *EditSubscriptionAddPhase) SetPhase(v SubscriptionPhaseCreate)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


