# PlanRefPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Version** | **int32** |  | 

## Methods

### NewPlanRefPayload

`func NewPlanRefPayload(id string, version int32, ) *PlanRefPayload`

NewPlanRefPayload instantiates a new PlanRefPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanRefPayloadWithDefaults

`func NewPlanRefPayloadWithDefaults() *PlanRefPayload`

NewPlanRefPayloadWithDefaults instantiates a new PlanRefPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlanRefPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanRefPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanRefPayload) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *PlanRefPayload) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlanRefPayload) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlanRefPayload) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PlanRefPayload) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetVersion

`func (o *PlanRefPayload) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PlanRefPayload) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PlanRefPayload) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


