# CaseInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CaseId** | **string** |  | 
**InvoiceId** | **string** |  | 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**PaidMicros** | **int32** |  | 
**DisputedMicros** | **int32** |  | 
**CreditedMicros** | **int32** |  | 
**WrittenOffMicros** | **int32** |  | 
**CollectibleMicros** | Pointer to **int32** |  | [optional] 
**CurrencyCode** | **string** |  | 
**BaseCurrencyCode** | **string** |  | 
**FxRateSnapshot** | Pointer to **string** |  | [optional] 
**AmountBaseMicros** | **int32** |  | 
**PaidBaseMicros** | **int32** |  | 
**DisputedBaseMicros** | **int32** |  | 
**CreditedBaseMicros** | **int32** |  | 
**WrittenOffBaseMicros** | **int32** |  | 
**CollectibleBaseMicros** | Pointer to **int32** |  | [optional] 
**DueDate** | **string** |  | 
**IssueDate** | Pointer to **string** |  | [optional] 
**PredictedPayDate** | Pointer to **string** |  | [optional] 
**InvoiceStatus** | Pointer to **string** |  | [optional] [default to "payment_pending"]
**AgingBucket** | **string** |  | 
**InvoiceType** | **string** |  | 
**VoidedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCaseInvoiceResponse

`func NewCaseInvoiceResponse(id string, caseId string, invoiceId string, amountMicros int32, paidMicros int32, disputedMicros int32, creditedMicros int32, writtenOffMicros int32, currencyCode string, baseCurrencyCode string, amountBaseMicros int32, paidBaseMicros int32, disputedBaseMicros int32, creditedBaseMicros int32, writtenOffBaseMicros int32, dueDate string, agingBucket string, invoiceType string, createdAt time.Time, updatedAt time.Time, ) *CaseInvoiceResponse`

NewCaseInvoiceResponse instantiates a new CaseInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseInvoiceResponseWithDefaults

`func NewCaseInvoiceResponseWithDefaults() *CaseInvoiceResponse`

NewCaseInvoiceResponseWithDefaults instantiates a new CaseInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CaseInvoiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaseInvoiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaseInvoiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCaseId

`func (o *CaseInvoiceResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CaseInvoiceResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CaseInvoiceResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetInvoiceId

`func (o *CaseInvoiceResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CaseInvoiceResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CaseInvoiceResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceNumber

`func (o *CaseInvoiceResponse) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *CaseInvoiceResponse) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *CaseInvoiceResponse) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *CaseInvoiceResponse) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetAmountMicros

`func (o *CaseInvoiceResponse) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *CaseInvoiceResponse) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *CaseInvoiceResponse) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetPaidMicros

`func (o *CaseInvoiceResponse) GetPaidMicros() int32`

GetPaidMicros returns the PaidMicros field if non-nil, zero value otherwise.

### GetPaidMicrosOk

`func (o *CaseInvoiceResponse) GetPaidMicrosOk() (*int32, bool)`

GetPaidMicrosOk returns a tuple with the PaidMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidMicros

`func (o *CaseInvoiceResponse) SetPaidMicros(v int32)`

SetPaidMicros sets PaidMicros field to given value.


### GetDisputedMicros

`func (o *CaseInvoiceResponse) GetDisputedMicros() int32`

GetDisputedMicros returns the DisputedMicros field if non-nil, zero value otherwise.

### GetDisputedMicrosOk

`func (o *CaseInvoiceResponse) GetDisputedMicrosOk() (*int32, bool)`

GetDisputedMicrosOk returns a tuple with the DisputedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedMicros

`func (o *CaseInvoiceResponse) SetDisputedMicros(v int32)`

SetDisputedMicros sets DisputedMicros field to given value.


### GetCreditedMicros

`func (o *CaseInvoiceResponse) GetCreditedMicros() int32`

GetCreditedMicros returns the CreditedMicros field if non-nil, zero value otherwise.

### GetCreditedMicrosOk

`func (o *CaseInvoiceResponse) GetCreditedMicrosOk() (*int32, bool)`

GetCreditedMicrosOk returns a tuple with the CreditedMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedMicros

`func (o *CaseInvoiceResponse) SetCreditedMicros(v int32)`

