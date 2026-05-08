# NotificationRule

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

### NewNotificationRule

`func NewNotificationRule(createdAt time.Time, updatedAt time.Time, id string, type_ string, name string, channels []NotificationChannelMeta, thresholds []NotificationRuleBalanceThresholdValue, ) *NotificationRule`

NewNotificationRule instantiates a new NotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleWithDefaults

`func NewNotificationRuleWithDefaults() *NotificationRule`

NewNotificationRuleWithDefaults instantiates a new NotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NotificationRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NotificationRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *NotificationRule) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NotificationRule) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NotificationRule) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NotificationRule) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *NotificationRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRule) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NotificationRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRule) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRule) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetChannels

`func (o *NotificationRule) GetChannels() []NotificationChannelMeta`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRule) GetChannelsOk() (*[]NotificationChannelMeta, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRule) SetChannels(v []NotificationChannelMeta)`

SetChannels sets Channels field to given value.


### GetAnnotations

`func (o *NotificationRule) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *NotificationRule) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *NotificationRule) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *NotificationRule) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRule) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRule) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRule) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetThresholds

`func (o *NotificationRule) GetThresholds() []NotificationRuleBalanceThresholdValue`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *NotificationRule) GetThresholdsOk() (*[]NotificationRuleBalanceThresholdValue, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *NotificationRule) SetThresholds(v []NotificationRuleBalanceThresholdValue)`

SetThresholds sets Thresholds field to given value.


### GetFeatures

`func (o *NotificationRule) GetFeatures() []FeatureMeta`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NotificationRule) GetFeaturesOk() (*[]FeatureMeta, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NotificationRule) SetFeatures(v []FeatureMeta)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *NotificationRule) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


