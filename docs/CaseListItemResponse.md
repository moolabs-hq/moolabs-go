# CaseListItemResponse

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
**InvoiceCount** | Pointer to **int32** |  | [optional] [default to 0]
**Invoices** | Pointer to [**[]CaseInvoiceResponse**](CaseInvoiceResponse.md) |  | [optional] 
**Promises** | Pointer to [**[]PromiseResponse**](PromiseResponse.md) |  | [optional] 
**Account** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCaseListItemResponse

`func NewCaseListItemResponse(id string, tenantId string, accountId string, customerId string, scopeHash string, currencyCode string, baseCurrencyCode string, status string, riskTier string, autonomyLevel int32, totalOutstandingMicros int32, totalOutstandingBaseMicros int32, dunningStage string, createdAt time.Time, updatedAt time.Time, ) *CaseListItemResponse`

NewCaseListItemResponse instantiates a new CaseListItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseListItemResponseWithDefaults

`func NewCaseListItemResponseWithDefaults() *CaseListItemResponse`

NewCaseListItemResponseWithDefaults instantiates a new CaseListItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaseListItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaseListItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaseListItemResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CaseListItemResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CaseListItemResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CaseListItemResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAccountId

`func (o *CaseListItemResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CaseListItemResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CaseListItemResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *CaseListItemResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CaseListItemResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CaseListItemResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetScopeHash

`func (o *CaseListItemResponse) GetScopeHash() string`

GetScopeHash returns the ScopeHash field if non-nil, zero value otherwise.

### GetScopeHashOk

`func (o *CaseListItemResponse) GetScopeHashOk() (*string, bool)`

GetScopeHashOk returns a tuple with the ScopeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeHash

`func (o *CaseListItemResponse) SetScopeHash(v string)`

SetScopeHash sets ScopeHash field to given value.


### GetCurrencyCode

`func (o *CaseListItemResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CaseListItemResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CaseListItemResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetBaseCurrencyCode

`func (o *CaseListItemResponse) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *CaseListItemResponse) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *CaseListItemResponse) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.


### GetStatus

`func (o *CaseListItemResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CaseListItemResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CaseListItemResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRiskTier

`func (o *CaseListItemResponse) GetRiskTier() string`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *CaseListItemResponse) GetRiskTierOk() (*string, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *CaseListItemResponse) SetRiskTier(v string)`

SetRiskTier sets RiskTier field to given value.


### GetAutonomyLevel

`func (o *CaseListItemResponse) GetAutonomyLevel() int32`

GetAutonomyLevel returns the AutonomyLevel field if non-nil, zero value otherwise.

### GetAutonomyLevelOk

`func (o *CaseListItemResponse) GetAutonomyLevelOk() (*int32, bool)`

GetAutonomyLevelOk returns a tuple with the AutonomyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomyLevel

`func (o *CaseListItemResponse) SetAutonomyLevel(v int32)`

SetAutonomyLevel sets AutonomyLevel field to given value.


### GetTotalOutstandingMicros

`func (o *CaseListItemResponse) GetTotalOutstandingMicros() int32`

GetTotalOutstandingMicros returns the TotalOutstandingMicros field if non-nil, zero value otherwise.

### GetTotalOutstandingMicrosOk

`func (o *CaseListItemResponse) GetTotalOutstandingMicrosOk() (*int32, bool)`

GetTotalOutstandingMicrosOk returns a tuple with the TotalOutstandingMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstandingMicros

`func (o *CaseListItemResponse) SetTotalOutstandingMicros(v int32)`

SetTotalOutstandingMicros sets TotalOutstandingMicros field to given value.


### GetTotalOutstandingBaseMicros

`func (o *CaseListItemResponse) GetTotalOutstandingBaseMicros() int32`

GetTotalOutstandingBaseMicros returns the TotalOutstandingBaseMicros field if non-nil, zero value otherwise.

### GetTotalOutstandingBaseMicrosOk

`func (o *CaseListItemResponse) GetTotalOutstandingBaseMicrosOk() (*int32, bool)`

GetTotalOutstandingBaseMicrosOk returns a tuple with the TotalOutstandingBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstandingBaseMicros

`func (o *CaseListItemResponse) SetTotalOutstandingBaseMicros(v int32)`

SetTotalOutstandingBaseMicros sets TotalOutstandingBaseMicros field to given value.


### GetTotalDisputedMicros

`func (o *CaseListItemResponse) GetTotalDisputedMicros() int32`

GetTotalDisputedMicros returns the TotalDisputedMicros field if non-nil, zero value otherwise.

### GetTotalDisputedMicrosOk

`func (o *CaseListItemResponse) GetTotalDisputedMicrosOk() (*int32, bool)`

GetTotalDisputedMicrosOk returns a tuple with the TotalDisputedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisputedMicros

`func (o *CaseListItemResponse) SetTotalDisputedMicros(v int32)`

SetTotalDisputedMicros sets TotalDisputedMicros field to given value.

### HasTotalDisputedMicros

`func (o *CaseListItemResponse) HasTotalDisputedMicros() bool`

HasTotalDisputedMicros returns a boolean if a field has been set.

### GetOldestDueDate

`func (o *CaseListItemResponse) GetOldestDueDate() string`

GetOldestDueDate returns the OldestDueDate field if non-nil, zero value otherwise.

### GetOldestDueDateOk

`func (o *CaseListItemResponse) GetOldestDueDateOk() (*string, bool)`

GetOldestDueDateOk returns a tuple with the OldestDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestDueDate

`func (o *CaseListItemResponse) SetOldestDueDate(v string)`

SetOldestDueDate sets OldestDueDate field to given value.

### HasOldestDueDate

`func (o *CaseListItemResponse) HasOldestDueDate() bool`

HasOldestDueDate returns a boolean if a field has been set.

### GetLastContactAt

`func (o *CaseListItemResponse) GetLastContactAt() time.Time`

GetLastContactAt returns the LastContactAt field if non-nil, zero value otherwise.

### GetLastContactAtOk

`func (o *CaseListItemResponse) GetLastContactAtOk() (*time.Time, bool)`

GetLastContactAtOk returns a tuple with the LastContactAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactAt

`func (o *CaseListItemResponse) SetLastContactAt(v time.Time)`

SetLastContactAt sets LastContactAt field to given value.

### HasLastContactAt

`func (o *CaseListItemResponse) HasLastContactAt() bool`

HasLastContactAt returns a boolean if a field has been set.

### GetLastPaymentAt

`func (o *CaseListItemResponse) GetLastPaymentAt() time.Time`

GetLastPaymentAt returns the LastPaymentAt field if non-nil, zero value otherwise.

### GetLastPaymentAtOk

`func (o *CaseListItemResponse) GetLastPaymentAtOk() (*time.Time, bool)`

GetLastPaymentAtOk returns a tuple with the LastPaymentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAt

`func (o *CaseListItemResponse) SetLastPaymentAt(v time.Time)`

SetLastPaymentAt sets LastPaymentAt field to given value.

### HasLastPaymentAt

`func (o *CaseListItemResponse) HasLastPaymentAt() bool`

HasLastPaymentAt returns a boolean if a field has been set.

### GetNextActionAt

`func (o *CaseListItemResponse) GetNextActionAt() time.Time`

GetNextActionAt returns the NextActionAt field if non-nil, zero value otherwise.

### GetNextActionAtOk

`func (o *CaseListItemResponse) GetNextActionAtOk() (*time.Time, bool)`

GetNextActionAtOk returns a tuple with the NextActionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextActionAt

`func (o *CaseListItemResponse) SetNextActionAt(v time.Time)`

SetNextActionAt sets NextActionAt field to given value.

### HasNextActionAt

`func (o *CaseListItemResponse) HasNextActionAt() bool`

HasNextActionAt returns a boolean if a field has been set.

### GetNextActionType

`func (o *CaseListItemResponse) GetNextActionType() NextActionType`

GetNextActionType returns the NextActionType field if non-nil, zero value otherwise.

### GetNextActionTypeOk

`func (o *CaseListItemResponse) GetNextActionTypeOk() (*NextActionType, bool)`

GetNextActionTypeOk returns a tuple with the NextActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextActionType

`func (o *CaseListItemResponse) SetNextActionType(v NextActionType)`

SetNextActionType sets NextActionType field to given value.

### HasNextActionType

`func (o *CaseListItemResponse) HasNextActionType() bool`

HasNextActionType returns a boolean if a field has been set.

### GetNextPromiseDueAt

`func (o *CaseListItemResponse) GetNextPromiseDueAt() time.Time`

GetNextPromiseDueAt returns the NextPromiseDueAt field if non-nil, zero value otherwise.

### GetNextPromiseDueAtOk

`func (o *CaseListItemResponse) GetNextPromiseDueAtOk() (*time.Time, bool)`

GetNextPromiseDueAtOk returns a tuple with the NextPromiseDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPromiseDueAt

`func (o *CaseListItemResponse) SetNextPromiseDueAt(v time.Time)`

SetNextPromiseDueAt sets NextPromiseDueAt field to given value.

### HasNextPromiseDueAt

`func (o *CaseListItemResponse) HasNextPromiseDueAt() bool`

HasNextPromiseDueAt returns a boolean if a field has been set.

### GetDunningStage

`func (o *CaseListItemResponse) GetDunningStage() string`

GetDunningStage returns the DunningStage field if non-nil, zero value otherwise.

### GetDunningStageOk

`func (o *CaseListItemResponse) GetDunningStageOk() (*string, bool)`

GetDunningStageOk returns a tuple with the DunningStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningStage

`func (o *CaseListItemResponse) SetDunningStage(v string)`

SetDunningStage sets DunningStage field to given value.


### GetCurrentStrategyStep

`func (o *CaseListItemResponse) GetCurrentStrategyStep() string`

GetCurrentStrategyStep returns the CurrentStrategyStep field if non-nil, zero value otherwise.

### GetCurrentStrategyStepOk

`func (o *CaseListItemResponse) GetCurrentStrategyStepOk() (*string, bool)`

GetCurrentStrategyStepOk returns a tuple with the CurrentStrategyStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStrategyStep

`func (o *CaseListItemResponse) SetCurrentStrategyStep(v string)`

SetCurrentStrategyStep sets CurrentStrategyStep field to given value.

### HasCurrentStrategyStep

`func (o *CaseListItemResponse) HasCurrentStrategyStep() bool`

HasCurrentStrategyStep returns a boolean if a field has been set.

### GetStrategyId

`func (o *CaseListItemResponse) GetStrategyId() string`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *CaseListItemResponse) GetStrategyIdOk() (*string, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *CaseListItemResponse) SetStrategyId(v string)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *CaseListItemResponse) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetAssignedAgent

`func (o *CaseListItemResponse) GetAssignedAgent() string`

GetAssignedAgent returns the AssignedAgent field if non-nil, zero value otherwise.

### GetAssignedAgentOk

`func (o *CaseListItemResponse) GetAssignedAgentOk() (*string, bool)`

GetAssignedAgentOk returns a tuple with the AssignedAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgent

`func (o *CaseListItemResponse) SetAssignedAgent(v string)`

SetAssignedAgent sets AssignedAgent field to given value.

### HasAssignedAgent

`func (o *CaseListItemResponse) HasAssignedAgent() bool`

HasAssignedAgent returns a boolean if a field has been set.

### GetAssignedHuman

`func (o *CaseListItemResponse) GetAssignedHuman() string`

GetAssignedHuman returns the AssignedHuman field if non-nil, zero value otherwise.

### GetAssignedHumanOk

`func (o *CaseListItemResponse) GetAssignedHumanOk() (*string, bool)`

GetAssignedHumanOk returns a tuple with the AssignedHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedHuman

`func (o *CaseListItemResponse) SetAssignedHuman(v string)`

SetAssignedHuman sets AssignedHuman field to given value.

### HasAssignedHuman

`func (o *CaseListItemResponse) HasAssignedHuman() bool`

HasAssignedHuman returns a boolean if a field has been set.

### GetCaseOwnerTeam

`func (o *CaseListItemResponse) GetCaseOwnerTeam() string`

GetCaseOwnerTeam returns the CaseOwnerTeam field if non-nil, zero value otherwise.

### GetCaseOwnerTeamOk

`func (o *CaseListItemResponse) GetCaseOwnerTeamOk() (*string, bool)`

GetCaseOwnerTeamOk returns a tuple with the CaseOwnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseOwnerTeam

`func (o *CaseListItemResponse) SetCaseOwnerTeam(v string)`

SetCaseOwnerTeam sets CaseOwnerTeam field to given value.

### HasCaseOwnerTeam

`func (o *CaseListItemResponse) HasCaseOwnerTeam() bool`

HasCaseOwnerTeam returns a boolean if a field has been set.

### GetMetadata

`func (o *CaseListItemResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CaseListItemResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CaseListItemResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CaseListItemResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInvoiceCount

`func (o *CaseListItemResponse) GetInvoiceCount() int32`

GetInvoiceCount returns the InvoiceCount field if non-nil, zero value otherwise.

### GetInvoiceCountOk

`func (o *CaseListItemResponse) GetInvoiceCountOk() (*int32, bool)`

GetInvoiceCountOk returns a tuple with the InvoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCount

`func (o *CaseListItemResponse) SetInvoiceCount(v int32)`

SetInvoiceCount sets InvoiceCount field to given value.

### HasInvoiceCount

`func (o *CaseListItemResponse) HasInvoiceCount() bool`

HasInvoiceCount returns a boolean if a field has been set.

### GetInvoices

`func (o *CaseListItemResponse) GetInvoices() []CaseInvoiceResponse`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *CaseListItemResponse) GetInvoicesOk() (*[]CaseInvoiceResponse, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *CaseListItemResponse) SetInvoices(v []CaseInvoiceResponse)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *CaseListItemResponse) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetPromises

`func (o *CaseListItemResponse) GetPromises() []PromiseResponse`

GetPromises returns the Promises field if non-nil, zero value otherwise.

### GetPromisesOk

`func (o *CaseListItemResponse) GetPromisesOk() (*[]PromiseResponse, bool)`

GetPromisesOk returns a tuple with the Promises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromises

`func (o *CaseListItemResponse) SetPromises(v []PromiseResponse)`

SetPromises sets Promises field to given value.

### HasPromises

`func (o *CaseListItemResponse) HasPromises() bool`

HasPromises returns a boolean if a field has been set.

### GetAccount

`func (o *CaseListItemResponse) GetAccount() AccountSummary`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CaseListItemResponse) GetAccountOk() (*AccountSummary, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CaseListItemResponse) SetAccount(v AccountSummary)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CaseListItemResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaseListItemResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaseListItemResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaseListItemResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CaseListItemResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CaseListItemResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CaseListItemResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


