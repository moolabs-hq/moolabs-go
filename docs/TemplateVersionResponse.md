# TemplateVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CanonicalTemplateKey** | **string** |  | 
**Channel** | Pointer to **string** |  | [optional] [default to "email"]
**Version** | **int32** |  | 
**Status** | **string** |  | 
**DraftGroupId** | Pointer to **string** |  | [optional] 
**DraftState** | Pointer to **string** |  | [optional] 
**StaleReason** | Pointer to **string** |  | [optional] 
**BasePublishedVersionId** | Pointer to **string** |  | [optional] 
**SubjectTemplate** | **string** |  | 
**BodyTemplate** | **string** |  | 
**ContentHash** | **string** |  | 
**ValidationSnapshot** | **map[string]interface{}** |  | 
**DisclosurePolicyHash** | **string** |  | 
**ProviderReadinessHash** | **string** |  | 
**PaymentInstructionsHash** | **string** |  | 
**CreatedByActorId** | **string** |  | 
**CreatedByActorDisplay** | Pointer to **string** |  | [optional] 
**PublishedByActorId** | Pointer to **string** |  | [optional] 
**PublishedAt** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] [default to false]
**LockOwnerActorId** | Pointer to **string** |  | [optional] 
**LockOwnerDisplay** | Pointer to **string** |  | [optional] 
**LockExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateVersionResponse

`func NewTemplateVersionResponse(id string, canonicalTemplateKey string, version int32, status string, subjectTemplate string, bodyTemplate string, contentHash string, validationSnapshot map[string]interface{}, disclosurePolicyHash string, providerReadinessHash string, paymentInstructionsHash string, createdByActorId string, ) *TemplateVersionResponse`

NewTemplateVersionResponse instantiates a new TemplateVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionResponseWithDefaults

`func NewTemplateVersionResponseWithDefaults() *TemplateVersionResponse`

NewTemplateVersionResponseWithDefaults instantiates a new TemplateVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateVersionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateVersionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateVersionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCanonicalTemplateKey

`func (o *TemplateVersionResponse) GetCanonicalTemplateKey() string`

GetCanonicalTemplateKey returns the CanonicalTemplateKey field if non-nil, zero value otherwise.

### GetCanonicalTemplateKeyOk

`func (o *TemplateVersionResponse) GetCanonicalTemplateKeyOk() (*string, bool)`

GetCanonicalTemplateKeyOk returns a tuple with the CanonicalTemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalTemplateKey

`func (o *TemplateVersionResponse) SetCanonicalTemplateKey(v string)`

SetCanonicalTemplateKey sets CanonicalTemplateKey field to given value.


### GetChannel

