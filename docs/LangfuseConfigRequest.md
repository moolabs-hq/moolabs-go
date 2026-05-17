# LangfuseConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** |  | 
**SecretKey** | **string** |  | 
**Host** | Pointer to **string** |  | [optional] [default to "https://cloud.langfuse.com"]

## Methods

### NewLangfuseConfigRequest

`func NewLangfuseConfigRequest(publicKey string, secretKey string, ) *LangfuseConfigRequest`

NewLangfuseConfigRequest instantiates a new LangfuseConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLangfuseConfigRequestWithDefaults

`func NewLangfuseConfigRequestWithDefaults() *LangfuseConfigRequest`

NewLangfuseConfigRequestWithDefaults instantiates a new LangfuseConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *LangfuseConfigRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LangfuseConfigRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LangfuseConfigRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSecretKey

`func (o *LangfuseConfigRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LangfuseConfigRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LangfuseConfigRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetHost

`func (o *LangfuseConfigRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LangfuseConfigRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LangfuseConfigRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LangfuseConfigRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


