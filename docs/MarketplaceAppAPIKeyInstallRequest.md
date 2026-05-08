# MarketplaceAppAPIKeyInstallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the application to install.  If name is not provided defaults to the marketplace listing&#39;s name. | [optional] 
**CreateBillingProfile** | Pointer to **bool** | If true, a billing profile will be created for the app. The Stripe app will be also set as the default billing profile if the current default is a Sandbox app. | [optional] [default to true]
**ApiKey** | **string** | The API key for the provider. For example, the Stripe API key. | 

## Methods

### NewMarketplaceAppAPIKeyInstallRequest

`func NewMarketplaceAppAPIKeyInstallRequest(apiKey string, ) *MarketplaceAppAPIKeyInstallRequest`

NewMarketplaceAppAPIKeyInstallRequest instantiates a new MarketplaceAppAPIKeyInstallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceAppAPIKeyInstallRequestWithDefaults

`func NewMarketplaceAppAPIKeyInstallRequestWithDefaults() *MarketplaceAppAPIKeyInstallRequest`

NewMarketplaceAppAPIKeyInstallRequestWithDefaults instantiates a new MarketplaceAppAPIKeyInstallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketplaceAppAPIKeyInstallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceAppAPIKeyInstallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceAppAPIKeyInstallRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceAppAPIKeyInstallRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreateBillingProfile

`func (o *MarketplaceAppAPIKeyInstallRequest) GetCreateBillingProfile() bool`

GetCreateBillingProfile returns the CreateBillingProfile field if non-nil, zero value otherwise.

### GetCreateBillingProfileOk

`func (o *MarketplaceAppAPIKeyInstallRequest) GetCreateBillingProfileOk() (*bool, bool)`

GetCreateBillingProfileOk returns a tuple with the CreateBillingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBillingProfile

`func (o *MarketplaceAppAPIKeyInstallRequest) SetCreateBillingProfile(v bool)`

SetCreateBillingProfile sets CreateBillingProfile field to given value.

### HasCreateBillingProfile

`func (o *MarketplaceAppAPIKeyInstallRequest) HasCreateBillingProfile() bool`

HasCreateBillingProfile returns a boolean if a field has been set.

### GetApiKey

`func (o *MarketplaceAppAPIKeyInstallRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *MarketplaceAppAPIKeyInstallRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *MarketplaceAppAPIKeyInstallRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


