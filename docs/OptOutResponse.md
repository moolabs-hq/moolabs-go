# OptOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | **string** |  | 
**Channel** | **string** |  | 
**OptOut** | Pointer to **bool** |  | [optional] [default to true]
**Message** | **string** |  | 

## Methods

### NewOptOutResponse

`func NewOptOutResponse(contactId string, channel string, message string, ) *OptOutResponse`

NewOptOutResponse instantiates a new OptOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptOutResponseWithDefaults

`func NewOptOutResponseWithDefaults() *OptOutResponse`

NewOptOutResponseWithDefaults instantiates a new OptOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *OptOutResponse) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *OptOutResponse) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *OptOutResponse) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetChannel

`func (o *OptOutResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OptOutResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OptOutResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetOptOut

`func (o *OptOutResponse) GetOptOut() bool`

GetOptOut returns the OptOut field if non-nil, zero value otherwise.

### GetOptOutOk

`func (o *OptOutResponse) GetOptOutOk() (*bool, bool)`

GetOptOutOk returns a tuple with the OptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOut

`func (o *OptOutResponse) SetOptOut(v bool)`

SetOptOut sets OptOut field to given value.

### HasOptOut

`func (o *OptOutResponse) HasOptOut() bool`

HasOptOut returns a boolean if a field has been set.

### GetMessage

`func (o *OptOutResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OptOutResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OptOutResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


