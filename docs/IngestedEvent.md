# IngestedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**Event**](Event.md) | The original event ingested. | 
**CustomerId** | Pointer to **string** | The customer ID if the event is associated with a customer. | [optional] 
**ValidationError** | Pointer to **string** | The validation error if the event failed validation. | [optional] 
**IngestedAt** | **time.Time** | The date and time the event was ingested. | 
**StoredAt** | **time.Time** | The date and time the event was stored. | 

## Methods

### NewIngestedEvent

`func NewIngestedEvent(event Event, ingestedAt time.Time, storedAt time.Time, ) *IngestedEvent`

NewIngestedEvent instantiates a new IngestedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestedEventWithDefaults

`func NewIngestedEventWithDefaults() *IngestedEvent`

NewIngestedEventWithDefaults instantiates a new IngestedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *IngestedEvent) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IngestedEvent) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IngestedEvent) SetEvent(v Event)`

SetEvent sets Event field to given value.


### GetCustomerId

`func (o *IngestedEvent) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *IngestedEvent) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *IngestedEvent) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *IngestedEvent) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetValidationError

`func (o *IngestedEvent) GetValidationError() string`

GetValidationError returns the ValidationError field if non-nil, zero value otherwise.

### GetValidationErrorOk

`func (o *IngestedEvent) GetValidationErrorOk() (*string, bool)`

GetValidationErrorOk returns a tuple with the ValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationError

`func (o *IngestedEvent) SetValidationError(v string)`

SetValidationError sets ValidationError field to given value.

### HasValidationError

`func (o *IngestedEvent) HasValidationError() bool`

HasValidationError returns a boolean if a field has been set.

### GetIngestedAt

`func (o *IngestedEvent) GetIngestedAt() time.Time`

GetIngestedAt returns the IngestedAt field if non-nil, zero value otherwise.

### GetIngestedAtOk

`func (o *IngestedEvent) GetIngestedAtOk() (*time.Time, bool)`

GetIngestedAtOk returns a tuple with the IngestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedAt

`func (o *IngestedEvent) SetIngestedAt(v time.Time)`

SetIngestedAt sets IngestedAt field to given value.


### GetStoredAt

`func (o *IngestedEvent) GetStoredAt() time.Time`

GetStoredAt returns the StoredAt field if non-nil, zero value otherwise.

### GetStoredAtOk

`func (o *IngestedEvent) GetStoredAtOk() (*time.Time, bool)`

GetStoredAtOk returns a tuple with the StoredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredAt

`func (o *IngestedEvent) SetStoredAt(v time.Time)`

SetStoredAt sets StoredAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


