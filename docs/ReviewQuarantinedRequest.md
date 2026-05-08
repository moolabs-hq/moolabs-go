# ReviewQuarantinedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant identifier | 
**PoolId** | **string** | Pool identifier | 
**UsageEventId** | **string** | Usage event ID to review | 
**Action** | **string** | Action: &#39;retry&#39;, &#39;manual_rate&#39;, or &#39;ignore&#39; | 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReviewQuarantinedRequest

`func NewReviewQuarantinedRequest(tenantId string, poolId string, usageEventId string, action string, ) *ReviewQuarantinedRequest`

NewReviewQuarantinedRequest instantiates a new ReviewQuarantinedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewQuarantinedRequestWithDefaults

`func NewReviewQuarantinedRequestWithDefaults() *ReviewQuarantinedRequest`

NewReviewQuarantinedRequestWithDefaults instantiates a new ReviewQuarantinedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ReviewQuarantinedRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ReviewQuarantinedRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ReviewQuarantinedRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetPoolId

`func (o *ReviewQuarantinedRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ReviewQuarantinedRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ReviewQuarantinedRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetUsageEventId

`func (o *ReviewQuarantinedRequest) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *ReviewQuarantinedRequest) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *ReviewQuarantinedRequest) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.


### GetAction

`func (o *ReviewQuarantinedRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReviewQuarantinedRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReviewQuarantinedRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetNotes

`func (o *ReviewQuarantinedRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ReviewQuarantinedRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ReviewQuarantinedRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ReviewQuarantinedRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ReviewQuarantinedRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ReviewQuarantinedRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


