# InboundEmailProcessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboundMessageId** | **string** |  | 
**Status** | **string** |  | 
**Created** | **bool** |  | 
**Version** | Pointer to [**ContractVersionResponse**](ContractVersionResponse.md) |  | [optional] 

## Methods

### NewInboundEmailProcessResponse

`func NewInboundEmailProcessResponse(inboundMessageId string, status string, created bool, ) *InboundEmailProcessResponse`

NewInboundEmailProcessResponse instantiates a new InboundEmailProcessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundEmailProcessResponseWithDefaults

`func NewInboundEmailProcessResponseWithDefaults() *InboundEmailProcessResponse`

NewInboundEmailProcessResponseWithDefaults instantiates a new InboundEmailProcessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundMessageId

`func (o *InboundEmailProcessResponse) GetInboundMessageId() string`

GetInboundMessageId returns the InboundMessageId field if non-nil, zero value otherwise.

### GetInboundMessageIdOk

`func (o *InboundEmailProcessResponse) GetInboundMessageIdOk() (*string, bool)`

GetInboundMessageIdOk returns a tuple with the InboundMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMessageId

`func (o *InboundEmailProcessResponse) SetInboundMessageId(v string)`

SetInboundMessageId sets InboundMessageId field to given value.


### GetStatus

`func (o *InboundEmailProcessResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InboundEmailProcessResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InboundEmailProcessResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *InboundEmailProcessResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InboundEmailProcessResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InboundEmailProcessResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.


### GetVersion

`func (o *InboundEmailProcessResponse) GetVersion() ContractVersionResponse`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InboundEmailProcessResponse) GetVersionOk() (*ContractVersionResponse, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InboundEmailProcessResponse) SetVersion(v ContractVersionResponse)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InboundEmailProcessResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


