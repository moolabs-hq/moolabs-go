# SyncConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** |  | 
**Schedule** | Pointer to **string** |  | [optional] [default to "0 6 * * *"]

## Methods

### NewSyncConfigRequest

`func NewSyncConfigRequest(cloudProvider string, ) *SyncConfigRequest`

NewSyncConfigRequest instantiates a new SyncConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncConfigRequestWithDefaults

`func NewSyncConfigRequestWithDefaults() *SyncConfigRequest`

NewSyncConfigRequestWithDefaults instantiates a new SyncConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *SyncConfigRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SyncConfigRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SyncConfigRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetSchedule

`func (o *SyncConfigRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SyncConfigRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SyncConfigRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *SyncConfigRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


