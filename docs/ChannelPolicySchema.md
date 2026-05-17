# ChannelPolicySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPerWeek** | Pointer to **int32** |  | [optional] [default to 3]
**QuietHoursStart** | Pointer to **string** |  | [optional] [default to "21:00"]
**QuietHoursEnd** | Pointer to **string** |  | [optional] [default to "08:00"]

## Methods

### NewChannelPolicySchema

`func NewChannelPolicySchema() *ChannelPolicySchema`

NewChannelPolicySchema instantiates a new ChannelPolicySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelPolicySchemaWithDefaults

`func NewChannelPolicySchemaWithDefaults() *ChannelPolicySchema`

NewChannelPolicySchemaWithDefaults instantiates a new ChannelPolicySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPerWeek

`func (o *ChannelPolicySchema) GetMaxPerWeek() int32`

GetMaxPerWeek returns the MaxPerWeek field if non-nil, zero value otherwise.

### GetMaxPerWeekOk

`func (o *ChannelPolicySchema) GetMaxPerWeekOk() (*int32, bool)`

GetMaxPerWeekOk returns a tuple with the MaxPerWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerWeek

`func (o *ChannelPolicySchema) SetMaxPerWeek(v int32)`

SetMaxPerWeek sets MaxPerWeek field to given value.

### HasMaxPerWeek

`func (o *ChannelPolicySchema) HasMaxPerWeek() bool`

HasMaxPerWeek returns a boolean if a field has been set.

### GetQuietHoursStart

`func (o *ChannelPolicySchema) GetQuietHoursStart() string`

GetQuietHoursStart returns the QuietHoursStart field if non-nil, zero value otherwise.

### GetQuietHoursStartOk

`func (o *ChannelPolicySchema) GetQuietHoursStartOk() (*string, bool)`

GetQuietHoursStartOk returns a tuple with the QuietHoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHoursStart

`func (o *ChannelPolicySchema) SetQuietHoursStart(v string)`

SetQuietHoursStart sets QuietHoursStart field to given value.

### HasQuietHoursStart

`func (o *ChannelPolicySchema) HasQuietHoursStart() bool`

HasQuietHoursStart returns a boolean if a field has been set.

### GetQuietHoursEnd

`func (o *ChannelPolicySchema) GetQuietHoursEnd() string`

GetQuietHoursEnd returns the QuietHoursEnd field if non-nil, zero value otherwise.

### GetQuietHoursEndOk

`func (o *ChannelPolicySchema) GetQuietHoursEndOk() (*string, bool)`

GetQuietHoursEndOk returns a tuple with the QuietHoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHoursEnd

`func (o *ChannelPolicySchema) SetQuietHoursEnd(v string)`

SetQuietHoursEnd sets QuietHoursEnd field to given value.

### HasQuietHoursEnd

`func (o *ChannelPolicySchema) HasQuietHoursEnd() bool`

HasQuietHoursEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


