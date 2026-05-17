# CEIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**AsOf** | **time.Time** |  | 
**CeiPercentage** | **float32** |  | 
**BeginningReceivablesMicros** | **int32** |  | 
**EndingReceivablesMicros** | **int32** |  | 
**CreditSalesMicros** | **int32** |  | 
**CashReceiptsMicros** | **int32** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]

## Methods

### NewCEIResponse

`func NewCEIResponse(tenantId string, asOf time.Time, ceiPercentage float32, beginningReceivablesMicros int32, endingReceivablesMicros int32, creditSalesMicros int32, cashReceiptsMicros int32, ) *CEIResponse`

NewCEIResponse instantiates a new CEIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCEIResponseWithDefaults

`func NewCEIResponseWithDefaults() *CEIResponse`

NewCEIResponseWithDefaults instantiates a new CEIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CEIResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CEIResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CEIResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsOf

`func (o *CEIResponse) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CEIResponse) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CEIResponse) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.


### GetCeiPercentage

`func (o *CEIResponse) GetCeiPercentage() float32`

GetCeiPercentage returns the CeiPercentage field if non-nil, zero value otherwise.

### GetCeiPercentageOk

`func (o *CEIResponse) GetCeiPercentageOk() (*float32, bool)`

GetCeiPercentageOk returns a tuple with the CeiPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiPercentage

`func (o *CEIResponse) SetCeiPercentage(v float32)`

SetCeiPercentage sets CeiPercentage field to given value.


### GetBeginningReceivablesMicros

`func (o *CEIResponse) GetBeginningReceivablesMicros() int32`

GetBeginningReceivablesMicros returns the BeginningReceivablesMicros field if non-nil, zero value otherwise.

### GetBeginningReceivablesMicrosOk

`func (o *CEIResponse) GetBeginningReceivablesMicrosOk() (*int32, bool)`

GetBeginningReceivablesMicrosOk returns a tuple with the BeginningReceivablesMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginningReceivablesMicros

`func (o *CEIResponse) SetBeginningReceivablesMicros(v int32)`

SetBeginningReceivablesMicros sets BeginningReceivablesMicros field to given value.


### GetEndingReceivablesMicros

`func (o *CEIResponse) GetEndingReceivablesMicros() int32`

GetEndingReceivablesMicros returns the EndingReceivablesMicros field if non-nil, zero value otherwise.

### GetEndingReceivablesMicrosOk

`func (o *CEIResponse) GetEndingReceivablesMicrosOk() (*int32, bool)`

GetEndingReceivablesMicrosOk returns a tuple with the EndingReceivablesMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingReceivablesMicros

`func (o *CEIResponse) SetEndingReceivablesMicros(v int32)`

SetEndingReceivablesMicros sets EndingReceivablesMicros field to given value.


### GetCreditSalesMicros

`func (o *CEIResponse) GetCreditSalesMicros() int32`

GetCreditSalesMicros returns the CreditSalesMicros field if non-nil, zero value otherwise.

### GetCreditSalesMicrosOk

`func (o *CEIResponse) GetCreditSalesMicrosOk() (*int32, bool)`

GetCreditSalesMicrosOk returns a tuple with the CreditSalesMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditSalesMicros

`func (o *CEIResponse) SetCreditSalesMicros(v int32)`

SetCreditSalesMicros sets CreditSalesMicros field to given value.


### GetCashReceiptsMicros

`func (o *CEIResponse) GetCashReceiptsMicros() int32`

GetCashReceiptsMicros returns the CashReceiptsMicros field if non-nil, zero value otherwise.

### GetCashReceiptsMicrosOk

`func (o *CEIResponse) GetCashReceiptsMicrosOk() (*int32, bool)`

GetCashReceiptsMicrosOk returns a tuple with the CashReceiptsMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashReceiptsMicros

`func (o *CEIResponse) SetCashReceiptsMicros(v int32)`

SetCashReceiptsMicros sets CashReceiptsMicros field to given value.


### GetCurrencyCode

`func (o *CEIResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CEIResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CEIResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CEIResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


