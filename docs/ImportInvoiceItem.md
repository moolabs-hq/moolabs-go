# ImportInvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **string** |  | 
**InvoiceId** | **string** |  | 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**PaidMicros** | Pointer to **int32** |  | [optional] [default to 0]
**DisputedMicros** | Pointer to **int32** |  | [optional] [default to 0]
**CreditedMicros** | Pointer to **int32** |  | [optional] [default to 0]
**WrittenOffMicros** | Pointer to **int32** |  | [optional] [default to 0]
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**DueDate** | **string** |  | 
**IssueDate** | Pointer to **string** |  | [optional] 
**AgingBucket** | Pointer to **string** |  | [optional] [default to "current"]
**InvoiceType** | Pointer to **string** |  | [optional] [default to "usage"]

## Methods

### NewImportInvoiceItem

`func NewImportInvoiceItem(caseId string, invoiceId string, amountMicros int32, dueDate string, ) *ImportInvoiceItem`

NewImportInvoiceItem instantiates a new ImportInvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportInvoiceItemWithDefaults

`func NewImportInvoiceItemWithDefaults() *ImportInvoiceItem`

NewImportInvoiceItemWithDefaults instantiates a new ImportInvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *ImportInvoiceItem) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ImportInvoiceItem) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ImportInvoiceItem) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetInvoiceId

`func (o *ImportInvoiceItem) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ImportInvoiceItem) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ImportInvoiceItem) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceNumber

`func (o *ImportInvoiceItem) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *ImportInvoiceItem) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *ImportInvoiceItem) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *ImportInvoiceItem) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetAmountMicros

`func (o *ImportInvoiceItem) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *ImportInvoiceItem) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *ImportInvoiceItem) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetPaidMicros

`func (o *ImportInvoiceItem) GetPaidMicros() int32`

GetPaidMicros returns the PaidMicros field if non-nil, zero value otherwise.

### GetPaidMicrosOk

`func (o *ImportInvoiceItem) GetPaidMicrosOk() (*int32, bool)`

GetPaidMicrosOk returns a tuple with the PaidMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMicros

`func (o *ImportInvoiceItem) SetPaidMicros(v int32)`

SetPaidMicros sets PaidMicros field to given value.

### HasPaidMicros

`func (o *ImportInvoiceItem) HasPaidMicros() bool`

HasPaidMicros returns a boolean if a field has been set.

### GetDisputedMicros

`func (o *ImportInvoiceItem) GetDisputedMicros() int32`

GetDisputedMicros returns the DisputedMicros field if non-nil, zero value otherwise.

### GetDisputedMicrosOk

`func (o *ImportInvoiceItem) GetDisputedMicrosOk() (*int32, bool)`

GetDisputedMicrosOk returns a tuple with the DisputedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedMicros

`func (o *ImportInvoiceItem) SetDisputedMicros(v int32)`

SetDisputedMicros sets DisputedMicros field to given value.

### HasDisputedMicros

`func (o *ImportInvoiceItem) HasDisputedMicros() bool`

HasDisputedMicros returns a boolean if a field has been set.

### GetCreditedMicros

`func (o *ImportInvoiceItem) GetCreditedMicros() int32`

GetCreditedMicros returns the CreditedMicros field if non-nil, zero value otherwise.

### GetCreditedMicrosOk

`func (o *ImportInvoiceItem) GetCreditedMicrosOk() (*int32, bool)`

GetCreditedMicrosOk returns a tuple with the CreditedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedMicros

`func (o *ImportInvoiceItem) SetCreditedMicros(v int32)`

SetCreditedMicros sets CreditedMicros field to given value.

### HasCreditedMicros

`func (o *ImportInvoiceItem) HasCreditedMicros() bool`

HasCreditedMicros returns a boolean if a field has been set.

### GetWrittenOffMicros

`func (o *ImportInvoiceItem) GetWrittenOffMicros() int32`

GetWrittenOffMicros returns the WrittenOffMicros field if non-nil, zero value otherwise.

### GetWrittenOffMicrosOk

`func (o *ImportInvoiceItem) GetWrittenOffMicrosOk() (*int32, bool)`

GetWrittenOffMicrosOk returns a tuple with the WrittenOffMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenOffMicros

`func (o *ImportInvoiceItem) SetWrittenOffMicros(v int32)`

SetWrittenOffMicros sets WrittenOffMicros field to given value.

### HasWrittenOffMicros

`func (o *ImportInvoiceItem) HasWrittenOffMicros() bool`

HasWrittenOffMicros returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ImportInvoiceItem) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ImportInvoiceItem) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ImportInvoiceItem) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ImportInvoiceItem) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDueDate

`func (o *ImportInvoiceItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ImportInvoiceItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ImportInvoiceItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetIssueDate

`func (o *ImportInvoiceItem) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ImportInvoiceItem) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ImportInvoiceItem) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *ImportInvoiceItem) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetAgingBucket

`func (o *ImportInvoiceItem) GetAgingBucket() string`

GetAgingBucket returns the AgingBucket field if non-nil, zero value otherwise.

### GetAgingBucketOk

`func (o *ImportInvoiceItem) GetAgingBucketOk() (*string, bool)`

GetAgingBucketOk returns a tuple with the AgingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingBucket

`func (o *ImportInvoiceItem) SetAgingBucket(v string)`

SetAgingBucket sets AgingBucket field to given value.

### HasAgingBucket

`func (o *ImportInvoiceItem) HasAgingBucket() bool`

HasAgingBucket returns a boolean if a field has been set.

### GetInvoiceType

`func (o *ImportInvoiceItem) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *ImportInvoiceItem) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *ImportInvoiceItem) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *ImportInvoiceItem) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


