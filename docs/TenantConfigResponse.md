# TenantConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**Region** | **string** |  | 
**Endpoints** | [**EndpointsResponse**](EndpointsResponse.md) |  | 

## Methods

### NewTenantConfigResponse

`func NewTenantConfigResponse(tenantId string, region string, endpoints EndpointsResponse, ) *TenantConfigResponse`

NewTenantConfigResponse instantiates a new TenantConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantConfigResponseWithDefaults

`func NewTenantConfigResponseWithDefaults() *TenantConfigResponse`

NewTenantConfigResponseWithDefaults instantiates a new TenantConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantConfigResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantConfigResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantConfigResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetRegion

`func (o *TenantConfigResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TenantConfigResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TenantConfigResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetEndpoints

`func (o *TenantConfigResponse) GetEndpoints() EndpointsResponse`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *TenantConfigResponse) GetEndpointsOk() (*EndpointsResponse, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *TenantConfigResponse) SetEndpoints(v EndpointsResponse)`

SetEndpoints sets Endpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


