# MarketplaceInstallRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the application to install.  If name is not provided defaults to the marketplace listing&#39;s name. | [optional] 
**CreateBillingProfile** | Pointer to **bool** | If true, a billing profile will be created for the app. The Stripe app will be also set as the default billing profile if the current default is a Sandbox app. | [optional] [default to true]

## Methods

### NewMarketplaceInstallRequestPayload

`func NewMarketplaceInstallRequestPayload() *MarketplaceInstallRequestPayload`

NewMarketplaceInstallRequestPayload instantiates a new MarketplaceInstallRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceInstallRequestPayloadWithDefaults

`func NewMarketplaceInstallRequestPayloadWithDefaults() *MarketplaceInstallRequestPayload`

NewMarketplaceInstallRequestPayloadWithDefaults instantiates a new MarketplaceInstallRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketplaceInstallRequestPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceInstallRequestPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceInstallRequestPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceInstallRequestPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreateBillingProfile

`func (o *MarketplaceInstallRequestPayload) GetCreateBillingProfile() bool`

GetCreateBillingProfile returns the CreateBillingProfile field if non-nil, zero value otherwise.

### GetCreateBillingProfileOk

`func (o *MarketplaceInstallRequestPayload) GetCreateBillingProfileOk() (*bool, bool)`

GetCreateBillingProfileOk returns a tuple with the CreateBillingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBillingProfile

`func (o *MarketplaceInstallRequestPayload) SetCreateBillingProfile(v bool)`

SetCreateBillingProfile sets CreateBillingProfile field to given value.

### HasCreateBillingProfile

`func (o *MarketplaceInstallRequestPayload) HasCreateBillingProfile() bool`

HasCreateBillingProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


