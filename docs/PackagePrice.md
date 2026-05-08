# PackagePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Amount** | **string** | The price of one package. | 
**QuantityPerPackage** | **string** | The quantity per package. | 

## Methods

### NewPackagePrice

`func NewPackagePrice(type_ string, amount string, quantityPerPackage string, ) *PackagePrice`

NewPackagePrice instantiates a new PackagePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagePriceWithDefaults

`func NewPackagePriceWithDefaults() *PackagePrice`

NewPackagePriceWithDefaults instantiates a new PackagePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PackagePrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackagePrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackagePrice) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *PackagePrice) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PackagePrice) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PackagePrice) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetQuantityPerPackage

`func (o *PackagePrice) GetQuantityPerPackage() string`

GetQuantityPerPackage returns the QuantityPerPackage field if non-nil, zero value otherwise.

### GetQuantityPerPackageOk

`func (o *PackagePrice) GetQuantityPerPackageOk() (*string, bool)`

GetQuantityPerPackageOk returns a tuple with the QuantityPerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPerPackage

`func (o *PackagePrice) SetQuantityPerPackage(v string)`

SetQuantityPerPackage sets QuantityPerPackage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


