# MeterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Name** | Pointer to **string** | Human-readable name for the resource. Between 1 and 256 characters. Defaults to the slug if not specified. | [optional] 
**Slug** | **string** | A unique, human-readable identifier for the meter. Must consist only alphanumeric and underscore characters. | 
**Aggregation** | [**MeterAggregation**](MeterAggregation.md) | The aggregation type to use for the meter. | 
**EventType** | **string** | The event type to aggregate. | 
**EventFrom** | Pointer to **time.Time** | The date since the meter should include events. Useful to skip old events. If not specified, all historical events are included. | [optional] 
**ValueProperty** | Pointer to **string** | JSONPath expression to extract the value from the ingested event&#39;s data property.  The ingested value for SUM, AVG, MIN, and MAX aggregations is a number or a string that can be parsed to a number.  For UNIQUE_COUNT aggregation, the ingested value must be a string. For COUNT aggregation the valueProperty is ignored. | [optional] 
**GroupBy** | Pointer to **map[string]string** | Named JSONPath expressions to extract the group by values from the event data.  Keys must be unique and consist only alphanumeric and underscore characters. | [optional] 

## Methods

### NewMeterCreate

`func NewMeterCreate(slug string, aggregation MeterAggregation, eventType string, ) *MeterCreate`

NewMeterCreate instantiates a new MeterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterCreateWithDefaults

`func NewMeterCreateWithDefaults() *MeterCreate`

NewMeterCreateWithDefaults instantiates a new MeterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MeterCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeterCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeterCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MeterCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *MeterCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MeterCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MeterCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MeterCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *MeterCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeterCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeterCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeterCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *MeterCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MeterCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MeterCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetAggregation

`func (o *MeterCreate) GetAggregation() MeterAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MeterCreate) GetAggregationOk() (*MeterAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MeterCreate) SetAggregation(v MeterAggregation)`

SetAggregation sets Aggregation field to given value.


### GetEventType

`func (o *MeterCreate) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MeterCreate) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MeterCreate) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventFrom

`func (o *MeterCreate) GetEventFrom() time.Time`

GetEventFrom returns the EventFrom field if non-nil, zero value otherwise.

### GetEventFromOk

`func (o *MeterCreate) GetEventFromOk() (*time.Time, bool)`

GetEventFromOk returns a tuple with the EventFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFrom

`func (o *MeterCreate) SetEventFrom(v time.Time)`

SetEventFrom sets EventFrom field to given value.

### HasEventFrom

`func (o *MeterCreate) HasEventFrom() bool`

HasEventFrom returns a boolean if a field has been set.

### GetValueProperty

`func (o *MeterCreate) GetValueProperty() string`

GetValueProperty returns the ValueProperty field if non-nil, zero value otherwise.

### GetValuePropertyOk

`func (o *MeterCreate) GetValuePropertyOk() (*string, bool)`

GetValuePropertyOk returns a tuple with the ValueProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueProperty

`func (o *MeterCreate) SetValueProperty(v string)`

SetValueProperty sets ValueProperty field to given value.

### HasValueProperty

`func (o *MeterCreate) HasValueProperty() bool`

HasValueProperty returns a boolean if a field has been set.

### GetGroupBy

`func (o *MeterCreate) GetGroupBy() map[string]string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MeterCreate) GetGroupByOk() (*map[string]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MeterCreate) SetGroupBy(v map[string]string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MeterCreate) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


