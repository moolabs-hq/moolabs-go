# FlatPriceWithPaymentTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the price. | 
**Amount** | **string** | The amount of the flat price. | 
**PaymentTerm** | Pointer to [**PricePaymentTerm**](PricePaymentTerm.md) | The payment term of the flat price. Defaults to in advance. | [optional] 

## Methods

### NewFlatPriceWithPaymentTerm

`func NewFlatPriceWithPaymentTerm(type_ string, amount string, ) *FlatPriceWithPaymentTerm`

NewFlatPriceWithPaymentTerm instantiates a new FlatPriceWithPaymentTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatPriceWithPaymentTermWithDefaults

`func NewFlatPriceWithPaymentTermWithDefaults() *FlatPriceWithPaymentTerm`

NewFlatPriceWithPaymentTermWithDefaults instantiates a new FlatPriceWithPaymentTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlatPriceWithPaymentTerm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlatPriceWithPaymentTerm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlatPriceWithPaymentTerm) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *FlatPriceWithPaymentTerm) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FlatPriceWithPaymentTerm) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FlatPriceWithPaymentTerm) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetPaymentTerm

`func (o *FlatPriceWithPaymentTerm) GetPaymentTerm() PricePaymentTerm`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *FlatPriceWithPaymentTerm) GetPaymentTermOk() (*PricePaymentTerm, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *FlatPriceWithPaymentTerm) SetPaymentTerm(v PricePaymentTerm)`

SetPaymentTerm sets PaymentTerm field to given value.

### HasPaymentTerm

`func (o *FlatPriceWithPaymentTerm) HasPaymentTerm() bool`

HasPaymentTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


