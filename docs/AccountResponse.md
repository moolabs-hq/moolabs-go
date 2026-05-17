# AccountResponse

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

## Methods

### NewAccountResponse

`func NewAccountResponse(id string, tenantId string, customerId string, legalName string, defaultLanguageCode string, defaultCurrencyCode string, portalRequired bool, isStrategic bool, sourceSystem string, createdAt time.Time, updatedAt time.Time, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *AccountResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *AccountResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBillingAccountId

`func (o *AccountResponse) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *AccountResponse) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *AccountResponse) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *AccountResponse) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetBillToEntityId

`func (o *AccountResponse) GetBillToEntityId() string`

GetBillToEntityId returns the BillToEntityId field if non-nil, zero value otherwise.

### GetBillToEntityIdOk

`func (o *AccountResponse) GetBillToEntityIdOk() (*string, bool)`

GetBillToEntityIdOk returns a tuple with the BillToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToEntityId

`func (o *AccountResponse) SetBillToEntityId(v string)`

SetBillToEntityId sets BillToEntityId field to given value.

### HasBillToEntityId

`func (o *AccountResponse) HasBillToEntityId() bool`

HasBillToEntityId returns a boolean if a field has been set.

### GetLegalName

`func (o *AccountResponse) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *AccountResponse) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *AccountResponse) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetBillingCountryCode

`func (o *AccountResponse) GetBillingCountryCode() string`

GetBillingCountryCode returns the BillingCountryCode field if non-nil, zero value otherwise.

### GetBillingCountryCodeOk

`func (o *AccountResponse) GetBillingCountryCodeOk() (*string, bool)`

GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountryCode

`func (o *AccountResponse) SetBillingCountryCode(v string)`

SetBillingCountryCode sets BillingCountryCode field to given value.

### HasBillingCountryCode

`func (o *AccountResponse) HasBillingCountryCode() bool`

HasBillingCountryCode returns a boolean if a field has been set.

### GetBillingRegionCode

`func (o *AccountResponse) GetBillingRegionCode() string`

GetBillingRegionCode returns the BillingRegionCode field if non-nil, zero value otherwise.

### GetBillingRegionCodeOk

`func (o *AccountResponse) GetBillingRegionCodeOk() (*string, bool)`

GetBillingRegionCodeOk returns a tuple with the BillingRegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRegionCode

`func (o *AccountResponse) SetBillingRegionCode(v string)`

SetBillingRegionCode sets BillingRegionCode field to given value.

### HasBillingRegionCode

`func (o *AccountResponse) HasBillingRegionCode() bool`

HasBillingRegionCode returns a boolean if a field has been set.

### GetBillingTimezone

`func (o *AccountResponse) GetBillingTimezone() string`

GetBillingTimezone returns the BillingTimezone field if non-nil, zero value otherwise.

### GetBillingTimezoneOk

`func (o *AccountResponse) GetBillingTimezoneOk() (*string, bool)`

GetBillingTimezoneOk returns a tuple with the BillingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTimezone

`func (o *AccountResponse) SetBillingTimezone(v string)`

SetBillingTimezone sets BillingTimezone field to given value.

### HasBillingTimezone

`func (o *AccountResponse) HasBillingTimezone() bool`

HasBillingTimezone returns a boolean if a field has been set.

### GetDefaultLanguageCode

`func (o *AccountResponse) GetDefaultLanguageCode() string`

GetDefaultLanguageCode returns the DefaultLanguageCode field if non-nil, zero value otherwise.

### GetDefaultLanguageCodeOk

`func (o *AccountResponse) GetDefaultLanguageCodeOk() (*string, bool)`

GetDefaultLanguageCodeOk returns a tuple with the DefaultLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageCode

`func (o *AccountResponse) SetDefaultLanguageCode(v string)`

SetDefaultLanguageCode sets DefaultLanguageCode field to given value.


### GetDefaultCurrencyCode

`func (o *AccountResponse) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *AccountResponse) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *AccountResponse) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.


### GetPortalRequired

`func (o *AccountResponse) GetPortalRequired() bool`

GetPortalRequired returns the PortalRequired field if non-nil, zero value otherwise.

### GetPortalRequiredOk

`func (o *AccountResponse) GetPortalRequiredOk() (*bool, bool)`

GetPortalRequiredOk returns a tuple with the PortalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalRequired

`func (o *AccountResponse) SetPortalRequired(v bool)`

SetPortalRequired sets PortalRequired field to given value.


### GetPaymentTerms

`func (o *AccountResponse) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *AccountResponse) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *AccountResponse) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *AccountResponse) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetIsStrategic

`func (o *AccountResponse) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *AccountResponse) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *AccountResponse) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.


### GetSourceSystem

`func (o *AccountResponse) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *AccountResponse) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *AccountResponse) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.


### GetMetadata

`func (o *AccountResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubsidiaries

`func (o *AccountResponse) GetSubsidiaries() []AccountListItemResponseSubsidiariesInner`

GetSubsidiaries returns the Subsidiaries field if non-nil, zero value otherwise.

### GetSubsidiariesOk

`func (o *AccountResponse) GetSubsidiariesOk() (*[]AccountListItemResponseSubsidiariesInner, bool)`

GetSubsidiariesOk returns a tuple with the Subsidiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidiaries

`func (o *AccountResponse) SetSubsidiaries(v []AccountListItemResponseSubsidiariesInner)`

SetSubsidiaries sets Subsidiaries field to given value.

### HasSubsidiaries

`func (o *AccountResponse) HasSubsidiaries() bool`

HasSubsidiaries returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AccountResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


