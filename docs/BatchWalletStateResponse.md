# BatchWalletStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**States** | [**[]WalletStateResponse**](WalletStateResponse.md) | List of wallet states | 
**Errors** | Pointer to **[]map[string]interface{}** | List of errors for wallets that failed | [optional] 

## Methods

### NewBatchWalletStateResponse

`func NewBatchWalletStateResponse(states []WalletStateResponse, ) *BatchWalletStateResponse`

NewBatchWalletStateResponse instantiates a new BatchWalletStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWalletStateResponseWithDefaults

`func NewBatchWalletStateResponseWithDefaults() *BatchWalletStateResponse`

NewBatchWalletStateResponseWithDefaults instantiates a new BatchWalletStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStates

`func (o *BatchWalletStateResponse) GetStates() []WalletStateResponse`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *BatchWalletStateResponse) GetStatesOk() (*[]WalletStateResponse, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *BatchWalletStateResponse) SetStates(v []WalletStateResponse)`

SetStates sets States field to given value.


### GetErrors

`func (o *BatchWalletStateResponse) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchWalletStateResponse) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchWalletStateResponse) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchWalletStateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


