# SDKEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageEventId** | **string** |  | 
**RequestId** | **string** |  | 
**TenantId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**FeatureKey** | Pointer to **string** |  | [optional] 
**Spans** | [**[]SDKSpan**](SDKSpan.md) |  | 
**ExpectedSpanCount** | Pointer to **int32** | Informational only — the server does NOT reject the event when len(spans) !&#x3D; expectedSpanCount. Provided by the SDK for client-side reconciliation / observability; the contract is documented in the rev 3 sdk-cost-capability-acute-backing contracts doc §4.4 / CONTRACT-Q6 RESOLVED. | [optional] 
**Timestamp** | **time.Time** |  | 

## Methods

### NewSDKEvent

`func NewSDKEvent(usageEventId string, requestId string, spans []SDKSpan, timestamp time.Time, ) *SDKEvent`

NewSDKEvent instantiates a new SDKEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDKEventWithDefaults

`func NewSDKEventWithDefaults() *SDKEvent`

NewSDKEventWithDefaults instantiates a new SDKEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageEventId

`func (o *SDKEvent) GetUsageEventId() string`

GetUsageEventId returns the UsageEventId field if non-nil, zero value otherwise.

### GetUsageEventIdOk

`func (o *SDKEvent) GetUsageEventIdOk() (*string, bool)`

GetUsageEventIdOk returns a tuple with the UsageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventId

`func (o *SDKEvent) SetUsageEventId(v string)`

SetUsageEventId sets UsageEventId field to given value.


### GetRequestId

`func (o *SDKEvent) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SDKEvent) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SDKEvent) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTenantId

`func (o *SDKEvent) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SDKEvent) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SDKEvent) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SDKEvent) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetCustomerId

`func (o *SDKEvent) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SDKEvent) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SDKEvent) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SDKEvent) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFeatureKey

`func (o *SDKEvent) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *SDKEvent) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *SDKEvent) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *SDKEvent) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetSpans

`func (o *SDKEvent) GetSpans() []SDKSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *SDKEvent) GetSpansOk() (*[]SDKSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *SDKEvent) SetSpans(v []SDKSpan)`

SetSpans sets Spans field to given value.


### GetExpectedSpanCount

`func (o *SDKEvent) GetExpectedSpanCount() int32`

GetExpectedSpanCount returns the ExpectedSpanCount field if non-nil, zero value otherwise.

### GetExpectedSpanCountOk

`func (o *SDKEvent) GetExpectedSpanCountOk() (*int32, bool)`

GetExpectedSpanCountOk returns a tuple with the ExpectedSpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSpanCount

`func (o *SDKEvent) SetExpectedSpanCount(v int32)`

SetExpectedSpanCount sets ExpectedSpanCount field to given value.

### HasExpectedSpanCount

`func (o *SDKEvent) HasExpectedSpanCount() bool`

HasExpectedSpanCount returns a boolean if a field has been set.

### GetTimestamp

`func (o *SDKEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SDKEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SDKEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


