# AppApiV1PortalRouterCreatePortalTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Subject the token is scoped to | 
**ExpiresInSeconds** | Pointer to **int32** | Token TTL (1m to 24h) | [optional] [default to 3600]
**Scopes** | Pointer to **[]string** | Optional scope list | [optional] 

## Methods

### NewAppApiV1PortalRouterCreatePortalTokenRequest

`func NewAppApiV1PortalRouterCreatePortalTokenRequest(subject string, ) *AppApiV1PortalRouterCreatePortalTokenRequest`

NewAppApiV1PortalRouterCreatePortalTokenRequest instantiates a new AppApiV1PortalRouterCreatePortalTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApiV1PortalRouterCreatePortalTokenRequestWithDefaults

`func NewAppApiV1PortalRouterCreatePortalTokenRequestWithDefaults() *AppApiV1PortalRouterCreatePortalTokenRequest`

NewAppApiV1PortalRouterCreatePortalTokenRequestWithDefaults instantiates a new AppApiV1PortalRouterCreatePortalTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetExpiresInSeconds

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) GetExpiresInSeconds() int32`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) GetExpiresInSecondsOk() (*int32, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) SetExpiresInSeconds(v int32)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.

### GetScopes

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AppApiV1PortalRouterCreatePortalTokenRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


