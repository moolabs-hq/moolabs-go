# TenantEscalationConfigPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** | Sparse partial config | 
**ChangeReason** | **string** |  | 
**ChangeLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantEscalationConfigPatch

`func NewTenantEscalationConfigPatch(config map[string]interface{}, changeReason string, ) *TenantEscalationConfigPatch`

NewTenantEscalationConfigPatch instantiates a new TenantEscalationConfigPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantEscalationConfigPatchWithDefaults

`func NewTenantEscalationConfigPatchWithDefaults() *TenantEscalationConfigPatch`

NewTenantEscalationConfigPatchWithDefaults instantiates a new TenantEscalationConfigPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *TenantEscalationConfigPatch) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TenantEscalationConfigPatch) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TenantEscalationConfigPatch) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetChangeReason

`func (o *TenantEscalationConfigPatch) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *TenantEscalationConfigPatch) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *TenantEscalationConfigPatch) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.


### GetChangeLabel

`func (o *TenantEscalationConfigPatch) GetChangeLabel() string`

GetChangeLabel returns the ChangeLabel field if non-nil, zero value otherwise.

### GetChangeLabelOk

`func (o *TenantEscalationConfigPatch) GetChangeLabelOk() (*string, bool)`

GetChangeLabelOk returns a tuple with the ChangeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLabel

`func (o *TenantEscalationConfigPatch) SetChangeLabel(v string)`

SetChangeLabel sets ChangeLabel field to given value.

### HasChangeLabel

`func (o *TenantEscalationConfigPatch) HasChangeLabel() bool`

HasChangeLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


