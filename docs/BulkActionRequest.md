# BulkActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action: pause, resume, escalate, assign | 
**CaseIds** | **[]string** |  | 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBulkActionRequest

`func NewBulkActionRequest(action string, caseIds []string, ) *BulkActionRequest`

NewBulkActionRequest instantiates a new BulkActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkActionRequestWithDefaults

`func NewBulkActionRequestWithDefaults() *BulkActionRequest`

NewBulkActionRequestWithDefaults instantiates a new BulkActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetCaseIds

`func (o *BulkActionRequest) GetCaseIds() []string`

GetCaseIds returns the CaseIds field if non-nil, zero value otherwise.

### GetCaseIdsOk

`func (o *BulkActionRequest) GetCaseIdsOk() (*[]string, bool)`

GetCaseIdsOk returns a tuple with the CaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIds

`func (o *BulkActionRequest) SetCaseIds(v []string)`

SetCaseIds sets CaseIds field to given value.


### GetParams

`func (o *BulkActionRequest) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *BulkActionRequest) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *BulkActionRequest) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *BulkActionRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


