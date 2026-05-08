# OverdraftSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletPolicy** | **string** |  | 
**LocalHardCapMicros** | Pointer to **int32** |  | [optional] 
**GracePeriodSeconds** | Pointer to **int32** |  | [optional] 
**MaxSpendRatePerHourMicros** | Pointer to **int32** |  | [optional] 

## Methods

### NewOverdraftSettings

`func NewOverdraftSettings(walletPolicy string, ) *OverdraftSettings`

NewOverdraftSettings instantiates a new OverdraftSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverdraftSettingsWithDefaults

`func NewOverdraftSettingsWithDefaults() *OverdraftSettings`

NewOverdraftSettingsWithDefaults instantiates a new OverdraftSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletPolicy

`func (o *OverdraftSettings) GetWalletPolicy() string`

GetWalletPolicy returns the WalletPolicy field if non-nil, zero value otherwise.

### GetWalletPolicyOk

`func (o *OverdraftSettings) GetWalletPolicyOk() (*string, bool)`

GetWalletPolicyOk returns a tuple with the WalletPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletPolicy

`func (o *OverdraftSettings) SetWalletPolicy(v string)`

SetWalletPolicy sets WalletPolicy field to given value.


### GetLocalHardCapMicros

`func (o *OverdraftSettings) GetLocalHardCapMicros() int32`

GetLocalHardCapMicros returns the LocalHardCapMicros field if non-nil, zero value otherwise.

### GetLocalHardCapMicrosOk

`func (o *OverdraftSettings) GetLocalHardCapMicrosOk() (*int32, bool)`

GetLocalHardCapMicrosOk returns a tuple with the LocalHardCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardCapMicros

`func (o *OverdraftSettings) SetLocalHardCapMicros(v int32)`

SetLocalHardCapMicros sets LocalHardCapMicros field to given value.

### HasLocalHardCapMicros

`func (o *OverdraftSettings) HasLocalHardCapMicros() bool`

HasLocalHardCapMicros returns a boolean if a field has been set.

### GetGracePeriodSeconds

`func (o *OverdraftSettings) GetGracePeriodSeconds() int32`

GetGracePeriodSeconds returns the GracePeriodSeconds field if non-nil, zero value otherwise.

### GetGracePeriodSecondsOk

`func (o *OverdraftSettings) GetGracePeriodSecondsOk() (*int32, bool)`

GetGracePeriodSecondsOk returns a tuple with the GracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodSeconds

`func (o *OverdraftSettings) SetGracePeriodSeconds(v int32)`

SetGracePeriodSeconds sets GracePeriodSeconds field to given value.

### HasGracePeriodSeconds

`func (o *OverdraftSettings) HasGracePeriodSeconds() bool`

HasGracePeriodSeconds returns a boolean if a field has been set.

### GetMaxSpendRatePerHourMicros

`func (o *OverdraftSettings) GetMaxSpendRatePerHourMicros() int32`

GetMaxSpendRatePerHourMicros returns the MaxSpendRatePerHourMicros field if non-nil, zero value otherwise.

### GetMaxSpendRatePerHourMicrosOk

`func (o *OverdraftSettings) GetMaxSpendRatePerHourMicrosOk() (*int32, bool)`

GetMaxSpendRatePerHourMicrosOk returns a tuple with the MaxSpendRatePerHourMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpendRatePerHourMicros

`func (o *OverdraftSettings) SetMaxSpendRatePerHourMicros(v int32)`

SetMaxSpendRatePerHourMicros sets MaxSpendRatePerHourMicros field to given value.

### HasMaxSpendRatePerHourMicros

`func (o *OverdraftSettings) HasMaxSpendRatePerHourMicros() bool`

HasMaxSpendRatePerHourMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


