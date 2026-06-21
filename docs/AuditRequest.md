# AuditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndustryVertical** | **string** |  | 
**ContractUploadIds** | **[]string** |  | 
**KbPositions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAuditRequest

`func NewAuditRequest(industryVertical string, contractUploadIds []string, ) *AuditRequest`

NewAuditRequest instantiates a new AuditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditRequestWithDefaults

`func NewAuditRequestWithDefaults() *AuditRequest`

NewAuditRequestWithDefaults instantiates a new AuditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndustryVertical

`func (o *AuditRequest) GetIndustryVertical() string`

GetIndustryVertical returns the IndustryVertical field if non-nil, zero value otherwise.

### GetIndustryVerticalOk

`func (o *AuditRequest) GetIndustryVerticalOk() (*string, bool)`

GetIndustryVerticalOk returns a tuple with the IndustryVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryVertical

`func (o *AuditRequest) SetIndustryVertical(v string)`

SetIndustryVertical sets IndustryVertical field to given value.


### GetContractUploadIds

`func (o *AuditRequest) GetContractUploadIds() []string`

GetContractUploadIds returns the ContractUploadIds field if non-nil, zero value otherwise.

### GetContractUploadIdsOk

`func (o *AuditRequest) GetContractUploadIdsOk() (*[]string, bool)`

GetContractUploadIdsOk returns a tuple with the ContractUploadIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractUploadIds

`func (o *AuditRequest) SetContractUploadIds(v []string)`

SetContractUploadIds sets ContractUploadIds field to given value.


### GetKbPositions

`func (o *AuditRequest) GetKbPositions() map[string]string`

GetKbPositions returns the KbPositions field if non-nil, zero value otherwise.

### GetKbPositionsOk

`func (o *AuditRequest) GetKbPositionsOk() (*map[string]string, bool)`

GetKbPositionsOk returns a tuple with the KbPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbPositions

`func (o *AuditRequest) SetKbPositions(v map[string]string)`

SetKbPositions sets KbPositions field to given value.

### HasKbPositions

`func (o *AuditRequest) HasKbPositions() bool`

HasKbPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


