# SDKBatchIngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]SDKEvent**](SDKEvent.md) |  | 

## Methods

### NewSDKBatchIngestRequest

`func NewSDKBatchIngestRequest(events []SDKEvent, ) *SDKBatchIngestRequest`

NewSDKBatchIngestRequest instantiates a new SDKBatchIngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDKBatchIngestRequestWithDefaults

`func NewSDKBatchIngestRequestWithDefaults() *SDKBatchIngestRequest`

NewSDKBatchIngestRequestWithDefaults instantiates a new SDKBatchIngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *SDKBatchIngestRequest) GetEvents() []SDKEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SDKBatchIngestRequest) GetEventsOk() (*[]SDKEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SDKBatchIngestRequest) SetEvents(v []SDKEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


