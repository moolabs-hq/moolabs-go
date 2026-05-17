# ReportCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] [default to true]
**ReportId** | **string** |  | 
**Status** | **string** |  | 
**ReportType** | **string** |  | 
**RequestedFormat** | **string** |  | 
**ActualFormat** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**DownloadUrl** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Summary** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewReportCreateResponse

`func NewReportCreateResponse(reportId string, status string, reportType string, requestedFormat string, ) *ReportCreateResponse`

NewReportCreateResponse instantiates a new ReportCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCreateResponseWithDefaults

`func NewReportCreateResponseWithDefaults() *ReportCreateResponse`

NewReportCreateResponseWithDefaults instantiates a new ReportCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *ReportCreateResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ReportCreateResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ReportCreateResponse) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *ReportCreateResponse) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetReportId

`func (o *ReportCreateResponse) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ReportCreateResponse) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ReportCreateResponse) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetStatus

`func (o *ReportCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReportType

`func (o *ReportCreateResponse) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportCreateResponse) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportCreateResponse) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetRequestedFormat

`func (o *ReportCreateResponse) GetRequestedFormat() string`

GetRequestedFormat returns the RequestedFormat field if non-nil, zero value otherwise.

### GetRequestedFormatOk

`func (o *ReportCreateResponse) GetRequestedFormatOk() (*string, bool)`

GetRequestedFormatOk returns a tuple with the RequestedFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFormat

`func (o *ReportCreateResponse) SetRequestedFormat(v string)`

SetRequestedFormat sets RequestedFormat field to given value.


### GetActualFormat

`func (o *ReportCreateResponse) GetActualFormat() string`

GetActualFormat returns the ActualFormat field if non-nil, zero value otherwise.

### GetActualFormatOk

`func (o *ReportCreateResponse) GetActualFormatOk() (*string, bool)`

GetActualFormatOk returns a tuple with the ActualFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFormat

`func (o *ReportCreateResponse) SetActualFormat(v string)`

SetActualFormat sets ActualFormat field to given value.

### HasActualFormat

`func (o *ReportCreateResponse) HasActualFormat() bool`

HasActualFormat returns a boolean if a field has been set.

### GetFilename

`func (o *ReportCreateResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ReportCreateResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ReportCreateResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ReportCreateResponse) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *ReportCreateResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ReportCreateResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ReportCreateResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *ReportCreateResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *ReportCreateResponse) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *ReportCreateResponse) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *ReportCreateResponse) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *ReportCreateResponse) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ReportCreateResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ReportCreateResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ReportCreateResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ReportCreateResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSummary

`func (o *ReportCreateResponse) GetSummary() map[string]interface{}`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportCreateResponse) GetSummaryOk() (*map[string]interface{}, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportCreateResponse) SetSummary(v map[string]interface{})`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ReportCreateResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


