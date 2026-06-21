# AskMooRedlineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FindingId** | **string** |  | 
**FamilyId** | Pointer to **string** |  | [optional] 
**AnchorCharStart** | **int32** |  | 
**AnchorCharEnd** | **int32** |  | 
**Instruction** | **string** |  | 
**AnchorQuote** | Pointer to **string** |  | [optional] 

## Methods

### NewAskMooRedlineRequest

`func NewAskMooRedlineRequest(findingId string, anchorCharStart int32, anchorCharEnd int32, instruction string, ) *AskMooRedlineRequest`

NewAskMooRedlineRequest instantiates a new AskMooRedlineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskMooRedlineRequestWithDefaults

`func NewAskMooRedlineRequestWithDefaults() *AskMooRedlineRequest`

NewAskMooRedlineRequestWithDefaults instantiates a new AskMooRedlineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindingId

`func (o *AskMooRedlineRequest) GetFindingId() string`

GetFindingId returns the FindingId field if non-nil, zero value otherwise.

### GetFindingIdOk

`func (o *AskMooRedlineRequest) GetFindingIdOk() (*string, bool)`

GetFindingIdOk returns a tuple with the FindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingId

`func (o *AskMooRedlineRequest) SetFindingId(v string)`

SetFindingId sets FindingId field to given value.


### GetFamilyId

`func (o *AskMooRedlineRequest) GetFamilyId() string`

GetFamilyId returns the FamilyId field if non-nil, zero value otherwise.

### GetFamilyIdOk

`func (o *AskMooRedlineRequest) GetFamilyIdOk() (*string, bool)`

GetFamilyIdOk returns a tuple with the FamilyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyId

`func (o *AskMooRedlineRequest) SetFamilyId(v string)`

SetFamilyId sets FamilyId field to given value.

### HasFamilyId

`func (o *AskMooRedlineRequest) HasFamilyId() bool`

HasFamilyId returns a boolean if a field has been set.

### GetAnchorCharStart

`func (o *AskMooRedlineRequest) GetAnchorCharStart() int32`

GetAnchorCharStart returns the AnchorCharStart field if non-nil, zero value otherwise.

### GetAnchorCharStartOk

`func (o *AskMooRedlineRequest) GetAnchorCharStartOk() (*int32, bool)`

GetAnchorCharStartOk returns a tuple with the AnchorCharStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCharStart

`func (o *AskMooRedlineRequest) SetAnchorCharStart(v int32)`

SetAnchorCharStart sets AnchorCharStart field to given value.


### GetAnchorCharEnd

`func (o *AskMooRedlineRequest) GetAnchorCharEnd() int32`

GetAnchorCharEnd returns the AnchorCharEnd field if non-nil, zero value otherwise.

### GetAnchorCharEndOk

`func (o *AskMooRedlineRequest) GetAnchorCharEndOk() (*int32, bool)`

GetAnchorCharEndOk returns a tuple with the AnchorCharEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCharEnd

`func (o *AskMooRedlineRequest) SetAnchorCharEnd(v int32)`

SetAnchorCharEnd sets AnchorCharEnd field to given value.


### GetInstruction

`func (o *AskMooRedlineRequest) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *AskMooRedlineRequest) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *AskMooRedlineRequest) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.


### GetAnchorQuote

`func (o *AskMooRedlineRequest) GetAnchorQuote() string`

GetAnchorQuote returns the AnchorQuote field if non-nil, zero value otherwise.

### GetAnchorQuoteOk

`func (o *AskMooRedlineRequest) GetAnchorQuoteOk() (*string, bool)`

GetAnchorQuoteOk returns a tuple with the AnchorQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorQuote

`func (o *AskMooRedlineRequest) SetAnchorQuote(v string)`

SetAnchorQuote sets AnchorQuote field to given value.

### HasAnchorQuote

`func (o *AskMooRedlineRequest) HasAnchorQuote() bool`

HasAnchorQuote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


