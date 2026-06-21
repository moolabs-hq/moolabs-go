# ContractUploadDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Filename** | **string** |  | 
**PageCount** | **int32** |  | 
**ExtractedText** | **string** |  | 
**Segments** | [**[]ContractSegmentResponse**](ContractSegmentResponse.md) |  | 

## Methods

### NewContractUploadDetailResponse

`func NewContractUploadDetailResponse(id string, filename string, pageCount int32, extractedText string, segments []ContractSegmentResponse, ) *ContractUploadDetailResponse`

NewContractUploadDetailResponse instantiates a new ContractUploadDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractUploadDetailResponseWithDefaults

`func NewContractUploadDetailResponseWithDefaults() *ContractUploadDetailResponse`

NewContractUploadDetailResponseWithDefaults instantiates a new ContractUploadDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractUploadDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractUploadDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractUploadDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetFilename

`func (o *ContractUploadDetailResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ContractUploadDetailResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ContractUploadDetailResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetPageCount

`func (o *ContractUploadDetailResponse) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ContractUploadDetailResponse) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ContractUploadDetailResponse) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetExtractedText

`func (o *ContractUploadDetailResponse) GetExtractedText() string`

GetExtractedText returns the ExtractedText field if non-nil, zero value otherwise.

### GetExtractedTextOk

`func (o *ContractUploadDetailResponse) GetExtractedTextOk() (*string, bool)`

GetExtractedTextOk returns a tuple with the ExtractedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedText

`func (o *ContractUploadDetailResponse) SetExtractedText(v string)`

SetExtractedText sets ExtractedText field to given value.


### GetSegments

`func (o *ContractUploadDetailResponse) GetSegments() []ContractSegmentResponse`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ContractUploadDetailResponse) GetSegmentsOk() (*[]ContractSegmentResponse, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ContractUploadDetailResponse) SetSegments(v []ContractSegmentResponse)`

SetSegments sets Segments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


