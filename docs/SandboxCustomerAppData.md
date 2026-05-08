# SandboxCustomerAppData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**SandboxApp**](SandboxApp.md) | The installed sandbox app this data belongs to. | [optional] [readonly] 
**Id** | Pointer to **string** | The app ID. If not provided, it will use the global default for the app type. | [optional] 
**Type** | **string** | The app name. | 

## Methods

### NewSandboxCustomerAppData

`func NewSandboxCustomerAppData(type_ string, ) *SandboxCustomerAppData`

NewSandboxCustomerAppData instantiates a new SandboxCustomerAppData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCustomerAppDataWithDefaults

`func NewSandboxCustomerAppDataWithDefaults() *SandboxCustomerAppData`

NewSandboxCustomerAppDataWithDefaults instantiates a new SandboxCustomerAppData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *SandboxCustomerAppData) GetApp() SandboxApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SandboxCustomerAppData) GetAppOk() (*SandboxApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SandboxCustomerAppData) SetApp(v SandboxApp)`

SetApp sets App field to given value.

### HasApp

`func (o *SandboxCustomerAppData) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *SandboxCustomerAppData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxCustomerAppData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxCustomerAppData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxCustomerAppData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SandboxCustomerAppData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SandboxCustomerAppData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SandboxCustomerAppData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


