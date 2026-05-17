# HandoffCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandoffType** | **string** |  | 
**PartnerName** | **string** |  | 
**CommissionBps** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandoffCreateRequest

`func NewHandoffCreateRequest(handoffType string, partnerName string, ) *HandoffCreateRequest`

NewHandoffCreateRequest instantiates a new HandoffCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandoffCreateRequestWithDefaults

`func NewHandoffCreateRequestWithDefaults() *HandoffCreateRequest`

NewHandoffCreateRequestWithDefaults instantiates a new HandoffCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandoffType

`func (o *HandoffCreateRequest) GetHandoffType() string`

GetHandoffType returns the HandoffType field if non-nil, zero value otherwise.

### GetHandoffTypeOk

`func (o *HandoffCreateRequest) GetHandoffTypeOk() (*string, bool)`

GetHandoffTypeOk returns a tuple with the HandoffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffType

`func (o *HandoffCreateRequest) SetHandoffType(v string)`

SetHandoffType sets HandoffType field to given value.


### GetPartnerName

`func (o *HandoffCreateRequest) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *HandoffCreateRequest) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *HandoffCreateRequest) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### GetCommissionBps

`func (o *HandoffCreateRequest) GetCommissionBps() int32`

GetCommissionBps returns the CommissionBps field if non-nil, zero value otherwise.

### GetCommissionBpsOk

`func (o *HandoffCreateRequest) GetCommissionBpsOk() (*int32, bool)`

GetCommissionBpsOk returns a tuple with the CommissionBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionBps

`func (o *HandoffCreateRequest) SetCommissionBps(v int32)`

SetCommissionBps sets CommissionBps field to given value.

### HasCommissionBps

`func (o *HandoffCreateRequest) HasCommissionBps() bool`

HasCommissionBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