SetCreditedMicros sets CreditedMicros field to given value.


### GetWrittenOffMicros

`func (o *CaseInvoiceResponse) GetWrittenOffMicros() int32`

GetWrittenOffMicros returns the WrittenOffMicros field if non-nil, zero value otherwise.

### GetWrittenOffMicrosOk

`func (o *CaseInvoiceResponse) GetWrittenOffMicrosOk() (*int32, bool)`

GetWrittenOffMicrosOk returns a tuple with the WrittenOffMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenOffMicros

`func (o *CaseInvoiceResponse) SetWrittenOffMicros(v int32)`

SetWrittenOffMicros sets WrittenOffMicros field to given value.


### GetCollectibleMicros

`func (o *CaseInvoiceResponse) GetCollectibleMicros() int32`

GetCollectibleMicros returns the CollectibleMicros field if non-nil, zero value otherwise.

### GetCollectibleMicrosOk

`func (o *CaseInvoiceResponse) GetCollectibleMicrosOk() (*int32, bool)`

GetCollectibleMicrosOk returns a tuple with the CollectibleMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibleMicros

`func (o *CaseInvoiceResponse) SetCollectibleMicros(v int32)`

SetCollectibleMicros sets CollectibleMicros field to given value.

### HasCollectibleMicros

`func (o *CaseInvoiceResponse) HasCollectibleMicros() bool`

HasCollectibleMicros returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CaseInvoiceResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CaseInvoiceResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CaseInvoiceResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetBaseCurrencyCode

`func (o *CaseInvoiceResponse) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *CaseInvoiceResponse) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *CaseInvoiceResponse) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.


### GetFxRateSnapshot

`func (o *CaseInvoiceResponse) GetFxRateSnapshot() string`

GetFxRateSnapshot returns the FxRateSnapshot field if non-nil, zero value otherwise.

### GetFxRateSnapshotOk

`func (o *CaseInvoiceResponse) GetFxRateSnapshotOk() (*string, bool)`

GetFxRateSnapshotOk returns a tuple with the FxRateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRateSnapshot

`func (o *CaseInvoiceResponse) SetFxRateSnapshot(v string)`

SetFxRateSnapshot sets FxRateSnapshot field to given value.

### HasFxRateSnapshot

`func (o *CaseInvoiceResponse) HasFxRateSnapshot() bool`

HasFxRateSnapshot returns a boolean if a field has been set.

### GetAmountBaseMicros

`func (o *CaseInvoiceResponse) GetAmountBaseMicros() int32`

GetAmountBaseMicros returns the AmountBaseMicros field if non-nil, zero value otherwise.

### GetAmountBaseMicrosOk

`func (o *CaseInvoiceResponse) GetAmountBaseMicrosOk() (*int32, bool)`

GetAmountBaseMicrosOk returns a tuple with the AmountBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBaseMicros

`func (o *CaseInvoiceResponse) SetAmountBaseMicros(v int32)`

SetAmountBaseMicros sets AmountBaseMicros field to given value.


### GetPaidBaseMicros

`func (o *CaseInvoiceResponse) GetPaidBaseMicros() int32`

GetPaidBaseMicros returns the PaidBaseMicros field if non-nil, zero value otherwise.

### GetPaidBaseMicrosOk

`func (o *CaseInvoiceResponse) GetPaidBaseMicrosOk() (*int32, bool)`

GetPaidBaseMicrosOk returns a tuple with the PaidBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidBaseMicros

`func (o *CaseInvoiceResponse) SetPaidBaseMicros(v int32)`

SetPaidBaseMicros sets PaidBaseMicros field to given value.


### GetDisputedBaseMicros

`func (o *CaseInvoiceResponse) GetDisputedBaseMicros() int32`

GetDisputedBaseMicros returns the DisputedBaseMicros field if non-nil, zero value otherwise.

### GetDisputedBaseMicrosOk

`func (o *CaseInvoiceResponse) GetDisputedBaseMicrosOk() (*int32, bool)`

GetDisputedBaseMicrosOk returns a tuple with the DisputedBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedBaseMicros

`func (o *CaseInvoiceResponse) SetDisputedBaseMicros(v int32)`

SetDisputedBaseMicros sets DisputedBaseMicros field to given value.


### GetCreditedBaseMicros

`func (o *CaseInvoiceResponse) GetCreditedBaseMicros() int32`

GetCreditedBaseMicros returns the CreditedBaseMicros field if non-nil, zero value otherwise.

### GetCreditedBaseMicrosOk

`func (o *CaseInvoiceResponse) GetCreditedBaseMicrosOk() (*int32, bool)`

GetCreditedBaseMicrosOk returns a tuple with the CreditedBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedBaseMicros

`func (o *CaseInvoiceResponse) SetCreditedBaseMicros(v int32)`

SetCreditedBaseMicros sets CreditedBaseMicros field to given value.


### GetWrittenOffBaseMicros

`func (o *CaseInvoiceResponse) GetWrittenOffBaseMicros() int32`

GetWrittenOffBaseMicros returns the WrittenOffBaseMicros field if non-nil, zero value otherwise.

### GetWrittenOffBaseMicrosOk

`func (o *CaseInvoiceResponse) GetWrittenOffBaseMicrosOk() (*int32, bool)`

GetWrittenOffBaseMicrosOk returns a tuple with the WrittenOffBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenOffBaseMicros

`func (o *CaseInvoiceResponse) SetWrittenOffBaseMicros(v int32)`

SetWrittenOffBaseMicros sets WrittenOffBaseMicros field to given value.


### GetCollectibleBaseMicros

`func (o *CaseInvoiceResponse) GetCollectibleBaseMicros() int32`

GetCollectibleBaseMicros returns the CollectibleBaseMicros field if non-nil, zero value otherwise.

### GetCollectibleBaseMicrosOk

`func (o *CaseInvoiceResponse) GetCollectibleBaseMicrosOk() (*int32, bool)`

GetCollectibleBaseMicrosOk returns a tuple with the CollectibleBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectibleBaseMicros

`func (o *CaseInvoiceResponse) SetCollectibleBaseMicros(v int32)`

SetCollectibleBaseMicros sets CollectibleBaseMicros field to given value.

### HasCollectibleBaseMicros

`func (o *CaseInvoiceResponse) HasCollectibleBaseMicros() bool`

HasCollectibleBaseMicros returns a boolean if a field has been set.

### GetDueDate

`func (o *CaseInvoiceResponse) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CaseInvoiceResponse) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CaseInvoiceResponse) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetIssueDate

