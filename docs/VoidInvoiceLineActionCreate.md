# VoidInvoiceLineActionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The action to take on the line item. | 
**NextInvoiceAt** | Pointer to **time.Time** | The time at which the line item should be invoiced again.  If not provided, the line item will be re-invoiced now. | [optional] 

## Methods

### NewVoidInvoiceLineActionCreate

`func NewVoidInvoiceLineActionCreate(type_ string, ) *VoidInvoiceLineActionCreate`

NewVoidInvoiceLineActionCreate instantiates a new VoidInvoiceLineActionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidInvoiceLineActionCreateWithDefaults

`func NewVoidInvoiceLineActionCreateWithDefaults() *VoidInvoiceLineActionCreate`

NewVoidInvoiceLineActionCreateWithDefaults instantiates a new VoidInvoiceLineActionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VoidInvoiceLineActionCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoidInvoiceLineActionCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoidInvoiceLineActionCreate) SetType(v string)`

SetType sets Type field to given value.


### GetNextInvoiceAt

`func (o *VoidInvoiceLineActionCreate) GetNextInvoiceAt() time.Time`

GetNextInvoiceAt returns the NextInvoiceAt field if non-nil, zero value otherwise.

### GetNextInvoiceAtOk

`func (o *VoidInvoiceLineActionCreate) GetNextInvoiceAtOk() (*time.Time, bool)`

GetNextInvoiceAtOk returns a tuple with the NextInvoiceAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceAt

`func (o *VoidInvoiceLineActionCreate) SetNextInvoiceAt(v time.Time)`

SetNextInvoiceAt sets NextInvoiceAt field to given value.

### HasNextInvoiceAt

`func (o *VoidInvoiceLineActionCreate) HasNextInvoiceAt() bool`

HasNextInvoiceAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


