# InvoiceAvailableActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advance** | Pointer to [**InvoiceAvailableActionDetails**](InvoiceAvailableActionDetails.md) | Advance the invoice to the next status. | [optional] [readonly] 
**Approve** | Pointer to [**InvoiceAvailableActionDetails**](InvoiceAvailableActionDetails.md) | Approve an invoice that requires manual approval. | [optional] [readonly] 
**Delete** | Pointer to [**InvoiceAvailableActionDetails**](InvoiceAvailableActionDetails.md) | Delete the invoice (only non-issued invoices can be deleted). | [optional] [readonly] 
**Retry** | Pointer to [**InvoiceAvailableActionDetails**](InvoiceAvailableActionDetails.md) | Retry an invoice issuing step that failed. | [optional] [readonly] 
**SnapshotQuantities** | Pointer to [**InvoiceAvailableActionDetails**](InvoiceAvailableActionDetails.md) | Snapshot quantities for usage based line items. | [optional] [readonly] 
**Void** | Pointer to [**InvoiceAvailableActionDetails**](InvoiceAvailableActionDetails.md) | Void an already issued invoice. | [optional] [readonly] 
**Invoice** | Pointer to **map[string]interface{}** | Invoice a gathering invoice | [optional] [readonly] 

## Methods

### NewInvoiceAvailableActions

`func NewInvoiceAvailableActions() *InvoiceAvailableActions`

NewInvoiceAvailableActions instantiates a new InvoiceAvailableActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAvailableActionsWithDefaults

`func NewInvoiceAvailableActionsWithDefaults() *InvoiceAvailableActions`

NewInvoiceAvailableActionsWithDefaults instantiates a new InvoiceAvailableActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvance

`func (o *InvoiceAvailableActions) GetAdvance() InvoiceAvailableActionDetails`

GetAdvance returns the Advance field if non-nil, zero value otherwise.

### GetAdvanceOk

`func (o *InvoiceAvailableActions) GetAdvanceOk() (*InvoiceAvailableActionDetails, bool)`

GetAdvanceOk returns a tuple with the Advance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvance

`func (o *InvoiceAvailableActions) SetAdvance(v InvoiceAvailableActionDetails)`

SetAdvance sets Advance field to given value.

### HasAdvance

`func (o *InvoiceAvailableActions) HasAdvance() bool`

HasAdvance returns a boolean if a field has been set.

### GetApprove

`func (o *InvoiceAvailableActions) GetApprove() InvoiceAvailableActionDetails`

GetApprove returns the Approve field if non-nil, zero value otherwise.

### GetApproveOk

`func (o *InvoiceAvailableActions) GetApproveOk() (*InvoiceAvailableActionDetails, bool)`

GetApproveOk returns a tuple with the Approve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprove

`func (o *InvoiceAvailableActions) SetApprove(v InvoiceAvailableActionDetails)`

SetApprove sets Approve field to given value.

### HasApprove

`func (o *InvoiceAvailableActions) HasApprove() bool`

HasApprove returns a boolean if a field has been set.

### GetDelete

`func (o *InvoiceAvailableActions) GetDelete() InvoiceAvailableActionDetails`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *InvoiceAvailableActions) GetDeleteOk() (*InvoiceAvailableActionDetails, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *InvoiceAvailableActions) SetDelete(v InvoiceAvailableActionDetails)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *InvoiceAvailableActions) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetRetry

`func (o *InvoiceAvailableActions) GetRetry() InvoiceAvailableActionDetails`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *InvoiceAvailableActions) GetRetryOk() (*InvoiceAvailableActionDetails, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *InvoiceAvailableActions) SetRetry(v InvoiceAvailableActionDetails)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *InvoiceAvailableActions) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSnapshotQuantities

`func (o *InvoiceAvailableActions) GetSnapshotQuantities() InvoiceAvailableActionDetails`

GetSnapshotQuantities returns the SnapshotQuantities field if non-nil, zero value otherwise.

### GetSnapshotQuantitiesOk

`func (o *InvoiceAvailableActions) GetSnapshotQuantitiesOk() (*InvoiceAvailableActionDetails, bool)`

GetSnapshotQuantitiesOk returns a tuple with the SnapshotQuantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotQuantities

`func (o *InvoiceAvailableActions) SetSnapshotQuantities(v InvoiceAvailableActionDetails)`

SetSnapshotQuantities sets SnapshotQuantities field to given value.

### HasSnapshotQuantities

`func (o *InvoiceAvailableActions) HasSnapshotQuantities() bool`

HasSnapshotQuantities returns a boolean if a field has been set.

### GetVoid

`func (o *InvoiceAvailableActions) GetVoid() InvoiceAvailableActionDetails`

GetVoid returns the Void field if non-nil, zero value otherwise.

### GetVoidOk

`func (o *InvoiceAvailableActions) GetVoidOk() (*InvoiceAvailableActionDetails, bool)`

GetVoidOk returns a tuple with the Void field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoid

`func (o *InvoiceAvailableActions) SetVoid(v InvoiceAvailableActionDetails)`

SetVoid sets Void field to given value.

### HasVoid

`func (o *InvoiceAvailableActions) HasVoid() bool`

HasVoid returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoiceAvailableActions) GetInvoice() map[string]interface{}`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceAvailableActions) GetInvoiceOk() (*map[string]interface{}, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceAvailableActions) SetInvoice(v map[string]interface{})`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceAvailableActions) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


