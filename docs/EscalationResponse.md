# EscalationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CaseId** | **string** |  | 
**Reason** | **string** |  | 
**EscalatedTo** | Pointer to **string** |  | [optional] 
**EscalatedFrom** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Priority** | **string** |  | 
**SlaDeadline** | Pointer to **time.Time** |  | [optional] 
**ResolvedAt** | Pointer to **time.Time** |  | [optional] 
**ResolutionNotes** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewEscalationResponse

`func NewEscalationResponse(id string, tenantId string, caseId string, reason string, status string, priority string, createdAt time.Time, updatedAt time.Time, ) *EscalationResponse`

NewEscalationResponse instantiates a new EscalationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationResponseWithDefaults

`func NewEscalationResponseWithDefaults() *EscalationResponse`

NewEscalationResponseWithDefaults instantiates a new EscalationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EscalationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EscalationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EscalationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *EscalationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EscalationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EscalationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCaseId

`func (o *EscalationResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *EscalationResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *EscalationResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetReason

`func (o *EscalationResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EscalationResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EscalationResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetEscalatedTo

`func (o *EscalationResponse) GetEscalatedTo() string`

GetEscalatedTo returns the EscalatedTo field if non-nil, zero value otherwise.

### GetEscalatedToOk

`func (o *EscalationResponse) GetEscalatedToOk() (*string, bool)`

GetEscalatedToOk returns a tuple with the EscalatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalatedTo

`func (o *EscalationResponse) SetEscalatedTo(v string)`

SetEscalatedTo sets EscalatedTo field to given value.

### HasEscalatedTo

`func (o *EscalationResponse) HasEscalatedTo() bool`

HasEscalatedTo returns a boolean if a field has been set.

### GetEscalatedFrom

`func (o *EscalationResponse) GetEscalatedFrom() string`

GetEscalatedFrom returns the EscalatedFrom field if non-nil, zero value otherwise.

### GetEscalatedFromOk

`func (o *EscalationResponse) GetEscalatedFromOk() (*string, bool)`

GetEscalatedFromOk returns a tuple with the EscalatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalatedFrom

`func (o *EscalationResponse) SetEscalatedFrom(v string)`

SetEscalatedFrom sets EscalatedFrom field to given value.

### HasEscalatedFrom

`func (o *EscalationResponse) HasEscalatedFrom() bool`

HasEscalatedFrom returns a boolean if a field has been set.

### GetStatus

`func (o *EscalationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EscalationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EscalationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *EscalationResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EscalationResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EscalationResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetSlaDeadline

`func (o *EscalationResponse) GetSlaDeadline() time.Time`

GetSlaDeadline returns the SlaDeadline field if non-nil, zero value otherwise.

### GetSlaDeadlineOk

`func (o *EscalationResponse) GetSlaDeadlineOk() (*time.Time, bool)`

GetSlaDeadlineOk returns a tuple with the SlaDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaDeadline

`func (o *EscalationResponse) SetSlaDeadline(v time.Time)`

SetSlaDeadline sets SlaDeadline field to given value.

### HasSlaDeadline

`func (o *EscalationResponse) HasSlaDeadline() bool`

HasSlaDeadline returns a boolean if a field has been set.

### GetResolvedAt

`func (o *EscalationResponse) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *EscalationResponse) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *EscalationResponse) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *EscalationResponse) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetResolutionNotes

`func (o *EscalationResponse) GetResolutionNotes() string`

GetResolutionNotes returns the ResolutionNotes field if non-nil, zero value otherwise.

### GetResolutionNotesOk

`func (o *EscalationResponse) GetResolutionNotesOk() (*string, bool)`

GetResolutionNotesOk returns a tuple with the ResolutionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionNotes

`func (o *EscalationResponse) SetResolutionNotes(v string)`

SetResolutionNotes sets ResolutionNotes field to given value.

### HasResolutionNotes

`func (o *EscalationResponse) HasResolutionNotes() bool`

HasResolutionNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EscalationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EscalationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EscalationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EscalationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EscalationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EscalationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


