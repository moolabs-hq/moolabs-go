# ConnectorHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**ConnectorType** | **string** |  | 
**Status** | **string** |  | 
**LastSyncedAt** | **time.Time** |  | 
**EventsIngestedTotal** | **int32** |  | 
**EventsLast24h** | **int32** |  | 
**ErrorsLast24h** | **int32** |  | 
**IsAggregateOnly** | **bool** |  | 

## Methods

### NewConnectorHealthResponse

`func NewConnectorHealthResponse(tenantId string, connectorType string, status string, lastSyncedAt time.Time, eventsIngestedTotal int32, eventsLast24h int32, errorsLast24h int32, isAggregateOnly bool, ) *ConnectorHealthResponse`

NewConnectorHealthResponse instantiates a new ConnectorHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorHealthResponseWithDefaults

`func NewConnectorHealthResponseWithDefaults() *ConnectorHealthResponse`

NewConnectorHealthResponseWithDefaults instantiates a new ConnectorHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ConnectorHealthResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ConnectorHealthResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ConnectorHealthResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetConnectorType

`func (o *ConnectorHealthResponse) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *ConnectorHealthResponse) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *ConnectorHealthResponse) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.


### GetStatus

`func (o *ConnectorHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSyncedAt

`func (o *ConnectorHealthResponse) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *ConnectorHealthResponse) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *ConnectorHealthResponse) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.


### GetEventsIngestedTotal

`func (o *ConnectorHealthResponse) GetEventsIngestedTotal() int32`

GetEventsIngestedTotal returns the EventsIngestedTotal field if non-nil, zero value otherwise.

### GetEventsIngestedTotalOk

`func (o *ConnectorHealthResponse) GetEventsIngestedTotalOk() (*int32, bool)`

GetEventsIngestedTotalOk returns a tuple with the EventsIngestedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsIngestedTotal

`func (o *ConnectorHealthResponse) SetEventsIngestedTotal(v int32)`

SetEventsIngestedTotal sets EventsIngestedTotal field to given value.


### GetEventsLast24h

`func (o *ConnectorHealthResponse) GetEventsLast24h() int32`

GetEventsLast24h returns the EventsLast24h field if non-nil, zero value otherwise.

### GetEventsLast24hOk

`func (o *ConnectorHealthResponse) GetEventsLast24hOk() (*int32, bool)`

GetEventsLast24hOk returns a tuple with the EventsLast24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsLast24h

`func (o *ConnectorHealthResponse) SetEventsLast24h(v int32)`

SetEventsLast24h sets EventsLast24h field to given value.


### GetErrorsLast24h

`func (o *ConnectorHealthResponse) GetErrorsLast24h() int32`

GetErrorsLast24h returns the ErrorsLast24h field if non-nil, zero value otherwise.

### GetErrorsLast24hOk

`func (o *ConnectorHealthResponse) GetErrorsLast24hOk() (*int32, bool)`

GetErrorsLast24hOk returns a tuple with the ErrorsLast24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLast24h

`func (o *ConnectorHealthResponse) SetErrorsLast24h(v int32)`

SetErrorsLast24h sets ErrorsLast24h field to given value.


### GetIsAggregateOnly

`func (o *ConnectorHealthResponse) GetIsAggregateOnly() bool`

GetIsAggregateOnly returns the IsAggregateOnly field if non-nil, zero value otherwise.

### GetIsAggregateOnlyOk

`func (o *ConnectorHealthResponse) GetIsAggregateOnlyOk() (*bool, bool)`

GetIsAggregateOnlyOk returns a tuple with the IsAggregateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAggregateOnly

`func (o *ConnectorHealthResponse) SetIsAggregateOnly(v bool)`

SetIsAggregateOnly sets IsAggregateOnly field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


