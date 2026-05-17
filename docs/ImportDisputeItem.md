# ImportDisputeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **string** |  | 
**InvoiceId** | **string** |  | 
**DisputedAmountMicros** | **int32** |  | 
**DisputeType** | Pointer to **string** |  | [optional] [default to "admin"]
**Status** | Pointer to **string** |  | [optional] [default to "open"]
**Description** | Pointer to **string** |  | [optional] [default to "Imported from Growfin"]
**ExternalDisputeId** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 

## Methods

### NewImportDisputeItem

`func NewImportDisputeItem(caseId string, invoiceId string, disputedAmountMicros int32, ) *ImportDisputeItem`

NewImportDisputeItem instantiates a new ImportDisputeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportDisputeItemWithDefaults

`func NewImportDisputeItemWithDefaults() *ImportDisputeItem`

NewImportDisputeItemWithDefaults instantiates a new ImportDisputeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *ImportDisputeItem) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ImportDisputeItem) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ImportDisputeItem) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetInvoiceId

`func (o *ImportDisputeItem) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ImportDisputeItem) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ImportDisputeItem) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetDisputedAmountMicros

`func (o *ImportDisputeItem) GetDisputedAmountMicros() int32`

GetDisputedAmountMicros returns the DisputedAmountMicros field if non-nil, zero value otherwise.

### GetDisputedAmountMicrosOk

`func (o *ImportDisputeItem) GetDisputedAmountMicrosOk() (*int32, bool)`

GetDisputedAmountMicrosOk returns a tuple with the DisputedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmountMicros

`func (o *ImportDisputeItem) SetDisputedAmountMicros(v int32)`

SetDisputedAmountMicros sets DisputedAmountMicros field to given value.


### GetDisputeType

`func (o *ImportDisputeItem) GetDisputeType() string`

GetDisputeType returns the DisputeType field if non-nil, zero value otherwise.

### GetDisputeTypeOk

`func (o *ImportDisputeItem) GetDisputeTypeOk() (*string, bool)`

GetDisputeTypeOk returns a tuple with the DisputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeType

`func (o *ImportDisputeItem) SetDisputeType(v string)`

SetDisputeType sets DisputeType field to given value.

### HasDisputeType

`func (o *ImportDisputeItem) HasDisputeType() bool`

HasDisputeType returns a boolean if a field has been set.

### GetStatus

`func (o *ImportDisputeItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportDisputeItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportDisputeItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportDisputeItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *ImportDisputeItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImportDisputeItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImportDisputeItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImportDisputeItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalDisputeId

`func (o *ImportDisputeItem) GetExternalDisputeId() string`

GetExternalDisputeId returns the ExternalDisputeId field if non-nil, zero value otherwise.

### GetExternalDisputeIdOk

`func (o *ImportDisputeItem) GetExternalDisputeIdOk() (*string, bool)`

GetExternalDisputeIdOk returns a tuple with the ExternalDisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDisputeId

`func (o *ImportDisputeItem) SetExternalDisputeId(v string)`

SetExternalDisputeId sets ExternalDisputeId field to given value.

### HasExternalDisputeId

`func (o *ImportDisputeItem) HasExternalDisputeId() bool`

HasExternalDisputeId returns a boolean if a field has been set.

### GetAssignedTo

`func (o *ImportDisputeItem) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *ImportDisputeItem) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *ImportDisputeItem) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *ImportDisputeItem) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


