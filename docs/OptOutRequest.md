# OptOutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | Channel to opt out of: email, sms, phone, portal | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewOptOutRequest

`func NewOptOutRequest(channel string, ) *OptOutRequest`

NewOptOutRequest instantiates a new OptOutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptOutRequestWithDefaults

`func NewOptOutRequestWithDefaults() *OptOutRequest`

NewOptOutRequestWithDefaults instantiates a new OptOutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *OptOutRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OptOutRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OptOutRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetReason

`func (o *OptOutRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OptOutRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OptOutRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OptOutRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


