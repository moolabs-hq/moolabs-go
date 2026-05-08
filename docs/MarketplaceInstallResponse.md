# MarketplaceInstallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**App**](App.md) |  | 
**DefaultForCapabilityTypes** | [**[]AppCapabilityType**](AppCapabilityType.md) | Default for capabilities | 

## Methods

### NewMarketplaceInstallResponse

`func NewMarketplaceInstallResponse(app App, defaultForCapabilityTypes []AppCapabilityType, ) *MarketplaceInstallResponse`

NewMarketplaceInstallResponse instantiates a new MarketplaceInstallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceInstallResponseWithDefaults

`func NewMarketplaceInstallResponseWithDefaults() *MarketplaceInstallResponse`

NewMarketplaceInstallResponseWithDefaults instantiates a new MarketplaceInstallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *MarketplaceInstallResponse) GetApp() App`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *MarketplaceInstallResponse) GetAppOk() (*App, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *MarketplaceInstallResponse) SetApp(v App)`

SetApp sets App field to given value.


### GetDefaultForCapabilityTypes

`func (o *MarketplaceInstallResponse) GetDefaultForCapabilityTypes() []AppCapabilityType`

GetDefaultForCapabilityTypes returns the DefaultForCapabilityTypes field if non-nil, zero value otherwise.

### GetDefaultForCapabilityTypesOk

`func (o *MarketplaceInstallResponse) GetDefaultForCapabilityTypesOk() (*[]AppCapabilityType, bool)`

GetDefaultForCapabilityTypesOk returns a tuple with the DefaultForCapabilityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForCapabilityTypes

`func (o *MarketplaceInstallResponse) SetDefaultForCapabilityTypes(v []AppCapabilityType)`

SetDefaultForCapabilityTypes sets DefaultForCapabilityTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


