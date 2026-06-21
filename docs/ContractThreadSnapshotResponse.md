# ContractThreadSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thread** | [**ContractThreadResponse**](ContractThreadResponse.md) |  | 
**Participants** | [**[]ContractParticipantResponse**](ContractParticipantResponse.md) |  | 
**Documents** | [**[]ContractDocumentResponse**](ContractDocumentResponse.md) |  | 
**Events** | [**[]ContractEventResponse**](ContractEventResponse.md) |  | 

## Methods

### NewContractThreadSnapshotResponse

`func NewContractThreadSnapshotResponse(thread ContractThreadResponse, participants []ContractParticipantResponse, documents []ContractDocumentResponse, events []ContractEventResponse, ) *ContractThreadSnapshotResponse`

NewContractThreadSnapshotResponse instantiates a new ContractThreadSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractThreadSnapshotResponseWithDefaults

`func NewContractThreadSnapshotResponseWithDefaults() *ContractThreadSnapshotResponse`

NewContractThreadSnapshotResponseWithDefaults instantiates a new ContractThreadSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThread

`func (o *ContractThreadSnapshotResponse) GetThread() ContractThreadResponse`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *ContractThreadSnapshotResponse) GetThreadOk() (*ContractThreadResponse, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *ContractThreadSnapshotResponse) SetThread(v ContractThreadResponse)`

SetThread sets Thread field to given value.


### GetParticipants

`func (o *ContractThreadSnapshotResponse) GetParticipants() []ContractParticipantResponse`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ContractThreadSnapshotResponse) GetParticipantsOk() (*[]ContractParticipantResponse, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ContractThreadSnapshotResponse) SetParticipants(v []ContractParticipantResponse)`

SetParticipants sets Participants field to given value.


### GetDocuments

`func (o *ContractThreadSnapshotResponse) GetDocuments() []ContractDocumentResponse`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ContractThreadSnapshotResponse) GetDocumentsOk() (*[]ContractDocumentResponse, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ContractThreadSnapshotResponse) SetDocuments(v []ContractDocumentResponse)`

SetDocuments sets Documents field to given value.


### GetEvents

`func (o *ContractThreadSnapshotResponse) GetEvents() []ContractEventResponse`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ContractThreadSnapshotResponse) GetEventsOk() (*[]ContractEventResponse, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ContractThreadSnapshotResponse) SetEvents(v []ContractEventResponse)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


