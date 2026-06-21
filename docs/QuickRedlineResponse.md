# QuickRedlineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** |  | 
**AuditId** | **string** |  | 
**Filename** | **string** |  | 
**PageCount** | **int32** |  | 
**ExtractedText** | **string** |  | 
**Segments** | [**[]SegmentOut**](SegmentOut.md) |  | 
**Findings** | [**[]ClauseFinding**](ClauseFinding.md) |  | 
**Divergences** | **[]interface{}** |  | 
**Summary** | [**RedlineSummary**](RedlineSummary.md) |  | 

## Methods

### NewQuickRedlineResponse

`func NewQuickRedlineResponse(uploadId string, auditId string, filename string, pageCount int32, extractedText string, segments []SegmentOut, findings []ClauseFinding, divergences []interface{}, summary RedlineSummary, ) *QuickRedlineResponse`

NewQuickRedlineResponse instantiates a new QuickRedlineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickRedlineResponseWithDefaults

`func NewQuickRedlineResponseWithDefaults() *QuickRedlineResponse`

NewQuickRedlineResponseWithDefaults instantiates a new QuickRedlineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *QuickRedlineResponse) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *QuickRedlineResponse) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *QuickRedlineResponse) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetAuditId

`func (o *QuickRedlineResponse) GetAuditId() string`

GetAuditId returns the AuditId field if non-nil, zero value otherwise.

### GetAuditIdOk

`func (o *QuickRedlineResponse) GetAuditIdOk() (*string, bool)`

GetAuditIdOk returns a tuple with the AuditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditId

`func (o *QuickRedlineResponse) SetAuditId(v string)`

SetAuditId sets AuditId field to given value.


### GetFilename

`func (o *QuickRedlineResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *QuickRedlineResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *QuickRedlineResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPageCount

`func (o *QuickRedlineResponse) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *QuickRedlineResponse) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *QuickRedlineResponse) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetExtractedText

`func (o *QuickRedlineResponse) GetExtractedText() string`

GetExtractedText returns the ExtractedText field if non-nil, zero value otherwise.

### GetExtractedTextOk

`func (o *QuickRedlineResponse) GetExtractedTextOk() (*string, bool)`

GetExtractedTextOk returns a tuple with the ExtractedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedText

`func (o *QuickRedlineResponse) SetExtractedText(v string)`

SetExtractedText sets ExtractedText field to given value.


### GetSegments

`func (o *QuickRedlineResponse) GetSegments() []SegmentOut`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *QuickRedlineResponse) GetSegmentsOk() (*[]SegmentOut, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *QuickRedlineResponse) SetSegments(v []SegmentOut)`

SetSegments sets Segments field to given value.


### GetFindings

`func (o *QuickRedlineResponse) GetFindings() []ClauseFinding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *QuickRedlineResponse) GetFindingsOk() (*[]ClauseFinding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *QuickRedlineResponse) SetFindings(v []ClauseFinding)`

SetFindings sets Findings field to given value.


### GetDivergences

`func (o *QuickRedlineResponse) GetDivergences() []interface{}`

GetDivergences returns the Divergences field if non-nil, zero value otherwise.

### GetDivergencesOk

`func (o *QuickRedlineResponse) GetDivergencesOk() (*[]interface{}, bool)`

GetDivergencesOk returns a tuple with the Divergences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivergences

`func (o *QuickRedlineResponse) SetDivergences(v []interface{})`

SetDivergences sets Divergences field to given value.


### GetSummary

`func (o *QuickRedlineResponse) GetSummary() RedlineSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *QuickRedlineResponse) GetSummaryOk() (*RedlineSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *QuickRedlineResponse) SetSummary(v RedlineSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


