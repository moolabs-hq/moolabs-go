# ClausePack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackId** | **string** |  | 
**Version** | **int32** |  | 
**LifecycleState** | **string** |  | 
**IndustryVertical** | **string** |  | 
**AppliesToContractTypes** | **[]string** |  | 
**Provenance** | [**PackProvenance**](PackProvenance.md) |  | 
**ClauseFamilies** | [**[]ClauseFamily**](ClauseFamily.md) |  | 

## Methods

### NewClausePack

`func NewClausePack(packId string, version int32, lifecycleState string, industryVertical string, appliesToContractTypes []string, provenance PackProvenance, clauseFamilies []ClauseFamily, ) *ClausePack`

NewClausePack instantiates a new ClausePack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClausePackWithDefaults

`func NewClausePackWithDefaults() *ClausePack`

NewClausePackWithDefaults instantiates a new ClausePack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackId

`func (o *ClausePack) GetPackId() string`

GetPackId returns the PackId field if non-nil, zero value otherwise.

### GetPackIdOk

`func (o *ClausePack) GetPackIdOk() (*string, bool)`

GetPackIdOk returns a tuple with the PackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackId

`func (o *ClausePack) SetPackId(v string)`

SetPackId sets PackId field to given value.


### GetVersion

`func (o *ClausePack) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClausePack) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClausePack) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetLifecycleState

`func (o *ClausePack) GetLifecycleState() string`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *ClausePack) GetLifecycleStateOk() (*string, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *ClausePack) SetLifecycleState(v string)`

SetLifecycleState sets LifecycleState field to given value.


### GetIndustryVertical

`func (o *ClausePack) GetIndustryVertical() string`

GetIndustryVertical returns the IndustryVertical field if non-nil, zero value otherwise.

### GetIndustryVerticalOk

`func (o *ClausePack) GetIndustryVerticalOk() (*string, bool)`

GetIndustryVerticalOk returns a tuple with the IndustryVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryVertical

`func (o *ClausePack) SetIndustryVertical(v string)`

SetIndustryVertical sets IndustryVertical field to given value.


### GetAppliesToContractTypes

`func (o *ClausePack) GetAppliesToContractTypes() []string`

GetAppliesToContractTypes returns the AppliesToContractTypes field if non-nil, zero value otherwise.

### GetAppliesToContractTypesOk

`func (o *ClausePack) GetAppliesToContractTypesOk() (*[]string, bool)`

GetAppliesToContractTypesOk returns a tuple with the AppliesToContractTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToContractTypes

`func (o *ClausePack) SetAppliesToContractTypes(v []string)`

SetAppliesToContractTypes sets AppliesToContractTypes field to given value.


### GetProvenance

`func (o *ClausePack) GetProvenance() PackProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *ClausePack) GetProvenanceOk() (*PackProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *ClausePack) SetProvenance(v PackProvenance)`

SetProvenance sets Provenance field to given value.


### GetClauseFamilies

`func (o *ClausePack) GetClauseFamilies() []ClauseFamily`

GetClauseFamilies returns the ClauseFamilies field if non-nil, zero value otherwise.

### GetClauseFamiliesOk

`func (o *ClausePack) GetClauseFamiliesOk() (*[]ClauseFamily, bool)`

GetClauseFamiliesOk returns a tuple with the ClauseFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseFamilies

`func (o *ClausePack) SetClauseFamilies(v []ClauseFamily)`

SetClauseFamilies sets ClauseFamilies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


