# ImportPromiseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalPtpId** | **string** |  | 
**CaseId** | **string** |  | 
**ContactId** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**PromisedAmountMicros** | **int32** |  | 
**FulfilledAmountMicros** | Pointer to **int32** |  | [optional] [default to 0]
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**PromisedDate** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] [default to "open"]
**CapturedFromChannel** | Pointer to **string** |  | [optional] [default to "email"]
**CapturedBy** | Pointer to **string** |  | [optional] [default to "human"]
**Confidence** | Pointer to **float32** |  | [optional] 
**PtpNo** | Pointer to **string** |  | [optional] 
**BalanceAmountMicros** | Pointer to **int32** |  | [optional] 
**AmountBaseCurrencyMicros** | Pointer to **int32** |  | [optional] 
**BalanceAmountBaseCurrencyMicros** | Pointer to **int32** |  | [optional] 
**SlippageCount** | Pointer to **int32** |  | [optional] [default to 0]
**SourceApp** | Pointer to **string** |  | [optional] 
**PtpType** | Pointer to **string** |  | [optional] 

## Methods

### NewImportPromiseItem

`func NewImportPromiseItem(externalPtpId string, caseId string, promisedAmountMicros int32, promisedDate string, ) *ImportPromiseItem`

NewImportPromiseItem instantiates a new ImportPromiseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportPromiseItemWithDefaults

`func NewImportPromiseItemWithDefaults() *ImportPromiseItem`

NewImportPromiseItemWithDefaults instantiates a new ImportPromiseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalPtpId

`func (o *ImportPromiseItem) GetExternalPtpId() string`

GetExternalPtpId returns the ExternalPtpId field if non-nil, zero value otherwise.

### GetExternalPtpIdOk

`func (o *ImportPromiseItem) GetExternalPtpIdOk() (*string, bool)`

GetExternalPtpIdOk returns a tuple with the ExternalPtpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPtpId

`func (o *ImportPromiseItem) SetExternalPtpId(v string)`

SetExternalPtpId sets ExternalPtpId field to given value.


### GetCaseId

`func (o *ImportPromiseItem) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ImportPromiseItem) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ImportPromiseItem) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetContactId

`func (o *ImportPromiseItem) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ImportPromiseItem) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ImportPromiseItem) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ImportPromiseItem) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *ImportPromiseItem) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ImportPromiseItem) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ImportPromiseItem) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ImportPromiseItem) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPromisedAmountMicros

`func (o *ImportPromiseItem) GetPromisedAmountMicros() int32`

GetPromisedAmountMicros returns the PromisedAmountMicros field if non-nil, zero value otherwise.

### GetPromisedAmountMicrosOk

`func (o *ImportPromiseItem) GetPromisedAmountMicrosOk() (*int32, bool)`

GetPromisedAmountMicrosOk returns a tuple with the PromisedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedAmountMicros

`func (o *ImportPromiseItem) SetPromisedAmountMicros(v int32)`

SetPromisedAmountMicros sets PromisedAmountMicros field to given value.


### GetFulfilledAmountMicros

`func (o *ImportPromiseItem) GetFulfilledAmountMicros() int32`

GetFulfilledAmountMicros returns the FulfilledAmountMicros field if non-nil, zero value otherwise.

### GetFulfilledAmountMicrosOk

`func (o *ImportPromiseItem) GetFulfilledAmountMicrosOk() (*int32, bool)`

GetFulfilledAmountMicrosOk returns a tuple with the FulfilledAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAmountMicros

`func (o *ImportPromiseItem) SetFulfilledAmountMicros(v int32)`

SetFulfilledAmountMicros sets FulfilledAmountMicros field to given value.

### HasFulfilledAmountMicros

`func (o *ImportPromiseItem) HasFulfilledAmountMicros() bool`

HasFulfilledAmountMicros returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ImportPromiseItem) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ImportPromiseItem) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ImportPromiseItem) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ImportPromiseItem) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPromisedDate

`func (o *ImportPromiseItem) GetPromisedDate() string`

GetPromisedDate returns the PromisedDate field if non-nil, zero value otherwise.

### GetPromisedDateOk

`func (o *ImportPromiseItem) GetPromisedDateOk() (*string, bool)`

GetPromisedDateOk returns a tuple with the PromisedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedDate

`func (o *ImportPromiseItem) SetPromisedDate(v string)`

SetPromisedDate sets PromisedDate field to given value.


### GetStatus

`func (o *ImportPromiseItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportPromiseItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportPromiseItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportPromiseItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCapturedFromChannel

`func (o *ImportPromiseItem) GetCapturedFromChannel() string`

GetCapturedFromChannel returns the CapturedFromChannel field if non-nil, zero value otherwise.

### GetCapturedFromChannelOk

`func (o *ImportPromiseItem) GetCapturedFromChannelOk() (*string, bool)`

GetCapturedFromChannelOk returns a tuple with the CapturedFromChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedFromChannel

`func (o *ImportPromiseItem) SetCapturedFromChannel(v string)`

SetCapturedFromChannel sets CapturedFromChannel field to given value.

### HasCapturedFromChannel

`func (o *ImportPromiseItem) HasCapturedFromChannel() bool`

HasCapturedFromChannel returns a boolean if a field has been set.

### GetCapturedBy

`func (o *ImportPromiseItem) GetCapturedBy() string`

GetCapturedBy returns the CapturedBy field if non-nil, zero value otherwise.

### GetCapturedByOk

`func (o *ImportPromiseItem) GetCapturedByOk() (*string, bool)`

GetCapturedByOk returns a tuple with the CapturedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedBy

`func (o *ImportPromiseItem) SetCapturedBy(v string)`

SetCapturedBy sets CapturedBy field to given value.

### HasCapturedBy

`func (o *ImportPromiseItem) HasCapturedBy() bool`

HasCapturedBy returns a boolean if a field has been set.

### GetConfidence

`func (o *ImportPromiseItem) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ImportPromiseItem) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ImportPromiseItem) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ImportPromiseItem) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetPtpNo

