# IngestEventsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifies the event. | 
**Source** | **string** | Identifies the context in which an event happened. | 
**Specversion** | **string** | The version of the CloudEvents specification which the event uses. | [default to "1.0"]
**Type** | **string** | Contains a value describing the type of event related to the originating occurrence. | 
**Datacontenttype** | Pointer to **string** | Content type of the CloudEvents data value. Only the value \&quot;application/json\&quot; is allowed over HTTP. | [optional] 
**Dataschema** | Pointer to **string** | Identifies the schema that data adheres to. | [optional] 
**Subject** | **string** | Describes the subject of the event in the context of the event producer (identified by source). | 
**Time** | Pointer to **time.Time** | Timestamp of when the occurrence happened. Must adhere to RFC 3339. | [optional] 
**Data** | Pointer to **map[string]interface{}** | The event payload. Optional, if present it must be a JSON object. | [optional] 

## Methods

### NewIngestEventsBody

`func NewIngestEventsBody(id string, source string, specversion string, type_ string, subject string, ) *IngestEventsBody`

NewIngestEventsBody instantiates a new IngestEventsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestEventsBodyWithDefaults

`func NewIngestEventsBodyWithDefaults() *IngestEventsBody`

NewIngestEventsBodyWithDefaults instantiates a new IngestEventsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestEventsBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestEventsBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestEventsBody) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *IngestEventsBody) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IngestEventsBody) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IngestEventsBody) SetSource(v string)`

SetSource sets Source field to given value.


### GetSpecversion

`func (o *IngestEventsBody) GetSpecversion() string`

GetSpecversion returns the Specversion field if non-nil, zero value otherwise.

### GetSpecversionOk

`func (o *IngestEventsBody) GetSpecversionOk() (*string, bool)`

GetSpecversionOk returns a tuple with the Specversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecversion

`func (o *IngestEventsBody) SetSpecversion(v string)`

SetSpecversion sets Specversion field to given value.


### GetType

`func (o *IngestEventsBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestEventsBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestEventsBody) SetType(v string)`

SetType sets Type field to given value.


### GetDatacontenttype

`func (o *IngestEventsBody) GetDatacontenttype() string`

GetDatacontenttype returns the Datacontenttype field if non-nil, zero value otherwise.

### GetDatacontenttypeOk

`func (o *IngestEventsBody) GetDatacontenttypeOk() (*string, bool)`

GetDatacontenttypeOk returns a tuple with the Datacontenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacontenttype

`func (o *IngestEventsBody) SetDatacontenttype(v string)`

SetDatacontenttype sets Datacontenttype field to given value.

### HasDatacontenttype

`func (o *IngestEventsBody) HasDatacontenttype() bool`

HasDatacontenttype returns a boolean if a field has been set.

### GetDataschema

`func (o *IngestEventsBody) GetDataschema() string`

GetDataschema returns the Dataschema field if non-nil, zero value otherwise.

### GetDataschemaOk

`func (o *IngestEventsBody) GetDataschemaOk() (*string, bool)`

GetDataschemaOk returns a tuple with the Dataschema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataschema

`func (o *IngestEventsBody) SetDataschema(v string)`

SetDataschema sets Dataschema field to given value.

### HasDataschema

`func (o *IngestEventsBody) HasDataschema() bool`

HasDataschema returns a boolean if a field has been set.

### GetSubject

`func (o *IngestEventsBody) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IngestEventsBody) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IngestEventsBody) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTime

`func (o *IngestEventsBody) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IngestEventsBody) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IngestEventsBody) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *IngestEventsBody) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetData

`func (o *IngestEventsBody) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IngestEventsBody) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IngestEventsBody) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *IngestEventsBody) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


