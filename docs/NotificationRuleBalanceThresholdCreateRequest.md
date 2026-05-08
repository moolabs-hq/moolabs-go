# NotificationRuleBalanceThresholdCreateRequest

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

### NewNotificationRuleBalanceThresholdCreateRequest

`func NewNotificationRuleBalanceThresholdCreateRequest(type_ string, name string, thresholds []NotificationRuleBalanceThresholdValue, channels []string, ) *NotificationRuleBalanceThresholdCreateRequest`

NewNotificationRuleBalanceThresholdCreateRequest instantiates a new NotificationRuleBalanceThresholdCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleBalanceThresholdCreateRequestWithDefaults

`func NewNotificationRuleBalanceThresholdCreateRequestWithDefaults() *NotificationRuleBalanceThresholdCreateRequest`

NewNotificationRuleBalanceThresholdCreateRequestWithDefaults instantiates a new NotificationRuleBalanceThresholdCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRuleBalanceThresholdCreateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRuleBalanceThresholdCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetThresholds

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetThresholds() []NotificationRuleBalanceThresholdValue`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetThresholdsOk() (*[]NotificationRuleBalanceThresholdValue, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetThresholds(v []NotificationRuleBalanceThresholdValue)`

SetThresholds sets Thresholds field to given value.


### GetChannels

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetFeatures

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NotificationRuleBalanceThresholdCreateRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NotificationRuleBalanceThresholdCreateRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *NotificationRuleBalanceThresholdCreateRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


