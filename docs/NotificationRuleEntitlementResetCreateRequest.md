# NotificationRuleEntitlementResetCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Notification rule type. | 
**Name** | **string** | The user friendly name of the notification rule. | 
**Disabled** | Pointer to **bool** | Whether the rule is disabled or not. | [optional] [default to false]
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Channels** | **[]string** | List of notification channels the rule is applied to. | 
**Features** | Pointer to **[]string** | Optional field for defining the scope of notification by feature. It may contain features by id or key. | [optional] 

## Methods

### NewNotificationRuleEntitlementResetCreateRequest

`func NewNotificationRuleEntitlementResetCreateRequest(type_ string, name string, channels []string, ) *NotificationRuleEntitlementResetCreateRequest`

NewNotificationRuleEntitlementResetCreateRequest instantiates a new NotificationRuleEntitlementResetCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleEntitlementResetCreateRequestWithDefaults

`func NewNotificationRuleEntitlementResetCreateRequestWithDefaults() *NotificationRuleEntitlementResetCreateRequest`

NewNotificationRuleEntitlementResetCreateRequestWithDefaults instantiates a new NotificationRuleEntitlementResetCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationRuleEntitlementResetCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleEntitlementResetCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleEntitlementResetCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRuleEntitlementResetCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleEntitlementResetCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleEntitlementResetCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRuleEntitlementResetCreateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRuleEntitlementResetCreateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRuleEntitlementResetCreateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRuleEntitlementResetCreateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRuleEntitlementResetCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRuleEntitlementResetCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRuleEntitlementResetCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRuleEntitlementResetCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetChannels

`func (o *NotificationRuleEntitlementResetCreateRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleEntitlementResetCreateRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleEntitlementResetCreateRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetFeatures

`func (o *NotificationRuleEntitlementResetCreateRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *NotificationRuleEntitlementResetCreateRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *NotificationRuleEntitlementResetCreateRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *NotificationRuleEntitlementResetCreateRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


