# SubscriptionPhaseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAfter** | **string** | Interval after the subscription starts to transition to the phase. When null, the phase starts immediately after the subscription starts. | 
**Duration** | Pointer to **string** | The intended duration of the new phase. Duration is required when the phase will not be the last phase. | [optional] 
**Discounts** | Pointer to [**Discounts**](Discounts.md) | The discounts on the plan. | [optional] 
**Key** | **string** | A locally unique identifier for the phase. | 
**Name** | **string** | The name of the phase. | 
**Description** | Pointer to **string** | The description of the phase. | [optional] 

## Methods

### NewSubscriptionPhaseCreate

`func NewSubscriptionPhaseCreate(startAfter string, key string, name string, ) *SubscriptionPhaseCreate`

NewSubscriptionPhaseCreate instantiates a new SubscriptionPhaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPhaseCreateWithDefaults

`func NewSubscriptionPhaseCreateWithDefaults() *SubscriptionPhaseCreate`

NewSubscriptionPhaseCreateWithDefaults instantiates a new SubscriptionPhaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAfter

`func (o *SubscriptionPhaseCreate) GetStartAfter() string`

GetStartAfter returns the StartAfter field if non-nil, zero value otherwise.

### GetStartAfterOk

`func (o *SubscriptionPhaseCreate) GetStartAfterOk() (*string, bool)`

GetStartAfterOk returns a tuple with the StartAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAfter

`func (o *SubscriptionPhaseCreate) SetStartAfter(v string)`

SetStartAfter sets StartAfter field to given value.


### GetDuration

`func (o *SubscriptionPhaseCreate) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SubscriptionPhaseCreate) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SubscriptionPhaseCreate) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SubscriptionPhaseCreate) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDiscounts

`func (o *SubscriptionPhaseCreate) GetDiscounts() Discounts`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *SubscriptionPhaseCreate) GetDiscountsOk() (*Discounts, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *SubscriptionPhaseCreate) SetDiscounts(v Discounts)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *SubscriptionPhaseCreate) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetKey

`func (o *SubscriptionPhaseCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SubscriptionPhaseCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SubscriptionPhaseCreate) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SubscriptionPhaseCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionPhaseCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionPhaseCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionPhaseCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionPhaseCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionPhaseCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionPhaseCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


