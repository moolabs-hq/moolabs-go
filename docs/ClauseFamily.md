# ClauseFamily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FamilyId** | **string** |  | 
**DisplayName** | **string** |  | 
**OwnerPersona** | **string** |  | 
**Required** | Pointer to **bool** |  | [optional] [default to false]
**Detector** | **map[string]interface{}** |  | 
**Positions** | [**map[string]ClausePosition**](ClausePosition.md) |  | 
**DealContext** | Pointer to [**DealContext**](DealContext.md) |  | [optional] 
**Citations** | Pointer to [**ClauseCitations**](ClauseCitations.md) |  | [optional] 

## Methods

### NewClauseFamily

`func NewClauseFamily(familyId string, displayName string, ownerPersona string, detector map[string]interface{}, positions map[string]ClausePosition, ) *ClauseFamily`

NewClauseFamily instantiates a new ClauseFamily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClauseFamilyWithDefaults

`func NewClauseFamilyWithDefaults() *ClauseFamily`

NewClauseFamilyWithDefaults instantiates a new ClauseFamily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamilyId

`func (o *ClauseFamily) GetFamilyId() string`

GetFamilyId returns the FamilyId field if non-nil, zero value otherwise.

### GetFamilyIdOk

`func (o *ClauseFamily) GetFamilyIdOk() (*string, bool)`

GetFamilyIdOk returns a tuple with the FamilyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyId

`func (o *ClauseFamily) SetFamilyId(v string)`

SetFamilyId sets FamilyId field to given value.


### GetDisplayName

`func (o *ClauseFamily) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClauseFamily) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClauseFamily) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetOwnerPersona

`func (o *ClauseFamily) GetOwnerPersona() string`

GetOwnerPersona returns the OwnerPersona field if non-nil, zero value otherwise.

### GetOwnerPersonaOk

`func (o *ClauseFamily) GetOwnerPersonaOk() (*string, bool)`

GetOwnerPersonaOk returns a tuple with the OwnerPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPersona

`func (o *ClauseFamily) SetOwnerPersona(v string)`

SetOwnerPersona sets OwnerPersona field to given value.


### GetRequired

`func (o *ClauseFamily) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ClauseFamily) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ClauseFamily) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ClauseFamily) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDetector

`func (o *ClauseFamily) GetDetector() map[string]interface{}`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *ClauseFamily) GetDetectorOk() (*map[string]interface{}, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *ClauseFamily) SetDetector(v map[string]interface{})`

SetDetector sets Detector field to given value.


### GetPositions

`func (o *ClauseFamily) GetPositions() map[string]ClausePosition`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *ClauseFamily) GetPositionsOk() (*map[string]ClausePosition, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *ClauseFamily) SetPositions(v map[string]ClausePosition)`

SetPositions sets Positions field to given value.


### GetDealContext

`func (o *ClauseFamily) GetDealContext() DealContext`

GetDealContext returns the DealContext field if non-nil, zero value otherwise.

### GetDealContextOk

`func (o *ClauseFamily) GetDealContextOk() (*DealContext, bool)`

GetDealContextOk returns a tuple with the DealContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealContext

`func (o *ClauseFamily) SetDealContext(v DealContext)`

SetDealContext sets DealContext field to given value.

### HasDealContext

`func (o *ClauseFamily) HasDealContext() bool`

HasDealContext returns a boolean if a field has been set.

### GetCitations

`func (o *ClauseFamily) GetCitations() ClauseCitations`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *ClauseFamily) GetCitationsOk() (*ClauseCitations, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *ClauseFamily) SetCitations(v ClauseCitations)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *ClauseFamily) HasCitations() bool`

HasCitations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


