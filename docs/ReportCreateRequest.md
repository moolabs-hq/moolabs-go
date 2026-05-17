# ReportCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** |  | 
**Format** | **string** |  | 
**AllowEmpty** | Pointer to **bool** |  | [optional] [default to false]
**Filters** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewReportCreateRequest

`func NewReportCreateRequest(reportType string, format string, ) *ReportCreateRequest`

NewReportCreateRequest instantiates a new ReportCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCreateRequestWithDefaults

`func NewReportCreateRequestWithDefaults() *ReportCreateRequest`

NewReportCreateRequestWithDefaults instantiates a new ReportCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *ReportCreateRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportCreateRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportCreateRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetFormat

`func (o *ReportCreateRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ReportCreateRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ReportCreateRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetAllowEmpty

`func (o *ReportCreateRequest) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *ReportCreateRequest) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *ReportCreateRequest) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *ReportCreateRequest) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetFilters

`func (o *ReportCreateRequest) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReportCreateRequest) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReportCreateRequest) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ReportCreateRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSource

`func (o *ReportCreateRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReportCreateRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReportCreateRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReportCreateRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


