# QuickBooksConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** |  | 
**SyncDirection** | **string** |  | 
**SyncFrequency** | **string** |  | 
**Owner** | Pointer to **string** |  | [optional] 

## Methods

### NewQuickBooksConnectRequest

`func NewQuickBooksConnectRequest(environment string, syncDirection string, syncFrequency string, ) *QuickBooksConnectRequest`

NewQuickBooksConnectRequest instantiates a new QuickBooksConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickBooksConnectRequestWithDefaults

`func NewQuickBooksConnectRequestWithDefaults() *QuickBooksConnectRequest`

NewQuickBooksConnectRequestWithDefaults instantiates a new QuickBooksConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *QuickBooksConnectRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *QuickBooksConnectRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *QuickBooksConnectRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetSyncDirection

`func (o *QuickBooksConnectRequest) GetSyncDirection() string`

GetSyncDirection returns the SyncDirection field if non-nil, zero value otherwise.

### GetSyncDirectionOk

`func (o *QuickBooksConnectRequest) GetSyncDirectionOk() (*string, bool)`

GetSyncDirectionOk returns a tuple with the SyncDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDirection

`func (o *QuickBooksConnectRequest) SetSyncDirection(v string)`

SetSyncDirection sets SyncDirection field to given value.


### GetSyncFrequency

`func (o *QuickBooksConnectRequest) GetSyncFrequency() string`

GetSyncFrequency returns the SyncFrequency field if non-nil, zero value otherwise.

### GetSyncFrequencyOk

`func (o *QuickBooksConnectRequest) GetSyncFrequencyOk() (*string, bool)`

GetSyncFrequencyOk returns a tuple with the SyncFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFrequency

`func (o *QuickBooksConnectRequest) SetSyncFrequency(v string)`

SetSyncFrequency sets SyncFrequency field to given value.


### GetOwner

`func (o *QuickBooksConnectRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *QuickBooksConnectRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *QuickBooksConnectRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *QuickBooksConnectRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


