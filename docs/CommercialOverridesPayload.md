# CommercialOverridesPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RateCardDiscount** | Pointer to [**RateCardDiscountPayload**](RateCardDiscountPayload.md) |  | [optional] 
**CreditOverride** | Pointer to [**CreditOverridePayload**](CreditOverridePayload.md) |  | [optional] 
**EntitlementOverrides** | Pointer to [**[]EntitlementOverridePayload**](EntitlementOverridePayload.md) |  | [optional] 
**OverdraftSettings** | Pointer to [**OverdraftSettingsPayload**](OverdraftSettingsPayload.md) |  | [optional] 

## Methods

### NewCommercialOverridesPayload

`func NewCommercialOverridesPayload() *CommercialOverridesPayload`

NewCommercialOverridesPayload instantiates a new CommercialOverridesPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercialOverridesPayloadWithDefaults

`func NewCommercialOverridesPayloadWithDefaults() *CommercialOverridesPayload`

NewCommercialOverridesPayloadWithDefaults instantiates a new CommercialOverridesPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRateCardDiscount

`func (o *CommercialOverridesPayload) GetRateCardDiscount() RateCardDiscountPayload`

GetRateCardDiscount returns the RateCardDiscount field if non-nil, zero value otherwise.

### GetRateCardDiscountOk

`func (o *CommercialOverridesPayload) GetRateCardDiscountOk() (*RateCardDiscountPayload, bool)`

GetRateCardDiscountOk returns a tuple with the RateCardDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCardDiscount

`func (o *CommercialOverridesPayload) SetRateCardDiscount(v RateCardDiscountPayload)`

SetRateCardDiscount sets RateCardDiscount field to given value.

### HasRateCardDiscount

`func (o *CommercialOverridesPayload) HasRateCardDiscount() bool`

HasRateCardDiscount returns a boolean if a field has been set.

### GetCreditOverride

`func (o *CommercialOverridesPayload) GetCreditOverride() CreditOverridePayload`

GetCreditOverride returns the CreditOverride field if non-nil, zero value otherwise.

### GetCreditOverrideOk

`func (o *CommercialOverridesPayload) GetCreditOverrideOk() (*CreditOverridePayload, bool)`

GetCreditOverrideOk returns a tuple with the CreditOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditOverride

`func (o *CommercialOverridesPayload) SetCreditOverride(v CreditOverridePayload)`

SetCreditOverride sets CreditOverride field to given value.

### HasCreditOverride

`func (o *CommercialOverridesPayload) HasCreditOverride() bool`

HasCreditOverride returns a boolean if a field has been set.

### GetEntitlementOverrides

`func (o *CommercialOverridesPayload) GetEntitlementOverrides() []EntitlementOverridePayload`

GetEntitlementOverrides returns the EntitlementOverrides field if non-nil, zero value otherwise.

### GetEntitlementOverridesOk

`func (o *CommercialOverridesPayload) GetEntitlementOverridesOk() (*[]EntitlementOverridePayload, bool)`

GetEntitlementOverridesOk returns a tuple with the EntitlementOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementOverrides

`func (o *CommercialOverridesPayload) SetEntitlementOverrides(v []EntitlementOverridePayload)`

SetEntitlementOverrides sets EntitlementOverrides field to given value.

### HasEntitlementOverrides

`func (o *CommercialOverridesPayload) HasEntitlementOverrides() bool`

HasEntitlementOverrides returns a boolean if a field has been set.

### GetOverdraftSettings

`func (o *CommercialOverridesPayload) GetOverdraftSettings() OverdraftSettingsPayload`

GetOverdraftSettings returns the OverdraftSettings field if non-nil, zero value otherwise.

### GetOverdraftSettingsOk

`func (o *CommercialOverridesPayload) GetOverdraftSettingsOk() (*OverdraftSettingsPayload, bool)`

GetOverdraftSettingsOk returns a tuple with the OverdraftSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftSettings

`func (o *CommercialOverridesPayload) SetOverdraftSettings(v OverdraftSettingsPayload)`

SetOverdraftSettings sets OverdraftSettings field to given value.

### HasOverdraftSettings

`func (o *CommercialOverridesPayload) HasOverdraftSettings() bool`

HasOverdraftSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


