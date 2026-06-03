# QuoteAgentEvaluationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AgentRunId** | **string** |  | 
**CommandId** | **string** |  | 
**Outcome** | **string** |  | 
**ReviewerActorId** | **string** |  | 
**HumanOverrideReason** | **string** |  | 
**EvidenceCoverage** | **int32** |  | 
**Metrics** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewQuoteAgentEvaluationResponse

`func NewQuoteAgentEvaluationResponse(id string, tenantId string, agentRunId string, commandId string, outcome string, reviewerActorId string, humanOverrideReason string, evidenceCoverage int32, metrics map[string]interface{}, createdAt time.Time, ) *QuoteAgentEvaluationResponse`

NewQuoteAgentEvaluationResponse instantiates a new QuoteAgentEvaluationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteAgentEvaluationResponseWithDefaults

`func NewQuoteAgentEvaluationResponseWithDefaults() *QuoteAgentEvaluationResponse`

NewQuoteAgentEvaluationResponseWithDefaults instantiates a new QuoteAgentEvaluationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteAgentEvaluationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteAgentEvaluationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteAgentEvaluationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteAgentEvaluationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteAgentEvaluationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteAgentEvaluationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAgentRunId

`func (o *QuoteAgentEvaluationResponse) GetAgentRunId() string`

GetAgentRunId returns the AgentRunId field if non-nil, zero value otherwise.

### GetAgentRunIdOk

`func (o *QuoteAgentEvaluationResponse) GetAgentRunIdOk() (*string, bool)`

GetAgentRunIdOk returns a tuple with the AgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRunId

`func (o *QuoteAgentEvaluationResponse) SetAgentRunId(v string)`

SetAgentRunId sets AgentRunId field to given value.


### GetCommandId

`func (o *QuoteAgentEvaluationResponse) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *QuoteAgentEvaluationResponse) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *QuoteAgentEvaluationResponse) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.


### GetOutcome

`func (o *QuoteAgentEvaluationResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *QuoteAgentEvaluationResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *QuoteAgentEvaluationResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetReviewerActorId

`func (o *QuoteAgentEvaluationResponse) GetReviewerActorId() string`

GetReviewerActorId returns the ReviewerActorId field if non-nil, zero value otherwise.

### GetReviewerActorIdOk

`func (o *QuoteAgentEvaluationResponse) GetReviewerActorIdOk() (*string, bool)`

GetReviewerActorIdOk returns a tuple with the ReviewerActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerActorId

`func (o *QuoteAgentEvaluationResponse) SetReviewerActorId(v string)`

SetReviewerActorId sets ReviewerActorId field to given value.


### GetHumanOverrideReason

`func (o *QuoteAgentEvaluationResponse) GetHumanOverrideReason() string`

GetHumanOverrideReason returns the HumanOverrideReason field if non-nil, zero value otherwise.

### GetHumanOverrideReasonOk

`func (o *QuoteAgentEvaluationResponse) GetHumanOverrideReasonOk() (*string, bool)`

GetHumanOverrideReasonOk returns a tuple with the HumanOverrideReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanOverrideReason

`func (o *QuoteAgentEvaluationResponse) SetHumanOverrideReason(v string)`

SetHumanOverrideReason sets HumanOverrideReason field to given value.


### GetEvidenceCoverage

`func (o *QuoteAgentEvaluationResponse) GetEvidenceCoverage() int32`

GetEvidenceCoverage returns the EvidenceCoverage field if non-nil, zero value otherwise.

### GetEvidenceCoverageOk

`func (o *QuoteAgentEvaluationResponse) GetEvidenceCoverageOk() (*int32, bool)`

GetEvidenceCoverageOk returns a tuple with the EvidenceCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCoverage

`func (o *QuoteAgentEvaluationResponse) SetEvidenceCoverage(v int32)`

SetEvidenceCoverage sets EvidenceCoverage field to given value.


### GetMetrics

`func (o *QuoteAgentEvaluationResponse) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *QuoteAgentEvaluationResponse) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *QuoteAgentEvaluationResponse) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.


### GetCreatedAt

`func (o *QuoteAgentEvaluationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuoteAgentEvaluationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuoteAgentEvaluationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


