# FeatureFlagItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TenantOverride** | **bool** |  | 
**DefaultValue** | **bool** |  | 
**Enabled** | **bool** |  | 
**LastChanged** | **string** |  | 

## Methods

### NewFeatureFlagItem

`func NewFeatureFlagItem(id string, name string, tenantOverride bool, defaultValue bool, enabled bool, lastChanged string, ) *FeatureFlagItem`

NewFeatureFlagItem instantiates a new FeatureFlagItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagItemWithDefaults

`func NewFeatureFlagItemWithDefaults() *FeatureFlagItem`

NewFeatureFlagItemWithDefaults instantiates a new FeatureFlagItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureFlagItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureFlagItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureFlagItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FeatureFlagItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlagItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlagItem) SetName(v string)`

SetName sets Name field to given value.


### GetTenantOverride

`func (o *FeatureFlagItem) GetTenantOverride() bool`

GetTenantOverride returns the TenantOverride field if non-nil, zero value otherwise.

### GetTenantOverrideOk

`func (o *FeatureFlagItem) GetTenantOverrideOk() (*bool, bool)`

GetTenantOverrideOk returns a tuple with the TenantOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOverride

`func (o *FeatureFlagItem) SetTenantOverride(v bool)`

SetTenantOverride sets TenantOverride field to given value.


### GetDefaultValue

`func (o *FeatureFlagItem) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FeatureFlagItem) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FeatureFlagItem) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.


### GetEnabled

`func (o *FeatureFlagItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureFlagItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureFlagItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLastChanged

`func (o *FeatureFlagItem) GetLastChanged() string`

GetLastChanged returns the LastChanged field if non-nil, zero value otherwise.

### GetLastChangedOk

`func (o *FeatureFlagItem) GetLastChangedOk() (*string, bool)`

GetLastChangedOk returns a tuple with the LastChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChanged

`func (o *FeatureFlagItem) SetLastChanged(v string)`

SetLastChanged sets LastChanged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


