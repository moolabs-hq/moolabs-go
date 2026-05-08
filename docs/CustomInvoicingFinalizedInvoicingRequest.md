# CustomInvoicingFinalizedInvoicingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | Pointer to **string** | If set the invoice&#39;s number will be set to this value. | [optional] 
**SentToCustomerAt** | Pointer to **time.Time** | If set the invoice&#39;s sent to customer at will be set to this value. | [optional] 

## Methods

### NewCustomInvoicingFinalizedInvoicingRequest

`func NewCustomInvoicingFinalizedInvoicingRequest() *CustomInvoicingFinalizedInvoicingRequest`

NewCustomInvoicingFinalizedInvoicingRequest instantiates a new CustomInvoicingFinalizedInvoicingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingFinalizedInvoicingRequestWithDefaults

`func NewCustomInvoicingFinalizedInvoicingRequestWithDefaults() *CustomInvoicingFinalizedInvoicingRequest`

NewCustomInvoicingFinalizedInvoicingRequestWithDefaults instantiates a new CustomInvoicingFinalizedInvoicingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *CustomInvoicingFinalizedInvoicingRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CustomInvoicingFinalizedInvoicingRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CustomInvoicingFinalizedInvoicingRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *CustomInvoicingFinalizedInvoicingRequest) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetSentToCustomerAt

`func (o *CustomInvoicingFinalizedInvoicingRequest) GetSentToCustomerAt() time.Time`

GetSentToCustomerAt returns the SentToCustomerAt field if non-nil, zero value otherwise.

### GetSentToCustomerAtOk

`func (o *CustomInvoicingFinalizedInvoicingRequest) GetSentToCustomerAtOk() (*time.Time, bool)`

GetSentToCustomerAtOk returns a tuple with the SentToCustomerAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToCustomerAt

`func (o *CustomInvoicingFinalizedInvoicingRequest) SetSentToCustomerAt(v time.Time)`

SetSentToCustomerAt sets SentToCustomerAt field to given value.

### HasSentToCustomerAt

`func (o *CustomInvoicingFinalizedInvoicingRequest) HasSentToCustomerAt() bool`

HasSentToCustomerAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


