# CreditMemoCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CaseId** | Pointer to **string** |  | [optional] 
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Reason** | **string** | billing_error, service_credit, commercial_concession, dispute_resolution, pricing_adjustment, other | 
**CreditAmountMicros** | **int32** |  | 
**CurrencyCode** | **string** |  | 
**Subsidiary** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCreditMemoCreate

`func NewCreditMemoCreate(accountId string, reason string, creditAmountMicros int32, currencyCode string, ) *CreditMemoCreate`

NewCreditMemoCreate instantiates a new CreditMemoCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditMemoCreateWithDefaults

`func NewCreditMemoCreateWithDefaults() *CreditMemoCreate`

NewCreditMemoCreateWithDefaults instantiates a new CreditMemoCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreditMemoCreate) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditMemoCreate) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditMemoCreate) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCaseId

`func (o *CreditMemoCreate) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CreditMemoCreate) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CreditMemoCreate) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *CreditMemoCreate) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *CreditMemoCreate) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *CreditMemoCreate) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *CreditMemoCreate) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *CreditMemoCreate) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetReason

`func (o *CreditMemoCreate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditMemoCreate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditMemoCreate) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreditAmountMicros

`func (o *CreditMemoCreate) GetCreditAmountMicros() int32`

GetCreditAmountMicros returns the CreditAmountMicros field if non-nil, zero value otherwise.

### GetCreditAmountMicrosOk

`func (o *CreditMemoCreate) GetCreditAmountMicrosOk() (*int32, bool)`

GetCreditAmountMicrosOk returns a tuple with the CreditAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountMicros

`func (o *CreditMemoCreate) SetCreditAmountMicros(v int32)`

SetCreditAmountMicros sets CreditAmountMicros field to given value.


### GetCurrencyCode

`func (o *CreditMemoCreate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CreditMemoCreate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CreditMemoCreate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetSubsidiary

`func (o *CreditMemoCreate) GetSubsidiary() string`

GetSubsidiary returns the Subsidiary field if non-nil, zero value otherwise.

### GetSubsidiaryOk

`func (o *CreditMemoCreate) GetSubsidiaryOk() (*string, bool)`

GetSubsidiaryOk returns a tuple with the Subsidiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidiary

`func (o *CreditMemoCreate) SetSubsidiary(v string)`

SetSubsidiary sets Subsidiary field to given value.

### HasSubsidiary

`func (o *CreditMemoCreate) HasSubsidiary() bool`

HasSubsidiary returns a boolean if a field has been set.

### GetDescription

`func (o *CreditMemoCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditMemoCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditMemoCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditMemoCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreditMemoCreate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreditMemoCreate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreditMemoCreate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreditMemoCreate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