`func (o *ImportPromiseItem) GetPtpNo() string`

GetPtpNo returns the PtpNo field if non-nil, zero value otherwise.

### GetPtpNoOk

`func (o *ImportPromiseItem) GetPtpNoOk() (*string, bool)`

GetPtpNoOk returns a tuple with the PtpNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpNo

`func (o *ImportPromiseItem) SetPtpNo(v string)`

SetPtpNo sets PtpNo field to given value.

### HasPtpNo

`func (o *ImportPromiseItem) HasPtpNo() bool`

HasPtpNo returns a boolean if a field has been set.

### GetBalanceAmountMicros

`func (o *ImportPromiseItem) GetBalanceAmountMicros() int32`

GetBalanceAmountMicros returns the BalanceAmountMicros field if non-nil, zero value otherwise.

### GetBalanceAmountMicrosOk

`func (o *ImportPromiseItem) GetBalanceAmountMicrosOk() (*int32, bool)`

GetBalanceAmountMicrosOk returns a tuple with the BalanceAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmountMicros

`func (o *ImportPromiseItem) SetBalanceAmountMicros(v int32)`

SetBalanceAmountMicros sets BalanceAmountMicros field to given value.

### HasBalanceAmountMicros

`func (o *ImportPromiseItem) HasBalanceAmountMicros() bool`

HasBalanceAmountMicros returns a boolean if a field has been set.

### GetAmountBaseCurrencyMicros

`func (o *ImportPromiseItem) GetAmountBaseCurrencyMicros() int32`

GetAmountBaseCurrencyMicros returns the AmountBaseCurrencyMicros field if non-nil, zero value otherwise.

### GetAmountBaseCurrencyMicrosOk

`func (o *ImportPromiseItem) GetAmountBaseCurrencyMicrosOk() (*int32, bool)`

GetAmountBaseCurrencyMicrosOk returns a tuple with the AmountBaseCurrencyMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBaseCurrencyMicros

`func (o *ImportPromiseItem) SetAmountBaseCurrencyMicros(v int32)`

SetAmountBaseCurrencyMicros sets AmountBaseCurrencyMicros field to given value.

### HasAmountBaseCurrencyMicros

`func (o *ImportPromiseItem) HasAmountBaseCurrencyMicros() bool`

HasAmountBaseCurrencyMicros returns a boolean if a field has been set.

### GetBalanceAmountBaseCurrencyMicros

`func (o *ImportPromiseItem) GetBalanceAmountBaseCurrencyMicros() int32`

GetBalanceAmountBaseCurrencyMicros returns the BalanceAmountBaseCurrencyMicros field if non-nil, zero value otherwise.

### GetBalanceAmountBaseCurrencyMicrosOk

`func (o *ImportPromiseItem) GetBalanceAmountBaseCurrencyMicrosOk() (*int32, bool)`

GetBalanceAmountBaseCurrencyMicrosOk returns a tuple with the BalanceAmountBaseCurrencyMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmountBaseCurrencyMicros

`func (o *ImportPromiseItem) SetBalanceAmountBaseCurrencyMicros(v int32)`

SetBalanceAmountBaseCurrencyMicros sets BalanceAmountBaseCurrencyMicros field to given value.

### HasBalanceAmountBaseCurrencyMicros

`func (o *ImportPromiseItem) HasBalanceAmountBaseCurrencyMicros() bool`

HasBalanceAmountBaseCurrencyMicros returns a boolean if a field has been set.

### GetSlippageCount

`func (o *ImportPromiseItem) GetSlippageCount() int32`

GetSlippageCount returns the SlippageCount field if non-nil, zero value otherwise.

### GetSlippageCountOk

`func (o *ImportPromiseItem) GetSlippageCountOk() (*int32, bool)`

GetSlippageCountOk returns a tuple with the SlippageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageCount

`func (o *ImportPromiseItem) SetSlippageCount(v int32)`

SetSlippageCount sets SlippageCount field to given value.

### HasSlippageCount

`func (o *ImportPromiseItem) HasSlippageCount() bool`

HasSlippageCount returns a boolean if a field has been set.

### GetSourceApp

`func (o *ImportPromiseItem) GetSourceApp() string`

GetSourceApp returns the SourceApp field if non-nil, zero value otherwise.

### GetSourceAppOk

`func (o *ImportPromiseItem) GetSourceAppOk() (*string, bool)`

GetSourceAppOk returns a tuple with the SourceApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceApp

`func (o *ImportPromiseItem) SetSourceApp(v string)`

SetSourceApp sets SourceApp field to given value.

### HasSourceApp

`func (o *ImportPromiseItem) HasSourceApp() bool`

HasSourceApp returns a boolean if a field has been set.

### GetPtpType

`func (o *ImportPromiseItem) GetPtpType() string`

GetPtpType returns the PtpType field if non-nil, zero value otherwise.

### GetPtpTypeOk

`func (o *ImportPromiseItem) GetPtpTypeOk() (*string, bool)`

GetPtpTypeOk returns a tuple with the PtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpType

`func (o *ImportPromiseItem) SetPtpType(v string)`

SetPtpType sets PtpType field to given value.

### HasPtpType

`func (o *ImportPromiseItem) HasPtpType() bool`

HasPtpType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