`func (o *CaseInvoiceResponse) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *CaseInvoiceResponse) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *CaseInvoiceResponse) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *CaseInvoiceResponse) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPredictedPayDate

`func (o *CaseInvoiceResponse) GetPredictedPayDate() string`

GetPredictedPayDate returns the PredictedPayDate field if non-nil, zero value otherwise.

### GetPredictedPayDateOk

`func (o *CaseInvoiceResponse) GetPredictedPayDateOk() (*string, bool)`

GetPredictedPayDateOk returns a tuple with the PredictedPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedPayDate

`func (o *CaseInvoiceResponse) SetPredictedPayDate(v string)`

SetPredictedPayDate sets PredictedPayDate field to given value.

### HasPredictedPayDate

`func (o *CaseInvoiceResponse) HasPredictedPayDate() bool`

HasPredictedPayDate returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *CaseInvoiceResponse) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *CaseInvoiceResponse) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *CaseInvoiceResponse) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *CaseInvoiceResponse) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetAgingBucket

`func (o *CaseInvoiceResponse) GetAgingBucket() string`

GetAgingBucket returns the AgingBucket field if non-nil, zero value otherwise.

### GetAgingBucketOk

`func (o *CaseInvoiceResponse) GetAgingBucketOk() (*string, bool)`

GetAgingBucketOk returns a tuple with the AgingBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingBucket

`func (o *CaseInvoiceResponse) SetAgingBucket(v string)`

SetAgingBucket sets AgingBucket field to given value.


### GetInvoiceType

`func (o *CaseInvoiceResponse) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *CaseInvoiceResponse) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *CaseInvoiceResponse) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.


### GetVoidedAt

`func (o *CaseInvoiceResponse) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *CaseInvoiceResponse) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *CaseInvoiceResponse) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *CaseInvoiceResponse) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaseInvoiceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaseInvoiceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaseInvoiceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CaseInvoiceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CaseInvoiceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CaseInvoiceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


