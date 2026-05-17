# CreditConfigPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**GlobalPoolEnabled** | **bool** |  | 
**FeaturePoolsEnabled** | **bool** |  | 
**InitialAmount** | [**Initialamount**](Initialamount.md) |  | 
**RecurringAmount** | [**Recurringamount**](Recurringamount.md) |  | 
**Cadence** | **string** |  | 
**ExpiresInDays** | **int32** |  | 
**RolloverEnabled** | **bool** |  | 
**RolloverPercent** | Pointer to [**NullableRolloverpercent**](Rolloverpercent.md) |  | [optional] 
**RolloverCapAmount** | Pointer to [**NullableRollovercapamount**](Rollovercapamount.md) |  | [optional] 
**RolloverExpiresAfterDays** | Pointer to **int32** |  | [optional] 
**RolloverPriorityDelta** | Pointer to **int32** |  | [optional] 
**FeaturePools** | Pointer to [**[]FeaturePoolPayload**](FeaturePoolPayload.md) |  | [optional] 

## Methods

### NewCreditConfigPayload

`func NewCreditConfigPayload(enabled bool, globalPoolEnabled bool, featurePoolsEnabled bool, initialAmount Initialamount, recurringAmount Recurringamount, cadence string, expiresInDays int32, rolloverEnabled bool, ) *CreditConfigPayload`

NewCreditConfigPayload instantiates a new CreditConfigPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditConfigPayloadWithDefaults

`func NewCreditConfigPayloadWithDefaults() *CreditConfigPayload`

NewCreditConfigPayloadWithDefaults instantiates a new CreditConfigPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CreditConfigPayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreditConfigPayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreditConfigPayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGlobalPoolEnabled

`func (o *CreditConfigPayload) GetGlobalPoolEnabled() bool`

GetGlobalPoolEnabled returns the GlobalPoolEnabled field if non-nil, zero value otherwise.

### GetGlobalPoolEnabledOk

`func (o *CreditConfigPayload) GetGlobalPoolEnabledOk() (*bool, bool)`

GetGlobalPoolEnabledOk returns a tuple with the GlobalPoolEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPoolEnabled

`func (o *CreditConfigPayload) SetGlobalPoolEnabled(v bool)`

SetGlobalPoolEnabled sets GlobalPoolEnabled field to given value.


### GetFeaturePoolsEnabled

`func (o *CreditConfigPayload) GetFeaturePoolsEnabled() bool`

GetFeaturePoolsEnabled returns the FeaturePoolsEnabled field if non-nil, zero value otherwise.

### GetFeaturePoolsEnabledOk

`func (o *CreditConfigPayload) GetFeaturePoolsEnabledOk() (*bool, bool)`

GetFeaturePoolsEnabledOk returns a tuple with the FeaturePoolsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePoolsEnabled

`func (o *CreditConfigPayload) SetFeaturePoolsEnabled(v bool)`

SetFeaturePoolsEnabled sets FeaturePoolsEnabled field to given value.


### GetInitialAmount

`func (o *CreditConfigPayload) GetInitialAmount() Initialamount`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *CreditConfigPayload) GetInitialAmountOk() (*Initialamount, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *CreditConfigPayload) SetInitialAmount(v Initialamount)`

SetInitialAmount sets InitialAmount field to given value.


### GetRecurringAmount

`func (o *CreditConfigPayload) GetRecurringAmount() Recurringamount`

GetRecurringAmount returns the RecurringAmount field if non-nil, zero value otherwise.

### GetRecurringAmountOk

`func (o *CreditConfigPayload) GetRecurringAmountOk() (*Recurringamount, bool)`

GetRecurringAmountOk returns a tuple with the RecurringAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmount

`func (o *CreditConfigPayload) SetRecurringAmount(v Recurringamount)`

SetRecurringAmount sets RecurringAmount field to given value.


### GetCadence

`func (o *CreditConfigPayload) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *CreditConfigPayload) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *CreditConfigPayload) SetCadence(v string)`

SetCadence sets Cadence field to given value.


### GetExpiresInDays

`func (o *CreditConfigPayload) GetExpiresInDays() int32`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *CreditConfigPayload) GetExpiresInDaysOk() (*int32, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *CreditConfigPayload) SetExpiresInDays(v int32)`

SetExpiresInDays sets ExpiresInDays field to given value.


### GetRolloverEnabled

`func (o *CreditConfigPayload) GetRolloverEnabled() bool`

GetRolloverEnabled returns the RolloverEnabled field if non-nil, zero value otherwise.

### GetRolloverEnabledOk

`func (o *CreditConfigPayload) GetRolloverEnabledOk() (*bool, bool)`

GetRolloverEnabledOk returns a tuple with the RolloverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverEnabled

`func (o *CreditConfigPayload) SetRolloverEnabled(v bool)`

SetRolloverEnabled sets RolloverEnabled field to given value.


