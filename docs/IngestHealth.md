# IngestHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsLastHour** | **int32** |  | 
**EventsLast24h** | **int32** |  | 
**ConsumerLagEstimate** | **int32** |  | 
**CacheHitRatio** | **float32** |  | 
**DlqCount24h** | **int32** |  | 
**Status** | **string** |  | 

## Methods

### NewIngestHealth

`func NewIngestHealth(eventsLastHour int32, eventsLast24h int32, consumerLagEstimate int32, cacheHitRatio float32, dlqCount24h int32, status string, ) *IngestHealth`

NewIngestHealth instantiates a new IngestHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestHealthWithDefaults

`func NewIngestHealthWithDefaults() *IngestHealth`

NewIngestHealthWithDefaults instantiates a new IngestHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsLastHour

`func (o *IngestHealth) GetEventsLastHour() int32`

GetEventsLastHour returns the EventsLastHour field if non-nil, zero value otherwise.

### GetEventsLastHourOk

`func (o *IngestHealth) GetEventsLastHourOk() (*int32, bool)`

GetEventsLastHourOk returns a tuple with the EventsLastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsLastHour

`func (o *IngestHealth) SetEventsLastHour(v int32)`

SetEventsLastHour sets EventsLastHour field to given value.


### GetEventsLast24h

`func (o *IngestHealth) GetEventsLast24h() int32`

GetEventsLast24h returns the EventsLast24h field if non-nil, zero value otherwise.

### GetEventsLast24hOk

`func (o *IngestHealth) GetEventsLast24hOk() (*int32, bool)`

GetEventsLast24hOk returns a tuple with the EventsLast24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsLast24h

`func (o *IngestHealth) SetEventsLast24h(v int32)`

SetEventsLast24h sets EventsLast24h field to given value.


### GetConsumerLagEstimate

`func (o *IngestHealth) GetConsumerLagEstimate() int32`

GetConsumerLagEstimate returns the ConsumerLagEstimate field if non-nil, zero value otherwise.

### GetConsumerLagEstimateOk

`func (o *IngestHealth) GetConsumerLagEstimateOk() (*int32, bool)`

GetConsumerLagEstimateOk returns a tuple with the ConsumerLagEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerLagEstimate

`func (o *IngestHealth) SetConsumerLagEstimate(v int32)`

SetConsumerLagEstimate sets ConsumerLagEstimate field to given value.


### GetCacheHitRatio

`func (o *IngestHealth) GetCacheHitRatio() float32`

GetCacheHitRatio returns the CacheHitRatio field if non-nil, zero value otherwise.

### GetCacheHitRatioOk

`func (o *IngestHealth) GetCacheHitRatioOk() (*float32, bool)`

GetCacheHitRatioOk returns a tuple with the CacheHitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheHitRatio

`func (o *IngestHealth) SetCacheHitRatio(v float32)`

SetCacheHitRatio sets CacheHitRatio field to given value.


### GetDlqCount24h

`func (o *IngestHealth) GetDlqCount24h() int32`

GetDlqCount24h returns the DlqCount24h field if non-nil, zero value otherwise.

### GetDlqCount24hOk

`func (o *IngestHealth) GetDlqCount24hOk() (*int32, bool)`

GetDlqCount24hOk returns a tuple with the DlqCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlqCount24h

`func (o *IngestHealth) SetDlqCount24h(v int32)`

SetDlqCount24h sets DlqCount24h field to given value.


### GetStatus

`func (o *IngestHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IngestHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IngestHealth) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


