# TestSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**Mode** | Pointer to **string** |  | [optional] 
**CaseId** | Pointer to **string** |  | [optional] 
**RecipientEmail** | **string** |  | 

## Methods

### NewTestSendRequest

`func NewTestSendRequest(versionId string, recipientEmail string, ) *TestSendRequest`

NewTestSendRequest instantiates a new TestSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSendRequestWithDefaults

`func NewTestSendRequestWithDefaults() *TestSendRequest`

NewTestSendRequestWithDefaults instantiates a new TestSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *TestSendRequest) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TestSendRequest) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TestSendRequest) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetMode

`func (o *TestSendRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TestSendRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TestSendRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TestSendRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCaseId

`func (o *TestSendRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *TestSendRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *TestSendRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *TestSendRequest) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetRecipientEmail

`func (o *TestSendRequest) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *TestSendRequest) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *TestSendRequest) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


