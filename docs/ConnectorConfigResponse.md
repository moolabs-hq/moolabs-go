# ConnectorConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**ConnectorType** | **string** |  | 
**Status** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewConnectorConfigResponse

`func NewConnectorConfigResponse(tenantId string, connectorType string, status string, message string, ) *ConnectorConfigResponse`

NewConnectorConfigResponse instantiates a new ConnectorConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorConfigResponseWithDefaults

`func NewConnectorConfigResponseWithDefaults() *ConnectorConfigResponse`

NewConnectorConfigResponseWithDefaults instantiates a new ConnectorConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ConnectorConfigResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ConnectorConfigResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ConnectorConfigResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetConnectorType

`func (o *ConnectorConfigResponse) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *ConnectorConfigResponse) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *ConnectorConfigResponse) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.


### GetStatus

`func (o *ConnectorConfigResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorConfigResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorConfigResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ConnectorConfigResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectorConfigResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectorConfigResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


