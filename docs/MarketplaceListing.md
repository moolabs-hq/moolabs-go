# MarketplaceListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AppType**](AppType.md) | The app&#39;s type | 
**Name** | **string** | The app&#39;s name. | 
**Description** | **string** | The app&#39;s description. | 
**Capabilities** | [**[]AppCapability**](AppCapability.md) | The app&#39;s capabilities. | 
**InstallMethods** | [**[]InstallMethod**](InstallMethod.md) | Install methods.  List of methods to install the app. | 

## Methods

### NewMarketplaceListing

`func NewMarketplaceListing(type_ AppType, name string, description string, capabilities []AppCapability, installMethods []InstallMethod, ) *MarketplaceListing`

NewMarketplaceListing instantiates a new MarketplaceListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceListingWithDefaults

`func NewMarketplaceListingWithDefaults() *MarketplaceListing`

NewMarketplaceListingWithDefaults instantiates a new MarketplaceListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketplaceListing) GetType() AppType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceListing) GetTypeOk() (*AppType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceListing) SetType(v AppType)`

SetType sets Type field to given value.


### GetName

`func (o *MarketplaceListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceListing) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MarketplaceListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceListing) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCapabilities

`func (o *MarketplaceListing) GetCapabilities() []AppCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MarketplaceListing) GetCapabilitiesOk() (*[]AppCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MarketplaceListing) SetCapabilities(v []AppCapability)`

SetCapabilities sets Capabilities field to given value.


### GetInstallMethods

`func (o *MarketplaceListing) GetInstallMethods() []InstallMethod`

GetInstallMethods returns the InstallMethods field if non-nil, zero value otherwise.

### GetInstallMethodsOk

`func (o *MarketplaceListing) GetInstallMethodsOk() (*[]InstallMethod, bool)`

GetInstallMethodsOk returns a tuple with the InstallMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallMethods

`func (o *MarketplaceListing) SetInstallMethods(v []InstallMethod)`

SetInstallMethods sets InstallMethods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


