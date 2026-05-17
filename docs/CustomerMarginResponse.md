# CustomerMarginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CustomerMarginItem**](CustomerMarginItem.md) |  | 
**ReconciliationBasis** | [**ReconciliationBasis**](ReconciliationBasis.md) |  | 

## Methods

### NewCustomerMarginResponse

`func NewCustomerMarginResponse(items []CustomerMarginItem, reconciliationBasis ReconciliationBasis, ) *CustomerMarginResponse`

NewCustomerMarginResponse instantiates a new CustomerMarginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerMarginResponseWithDefaults

`func NewCustomerMarginResponseWithDefaults() *CustomerMarginResponse`

NewCustomerMarginResponseWithDefaults instantiates a new CustomerMarginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CustomerMarginResponse) GetItems() []CustomerMarginItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomerMarginResponse) GetItemsOk() (*[]CustomerMarginItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomerMarginResponse) SetItems(v []CustomerMarginItem)`

SetItems sets Items field to given value.


### GetReconciliationBasis

`func (o *CustomerMarginResponse) GetReconciliationBasis() ReconciliationBasis`

GetReconciliationBasis returns the ReconciliationBasis field if non-nil, zero value otherwise.

### GetReconciliationBasisOk

`func (o *CustomerMarginResponse) GetReconciliationBasisOk() (*ReconciliationBasis, bool)`

GetReconciliationBasisOk returns a tuple with the ReconciliationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationBasis

`func (o *CustomerMarginResponse) SetReconciliationBasis(v ReconciliationBasis)`

SetReconciliationBasis sets ReconciliationBasis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


