# SalesforceConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** |  | 
**SyncDirection** | **string** |  | 
**SyncFrequency** | **string** |  | 
**Owner** | Pointer to **string** |  | [optional] 

## Methods

### NewSalesforceConnectRequest

`func NewSalesforceConnectRequest(environment string, syncDirection string, syncFrequency string, ) *SalesforceConnectRequest`

NewSalesforceConnectRequest instantiates a new SalesforceConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceConnectRequestWithDefaults

`func NewSalesforceConnectRequestWithDefaults() *SalesforceConnectRequest`

NewSalesforceConnectRequestWithDefaults instantiates a new SalesforceConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *SalesforceConnectRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SalesforceConnectRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SalesforceConnectRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetSyncDirection

`func (o *SalesforceConnectRequest) GetSyncDirection() string`

GetSyncDirection returns the SyncDirection field if non-nil, zero value otherwise.

### GetSyncDirectionOk

`func (o *SalesforceConnectRequest) GetSyncDirectionOk() (*string, bool)`

GetSyncDirectionOk returns a tuple with the SyncDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDirection

`func (o *SalesforceConnectRequest) SetSyncDirection(v string)`

SetSyncDirection sets SyncDirection field to given value.


### GetSyncFrequency

`func (o *SalesforceConnectRequest) GetSyncFrequency() string`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *SalesforceConnectRequest) GetSyncFrequencyOk() (*string, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *SalesforceConnectRequest) SetSyncFrequency(v string)`

SetSyncFrequency sets SyncFrequency field to given value.


### GetOwner

`func (o *SalesforceConnectRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SalesforceConnectRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SalesforceConnectRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SalesforceConnectRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


