# CustomerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NetsuiteId** | **string** |  | 
**EntityIdDisplay** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsPerson** | **bool** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **string** |  | [optional] 
**ManagedThroughPortal** | Pointer to **map[string]interface{}** |  | [optional] 
**ManagedThroughPortalSourceFieldId** | Pointer to **string** |  | [optional] 
**NsLastModified** | Pointer to **time.Time** |  | [optional] 
**SyncedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCustomerItem

`func NewCustomerItem(id string, netsuiteId string, isPerson bool, ) *CustomerItem`

NewCustomerItem instantiates a new CustomerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerItemWithDefaults

`func NewCustomerItemWithDefaults() *CustomerItem`

NewCustomerItemWithDefaults instantiates a new CustomerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerItem) SetId(v string)`

SetId sets Id field to given value.


### GetNetsuiteId

`func (o *CustomerItem) GetNetsuiteId() string`

GetNetsuiteId returns the NetsuiteId field if non-nil, zero value otherwise.

### GetNetsuiteIdOk

`func (o *CustomerItem) GetNetsuiteIdOk() (*string, bool)`

GetNetsuiteIdOk returns a tuple with the NetsuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteId

`func (o *CustomerItem) SetNetsuiteId(v string)`

SetNetsuiteId sets NetsuiteId field to given value.


### GetEntityIdDisplay

`func (o *CustomerItem) GetEntityIdDisplay() string`

GetEntityIdDisplay returns the EntityIdDisplay field if non-nil, zero value otherwise.

### GetEntityIdDisplayOk

`func (o *CustomerItem) GetEntityIdDisplayOk() (*string, bool)`

GetEntityIdDisplayOk returns a tuple with the EntityIdDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIdDisplay

`func (o *CustomerItem) SetEntityIdDisplay(v string)`

SetEntityIdDisplay sets EntityIdDisplay field to given value.

### HasEntityIdDisplay

`func (o *CustomerItem) HasEntityIdDisplay() bool`

HasEntityIdDisplay returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerItem) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerItem) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerItem) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerItem) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerItem) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerItem) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerItem) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerItem) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomerItem) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerItem) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerItem) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomerItem) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsPerson

`func (o *CustomerItem) GetIsPerson() bool`

GetIsPerson returns the IsPerson field if non-nil, zero value otherwise.

### GetIsPersonOk

`func (o *CustomerItem) GetIsPersonOk() (*bool, bool)`

GetIsPersonOk returns a tuple with the IsPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerson

`func (o *CustomerItem) SetIsPerson(v bool)`

SetIsPerson sets IsPerson field to given value.


### GetCurrencyCode

`func (o *CustomerItem) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CustomerItem) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CustomerItem) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CustomerItem) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTerms

`func (o *CustomerItem) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *CustomerItem) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *CustomerItem) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *CustomerItem) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetManagedThroughPortal

`func (o *CustomerItem) GetManagedThroughPortal() map[string]interface{}`

GetManagedThroughPortal returns the ManagedThroughPortal field if non-nil, zero value otherwise.

### GetManagedThroughPortalOk

`func (o *CustomerItem) GetManagedThroughPortalOk() (*map[string]interface{}, bool)`

GetManagedThroughPortalOk returns a tuple with the ManagedThroughPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedThroughPortal

`func (o *CustomerItem) SetManagedThroughPortal(v map[string]interface{})`

SetManagedThroughPortal sets ManagedThroughPortal field to given value.

### HasManagedThroughPortal

`func (o *CustomerItem) HasManagedThroughPortal() bool`

HasManagedThroughPortal returns a boolean if a field has been set.

### GetManagedThroughPortalSourceFieldId

`func (o *CustomerItem) GetManagedThroughPortalSourceFieldId() string`

GetManagedThroughPortalSourceFieldId returns the ManagedThroughPortalSourceFieldId field if non-nil, zero value otherwise.

### GetManagedThroughPortalSourceFieldIdOk

`func (o *CustomerItem) GetManagedThroughPortalSourceFieldIdOk() (*string, bool)`

GetManagedThroughPortalSourceFieldIdOk returns a tuple with the ManagedThroughPortalSourceFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedThroughPortalSourceFieldId

`func (o *CustomerItem) SetManagedThroughPortalSourceFieldId(v string)`

SetManagedThroughPortalSourceFieldId sets ManagedThroughPortalSourceFieldId field to given value.

### HasManagedThroughPortalSourceFieldId

`func (o *CustomerItem) HasManagedThroughPortalSourceFieldId() bool`

HasManagedThroughPortalSourceFieldId returns a boolean if a field has been set.

### GetNsLastModified

`func (o *CustomerItem) GetNsLastModified() time.Time`

GetNsLastModified returns the NsLastModified field if non-nil, zero value otherwise.

### GetNsLastModifiedOk

`func (o *CustomerItem) GetNsLastModifiedOk() (*time.Time, bool)`

GetNsLastModifiedOk returns a tuple with the NsLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsLastModified

`func (o *CustomerItem) SetNsLastModified(v time.Time)`

SetNsLastModified sets NsLastModified field to given value.

### HasNsLastModified

`func (o *CustomerItem) HasNsLastModified() bool`

HasNsLastModified returns a boolean if a field has been set.

### GetSyncedAt

`func (o *CustomerItem) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *CustomerItem) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *CustomerItem) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *CustomerItem) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


