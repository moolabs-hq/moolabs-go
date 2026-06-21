# PackProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthoredBy** | Pointer to **string** |  | [optional] [default to "moolabs"]
**ForkedFrom** | Pointer to **string** |  | [optional] 
**Signoffs** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewPackProvenance

`func NewPackProvenance() *PackProvenance`

NewPackProvenance instantiates a new PackProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackProvenanceWithDefaults

`func NewPackProvenanceWithDefaults() *PackProvenance`

NewPackProvenanceWithDefaults instantiates a new PackProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthoredBy

`func (o *PackProvenance) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PackProvenance) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PackProvenance) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PackProvenance) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetForkedFrom

`func (o *PackProvenance) GetForkedFrom() string`

GetForkedFrom returns the ForkedFrom field if non-nil, zero value otherwise.

### GetForkedFromOk

`func (o *PackProvenance) GetForkedFromOk() (*string, bool)`

GetForkedFromOk returns a tuple with the ForkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkedFrom

`func (o *PackProvenance) SetForkedFrom(v string)`

SetForkedFrom sets ForkedFrom field to given value.

### HasForkedFrom

`func (o *PackProvenance) HasForkedFrom() bool`

HasForkedFrom returns a boolean if a field has been set.

### GetSignoffs

`func (o *PackProvenance) GetSignoffs() []map[string]interface{}`

GetSignoffs returns the Signoffs field if non-nil, zero value otherwise.

### GetSignoffsOk

`func (o *PackProvenance) GetSignoffsOk() (*[]map[string]interface{}, bool)`

GetSignoffsOk returns a tuple with the Signoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignoffs

`func (o *PackProvenance) SetSignoffs(v []map[string]interface{})`

SetSignoffs sets Signoffs field to given value.

### HasSignoffs

`func (o *PackProvenance) HasSignoffs() bool`

HasSignoffs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


