# AccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | Pointer to **string** |  | [optional] 
**BillingCountryCode** | Pointer to **string** |  | [optional] 
**BillingRegionCode** | Pointer to **string** |  | [optional] 
**BillingTimezone** | Pointer to **string** |  | [optional] 
**DefaultLanguageCode** | Pointer to **string** |  | [optional] 
**DefaultCurrencyCode** | Pointer to **string** |  | [optional] 
**PortalRequired** | Pointer to **bool** |  | [optional] 
**PaymentTerms** | Pointer to **string** |  | [optional] 
**IsStrategic** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAccountUpdate

`func NewAccountUpdate() *AccountUpdate`

NewAccountUpdate instantiates a new AccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateWithDefaults

`func NewAccountUpdateWithDefaults() *AccountUpdate`

NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *AccountUpdate) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *AccountUpdate) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *AccountUpdate) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *AccountUpdate) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetBillingCountryCode

`func (o *AccountUpdate) GetBillingCountryCode() string`

GetBillingCountryCode returns the BillingCountryCode field if non-nil, zero value otherwise.

### GetBillingCountryCodeOk

`func (o *AccountUpdate) GetBillingCountryCodeOk() (*string, bool)`

GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountryCode

`func (o *AccountUpdate) SetBillingCountryCode(v string)`

SetBillingCountryCode sets BillingCountryCode field to given value.

### HasBillingCountryCode

`func (o *AccountUpdate) HasBillingCountryCode() bool`

HasBillingCountryCode returns a boolean if a field has been set.

### GetBillingRegionCode

`func (o *AccountUpdate) GetBillingRegionCode() string`

GetBillingRegionCode returns the BillingRegionCode field if non-nil, zero value otherwise.

### GetBillingRegionCodeOk

`func (o *AccountUpdate) GetBillingRegionCodeOk() (*string, bool)`

GetBillingRegionCodeOk returns a tuple with the BillingRegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRegionCode

`func (o *AccountUpdate) SetBillingRegionCode(v string)`

SetBillingRegionCode sets BillingRegionCode field to given value.

### HasBillingRegionCode

`func (o *AccountUpdate) HasBillingRegionCode() bool`

HasBillingRegionCode returns a boolean if a field has been set.

### GetBillingTimezone

`func (o *AccountUpdate) GetBillingTimezone() string`

GetBillingTimezone returns the BillingTimezone field if non-nil, zero value otherwise.

### GetBillingTimezoneOk

`func (o *AccountUpdate) GetBillingTimezoneOk() (*string, bool)`

GetBillingTimezoneOk returns a tuple with the BillingTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTimezone

`func (o *AccountUpdate) SetBillingTimezone(v string)`

SetBillingTimezone sets BillingTimezone field to given value.

### HasBillingTimezone

`func (o *AccountUpdate) HasBillingTimezone() bool`

HasBillingTimezone returns a boolean if a field has been set.

### GetDefaultLanguageCode

`func (o *AccountUpdate) GetDefaultLanguageCode() string`

GetDefaultLanguageCode returns the DefaultLanguageCode field if non-nil, zero value otherwise.

### GetDefaultLanguageCodeOk

`func (o *AccountUpdate) GetDefaultLanguageCodeOk() (*string, bool)`

GetDefaultLanguageCodeOk returns a tuple with the DefaultLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguageCode

`func (o *AccountUpdate) SetDefaultLanguageCode(v string)`

SetDefaultLanguageCode sets DefaultLanguageCode field to given value.

### HasDefaultLanguageCode

`func (o *AccountUpdate) HasDefaultLanguageCode() bool`

HasDefaultLanguageCode returns a boolean if a field has been set.

### GetDefaultCurrencyCode

`func (o *AccountUpdate) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *AccountUpdate) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *AccountUpdate) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *AccountUpdate) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetPortalRequired

`func (o *AccountUpdate) GetPortalRequired() bool`

GetPortalRequired returns the PortalRequired field if non-nil, zero value otherwise.

### GetPortalRequiredOk

`func (o *AccountUpdate) GetPortalRequiredOk() (*bool, bool)`

GetPortalRequiredOk returns a tuple with the PortalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalRequired

`func (o *AccountUpdate) SetPortalRequired(v bool)`

SetPortalRequired sets PortalRequired field to given value.

### HasPortalRequired

`func (o *AccountUpdate) HasPortalRequired() bool`

HasPortalRequired returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *AccountUpdate) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *AccountUpdate) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *AccountUpdate) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *AccountUpdate) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetIsStrategic

`func (o *AccountUpdate) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *AccountUpdate) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *AccountUpdate) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *AccountUpdate) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountUpdate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountUpdate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountUpdate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


