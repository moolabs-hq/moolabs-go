# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Type** | [**InvoiceType**](InvoiceType.md) | Type of the invoice.  The type of invoice determines the purpose of the invoice and how it should be handled.  Supported types: - standard: A regular commercial invoice document between a supplier and customer. - credit_note: Reflects a refund either partial or complete of the preceding document. A credit note effectively *extends* the previous document. | [readonly] 
**Supplier** | [**BillingParty**](BillingParty.md) | The taxable entity supplying the goods or services. | 
**Customer** | [**BillingInvoiceCustomerExtendedDetails**](BillingInvoiceCustomerExtendedDetails.md) | Legal entity receiving the goods or services. | 
**Number** | **string** | Number specifies the human readable key used to reference this Invoice.  The invoice number can change in the draft phases, as we are allocating temporary draft invoice numbers, but it&#39;s final as soon as the invoice gets finalized (issued state).  Please note that the number is (depending on the upstream settings) either unique for the whole organization or unique for the customer, or in multi (stripe) account setups unique for the account. | [readonly] 
**Currency** | **string** | Currency for all invoice line items.  Multi currency invoices are not supported yet. | 
**Preceding** | Pointer to [**[]CreditNoteOriginalInvoiceRef**](CreditNoteOriginalInvoiceRef.md) | Key information regarding previous invoices and potentially details as to why they were corrected. | [optional] 
**Totals** | [**InvoiceTotals**](InvoiceTotals.md) | Summary of all the invoice totals, including taxes (calculated). | [readonly] 
**Status** | [**InvoiceStatus**](InvoiceStatus.md) | The status of the invoice.  This field only conatins a simplified status, for more detailed information use the statusDetails field. | [readonly] 
**StatusDetails** | [**InvoiceStatusDetails**](InvoiceStatusDetails.md) | The details of the current invoice status. | [readonly] 
**IssuedAt** | Pointer to **time.Time** | The time the invoice was issued.  Depending on the status of the invoice this can mean multiple things: - draft, gathering: The time the invoice will be issued based on the workflow settings. - issued: The time the invoice was issued. | [optional] [readonly] 
**DraftUntil** | Pointer to **time.Time** | The time until the invoice is in draft status.  On draft invoice creation it is calculated from the workflow settings.  If manual approval is required, the draftUntil time is set. | [optional] 
**QuantitySnapshotedAt** | Pointer to **time.Time** | The time when the quantity snapshots on the invoice lines were taken. | [optional] [readonly] 
**CollectionAt** | Pointer to **time.Time** | The time when the invoice will be/has been collected. | [optional] [readonly] 
**DueAt** | Pointer to **time.Time** | Due time of the fulfillment of the invoice (if available). | [optional] [readonly] 
**Period** | Pointer to [**Period**](Period.md) | The period the invoice covers. If the invoice has no line items, it&#39;s not set. | [optional] 
**VoidedAt** | Pointer to **time.Time** | The time the invoice was voided.  If the invoice was voided, this field will be set to the time the invoice was voided. | [optional] [readonly] 
**SentToCustomerAt** | Pointer to **time.Time** | The time the invoice was sent to customer. | [optional] [readonly] 
**Workflow** | [**InvoiceWorkflowSettings**](InvoiceWorkflowSettings.md) | The workflow associated with the invoice.  It is always a snapshot of the workflow settings at the time of invoice creation. The field is optional as it should be explicitly requested with expand options. | 
**Lines** | Pointer to [**[]InvoiceLine**](InvoiceLine.md) | List of invoice lines representing each of the items sold to the customer. | [optional] 
**Payment** | Pointer to [**InvoicePaymentTerms**](InvoicePaymentTerms.md) | Information on when, how, and to whom the invoice should be paid. | [optional] [readonly] 
**ValidationIssues** | Pointer to [**[]ValidationIssue**](ValidationIssue.md) | Validation issues reported by the invoice workflow. | [optional] 
**ExternalIds** | Pointer to [**InvoiceAppExternalIds**](InvoiceAppExternalIds.md) | External IDs of the invoice in other apps such as Stripe. | [optional] [readonly] 

## Methods

### NewInvoice

`func NewInvoice(id string, createdAt time.Time, updatedAt time.Time, type_ InvoiceType, supplier BillingParty, customer BillingInvoiceCustomerExtendedDetails, number string, currency string, totals InvoiceTotals, status InvoiceStatus, statusDetails InvoiceStatusDetails, workflow InvoiceWorkflowSettings, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *Invoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Invoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Invoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Invoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Invoice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Invoice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Invoice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Invoice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Invoice) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Invoice) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Invoice) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Invoice) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Invoice) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Invoice) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Invoice) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetType

`func (o *Invoice) GetType() InvoiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*InvoiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v InvoiceType)`

SetType sets Type field to given value.


### GetSupplier

`func (o *Invoice) GetSupplier() BillingParty`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *Invoice) GetSupplierOk() (*BillingParty, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *Invoice) SetSupplier(v BillingParty)`

SetSupplier sets Supplier field to given value.


### GetCustomer

`func (o *Invoice) GetCustomer() BillingInvoiceCustomerExtendedDetails`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Invoice) GetCustomerOk() (*BillingInvoiceCustomerExtendedDetails, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Invoice) SetCustomer(v BillingInvoiceCustomerExtendedDetails)`

SetCustomer sets Customer field to given value.


### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPreceding

`func (o *Invoice) GetPreceding() []CreditNoteOriginalInvoiceRef`

GetPreceding returns the Preceding field if non-nil, zero value otherwise.

### GetPrecedingOk

`func (o *Invoice) GetPrecedingOk() (*[]CreditNoteOriginalInvoiceRef, bool)`

GetPrecedingOk returns a tuple with the Preceding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreceding

`func (o *Invoice) SetPreceding(v []CreditNoteOriginalInvoiceRef)`

SetPreceding sets Preceding field to given value.

### HasPreceding

`func (o *Invoice) HasPreceding() bool`

HasPreceding returns a boolean if a field has been set.

### GetTotals

`func (o *Invoice) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Invoice) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Invoice) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.


