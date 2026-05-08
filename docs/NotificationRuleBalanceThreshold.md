# NotificationRuleBalanceThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Id** | **string** | Identifies the notification rule. | [readonly] 
**Type** | **string** | Notification rule type. | 
**Name** | **string** | The user friendly name of the notification rule. | 
**Disabled** | Pointer to **bool** | Whether the rule is disabled or not. | [optional] [default to false]
**Channels** | [**[]NotificationChannelMeta**](NotificationChannelMeta.md) | List of notification channels the rule applies to. | 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Thresholds** | [**[]NotificationRuleBalanceThresholdValue**](NotificationRuleBalanceThresholdValue.md) | List of thresholds the rule suppose to be triggered. | 
**Features** | Pointer to [**[]FeatureMeta**](FeatureMeta.md) | Optional field containing list of features the rule applies to. | [optional] 

## Methods

### NewNotificationRuleBalanceThreshold

`func NewNotificationRuleBalanceThreshold(createdAt time.Time, updatedAt time.Time, id string, type_ string, name string, channels []NotificationChannelMeta, thresholds []NotificationRuleBalanceThresholdValue, ) *NotificationRuleBalanceThreshold`

NewNotificationRuleBalanceThreshold instantiates a new NotificationRuleBalanceThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleBalanceThresholdWithDefaults

`func NewNotificationRuleBalanceThresholdWithDefaults() *NotificationRuleBalanceThreshold`

NewNotificationRuleBalanceThresholdWithDefaults instantiates a new NotificationRuleBalanceThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NotificationRuleBalanceThreshold) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationRuleBalanceThreshold) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationRuleBalanceThreshold) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NotificationRuleBalanceThreshold) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationRuleBalanceThreshold) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationRuleBalanceThreshold) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *NotificationRuleBalanceThreshold) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NotificationRuleBalanceThreshold) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NotificationRuleBalanceThreshold) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NotificationRuleBalanceThreshold) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *NotificationRuleBalanceThreshold) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRuleBalanceThreshold) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRuleBalanceThreshold) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationRuleBalanceThreshold) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleBalanceThreshold) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleBalanceThreshold) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRuleBalanceThreshold) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleBalanceThreshold) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleBalanceThreshold) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRuleBalanceThreshold) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRuleBalanceThreshold) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRuleBalanceThreshold) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRuleBalanceThreshold) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetChannels

`func (o *NotificationRuleBalanceThreshold) GetChannels() []NotificationChannelMeta`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleBalanceThreshold) GetChannelsOk() (*[]NotificationChannelMeta, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleBalanceThreshold) SetChannels(v []NotificationChannelMeta)`

SetChannels sets Channels field to given value.


### GetAnnotations

`func (o *NotificationRuleBalanceThreshold) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NotificationRuleBalanceThreshold) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NotificationRuleBalanceThreshold) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NotificationRuleBalanceThreshold) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRuleBalanceThreshold) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRuleBalanceThreshold) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRuleBalanceThreshold) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRuleBalanceThreshold) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetThresholds

`func (o *NotificationRuleBalanceThreshold) GetThresholds() []NotificationRuleBalanceThresholdValue`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *NotificationRuleBalanceThreshold) GetThresholdsOk() (*[]NotificationRuleBalanceThresholdValue, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *NotificationRuleBalanceThreshold) SetThresholds(v []NotificationRuleBalanceThresholdValue)`

SetThresholds sets Thresholds field to given value.


### GetFeatures

`func (o *NotificationRuleBalanceThreshold) GetFeatures() []FeatureMeta`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NotificationRuleBalanceThreshold) GetFeaturesOk() (*[]FeatureMeta, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NotificationRuleBalanceThreshold) SetFeatures(v []FeatureMeta)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *NotificationRuleBalanceThreshold) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


