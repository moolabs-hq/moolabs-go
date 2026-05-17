# AccountCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | MooMeter customer_id | 
**BillingAccountId** | Pointer to **string** |  | [optional] 
**BillToEntityId** | Pointer to **string** |  | [optional] 
**LegalName** | **string** |  | 
**BillingCountryCode** | Pointer to **string** |  | [optional] 
**BillingRegionCode** | Pointer to **string** |  | [optional] 
**BillingTimezone** | Pointer to **string** |  | [optional] 
**DefaultLanguageCode** | Pointer to **string** |  | [optional] [default to "en"]
**DefaultCurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**PortalRequired** | Pointer to **bool** |  | [optional] [default to false]
**IsStrategic** | Pointer to **bool** |  | [optional] [default to false]
**PaymentTerms** | Pointer to **string** |  | [optional] 
**SourceSystem** | Pointer to **string** |  | [optional] [default to "moometer"]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAccountCreate

`func NewAccountCreate(customerId string, legalName string, ) *AccountCreate`

NewAccountCreate instantiates a new AccountCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateWithDefaults

`func NewAccountCreateWithDefaults() *AccountCreate`

NewAccountCreateWithDefaults instantiates a new AccountCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AccountCreate) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountCreate) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountCreate) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBillingAccountId

`func (o *AccountCreate) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *AccountCreate) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *AccountCreate) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *AccountCreate) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetBillToEntityId

`func (o *AccountCreate) GetBillToEntityId() string`

GetBillToEntityId returns the BillToEntityId field if non-nil, zero value otherwise.

### GetBillToEntityIdOk

`func (o *AccountCreate) GetBillToEntityIdOk() (*string, bool)`

GetBillToEntityIdOk returns a tuple with the BillToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToEntityId

`func (o *AccountCreate) SetBillToEntityId(v string)`

SetBillToEntityId sets BillToEntityId field to given value.

### HasBillToEntityId

`func (o *AccountCreate) HasBillToEntityId() bool`

HasBillToEntityId returns a boolean if a field has been set.

### GetLegalName

`func (o *AccountCreate) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *AccountCreate) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *AccountCreate) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetBillingCountryCode

`func (o *AccountCreate) GetBillingCountryCode() string`

GetBillingCountryCode returns the BillingCountryCode field if non-nil, zero value otherwise.

### GetBillingCountryCodeOk

`func (o *AccountCreate) GetBillingCountryCodeOk() (*string, bool)`

GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountryCode

`func (o *AccountCreate) SetBillingCountryCode(v string)`

SetBillingCountryCode sets BillingCountryCode field to given value.

### HasBillingCountryCode

`func (o *AccountCreate) HasBillingCountryCode() bool`

HasBillingCountryCode returns a boolean if a field has been set.

### GetBillingRegionCode

`func (o *AccountCreate) GetBillingRegionCode() string`

GetBillingRegionCode returns the BillingRegionCode field if non-nil, zero value otherwise.

### GetBillingRegionCodeOk

`func (o *AccountCreate) GetBillingRegionCodeOk() (*string, bool)`

GetBillingRegionCodeOk returns a tuple with the BillingRegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRegionCode

`func (o *AccountCreate) SetBillingRegionCode(v string)`

SetBillingRegionCode sets BillingRegionCode field to given value.

### HasBillingRegionCode

`func (o *AccountCreate) HasBillingRegionCode() bool`

HasBillingRegionCode returns a boolean if a field has been set.

### GetBillingTimezone

`func (o *AccountCreate) GetBillingTimezone() string`

GetBillingTimezone returns the BillingTimezone field if non-nil, zero value otherwise.

### GetBillingTimezoneOk

`func (o *AccountCreate) GetBillingTimezoneOk() (*string, bool)`

GetBillingTimezoneOk returns a tuple with the BillingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTimezone

`func (o *AccountCreate) SetBillingTimezone(v string)`

SetBillingTimezone sets BillingTimezone field to given value.

### HasBillingTimezone

`func (o *AccountCreate) HasBillingTimezone() bool`

HasBillingTimezone returns a boolean if a field has been set.

### GetDefaultLanguageCode

`func (o *AccountCreate) GetDefaultLanguageCode() string`

GetDefaultLanguageCode returns the DefaultLanguageCode field if non-nil, zero value otherwise.

### GetDefaultLanguageCodeOk

`func (o *AccountCreate) GetDefaultLanguageCodeOk() (*string, bool)`

GetDefaultLanguageCodeOk returns a tuple with the DefaultLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageCode

`func (o *AccountCreate) SetDefaultLanguageCode(v string)`

SetDefaultLanguageCode sets DefaultLanguageCode field to given value.

### HasDefaultLanguageCode

`func (o *AccountCreate) HasDefaultLanguageCode() bool`

HasDefaultLanguageCode returns a boolean if a field has been set.

### GetDefaultCurrencyCode

`func (o *AccountCreate) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *AccountCreate) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *AccountCreate) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *AccountCreate) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetPortalRequired

`func (o *AccountCreate) GetPortalRequired() bool`

GetPortalRequired returns the PortalRequired field if non-nil, zero value otherwise.

### GetPortalRequiredOk

`func (o *AccountCreate) GetPortalRequiredOk() (*bool, bool)`

GetPortalRequiredOk returns a tuple with the PortalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalRequired

`func (o *AccountCreate) SetPortalRequired(v bool)`

SetPortalRequired sets PortalRequired field to given value.

### HasPortalRequired

`func (o *AccountCreate) HasPortalRequired() bool`

HasPortalRequired returns a boolean if a field has been set.

### GetIsStrategic

`func (o *AccountCreate) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *AccountCreate) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *AccountCreate) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *AccountCreate) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *AccountCreate) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *AccountCreate) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *AccountCreate) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *AccountCreate) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetSourceSystem

`func (o *AccountCreate) GetSourceSystem() string`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *AccountCreate) GetSourceSystemOk() (*string, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *AccountCreate) SetSourceSystem(v string)`

SetSourceSystem sets SourceSystem field to given value.

### HasSourceSystem

`func (o *AccountCreate) HasSourceSystem() bool`

HasSourceSystem returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


