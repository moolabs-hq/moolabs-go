# CaseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CustomerId** | **string** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**BaseCurrencyCode** | Pointer to **string** |  | [optional] 
**RiskTier** | Pointer to **string** | low, medium, high, critical | [optional] [default to "medium"]
**AssignedHuman** | Pointer to **string** |  | [optional] 
**CaseOwnerTeam** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Invoices** | Pointer to [**[]CaseInvoiceCreate**](CaseInvoiceCreate.md) |  | [optional] 

## Methods

### NewCaseCreate

`func NewCaseCreate(accountId string, customerId string, ) *CaseCreate`

NewCaseCreate instantiates a new CaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseCreateWithDefaults

`func NewCaseCreateWithDefaults() *CaseCreate`

NewCaseCreateWithDefaults instantiates a new CaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CaseCreate) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CaseCreate) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CaseCreate) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *CaseCreate) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CaseCreate) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CaseCreate) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCurrencyCode

`func (o *CaseCreate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CaseCreate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CaseCreate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CaseCreate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBaseCurrencyCode

`func (o *CaseCreate) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *CaseCreate) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *CaseCreate) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.

### HasBaseCurrencyCode

`func (o *CaseCreate) HasBaseCurrencyCode() bool`

HasBaseCurrencyCode returns a boolean if a field has been set.

### GetRiskTier

`func (o *CaseCreate) GetRiskTier() string`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *CaseCreate) GetRiskTierOk() (*string, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *CaseCreate) SetRiskTier(v string)`

SetRiskTier sets RiskTier field to given value.

### HasRiskTier

`func (o *CaseCreate) HasRiskTier() bool`

HasRiskTier returns a boolean if a field has been set.

### GetAssignedHuman

`func (o *CaseCreate) GetAssignedHuman() string`

GetAssignedHuman returns the AssignedHuman field if non-nil, zero value otherwise.

### GetAssignedHumanOk

`func (o *CaseCreate) GetAssignedHumanOk() (*string, bool)`

GetAssignedHumanOk returns a tuple with the AssignedHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedHuman

`func (o *CaseCreate) SetAssignedHuman(v string)`

SetAssignedHuman sets AssignedHuman field to given value.

### HasAssignedHuman

`func (o *CaseCreate) HasAssignedHuman() bool`

HasAssignedHuman returns a boolean if a field has been set.

### GetCaseOwnerTeam

`func (o *CaseCreate) GetCaseOwnerTeam() string`

GetCaseOwnerTeam returns the CaseOwnerTeam field if non-nil, zero value otherwise.

### GetCaseOwnerTeamOk

`func (o *CaseCreate) GetCaseOwnerTeamOk() (*string, bool)`

GetCaseOwnerTeamOk returns a tuple with the CaseOwnerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseOwnerTeam

`func (o *CaseCreate) SetCaseOwnerTeam(v string)`

SetCaseOwnerTeam sets CaseOwnerTeam field to given value.

### HasCaseOwnerTeam

`func (o *CaseCreate) HasCaseOwnerTeam() bool`

HasCaseOwnerTeam returns a boolean if a field has been set.

### GetMetadata

`func (o *CaseCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CaseCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CaseCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CaseCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInvoices

`func (o *CaseCreate) GetInvoices() []CaseInvoiceCreate`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *CaseCreate) GetInvoicesOk() (*[]CaseInvoiceCreate, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *CaseCreate) SetInvoices(v []CaseInvoiceCreate)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *CaseCreate) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


