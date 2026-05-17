# HostedCheckoutSettlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **string** |  | 
**PtpId** | **string** |  | 
**InvoiceId** | **string** |  | 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**BaseCurrencyCode** | Pointer to **string** |  | [optional] 
**AmountBaseMicros** | Pointer to **int32** |  | [optional] 
**ExternalPaymentId** | **string** |  | 
**ReceivedAt** | **time.Time** |  | 
**PaymentMethod** | Pointer to **string** | stripe, ach, wire, check, credit_adjustment, other | [optional] [default to "stripe"]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewHostedCheckoutSettlementRequest

`func NewHostedCheckoutSettlementRequest(caseId string, ptpId string, invoiceId string, amountMicros int32, externalPaymentId string, receivedAt time.Time, ) *HostedCheckoutSettlementRequest`

NewHostedCheckoutSettlementRequest instantiates a new HostedCheckoutSettlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCheckoutSettlementRequestWithDefaults

`func NewHostedCheckoutSettlementRequestWithDefaults() *HostedCheckoutSettlementRequest`

NewHostedCheckoutSettlementRequestWithDefaults instantiates a new HostedCheckoutSettlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *HostedCheckoutSettlementRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *HostedCheckoutSettlementRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *HostedCheckoutSettlementRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetPtpId

`func (o *HostedCheckoutSettlementRequest) GetPtpId() string`

GetPtpId returns the PtpId field if non-nil, zero value otherwise.

### GetPtpIdOk

`func (o *HostedCheckoutSettlementRequest) GetPtpIdOk() (*string, bool)`

GetPtpIdOk returns a tuple with the PtpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpId

`func (o *HostedCheckoutSettlementRequest) SetPtpId(v string)`

SetPtpId sets PtpId field to given value.


### GetInvoiceId

`func (o *HostedCheckoutSettlementRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *HostedCheckoutSettlementRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *HostedCheckoutSettlementRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmountMicros

`func (o *HostedCheckoutSettlementRequest) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *HostedCheckoutSettlementRequest) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *HostedCheckoutSettlementRequest) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *HostedCheckoutSettlementRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *HostedCheckoutSettlementRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *HostedCheckoutSettlementRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *HostedCheckoutSettlementRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetBaseCurrencyCode

`func (o *HostedCheckoutSettlementRequest) GetBaseCurrencyCode() string`

GetBaseCurrencyCode returns the BaseCurrencyCode field if non-nil, zero value otherwise.

### GetBaseCurrencyCodeOk

`func (o *HostedCheckoutSettlementRequest) GetBaseCurrencyCodeOk() (*string, bool)`

GetBaseCurrencyCodeOk returns a tuple with the BaseCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrencyCode

`func (o *HostedCheckoutSettlementRequest) SetBaseCurrencyCode(v string)`

SetBaseCurrencyCode sets BaseCurrencyCode field to given value.

### HasBaseCurrencyCode

`func (o *HostedCheckoutSettlementRequest) HasBaseCurrencyCode() bool`

HasBaseCurrencyCode returns a boolean if a field has been set.

### GetAmountBaseMicros

`func (o *HostedCheckoutSettlementRequest) GetAmountBaseMicros() int32`

GetAmountBaseMicros returns the AmountBaseMicros field if non-nil, zero value otherwise.

### GetAmountBaseMicrosOk

`func (o *HostedCheckoutSettlementRequest) GetAmountBaseMicrosOk() (*int32, bool)`

GetAmountBaseMicrosOk returns a tuple with the AmountBaseMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBaseMicros

`func (o *HostedCheckoutSettlementRequest) SetAmountBaseMicros(v int32)`

SetAmountBaseMicros sets AmountBaseMicros field to given value.

### HasAmountBaseMicros

`func (o *HostedCheckoutSettlementRequest) HasAmountBaseMicros() bool`

HasAmountBaseMicros returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *HostedCheckoutSettlementRequest) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *HostedCheckoutSettlementRequest) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *HostedCheckoutSettlementRequest) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.


### GetReceivedAt

`func (o *HostedCheckoutSettlementRequest) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *HostedCheckoutSettlementRequest) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *HostedCheckoutSettlementRequest) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetPaymentMethod

`func (o *HostedCheckoutSettlementRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *HostedCheckoutSettlementRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *HostedCheckoutSettlementRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *HostedCheckoutSettlementRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetMetadata

`func (o *HostedCheckoutSettlementRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostedCheckoutSettlementRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostedCheckoutSettlementRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HostedCheckoutSettlementRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


