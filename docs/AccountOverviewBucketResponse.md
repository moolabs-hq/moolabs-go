# AccountOverviewBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**InvoiceCount** | **int32** |  | 
**TotalAmountMicros** | **int32** |  | 

## Methods

### NewAccountOverviewBucketResponse

`func NewAccountOverviewBucketResponse(bucket string, invoiceCount int32, totalAmountMicros int32, ) *AccountOverviewBucketResponse`

NewAccountOverviewBucketResponse instantiates a new AccountOverviewBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOverviewBucketResponseWithDefaults

`func NewAccountOverviewBucketResponseWithDefaults() *AccountOverviewBucketResponse`

NewAccountOverviewBucketResponseWithDefaults instantiates a new AccountOverviewBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *AccountOverviewBucketResponse) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AccountOverviewBucketResponse) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AccountOverviewBucketResponse) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetInvoiceCount

`func (o *AccountOverviewBucketResponse) GetInvoiceCount() int32`

GetInvoiceCount returns the InvoiceCount field if non-nil, zero value otherwise.

### GetInvoiceCountOk

`func (o *AccountOverviewBucketResponse) GetInvoiceCountOk() (*int32, bool)`

GetInvoiceCountOk returns a tuple with the InvoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCount

`func (o *AccountOverviewBucketResponse) SetInvoiceCount(v int32)`

SetInvoiceCount sets InvoiceCount field to given value.


### GetTotalAmountMicros

`func (o *AccountOverviewBucketResponse) GetTotalAmountMicros() int32`

GetTotalAmountMicros returns the TotalAmountMicros field if non-nil, zero value otherwise.

### GetTotalAmountMicrosOk

`func (o *AccountOverviewBucketResponse) GetTotalAmountMicrosOk() (*int32, bool)`

GetTotalAmountMicrosOk returns a tuple with the TotalAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMicros

`func (o *AccountOverviewBucketResponse) SetTotalAmountMicros(v int32)`

SetTotalAmountMicros sets TotalAmountMicros field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


