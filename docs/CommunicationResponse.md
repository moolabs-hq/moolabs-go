# CommunicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CaseId** | **string** |  | 
**AccountId** | Pointer to **string** |  | [optional] 
**ContactId** | Pointer to **string** |  | [optional] 
**Direction** | **string** |  | 
**Channel** | **string** |  | 
**Status** | **string** |  | 
**Audience** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**BodyHtml** | Pointer to **string** |  | [optional] 
**ThreadId** | Pointer to **string** |  | [optional] 
**ParentCommunicationId** | Pointer to **string** |  | [optional] 
**DunningStage** | **string** |  | 
**GeneratedBy** | **string** |  | 
**ApprovalRequestId** | Pointer to **string** |  | [optional] 
**SentAt** | Pointer to **time.Time** |  | [optional] 
**OpenedAt** | Pointer to **time.Time** |  | [optional] 
**ClickedAt** | Pointer to **time.Time** |  | [optional] 
**ReplyDetectedAt** | Pointer to **time.Time** |  | [optional] 
**ExternalMessageId** | Pointer to **string** |  | [optional] 
**ConsentSnapshot** | Pointer to **map[string]interface{}** |  | [optional] 
**JurisdictionSnapshot** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCommunicationResponse

`func NewCommunicationResponse(id string, tenantId string, caseId string, direction string, channel string, status string, dunningStage string, generatedBy string, createdAt time.Time, updatedAt time.Time, ) *CommunicationResponse`

NewCommunicationResponse instantiates a new CommunicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationResponseWithDefaults

`func NewCommunicationResponseWithDefaults() *CommunicationResponse`

NewCommunicationResponseWithDefaults instantiates a new CommunicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CommunicationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommunicationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommunicationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCaseId

`func (o *CommunicationResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CommunicationResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CommunicationResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetAccountId

`func (o *CommunicationResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CommunicationResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CommunicationResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CommunicationResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetContactId

`func (o *CommunicationResponse) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *CommunicationResponse) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *CommunicationResponse) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *CommunicationResponse) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetDirection

`func (o *CommunicationResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CommunicationResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CommunicationResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetChannel

`func (o *CommunicationResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CommunicationResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CommunicationResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetStatus

`func (o *CommunicationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommunicationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommunicationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAudience

`func (o *CommunicationResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CommunicationResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CommunicationResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CommunicationResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetSubject

`func (o *CommunicationResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommunicationResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommunicationResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CommunicationResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBodyText

`func (o *CommunicationResponse) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *CommunicationResponse) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *CommunicationResponse) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *CommunicationResponse) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBodyHtml

`func (o *CommunicationResponse) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *CommunicationResponse) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *CommunicationResponse) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *CommunicationResponse) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetThreadId

`func (o *CommunicationResponse) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *CommunicationResponse) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *CommunicationResponse) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *CommunicationResponse) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetParentCommunicationId

`func (o *CommunicationResponse) GetParentCommunicationId() string`

GetParentCommunicationId returns the ParentCommunicationId field if non-nil, zero value otherwise.

### GetParentCommunicationIdOk

`func (o *CommunicationResponse) GetParentCommunicationIdOk() (*string, bool)`

GetParentCommunicationIdOk returns a tuple with the ParentCommunicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommunicationId

`func (o *CommunicationResponse) SetParentCommunicationId(v string)`

SetParentCommunicationId sets ParentCommunicationId field to given value.

### HasParentCommunicationId

`func (o *CommunicationResponse) HasParentCommunicationId() bool`

HasParentCommunicationId returns a boolean if a field has been set.

### GetDunningStage

`func (o *CommunicationResponse) GetDunningStage() string`

GetDunningStage returns the DunningStage field if non-nil, zero value otherwise.

### GetDunningStageOk

`func (o *CommunicationResponse) GetDunningStageOk() (*string, bool)`

GetDunningStageOk returns a tuple with the DunningStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningStage

`func (o *CommunicationResponse) SetDunningStage(v string)`

SetDunningStage sets DunningStage field to given value.


### GetGeneratedBy

`func (o *CommunicationResponse) GetGeneratedBy() string`

GetGeneratedBy returns the GeneratedBy field if non-nil, zero value otherwise.

### GetGeneratedByOk

`func (o *CommunicationResponse) GetGeneratedByOk() (*string, bool)`

GetGeneratedByOk returns a tuple with the GeneratedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedBy

`func (o *CommunicationResponse) SetGeneratedBy(v string)`

SetGeneratedBy sets GeneratedBy field to given value.


### GetApprovalRequestId

`func (o *CommunicationResponse) GetApprovalRequestId() string`

GetApprovalRequestId returns the ApprovalRequestId field if non-nil, zero value otherwise.

### GetApprovalRequestIdOk

`func (o *CommunicationResponse) GetApprovalRequestIdOk() (*string, bool)`

GetApprovalRequestIdOk returns a tuple with the ApprovalRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequestId

`func (o *CommunicationResponse) SetApprovalRequestId(v string)`

SetApprovalRequestId sets ApprovalRequestId field to given value.

### HasApprovalRequestId

`func (o *CommunicationResponse) HasApprovalRequestId() bool`

HasApprovalRequestId returns a boolean if a field has been set.

### GetSentAt

`func (o *CommunicationResponse) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *CommunicationResponse) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *CommunicationResponse) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *CommunicationResponse) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetOpenedAt

