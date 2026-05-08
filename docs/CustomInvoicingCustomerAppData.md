# CustomInvoicingCustomerAppData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**CustomInvoicingApp**](CustomInvoicingApp.md) | The installed custom invoicing app this data belongs to. | [optional] [readonly] 
**Id** | Pointer to **string** | The app ID. If not provided, it will use the global default for the app type. | [optional] 
**Type** | **string** | The app name. | 
**Metadata** | Pointer to **map[string]string** | Metadata to be used by the custom invoicing provider. | [optional] 

## Methods

### NewCustomInvoicingCustomerAppData

`func NewCustomInvoicingCustomerAppData(type_ string, ) *CustomInvoicingCustomerAppData`

NewCustomInvoicingCustomerAppData instantiates a new CustomInvoicingCustomerAppData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomInvoicingCustomerAppDataWithDefaults

`func NewCustomInvoicingCustomerAppDataWithDefaults() *CustomInvoicingCustomerAppData`

NewCustomInvoicingCustomerAppDataWithDefaults instantiates a new CustomInvoicingCustomerAppData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *CustomInvoicingCustomerAppData) GetApp() CustomInvoicingApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CustomInvoicingCustomerAppData) GetAppOk() (*CustomInvoicingApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CustomInvoicingCustomerAppData) SetApp(v CustomInvoicingApp)`

SetApp sets App field to given value.

### HasApp

`func (o *CustomInvoicingCustomerAppData) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *CustomInvoicingCustomerAppData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomInvoicingCustomerAppData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomInvoicingCustomerAppData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomInvoicingCustomerAppData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomInvoicingCustomerAppData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomInvoicingCustomerAppData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomInvoicingCustomerAppData) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *CustomInvoicingCustomerAppData) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomInvoicingCustomerAppData) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomInvoicingCustomerAppData) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomInvoicingCustomerAppData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


