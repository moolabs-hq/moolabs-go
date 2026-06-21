# ApprovalPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**PolicyVersionStamp** | **string** |  | 
**Config** | **map[string]interface{}** |  | 

## Methods

### NewApprovalPolicyResponse

`func NewApprovalPolicyResponse(version int32, policyVersionStamp string, config map[string]interface{}, ) *ApprovalPolicyResponse`

NewApprovalPolicyResponse instantiates a new ApprovalPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalPolicyResponseWithDefaults

`func NewApprovalPolicyResponseWithDefaults() *ApprovalPolicyResponse`

NewApprovalPolicyResponseWithDefaults instantiates a new ApprovalPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ApprovalPolicyResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApprovalPolicyResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApprovalPolicyResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetPolicyVersionStamp

`func (o *ApprovalPolicyResponse) GetPolicyVersionStamp() string`

GetPolicyVersionStamp returns the PolicyVersionStamp field if non-nil, zero value otherwise.

### GetPolicyVersionStampOk

`func (o *ApprovalPolicyResponse) GetPolicyVersionStampOk() (*string, bool)`

GetPolicyVersionStampOk returns a tuple with the PolicyVersionStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersionStamp

`func (o *ApprovalPolicyResponse) SetPolicyVersionStamp(v string)`

SetPolicyVersionStamp sets PolicyVersionStamp field to given value.


### GetConfig

`func (o *ApprovalPolicyResponse) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApprovalPolicyResponse) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApprovalPolicyResponse) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


