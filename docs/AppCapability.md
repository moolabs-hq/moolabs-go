# AppCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AppCapabilityType**](AppCapabilityType.md) | The capability type. | 
**Key** | **string** | Key | 
**Name** | **string** | The capability name. | 
**Description** | **string** | The capability description. | 

## Methods

### NewAppCapability

`func NewAppCapability(type_ AppCapabilityType, key string, name string, description string, ) *AppCapability`

NewAppCapability instantiates a new AppCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCapabilityWithDefaults

`func NewAppCapabilityWithDefaults() *AppCapability`

NewAppCapabilityWithDefaults instantiates a new AppCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppCapability) GetType() AppCapabilityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppCapability) GetTypeOk() (*AppCapabilityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppCapability) SetType(v AppCapabilityType)`

SetType sets Type field to given value.


### GetKey

`func (o *AppCapability) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppCapability) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppCapability) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *AppCapability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppCapability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppCapability) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppCapability) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppCapability) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppCapability) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


