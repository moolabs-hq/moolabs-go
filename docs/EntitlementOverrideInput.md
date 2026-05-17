# EntitlementOverrideInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | **string** |  | 
**EntitlementType** | Pointer to **string** |  | [optional] [default to "metered"]
**DefaultValue** | Pointer to **string** |  | [optional] [default to ""]
**OverrideValue** | **string** |  | 
**Window** | Pointer to **string** |  | [optional] [default to "Monthly"]

## Methods

### NewEntitlementOverrideInput

`func NewEntitlementOverrideInput(featureKey string, overrideValue string, ) *EntitlementOverrideInput`

NewEntitlementOverrideInput instantiates a new EntitlementOverrideInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementOverrideInputWithDefaults

`func NewEntitlementOverrideInputWithDefaults() *EntitlementOverrideInput`

NewEntitlementOverrideInputWithDefaults instantiates a new EntitlementOverrideInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *EntitlementOverrideInput) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementOverrideInput) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementOverrideInput) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetEntitlementType

`func (o *EntitlementOverrideInput) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *EntitlementOverrideInput) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *EntitlementOverrideInput) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.

### HasEntitlementType

`func (o *EntitlementOverrideInput) HasEntitlementType() bool`

HasEntitlementType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *EntitlementOverrideInput) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *EntitlementOverrideInput) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *EntitlementOverrideInput) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *EntitlementOverrideInput) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetOverrideValue

`func (o *EntitlementOverrideInput) GetOverrideValue() string`

GetOverrideValue returns the OverrideValue field if non-nil, zero value otherwise.

### GetOverrideValueOk

`func (o *EntitlementOverrideInput) GetOverrideValueOk() (*string, bool)`

GetOverrideValueOk returns a tuple with the OverrideValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideValue

`func (o *EntitlementOverrideInput) SetOverrideValue(v string)`

SetOverrideValue sets OverrideValue field to given value.


### GetWindow

`func (o *EntitlementOverrideInput) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *EntitlementOverrideInput) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *EntitlementOverrideInput) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *EntitlementOverrideInput) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


