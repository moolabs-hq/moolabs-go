# Feature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**ArchivedAt** | Pointer to **time.Time** | Timestamp of when the resource was archived. | [optional] [readonly] 
**Key** | **string** | A key is a unique string that is used to identify a resource. | 
**Name** | **string** |  | 
**Metadata** | Pointer to **map[string]string** | Set of key-value pairs. Metadata can be used to store additional information about a resource. | [optional] 
**MeterSlug** | Pointer to **string** | A key is a unique string that is used to identify a resource. | [optional] 
**MeterGroupByFilters** | Pointer to **map[string]string** | Optional meter group by filters. Useful if the meter scope is broader than what feature tracks. Example scenario would be a meter tracking all token use with groupBy fields for the model, then the feature could filter for model&#x3D;gpt-4. | [optional] 
**AdvancedMeterGroupByFilters** | Pointer to [**map[string]FilterString**](FilterString.md) | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. | [optional] 
**Id** | **string** | Readonly unique ULID identifier. | [readonly] 

## Methods

### NewFeature

`func NewFeature(createdAt time.Time, updatedAt time.Time, key string, name string, id string, ) *Feature`

NewFeature instantiates a new Feature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureWithDefaults

`func NewFeatureWithDefaults() *Feature`

NewFeatureWithDefaults instantiates a new Feature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Feature) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Feature) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Feature) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Feature) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Feature) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Feature) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Feature) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Feature) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Feature) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Feature) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Feature) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Feature) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Feature) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Feature) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetKey

`func (o *Feature) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Feature) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Feature) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Feature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Feature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Feature) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *Feature) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Feature) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Feature) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Feature) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeterSlug

`func (o *Feature) GetMeterSlug() string`

GetMeterSlug returns the MeterSlug field if non-nil, zero value otherwise.

### GetMeterSlugOk

`func (o *Feature) GetMeterSlugOk() (*string, bool)`

GetMeterSlugOk returns a tuple with the MeterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterSlug

`func (o *Feature) SetMeterSlug(v string)`

SetMeterSlug sets MeterSlug field to given value.

### HasMeterSlug

`func (o *Feature) HasMeterSlug() bool`

HasMeterSlug returns a boolean if a field has been set.

### GetMeterGroupByFilters

`func (o *Feature) GetMeterGroupByFilters() map[string]string`

GetMeterGroupByFilters returns the MeterGroupByFilters field if non-nil, zero value otherwise.

### GetMeterGroupByFiltersOk

`func (o *Feature) GetMeterGroupByFiltersOk() (*map[string]string, bool)`

GetMeterGroupByFiltersOk returns a tuple with the MeterGroupByFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterGroupByFilters

`func (o *Feature) SetMeterGroupByFilters(v map[string]string)`

SetMeterGroupByFilters sets MeterGroupByFilters field to given value.

### HasMeterGroupByFilters

`func (o *Feature) HasMeterGroupByFilters() bool`

HasMeterGroupByFilters returns a boolean if a field has been set.

### GetAdvancedMeterGroupByFilters

`func (o *Feature) GetAdvancedMeterGroupByFilters() map[string]FilterString`

GetAdvancedMeterGroupByFilters returns the AdvancedMeterGroupByFilters field if non-nil, zero value otherwise.

### GetAdvancedMeterGroupByFiltersOk

`func (o *Feature) GetAdvancedMeterGroupByFiltersOk() (*map[string]FilterString, bool)`

GetAdvancedMeterGroupByFiltersOk returns a tuple with the AdvancedMeterGroupByFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedMeterGroupByFilters

`func (o *Feature) SetAdvancedMeterGroupByFilters(v map[string]FilterString)`

SetAdvancedMeterGroupByFilters sets AdvancedMeterGroupByFilters field to given value.

### HasAdvancedMeterGroupByFilters

`func (o *Feature) HasAdvancedMeterGroupByFilters() bool`

HasAdvancedMeterGroupByFilters returns a boolean if a field has been set.

### GetId

`func (o *Feature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Feature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Feature) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


