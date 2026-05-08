# Addon1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the add-on. | 
**Key** | **string** | A semi-unique identifier for the resource. | [readonly] 
**Version** | **int32** | The version of the Add-on which templates this instance. | [readonly] [default to 1]
**InstanceType** | [**AddonInstanceType**](AddonInstanceType.md) | The instance type of the add-on. | [readonly] 

## Methods

### NewAddon1

`func NewAddon1(id string, key string, version int32, instanceType AddonInstanceType, ) *Addon1`

NewAddon1 instantiates a new Addon1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddon1WithDefaults

`func NewAddon1WithDefaults() *Addon1`

NewAddon1WithDefaults instantiates a new Addon1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Addon1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Addon1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Addon1) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *Addon1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Addon1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Addon1) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *Addon1) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Addon1) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Addon1) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetInstanceType

`func (o *Addon1) GetInstanceType() AddonInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Addon1) GetInstanceTypeOk() (*AddonInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Addon1) SetInstanceType(v AddonInstanceType)`

SetInstanceType sets InstanceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


