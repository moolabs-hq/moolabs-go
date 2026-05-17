# AccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CustomerId** | **string** |  | 
**LegalName** | **string** |  | 
**BillingCountryCode** | Pointer to **string** |  | [optional] 
**DefaultCurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**IsStrategic** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAccountSummary

`func NewAccountSummary(id string, customerId string, legalName string, ) *AccountSummary`

NewAccountSummary instantiates a new AccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSummaryWithDefaults

`func NewAccountSummaryWithDefaults() *AccountSummary`

NewAccountSummaryWithDefaults instantiates a new AccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSummary) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *AccountSummary) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountSummary) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountSummary) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetLegalName

`func (o *AccountSummary) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *AccountSummary) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *AccountSummary) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetBillingCountryCode

`func (o *AccountSummary) GetBillingCountryCode() string`

GetBillingCountryCode returns the BillingCountryCode field if non-nil, zero value otherwise.

### GetBillingCountryCodeOk

`func (o *AccountSummary) GetBillingCountryCodeOk() (*string, bool)`

GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountryCode

`func (o *AccountSummary) SetBillingCountryCode(v string)`

SetBillingCountryCode sets BillingCountryCode field to given value.

### HasBillingCountryCode

`func (o *AccountSummary) HasBillingCountryCode() bool`

HasBillingCountryCode returns a boolean if a field has been set.

### GetDefaultCurrencyCode

`func (o *AccountSummary) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *AccountSummary) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *AccountSummary) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.

### HasDefaultCurrencyCode

`func (o *AccountSummary) HasDefaultCurrencyCode() bool`

HasDefaultCurrencyCode returns a boolean if a field has been set.

### GetIsStrategic

`func (o *AccountSummary) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *AccountSummary) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *AccountSummary) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *AccountSummary) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


