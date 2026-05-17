# ContactCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Role** | **string** | One of: ap, finance_admin, procurement, buyer, legal, exec, other | 
**IsPrimaryAp** | Pointer to **bool** |  | [optional] [default to false]
**Locale** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewContactCreate

`func NewContactCreate(fullName string, role string, ) *ContactCreate`

NewContactCreate instantiates a new ContactCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreateWithDefaults

`func NewContactCreateWithDefaults() *ContactCreate`

NewContactCreateWithDefaults instantiates a new ContactCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *ContactCreate) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ContactCreate) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ContactCreate) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *ContactCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *ContactCreate) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactCreate) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactCreate) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ContactCreate) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetRole

`func (o *ContactCreate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContactCreate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContactCreate) SetRole(v string)`

SetRole sets Role field to given value.


### GetIsPrimaryAp

`func (o *ContactCreate) GetIsPrimaryAp() bool`

GetIsPrimaryAp returns the IsPrimaryAp field if non-nil, zero value otherwise.

### GetIsPrimaryApOk

`func (o *ContactCreate) GetIsPrimaryApOk() (*bool, bool)`

GetIsPrimaryApOk returns a tuple with the IsPrimaryAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryAp

`func (o *ContactCreate) SetIsPrimaryAp(v bool)`

SetIsPrimaryAp sets IsPrimaryAp field to given value.

### HasIsPrimaryAp

`func (o *ContactCreate) HasIsPrimaryAp() bool`

HasIsPrimaryAp returns a boolean if a field has been set.

### GetLocale

`func (o *ContactCreate) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ContactCreate) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ContactCreate) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ContactCreate) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTimezone

`func (o *ContactCreate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ContactCreate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ContactCreate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ContactCreate) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


