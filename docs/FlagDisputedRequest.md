# FlagDisputedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisputedAmountMicros** | **int32** |  | 
**Description** | **string** |  | 
**DisputeType** | Pointer to **string** | Dispute type for classification | [optional] [default to "admin"]

## Methods

### NewFlagDisputedRequest

`func NewFlagDisputedRequest(disputedAmountMicros int32, description string, ) *FlagDisputedRequest`

NewFlagDisputedRequest instantiates a new FlagDisputedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagDisputedRequestWithDefaults

`func NewFlagDisputedRequestWithDefaults() *FlagDisputedRequest`

NewFlagDisputedRequestWithDefaults instantiates a new FlagDisputedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputedAmountMicros

`func (o *FlagDisputedRequest) GetDisputedAmountMicros() int32`

GetDisputedAmountMicros returns the DisputedAmountMicros field if non-nil, zero value otherwise.

### GetDisputedAmountMicrosOk

`func (o *FlagDisputedRequest) GetDisputedAmountMicrosOk() (*int32, bool)`

GetDisputedAmountMicrosOk returns a tuple with the DisputedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmountMicros

`func (o *FlagDisputedRequest) SetDisputedAmountMicros(v int32)`

SetDisputedAmountMicros sets DisputedAmountMicros field to given value.


### GetDescription

`func (o *FlagDisputedRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagDisputedRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagDisputedRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisputeType

`func (o *FlagDisputedRequest) GetDisputeType() string`

GetDisputeType returns the DisputeType field if non-nil, zero value otherwise.

### GetDisputeTypeOk

`func (o *FlagDisputedRequest) GetDisputeTypeOk() (*string, bool)`

GetDisputeTypeOk returns a tuple with the DisputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeType

`func (o *FlagDisputedRequest) SetDisputeType(v string)`

SetDisputeType sets DisputeType field to given value.

### HasDisputeType

`func (o *FlagDisputedRequest) HasDisputeType() bool`

HasDisputeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


