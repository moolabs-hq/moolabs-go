# PortalPreferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **map[string]interface{}** | {\&quot;email\&quot;: {\&quot;opt_out\&quot;: false}, \&quot;sms\&quot;: {\&quot;opt_out\&quot;: true}} | 

## Methods

### NewPortalPreferencesRequest

`func NewPortalPreferencesRequest(channels map[string]interface{}, ) *PortalPreferencesRequest`

NewPortalPreferencesRequest instantiates a new PortalPreferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalPreferencesRequestWithDefaults

`func NewPortalPreferencesRequestWithDefaults() *PortalPreferencesRequest`

NewPortalPreferencesRequestWithDefaults instantiates a new PortalPreferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *PortalPreferencesRequest) GetChannels() map[string]interface{}`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *PortalPreferencesRequest) GetChannelsOk() (*map[string]interface{}, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *PortalPreferencesRequest) SetChannels(v map[string]interface{})`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