`func (o *CommunicationResponse) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *CommunicationResponse) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *CommunicationResponse) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.

### HasOpenedAt

`func (o *CommunicationResponse) HasOpenedAt() bool`

HasOpenedAt returns a boolean if a field has been set.

### GetClickedAt

`func (o *CommunicationResponse) GetClickedAt() time.Time`

GetClickedAt returns the ClickedAt field if non-nil, zero value otherwise.

### GetClickedAtOk

`func (o *CommunicationResponse) GetClickedAtOk() (*time.Time, bool)`

GetClickedAtOk returns a tuple with the ClickedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedAt

`func (o *CommunicationResponse) SetClickedAt(v time.Time)`

SetClickedAt sets ClickedAt field to given value.

### HasClickedAt

`func (o *CommunicationResponse) HasClickedAt() bool`

HasClickedAt returns a boolean if a field has been set.

### GetReplyDetectedAt

`func (o *CommunicationResponse) GetReplyDetectedAt() time.Time`

GetReplyDetectedAt returns the ReplyDetectedAt field if non-nil, zero value otherwise.

### GetReplyDetectedAtOk

`func (o *CommunicationResponse) GetReplyDetectedAtOk() (*time.Time, bool)`

GetReplyDetectedAtOk returns a tuple with the ReplyDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyDetectedAt

`func (o *CommunicationResponse) SetReplyDetectedAt(v time.Time)`

SetReplyDetectedAt sets ReplyDetectedAt field to given value.

### HasReplyDetectedAt

`func (o *CommunicationResponse) HasReplyDetectedAt() bool`

HasReplyDetectedAt returns a boolean if a field has been set.

### GetExternalMessageId

`func (o *CommunicationResponse) GetExternalMessageId() string`

GetExternalMessageId returns the ExternalMessageId field if non-nil, zero value otherwise.

### GetExternalMessageIdOk

`func (o *CommunicationResponse) GetExternalMessageIdOk() (*string, bool)`

GetExternalMessageIdOk returns a tuple with the ExternalMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMessageId

`func (o *CommunicationResponse) SetExternalMessageId(v string)`

SetExternalMessageId sets ExternalMessageId field to given value.

### HasExternalMessageId

`func (o *CommunicationResponse) HasExternalMessageId() bool`

HasExternalMessageId returns a boolean if a field has been set.

### GetConsentSnapshot

`func (o *CommunicationResponse) GetConsentSnapshot() map[string]interface{}`

GetConsentSnapshot returns the ConsentSnapshot field if non-nil, zero value otherwise.

### GetConsentSnapshotOk

`func (o *CommunicationResponse) GetConsentSnapshotOk() (*map[string]interface{}, bool)`

GetConsentSnapshotOk returns a tuple with the ConsentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSnapshot

`func (o *CommunicationResponse) SetConsentSnapshot(v map[string]interface{})`

SetConsentSnapshot sets ConsentSnapshot field to given value.

### HasConsentSnapshot

`func (o *CommunicationResponse) HasConsentSnapshot() bool`

HasConsentSnapshot returns a boolean if a field has been set.

### GetJurisdictionSnapshot

`func (o *CommunicationResponse) GetJurisdictionSnapshot() map[string]interface{}`

GetJurisdictionSnapshot returns the JurisdictionSnapshot field if non-nil, zero value otherwise.

### GetJurisdictionSnapshotOk

`func (o *CommunicationResponse) GetJurisdictionSnapshotOk() (*map[string]interface{}, bool)`

GetJurisdictionSnapshotOk returns a tuple with the JurisdictionSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionSnapshot

`func (o *CommunicationResponse) SetJurisdictionSnapshot(v map[string]interface{})`

SetJurisdictionSnapshot sets JurisdictionSnapshot field to given value.

### HasJurisdictionSnapshot

`func (o *CommunicationResponse) HasJurisdictionSnapshot() bool`

HasJurisdictionSnapshot returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommunicationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommunicationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommunicationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CommunicationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommunicationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommunicationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


