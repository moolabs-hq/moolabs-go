# MeterUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Name** | Pointer to **string** | Human-readable name for the resource. Between 1 and 256 characters. Defaults to the slug if not specified. | [optional] 
**GroupBy** | Pointer to **map[string]string** | Named JSONPath expressions to extract the group by values from the event data.  Keys must be unique and consist only alphanumeric and underscore characters. | [optional] 

## Methods

### NewMeterUpdate

`func NewMeterUpdate() *MeterUpdate`

NewMeterUpdate instantiates a new MeterUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterUpdateWithDefaults

`func NewMeterUpdateWithDefaults() *MeterUpdate`

NewMeterUpdateWithDefaults instantiates a new MeterUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MeterUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MeterUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MeterUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MeterUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *MeterUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MeterUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MeterUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MeterUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *MeterUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeterUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeterUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeterUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroupBy

`func (o *MeterUpdate) GetGroupBy() map[string]string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MeterUpdate) GetGroupByOk() (*map[string]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MeterUpdate) SetGroupBy(v map[string]string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MeterUpdate) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


