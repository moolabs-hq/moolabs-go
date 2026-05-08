# PlanReferenceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The plan key. | 
**Version** | Pointer to **int32** | The plan version. | [optional] 

## Methods

### NewPlanReferenceInput

`func NewPlanReferenceInput(key string, ) *PlanReferenceInput`

NewPlanReferenceInput instantiates a new PlanReferenceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanReferenceInputWithDefaults

`func NewPlanReferenceInputWithDefaults() *PlanReferenceInput`

NewPlanReferenceInputWithDefaults instantiates a new PlanReferenceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlanReferenceInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlanReferenceInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlanReferenceInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *PlanReferenceInput) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlanReferenceInput) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlanReferenceInput) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PlanReferenceInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


