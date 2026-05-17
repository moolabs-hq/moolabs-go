# DisputeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CaseId** | **string** |  | 
**InvoiceId** | **string** |  | 
**DisputeType** | Pointer to **string** |  | [optional] 
**ClassificationStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AccountingClass** | Pointer to **string** |  | [optional] 
**AccountingReasonCode** | Pointer to **string** |  | [optional] 
**DisputedAmountMicros** | **int32** |  | 
**ResolvedAmountMicros** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  | 
**ResolutionNotes** | Pointer to **string** |  | [optional] 
**ExternalDisputeId** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**ResolutionDeadline** | Pointer to **time.Time** |  | [optional] 
**ResolvedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomerName** | Pointer to **string** |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**AgingBucket** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**InvoiceDueDate** | Pointer to **string** |  | [optional] 

## Methods

### NewDisputeResponse

`func NewDisputeResponse(id string, tenantId string, caseId string, invoiceId string, disputedAmountMicros int32, description string, ) *DisputeResponse`

NewDisputeResponse instantiates a new DisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeResponseWithDefaults

`func NewDisputeResponseWithDefaults() *DisputeResponse`

NewDisputeResponseWithDefaults instantiates a new DisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisputeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisputeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisputeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *DisputeResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DisputeResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DisputeResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCaseId

`func (o *DisputeResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *DisputeResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *DisputeResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetInvoiceId

`func (o *DisputeResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DisputeResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DisputeResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetDisputeType

`func (o *DisputeResponse) GetDisputeType() string`

GetDisputeType returns the DisputeType field if non-nil, zero value otherwise.

### GetDisputeTypeOk

`func (o *DisputeResponse) GetDisputeTypeOk() (*string, bool)`

GetDisputeTypeOk returns a tuple with the DisputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeType

`func (o *DisputeResponse) SetDisputeType(v string)`

SetDisputeType sets DisputeType field to given value.

### HasDisputeType

`func (o *DisputeResponse) HasDisputeType() bool`

HasDisputeType returns a boolean if a field has been set.

### GetClassificationStatus

`func (o *DisputeResponse) GetClassificationStatus() string`

GetClassificationStatus returns the ClassificationStatus field if non-nil, zero value otherwise.

### GetClassificationStatusOk

`func (o *DisputeResponse) GetClassificationStatusOk() (*string, bool)`

GetClassificationStatusOk returns a tuple with the ClassificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationStatus

`func (o *DisputeResponse) SetClassificationStatus(v string)`

SetClassificationStatus sets ClassificationStatus field to given value.

### HasClassificationStatus

`func (o *DisputeResponse) HasClassificationStatus() bool`

HasClassificationStatus returns a boolean if a field has been set.

### GetStatus

`func (o *DisputeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DisputeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccountingClass

`func (o *DisputeResponse) GetAccountingClass() string`

GetAccountingClass returns the AccountingClass field if non-nil, zero value otherwise.

### GetAccountingClassOk

`func (o *DisputeResponse) GetAccountingClassOk() (*string, bool)`

GetAccountingClassOk returns a tuple with the AccountingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingClass

`func (o *DisputeResponse) SetAccountingClass(v string)`

SetAccountingClass sets AccountingClass field to given value.

### HasAccountingClass

`func (o *DisputeResponse) HasAccountingClass() bool`

HasAccountingClass returns a boolean if a field has been set.

### GetAccountingReasonCode

`func (o *DisputeResponse) GetAccountingReasonCode() string`

GetAccountingReasonCode returns the AccountingReasonCode field if non-nil, zero value otherwise.

### GetAccountingReasonCodeOk

`func (o *DisputeResponse) GetAccountingReasonCodeOk() (*string, bool)`

GetAccountingReasonCodeOk returns a tuple with the AccountingReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingReasonCode

`func (o *DisputeResponse) SetAccountingReasonCode(v string)`

SetAccountingReasonCode sets AccountingReasonCode field to given value.

### HasAccountingReasonCode

`func (o *DisputeResponse) HasAccountingReasonCode() bool`

HasAccountingReasonCode returns a boolean if a field has been set.

### GetDisputedAmountMicros

`func (o *DisputeResponse) GetDisputedAmountMicros() int32`

GetDisputedAmountMicros returns the DisputedAmountMicros field if non-nil, zero value otherwise.

### GetDisputedAmountMicrosOk

`func (o *DisputeResponse) GetDisputedAmountMicrosOk() (*int32, bool)`

GetDisputedAmountMicrosOk returns a tuple with the DisputedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmountMicros

`func (o *DisputeResponse) SetDisputedAmountMicros(v int32)`

SetDisputedAmountMicros sets DisputedAmountMicros field to given value.


### GetResolvedAmountMicros

`func (o *DisputeResponse) GetResolvedAmountMicros() int32`

GetResolvedAmountMicros returns the ResolvedAmountMicros field if non-nil, zero value otherwise.

### GetResolvedAmountMicrosOk

`func (o *DisputeResponse) GetResolvedAmountMicrosOk() (*int32, bool)`

GetResolvedAmountMicrosOk returns a tuple with the ResolvedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAmountMicros

`func (o *DisputeResponse) SetResolvedAmountMicros(v int32)`

SetResolvedAmountMicros sets ResolvedAmountMicros field to given value.

### HasResolvedAmountMicros

`func (o *DisputeResponse) HasResolvedAmountMicros() bool`

HasResolvedAmountMicros returns a boolean if a field has been set.

### GetDescription

`func (o *DisputeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DisputeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DisputeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetResolutionNotes

`func (o *DisputeResponse) GetResolutionNotes() string`

GetResolutionNotes returns the ResolutionNotes field if non-nil, zero value otherwise.

### GetResolutionNotesOk

`func (o *DisputeResponse) GetResolutionNotesOk() (*string, bool)`

GetResolutionNotesOk returns a tuple with the ResolutionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionNotes

`func (o *DisputeResponse) SetResolutionNotes(v string)`

SetResolutionNotes sets ResolutionNotes field to given value.

### HasResolutionNotes

`func (o *DisputeResponse) HasResolutionNotes() bool`

HasResolutionNotes returns a boolean if a field has been set.

### GetExternalDisputeId

`func (o *DisputeResponse) GetExternalDisputeId() string`

GetExternalDisputeId returns the ExternalDisputeId field if non-nil, zero value otherwise.

### GetExternalDisputeIdOk

`func (o *DisputeResponse) GetExternalDisputeIdOk() (*string, bool)`

GetExternalDisputeIdOk returns a tuple with the ExternalDisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDisputeId

`func (o *DisputeResponse) SetExternalDisputeId(v string)`

SetExternalDisputeId sets ExternalDisputeId field to given value.

### HasExternalDisputeId

`func (o *DisputeResponse) HasExternalDisputeId() bool`

HasExternalDisputeId returns a boolean if a field has been set.

### GetAssignedTo

`func (o *DisputeResponse) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *DisputeResponse) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *DisputeResponse) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *DisputeResponse) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetResolutionDeadline

`func (o *DisputeResponse) GetResolutionDeadline() time.Time`

GetResolutionDeadline returns the ResolutionDeadline field if non-nil, zero value otherwise.

### GetResolutionDeadlineOk

`func (o *DisputeResponse) GetResolutionDeadlineOk() (*time.Time, bool)`

GetResolutionDeadlineOk returns a tuple with the ResolutionDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDeadline

`func (o *DisputeResponse) SetResolutionDeadline(v time.Time)`

SetResolutionDeadline sets ResolutionDeadline field to given value.

### HasResolutionDeadline

`func (o *DisputeResponse) HasResolutionDeadline() bool`

HasResolutionDeadline returns a boolean if a field has been set.

### GetResolvedAt

`func (o *DisputeResponse) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *DisputeResponse) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *DisputeResponse) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *DisputeResponse) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DisputeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DisputeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DisputeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DisputeResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DisputeResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DisputeResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DisputeResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DisputeResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCustomerName

