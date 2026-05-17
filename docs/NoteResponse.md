# NoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CaseId** | **string** |  | 
**AccountId** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**Author** | **string** |  | 
**NoteType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewNoteResponse

`func NewNoteResponse(id string, tenantId string, caseId string, content string, author string, noteType string, createdAt time.Time, updatedAt time.Time, ) *NoteResponse`

NewNoteResponse instantiates a new NoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteResponseWithDefaults

`func NewNoteResponseWithDefaults() *NoteResponse`

NewNoteResponseWithDefaults instantiates a new NoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *NoteResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NoteResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NoteResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCaseId

`func (o *NoteResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *NoteResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *NoteResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetAccountId

`func (o *NoteResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NoteResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NoteResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *NoteResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetContent

`func (o *NoteResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetAuthor

`func (o *NoteResponse) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NoteResponse) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NoteResponse) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetNoteType

`func (o *NoteResponse) GetNoteType() string`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *NoteResponse) GetNoteTypeOk() (*string, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *NoteResponse) SetNoteType(v string)`

SetNoteType sets NoteType field to given value.


### GetCreatedAt

`func (o *NoteResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NoteResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NoteResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NoteResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NoteResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NoteResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


