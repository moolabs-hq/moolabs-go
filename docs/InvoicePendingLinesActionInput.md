# InvoicePendingLinesActionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**InvoicePendingLinesActionFiltersInput**](InvoicePendingLinesActionFiltersInput.md) | Filters to apply when creating the invoice. | [optional] 
**AsOf** | Pointer to **time.Time** | The time as of which the invoice is created.  If not provided, the current time is used. | [optional] 
**CustomerId** | **string** | The customer ID for which to create the invoice. | 
**ProgressiveBillingOverride** | Pointer to **bool** | Override the progressive billing setting of the customer.  Can be used to disable/enable progressive billing in case the business logic requires it, if not provided the billing profile&#39;s progressive billing setting will be used. | [optional] 

## Methods

### NewInvoicePendingLinesActionInput

`func NewInvoicePendingLinesActionInput(customerId string, ) *InvoicePendingLinesActionInput`

NewInvoicePendingLinesActionInput instantiates a new InvoicePendingLinesActionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePendingLinesActionInputWithDefaults

`func NewInvoicePendingLinesActionInputWithDefaults() *InvoicePendingLinesActionInput`

NewInvoicePendingLinesActionInputWithDefaults instantiates a new InvoicePendingLinesActionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *InvoicePendingLinesActionInput) GetFilters() InvoicePendingLinesActionFiltersInput`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InvoicePendingLinesActionInput) GetFiltersOk() (*InvoicePendingLinesActionFiltersInput, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InvoicePendingLinesActionInput) SetFilters(v InvoicePendingLinesActionFiltersInput)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InvoicePendingLinesActionInput) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetAsOf

`func (o *InvoicePendingLinesActionInput) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *InvoicePendingLinesActionInput) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *InvoicePendingLinesActionInput) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *InvoicePendingLinesActionInput) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetCustomerId

`func (o *InvoicePendingLinesActionInput) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoicePendingLinesActionInput) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoicePendingLinesActionInput) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetProgressiveBillingOverride

`func (o *InvoicePendingLinesActionInput) GetProgressiveBillingOverride() bool`

GetProgressiveBillingOverride returns the ProgressiveBillingOverride field if non-nil, zero value otherwise.

### GetProgressiveBillingOverrideOk

`func (o *InvoicePendingLinesActionInput) GetProgressiveBillingOverrideOk() (*bool, bool)`

GetProgressiveBillingOverrideOk returns a tuple with the ProgressiveBillingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveBillingOverride

`func (o *InvoicePendingLinesActionInput) SetProgressiveBillingOverride(v bool)`

SetProgressiveBillingOverride sets ProgressiveBillingOverride field to given value.

### HasProgressiveBillingOverride

`func (o *InvoicePendingLinesActionInput) HasProgressiveBillingOverride() bool`

HasProgressiveBillingOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