`func (o *DisputeResponse) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *DisputeResponse) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *DisputeResponse) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *DisputeResponse) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *DisputeResponse) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *DisputeResponse) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *DisputeResponse) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *DisputeResponse) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetAgingBucket

`func (o *DisputeResponse) GetAgingBucket() string`

GetAgingBucket returns the AgingBucket field if non-nil, zero value otherwise.

### GetAgingBucketOk

`func (o *DisputeResponse) GetAgingBucketOk() (*string, bool)`

GetAgingBucketOk returns a tuple with the AgingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingBucket

`func (o *DisputeResponse) SetAgingBucket(v string)`

SetAgingBucket sets AgingBucket field to given value.

### HasAgingBucket

`func (o *DisputeResponse) HasAgingBucket() bool`

HasAgingBucket returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *DisputeResponse) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *DisputeResponse) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *DisputeResponse) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *DisputeResponse) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceDueDate

`func (o *DisputeResponse) GetInvoiceDueDate() string`

GetInvoiceDueDate returns the InvoiceDueDate field if non-nil, zero value otherwise.

### GetInvoiceDueDateOk

`func (o *DisputeResponse) GetInvoiceDueDateOk() (*string, bool)`

GetInvoiceDueDateOk returns a tuple with the InvoiceDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDueDate

`func (o *DisputeResponse) SetInvoiceDueDate(v string)`

SetInvoiceDueDate sets InvoiceDueDate field to given value.

### HasInvoiceDueDate

`func (o *DisputeResponse) HasInvoiceDueDate() bool`

HasInvoiceDueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


