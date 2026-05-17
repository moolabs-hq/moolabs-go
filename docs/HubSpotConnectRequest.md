# HubSpotConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncDirection** | **string** |  | 
**SyncFrequency** | **string** |  | 
**Owner** | Pointer to **string** |  | [optional] 

## Methods

### NewHubSpotConnectRequest

`func NewHubSpotConnectRequest(syncDirection string, syncFrequency string, ) *HubSpotConnectRequest`

NewHubSpotConnectRequest instantiates a new HubSpotConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubSpotConnectRequestWithDefaults

`func NewHubSpotConnectRequestWithDefaults() *HubSpotConnectRequest`

NewHubSpotConnectRequestWithDefaults instantiates a new HubSpotConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncDirection

`func (o *HubSpotConnectRequest) GetSyncDirection() string`

GetSyncDirection returns the SyncDirection field if non-nil, zero value otherwise.

### GetSyncDirectionOk

`func (o *HubSpotConnectRequest) GetSyncDirectionOk() (*string, bool)`

GetSyncDirectionOk returns a tuple with the SyncDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDirection

`func (o *HubSpotConnectRequest) SetSyncDirection(v string)`

SetSyncDirection sets SyncDirection field to given value.


### GetSyncFrequency

`func (o *HubSpotConnectRequest) GetSyncFrequency() string`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *HubSpotConnectRequest) GetSyncFrequencyOk() (*string, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *HubSpotConnectRequest) SetSyncFrequency(v string)`

SetSyncFrequency sets SyncFrequency field to given value.


### GetOwner

`func (o *HubSpotConnectRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HubSpotConnectRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HubSpotConnectRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HubSpotConnectRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