`func (o *TemplateVersionResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TemplateVersionResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TemplateVersionResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TemplateVersionResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetVersion

`func (o *TemplateVersionResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TemplateVersionResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TemplateVersionResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *TemplateVersionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateVersionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateVersionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDraftGroupId

`func (o *TemplateVersionResponse) GetDraftGroupId() string`

GetDraftGroupId returns the DraftGroupId field if non-nil, zero value otherwise.

### GetDraftGroupIdOk

`func (o *TemplateVersionResponse) GetDraftGroupIdOk() (*string, bool)`

GetDraftGroupIdOk returns a tuple with the DraftGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftGroupId

`func (o *TemplateVersionResponse) SetDraftGroupId(v string)`

SetDraftGroupId sets DraftGroupId field to given value.

### HasDraftGroupId

`func (o *TemplateVersionResponse) HasDraftGroupId() bool`

HasDraftGroupId returns a boolean if a field has been set.

### GetDraftState

`func (o *TemplateVersionResponse) GetDraftState() string`

GetDraftState returns the DraftState field if non-nil, zero value otherwise.

### GetDraftStateOk

`func (o *TemplateVersionResponse) GetDraftStateOk() (*string, bool)`

GetDraftStateOk returns a tuple with the DraftState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftState

`func (o *TemplateVersionResponse) SetDraftState(v string)`

SetDraftState sets DraftState field to given value.

### HasDraftState

`func (o *TemplateVersionResponse) HasDraftState() bool`

HasDraftState returns a boolean if a field has been set.

### GetStaleReason

`func (o *TemplateVersionResponse) GetStaleReason() string`

GetStaleReason returns the StaleReason field if non-nil, zero value otherwise.

### GetStaleReasonOk

`func (o *TemplateVersionResponse) GetStaleReasonOk() (*string, bool)`

GetStaleReasonOk returns a tuple with the StaleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleReason

`func (o *TemplateVersionResponse) SetStaleReason(v string)`

SetStaleReason sets StaleReason field to given value.

### HasStaleReason

`func (o *TemplateVersionResponse) HasStaleReason() bool`

HasStaleReason returns a boolean if a field has been set.

### GetBasePublishedVersionId

`func (o *TemplateVersionResponse) GetBasePublishedVersionId() string`

GetBasePublishedVersionId returns the BasePublishedVersionId field if non-nil, zero value otherwise.

### GetBasePublishedVersionIdOk

`func (o *TemplateVersionResponse) GetBasePublishedVersionIdOk() (*string, bool)`

GetBasePublishedVersionIdOk returns a tuple with the BasePublishedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePublishedVersionId

`func (o *TemplateVersionResponse) SetBasePublishedVersionId(v string)`

SetBasePublishedVersionId sets BasePublishedVersionId field to given value.

### HasBasePublishedVersionId

`func (o *TemplateVersionResponse) HasBasePublishedVersionId() bool`

HasBasePublishedVersionId returns a boolean if a field has been set.

### GetSubjectTemplate

`func (o *TemplateVersionResponse) GetSubjectTemplate() string`

GetSubjectTemplate returns the SubjectTemplate field if non-nil, zero value otherwise.

### GetSubjectTemplateOk

`func (o *TemplateVersionResponse) GetSubjectTemplateOk() (*string, bool)`

GetSubjectTemplateOk returns a tuple with the SubjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTemplate

`func (o *TemplateVersionResponse) SetSubjectTemplate(v string)`

SetSubjectTemplate sets SubjectTemplate field to given value.


### GetBodyTemplate

`func (o *TemplateVersionResponse) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *TemplateVersionResponse) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *TemplateVersionResponse) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.


### GetContentHash

`func (o *TemplateVersionResponse) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *TemplateVersionResponse) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *TemplateVersionResponse) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.


### GetValidationSnapshot

`func (o *TemplateVersionResponse) GetValidationSnapshot() map[string]interface{}`

GetValidationSnapshot returns the ValidationSnapshot field if non-nil, zero value otherwise.

### GetValidationSnapshotOk

`func (o *TemplateVersionResponse) GetValidationSnapshotOk() (*map[string]interface{}, bool)`

GetValidationSnapshotOk returns a tuple with the ValidationSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationSnapshot

`func (o *TemplateVersionResponse) SetValidationSnapshot(v map[string]interface{})`

SetValidationSnapshot sets ValidationSnapshot field to given value.


### GetDisclosurePolicyHash

`func (o *TemplateVersionResponse) GetDisclosurePolicyHash() string`

GetDisclosurePolicyHash returns the DisclosurePolicyHash field if non-nil, zero value otherwise.

### GetDisclosurePolicyHashOk

`func (o *TemplateVersionResponse) GetDisclosurePolicyHashOk() (*string, bool)`

GetDisclosurePolicyHashOk returns a tuple with the DisclosurePolicyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosurePolicyHash

`func (o *TemplateVersionResponse) SetDisclosurePolicyHash(v string)`

SetDisclosurePolicyHash sets DisclosurePolicyHash field to given value.


### GetProviderReadinessHash

`func (o *TemplateVersionResponse) GetProviderReadinessHash() string`

GetProviderReadinessHash returns the ProviderReadinessHash field if non-nil, zero value otherwise.

### GetProviderReadinessHashOk

`func (o *TemplateVersionResponse) GetProviderReadinessHashOk() (*string, bool)`

GetProviderReadinessHashOk returns a tuple with the ProviderReadinessHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderReadinessHash

`func (o *TemplateVersionResponse) SetProviderReadinessHash(v string)`

SetProviderReadinessHash sets ProviderReadinessHash field to given value.


### GetPaymentInstructionsHash

`func (o *TemplateVersionResponse) GetPaymentInstructionsHash() string`

GetPaymentInstructionsHash returns the PaymentInstructionsHash field if non-nil, zero value otherwise.

### GetPaymentInstructionsHashOk

`func (o *TemplateVersionResponse) GetPaymentInstructionsHashOk() (*string, bool)`

GetPaymentInstructionsHashOk returns a tuple with the PaymentInstructionsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructionsHash

`func (o *TemplateVersionResponse) SetPaymentInstructionsHash(v string)`

SetPaymentInstructionsHash sets PaymentInstructionsHash field to given value.


### GetCreatedByActorId

`func (o *TemplateVersionResponse) GetCreatedByActorId() string`

