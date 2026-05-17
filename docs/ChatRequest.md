# ChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatMessageIn**](ChatMessageIn.md) |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewChatRequest

`func NewChatRequest(messages []ChatMessageIn, ) *ChatRequest`

NewChatRequest instantiates a new ChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatRequestWithDefaults

`func NewChatRequestWithDefaults() *ChatRequest`

NewChatRequestWithDefaults instantiates a new ChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatRequest) GetMessages() []ChatMessageIn`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatRequest) GetMessagesOk() (*[]ChatMessageIn, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatRequest) SetMessages(v []ChatMessageIn)`

SetMessages sets Messages field to given value.


### GetSessionId

`func (o *ChatRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStream

`func (o *ChatRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


