# JurisdictionProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AccountId** | **string** |  | 
**EffectiveCountryCode** | **string** |  | 
**EffectiveRegionCode** | Pointer to **string** |  | [optional] 
**EffectiveTimezone** | **string** |  | 
**RuleProfile** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewJurisdictionProfileResponse

`func NewJurisdictionProfileResponse(id string, accountId string, effectiveCountryCode string, effectiveTimezone string, ruleProfile string, createdAt time.Time, updatedAt time.Time, ) *JurisdictionProfileResponse`

NewJurisdictionProfileResponse instantiates a new JurisdictionProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJurisdictionProfileResponseWithDefaults

`func NewJurisdictionProfileResponseWithDefaults() *JurisdictionProfileResponse`

NewJurisdictionProfileResponseWithDefaults instantiates a new JurisdictionProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JurisdictionProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JurisdictionProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JurisdictionProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *JurisdictionProfileResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *JurisdictionProfileResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *JurisdictionProfileResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEffectiveCountryCode

`func (o *JurisdictionProfileResponse) GetEffectiveCountryCode() string`

GetEffectiveCountryCode returns the EffectiveCountryCode field if non-nil, zero value otherwise.

### GetEffectiveCountryCodeOk

`func (o *JurisdictionProfileResponse) GetEffectiveCountryCodeOk() (*string, bool)`

GetEffectiveCountryCodeOk returns a tuple with the EffectiveCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCountryCode

`func (o *JurisdictionProfileResponse) SetEffectiveCountryCode(v string)`

SetEffectiveCountryCode sets EffectiveCountryCode field to given value.


### GetEffectiveRegionCode

`func (o *JurisdictionProfileResponse) GetEffectiveRegionCode() string`

GetEffectiveRegionCode returns the EffectiveRegionCode field if non-nil, zero value otherwise.

### GetEffectiveRegionCodeOk

`func (o *JurisdictionProfileResponse) GetEffectiveRegionCodeOk() (*string, bool)`

GetEffectiveRegionCodeOk returns a tuple with the EffectiveRegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRegionCode

`func (o *JurisdictionProfileResponse) SetEffectiveRegionCode(v string)`

SetEffectiveRegionCode sets EffectiveRegionCode field to given value.

### HasEffectiveRegionCode

`func (o *JurisdictionProfileResponse) HasEffectiveRegionCode() bool`

HasEffectiveRegionCode returns a boolean if a field has been set.

### GetEffectiveTimezone

`func (o *JurisdictionProfileResponse) GetEffectiveTimezone() string`

GetEffectiveTimezone returns the EffectiveTimezone field if non-nil, zero value otherwise.

### GetEffectiveTimezoneOk

`func (o *JurisdictionProfileResponse) GetEffectiveTimezoneOk() (*string, bool)`

GetEffectiveTimezoneOk returns a tuple with the EffectiveTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTimezone

`func (o *JurisdictionProfileResponse) SetEffectiveTimezone(v string)`

SetEffectiveTimezone sets EffectiveTimezone field to given value.


### GetRuleProfile

`func (o *JurisdictionProfileResponse) GetRuleProfile() string`

GetRuleProfile returns the RuleProfile field if non-nil, zero value otherwise.

### GetRuleProfileOk

`func (o *JurisdictionProfileResponse) GetRuleProfileOk() (*string, bool)`

GetRuleProfileOk returns a tuple with the RuleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleProfile

`func (o *JurisdictionProfileResponse) SetRuleProfile(v string)`

SetRuleProfile sets RuleProfile field to given value.


### GetCreatedAt

`func (o *JurisdictionProfileResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JurisdictionProfileResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JurisdictionProfileResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *JurisdictionProfileResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JurisdictionProfileResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JurisdictionProfileResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


