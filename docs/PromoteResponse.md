# PromoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromotedRows** | **int32** |  | 
**ShadowRunId** | **string** |  | 

## Methods

### NewPromoteResponse

`func NewPromoteResponse(promotedRows int32, shadowRunId string, ) *PromoteResponse`

NewPromoteResponse instantiates a new PromoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoteResponseWithDefaults

`func NewPromoteResponseWithDefaults() *PromoteResponse`

NewPromoteResponseWithDefaults instantiates a new PromoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromotedRows

`func (o *PromoteResponse) GetPromotedRows() int32`

GetPromotedRows returns the PromotedRows field if non-nil, zero value otherwise.

### GetPromotedRowsOk

`func (o *PromoteResponse) GetPromotedRowsOk() (*int32, bool)`

GetPromotedRowsOk returns a tuple with the PromotedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedRows

`func (o *PromoteResponse) SetPromotedRows(v int32)`

SetPromotedRows sets PromotedRows field to given value.


### GetShadowRunId

`func (o *PromoteResponse) GetShadowRunId() string`

GetShadowRunId returns the ShadowRunId field if non-nil, zero value otherwise.

### GetShadowRunIdOk

`func (o *PromoteResponse) GetShadowRunIdOk() (*string, bool)`

GetShadowRunIdOk returns a tuple with the ShadowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowRunId

`func (o *PromoteResponse) SetShadowRunId(v string)`

SetShadowRunId sets ShadowRunId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


