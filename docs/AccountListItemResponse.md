# AccountListItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CustomerId** | **string** |  | 
**BillingAccountId** | Pointer to **string** |  | [optional] 
**BillToEntityId** | Pointer to **string** |  | [optional] 
**LegalName** | **string** |  | 
**BillingCountryCode** | Pointer to **string** |  | [optional] 
**BillingRegionCode** | Pointer to **string** |  | [optional] 
**BillingTimezone** | Pointer to **string** |  | [optional] 
**DefaultLanguageCode** | **string** |  | 
**DefaultCurrencyCode** | **string** |  | 
**PortalRequired** | **bool** |  | 
**PaymentTerms** | Pointer to **string** |  | [optional] 
**IsStrategic** | **bool** |  | 
**SourceSystem** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Subsidiaries** | Pointer to [**[]AccountListItemResponseSubsidiariesInner**](AccountListItemResponseSubsidiariesInner.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**OpenBalanceMicros** | Pointer to **int32** | Display open balance. Uses the invoice currency when a single open currency exists; uses base currency when open invoices are mixed-currency. | [optional] [default to 0]
**OpenBalanceBaseMicros** | Pointer to **int32** | Base-currency open balance used for filtering and sorting. | [optional] [default to 0]
**OpenBalanceCurrencyCode** | Pointer to **string** | Currency code for open_balance_micros. | [optional] [default to "USD"]
**OpenBalanceBaseCurrencyCode** | Pointer to **string** | Base currency code for open_balance_base_micros. | [optional] [default to "USD"]
**OpenBalanceHasMixedCurrencies** | Pointer to **bool** |  | [optional] [default to false]
**SubsidiaryOptions** | Pointer to [**[]AccountSubsidiaryOption**](AccountSubsidiaryOption.md) |  | [optional] 

## Methods

### NewAccountListItemResponse

`func NewAccountListItemResponse(id string, tenantId string, customerId string, legalName string, defaultLanguageCode string, defaultCurrencyCode string, portalRequired bool, isStrategic bool, sourceSystem string, createdAt time.Time, updatedAt time.Time, ) *AccountListItemResponse`

NewAccountListItemResponse instantiates a new AccountListItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListItemResponseWithDefaults

`func NewAccountListItemResponseWithDefaults() *AccountListItemResponse`

NewAccountListItemResponseWithDefaults instantiates a new AccountListItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountListItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountListItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountListItemResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *AccountListItemResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountListItemResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountListItemResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *AccountListItemResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountListItemResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountListItemResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBillingAccountId

`func (o *AccountListItemResponse) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *AccountListItemResponse) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *AccountListItemResponse) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *AccountListItemResponse) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetBillToEntityId

`func (o *AccountListItemResponse) GetBillToEntityId() string`

GetBillToEntityId returns the BillToEntityId field if non-nil, zero value otherwise.

### GetBillToEntityIdOk

`func (o *AccountListItemResponse) GetBillToEntityIdOk() (*string, bool)`

GetBillToEntityIdOk returns a tuple with the BillToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToEntityId

`func (o *AccountListItemResponse) SetBillToEntityId(v string)`

SetBillToEntityId sets BillToEntityId field to given value.

### HasBillToEntityId

`func (o *AccountListItemResponse) HasBillToEntityId() bool`

HasBillToEntityId returns a boolean if a field has been set.

### GetLegalName

`func (o *AccountListItemResponse) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *AccountListItemResponse) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *AccountListItemResponse) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetBillingCountryCode

`func (o *AccountListItemResponse) GetBillingCountryCode() string`

GetBillingCountryCode returns the BillingCountryCode field if non-nil, zero value otherwise.

### GetBillingCountryCodeOk

`func (o *AccountListItemResponse) GetBillingCountryCodeOk() (*string, bool)`

GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountryCode

`func (o *AccountListItemResponse) SetBillingCountryCode(v string)`

SetBillingCountryCode sets BillingCountryCode field to given value.

### HasBillingCountryCode

`func (o *AccountListItemResponse) HasBillingCountryCode() bool`

HasBillingCountryCode returns a boolean if a field has been set.

### GetBillingRegionCode

`func (o *AccountListItemResponse) GetBillingRegionCode() string`

GetBillingRegionCode returns the BillingRegionCode field if non-nil, zero value otherwise.

### GetBillingRegionCodeOk

`func (o *AccountListItemResponse) GetBillingRegionCodeOk() (*string, bool)`

GetBillingRegionCodeOk returns a tuple with the BillingRegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRegionCode

`func (o *AccountListItemResponse) SetBillingRegionCode(v string)`

