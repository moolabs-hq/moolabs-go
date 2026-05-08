# FeaturePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** | Feature key this pool applies to. | 
**Enabled** | **bool** |  | 
**InitialAmount** | **float64** |  | 
**RecurringAmount** | **float64** |  | 
**Cadence** | Pointer to **string** | Optional cadence override for this pool (e.g. same as plan billing when empty). | [optional] 
**ExpiresInDays** | **int32** |  | 
**RolloverEnabled** | **bool** |  | 

## Methods

### NewFeaturePool

`func NewFeaturePool(featureKey string, enabled bool, initialAmount float64, recurringAmount float64, expiresInDays int32, rolloverEnabled bool, ) *FeaturePool`

NewFeaturePool instantiates a new FeaturePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturePoolWithDefaults

`func NewFeaturePoolWithDefaults() *FeaturePool`

NewFeaturePoolWithDefaults instantiates a new FeaturePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *FeaturePool) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *FeaturePool) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *FeaturePool) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetEnabled

`func (o *FeaturePool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeaturePool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeaturePool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInitialAmount

`func (o *FeaturePool) GetInitialAmount() float64`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *FeaturePool) GetInitialAmountOk() (*float64, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *FeaturePool) SetInitialAmount(v float64)`

SetInitialAmount sets InitialAmount field to given value.


### GetRecurringAmount

`func (o *FeaturePool) GetRecurringAmount() float64`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *FeaturePool) GetRecurringAmountOk() (*float64, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *FeaturePool) SetRecurringAmount(v float64)`

SetRecurringAmount sets RecurringAmount field to given value.


### GetCadence

`func (o *FeaturePool) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *FeaturePool) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *FeaturePool) SetCadence(v string)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *FeaturePool) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetExpiresInDays

`func (o *FeaturePool) GetExpiresInDays() int32`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *FeaturePool) GetExpiresInDaysOk() (*int32, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *FeaturePool) SetExpiresInDays(v int32)`

SetExpiresInDays sets ExpiresInDays field to given value.


### GetRolloverEnabled

`func (o *FeaturePool) GetRolloverEnabled() bool`

GetRolloverEnabled returns the RolloverEnabled field if non-nil, zero value otherwise.

### GetRolloverEnabledOk

`func (o *FeaturePool) GetRolloverEnabledOk() (*bool, bool)`

GetRolloverEnabledOk returns a tuple with the RolloverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverEnabled

`func (o *FeaturePool) SetRolloverEnabled(v bool)`

SetRolloverEnabled sets RolloverEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


