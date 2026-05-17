# CustomerUpsertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetsuiteCustomerId** | **string** |  | 
**MoometerCustomerId** | **string** |  | 
**CompanyName** | **string** |  | 
**CustomerCode** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**Status** | Pointer to **string** |  | [optional] [default to "active"]
**SourceUpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCustomerUpsertRequest

`func NewCustomerUpsertRequest(netsuiteCustomerId string, moometerCustomerId string, companyName string, ) *CustomerUpsertRequest`

NewCustomerUpsertRequest instantiates a new CustomerUpsertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpsertRequestWithDefaults

`func NewCustomerUpsertRequestWithDefaults() *CustomerUpsertRequest`

NewCustomerUpsertRequestWithDefaults instantiates a new CustomerUpsertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetsuiteCustomerId

`func (o *CustomerUpsertRequest) GetNetsuiteCustomerId() string`

GetNetsuiteCustomerId returns the NetsuiteCustomerId field if non-nil, zero value otherwise.

### GetNetsuiteCustomerIdOk

`func (o *CustomerUpsertRequest) GetNetsuiteCustomerIdOk() (*string, bool)`

GetNetsuiteCustomerIdOk returns a tuple with the NetsuiteCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteCustomerId

`func (o *CustomerUpsertRequest) SetNetsuiteCustomerId(v string)`

SetNetsuiteCustomerId sets NetsuiteCustomerId field to given value.


### GetMoometerCustomerId

`func (o *CustomerUpsertRequest) GetMoometerCustomerId() string`

GetMoometerCustomerId returns the MoometerCustomerId field if non-nil, zero value otherwise.

### GetMoometerCustomerIdOk

`func (o *CustomerUpsertRequest) GetMoometerCustomerIdOk() (*string, bool)`

GetMoometerCustomerIdOk returns a tuple with the MoometerCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoometerCustomerId

`func (o *CustomerUpsertRequest) SetMoometerCustomerId(v string)`

SetMoometerCustomerId sets MoometerCustomerId field to given value.


### GetCompanyName

`func (o *CustomerUpsertRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerUpsertRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerUpsertRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCustomerCode

`func (o *CustomerUpsertRequest) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *CustomerUpsertRequest) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *CustomerUpsertRequest) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### HasCustomerCode

`func (o *CustomerUpsertRequest) HasCustomerCode() bool`

HasCustomerCode returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerUpsertRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerUpsertRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerUpsertRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerUpsertRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CustomerUpsertRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CustomerUpsertRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CustomerUpsertRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CustomerUpsertRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerUpsertRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerUpsertRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerUpsertRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerUpsertRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSourceUpdatedAt

`func (o *CustomerUpsertRequest) GetSourceUpdatedAt() time.Time`

GetSourceUpdatedAt returns the SourceUpdatedAt field if non-nil, zero value otherwise.

### GetSourceUpdatedAtOk

`func (o *CustomerUpsertRequest) GetSourceUpdatedAtOk() (*time.Time, bool)`

GetSourceUpdatedAtOk returns a tuple with the SourceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUpdatedAt

`func (o *CustomerUpsertRequest) SetSourceUpdatedAt(v time.Time)`

SetSourceUpdatedAt sets SourceUpdatedAt field to given value.

### HasSourceUpdatedAt

`func (o *CustomerUpsertRequest) HasSourceUpdatedAt() bool`

HasSourceUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


