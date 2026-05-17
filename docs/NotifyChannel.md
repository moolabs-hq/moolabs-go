# NotifyChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Channel type — currently only &#39;email&#39; is supported | 
**Address** | **string** | Email address to notify | 

## Methods

### NewNotifyChannel

`func NewNotifyChannel(type_ string, address string, ) *NotifyChannel`

NewNotifyChannel instantiates a new NotifyChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyChannelWithDefaults

`func NewNotifyChannelWithDefaults() *NotifyChannel`

NewNotifyChannelWithDefaults instantiates a new NotifyChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotifyChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotifyChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotifyChannel) SetType(v string)`

SetType sets Type field to given value.


### GetAddress

`func (o *NotifyChannel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NotifyChannel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NotifyChannel) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