SetBillingRegionCode sets BillingRegionCode field to given value.

### HasBillingRegionCode

`func (o *AccountListItemResponse) HasBillingRegionCode() bool`

HasBillingRegionCode returns a boolean if a field has been set.

### GetBillingTimezone

`func (o *AccountListItemResponse) GetBillingTimezone() string`

GetBillingTimezone returns the BillingTimezone field if non-nil, zero value otherwise.

### GetBillingTimezoneOk

`func (o *AccountListItemResponse) GetBillingTimezoneOk() (*string, bool)`

GetBillingTimezoneOk returns a tuple with the BillingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTimezone

`func (o *AccountListItemResponse) SetBillingTimezone(v string)`

SetBillingTimezone sets BillingTimezone field to given value.

### HasBillingTimezone

`func (o *AccountListItemResponse) HasBillingTimezone() bool`

HasBillingTimezone returns a boolean if a field has been set.

### GetDefaultLanguageCode

`func (o *AccountListItemResponse) GetDefaultLanguageCode() string`

GetDefaultLanguageCode returns the DefaultLanguageCode field if non-nil, zero value otherwise.

### GetDefaultLanguageCodeOk

`func (o *AccountListItemResponse) GetDefaultLanguageCodeOk() (*string, bool)`

GetDefaultLanguageCodeOk returns a tuple with the DefaultLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageCode

`func (o *AccountListItemResponse) SetDefaultLanguageCode(v string)`

SetDefaultLanguageCode sets DefaultLanguageCode field to given value.


### GetDefaultCurrencyCode

`func (o *AccountListItemResponse) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *AccountListItemResponse) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *AccountListItemResponse) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.


### GetPortalRequired

`func (o *AccountListItemResponse) GetPortalRequired() bool`

GetPortalRequired returns the PortalRequired field if non-nil, zero value otherwise.

### GetPortalRequiredOk

`func (o *AccountListItemResponse) GetPortalRequiredOk() (*bool, bool)`

GetPortalRequiredOk returns a tuple with the PortalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalRequired

`func (o *AccountListItemResponse) SetPortalRequired(v bool)`

SetPortalRequired sets PortalRequired field to given value.


### GetPaymentTerms

`func (o *AccountListItemResponse) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *AccountListItemResponse) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *AccountListItemResponse) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *AccountListItemResponse) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetIsStrategic

`func (o *AccountListItemResponse) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *AccountListItemResponse) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *AccountListItemResponse) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.


### GetSourceSystem

`func (o *AccountListItemResponse) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *AccountListItemResponse) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *AccountListItemResponse) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.


### GetMetadata

