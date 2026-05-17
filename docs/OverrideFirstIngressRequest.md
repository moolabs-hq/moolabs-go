# OverrideFirstIngressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**UsageEventId** | **string** | External usage event id | 
**NewFirstIngressAt** | **time.Time** | Replacement ingress timestamp | 
**OverrideTicketId** | **string** | Break-glass ticket id (INC/CHG/JIRA) | 

## Methods

### NewOverrideFirstIngressRequest

`func NewOverrideFirstIngressRequest(tenantId string, usageEventId string, newFirstIngressAt time.Time, overrideTicketId string, ) *OverrideFirstIngressRequest`

NewOverrideFirstIngressRequest instantiates a new OverrideFirstIngressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideFirstIngressRequestWithDefaults

`func NewOverrideFirstIngressRequestWithDefaults() *OverrideFirstIngressRequest`

NewOverrideFirstIngressRequestWithDefaults instantiates a new OverrideFirstIngressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *OverrideFirstIngressRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OverrideFirstIngressRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OverrideFirstIngressRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUsageEventId

`func (o *OverrideFirstIngressRequest) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *OverrideFirstIngressRequest) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *OverrideFirstIngressRequest) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.


### GetNewFirstIngressAt

`func (o *OverrideFirstIngressRequest) GetNewFirstIngressAt() time.Time`

GetNewFirstIngressAt returns the NewFirstIngressAt field if non-nil, zero value otherwise.

### GetNewFirstIngressAtOk

`func (o *OverrideFirstIngressRequest) GetNewFirstIngressAtOk() (*time.Time, bool)`

GetNewFirstIngressAtOk returns a tuple with the NewFirstIngressAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFirstIngressAt

`func (o *OverrideFirstIngressRequest) SetNewFirstIngressAt(v time.Time)`

SetNewFirstIngressAt sets NewFirstIngressAt field to given value.


### GetOverrideTicketId

`func (o *OverrideFirstIngressRequest) GetOverrideTicketId() string`

GetOverrideTicketId returns the OverrideTicketId field if non-nil, zero value otherwise.

### GetOverrideTicketIdOk

`func (o *OverrideFirstIngressRequest) GetOverrideTicketIdOk() (*string, bool)`

GetOverrideTicketIdOk returns a tuple with the OverrideTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideTicketId

`func (o *OverrideFirstIngressRequest) SetOverrideTicketId(v string)`

SetOverrideTicketId sets OverrideTicketId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


