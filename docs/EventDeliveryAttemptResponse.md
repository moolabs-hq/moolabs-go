# EventDeliveryAttemptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | Status code of the response if available. | [optional] [readonly] 
**Body** | **string** | The body of the response. | [readonly] 
**DurationMs** | **int32** | The duration of the response in milliseconds. | [readonly] 
**Url** | Pointer to **string** | URL where the event was sent in case of notification channel with webhook type. | [optional] [readonly] 

## Methods

### NewEventDeliveryAttemptResponse

`func NewEventDeliveryAttemptResponse(body string, durationMs int32, ) *EventDeliveryAttemptResponse`

NewEventDeliveryAttemptResponse instantiates a new EventDeliveryAttemptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDeliveryAttemptResponseWithDefaults

`func NewEventDeliveryAttemptResponseWithDefaults() *EventDeliveryAttemptResponse`

NewEventDeliveryAttemptResponseWithDefaults instantiates a new EventDeliveryAttemptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *EventDeliveryAttemptResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EventDeliveryAttemptResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EventDeliveryAttemptResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *EventDeliveryAttemptResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetBody

`func (o *EventDeliveryAttemptResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EventDeliveryAttemptResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EventDeliveryAttemptResponse) SetBody(v string)`

SetBody sets Body field to given value.


### GetDurationMs

`func (o *EventDeliveryAttemptResponse) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *EventDeliveryAttemptResponse) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *EventDeliveryAttemptResponse) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetUrl

`func (o *EventDeliveryAttemptResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventDeliveryAttemptResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventDeliveryAttemptResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventDeliveryAttemptResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


