# ChannelPreferenceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | One of: email, sms, phone, portal | 
**ConsentStatus** | Pointer to **string** | One of: unknown, granted, denied, revoked | [optional] [default to "unknown"]
**OptOut** | Pointer to **bool** |  | [optional] [default to false]
**SuppressedUntil** | Pointer to **time.Time** |  | [optional] 
**ConsentSource** | Pointer to **string** |  | [optional] 
**ConsentCapturedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChannelPreferenceCreate

`func NewChannelPreferenceCreate(channel string, ) *ChannelPreferenceCreate`

NewChannelPreferenceCreate instantiates a new ChannelPreferenceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPreferenceCreateWithDefaults

`func NewChannelPreferenceCreateWithDefaults() *ChannelPreferenceCreate`

NewChannelPreferenceCreateWithDefaults instantiates a new ChannelPreferenceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelPreferenceCreate) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelPreferenceCreate) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelPreferenceCreate) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetConsentStatus

`func (o *ChannelPreferenceCreate) GetConsentStatus() string`

GetConsentStatus returns the ConsentStatus field if non-nil, zero value otherwise.

### GetConsentStatusOk

`func (o *ChannelPreferenceCreate) GetConsentStatusOk() (*string, bool)`

GetConsentStatusOk returns a tuple with the ConsentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentStatus

`func (o *ChannelPreferenceCreate) SetConsentStatus(v string)`

SetConsentStatus sets ConsentStatus field to given value.

### HasConsentStatus

`func (o *ChannelPreferenceCreate) HasConsentStatus() bool`

HasConsentStatus returns a boolean if a field has been set.

### GetOptOut

`func (o *ChannelPreferenceCreate) GetOptOut() bool`

GetOptOut returns the OptOut field if non-nil, zero value otherwise.

### GetOptOutOk

`func (o *ChannelPreferenceCreate) GetOptOutOk() (*bool, bool)`

GetOptOutOk returns a tuple with the OptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOut

`func (o *ChannelPreferenceCreate) SetOptOut(v bool)`

SetOptOut sets OptOut field to given value.

### HasOptOut

`func (o *ChannelPreferenceCreate) HasOptOut() bool`

HasOptOut returns a boolean if a field has been set.

### GetSuppressedUntil

`func (o *ChannelPreferenceCreate) GetSuppressedUntil() time.Time`

GetSuppressedUntil returns the SuppressedUntil field if non-nil, zero value otherwise.

### GetSuppressedUntilOk

`func (o *ChannelPreferenceCreate) GetSuppressedUntilOk() (*time.Time, bool)`

GetSuppressedUntilOk returns a tuple with the SuppressedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedUntil

`func (o *ChannelPreferenceCreate) SetSuppressedUntil(v time.Time)`

SetSuppressedUntil sets SuppressedUntil field to given value.

### HasSuppressedUntil

`func (o *ChannelPreferenceCreate) HasSuppressedUntil() bool`

HasSuppressedUntil returns a boolean if a field has been set.

### GetConsentSource

`func (o *ChannelPreferenceCreate) GetConsentSource() string`

GetConsentSource returns the ConsentSource field if non-nil, zero value otherwise.

### GetConsentSourceOk

`func (o *ChannelPreferenceCreate) GetConsentSourceOk() (*string, bool)`

GetConsentSourceOk returns a tuple with the ConsentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSource

`func (o *ChannelPreferenceCreate) SetConsentSource(v string)`

SetConsentSource sets ConsentSource field to given value.

### HasConsentSource

`func (o *ChannelPreferenceCreate) HasConsentSource() bool`

HasConsentSource returns a boolean if a field has been set.

### GetConsentCapturedAt

`func (o *ChannelPreferenceCreate) GetConsentCapturedAt() time.Time`

GetConsentCapturedAt returns the ConsentCapturedAt field if non-nil, zero value otherwise.

### GetConsentCapturedAtOk

`func (o *ChannelPreferenceCreate) GetConsentCapturedAtOk() (*time.Time, bool)`

GetConsentCapturedAtOk returns a tuple with the ConsentCapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentCapturedAt

`func (o *ChannelPreferenceCreate) SetConsentCapturedAt(v time.Time)`

SetConsentCapturedAt sets ConsentCapturedAt field to given value.

### HasConsentCapturedAt

`func (o *ChannelPreferenceCreate) HasConsentCapturedAt() bool`

HasConsentCapturedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


