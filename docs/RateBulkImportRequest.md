# RateBulkImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Rates** | [**[]RateCatalogCreate**](RateCatalogCreate.md) |  | 

## Methods

### NewRateBulkImportRequest

`func NewRateBulkImportRequest(tenantId string, rates []RateCatalogCreate, ) *RateBulkImportRequest`

NewRateBulkImportRequest instantiates a new RateBulkImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateBulkImportRequestWithDefaults

`func NewRateBulkImportRequestWithDefaults() *RateBulkImportRequest`

NewRateBulkImportRequestWithDefaults instantiates a new RateBulkImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RateBulkImportRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RateBulkImportRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RateBulkImportRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetRates

`func (o *RateBulkImportRequest) GetRates() []RateCatalogCreate`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *RateBulkImportRequest) GetRatesOk() (*[]RateCatalogCreate, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *RateBulkImportRequest) SetRates(v []RateCatalogCreate)`

SetRates sets Rates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


