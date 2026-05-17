# RateBulkImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inserted** | **int32** |  | 
**Skipped** | **int32** |  | 
**Details** | **[]map[string]interface{}** |  | 

## Methods

### NewRateBulkImportResponse

`func NewRateBulkImportResponse(inserted int32, skipped int32, details []map[string]interface{}, ) *RateBulkImportResponse`

NewRateBulkImportResponse instantiates a new RateBulkImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateBulkImportResponseWithDefaults

`func NewRateBulkImportResponseWithDefaults() *RateBulkImportResponse`

NewRateBulkImportResponseWithDefaults instantiates a new RateBulkImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInserted

`func (o *RateBulkImportResponse) GetInserted() int32`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *RateBulkImportResponse) GetInsertedOk() (*int32, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *RateBulkImportResponse) SetInserted(v int32)`

SetInserted sets Inserted field to given value.


### GetSkipped

`func (o *RateBulkImportResponse) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *RateBulkImportResponse) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *RateBulkImportResponse) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.


### GetDetails

`func (o *RateBulkImportResponse) GetDetails() []map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RateBulkImportResponse) GetDetailsOk() (*[]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RateBulkImportResponse) SetDetails(v []map[string]interface{})`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


