# PublishTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**ConfirmationAck** | **bool** |  | 
**ChangeReason** | **string** |  | 
**PromptSummary** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewPublishTemplateRequest

`func NewPublishTemplateRequest(versionId string, confirmationAck bool, changeReason string, ) *PublishTemplateRequest`

NewPublishTemplateRequest instantiates a new PublishTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishTemplateRequestWithDefaults

`func NewPublishTemplateRequestWithDefaults() *PublishTemplateRequest`

NewPublishTemplateRequestWithDefaults instantiates a new PublishTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *PublishTemplateRequest) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *PublishTemplateRequest) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *PublishTemplateRequest) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetConfirmationAck

`func (o *PublishTemplateRequest) GetConfirmationAck() bool`

GetConfirmationAck returns the ConfirmationAck field if non-nil, zero value otherwise.

### GetConfirmationAckOk

`func (o *PublishTemplateRequest) GetConfirmationAckOk() (*bool, bool)`

GetConfirmationAckOk returns a tuple with the ConfirmationAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationAck

`func (o *PublishTemplateRequest) SetConfirmationAck(v bool)`

SetConfirmationAck sets ConfirmationAck field to given value.


### GetChangeReason

`func (o *PublishTemplateRequest) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *PublishTemplateRequest) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *PublishTemplateRequest) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.


### GetPromptSummary

`func (o *PublishTemplateRequest) GetPromptSummary() string`

GetPromptSummary returns the PromptSummary field if non-nil, zero value otherwise.

### GetPromptSummaryOk

`func (o *PublishTemplateRequest) GetPromptSummaryOk() (*string, bool)`

GetPromptSummaryOk returns a tuple with the PromptSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptSummary

`func (o *PublishTemplateRequest) SetPromptSummary(v string)`

SetPromptSummary sets PromptSummary field to given value.

### HasPromptSummary

`func (o *PublishTemplateRequest) HasPromptSummary() bool`

HasPromptSummary returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *PublishTemplateRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *PublishTemplateRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *PublishTemplateRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *PublishTemplateRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


