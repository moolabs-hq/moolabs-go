# ContactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AccountId** | **string** |  | 
**FullName** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 
**IsPrimaryAp** | **bool** |  | 
**IsActive** | **bool** |  | 
**EmailVerifiedAt** | Pointer to **time.Time** |  | [optional] 
**PhoneVerifiedAt** | Pointer to **time.Time** |  | [optional] 
**VerificationSource** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewContactResponse

`func NewContactResponse(id string, tenantId string, accountId string, fullName string, role string, isPrimaryAp bool, isActive bool, createdAt time.Time, updatedAt time.Time, ) *ContactResponse`

NewContactResponse instantiates a new ContactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactResponseWithDefaults

`func NewContactResponseWithDefaults() *ContactResponse`

NewContactResponseWithDefaults instantiates a new ContactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ContactResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContactResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContactResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAccountId

`func (o *ContactResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ContactResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ContactResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetFullName

`func (o *ContactResponse) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ContactResponse) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ContactResponse) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *ContactResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *ContactResponse) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactResponse) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactResponse) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ContactResponse) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetRole

`func (o *ContactResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContactResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContactResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetIsPrimaryAp

`func (o *ContactResponse) GetIsPrimaryAp() bool`

GetIsPrimaryAp returns the IsPrimaryAp field if non-nil, zero value otherwise.

### GetIsPrimaryApOk

`func (o *ContactResponse) GetIsPrimaryApOk() (*bool, bool)`

GetIsPrimaryApOk returns a tuple with the IsPrimaryAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryAp

`func (o *ContactResponse) SetIsPrimaryAp(v bool)`

SetIsPrimaryAp sets IsPrimaryAp field to given value.


### GetIsActive

`func (o *ContactResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ContactResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ContactResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetEmailVerifiedAt

`func (o *ContactResponse) GetEmailVerifiedAt() time.Time`

GetEmailVerifiedAt returns the EmailVerifiedAt field if non-nil, zero value otherwise.

### GetEmailVerifiedAtOk

`func (o *ContactResponse) GetEmailVerifiedAtOk() (*time.Time, bool)`

GetEmailVerifiedAtOk returns a tuple with the EmailVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerifiedAt

`func (o *ContactResponse) SetEmailVerifiedAt(v time.Time)`

SetEmailVerifiedAt sets EmailVerifiedAt field to given value.

### HasEmailVerifiedAt

`func (o *ContactResponse) HasEmailVerifiedAt() bool`

HasEmailVerifiedAt returns a boolean if a field has been set.

### GetPhoneVerifiedAt

`func (o *ContactResponse) GetPhoneVerifiedAt() time.Time`

GetPhoneVerifiedAt returns the PhoneVerifiedAt field if non-nil, zero value otherwise.

### GetPhoneVerifiedAtOk

`func (o *ContactResponse) GetPhoneVerifiedAtOk() (*time.Time, bool)`

GetPhoneVerifiedAtOk returns a tuple with the PhoneVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneVerifiedAt

`func (o *ContactResponse) SetPhoneVerifiedAt(v time.Time)`

SetPhoneVerifiedAt sets PhoneVerifiedAt field to given value.

### HasPhoneVerifiedAt

`func (o *ContactResponse) HasPhoneVerifiedAt() bool`

HasPhoneVerifiedAt returns a boolean if a field has been set.

### GetVerificationSource

`func (o *ContactResponse) GetVerificationSource() string`

GetVerificationSource returns the VerificationSource field if non-nil, zero value otherwise.

### GetVerificationSourceOk

`func (o *ContactResponse) GetVerificationSourceOk() (*string, bool)`

GetVerificationSourceOk returns a tuple with the VerificationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationSource

`func (o *ContactResponse) SetVerificationSource(v string)`

SetVerificationSource sets VerificationSource field to given value.

### HasVerificationSource

`func (o *ContactResponse) HasVerificationSource() bool`

HasVerificationSource returns a boolean if a field has been set.

### GetLocale

`func (o *ContactResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ContactResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ContactResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ContactResponse) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTimezone

`func (o *ContactResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ContactResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ContactResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ContactResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContactResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContactResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContactResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ContactResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContactResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContactResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


