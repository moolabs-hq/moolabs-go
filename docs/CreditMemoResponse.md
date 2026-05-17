# CreditMemoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AccountId** | **string** |  | 
**CaseId** | Pointer to **string** |  | [optional] 
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**Status** | **string** |  | 
**CreditAmountMicros** | **int32** |  | 
**AppliedAmountMicros** | **int32** |  | 
**UnappliedAmountMicros** | **int32** |  | 
**CurrencyCode** | **string** |  | 
**Subsidiary** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ApprovedBy** | Pointer to **string** |  | [optional] 
**ApprovedAt** | Pointer to **time.Time** |  | [optional] 
**ClosedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCreditMemoResponse

`func NewCreditMemoResponse(id string, tenantId string, accountId string, reason string, status string, creditAmountMicros int32, appliedAmountMicros int32, unappliedAmountMicros int32, currencyCode string, createdAt time.Time, updatedAt time.Time, ) *CreditMemoResponse`

NewCreditMemoResponse instantiates a new CreditMemoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditMemoResponseWithDefaults

`func NewCreditMemoResponseWithDefaults() *CreditMemoResponse`

NewCreditMemoResponseWithDefaults instantiates a new CreditMemoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditMemoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditMemoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditMemoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CreditMemoResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreditMemoResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreditMemoResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAccountId

`func (o *CreditMemoResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditMemoResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditMemoResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCaseId

`func (o *CreditMemoResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CreditMemoResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CreditMemoResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *CreditMemoResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *CreditMemoResponse) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *CreditMemoResponse) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *CreditMemoResponse) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *CreditMemoResponse) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetReason

`func (o *CreditMemoResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditMemoResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditMemoResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *CreditMemoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditMemoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditMemoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreditAmountMicros

`func (o *CreditMemoResponse) GetCreditAmountMicros() int32`

GetCreditAmountMicros returns the CreditAmountMicros field if non-nil, zero value otherwise.

### GetCreditAmountMicrosOk

`func (o *CreditMemoResponse) GetCreditAmountMicrosOk() (*int32, bool)`

GetCreditAmountMicrosOk returns a tuple with the CreditAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountMicros

`func (o *CreditMemoResponse) SetCreditAmountMicros(v int32)`

SetCreditAmountMicros sets CreditAmountMicros field to given value.


### GetAppliedAmountMicros

`func (o *CreditMemoResponse) GetAppliedAmountMicros() int32`

GetAppliedAmountMicros returns the AppliedAmountMicros field if non-nil, zero value otherwise.

### GetAppliedAmountMicrosOk

`func (o *CreditMemoResponse) GetAppliedAmountMicrosOk() (*int32, bool)`

GetAppliedAmountMicrosOk returns a tuple with the AppliedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmountMicros

`func (o *CreditMemoResponse) SetAppliedAmountMicros(v int32)`

SetAppliedAmountMicros sets AppliedAmountMicros field to given value.


### GetUnappliedAmountMicros

`func (o *CreditMemoResponse) GetUnappliedAmountMicros() int32`

GetUnappliedAmountMicros returns the UnappliedAmountMicros field if non-nil, zero value otherwise.

### GetUnappliedAmountMicrosOk

`func (o *CreditMemoResponse) GetUnappliedAmountMicrosOk() (*int32, bool)`

GetUnappliedAmountMicrosOk returns a tuple with the UnappliedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnappliedAmountMicros

`func (o *CreditMemoResponse) SetUnappliedAmountMicros(v int32)`

SetUnappliedAmountMicros sets UnappliedAmountMicros field to given value.


### GetCurrencyCode

`func (o *CreditMemoResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CreditMemoResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CreditMemoResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetSubsidiary

`func (o *CreditMemoResponse) GetSubsidiary() string`

GetSubsidiary returns the Subsidiary field if non-nil, zero value otherwise.

### GetSubsidiaryOk

`func (o *CreditMemoResponse) GetSubsidiaryOk() (*string, bool)`

GetSubsidiaryOk returns a tuple with the Subsidiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidiary

`func (o *CreditMemoResponse) SetSubsidiary(v string)`

SetSubsidiary sets Subsidiary field to given value.

### HasSubsidiary

`func (o *CreditMemoResponse) HasSubsidiary() bool`

HasSubsidiary returns a boolean if a field has been set.

### GetDescription

`func (o *CreditMemoResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditMemoResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditMemoResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditMemoResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreditMemoResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreditMemoResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreditMemoResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreditMemoResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetApprovedBy

`func (o *CreditMemoResponse) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *CreditMemoResponse) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *CreditMemoResponse) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *CreditMemoResponse) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetApprovedAt

`func (o *CreditMemoResponse) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *CreditMemoResponse) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *CreditMemoResponse) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *CreditMemoResponse) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *CreditMemoResponse) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *CreditMemoResponse) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *CreditMemoResponse) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *CreditMemoResponse) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreditMemoResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditMemoResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditMemoResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreditMemoResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreditMemoResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreditMemoResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


