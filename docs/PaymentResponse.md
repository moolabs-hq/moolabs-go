# PaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**CaseId** | **string** |  | 
**RemittanceId** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | **string** |  | 
**BaseCurrencyCode** | **string** |  | 
**AmountBaseMicros** | **int32** |  | 
**PaymentMethod** | **string** |  | 
**ExternalPaymentId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**MatchConfidence** | Pointer to **float32** |  | [optional] 
**MatchMethod** | Pointer to **string** |  | [optional] 
**VerificationStatus** | **string** |  | 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 
**VerifiedSource** | Pointer to **string** |  | [optional] 
**VerifiedExternalPaymentId** | Pointer to **string** |  | [optional] 
**ReceivedAt** | **time.Time** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPaymentResponse

`func NewPaymentResponse(id string, tenantId string, caseId string, amountMicros int32, currencyCode string, baseCurrencyCode string, amountBaseMicros int32, paymentMethod string, status string, verificationStatus string, receivedAt time.Time, createdAt time.Time, updatedAt time.Time, ) *PaymentResponse`

NewPaymentResponse instantiates a new PaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseWithDefaults

`func NewPaymentResponseWithDefaults() *PaymentResponse`

NewPaymentResponseWithDefaults instantiates a new PaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *PaymentResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCaseId

`func (o *PaymentResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PaymentResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PaymentResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetRemittanceId

`func (o *PaymentResponse) GetRemittanceId() string`

GetRemittanceId returns the RemittanceId field if non-nil, zero value otherwise.

### GetRemittanceIdOk

`func (o *PaymentResponse) GetRemittanceIdOk() (*string, bool)`

GetRemittanceIdOk returns a tuple with the RemittanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemittanceId

`func (o *PaymentResponse) SetRemittanceId(v string)`

SetRemittanceId sets RemittanceId field to given value.

### HasRemittanceId

`func (o *PaymentResponse) HasRemittanceId() bool`

HasRemittanceId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PaymentResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaymentResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetAmountMicros

`func (o *PaymentResponse) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *PaymentResponse) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *PaymentResponse) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *PaymentResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PaymentResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PaymentResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetBaseCurrencyCode

`func (o *PaymentResponse) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *PaymentResponse) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *PaymentResponse) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.


### GetAmountBaseMicros

`func (o *PaymentResponse) GetAmountBaseMicros() int32`

GetAmountBaseMicros returns the AmountBaseMicros field if non-nil, zero value otherwise.

### GetAmountBaseMicrosOk

`func (o *PaymentResponse) GetAmountBaseMicrosOk() (*int32, bool)`

GetAmountBaseMicrosOk returns a tuple with the AmountBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBaseMicros

`func (o *PaymentResponse) SetAmountBaseMicros(v int32)`

SetAmountBaseMicros sets AmountBaseMicros field to given value.


### GetPaymentMethod

`func (o *PaymentResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetExternalPaymentId

`func (o *PaymentResponse) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *PaymentResponse) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *PaymentResponse) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *PaymentResponse) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMatchConfidence

`func (o *PaymentResponse) GetMatchConfidence() float32`

GetMatchConfidence returns the MatchConfidence field if non-nil, zero value otherwise.

### GetMatchConfidenceOk

`func (o *PaymentResponse) GetMatchConfidenceOk() (*float32, bool)`

GetMatchConfidenceOk returns a tuple with the MatchConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchConfidence

`func (o *PaymentResponse) SetMatchConfidence(v float32)`

SetMatchConfidence sets MatchConfidence field to given value.

### HasMatchConfidence

`func (o *PaymentResponse) HasMatchConfidence() bool`

HasMatchConfidence returns a boolean if a field has been set.

### GetMatchMethod

`func (o *PaymentResponse) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *PaymentResponse) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *PaymentResponse) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.

### HasMatchMethod

`func (o *PaymentResponse) HasMatchMethod() bool`

HasMatchMethod returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *PaymentResponse) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PaymentResponse) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PaymentResponse) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.


### GetVerifiedAt

`func (o *PaymentResponse) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *PaymentResponse) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *PaymentResponse) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *PaymentResponse) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetVerifiedSource

`func (o *PaymentResponse) GetVerifiedSource() string`

GetVerifiedSource returns the VerifiedSource field if non-nil, zero value otherwise.

### GetVerifiedSourceOk

`func (o *PaymentResponse) GetVerifiedSourceOk() (*string, bool)`

GetVerifiedSourceOk returns a tuple with the VerifiedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedSource

`func (o *PaymentResponse) SetVerifiedSource(v string)`

SetVerifiedSource sets VerifiedSource field to given value.

### HasVerifiedSource

`func (o *PaymentResponse) HasVerifiedSource() bool`

HasVerifiedSource returns a boolean if a field has been set.

### GetVerifiedExternalPaymentId

`func (o *PaymentResponse) GetVerifiedExternalPaymentId() string`

GetVerifiedExternalPaymentId returns the VerifiedExternalPaymentId field if non-nil, zero value otherwise.

### GetVerifiedExternalPaymentIdOk

`func (o *PaymentResponse) GetVerifiedExternalPaymentIdOk() (*string, bool)`

GetVerifiedExternalPaymentIdOk returns a tuple with the VerifiedExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedExternalPaymentId

`func (o *PaymentResponse) SetVerifiedExternalPaymentId(v string)`

SetVerifiedExternalPaymentId sets VerifiedExternalPaymentId field to given value.

### HasVerifiedExternalPaymentId

`func (o *PaymentResponse) HasVerifiedExternalPaymentId() bool`

HasVerifiedExternalPaymentId returns a boolean if a field has been set.

### GetReceivedAt

`func (o *PaymentResponse) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *PaymentResponse) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *PaymentResponse) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetMetadata

`func (o *PaymentResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PaymentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


