# NotificationRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Notification rule type. | 
**Name** | **string** | The user friendly name of the notification rule. | 
**Disabled** | Pointer to **bool** | Whether the rule is disabled or not. | [optional] [default to false]
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Thresholds** | [**[]NotificationRuleBalanceThresholdValue**](NotificationRuleBalanceThresholdValue.md) | List of thresholds the rule suppose to be triggered. | 
**Channels** | **[]string** | List of notification channels the rule is applied to. | 
**Features** | Pointer to **[]string** | Optional field for defining the scope of notification by feature. It may contain features by id or key. | [optional] 

## Methods

### NewNotificationRuleCreateRequest

`func NewNotificationRuleCreateRequest(type_ string, name string, thresholds []NotificationRuleBalanceThresholdValue, channels []string, ) *NotificationRuleCreateRequest`

NewNotificationRuleCreateRequest instantiates a new NotificationRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleCreateRequestWithDefaults

`func NewNotificationRuleCreateRequestWithDefaults() *NotificationRuleCreateRequest`

NewNotificationRuleCreateRequestWithDefaults instantiates a new NotificationRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationRuleCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRuleCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRuleCreateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRuleCreateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRuleCreateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRuleCreateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRuleCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRuleCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRuleCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRuleCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetThresholds

`func (o *NotificationRuleCreateRequest) GetThresholds() []NotificationRuleBalanceThresholdValue`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *NotificationRuleCreateRequest) GetThresholdsOk() (*[]NotificationRuleBalanceThresholdValue, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *NotificationRuleCreateRequest) SetThresholds(v []NotificationRuleBalanceThresholdValue)`

SetThresholds sets Thresholds field to given value.


### GetChannels

`func (o *NotificationRuleCreateRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleCreateRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleCreateRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetFeatures

`func (o *NotificationRuleCreateRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NotificationRuleCreateRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NotificationRuleCreateRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *NotificationRuleCreateRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


