# PackagePriceWithCommitments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Amount** | **string** | The price of one package. | 
**QuantityPerPackage** | **string** | The quantity per package. | 
**MinimumAmount** | Pointer to **string** | The customer is committed to spend at least the amount. | [optional] 
**MaximumAmount** | Pointer to **string** | The customer is limited to spend at most the amount. | [optional] 

## Methods

### NewPackagePriceWithCommitments

`func NewPackagePriceWithCommitments(type_ string, amount string, quantityPerPackage string, ) *PackagePriceWithCommitments`

NewPackagePriceWithCommitments instantiates a new PackagePriceWithCommitments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagePriceWithCommitmentsWithDefaults

`func NewPackagePriceWithCommitmentsWithDefaults() *PackagePriceWithCommitments`

NewPackagePriceWithCommitmentsWithDefaults instantiates a new PackagePriceWithCommitments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PackagePriceWithCommitments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackagePriceWithCommitments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackagePriceWithCommitments) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *PackagePriceWithCommitments) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PackagePriceWithCommitments) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PackagePriceWithCommitments) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetQuantityPerPackage

`func (o *PackagePriceWithCommitments) GetQuantityPerPackage() string`

GetQuantityPerPackage returns the QuantityPerPackage field if non-nil, zero value otherwise.

### GetQuantityPerPackageOk

`func (o *PackagePriceWithCommitments) GetQuantityPerPackageOk() (*string, bool)`

GetQuantityPerPackageOk returns a tuple with the QuantityPerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPerPackage

`func (o *PackagePriceWithCommitments) SetQuantityPerPackage(v string)`

SetQuantityPerPackage sets QuantityPerPackage field to given value.


### GetMinimumAmount

`func (o *PackagePriceWithCommitments) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *PackagePriceWithCommitments) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *PackagePriceWithCommitments) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *PackagePriceWithCommitments) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *PackagePriceWithCommitments) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *PackagePriceWithCommitments) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *PackagePriceWithCommitments) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *PackagePriceWithCommitments) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


