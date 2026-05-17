# SDKSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanId** | **string** |  | 
**Provider** | **string** |  | 
**Model** | **string** |  | 
**OperationType** | **string** |  | 
**InputTokens** | Pointer to **int32** |  | [optional] 
**OutputTokens** | Pointer to **int32** |  | [optional] 
**ToolCalls** | Pointer to **int32** |  | [optional] 
**LatencyMs** | Pointer to **int32** |  | [optional] 
**CacheHit** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSDKSpan

`func NewSDKSpan(spanId string, provider string, model string, operationType string, ) *SDKSpan`

NewSDKSpan instantiates a new SDKSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDKSpanWithDefaults

`func NewSDKSpanWithDefaults() *SDKSpan`

NewSDKSpanWithDefaults instantiates a new SDKSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanId

`func (o *SDKSpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SDKSpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SDKSpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetProvider

`func (o *SDKSpan) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SDKSpan) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SDKSpan) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *SDKSpan) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SDKSpan) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SDKSpan) SetModel(v string)`

SetModel sets Model field to given value.


### GetOperationType

`func (o *SDKSpan) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SDKSpan) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SDKSpan) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetInputTokens

`func (o *SDKSpan) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *SDKSpan) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *SDKSpan) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *SDKSpan) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### GetOutputTokens

`func (o *SDKSpan) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *SDKSpan) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *SDKSpan) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *SDKSpan) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### GetToolCalls

`func (o *SDKSpan) GetToolCalls() int32`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *SDKSpan) GetToolCallsOk() (*int32, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *SDKSpan) SetToolCalls(v int32)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *SDKSpan) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetLatencyMs

`func (o *SDKSpan) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *SDKSpan) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *SDKSpan) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *SDKSpan) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetCacheHit

`func (o *SDKSpan) GetCacheHit() bool`

GetCacheHit returns the CacheHit field if non-nil, zero value otherwise.

### GetCacheHitOk

`func (o *SDKSpan) GetCacheHitOk() (*bool, bool)`

GetCacheHitOk returns a tuple with the CacheHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHit

`func (o *SDKSpan) SetCacheHit(v bool)`

SetCacheHit sets CacheHit field to given value.

### HasCacheHit

`func (o *SDKSpan) HasCacheHit() bool`

HasCacheHit returns a boolean if a field has been set.

### GetMetadata

`func (o *SDKSpan) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SDKSpan) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SDKSpan) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SDKSpan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


