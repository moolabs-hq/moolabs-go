# InvalidatePortalTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvalidatePortalTokenRequest

`func NewInvalidatePortalTokenRequest() *InvalidatePortalTokenRequest`

NewInvalidatePortalTokenRequest instantiates a new InvalidatePortalTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidatePortalTokenRequestWithDefaults

`func NewInvalidatePortalTokenRequestWithDefaults() *InvalidatePortalTokenRequest`

NewInvalidatePortalTokenRequestWithDefaults instantiates a new InvalidatePortalTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *InvalidatePortalTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InvalidatePortalTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InvalidatePortalTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InvalidatePortalTokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *InvalidatePortalTokenRequest) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *InvalidatePortalTokenRequest) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSubject

`func (o *InvalidatePortalTokenRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InvalidatePortalTokenRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InvalidatePortalTokenRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *InvalidatePortalTokenRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *InvalidatePortalTokenRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *InvalidatePortalTokenRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