### GetStatus

`func (o *Invoice) GetStatus() InvoiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*InvoiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v InvoiceStatus)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *Invoice) GetStatusDetails() InvoiceStatusDetails`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Invoice) GetStatusDetailsOk() (*InvoiceStatusDetails, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Invoice) SetStatusDetails(v InvoiceStatusDetails)`

SetStatusDetails sets StatusDetails field to given value.


### GetIssuedAt

`func (o *Invoice) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Invoice) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Invoice) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Invoice) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetDraftUntil

`func (o *Invoice) GetDraftUntil() time.Time`

GetDraftUntil returns the DraftUntil field if non-nil, zero value otherwise.

### GetDraftUntilOk

`func (o *Invoice) GetDraftUntilOk() (*time.Time, bool)`

GetDraftUntilOk returns a tuple with the DraftUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftUntil

`func (o *Invoice) SetDraftUntil(v time.Time)`

SetDraftUntil sets DraftUntil field to given value.

### HasDraftUntil

`func (o *Invoice) HasDraftUntil() bool`

HasDraftUntil returns a boolean if a field has been set.

### GetQuantitySnapshotedAt

`func (o *Invoice) GetQuantitySnapshotedAt() time.Time`

GetQuantitySnapshotedAt returns the QuantitySnapshotedAt field if non-nil, zero value otherwise.

### GetQuantitySnapshotedAtOk

`func (o *Invoice) GetQuantitySnapshotedAtOk() (*time.Time, bool)`

GetQuantitySnapshotedAtOk returns a tuple with the QuantitySnapshotedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantitySnapshotedAt

`func (o *Invoice) SetQuantitySnapshotedAt(v time.Time)`

SetQuantitySnapshotedAt sets QuantitySnapshotedAt field to given value.

### HasQuantitySnapshotedAt

`func (o *Invoice) HasQuantitySnapshotedAt() bool`

HasQuantitySnapshotedAt returns a boolean if a field has been set.

### GetCollectionAt

`func (o *Invoice) GetCollectionAt() time.Time`

GetCollectionAt returns the CollectionAt field if non-nil, zero value otherwise.

### GetCollectionAtOk

`func (o *Invoice) GetCollectionAtOk() (*time.Time, bool)`

GetCollectionAtOk returns a tuple with the CollectionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAt

`func (o *Invoice) SetCollectionAt(v time.Time)`

SetCollectionAt sets CollectionAt field to given value.

### HasCollectionAt

`func (o *Invoice) HasCollectionAt() bool`

HasCollectionAt returns a boolean if a field has been set.

### GetDueAt

`func (o *Invoice) GetDueAt() time.Time`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *Invoice) GetDueAtOk() (*time.Time, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *Invoice) SetDueAt(v time.Time)`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *Invoice) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.

### GetPeriod

`func (o *Invoice) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Invoice) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Invoice) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Invoice) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetVoidedAt

`func (o *Invoice) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *Invoice) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *Invoice) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *Invoice) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### GetSentToCustomerAt

`func (o *Invoice) GetSentToCustomerAt() time.Time`

GetSentToCustomerAt returns the SentToCustomerAt field if non-nil, zero value otherwise.

### GetSentToCustomerAtOk

`func (o *Invoice) GetSentToCustomerAtOk() (*time.Time, bool)`

GetSentToCustomerAtOk returns a tuple with the SentToCustomerAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToCustomerAt

`func (o *Invoice) SetSentToCustomerAt(v time.Time)`

SetSentToCustomerAt sets SentToCustomerAt field to given value.

### HasSentToCustomerAt

`func (o *Invoice) HasSentToCustomerAt() bool`

HasSentToCustomerAt returns a boolean if a field has been set.

### GetWorkflow

`func (o *Invoice) GetWorkflow() InvoiceWorkflowSettings`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *Invoice) GetWorkflowOk() (*InvoiceWorkflowSettings, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *Invoice) SetWorkflow(v InvoiceWorkflowSettings)`

SetWorkflow sets Workflow field to given value.


### GetLines

`func (o *Invoice) GetLines() []InvoiceLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Invoice) GetLinesOk() (*[]InvoiceLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Invoice) SetLines(v []InvoiceLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Invoice) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetPayment

`func (o *Invoice) GetPayment() InvoicePaymentTerms`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *Invoice) GetPaymentOk() (*InvoicePaymentTerms, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *Invoice) SetPayment(v InvoicePaymentTerms)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *Invoice) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetValidationIssues

`func (o *Invoice) GetValidationIssues() []ValidationIssue`

GetValidationIssues returns the ValidationIssues field if non-nil, zero value otherwise.

### GetValidationIssuesOk

`func (o *Invoice) GetValidationIssuesOk() (*[]ValidationIssue, bool)`

GetValidationIssuesOk returns a tuple with the ValidationIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationIssues

`func (o *Invoice) SetValidationIssues(v []ValidationIssue)`

SetValidationIssues sets ValidationIssues field to given value.

### HasValidationIssues

`func (o *Invoice) HasValidationIssues() bool`

HasValidationIssues returns a boolean if a field has been set.

### GetExternalIds

`func (o *Invoice) GetExternalIds() InvoiceAppExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *Invoice) GetExternalIdsOk() (*InvoiceAppExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *Invoice) SetExternalIds(v InvoiceAppExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *Invoice) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


