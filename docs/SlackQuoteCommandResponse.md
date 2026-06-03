# SlackQuoteCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseType** | Pointer to **string** |  | [optional] [default to "ephemeral"]
**Text** | **string** |  | 
**CommandKind** | **string** |  | 
**DeepLinks** | Pointer to **[]map[string]string** |  | [optional] 

## Methods

### NewSlackQuoteCommandResponse

`func NewSlackQuoteCommandResponse(text string, commandKind string, ) *SlackQuoteCommandResponse`

NewSlackQuoteCommandResponse instantiates a new SlackQuoteCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackQuoteCommandResponseWithDefaults

`func NewSlackQuoteCommandResponseWithDefaults() *SlackQuoteCommandResponse`

NewSlackQuoteCommandResponseWithDefaults instantiates a new SlackQuoteCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseType

`func (o *SlackQuoteCommandResponse) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *SlackQuoteCommandResponse) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *SlackQuoteCommandResponse) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *SlackQuoteCommandResponse) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetText

`func (o *SlackQuoteCommandResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SlackQuoteCommandResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SlackQuoteCommandResponse) SetText(v string)`

SetText sets Text field to given value.


### GetCommandKind

`func (o *SlackQuoteCommandResponse) GetCommandKind() string`

GetCommandKind returns the CommandKind field if non-nil, zero value otherwise.

### GetCommandKindOk

`func (o *SlackQuoteCommandResponse) GetCommandKindOk() (*string, bool)`

GetCommandKindOk returns a tuple with the CommandKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandKind

`func (o *SlackQuoteCommandResponse) SetCommandKind(v string)`

SetCommandKind sets CommandKind field to given value.


### GetDeepLinks

`func (o *SlackQuoteCommandResponse) GetDeepLinks() []map[string]string`

GetDeepLinks returns the DeepLinks field if non-nil, zero value otherwise.

### GetDeepLinksOk

`func (o *SlackQuoteCommandResponse) GetDeepLinksOk() (*[]map[string]string, bool)`

GetDeepLinksOk returns a tuple with the DeepLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLinks

`func (o *SlackQuoteCommandResponse) SetDeepLinks(v []map[string]string)`

SetDeepLinks sets DeepLinks field to given value.

### HasDeepLinks

`func (o *SlackQuoteCommandResponse) HasDeepLinks() bool`

HasDeepLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


