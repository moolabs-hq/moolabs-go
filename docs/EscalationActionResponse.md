# EscalationActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewEscalationActionResponse

`func NewEscalationActionResponse(ok bool, ) *EscalationActionResponse`

NewEscalationActionResponse instantiates a new EscalationActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEscalationActionResponseWithDefaults

`func NewEscalationActionResponseWithDefaults() *EscalationActionResponse`

NewEscalationActionResponseWithDefaults instantiates a new EscalationActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *EscalationActionResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *EscalationActionResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *EscalationActionResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetResult

`func (o *EscalationActionResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EscalationActionResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EscalationActionResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *EscalationActionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *EscalationActionResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EscalationActionResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EscalationActionResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EscalationActionResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


