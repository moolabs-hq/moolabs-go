# CreditConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**GlobalPoolEnabled** | **bool** |  | 
**FeaturePoolsEnabled** | **bool** |  | 
**InitialAmount** | **float64** |  | 
**RecurringAmount** | **float64** |  | 
**Cadence** | **string** | Billing cadence key for credit allowance (e.g. MONTHLY), aligned with plan usage. | 
**ExpiresInDays** | **int32** |  | 
**RolloverEnabled** | **bool** |  | 
**RolloverPercent** | Pointer to **float64** |  | [optional] 
**RolloverCapAmount** | Pointer to **float64** |  | [optional] 
**RolloverExpiresAfterDays** | Pointer to **int32** |  | [optional] 
**RolloverPriorityDelta** | Pointer to **int32** |  | [optional] 
**FeaturePools** | Pointer to [**[]FeaturePool**](FeaturePool.md) |  | [optional] 

## Methods

### NewCreditConfig

`func NewCreditConfig(enabled bool, globalPoolEnabled bool, featurePoolsEnabled bool, initialAmount float64, recurringAmount float64, cadence string, expiresInDays int32, rolloverEnabled bool, ) *CreditConfig`

NewCreditConfig instantiates a new CreditConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditConfigWithDefaults

`func NewCreditConfigWithDefaults() *CreditConfig`

NewCreditConfigWithDefaults instantiates a new CreditConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreditConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreditConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreditConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGlobalPoolEnabled

`func (o *CreditConfig) GetGlobalPoolEnabled() bool`

GetGlobalPoolEnabled returns the GlobalPoolEnabled field if non-nil, zero value otherwise.

### GetGlobalPoolEnabledOk

`func (o *CreditConfig) GetGlobalPoolEnabledOk() (*bool, bool)`

GetGlobalPoolEnabledOk returns a tuple with the GlobalPoolEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPoolEnabled

`func (o *CreditConfig) SetGlobalPoolEnabled(v bool)`

SetGlobalPoolEnabled sets GlobalPoolEnabled field to given value.


### GetFeaturePoolsEnabled

`func (o *CreditConfig) GetFeaturePoolsEnabled() bool`

GetFeaturePoolsEnabled returns the FeaturePoolsEnabled field if non-nil, zero value otherwise.

### GetFeaturePoolsEnabledOk

`func (o *CreditConfig) GetFeaturePoolsEnabledOk() (*bool, bool)`

GetFeaturePoolsEnabledOk returns a tuple with the FeaturePoolsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePoolsEnabled

`func (o *CreditConfig) SetFeaturePoolsEnabled(v bool)`

SetFeaturePoolsEnabled sets FeaturePoolsEnabled field to given value.


### GetInitialAmount

`func (o *CreditConfig) GetInitialAmount() float64`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *CreditConfig) GetInitialAmountOk() (*float64, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *CreditConfig) SetInitialAmount(v float64)`

SetInitialAmount sets InitialAmount field to given value.


### GetRecurringAmount

`func (o *CreditConfig) GetRecurringAmount() float64`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *CreditConfig) GetRecurringAmountOk() (*float64, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *CreditConfig) SetRecurringAmount(v float64)`

SetRecurringAmount sets RecurringAmount field to given value.


### GetCadence

`func (o *CreditConfig) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *CreditConfig) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *CreditConfig) SetCadence(v string)`

SetCadence sets Cadence field to given value.


### GetExpiresInDays

