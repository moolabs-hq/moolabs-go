# QuoteAgentProvenanceRunResponse

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
**RationaleSummary** | **string** |  | 
**ErrorCode** | **string** |  | 

## Methods

### NewQuoteAgentProvenanceRunResponse

`func NewQuoteAgentProvenanceRunResponse(id string, tenantId string, quoteId string, quoteVersion int32, agentName string, trustTier int32, surface string, actorType string, actorId string, status string, startedAt time.Time, completedAt time.Time, rationaleSummary string, errorCode string, ) *QuoteAgentProvenanceRunResponse`

NewQuoteAgentProvenanceRunResponse instantiates a new QuoteAgentProvenanceRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteAgentProvenanceRunResponseWithDefaults

`func NewQuoteAgentProvenanceRunResponseWithDefaults() *QuoteAgentProvenanceRunResponse`

NewQuoteAgentProvenanceRunResponseWithDefaults instantiates a new QuoteAgentProvenanceRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteAgentProvenanceRunResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteAgentProvenanceRunResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteAgentProvenanceRunResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteAgentProvenanceRunResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteAgentProvenanceRunResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteAgentProvenanceRunResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetQuoteId

`func (o *QuoteAgentProvenanceRunResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteAgentProvenanceRunResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteAgentProvenanceRunResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *QuoteAgentProvenanceRunResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *QuoteAgentProvenanceRunResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *QuoteAgentProvenanceRunResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetAgentName

`func (o *QuoteAgentProvenanceRunResponse) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *QuoteAgentProvenanceRunResponse) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *QuoteAgentProvenanceRunResponse) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetTrustTier

`func (o *QuoteAgentProvenanceRunResponse) GetTrustTier() int32`

GetTrustTier returns the TrustTier field if non-nil, zero value otherwise.

### GetTrustTierOk

`func (o *QuoteAgentProvenanceRunResponse) GetTrustTierOk() (*int32, bool)`

GetTrustTierOk returns a tuple with the TrustTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustTier

`func (o *QuoteAgentProvenanceRunResponse) SetTrustTier(v int32)`

SetTrustTier sets TrustTier field to given value.


### GetSurface

`func (o *QuoteAgentProvenanceRunResponse) GetSurface() string`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *QuoteAgentProvenanceRunResponse) GetSurfaceOk() (*string, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *QuoteAgentProvenanceRunResponse) SetSurface(v string)`

SetSurface sets Surface field to given value.


### GetActorType

`func (o *QuoteAgentProvenanceRunResponse) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *QuoteAgentProvenanceRunResponse) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *QuoteAgentProvenanceRunResponse) SetActorType(v string)`

SetActorType sets ActorType field to given value.


### GetActorId

`func (o *QuoteAgentProvenanceRunResponse) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *QuoteAgentProvenanceRunResponse) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *QuoteAgentProvenanceRunResponse) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetStatus

`func (o *QuoteAgentProvenanceRunResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteAgentProvenanceRunResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteAgentProvenanceRunResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedAt

`func (o *QuoteAgentProvenanceRunResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *QuoteAgentProvenanceRunResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *QuoteAgentProvenanceRunResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *QuoteAgentProvenanceRunResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *QuoteAgentProvenanceRunResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *QuoteAgentProvenanceRunResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetRationaleSummary

`func (o *QuoteAgentProvenanceRunResponse) GetRationaleSummary() string`

GetRationaleSummary returns the RationaleSummary field if non-nil, zero value otherwise.

### GetRationaleSummaryOk

`func (o *QuoteAgentProvenanceRunResponse) GetRationaleSummaryOk() (*string, bool)`

GetRationaleSummaryOk returns a tuple with the RationaleSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationaleSummary

`func (o *QuoteAgentProvenanceRunResponse) SetRationaleSummary(v string)`

SetRationaleSummary sets RationaleSummary field to given value.


### GetErrorCode

`func (o *QuoteAgentProvenanceRunResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *QuoteAgentProvenanceRunResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *QuoteAgentProvenanceRunResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