`func (o *AccountListItemResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountListItemResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountListItemResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountListItemResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubsidiaries

`func (o *AccountListItemResponse) GetSubsidiaries() []AccountListItemResponseSubsidiariesInner`

GetSubsidiaries returns the Subsidiaries field if non-nil, zero value otherwise.

### GetSubsidiariesOk

`func (o *AccountListItemResponse) GetSubsidiariesOk() (*[]AccountListItemResponseSubsidiariesInner, bool)`

GetSubsidiariesOk returns a tuple with the Subsidiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidiaries

`func (o *AccountListItemResponse) SetSubsidiaries(v []AccountListItemResponseSubsidiariesInner)`

SetSubsidiaries sets Subsidiaries field to given value.

### HasSubsidiaries

`func (o *AccountListItemResponse) HasSubsidiaries() bool`

HasSubsidiaries returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountListItemResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountListItemResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountListItemResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AccountListItemResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountListItemResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountListItemResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOpenBalanceMicros

`func (o *AccountListItemResponse) GetOpenBalanceMicros() int32`

GetOpenBalanceMicros returns the OpenBalanceMicros field if non-nil, zero value otherwise.

### GetOpenBalanceMicrosOk

`func (o *AccountListItemResponse) GetOpenBalanceMicrosOk() (*int32, bool)`

GetOpenBalanceMicrosOk returns a tuple with the OpenBalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBalanceMicros

`func (o *AccountListItemResponse) SetOpenBalanceMicros(v int32)`

SetOpenBalanceMicros sets OpenBalanceMicros field to given value.

### HasOpenBalanceMicros

`func (o *AccountListItemResponse) HasOpenBalanceMicros() bool`

HasOpenBalanceMicros returns a boolean if a field has been set.

### GetOpenBalanceBaseMicros

`func (o *AccountListItemResponse) GetOpenBalanceBaseMicros() int32`

GetOpenBalanceBaseMicros returns the OpenBalanceBaseMicros field if non-nil, zero value otherwise.

### GetOpenBalanceBaseMicrosOk

`func (o *AccountListItemResponse) GetOpenBalanceBaseMicrosOk() (*int32, bool)`

GetOpenBalanceBaseMicrosOk returns a tuple with the OpenBalanceBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBalanceBaseMicros

`func (o *AccountListItemResponse) SetOpenBalanceBaseMicros(v int32)`

SetOpenBalanceBaseMicros sets OpenBalanceBaseMicros field to given value.

### HasOpenBalanceBaseMicros

`func (o *AccountListItemResponse) HasOpenBalanceBaseMicros() bool`

HasOpenBalanceBaseMicros returns a boolean if a field has been set.

### GetOpenBalanceCurrencyCode

`func (o *AccountListItemResponse) GetOpenBalanceCurrencyCode() string`

GetOpenBalanceCurrencyCode returns the OpenBalanceCurrencyCode field if non-nil, zero value otherwise.

### GetOpenBalanceCurrencyCodeOk

`func (o *AccountListItemResponse) GetOpenBalanceCurrencyCodeOk() (*string, bool)`

GetOpenBalanceCurrencyCodeOk returns a tuple with the OpenBalanceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBalanceCurrencyCode

`func (o *AccountListItemResponse) SetOpenBalanceCurrencyCode(v string)`

SetOpenBalanceCurrencyCode sets OpenBalanceCurrencyCode field to given value.

### HasOpenBalanceCurrencyCode

`func (o *AccountListItemResponse) HasOpenBalanceCurrencyCode() bool`

HasOpenBalanceCurrencyCode returns a boolean if a field has been set.

### GetOpenBalanceBaseCurrencyCode

`func (o *AccountListItemResponse) GetOpenBalanceBaseCurrencyCode() string`

GetOpenBalanceBaseCurrencyCode returns the OpenBalanceBaseCurrencyCode field if non-nil, zero value otherwise.

### GetOpenBalanceBaseCurrencyCodeOk

`func (o *AccountListItemResponse) GetOpenBalanceBaseCurrencyCodeOk() (*string, bool)`

GetOpenBalanceBaseCurrencyCodeOk returns a tuple with the OpenBalanceBaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBalanceBaseCurrencyCode

`func (o *AccountListItemResponse) SetOpenBalanceBaseCurrencyCode(v string)`

SetOpenBalanceBaseCurrencyCode sets OpenBalanceBaseCurrencyCode field to given value.

### HasOpenBalanceBaseCurrencyCode

`func (o *AccountListItemResponse) HasOpenBalanceBaseCurrencyCode() bool`

HasOpenBalanceBaseCurrencyCode returns a boolean if a field has been set.

### GetOpenBalanceHasMixedCurrencies

`func (o *AccountListItemResponse) GetOpenBalanceHasMixedCurrencies() bool`

GetOpenBalanceHasMixedCurrencies returns the OpenBalanceHasMixedCurrencies field if non-nil, zero value otherwise.

### GetOpenBalanceHasMixedCurrenciesOk

`func (o *AccountListItemResponse) GetOpenBalanceHasMixedCurrenciesOk() (*bool, bool)`

GetOpenBalanceHasMixedCurrenciesOk returns a tuple with the OpenBalanceHasMixedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBalanceHasMixedCurrencies

`func (o *AccountListItemResponse) SetOpenBalanceHasMixedCurrencies(v bool)`

SetOpenBalanceHasMixedCurrencies sets OpenBalanceHasMixedCurrencies field to given value.

### HasOpenBalanceHasMixedCurrencies

`func (o *AccountListItemResponse) HasOpenBalanceHasMixedCurrencies() bool`

HasOpenBalanceHasMixedCurrencies returns a boolean if a field has been set.

### GetSubsidiaryOptions

`func (o *AccountListItemResponse) GetSubsidiaryOptions() []AccountSubsidiaryOption`

GetSubsidiaryOptions returns the SubsidiaryOptions field if non-nil, zero value otherwise.

### GetSubsidiaryOptionsOk

`func (o *AccountListItemResponse) GetSubsidiaryOptionsOk() (*[]AccountSubsidiaryOption, bool)`

GetSubsidiaryOptionsOk returns a tuple with the SubsidiaryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidiaryOptions

`func (o *AccountListItemResponse) SetSubsidiaryOptions(v []AccountSubsidiaryOption)`

SetSubsidiaryOptions sets SubsidiaryOptions field to given value.

### HasSubsidiaryOptions

`func (o *AccountListItemResponse) HasSubsidiaryOptions() bool`

HasSubsidiaryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


