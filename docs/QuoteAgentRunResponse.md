# QuoteAgentRunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**QuoteId** | **string** |  | 
**QuoteVersion** | **int32** |  | 
**AgentName** | **string** |  | 
**TrustTier** | **int32** |  | 
**Surface** | **string** |  | 
**ActorType** | **string** |  | 
**ActorId** | **string** |  | 
**Status** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**CompletedAt** | **time.Time** |  | 
**InputHash** | **string** |  | 
**OutputHash** | **string** |  | 
**RationaleSummary** | **string** |  | 
**ToolTrace** | **[]map[string]interface{}** |  | 
**ErrorCode** | **string** |  | 
**TraceId** | **string** |  | 

## Methods

### NewQuoteAgentRunResponse

`func NewQuoteAgentRunResponse(id string, tenantId string, quoteId string, quoteVersion int32, agentName string, trustTier int32, surface string, actorType string, actorId string, status string, startedAt time.Time, completedAt time.Time, inputHash string, outputHash string, rationaleSummary string, toolTrace []map[string]interface{}, errorCode string, traceId string, ) *QuoteAgentRunResponse`

NewQuoteAgentRunResponse instantiates a new QuoteAgentRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteAgentRunResponseWithDefaults

`func NewQuoteAgentRunResponseWithDefaults() *QuoteAgentRunResponse`

NewQuoteAgentRunResponseWithDefaults instantiates a new QuoteAgentRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteAgentRunResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteAgentRunResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteAgentRunResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteAgentRunResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteAgentRunResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteAgentRunResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetQuoteId

`func (o *QuoteAgentRunResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteAgentRunResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteAgentRunResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *QuoteAgentRunResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *QuoteAgentRunResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *QuoteAgentRunResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetAgentName

`func (o *QuoteAgentRunResponse) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *QuoteAgentRunResponse) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *QuoteAgentRunResponse) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetTrustTier

`func (o *QuoteAgentRunResponse) GetTrustTier() int32`

GetTrustTier returns the TrustTier field if non-nil, zero value otherwise.

### GetTrustTierOk

`func (o *QuoteAgentRunResponse) GetTrustTierOk() (*int32, bool)`

GetTrustTierOk returns a tuple with the TrustTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustTier

`func (o *QuoteAgentRunResponse) SetTrustTier(v int32)`

SetTrustTier sets TrustTier field to given value.


### GetSurface

`func (o *QuoteAgentRunResponse) GetSurface() string`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *QuoteAgentRunResponse) GetSurfaceOk() (*string, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *QuoteAgentRunResponse) SetSurface(v string)`

SetSurface sets Surface field to given value.


### GetActorType

`func (o *QuoteAgentRunResponse) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *QuoteAgentRunResponse) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *QuoteAgentRunResponse) SetActorType(v string)`

SetActorType sets ActorType field to given value.


### GetActorId

`func (o *QuoteAgentRunResponse) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *QuoteAgentRunResponse) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *QuoteAgentRunResponse) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetStatus

`func (o *QuoteAgentRunResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteAgentRunResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteAgentRunResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *QuoteAgentRunResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *QuoteAgentRunResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *QuoteAgentRunResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *QuoteAgentRunResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *QuoteAgentRunResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *QuoteAgentRunResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetInputHash

`func (o *QuoteAgentRunResponse) GetInputHash() string`

GetInputHash returns the InputHash field if non-nil, zero value otherwise.

### GetInputHashOk

`func (o *QuoteAgentRunResponse) GetInputHashOk() (*string, bool)`

GetInputHashOk returns a tuple with the InputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputHash

`func (o *QuoteAgentRunResponse) SetInputHash(v string)`

SetInputHash sets InputHash field to given value.


### GetOutputHash

`func (o *QuoteAgentRunResponse) GetOutputHash() string`

GetOutputHash returns the OutputHash field if non-nil, zero value otherwise.

### GetOutputHashOk

`func (o *QuoteAgentRunResponse) GetOutputHashOk() (*string, bool)`

GetOutputHashOk returns a tuple with the OutputHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputHash

`func (o *QuoteAgentRunResponse) SetOutputHash(v string)`

SetOutputHash sets OutputHash field to given value.


### GetRationaleSummary

`func (o *QuoteAgentRunResponse) GetRationaleSummary() string`

GetRationaleSummary returns the RationaleSummary field if non-nil, zero value otherwise.

### GetRationaleSummaryOk

`func (o *QuoteAgentRunResponse) GetRationaleSummaryOk() (*string, bool)`

GetRationaleSummaryOk returns a tuple with the RationaleSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationaleSummary

`func (o *QuoteAgentRunResponse) SetRationaleSummary(v string)`

SetRationaleSummary sets RationaleSummary field to given value.


### GetToolTrace

`func (o *QuoteAgentRunResponse) GetToolTrace() []map[string]interface{}`

GetToolTrace returns the ToolTrace field if non-nil, zero value otherwise.

### GetToolTraceOk

`func (o *QuoteAgentRunResponse) GetToolTraceOk() (*[]map[string]interface{}, bool)`

GetToolTraceOk returns a tuple with the ToolTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTrace

`func (o *QuoteAgentRunResponse) SetToolTrace(v []map[string]interface{})`

SetToolTrace sets ToolTrace field to given value.


### GetErrorCode

`func (o *QuoteAgentRunResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *QuoteAgentRunResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *QuoteAgentRunResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetTraceId

`func (o *QuoteAgentRunResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *QuoteAgentRunResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *QuoteAgentRunResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


