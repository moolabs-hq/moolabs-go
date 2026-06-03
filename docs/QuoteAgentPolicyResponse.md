# QuoteAgentPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AgentName** | **string** |  | 
**TrustTier** | **int32** |  | 
**Enabled** | **bool** |  | 
**KillSwitch** | **bool** |  | 
**ShadowMode** | **bool** |  | 
**MaxAutonomy** | **string** |  | 
**EvidenceCoverageThreshold** | **int32** |  | 
**RateLimitPerMinute** | **int32** |  | 
**AllowedSurfaces** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewQuoteAgentPolicyResponse

`func NewQuoteAgentPolicyResponse(id string, tenantId string, agentName string, trustTier int32, enabled bool, killSwitch bool, shadowMode bool, maxAutonomy string, evidenceCoverageThreshold int32, rateLimitPerMinute int32, allowedSurfaces []string, createdAt time.Time, updatedAt time.Time, ) *QuoteAgentPolicyResponse`

NewQuoteAgentPolicyResponse instantiates a new QuoteAgentPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteAgentPolicyResponseWithDefaults

`func NewQuoteAgentPolicyResponseWithDefaults() *QuoteAgentPolicyResponse`

NewQuoteAgentPolicyResponseWithDefaults instantiates a new QuoteAgentPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteAgentPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteAgentPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteAgentPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteAgentPolicyResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteAgentPolicyResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteAgentPolicyResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAgentName

`func (o *QuoteAgentPolicyResponse) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *QuoteAgentPolicyResponse) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *QuoteAgentPolicyResponse) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.


### GetTrustTier

`func (o *QuoteAgentPolicyResponse) GetTrustTier() int32`

GetTrustTier returns the TrustTier field if non-nil, zero value otherwise.

### GetTrustTierOk

`func (o *QuoteAgentPolicyResponse) GetTrustTierOk() (*int32, bool)`

GetTrustTierOk returns a tuple with the TrustTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustTier

`func (o *QuoteAgentPolicyResponse) SetTrustTier(v int32)`

SetTrustTier sets TrustTier field to given value.


### GetEnabled

`func (o *QuoteAgentPolicyResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *QuoteAgentPolicyResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *QuoteAgentPolicyResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetKillSwitch

`func (o *QuoteAgentPolicyResponse) GetKillSwitch() bool`

GetKillSwitch returns the KillSwitch field if non-nil, zero value otherwise.

### GetKillSwitchOk

`func (o *QuoteAgentPolicyResponse) GetKillSwitchOk() (*bool, bool)`

GetKillSwitchOk returns a tuple with the KillSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillSwitch

`func (o *QuoteAgentPolicyResponse) SetKillSwitch(v bool)`

SetKillSwitch sets KillSwitch field to given value.


### GetShadowMode

`func (o *QuoteAgentPolicyResponse) GetShadowMode() bool`

GetShadowMode returns the ShadowMode field if non-nil, zero value otherwise.

### GetShadowModeOk

`func (o *QuoteAgentPolicyResponse) GetShadowModeOk() (*bool, bool)`

GetShadowModeOk returns a tuple with the ShadowMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowMode

`func (o *QuoteAgentPolicyResponse) SetShadowMode(v bool)`

SetShadowMode sets ShadowMode field to given value.


### GetMaxAutonomy

`func (o *QuoteAgentPolicyResponse) GetMaxAutonomy() string`

GetMaxAutonomy returns the MaxAutonomy field if non-nil, zero value otherwise.

### GetMaxAutonomyOk

`func (o *QuoteAgentPolicyResponse) GetMaxAutonomyOk() (*string, bool)`

GetMaxAutonomyOk returns a tuple with the MaxAutonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutonomy

`func (o *QuoteAgentPolicyResponse) SetMaxAutonomy(v string)`

SetMaxAutonomy sets MaxAutonomy field to given value.


### GetEvidenceCoverageThreshold

`func (o *QuoteAgentPolicyResponse) GetEvidenceCoverageThreshold() int32`

GetEvidenceCoverageThreshold returns the EvidenceCoverageThreshold field if non-nil, zero value otherwise.

### GetEvidenceCoverageThresholdOk

`func (o *QuoteAgentPolicyResponse) GetEvidenceCoverageThresholdOk() (*int32, bool)`

GetEvidenceCoverageThresholdOk returns a tuple with the EvidenceCoverageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCoverageThreshold

`func (o *QuoteAgentPolicyResponse) SetEvidenceCoverageThreshold(v int32)`

SetEvidenceCoverageThreshold sets EvidenceCoverageThreshold field to given value.


### GetRateLimitPerMinute

`func (o *QuoteAgentPolicyResponse) GetRateLimitPerMinute() int32`

GetRateLimitPerMinute returns the RateLimitPerMinute field if non-nil, zero value otherwise.

### GetRateLimitPerMinuteOk

`func (o *QuoteAgentPolicyResponse) GetRateLimitPerMinuteOk() (*int32, bool)`

GetRateLimitPerMinuteOk returns a tuple with the RateLimitPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerMinute

`func (o *QuoteAgentPolicyResponse) SetRateLimitPerMinute(v int32)`

SetRateLimitPerMinute sets RateLimitPerMinute field to given value.


### GetAllowedSurfaces

`func (o *QuoteAgentPolicyResponse) GetAllowedSurfaces() []string`

GetAllowedSurfaces returns the AllowedSurfaces field if non-nil, zero value otherwise.

### GetAllowedSurfacesOk

`func (o *QuoteAgentPolicyResponse) GetAllowedSurfacesOk() (*[]string, bool)`

GetAllowedSurfacesOk returns a tuple with the AllowedSurfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSurfaces

`func (o *QuoteAgentPolicyResponse) SetAllowedSurfaces(v []string)`

SetAllowedSurfaces sets AllowedSurfaces field to given value.


### GetCreatedAt

`func (o *QuoteAgentPolicyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuoteAgentPolicyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuoteAgentPolicyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QuoteAgentPolicyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QuoteAgentPolicyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QuoteAgentPolicyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


