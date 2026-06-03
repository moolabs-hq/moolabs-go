# TemplateCardActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**ChatSessionId** | Pointer to **string** |  | [optional] 
**Confirmation** | Pointer to **map[string]interface{}** |  | [optional] 
**ChangeReason** | Pointer to **string** |  | [optional] 
**PreviewCaseId** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateCardActionRequest

`func NewTemplateCardActionRequest(action string, idempotencyKey string, ) *TemplateCardActionRequest`

NewTemplateCardActionRequest instantiates a new TemplateCardActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateCardActionRequestWithDefaults

`func NewTemplateCardActionRequestWithDefaults() *TemplateCardActionRequest`

NewTemplateCardActionRequestWithDefaults instantiates a new TemplateCardActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TemplateCardActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TemplateCardActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TemplateCardActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetIdempotencyKey

`func (o *TemplateCardActionRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *TemplateCardActionRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *TemplateCardActionRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetChatSessionId

`func (o *TemplateCardActionRequest) GetChatSessionId() string`

GetChatSessionId returns the ChatSessionId field if non-nil, zero value otherwise.

### GetChatSessionIdOk

`func (o *TemplateCardActionRequest) GetChatSessionIdOk() (*string, bool)`

GetChatSessionIdOk returns a tuple with the ChatSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSessionId

`func (o *TemplateCardActionRequest) SetChatSessionId(v string)`

SetChatSessionId sets ChatSessionId field to given value.

### HasChatSessionId

`func (o *TemplateCardActionRequest) HasChatSessionId() bool`

HasChatSessionId returns a boolean if a field has been set.

### GetConfirmation

`func (o *TemplateCardActionRequest) GetConfirmation() map[string]interface{}`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *TemplateCardActionRequest) GetConfirmationOk() (*map[string]interface{}, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *TemplateCardActionRequest) SetConfirmation(v map[string]interface{})`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *TemplateCardActionRequest) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetChangeReason

`func (o *TemplateCardActionRequest) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *TemplateCardActionRequest) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *TemplateCardActionRequest) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.

### HasChangeReason

`func (o *TemplateCardActionRequest) HasChangeReason() bool`

HasChangeReason returns a boolean if a field has been set.

### GetPreviewCaseId

`func (o *TemplateCardActionRequest) GetPreviewCaseId() string`

GetPreviewCaseId returns the PreviewCaseId field if non-nil, zero value otherwise.

### GetPreviewCaseIdOk

`func (o *TemplateCardActionRequest) GetPreviewCaseIdOk() (*string, bool)`

GetPreviewCaseIdOk returns a tuple with the PreviewCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewCaseId

`func (o *TemplateCardActionRequest) SetPreviewCaseId(v string)`

SetPreviewCaseId sets PreviewCaseId field to given value.

### HasPreviewCaseId

`func (o *TemplateCardActionRequest) HasPreviewCaseId() bool`

HasPreviewCaseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


