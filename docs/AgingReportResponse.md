# AgingReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**AsOf** | **time.Time** |  | 
**TotalOutstandingMicros** | **int32** |  | 
**TotalCases** | **int32** |  | 
**Buckets** | [**[]AgingBucketDetail**](AgingBucketDetail.md) |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]

## Methods

### NewAgingReportResponse

`func NewAgingReportResponse(tenantId string, asOf time.Time, totalOutstandingMicros int32, totalCases int32, buckets []AgingBucketDetail, ) *AgingReportResponse`

NewAgingReportResponse instantiates a new AgingReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgingReportResponseWithDefaults

`func NewAgingReportResponseWithDefaults() *AgingReportResponse`

NewAgingReportResponseWithDefaults instantiates a new AgingReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AgingReportResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AgingReportResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AgingReportResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsOf

`func (o *AgingReportResponse) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *AgingReportResponse) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *AgingReportResponse) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.


### GetTotalOutstandingMicros

`func (o *AgingReportResponse) GetTotalOutstandingMicros() int32`

GetTotalOutstandingMicros returns the TotalOutstandingMicros field if non-nil, zero value otherwise.

### GetTotalOutstandingMicrosOk

`func (o *AgingReportResponse) GetTotalOutstandingMicrosOk() (*int32, bool)`

GetTotalOutstandingMicrosOk returns a tuple with the TotalOutstandingMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstandingMicros

`func (o *AgingReportResponse) SetTotalOutstandingMicros(v int32)`

SetTotalOutstandingMicros sets TotalOutstandingMicros field to given value.


### GetTotalCases

`func (o *AgingReportResponse) GetTotalCases() int32`

GetTotalCases returns the TotalCases field if non-nil, zero value otherwise.

### GetTotalCasesOk

`func (o *AgingReportResponse) GetTotalCasesOk() (*int32, bool)`

GetTotalCasesOk returns a tuple with the TotalCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCases

`func (o *AgingReportResponse) SetTotalCases(v int32)`

SetTotalCases sets TotalCases field to given value.


### GetBuckets

`func (o *AgingReportResponse) GetBuckets() []AgingBucketDetail`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AgingReportResponse) GetBucketsOk() (*[]AgingBucketDetail, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AgingReportResponse) SetBuckets(v []AgingBucketDetail)`

SetBuckets sets Buckets field to given value.


### GetCurrencyCode

`func (o *AgingReportResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AgingReportResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AgingReportResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AgingReportResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


