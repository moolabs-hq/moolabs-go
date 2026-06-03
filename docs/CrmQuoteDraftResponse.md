# CrmQuoteDraftResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**DraftStatus** | **string** |  | 
**ConfirmationRequired** | **bool** |  | 
**SurfaceResponse** | [**CrmSurfaceResponse**](CrmSurfaceResponse.md) |  | 
**OpenQuoteUrl** | **string** |  | 
**SourceEvidenceId** | **string** |  | 
**SurfaceSessionId** | **string** |  | 
**Command** | Pointer to **map[string]interface{}** |  | [optional] 
**Card** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCrmQuoteDraftResponse

`func NewCrmQuoteDraftResponse(ok bool, draftStatus string, confirmationRequired bool, surfaceResponse CrmSurfaceResponse, openQuoteUrl string, sourceEvidenceId string, surfaceSessionId string, ) *CrmQuoteDraftResponse`

NewCrmQuoteDraftResponse instantiates a new CrmQuoteDraftResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmQuoteDraftResponseWithDefaults

`func NewCrmQuoteDraftResponseWithDefaults() *CrmQuoteDraftResponse`

NewCrmQuoteDraftResponseWithDefaults instantiates a new CrmQuoteDraftResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *CrmQuoteDraftResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CrmQuoteDraftResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CrmQuoteDraftResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetDraftStatus

`func (o *CrmQuoteDraftResponse) GetDraftStatus() string`

GetDraftStatus returns the DraftStatus field if non-nil, zero value otherwise.

### GetDraftStatusOk

`func (o *CrmQuoteDraftResponse) GetDraftStatusOk() (*string, bool)`

GetDraftStatusOk returns a tuple with the DraftStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftStatus

`func (o *CrmQuoteDraftResponse) SetDraftStatus(v string)`

SetDraftStatus sets DraftStatus field to given value.


### GetConfirmationRequired

`func (o *CrmQuoteDraftResponse) GetConfirmationRequired() bool`

GetConfirmationRequired returns the ConfirmationRequired field if non-nil, zero value otherwise.

### GetConfirmationRequiredOk

`func (o *CrmQuoteDraftResponse) GetConfirmationRequiredOk() (*bool, bool)`

GetConfirmationRequiredOk returns a tuple with the ConfirmationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationRequired

`func (o *CrmQuoteDraftResponse) SetConfirmationRequired(v bool)`

SetConfirmationRequired sets ConfirmationRequired field to given value.


### GetSurfaceResponse

`func (o *CrmQuoteDraftResponse) GetSurfaceResponse() CrmSurfaceResponse`

GetSurfaceResponse returns the SurfaceResponse field if non-nil, zero value otherwise.

### GetSurfaceResponseOk

`func (o *CrmQuoteDraftResponse) GetSurfaceResponseOk() (*CrmSurfaceResponse, bool)`

GetSurfaceResponseOk returns a tuple with the SurfaceResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurfaceResponse

`func (o *CrmQuoteDraftResponse) SetSurfaceResponse(v CrmSurfaceResponse)`

SetSurfaceResponse sets SurfaceResponse field to given value.


### GetOpenQuoteUrl

`func (o *CrmQuoteDraftResponse) GetOpenQuoteUrl() string`

GetOpenQuoteUrl returns the OpenQuoteUrl field if non-nil, zero value otherwise.

### GetOpenQuoteUrlOk

`func (o *CrmQuoteDraftResponse) GetOpenQuoteUrlOk() (*string, bool)`

GetOpenQuoteUrlOk returns a tuple with the OpenQuoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenQuoteUrl

`func (o *CrmQuoteDraftResponse) SetOpenQuoteUrl(v string)`

SetOpenQuoteUrl sets OpenQuoteUrl field to given value.


### GetSourceEvidenceId

`func (o *CrmQuoteDraftResponse) GetSourceEvidenceId() string`

GetSourceEvidenceId returns the SourceEvidenceId field if non-nil, zero value otherwise.

### GetSourceEvidenceIdOk

`func (o *CrmQuoteDraftResponse) GetSourceEvidenceIdOk() (*string, bool)`

GetSourceEvidenceIdOk returns a tuple with the SourceEvidenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEvidenceId

`func (o *CrmQuoteDraftResponse) SetSourceEvidenceId(v string)`

SetSourceEvidenceId sets SourceEvidenceId field to given value.


### GetSurfaceSessionId

`func (o *CrmQuoteDraftResponse) GetSurfaceSessionId() string`

GetSurfaceSessionId returns the SurfaceSessionId field if non-nil, zero value otherwise.

### GetSurfaceSessionIdOk

`func (o *CrmQuoteDraftResponse) GetSurfaceSessionIdOk() (*string, bool)`

GetSurfaceSessionIdOk returns a tuple with the SurfaceSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurfaceSessionId

`func (o *CrmQuoteDraftResponse) SetSurfaceSessionId(v string)`

SetSurfaceSessionId sets SurfaceSessionId field to given value.


### GetCommand

`func (o *CrmQuoteDraftResponse) GetCommand() map[string]interface{}`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CrmQuoteDraftResponse) GetCommandOk() (*map[string]interface{}, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CrmQuoteDraftResponse) SetCommand(v map[string]interface{})`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CrmQuoteDraftResponse) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCard

`func (o *CrmQuoteDraftResponse) GetCard() map[string]interface{}`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CrmQuoteDraftResponse) GetCardOk() (*map[string]interface{}, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CrmQuoteDraftResponse) SetCard(v map[string]interface{})`

SetCard sets Card field to given value.

### HasCard

`func (o *CrmQuoteDraftResponse) HasCard() bool`

HasCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


