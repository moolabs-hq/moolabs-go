# CustomInvoicingDraftSynchronizedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoicing** | Pointer to [**CustomInvoicingSyncResult**](CustomInvoicingSyncResult.md) | The result of the synchronization. | [optional] 

## Methods

### NewCustomInvoicingDraftSynchronizedRequest

`func NewCustomInvoicingDraftSynchronizedRequest() *CustomInvoicingDraftSynchronizedRequest`

NewCustomInvoicingDraftSynchronizedRequest instantiates a new CustomInvoicingDraftSynchronizedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingDraftSynchronizedRequestWithDefaults

`func NewCustomInvoicingDraftSynchronizedRequestWithDefaults() *CustomInvoicingDraftSynchronizedRequest`

NewCustomInvoicingDraftSynchronizedRequestWithDefaults instantiates a new CustomInvoicingDraftSynchronizedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoicing

`func (o *CustomInvoicingDraftSynchronizedRequest) GetInvoicing() CustomInvoicingSyncResult`

GetInvoicing returns the Invoicing field if non-nil, zero value otherwise.

### GetInvoicingOk

`func (o *CustomInvoicingDraftSynchronizedRequest) GetInvoicingOk() (*CustomInvoicingSyncResult, bool)`

GetInvoicingOk returns a tuple with the Invoicing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicing

`func (o *CustomInvoicingDraftSynchronizedRequest) SetInvoicing(v CustomInvoicingSyncResult)`

SetInvoicing sets Invoicing field to given value.

### HasInvoicing

`func (o *CustomInvoicingDraftSynchronizedRequest) HasInvoicing() bool`

HasInvoicing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


