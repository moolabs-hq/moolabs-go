# RotateKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorType** | **string** |  | 
**Message** | **string** |  | 
**GracePeriodEndsAt** | **time.Time** |  | 

## Methods

### NewRotateKeyResponse

`func NewRotateKeyResponse(connectorType string, message string, gracePeriodEndsAt time.Time, ) *RotateKeyResponse`

NewRotateKeyResponse instantiates a new RotateKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateKeyResponseWithDefaults

`func NewRotateKeyResponseWithDefaults() *RotateKeyResponse`

NewRotateKeyResponseWithDefaults instantiates a new RotateKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorType

`func (o *RotateKeyResponse) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *RotateKeyResponse) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *RotateKeyResponse) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.


### GetMessage

`func (o *RotateKeyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RotateKeyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RotateKeyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetGracePeriodEndsAt

`func (o *RotateKeyResponse) GetGracePeriodEndsAt() time.Time`

GetGracePeriodEndsAt returns the GracePeriodEndsAt field if non-nil, zero value otherwise.

### GetGracePeriodEndsAtOk

`func (o *RotateKeyResponse) GetGracePeriodEndsAtOk() (*time.Time, bool)`

GetGracePeriodEndsAtOk returns a tuple with the GracePeriodEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndsAt

`func (o *RotateKeyResponse) SetGracePeriodEndsAt(v time.Time)`

SetGracePeriodEndsAt sets GracePeriodEndsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