GetCreatedByActorId returns the CreatedByActorId field if non-nil, zero value otherwise.

### GetCreatedByActorIdOk

`func (o *TemplateVersionResponse) GetCreatedByActorIdOk() (*string, bool)`

GetCreatedByActorIdOk returns a tuple with the CreatedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActorId

`func (o *TemplateVersionResponse) SetCreatedByActorId(v string)`

SetCreatedByActorId sets CreatedByActorId field to given value.


### GetCreatedByActorDisplay

`func (o *TemplateVersionResponse) GetCreatedByActorDisplay() string`

GetCreatedByActorDisplay returns the CreatedByActorDisplay field if non-nil, zero value otherwise.

### GetCreatedByActorDisplayOk

`func (o *TemplateVersionResponse) GetCreatedByActorDisplayOk() (*string, bool)`

GetCreatedByActorDisplayOk returns a tuple with the CreatedByActorDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActorDisplay

`func (o *TemplateVersionResponse) SetCreatedByActorDisplay(v string)`

SetCreatedByActorDisplay sets CreatedByActorDisplay field to given value.

### HasCreatedByActorDisplay

`func (o *TemplateVersionResponse) HasCreatedByActorDisplay() bool`

HasCreatedByActorDisplay returns a boolean if a field has been set.

### GetPublishedByActorId

`func (o *TemplateVersionResponse) GetPublishedByActorId() string`

GetPublishedByActorId returns the PublishedByActorId field if non-nil, zero value otherwise.

### GetPublishedByActorIdOk

`func (o *TemplateVersionResponse) GetPublishedByActorIdOk() (*string, bool)`

GetPublishedByActorIdOk returns a tuple with the PublishedByActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedByActorId

`func (o *TemplateVersionResponse) SetPublishedByActorId(v string)`

SetPublishedByActorId sets PublishedByActorId field to given value.

### HasPublishedByActorId

`func (o *TemplateVersionResponse) HasPublishedByActorId() bool`

HasPublishedByActorId returns a boolean if a field has been set.

### GetPublishedAt

`func (o *TemplateVersionResponse) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *TemplateVersionResponse) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *TemplateVersionResponse) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *TemplateVersionResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetLocked

`func (o *TemplateVersionResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *TemplateVersionResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *TemplateVersionResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *TemplateVersionResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockOwnerActorId

`func (o *TemplateVersionResponse) GetLockOwnerActorId() string`

GetLockOwnerActorId returns the LockOwnerActorId field if non-nil, zero value otherwise.

### GetLockOwnerActorIdOk

`func (o *TemplateVersionResponse) GetLockOwnerActorIdOk() (*string, bool)`

GetLockOwnerActorIdOk returns a tuple with the LockOwnerActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerActorId

`func (o *TemplateVersionResponse) SetLockOwnerActorId(v string)`

SetLockOwnerActorId sets LockOwnerActorId field to given value.

### HasLockOwnerActorId

`func (o *TemplateVersionResponse) HasLockOwnerActorId() bool`

HasLockOwnerActorId returns a boolean if a field has been set.

### GetLockOwnerDisplay

`func (o *TemplateVersionResponse) GetLockOwnerDisplay() string`

GetLockOwnerDisplay returns the LockOwnerDisplay field if non-nil, zero value otherwise.

### GetLockOwnerDisplayOk

`func (o *TemplateVersionResponse) GetLockOwnerDisplayOk() (*string, bool)`

GetLockOwnerDisplayOk returns a tuple with the LockOwnerDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerDisplay

`func (o *TemplateVersionResponse) SetLockOwnerDisplay(v string)`

SetLockOwnerDisplay sets LockOwnerDisplay field to given value.

### HasLockOwnerDisplay

`func (o *TemplateVersionResponse) HasLockOwnerDisplay() bool`

HasLockOwnerDisplay returns a boolean if a field has been set.

### GetLockExpiresAt

`func (o *TemplateVersionResponse) GetLockExpiresAt() string`

GetLockExpiresAt returns the LockExpiresAt field if non-nil, zero value otherwise.

### GetLockExpiresAtOk

`func (o *TemplateVersionResponse) GetLockExpiresAtOk() (*string, bool)`

GetLockExpiresAtOk returns a tuple with the LockExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiresAt

`func (o *TemplateVersionResponse) SetLockExpiresAt(v string)`

SetLockExpiresAt sets LockExpiresAt field to given value.

### HasLockExpiresAt

`func (o *TemplateVersionResponse) HasLockExpiresAt() bool`

HasLockExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


