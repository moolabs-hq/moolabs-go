# PlanPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A semi-unique identifier for the resource. | 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Duration** | **string** | The duration of the phase. | 
**RateCards** | [**[]RateCard**](RateCard.md) | The rate cards of the plan. | 

## Methods

### NewPlanPhase

`func NewPlanPhase(key string, name string, duration string, rateCards []RateCard, ) *PlanPhase`

NewPlanPhase instantiates a new PlanPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPhaseWithDefaults

`func NewPlanPhaseWithDefaults() *PlanPhase`

NewPlanPhaseWithDefaults instantiates a new PlanPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlanPhase) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlanPhase) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlanPhase) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *PlanPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanPhase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlanPhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanPhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanPhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanPhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *PlanPhase) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanPhase) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanPhase) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanPhase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDuration

`func (o *PlanPhase) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PlanPhase) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PlanPhase) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetRateCards

`func (o *PlanPhase) GetRateCards() []RateCard`

GetRateCards returns the RateCards field if non-nil, zero value otherwise.

### GetRateCardsOk

`func (o *PlanPhase) GetRateCardsOk() (*[]RateCard, bool)`

GetRateCardsOk returns a tuple with the RateCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCards

`func (o *PlanPhase) SetRateCards(v []RateCard)`

SetRateCards sets RateCards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


