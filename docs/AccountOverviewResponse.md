# AccountOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**TotalInvoices** | **int32** |  | 
**TotalAmountMicros** | **int32** |  | 
**LastPaymentAt** | Pointer to **time.Time** |  | [optional] 
**Buckets** | [**[]AccountOverviewBucketResponse**](AccountOverviewBucketResponse.md) |  | 

## Methods

### NewAccountOverviewResponse

`func NewAccountOverviewResponse(accountId string, totalInvoices int32, totalAmountMicros int32, buckets []AccountOverviewBucketResponse, ) *AccountOverviewResponse`

NewAccountOverviewResponse instantiates a new AccountOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOverviewResponseWithDefaults

`func NewAccountOverviewResponseWithDefaults() *AccountOverviewResponse`

NewAccountOverviewResponseWithDefaults instantiates a new AccountOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountOverviewResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountOverviewResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountOverviewResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCurrencyCode

`func (o *AccountOverviewResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountOverviewResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountOverviewResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountOverviewResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTotalInvoices

`func (o *AccountOverviewResponse) GetTotalInvoices() int32`

GetTotalInvoices returns the TotalInvoices field if non-nil, zero value otherwise.

### GetTotalInvoicesOk

`func (o *AccountOverviewResponse) GetTotalInvoicesOk() (*int32, bool)`

GetTotalInvoicesOk returns a tuple with the TotalInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvoices

`func (o *AccountOverviewResponse) SetTotalInvoices(v int32)`

SetTotalInvoices sets TotalInvoices field to given value.


### GetTotalAmountMicros

`func (o *AccountOverviewResponse) GetTotalAmountMicros() int32`

GetTotalAmountMicros returns the TotalAmountMicros field if non-nil, zero value otherwise.

### GetTotalAmountMicrosOk

`func (o *AccountOverviewResponse) GetTotalAmountMicrosOk() (*int32, bool)`

GetTotalAmountMicrosOk returns a tuple with the TotalAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMicros

`func (o *AccountOverviewResponse) SetTotalAmountMicros(v int32)`

SetTotalAmountMicros sets TotalAmountMicros field to given value.


### GetLastPaymentAt

`func (o *AccountOverviewResponse) GetLastPaymentAt() time.Time`

GetLastPaymentAt returns the LastPaymentAt field if non-nil, zero value otherwise.

### GetLastPaymentAtOk

`func (o *AccountOverviewResponse) GetLastPaymentAtOk() (*time.Time, bool)`

GetLastPaymentAtOk returns a tuple with the LastPaymentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAt

`func (o *AccountOverviewResponse) SetLastPaymentAt(v time.Time)`

SetLastPaymentAt sets LastPaymentAt field to given value.

### HasLastPaymentAt

`func (o *AccountOverviewResponse) HasLastPaymentAt() bool`

HasLastPaymentAt returns a boolean if a field has been set.

### GetBuckets

`func (o *AccountOverviewResponse) GetBuckets() []AccountOverviewBucketResponse`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AccountOverviewResponse) GetBucketsOk() (*[]AccountOverviewBucketResponse, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AccountOverviewResponse) SetBuckets(v []AccountOverviewBucketResponse)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


