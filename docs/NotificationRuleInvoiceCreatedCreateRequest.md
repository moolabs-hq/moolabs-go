# NotificationRuleInvoiceCreatedCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Notification rule type. | 
**Name** | **string** | The user friendly name of the notification rule. | 
**Disabled** | Pointer to **bool** | Whether the rule is disabled or not. | [optional] [default to false]
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Channels** | **[]string** | List of notification channels the rule is applied to. | 

## Methods

### NewNotificationRuleInvoiceCreatedCreateRequest

`func NewNotificationRuleInvoiceCreatedCreateRequest(type_ string, name string, channels []string, ) *NotificationRuleInvoiceCreatedCreateRequest`

NewNotificationRuleInvoiceCreatedCreateRequest instantiates a new NotificationRuleInvoiceCreatedCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleInvoiceCreatedCreateRequestWithDefaults

`func NewNotificationRuleInvoiceCreatedCreateRequestWithDefaults() *NotificationRuleInvoiceCreatedCreateRequest`

NewNotificationRuleInvoiceCreatedCreateRequestWithDefaults instantiates a new NotificationRuleInvoiceCreatedCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRuleInvoiceCreatedCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRuleInvoiceCreatedCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisabled

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *NotificationRuleInvoiceCreatedCreateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *NotificationRuleInvoiceCreatedCreateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationRuleInvoiceCreatedCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotificationRuleInvoiceCreatedCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetChannels

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationRuleInvoiceCreatedCreateRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationRuleInvoiceCreatedCreateRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


