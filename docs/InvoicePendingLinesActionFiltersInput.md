# InvoicePendingLinesActionFiltersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineIds** | Pointer to **[]string** | The pending line items to include in the invoice, if not provided: - all line items that have invoice_at &lt; asOf will be included - [progressive billing only] all usage based line items will be included up to asOf, new usage-based line items will be staged for the rest of the billing cycle  All lineIDs present in the list, must exists and must be invoicable as of asOf, or the action will fail. | [optional] 

## Methods

### NewInvoicePendingLinesActionFiltersInput

`func NewInvoicePendingLinesActionFiltersInput() *InvoicePendingLinesActionFiltersInput`

NewInvoicePendingLinesActionFiltersInput instantiates a new InvoicePendingLinesActionFiltersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePendingLinesActionFiltersInputWithDefaults

`func NewInvoicePendingLinesActionFiltersInputWithDefaults() *InvoicePendingLinesActionFiltersInput`

NewInvoicePendingLinesActionFiltersInputWithDefaults instantiates a new InvoicePendingLinesActionFiltersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineIds

`func (o *InvoicePendingLinesActionFiltersInput) GetLineIds() []string`

GetLineIds returns the LineIds field if non-nil, zero value otherwise.

### GetLineIdsOk

`func (o *InvoicePendingLinesActionFiltersInput) GetLineIdsOk() (*[]string, bool)`

GetLineIdsOk returns a tuple with the LineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIds

`func (o *InvoicePendingLinesActionFiltersInput) SetLineIds(v []string)`

SetLineIds sets LineIds field to given value.

### HasLineIds

`func (o *InvoicePendingLinesActionFiltersInput) HasLineIds() bool`

HasLineIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


