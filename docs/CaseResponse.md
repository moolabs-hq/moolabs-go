# CaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AccountId** | **string** |  | 
**CustomerId** | **string** |  | 
**ScopeHash** | **string** |  | 
**CurrencyCode** | **string** |  | 
**BaseCurrencyCode** | **string** |  | 
**Status** | **string** |  | 
**RiskTier** | **string** |  | 
**AutonomyLevel** | **int32** |  | 
**TotalOutstandingMicros** | **int32** |  | 
**TotalOutstandingBaseMicros** | **int32** |  | 
**TotalDisputedMicros** | Pointer to **int32** |  | [optional] [default to 0]
**OldestDueDate** | Pointer to **string** |  | [optional] 
**LastContactAt** | Pointer to **time.Time** |  | [optional] 
**LastPaymentAt** | Pointer to **time.Time** |  | [optional] 
**NextActionAt** | Pointer to **time.Time** |  | [optional] 
**NextActionType** | Pointer to [**NextActionType**](NextActionType.md) |  | [optional] 
**NextPromiseDueAt** | Pointer to **time.Time** |  | [optional] 
**DunningStage** | **string** |  | 
**CurrentStrategyStep** | Pointer to **string** |  | [optional] 
**StrategyId** | Pointer to **string** |  | [optional] 
**AssignedAgent** | Pointer to **string** |  | [optional] 
**AssignedHuman** | Pointer to **string** |  | [optional] 
**CaseOwnerTeam** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Invoices** | Pointer to [**[]CaseInvoiceResponse**](CaseInvoiceResponse.md) |  | [optional] 
**Promises** | Pointer to [**[]PromiseResponse**](PromiseResponse.md) |  | [optional] 
**Account** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCaseResponse

`func NewCaseResponse(id string, tenantId string, accountId string, customerId string, scopeHash string, currencyCode string, baseCurrencyCode string, status string, riskTier string, autonomyLevel int32, totalOutstandingMicros int32, totalOutstandingBaseMicros int32, dunningStage string, createdAt time.Time, updatedAt time.Time, ) *CaseResponse`

NewCaseResponse instantiates a new CaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseResponseWithDefaults

`func NewCaseResponseWithDefaults() *CaseResponse`

NewCaseResponseWithDefaults instantiates a new CaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CaseResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CaseResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CaseResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAccountId

