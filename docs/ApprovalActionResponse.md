# ApprovalActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationId** | **string** |  | 
**Action** | **string** |  | 
**NewStatus** | **string** |  | 
**ExternalMessageId** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewApprovalActionResponse

`func NewApprovalActionResponse(communicationId string, action string, newStatus string, message string, ) *ApprovalActionResponse`

NewApprovalActionResponse instantiates a new ApprovalActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalActionResponseWithDefaults

`func NewApprovalActionResponseWithDefaults() *ApprovalActionResponse`

NewApprovalActionResponseWithDefaults instantiates a new ApprovalActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationId

`func (o *ApprovalActionResponse) GetCommunicationId() string`

GetCommunicationId returns the CommunicationId field if non-nil, zero value otherwise.

### GetCommunicationIdOk

`func (o *ApprovalActionResponse) GetCommunicationIdOk() (*string, bool)`

GetCommunicationIdOk returns a tuple with the CommunicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationId

`func (o *ApprovalActionResponse) SetCommunicationId(v string)`

SetCommunicationId sets CommunicationId field to given value.


### GetAction

`func (o *ApprovalActionResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApprovalActionResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApprovalActionResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetNewStatus

`func (o *ApprovalActionResponse) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *ApprovalActionResponse) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *ApprovalActionResponse) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.


### GetExternalMessageId

`func (o *ApprovalActionResponse) GetExternalMessageId() string`

GetExternalMessageId returns the ExternalMessageId field if non-nil, zero value otherwise.

### GetExternalMessageIdOk

`func (o *ApprovalActionResponse) GetExternalMessageIdOk() (*string, bool)`

GetExternalMessageIdOk returns a tuple with the ExternalMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMessageId

`func (o *ApprovalActionResponse) SetExternalMessageId(v string)`

SetExternalMessageId sets ExternalMessageId field to given value.

### HasExternalMessageId

`func (o *ApprovalActionResponse) HasExternalMessageId() bool`

HasExternalMessageId returns a boolean if a field has been set.

### GetMessage

`func (o *ApprovalActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApprovalActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApprovalActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


