# ChannelPreferenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContactId** | **string** |  | 
**Channel** | **string** |  | 
**ConsentStatus** | **string** |  | 
**OptOut** | **bool** |  | 
**SuppressedUntil** | Pointer to **time.Time** |  | [optional] 
**ConsentSource** | Pointer to **string** |  | [optional] 
**ConsentCapturedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewChannelPreferenceResponse

`func NewChannelPreferenceResponse(id string, contactId string, channel string, consentStatus string, optOut bool, createdAt time.Time, updatedAt time.Time, ) *ChannelPreferenceResponse`

NewChannelPreferenceResponse instantiates a new ChannelPreferenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPreferenceResponseWithDefaults

`func NewChannelPreferenceResponseWithDefaults() *ChannelPreferenceResponse`

NewChannelPreferenceResponseWithDefaults instantiates a new ChannelPreferenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelPreferenceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelPreferenceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelPreferenceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetContactId

`func (o *ChannelPreferenceResponse) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ChannelPreferenceResponse) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ChannelPreferenceResponse) SetContactId(v string)`

SetContactId sets ContactId field to given value.


### GetChannel

`func (o *ChannelPreferenceResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelPreferenceResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelPreferenceResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetConsentStatus

`func (o *ChannelPreferenceResponse) GetConsentStatus() string`

GetConsentStatus returns the ConsentStatus field if non-nil, zero value otherwise.

### GetConsentStatusOk

`func (o *ChannelPreferenceResponse) GetConsentStatusOk() (*string, bool)`

GetConsentStatusOk returns a tuple with the ConsentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentStatus

`func (o *ChannelPreferenceResponse) SetConsentStatus(v string)`

SetConsentStatus sets ConsentStatus field to given value.


### GetOptOut

`func (o *ChannelPreferenceResponse) GetOptOut() bool`

GetOptOut returns the OptOut field if non-nil, zero value otherwise.

### GetOptOutOk

`func (o *ChannelPreferenceResponse) GetOptOutOk() (*bool, bool)`

GetOptOutOk returns a tuple with the OptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOut

`func (o *ChannelPreferenceResponse) SetOptOut(v bool)`

SetOptOut sets OptOut field to given value.


### GetSuppressedUntil

`func (o *ChannelPreferenceResponse) GetSuppressedUntil() time.Time`

GetSuppressedUntil returns the SuppressedUntil field if non-nil, zero value otherwise.

### GetSuppressedUntilOk

`func (o *ChannelPreferenceResponse) GetSuppressedUntilOk() (*time.Time, bool)`

GetSuppressedUntilOk returns a tuple with the SuppressedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedUntil

`func (o *ChannelPreferenceResponse) SetSuppressedUntil(v time.Time)`

SetSuppressedUntil sets SuppressedUntil field to given value.

### HasSuppressedUntil

`func (o *ChannelPreferenceResponse) HasSuppressedUntil() bool`

HasSuppressedUntil returns a boolean if a field has been set.

### GetConsentSource

`func (o *ChannelPreferenceResponse) GetConsentSource() string`

GetConsentSource returns the ConsentSource field if non-nil, zero value otherwise.

### GetConsentSourceOk

`func (o *ChannelPreferenceResponse) GetConsentSourceOk() (*string, bool)`

GetConsentSourceOk returns a tuple with the ConsentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSource

`func (o *ChannelPreferenceResponse) SetConsentSource(v string)`

SetConsentSource sets ConsentSource field to given value.

### HasConsentSource

`func (o *ChannelPreferenceResponse) HasConsentSource() bool`

HasConsentSource returns a boolean if a field has been set.

### GetConsentCapturedAt

`func (o *ChannelPreferenceResponse) GetConsentCapturedAt() time.Time`

GetConsentCapturedAt returns the ConsentCapturedAt field if non-nil, zero value otherwise.

### GetConsentCapturedAtOk

`func (o *ChannelPreferenceResponse) GetConsentCapturedAtOk() (*time.Time, bool)`

GetConsentCapturedAtOk returns a tuple with the ConsentCapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentCapturedAt

`func (o *ChannelPreferenceResponse) SetConsentCapturedAt(v time.Time)`

SetConsentCapturedAt sets ConsentCapturedAt field to given value.

### HasConsentCapturedAt

`func (o *ChannelPreferenceResponse) HasConsentCapturedAt() bool`

HasConsentCapturedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChannelPreferenceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelPreferenceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelPreferenceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ChannelPreferenceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChannelPreferenceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChannelPreferenceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


