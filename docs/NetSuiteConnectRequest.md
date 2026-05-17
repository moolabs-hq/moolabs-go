# NetSuiteConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**SyncDirection** | **string** |  | 
**SyncFrequency** | **string** |  | 
**Owner** | Pointer to **string** |  | [optional] 
**InvoicePdfRenderScriptId** | Pointer to **string** |  | [optional] 
**InvoicePdfRenderDeployId** | Pointer to **string** |  | [optional] 
**InvoicePdfFormId** | Pointer to **string** |  | [optional] 

## Methods

### NewNetSuiteConnectRequest

`func NewNetSuiteConnectRequest(accountId string, clientId string, clientSecret string, syncDirection string, syncFrequency string, ) *NetSuiteConnectRequest`

NewNetSuiteConnectRequest instantiates a new NetSuiteConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetSuiteConnectRequestWithDefaults

`func NewNetSuiteConnectRequestWithDefaults() *NetSuiteConnectRequest`

NewNetSuiteConnectRequestWithDefaults instantiates a new NetSuiteConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NetSuiteConnectRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NetSuiteConnectRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NetSuiteConnectRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetClientId

`func (o *NetSuiteConnectRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NetSuiteConnectRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NetSuiteConnectRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *NetSuiteConnectRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *NetSuiteConnectRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *NetSuiteConnectRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSyncDirection

`func (o *NetSuiteConnectRequest) GetSyncDirection() string`

GetSyncDirection returns the SyncDirection field if non-nil, zero value otherwise.

### GetSyncDirectionOk

`func (o *NetSuiteConnectRequest) GetSyncDirectionOk() (*string, bool)`

GetSyncDirectionOk returns a tuple with the SyncDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDirection

`func (o *NetSuiteConnectRequest) SetSyncDirection(v string)`

SetSyncDirection sets SyncDirection field to given value.


### GetSyncFrequency

`func (o *NetSuiteConnectRequest) GetSyncFrequency() string`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *NetSuiteConnectRequest) GetSyncFrequencyOk() (*string, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *NetSuiteConnectRequest) SetSyncFrequency(v string)`

SetSyncFrequency sets SyncFrequency field to given value.


### GetOwner

`func (o *NetSuiteConnectRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NetSuiteConnectRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NetSuiteConnectRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NetSuiteConnectRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetInvoicePdfRenderScriptId

`func (o *NetSuiteConnectRequest) GetInvoicePdfRenderScriptId() string`

GetInvoicePdfRenderScriptId returns the InvoicePdfRenderScriptId field if non-nil, zero value otherwise.

### GetInvoicePdfRenderScriptIdOk

`func (o *NetSuiteConnectRequest) GetInvoicePdfRenderScriptIdOk() (*string, bool)`

GetInvoicePdfRenderScriptIdOk returns a tuple with the InvoicePdfRenderScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfRenderScriptId

`func (o *NetSuiteConnectRequest) SetInvoicePdfRenderScriptId(v string)`

SetInvoicePdfRenderScriptId sets InvoicePdfRenderScriptId field to given value.

### HasInvoicePdfRenderScriptId

`func (o *NetSuiteConnectRequest) HasInvoicePdfRenderScriptId() bool`

HasInvoicePdfRenderScriptId returns a boolean if a field has been set.

### GetInvoicePdfRenderDeployId

`func (o *NetSuiteConnectRequest) GetInvoicePdfRenderDeployId() string`

GetInvoicePdfRenderDeployId returns the InvoicePdfRenderDeployId field if non-nil, zero value otherwise.

### GetInvoicePdfRenderDeployIdOk

`func (o *NetSuiteConnectRequest) GetInvoicePdfRenderDeployIdOk() (*string, bool)`

GetInvoicePdfRenderDeployIdOk returns a tuple with the InvoicePdfRenderDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfRenderDeployId

`func (o *NetSuiteConnectRequest) SetInvoicePdfRenderDeployId(v string)`

SetInvoicePdfRenderDeployId sets InvoicePdfRenderDeployId field to given value.

### HasInvoicePdfRenderDeployId

`func (o *NetSuiteConnectRequest) HasInvoicePdfRenderDeployId() bool`

HasInvoicePdfRenderDeployId returns a boolean if a field has been set.

### GetInvoicePdfFormId

`func (o *NetSuiteConnectRequest) GetInvoicePdfFormId() string`

GetInvoicePdfFormId returns the InvoicePdfFormId field if non-nil, zero value otherwise.

### GetInvoicePdfFormIdOk

`func (o *NetSuiteConnectRequest) GetInvoicePdfFormIdOk() (*string, bool)`

GetInvoicePdfFormIdOk returns a tuple with the InvoicePdfFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfFormId

`func (o *NetSuiteConnectRequest) SetInvoicePdfFormId(v string)`

SetInvoicePdfFormId sets InvoicePdfFormId field to given value.

### HasInvoicePdfFormId

`func (o *NetSuiteConnectRequest) HasInvoicePdfFormId() bool`

HasInvoicePdfFormId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


