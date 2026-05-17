# DryRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanAttributes** | **map[string]interface{}** |  | 
**Headers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDryRunRequest

`func NewDryRunRequest(spanAttributes map[string]interface{}, ) *DryRunRequest`

NewDryRunRequest instantiates a new DryRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunRequestWithDefaults

`func NewDryRunRequestWithDefaults() *DryRunRequest`

NewDryRunRequestWithDefaults instantiates a new DryRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanAttributes

`func (o *DryRunRequest) GetSpanAttributes() map[string]interface{}`

GetSpanAttributes returns the SpanAttributes field if non-nil, zero value otherwise.

### GetSpanAttributesOk

`func (o *DryRunRequest) GetSpanAttributesOk() (*map[string]interface{}, bool)`

GetSpanAttributesOk returns a tuple with the SpanAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanAttributes

`func (o *DryRunRequest) SetSpanAttributes(v map[string]interface{})`

SetSpanAttributes sets SpanAttributes field to given value.


### GetHeaders

`func (o *DryRunRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DryRunRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DryRunRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DryRunRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