### GetRolloverPercent

`func (o *CreditConfigPayload) GetRolloverPercent() Rolloverpercent`

GetRolloverPercent returns the RolloverPercent field if non-nil, zero value otherwise.

### GetRolloverPercentOk

`func (o *CreditConfigPayload) GetRolloverPercentOk() (*Rolloverpercent, bool)`

GetRolloverPercentOk returns a tuple with the RolloverPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPercent

`func (o *CreditConfigPayload) SetRolloverPercent(v Rolloverpercent)`

SetRolloverPercent sets RolloverPercent field to given value.

### HasRolloverPercent

`func (o *CreditConfigPayload) HasRolloverPercent() bool`

HasRolloverPercent returns a boolean if a field has been set.

### SetRolloverPercentNil

`func (o *CreditConfigPayload) SetRolloverPercentNil(b bool)`

 SetRolloverPercentNil sets the value for RolloverPercent to be an explicit nil

### UnsetRolloverPercent
`func (o *CreditConfigPayload) UnsetRolloverPercent()`

UnsetRolloverPercent ensures that no value is present for RolloverPercent, not even an explicit nil
### GetRolloverCapAmount

`func (o *CreditConfigPayload) GetRolloverCapAmount() Rollovercapamount`

GetRolloverCapAmount returns the RolloverCapAmount field if non-nil, zero value otherwise.

### GetRolloverCapAmountOk

`func (o *CreditConfigPayload) GetRolloverCapAmountOk() (*Rollovercapamount, bool)`

GetRolloverCapAmountOk returns a tuple with the RolloverCapAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverCapAmount

`func (o *CreditConfigPayload) SetRolloverCapAmount(v Rollovercapamount)`

SetRolloverCapAmount sets RolloverCapAmount field to given value.

### HasRolloverCapAmount

`func (o *CreditConfigPayload) HasRolloverCapAmount() bool`

HasRolloverCapAmount returns a boolean if a field has been set.

### SetRolloverCapAmountNil

`func (o *CreditConfigPayload) SetRolloverCapAmountNil(b bool)`

 SetRolloverCapAmountNil sets the value for RolloverCapAmount to be an explicit nil

### UnsetRolloverCapAmount
`func (o *CreditConfigPayload) UnsetRolloverCapAmount()`

UnsetRolloverCapAmount ensures that no value is present for RolloverCapAmount, not even an explicit nil
### GetRolloverExpiresAfterDays

`func (o *CreditConfigPayload) GetRolloverExpiresAfterDays() int32`

GetRolloverExpiresAfterDays returns the RolloverExpiresAfterDays field if non-nil, zero value otherwise.

### GetRolloverExpiresAfterDaysOk

`func (o *CreditConfigPayload) GetRolloverExpiresAfterDaysOk() (*int32, bool)`

GetRolloverExpiresAfterDaysOk returns a tuple with the RolloverExpiresAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverExpiresAfterDays

`func (o *CreditConfigPayload) SetRolloverExpiresAfterDays(v int32)`

SetRolloverExpiresAfterDays sets RolloverExpiresAfterDays field to given value.

### HasRolloverExpiresAfterDays

`func (o *CreditConfigPayload) HasRolloverExpiresAfterDays() bool`

HasRolloverExpiresAfterDays returns a boolean if a field has been set.

### GetRolloverPriorityDelta

`func (o *CreditConfigPayload) GetRolloverPriorityDelta() int32`

GetRolloverPriorityDelta returns the RolloverPriorityDelta field if non-nil, zero value otherwise.

### GetRolloverPriorityDeltaOk

`func (o *CreditConfigPayload) GetRolloverPriorityDeltaOk() (*int32, bool)`

GetRolloverPriorityDeltaOk returns a tuple with the RolloverPriorityDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPriorityDelta

`func (o *CreditConfigPayload) SetRolloverPriorityDelta(v int32)`

SetRolloverPriorityDelta sets RolloverPriorityDelta field to given value.

### HasRolloverPriorityDelta

`func (o *CreditConfigPayload) HasRolloverPriorityDelta() bool`

HasRolloverPriorityDelta returns a boolean if a field has been set.

### GetFeaturePools

`func (o *CreditConfigPayload) GetFeaturePools() []FeaturePoolPayload`

GetFeaturePools returns the FeaturePools field if non-nil, zero value otherwise.

### GetFeaturePoolsOk

`func (o *CreditConfigPayload) GetFeaturePoolsOk() (*[]FeaturePoolPayload, bool)`

GetFeaturePoolsOk returns a tuple with the FeaturePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePools

`func (o *CreditConfigPayload) SetFeaturePools(v []FeaturePoolPayload)`

SetFeaturePools sets FeaturePools field to given value.

### HasFeaturePools

`func (o *CreditConfigPayload) HasFeaturePools() bool`

HasFeaturePools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


