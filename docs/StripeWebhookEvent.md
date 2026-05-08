# StripeWebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The event ID. | 
**Type** | **string** | The event type. | 
**Livemode** | **bool** | Live mode. | 
**Created** | **int32** | The event created timestamp. | 
**Data** | [**StripeWebhookEventData**](StripeWebhookEventData.md) |  | 

## Methods

### NewStripeWebhookEvent

`func NewStripeWebhookEvent(id string, type_ string, livemode bool, created int32, data StripeWebhookEventData, ) *StripeWebhookEvent`

NewStripeWebhookEvent instantiates a new StripeWebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeWebhookEventWithDefaults

`func NewStripeWebhookEventWithDefaults() *StripeWebhookEvent`

NewStripeWebhookEventWithDefaults instantiates a new StripeWebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeWebhookEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeWebhookEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeWebhookEvent) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *StripeWebhookEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeWebhookEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeWebhookEvent) SetType(v string)`

SetType sets Type field to given value.


### GetLivemode

`func (o *StripeWebhookEvent) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *StripeWebhookEvent) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *StripeWebhookEvent) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetCreated

`func (o *StripeWebhookEvent) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StripeWebhookEvent) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StripeWebhookEvent) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetData

`func (o *StripeWebhookEvent) GetData() StripeWebhookEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StripeWebhookEvent) GetDataOk() (*StripeWebhookEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StripeWebhookEvent) SetData(v StripeWebhookEventData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


