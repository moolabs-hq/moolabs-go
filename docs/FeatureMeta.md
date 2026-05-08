# FeatureMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of a feature. | 
**Key** | **string** | The key is an immutable unique identifier of the feature used throughout the API, for example when interacting with a subject&#39;s entitlements. | 

## Methods

### NewFeatureMeta

`func NewFeatureMeta(id string, key string, ) *FeatureMeta`

NewFeatureMeta instantiates a new FeatureMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureMetaWithDefaults

`func NewFeatureMetaWithDefaults() *FeatureMeta`

NewFeatureMetaWithDefaults instantiates a new FeatureMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureMeta) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *FeatureMeta) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureMeta) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureMeta) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


