# ComputeVarianceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**TenantId** | **string** |  | 
**FeatureKey** | **string** |  | 
**PeriodStart** | **string** |  | 
**PeriodEnd** | **string** |  | 
**BomTemplateId** | **string** |  | 
**Status** | **string** |  | 
**Variance** | Pointer to [**VarianceDecompositionOut**](VarianceDecompositionOut.md) |  | [optional] 

## Methods

### NewComputeVarianceResponse

`func NewComputeVarianceResponse(jobId string, tenantId string, featureKey string, periodStart string, periodEnd string, bomTemplateId string, status string, ) *ComputeVarianceResponse`

NewComputeVarianceResponse instantiates a new ComputeVarianceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeVarianceResponseWithDefaults

`func NewComputeVarianceResponseWithDefaults() *ComputeVarianceResponse`

NewComputeVarianceResponseWithDefaults instantiates a new ComputeVarianceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ComputeVarianceResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ComputeVarianceResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ComputeVarianceResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetTenantId

`func (o *ComputeVarianceResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ComputeVarianceResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ComputeVarianceResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetFeatureKey

`func (o *ComputeVarianceResponse) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *ComputeVarianceResponse) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *ComputeVarianceResponse) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetPeriodStart

`func (o *ComputeVarianceResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ComputeVarianceResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ComputeVarianceResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *ComputeVarianceResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ComputeVarianceResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ComputeVarianceResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetBomTemplateId

`func (o *ComputeVarianceResponse) GetBomTemplateId() string`

GetBomTemplateId returns the BomTemplateId field if non-nil, zero value otherwise.

### GetBomTemplateIdOk

`func (o *ComputeVarianceResponse) GetBomTemplateIdOk() (*string, bool)`

GetBomTemplateIdOk returns a tuple with the BomTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomTemplateId

`func (o *ComputeVarianceResponse) SetBomTemplateId(v string)`

SetBomTemplateId sets BomTemplateId field to given value.


### GetStatus

`func (o *ComputeVarianceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeVarianceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeVarianceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVariance

`func (o *ComputeVarianceResponse) GetVariance() VarianceDecompositionOut`

GetVariance returns the Variance field if non-nil, zero value otherwise.

### GetVarianceOk

`func (o *ComputeVarianceResponse) GetVarianceOk() (*VarianceDecompositionOut, bool)`

GetVarianceOk returns a tuple with the Variance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariance

`func (o *ComputeVarianceResponse) SetVariance(v VarianceDecompositionOut)`

SetVariance sets Variance field to given value.

### HasVariance

`func (o *ComputeVarianceResponse) HasVariance() bool`

HasVariance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