`func (o *CaseResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CaseResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CaseResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *CaseResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CaseResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CaseResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetScopeHash

`func (o *CaseResponse) GetScopeHash() string`

GetScopeHash returns the ScopeHash field if non-nil, zero value otherwise.

### GetScopeHashOk

`func (o *CaseResponse) GetScopeHashOk() (*string, bool)`

GetScopeHashOk returns a tuple with the ScopeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeHash

`func (o *CaseResponse) SetScopeHash(v string)`

SetScopeHash sets ScopeHash field to given value.


### GetCurrencyCode

`func (o *CaseResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CaseResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CaseResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetBaseCurrencyCode

`func (o *CaseResponse) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *CaseResponse) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *CaseResponse) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.


### GetStatus

`func (o *CaseResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaseResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaseResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRiskTier

`func (o *CaseResponse) GetRiskTier() string`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *CaseResponse) GetRiskTierOk() (*string, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *CaseResponse) SetRiskTier(v string)`

SetRiskTier sets RiskTier field to given value.


### GetAutonomyLevel

`func (o *CaseResponse) GetAutonomyLevel() int32`

GetAutonomyLevel returns the AutonomyLevel field if non-nil, zero value otherwise.

### GetAutonomyLevelOk

`func (o *CaseResponse) GetAutonomyLevelOk() (*int32, bool)`

GetAutonomyLevelOk returns a tuple with the AutonomyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomyLevel

`func (o *CaseResponse) SetAutonomyLevel(v int32)`

SetAutonomyLevel sets AutonomyLevel field to given value.


### GetTotalOutstandingMicros

`func (o *CaseResponse) GetTotalOutstandingMicros() int32`

GetTotalOutstandingMicros returns the TotalOutstandingMicros field if non-nil, zero value otherwise.

### GetTotalOutstandingMicrosOk

`func (o *CaseResponse) GetTotalOutstandingMicrosOk() (*int32, bool)`

GetTotalOutstandingMicrosOk returns a tuple with the TotalOutstandingMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstandingMicros

`func (o *CaseResponse) SetTotalOutstandingMicros(v int32)`

SetTotalOutstandingMicros sets TotalOutstandingMicros field to given value.


### GetTotalOutstandingBaseMicros

`func (o *CaseResponse) GetTotalOutstandingBaseMicros() int32`

GetTotalOutstandingBaseMicros returns the TotalOutstandingBaseMicros field if non-nil, zero value otherwise.

### GetTotalOutstandingBaseMicrosOk

`func (o *CaseResponse) GetTotalOutstandingBaseMicrosOk() (*int32, bool)`

GetTotalOutstandingBaseMicrosOk returns a tuple with the TotalOutstandingBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstandingBaseMicros

`func (o *CaseResponse) SetTotalOutstandingBaseMicros(v int32)`

SetTotalOutstandingBaseMicros sets TotalOutstandingBaseMicros field to given value.


### GetTotalDisputedMicros

`func (o *CaseResponse) GetTotalDisputedMicros() int32`

GetTotalDisputedMicros returns the TotalDisputedMicros field if non-nil, zero value otherwise.

### GetTotalDisputedMicrosOk

`func (o *CaseResponse) GetTotalDisputedMicrosOk() (*int32, bool)`

GetTotalDisputedMicrosOk returns a tuple with the TotalDisputedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisputedMicros

`func (o *CaseResponse) SetTotalDisputedMicros(v int32)`

SetTotalDisputedMicros sets TotalDisputedMicros field to given value.

### HasTotalDisputedMicros

`func (o *CaseResponse) HasTotalDisputedMicros() bool`

HasTotalDisputedMicros returns a boolean if a field has been set.

### GetOldestDueDate

`func (o *CaseResponse) GetOldestDueDate() string`

GetOldestDueDate returns the OldestDueDate field if non-nil, zero value otherwise.

### GetOldestDueDateOk

`func (o *CaseResponse) GetOldestDueDateOk() (*string, bool)`

GetOldestDueDateOk returns a tuple with the OldestDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestDueDate

`func (o *CaseResponse) SetOldestDueDate(v string)`

SetOldestDueDate sets OldestDueDate field to given value.

### HasOldestDueDate

`func (o *CaseResponse) HasOldestDueDate() bool`

HasOldestDueDate returns a boolean if a field has been set.

### GetLastContactAt

`func (o *CaseResponse) GetLastContactAt() time.Time`

GetLastContactAt returns the LastContactAt field if non-nil, zero value otherwise.

### GetLastContactAtOk

`func (o *CaseResponse) GetLastContactAtOk() (*time.Time, bool)`

GetLastContactAtOk returns a tuple with the LastContactAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactAt

`func (o *CaseResponse) SetLastContactAt(v time.Time)`

SetLastContactAt sets LastContactAt field to given value.

### HasLastContactAt

`func (o *CaseResponse) HasLastContactAt() bool`

HasLastContactAt returns a boolean if a field has been set.

### GetLastPaymentAt

`func (o *CaseResponse) GetLastPaymentAt() time.Time`

GetLastPaymentAt returns the LastPaymentAt field if non-nil, zero value otherwise.

### GetLastPaymentAtOk

`func (o *CaseResponse) GetLastPaymentAtOk() (*time.Time, bool)`

GetLastPaymentAtOk returns a tuple with the LastPaymentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAt

`func (o *CaseResponse) SetLastPaymentAt(v time.Time)`

SetLastPaymentAt sets LastPaymentAt field to given value.

### HasLastPaymentAt

`func (o *CaseResponse) HasLastPaymentAt() bool`

HasLastPaymentAt returns a boolean if a field has been set.

### GetNextActionAt

`func (o *CaseResponse) GetNextActionAt() time.Time`

GetNextActionAt returns the NextActionAt field if non-nil, zero value otherwise.

### GetNextActionAtOk

`func (o *CaseResponse) GetNextActionAtOk() (*time.Time, bool)`

GetNextActionAtOk returns a tuple with the NextActionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextActionAt

`func (o *CaseResponse) SetNextActionAt(v time.Time)`

SetNextActionAt sets NextActionAt field to given value.

### HasNextActionAt

`func (o *CaseResponse) HasNextActionAt() bool`

HasNextActionAt returns a boolean if a field has been set.

### GetNextActionType

`func (o *CaseResponse) GetNextActionType() NextActionType`

GetNextActionType returns the NextActionType field if non-nil, zero value otherwise.

### GetNextActionTypeOk

`func (o *CaseResponse) GetNextActionTypeOk() (*NextActionType, bool)`

GetNextActionTypeOk returns a tuple with the NextActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextActionType

`func (o *CaseResponse) SetNextActionType(v NextActionType)`

SetNextActionType sets NextActionType field to given value.

### HasNextActionType

`func (o *CaseResponse) HasNextActionType() bool`

HasNextActionType returns a boolean if a field has been set.

### GetNextPromiseDueAt

`func (o *CaseResponse) GetNextPromiseDueAt() time.Time`

GetNextPromiseDueAt returns the NextPromiseDueAt field if non-nil, zero value otherwise.

### GetNextPromiseDueAtOk

`func (o *CaseResponse) GetNextPromiseDueAtOk() (*time.Time, bool)`

GetNextPromiseDueAtOk returns a tuple with the NextPromiseDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPromiseDueAt

`func (o *CaseResponse) SetNextPromiseDueAt(v time.Time)`

SetNextPromiseDueAt sets NextPromiseDueAt field to given value.

### HasNextPromiseDueAt

`func (o *CaseResponse) HasNextPromiseDueAt() bool`

HasNextPromiseDueAt returns a boolean if a field has been set.

### GetDunningStage

`func (o *CaseResponse) GetDunningStage() string`

GetDunningStage returns the DunningStage field if non-nil, zero value otherwise.

### GetDunningStageOk

`func (o *CaseResponse) GetDunningStageOk() (*string, bool)`

GetDunningStageOk returns a tuple with the DunningStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningStage

`func (o *CaseResponse) SetDunningStage(v string)`

SetDunningStage sets DunningStage field to given value.


### GetCurrentStrategyStep

`func (o *CaseResponse) GetCurrentStrategyStep() string`

GetCurrentStrategyStep returns the CurrentStrategyStep field if non-nil, zero value otherwise.

### GetCurrentStrategyStepOk

`func (o *CaseResponse) GetCurrentStrategyStepOk() (*string, bool)`

GetCurrentStrategyStepOk returns a tuple with the CurrentStrategyStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStrategyStep

`func (o *CaseResponse) SetCurrentStrategyStep(v string)`

SetCurrentStrategyStep sets CurrentStrategyStep field to given value.

### HasCurrentStrategyStep

`func (o *CaseResponse) HasCurrentStrategyStep() bool`

HasCurrentStrategyStep returns a boolean if a field has been set.

### GetStrategyId

`func (o *CaseResponse) GetStrategyId() string`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *CaseResponse) GetStrategyIdOk() (*string, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *CaseResponse) SetStrategyId(v string)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *CaseResponse) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetAssignedAgent

`func (o *CaseResponse) GetAssignedAgent() string`

GetAssignedAgent returns the AssignedAgent field if non-nil, zero value otherwise.

### GetAssignedAgentOk

`func (o *CaseResponse) GetAssignedAgentOk() (*string, bool)`

GetAssignedAgentOk returns a tuple with the AssignedAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgent

`func (o *CaseResponse) SetAssignedAgent(v string)`

SetAssignedAgent sets AssignedAgent field to given value.

### HasAssignedAgent

`func (o *CaseResponse) HasAssignedAgent() bool`

HasAssignedAgent returns a boolean if a field has been set.

### GetAssignedHuman

`func (o *CaseResponse) GetAssignedHuman() string`

GetAssignedHuman returns the AssignedHuman field if non-nil, zero value otherwise.

### GetAssignedHumanOk

`func (o *CaseResponse) GetAssignedHumanOk() (*string, bool)`

GetAssignedHumanOk returns a tuple with the AssignedHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedHuman

`func (o *CaseResponse) SetAssignedHuman(v string)`

SetAssignedHuman sets AssignedHuman field to given value.

### HasAssignedHuman

`func (o *CaseResponse) HasAssignedHuman() bool`

HasAssignedHuman returns a boolean if a field has been set.

### GetCaseOwnerTeam

`func (o *CaseResponse) GetCaseOwnerTeam() string`

GetCaseOwnerTeam returns the CaseOwnerTeam field if non-nil, zero value otherwise.

### GetCaseOwnerTeamOk

`func (o *CaseResponse) GetCaseOwnerTeamOk() (*string, bool)`

GetCaseOwnerTeamOk returns a tuple with the CaseOwnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseOwnerTeam

`func (o *CaseResponse) SetCaseOwnerTeam(v string)`

SetCaseOwnerTeam sets CaseOwnerTeam field to given value.

### HasCaseOwnerTeam

`func (o *CaseResponse) HasCaseOwnerTeam() bool`

HasCaseOwnerTeam returns a boolean if a field has been set.

### GetMetadata

`func (o *CaseResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CaseResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CaseResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CaseResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInvoices

`func (o *CaseResponse) GetInvoices() []CaseInvoiceResponse`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *CaseResponse) GetInvoicesOk() (*[]CaseInvoiceResponse, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *CaseResponse) SetInvoices(v []CaseInvoiceResponse)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *CaseResponse) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetPromises

`func (o *CaseResponse) GetPromises() []PromiseResponse`

GetPromises returns the Promises field if non-nil, zero value otherwise.

### GetPromisesOk

`func (o *CaseResponse) GetPromisesOk() (*[]PromiseResponse, bool)`

GetPromisesOk returns a tuple with the Promises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromises

`func (o *CaseResponse) SetPromises(v []PromiseResponse)`

SetPromises sets Promises field to given value.

### HasPromises

`func (o *CaseResponse) HasPromises() bool`

HasPromises returns a boolean if a field has been set.

### GetAccount

`func (o *CaseResponse) GetAccount() AccountSummary`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CaseResponse) GetAccountOk() (*AccountSummary, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CaseResponse) SetAccount(v AccountSummary)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CaseResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CaseResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CaseResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CaseResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


