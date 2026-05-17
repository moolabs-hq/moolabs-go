# OverdraftSettingsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletPolicy** | **string** |  | 
**LocalHardCapMicros** | Pointer to **int32** |  | [optional] 
**GracePeriodSeconds** | Pointer to **int32** |  | [optional] 
**MaxSpendRatePerHourMicros** | Pointer to **int32** |  | [optional] 

## Methods

### NewOverdraftSettingsPayload

`func NewOverdraftSettingsPayload(walletPolicy string, ) *OverdraftSettingsPayload`

NewOverdraftSettingsPayload instantiates a new OverdraftSettingsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverdraftSettingsPayloadWithDefaults

`func NewOverdraftSettingsPayloadWithDefaults() *OverdraftSettingsPayload`

NewOverdraftSettingsPayloadWithDefaults instantiates a new OverdraftSettingsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletPolicy

`func (o *OverdraftSettingsPayload) GetWalletPolicy() string`

GetWalletPolicy returns the WalletPolicy field if non-nil, zero value otherwise.

### GetWalletPolicyOk

`func (o *OverdraftSettingsPayload) GetWalletPolicyOk() (*string, bool)`

GetWalletPolicyOk returns a tuple with the WalletPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletPolicy

`func (o *OverdraftSettingsPayload) SetWalletPolicy(v string)`

SetWalletPolicy sets WalletPolicy field to given value.


### GetLocalHardCapMicros

`func (o *OverdraftSettingsPayload) GetLocalHardCapMicros() int32`

GetLocalHardCapMicros returns the LocalHardCapMicros field if non-nil, zero value otherwise.

### GetLocalHardCapMicrosOk

`func (o *OverdraftSettingsPayload) GetLocalHardCapMicrosOk() (*int32, bool)`

GetLocalHardCapMicrosOk returns a tuple with the LocalHardCapMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardCapMicros

`func (o *OverdraftSettingsPayload) SetLocalHardCapMicros(v int32)`

SetLocalHardCapMicros sets LocalHardCapMicros field to given value.

### HasLocalHardCapMicros

`func (o *OverdraftSettingsPayload) HasLocalHardCapMicros() bool`

HasLocalHardCapMicros returns a boolean if a field has been set.

### GetGracePeriodSeconds

`func (o *OverdraftSettingsPayload) GetGracePeriodSeconds() int32`

GetGracePeriodSeconds returns the GracePeriodSeconds field if non-nil, zero value otherwise.

### GetGracePeriodSecondsOk

`func (o *OverdraftSettingsPayload) GetGracePeriodSecondsOk() (*int32, bool)`

GetGracePeriodSecondsOk returns a tuple with the GracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodSeconds

`func (o *OverdraftSettingsPayload) SetGracePeriodSeconds(v int32)`

SetGracePeriodSeconds sets GracePeriodSeconds field to given value.

### HasGracePeriodSeconds

`func (o *OverdraftSettingsPayload) HasGracePeriodSeconds() bool`

HasGracePeriodSeconds returns a boolean if a field has been set.

### GetMaxSpendRatePerHourMicros

`func (o *OverdraftSettingsPayload) GetMaxSpendRatePerHourMicros() int32`

GetMaxSpendRatePerHourMicros returns the MaxSpendRatePerHourMicros field if non-nil, zero value otherwise.

### GetMaxSpendRatePerHourMicrosOk

`func (o *OverdraftSettingsPayload) GetMaxSpendRatePerHourMicrosOk() (*int32, bool)`

GetMaxSpendRatePerHourMicrosOk returns a tuple with the MaxSpendRatePerHourMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpendRatePerHourMicros

`func (o *OverdraftSettingsPayload) SetMaxSpendRatePerHourMicros(v int32)`

SetMaxSpendRatePerHourMicros sets MaxSpendRatePerHourMicros field to given value.

### HasMaxSpendRatePerHourMicros

`func (o *OverdraftSettingsPayload) HasMaxSpendRatePerHourMicros() bool`

HasMaxSpendRatePerHourMicros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


