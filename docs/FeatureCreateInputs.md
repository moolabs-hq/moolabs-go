# FeatureCreateInputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A key is a unique string that is used to identify a resource. | 
**Name** | **string** |  | 
**Metadata** | Pointer to **map[string]string** | Set of key-value pairs. Metadata can be used to store additional information about a resource. | [optional] 
**MeterSlug** | Pointer to **string** | A key is a unique string that is used to identify a resource. | [optional] 
**MeterGroupByFilters** | Pointer to **map[string]string** | Optional meter group by filters. Useful if the meter scope is broader than what feature tracks. Example scenario would be a meter tracking all token use with groupBy fields for the model, then the feature could filter for model&#x3D;gpt-4. | [optional] 
**AdvancedMeterGroupByFilters** | Pointer to [**map[string]FilterString**](FilterString.md) | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. | [optional] 

## Methods

### NewFeatureCreateInputs

`func NewFeatureCreateInputs(key string, name string, ) *FeatureCreateInputs`

NewFeatureCreateInputs instantiates a new FeatureCreateInputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureCreateInputsWithDefaults

`func NewFeatureCreateInputsWithDefaults() *FeatureCreateInputs`

NewFeatureCreateInputsWithDefaults instantiates a new FeatureCreateInputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FeatureCreateInputs) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureCreateInputs) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureCreateInputs) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *FeatureCreateInputs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureCreateInputs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureCreateInputs) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *FeatureCreateInputs) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeatureCreateInputs) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeatureCreateInputs) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeatureCreateInputs) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterSlug

`func (o *FeatureCreateInputs) GetMeterSlug() string`

GetMeterSlug returns the MeterSlug field if non-nil, zero value otherwise.

### GetMeterSlugOk

`func (o *FeatureCreateInputs) GetMeterSlugOk() (*string, bool)`

GetMeterSlugOk returns a tuple with the MeterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSlug

`func (o *FeatureCreateInputs) SetMeterSlug(v string)`

SetMeterSlug sets MeterSlug field to given value.

### HasMeterSlug

`func (o *FeatureCreateInputs) HasMeterSlug() bool`

HasMeterSlug returns a boolean if a field has been set.

### GetMeterGroupByFilters

`func (o *FeatureCreateInputs) GetMeterGroupByFilters() map[string]string`

GetMeterGroupByFilters returns the MeterGroupByFilters field if non-nil, zero value otherwise.

### GetMeterGroupByFiltersOk

`func (o *FeatureCreateInputs) GetMeterGroupByFiltersOk() (*map[string]string, bool)`

GetMeterGroupByFiltersOk returns a tuple with the MeterGroupByFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterGroupByFilters

`func (o *FeatureCreateInputs) SetMeterGroupByFilters(v map[string]string)`

SetMeterGroupByFilters sets MeterGroupByFilters field to given value.

### HasMeterGroupByFilters

`func (o *FeatureCreateInputs) HasMeterGroupByFilters() bool`

HasMeterGroupByFilters returns a boolean if a field has been set.

### GetAdvancedMeterGroupByFilters

`func (o *FeatureCreateInputs) GetAdvancedMeterGroupByFilters() map[string]FilterString`

GetAdvancedMeterGroupByFilters returns the AdvancedMeterGroupByFilters field if non-nil, zero value otherwise.

### GetAdvancedMeterGroupByFiltersOk

`func (o *FeatureCreateInputs) GetAdvancedMeterGroupByFiltersOk() (*map[string]FilterString, bool)`

GetAdvancedMeterGroupByFiltersOk returns a tuple with the AdvancedMeterGroupByFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedMeterGroupByFilters

`func (o *FeatureCreateInputs) SetAdvancedMeterGroupByFilters(v map[string]FilterString)`

SetAdvancedMeterGroupByFilters sets AdvancedMeterGroupByFilters field to given value.

### HasAdvancedMeterGroupByFilters

`func (o *FeatureCreateInputs) HasAdvancedMeterGroupByFilters() bool`

HasAdvancedMeterGroupByFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


