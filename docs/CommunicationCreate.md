# CommunicationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** | outbound or inbound | [optional] [default to "outbound"]
**Channel** | Pointer to **string** | email, sms, phone, portal, letter, internal_note | [optional] [default to "email"]
**Audience** | Pointer to **string** | customer, internal, or inbound_system (optional; server infers default when omitted) | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**DunningStage** | Pointer to **string** |  | [optional] [default to "pre_due"]
**GeneratedBy** | Pointer to **string** | agent, human, system, portal_customer | [optional] [default to "human"]
**ThreadId** | Pointer to **string** |  | [optional] 
**ParentCommunicationId** | Pointer to **string** |  | [optional] 
**ExternalMessageId** | Pointer to **string** | Optional pre-set external message ID (test/migration use). Production sends populate this from the email provider response. | [optional] 

## Methods

### NewCommunicationCreate

`func NewCommunicationCreate() *CommunicationCreate`

NewCommunicationCreate instantiates a new CommunicationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationCreateWithDefaults

`func NewCommunicationCreateWithDefaults() *CommunicationCreate`

NewCommunicationCreateWithDefaults instantiates a new CommunicationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *CommunicationCreate) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *CommunicationCreate) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *CommunicationCreate) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *CommunicationCreate) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetDirection

`func (o *CommunicationCreate) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CommunicationCreate) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CommunicationCreate) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CommunicationCreate) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetChannel

`func (o *CommunicationCreate) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CommunicationCreate) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CommunicationCreate) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CommunicationCreate) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetAudience

`func (o *CommunicationCreate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CommunicationCreate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CommunicationCreate) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CommunicationCreate) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetSubject

`func (o *CommunicationCreate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommunicationCreate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommunicationCreate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CommunicationCreate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBodyText

`func (o *CommunicationCreate) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *CommunicationCreate) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *CommunicationCreate) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *CommunicationCreate) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBodyHtml

`func (o *CommunicationCreate) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *CommunicationCreate) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *CommunicationCreate) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *CommunicationCreate) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetDunningStage

`func (o *CommunicationCreate) GetDunningStage() string`

GetDunningStage returns the DunningStage field if non-nil, zero value otherwise.

### GetDunningStageOk

`func (o *CommunicationCreate) GetDunningStageOk() (*string, bool)`

GetDunningStageOk returns a tuple with the DunningStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningStage

`func (o *CommunicationCreate) SetDunningStage(v string)`

SetDunningStage sets DunningStage field to given value.

### HasDunningStage

`func (o *CommunicationCreate) HasDunningStage() bool`

HasDunningStage returns a boolean if a field has been set.

### GetGeneratedBy

`func (o *CommunicationCreate) GetGeneratedBy() string`

GetGeneratedBy returns the GeneratedBy field if non-nil, zero value otherwise.

### GetGeneratedByOk

`func (o *CommunicationCreate) GetGeneratedByOk() (*string, bool)`

GetGeneratedByOk returns a tuple with the GeneratedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedBy

`func (o *CommunicationCreate) SetGeneratedBy(v string)`

SetGeneratedBy sets GeneratedBy field to given value.

### HasGeneratedBy

`func (o *CommunicationCreate) HasGeneratedBy() bool`

HasGeneratedBy returns a boolean if a field has been set.

### GetThreadId

`func (o *CommunicationCreate) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *CommunicationCreate) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *CommunicationCreate) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *CommunicationCreate) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetParentCommunicationId

`func (o *CommunicationCreate) GetParentCommunicationId() string`

GetParentCommunicationId returns the ParentCommunicationId field if non-nil, zero value otherwise.

### GetParentCommunicationIdOk

`func (o *CommunicationCreate) GetParentCommunicationIdOk() (*string, bool)`

GetParentCommunicationIdOk returns a tuple with the ParentCommunicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommunicationId

`func (o *CommunicationCreate) SetParentCommunicationId(v string)`

SetParentCommunicationId sets ParentCommunicationId field to given value.

### HasParentCommunicationId

`func (o *CommunicationCreate) HasParentCommunicationId() bool`

HasParentCommunicationId returns a boolean if a field has been set.

### GetExternalMessageId

`func (o *CommunicationCreate) GetExternalMessageId() string`

GetExternalMessageId returns the ExternalMessageId field if non-nil, zero value otherwise.

### GetExternalMessageIdOk

`func (o *CommunicationCreate) GetExternalMessageIdOk() (*string, bool)`

GetExternalMessageIdOk returns a tuple with the ExternalMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMessageId

`func (o *CommunicationCreate) SetExternalMessageId(v string)`

SetExternalMessageId sets ExternalMessageId field to given value.

### HasExternalMessageId

`func (o *CommunicationCreate) HasExternalMessageId() bool`

HasExternalMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


