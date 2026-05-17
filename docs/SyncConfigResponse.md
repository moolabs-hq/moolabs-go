# SyncConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Message** | **string** |  | 
**CloudProvider** | **string** |  | 
**Schedule** | **string** |  | 

## Methods

### NewSyncConfigResponse

`func NewSyncConfigResponse(status string, message string, cloudProvider string, schedule string, ) *SyncConfigResponse`

NewSyncConfigResponse instantiates a new SyncConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncConfigResponseWithDefaults

`func NewSyncConfigResponseWithDefaults() *SyncConfigResponse`

NewSyncConfigResponseWithDefaults instantiates a new SyncConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SyncConfigResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncConfigResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncConfigResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *SyncConfigResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyncConfigResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyncConfigResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCloudProvider

`func (o *SyncConfigResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SyncConfigResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SyncConfigResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetSchedule

`func (o *SyncConfigResponse) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SyncConfigResponse) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SyncConfigResponse) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


