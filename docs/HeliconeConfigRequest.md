# HeliconeConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** |  | 
**WebhookSecret** | Pointer to **string** |  | [optional] 
**PollIntervalMinutes** | Pointer to **int32** |  | [optional] [default to 5]

## Methods

### NewHeliconeConfigRequest

`func NewHeliconeConfigRequest(apiKey string, ) *HeliconeConfigRequest`

NewHeliconeConfigRequest instantiates a new HeliconeConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliconeConfigRequestWithDefaults

`func NewHeliconeConfigRequestWithDefaults() *HeliconeConfigRequest`

NewHeliconeConfigRequestWithDefaults instantiates a new HeliconeConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *HeliconeConfigRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HeliconeConfigRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HeliconeConfigRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetWebhookSecret

`func (o *HeliconeConfigRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *HeliconeConfigRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *HeliconeConfigRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *HeliconeConfigRequest) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### GetPollIntervalMinutes

`func (o *HeliconeConfigRequest) GetPollIntervalMinutes() int32`

GetPollIntervalMinutes returns the PollIntervalMinutes field if non-nil, zero value otherwise.

### GetPollIntervalMinutesOk

`func (o *HeliconeConfigRequest) GetPollIntervalMinutesOk() (*int32, bool)`

GetPollIntervalMinutesOk returns a tuple with the PollIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollIntervalMinutes

`func (o *HeliconeConfigRequest) SetPollIntervalMinutes(v int32)`

SetPollIntervalMinutes sets PollIntervalMinutes field to given value.

### HasPollIntervalMinutes

`func (o *HeliconeConfigRequest) HasPollIntervalMinutes() bool`

HasPollIntervalMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


