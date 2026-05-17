# NoteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Author** | **string** |  | 
**NoteType** | Pointer to **string** | general, escalation, dispute, internal, customer_interaction | [optional] [default to "general"]
**AccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewNoteCreate

`func NewNoteCreate(content string, author string, ) *NoteCreate`

NewNoteCreate instantiates a new NoteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteCreateWithDefaults

`func NewNoteCreateWithDefaults() *NoteCreate`

NewNoteCreateWithDefaults instantiates a new NoteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *NoteCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteCreate) SetContent(v string)`

SetContent sets Content field to given value.


### GetAuthor

`func (o *NoteCreate) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NoteCreate) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NoteCreate) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetNoteType

`func (o *NoteCreate) GetNoteType() string`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *NoteCreate) GetNoteTypeOk() (*string, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *NoteCreate) SetNoteType(v string)`

SetNoteType sets NoteType field to given value.

### HasNoteType

`func (o *NoteCreate) HasNoteType() bool`

HasNoteType returns a boolean if a field has been set.

### GetAccountId

`func (o *NoteCreate) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NoteCreate) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NoteCreate) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *NoteCreate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


