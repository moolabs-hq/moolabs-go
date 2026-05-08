# Meter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable name for the resource. Between 1 and 256 characters. Defaults to the slug if not specified. | [optional] 
**ProjectionStatus** | Pointer to **string** | Async ClickHouse parser reconciliation status for this meter. | [optional] [readonly] 
**ProjectionError** | Pointer to **string** | Last projection reconciliation error, if projectionStatus is failed. | [optional] [readonly] 
**DataHealth** | Pointer to **string** | Raw-event extraction health for this meter. | [optional] [readonly] 
**DataHealthReason** | Pointer to **string** | Last data-health transition reason. | [optional] [readonly] 
**Slug** | **string** | A unique, human-readable identifier for the meter. Must consist only alphanumeric and underscore characters. | 
**Aggregation** | [**MeterAggregation**](MeterAggregation.md) | The aggregation type to use for the meter. | 
**EventType** | **string** | The event type to aggregate. | 
**EventFrom** | Pointer to **time.Time** | The date since the meter should include events. Useful to skip old events. If not specified, all historical events are included. | [optional] 
**ValueProperty** | Pointer to **string** | JSONPath expression to extract the value from the ingested event&#39;s data property.  The ingested value for SUM, AVG, MIN, and MAX aggregations is a number or a string that can be parsed to a number.  For UNIQUE_COUNT aggregation, the ingested value must be a string. For COUNT aggregation the valueProperty is ignored. | [optional] 
**GroupBy** | Pointer to **map[string]string** | Named JSONPath expressions to extract the group by values from the event data.  Keys must be unique and consist only alphanumeric and underscore characters. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 

## Methods

### NewMeter

`func NewMeter(id string, createdAt time.Time, updatedAt time.Time, slug string, aggregation MeterAggregation, eventType string, ) *Meter`

NewMeter instantiates a new Meter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterWithDefaults

`func NewMeterWithDefaults() *Meter`

NewMeterWithDefaults instantiates a new Meter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Meter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Meter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Meter) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *Meter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Meter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Meter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Meter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Meter) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Meter) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Meter) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Meter) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Meter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Meter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Meter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Meter) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Meter) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Meter) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Meter) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Meter) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Meter) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Meter) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *Meter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Meter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Meter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Meter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectionStatus

`func (o *Meter) GetProjectionStatus() string`

GetProjectionStatus returns the ProjectionStatus field if non-nil, zero value otherwise.

### GetProjectionStatusOk

`func (o *Meter) GetProjectionStatusOk() (*string, bool)`

GetProjectionStatusOk returns a tuple with the ProjectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionStatus

`func (o *Meter) SetProjectionStatus(v string)`

SetProjectionStatus sets ProjectionStatus field to given value.

### HasProjectionStatus

`func (o *Meter) HasProjectionStatus() bool`

HasProjectionStatus returns a boolean if a field has been set.

### GetProjectionError

`func (o *Meter) GetProjectionError() string`

GetProjectionError returns the ProjectionError field if non-nil, zero value otherwise.

### GetProjectionErrorOk

`func (o *Meter) GetProjectionErrorOk() (*string, bool)`

GetProjectionErrorOk returns a tuple with the ProjectionError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionError

`func (o *Meter) SetProjectionError(v string)`

SetProjectionError sets ProjectionError field to given value.

### HasProjectionError

`func (o *Meter) HasProjectionError() bool`

HasProjectionError returns a boolean if a field has been set.

### GetDataHealth

`func (o *Meter) GetDataHealth() string`

GetDataHealth returns the DataHealth field if non-nil, zero value otherwise.

### GetDataHealthOk

`func (o *Meter) GetDataHealthOk() (*string, bool)`

GetDataHealthOk returns a tuple with the DataHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHealth

`func (o *Meter) SetDataHealth(v string)`

SetDataHealth sets DataHealth field to given value.

### HasDataHealth

`func (o *Meter) HasDataHealth() bool`

HasDataHealth returns a boolean if a field has been set.

### GetDataHealthReason

`func (o *Meter) GetDataHealthReason() string`

GetDataHealthReason returns the DataHealthReason field if non-nil, zero value otherwise.

### GetDataHealthReasonOk

`func (o *Meter) GetDataHealthReasonOk() (*string, bool)`

GetDataHealthReasonOk returns a tuple with the DataHealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHealthReason

`func (o *Meter) SetDataHealthReason(v string)`

SetDataHealthReason sets DataHealthReason field to given value.

### HasDataHealthReason

`func (o *Meter) HasDataHealthReason() bool`

HasDataHealthReason returns a boolean if a field has been set.

### GetSlug

`func (o *Meter) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Meter) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Meter) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetAggregation

`func (o *Meter) GetAggregation() MeterAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *Meter) GetAggregationOk() (*MeterAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *Meter) SetAggregation(v MeterAggregation)`

SetAggregation sets Aggregation field to given value.


### GetEventType

`func (o *Meter) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Meter) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Meter) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventFrom

`func (o *Meter) GetEventFrom() time.Time`

GetEventFrom returns the EventFrom field if non-nil, zero value otherwise.

### GetEventFromOk

`func (o *Meter) GetEventFromOk() (*time.Time, bool)`

GetEventFromOk returns a tuple with the EventFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFrom

`func (o *Meter) SetEventFrom(v time.Time)`

SetEventFrom sets EventFrom field to given value.

### HasEventFrom

`func (o *Meter) HasEventFrom() bool`

HasEventFrom returns a boolean if a field has been set.

### GetValueProperty

`func (o *Meter) GetValueProperty() string`

GetValueProperty returns the ValueProperty field if non-nil, zero value otherwise.

### GetValuePropertyOk

`func (o *Meter) GetValuePropertyOk() (*string, bool)`

GetValuePropertyOk returns a tuple with the ValueProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueProperty

`func (o *Meter) SetValueProperty(v string)`

SetValueProperty sets ValueProperty field to given value.

### HasValueProperty

`func (o *Meter) HasValueProperty() bool`

HasValueProperty returns a boolean if a field has been set.

### GetGroupBy

`func (o *Meter) GetGroupBy() map[string]string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Meter) GetGroupByOk() (*map[string]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Meter) SetGroupBy(v map[string]string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Meter) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetAnnotations

`func (o *Meter) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Meter) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Meter) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Meter) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


