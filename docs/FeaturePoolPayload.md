# FeaturePoolPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** |  | 
**Enabled** | **bool** |  | 
**InitialAmount** | [**Initialamount**](Initialamount.md) |  | 
**RecurringAmount** | [**Recurringamount**](Recurringamount.md) |  | 
**Cadence** | Pointer to **string** |  | [optional] 
**ExpiresInDays** | **int32** |  | 
**RolloverEnabled** | **bool** |  | 

## Methods

### NewFeaturePoolPayload

`func NewFeaturePoolPayload(featureKey string, enabled bool, initialAmount Initialamount, recurringAmount Recurringamount, expiresInDays int32, rolloverEnabled bool, ) *FeaturePoolPayload`

NewFeaturePoolPayload instantiates a new FeaturePoolPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturePoolPayloadWithDefaults

`func NewFeaturePoolPayloadWithDefaults() *FeaturePoolPayload`

NewFeaturePoolPayloadWithDefaults instantiates a new FeaturePoolPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *FeaturePoolPayload) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *FeaturePoolPayload) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *FeaturePoolPayload) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetEnabled

`func (o *FeaturePoolPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeaturePoolPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeaturePoolPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInitialAmount

`func (o *FeaturePoolPayload) GetInitialAmount() Initialamount`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *FeaturePoolPayload) GetInitialAmountOk() (*Initialamount, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *FeaturePoolPayload) SetInitialAmount(v Initialamount)`

SetInitialAmount sets InitialAmount field to given value.


### GetRecurringAmount

`func (o *FeaturePoolPayload) GetRecurringAmount() Recurringamount`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *FeaturePoolPayload) GetRecurringAmountOk() (*Recurringamount, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *FeaturePoolPayload) SetRecurringAmount(v Recurringamount)`

SetRecurringAmount sets RecurringAmount field to given value.


### GetCadence

`func (o *FeaturePoolPayload) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *FeaturePoolPayload) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *FeaturePoolPayload) SetCadence(v string)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *FeaturePoolPayload) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetExpiresInDays

`func (o *FeaturePoolPayload) GetExpiresInDays() int32`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *FeaturePoolPayload) GetExpiresInDaysOk() (*int32, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *FeaturePoolPayload) SetExpiresInDays(v int32)`

SetExpiresInDays sets ExpiresInDays field to given value.


### GetRolloverEnabled

`func (o *FeaturePoolPayload) GetRolloverEnabled() bool`

GetRolloverEnabled returns the RolloverEnabled field if non-nil, zero value otherwise.

### GetRolloverEnabledOk

`func (o *FeaturePoolPayload) GetRolloverEnabledOk() (*bool, bool)`

GetRolloverEnabledOk returns a tuple with the RolloverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverEnabled

`func (o *FeaturePoolPayload) SetRolloverEnabled(v bool)`

SetRolloverEnabled sets RolloverEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


