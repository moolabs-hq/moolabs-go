# AppApiV1TenantAccessRouterCreatePortalTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** |  | 
**ExpiresInSeconds** | Pointer to **int32** |  | [optional] [default to 3600]

## Methods

### NewAppApiV1TenantAccessRouterCreatePortalTokenRequest

`func NewAppApiV1TenantAccessRouterCreatePortalTokenRequest(subject string, ) *AppApiV1TenantAccessRouterCreatePortalTokenRequest`

NewAppApiV1TenantAccessRouterCreatePortalTokenRequest instantiates a new AppApiV1TenantAccessRouterCreatePortalTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiV1TenantAccessRouterCreatePortalTokenRequestWithDefaults

`func NewAppApiV1TenantAccessRouterCreatePortalTokenRequestWithDefaults() *AppApiV1TenantAccessRouterCreatePortalTokenRequest`

NewAppApiV1TenantAccessRouterCreatePortalTokenRequestWithDefaults instantiates a new AppApiV1TenantAccessRouterCreatePortalTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetExpiresInSeconds

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) GetExpiresInSeconds() int32`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) GetExpiresInSecondsOk() (*int32, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) SetExpiresInSeconds(v int32)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *AppApiV1TenantAccessRouterCreatePortalTokenRequest) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


