# NotificationRuleInvoiceUpdatedCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Notification rule type. | 
**Name** | **string** | The user friendly name of the notification rule. | 
**Disabled** | Pointer to **bool** | Whether the rule is disabled or not. | [optional] [default to false]
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Channels** | **[]string** | List of notification channels the rule is applied to. | 

## Methods

### NewNotificationRuleInvoiceUpdatedCreateRequest

`func NewNotificationRuleInvoiceUpdatedCreateRequest(type_ string, name string, channels []string, ) *NotificationRuleInvoiceUpdatedCreateRequest`

NewNotificationRuleInvoiceUpdatedCreateRequest instantiates a new NotificationRuleInvoiceUpdatedCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleInvoiceUpdatedCreateRequestWithDefaults

`func NewNotificationRuleInvoiceUpdatedCreateRequestWithDefaults() *NotificationRuleInvoiceUpdatedCreateRequest`

NewNotificationRuleInvoiceUpdatedCreateRequestWithDefaults instantiates a new NotificationRuleInvoiceUpdatedCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetChannels

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleInvoiceUpdatedCreateRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