`func (o *CreditConfig) GetExpiresInDays() int32`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *CreditConfig) GetExpiresInDaysOk() (*int32, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *CreditConfig) SetExpiresInDays(v int32)`

SetExpiresInDays sets ExpiresInDays field to given value.


### GetRolloverEnabled

`func (o *CreditConfig) GetRolloverEnabled() bool`

GetRolloverEnabled returns the RolloverEnabled field if non-nil, zero value otherwise.

### GetRolloverEnabledOk

`func (o *CreditConfig) GetRolloverEnabledOk() (*bool, bool)`

GetRolloverEnabledOk returns a tuple with the RolloverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverEnabled

`func (o *CreditConfig) SetRolloverEnabled(v bool)`

SetRolloverEnabled sets RolloverEnabled field to given value.


### GetRolloverPercent

`func (o *CreditConfig) GetRolloverPercent() float64`

GetRolloverPercent returns the RolloverPercent field if non-nil, zero value otherwise.

### GetRolloverPercentOk

`func (o *CreditConfig) GetRolloverPercentOk() (*float64, bool)`

GetRolloverPercentOk returns a tuple with the RolloverPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPercent

`func (o *CreditConfig) SetRolloverPercent(v float64)`

SetRolloverPercent sets RolloverPercent field to given value.

### HasRolloverPercent

`func (o *CreditConfig) HasRolloverPercent() bool`

HasRolloverPercent returns a boolean if a field has been set.

### GetRolloverCapAmount

`func (o *CreditConfig) GetRolloverCapAmount() float64`

GetRolloverCapAmount returns the RolloverCapAmount field if non-nil, zero value otherwise.

### GetRolloverCapAmountOk

`func (o *CreditConfig) GetRolloverCapAmountOk() (*float64, bool)`

GetRolloverCapAmountOk returns a tuple with the RolloverCapAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverCapAmount

`func (o *CreditConfig) SetRolloverCapAmount(v float64)`

SetRolloverCapAmount sets RolloverCapAmount field to given value.

### HasRolloverCapAmount

`func (o *CreditConfig) HasRolloverCapAmount() bool`

HasRolloverCapAmount returns a boolean if a field has been set.

### GetRolloverExpiresAfterDays

`func (o *CreditConfig) GetRolloverExpiresAfterDays() int32`

GetRolloverExpiresAfterDays returns the RolloverExpiresAfterDays field if non-nil, zero value otherwise.

### GetRolloverExpiresAfterDaysOk

`func (o *CreditConfig) GetRolloverExpiresAfterDaysOk() (*int32, bool)`

GetRolloverExpiresAfterDaysOk returns a tuple with the RolloverExpiresAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverExpiresAfterDays

`func (o *CreditConfig) SetRolloverExpiresAfterDays(v int32)`

SetRolloverExpiresAfterDays sets RolloverExpiresAfterDays field to given value.

### HasRolloverExpiresAfterDays

`func (o *CreditConfig) HasRolloverExpiresAfterDays() bool`

HasRolloverExpiresAfterDays returns a boolean if a field has been set.

### GetRolloverPriorityDelta

`func (o *CreditConfig) GetRolloverPriorityDelta() int32`

GetRolloverPriorityDelta returns the RolloverPriorityDelta field if non-nil, zero value otherwise.

### GetRolloverPriorityDeltaOk

`func (o *CreditConfig) GetRolloverPriorityDeltaOk() (*int32, bool)`

GetRolloverPriorityDeltaOk returns a tuple with the RolloverPriorityDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPriorityDelta

`func (o *CreditConfig) SetRolloverPriorityDelta(v int32)`

SetRolloverPriorityDelta sets RolloverPriorityDelta field to given value.

### HasRolloverPriorityDelta

`func (o *CreditConfig) HasRolloverPriorityDelta() bool`

HasRolloverPriorityDelta returns a boolean if a field has been set.

### GetFeaturePools

`func (o *CreditConfig) GetFeaturePools() []FeaturePool`

GetFeaturePools returns the FeaturePools field if non-nil, zero value otherwise.

### GetFeaturePoolsOk

`func (o *CreditConfig) GetFeaturePoolsOk() (*[]FeaturePool, bool)`

GetFeaturePoolsOk returns a tuple with the FeaturePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePools

`func (o *CreditConfig) SetFeaturePools(v []FeaturePool)`

SetFeaturePools sets FeaturePools field to given value.

### HasFeaturePools

`func (o *CreditConfig) HasFeaturePools() bool`

HasFeaturePools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


