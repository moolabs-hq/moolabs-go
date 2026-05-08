# RolloverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**PeriodStart** | **time.Time** |  | 
**PeriodEnd** | **time.Time** |  | 
**GrantsProcessed** | **int32** |  | 
**RolloverAmountMicros** | **int32** |  | 
**RolloverGrantsCreated** | **[]string** |  | 

## Methods

### NewRolloverResponse

`func NewRolloverResponse(subscriptionId string, periodStart time.Time, periodEnd time.Time, grantsProcessed int32, rolloverAmountMicros int32, rolloverGrantsCreated []string, ) *RolloverResponse`

NewRolloverResponse instantiates a new RolloverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloverResponseWithDefaults

`func NewRolloverResponseWithDefaults() *RolloverResponse`

NewRolloverResponseWithDefaults instantiates a new RolloverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *RolloverResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RolloverResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RolloverResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetPeriodStart

`func (o *RolloverResponse) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *RolloverResponse) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *RolloverResponse) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *RolloverResponse) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *RolloverResponse) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *RolloverResponse) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetGrantsProcessed

`func (o *RolloverResponse) GetGrantsProcessed() int32`

GetGrantsProcessed returns the GrantsProcessed field if non-nil, zero value otherwise.

### GetGrantsProcessedOk

`func (o *RolloverResponse) GetGrantsProcessedOk() (*int32, bool)`

GetGrantsProcessedOk returns a tuple with the GrantsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantsProcessed

`func (o *RolloverResponse) SetGrantsProcessed(v int32)`

SetGrantsProcessed sets GrantsProcessed field to given value.


### GetRolloverAmountMicros

`func (o *RolloverResponse) GetRolloverAmountMicros() int32`

GetRolloverAmountMicros returns the RolloverAmountMicros field if non-nil, zero value otherwise.

### GetRolloverAmountMicrosOk

`func (o *RolloverResponse) GetRolloverAmountMicrosOk() (*int32, bool)`

GetRolloverAmountMicrosOk returns a tuple with the RolloverAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverAmountMicros

`func (o *RolloverResponse) SetRolloverAmountMicros(v int32)`

SetRolloverAmountMicros sets RolloverAmountMicros field to given value.


### GetRolloverGrantsCreated

`func (o *RolloverResponse) GetRolloverGrantsCreated() []string`

GetRolloverGrantsCreated returns the RolloverGrantsCreated field if non-nil, zero value otherwise.

### GetRolloverGrantsCreatedOk

`func (o *RolloverResponse) GetRolloverGrantsCreatedOk() (*[]string, bool)`

GetRolloverGrantsCreatedOk returns a tuple with the RolloverGrantsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverGrantsCreated

`func (o *RolloverResponse) SetRolloverGrantsCreated(v []string)`

SetRolloverGrantsCreated sets RolloverGrantsCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


