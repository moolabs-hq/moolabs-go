# PortalToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ULID (Universally Unique Lexicographically Sortable Identifier). | [optional] [readonly] 
**Subject** | **string** |  | 
**ExpiresAt** | Pointer to **time.Time** | [RFC3339](https://tools.ietf.org/html/rfc3339) formatted date-time string in UTC. | [optional] [readonly] 
**Expired** | Pointer to **bool** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | [RFC3339](https://tools.ietf.org/html/rfc3339) formatted date-time string in UTC. | [optional] [readonly] 
**Token** | Pointer to **string** | The token is only returned at creation. | [optional] [readonly] 
**AllowedMeterSlugs** | Pointer to **[]string** | Optional, if defined only the specified meters will be allowed. | [optional] 

## Methods

### NewPortalToken

`func NewPortalToken(subject string, ) *PortalToken`

NewPortalToken instantiates a new PortalToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalTokenWithDefaults

`func NewPortalTokenWithDefaults() *PortalToken`

NewPortalTokenWithDefaults instantiates a new PortalToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortalToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *PortalToken) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PortalToken) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PortalToken) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetExpiresAt

`func (o *PortalToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PortalToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PortalToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PortalToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpired

`func (o *PortalToken) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *PortalToken) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *PortalToken) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *PortalToken) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortalToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortalToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortalToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortalToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetToken

`func (o *PortalToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PortalToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PortalToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PortalToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAllowedMeterSlugs

`func (o *PortalToken) GetAllowedMeterSlugs() []string`

GetAllowedMeterSlugs returns the AllowedMeterSlugs field if non-nil, zero value otherwise.

### GetAllowedMeterSlugsOk

`func (o *PortalToken) GetAllowedMeterSlugsOk() (*[]string, bool)`

GetAllowedMeterSlugsOk returns a tuple with the AllowedMeterSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMeterSlugs

`func (o *PortalToken) SetAllowedMeterSlugs(v []string)`

SetAllowedMeterSlugs sets AllowedMeterSlugs field to given value.

### HasAllowedMeterSlugs

`func (o *PortalToken) HasAllowedMeterSlugs() bool`

HasAllowedMeterSlugs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


