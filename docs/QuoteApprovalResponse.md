# QuoteApprovalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**TenantId** | **string** |  | 
**Version** | **int32** |  | 
**Status** | **string** |  | 
**ApprovalRequired** | **bool** |  | 
**RequiredRole** | Pointer to **string** |  | [optional] 
**ApprovalPolicyVersion** | **string** |  | 
**ApprovalBasisDigest** | Pointer to **string** |  | [optional] 
**Reasons** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ApprovalId** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteApprovalResponse

`func NewQuoteApprovalResponse(quoteId string, tenantId string, version int32, status string, approvalRequired bool, approvalPolicyVersion string, ) *QuoteApprovalResponse`

NewQuoteApprovalResponse instantiates a new QuoteApprovalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteApprovalResponseWithDefaults

`func NewQuoteApprovalResponseWithDefaults() *QuoteApprovalResponse`

NewQuoteApprovalResponseWithDefaults instantiates a new QuoteApprovalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *QuoteApprovalResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteApprovalResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteApprovalResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetTenantId

`func (o *QuoteApprovalResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteApprovalResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteApprovalResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetVersion

`func (o *QuoteApprovalResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QuoteApprovalResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QuoteApprovalResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *QuoteApprovalResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteApprovalResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteApprovalResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetApprovalRequired

`func (o *QuoteApprovalResponse) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *QuoteApprovalResponse) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *QuoteApprovalResponse) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.


### GetRequiredRole

`func (o *QuoteApprovalResponse) GetRequiredRole() string`

GetRequiredRole returns the RequiredRole field if non-nil, zero value otherwise.

### GetRequiredRoleOk

`func (o *QuoteApprovalResponse) GetRequiredRoleOk() (*string, bool)`

GetRequiredRoleOk returns a tuple with the RequiredRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRole

`func (o *QuoteApprovalResponse) SetRequiredRole(v string)`

SetRequiredRole sets RequiredRole field to given value.

### HasRequiredRole

`func (o *QuoteApprovalResponse) HasRequiredRole() bool`

HasRequiredRole returns a boolean if a field has been set.

### GetApprovalPolicyVersion

`func (o *QuoteApprovalResponse) GetApprovalPolicyVersion() string`

GetApprovalPolicyVersion returns the ApprovalPolicyVersion field if non-nil, zero value otherwise.

### GetApprovalPolicyVersionOk

`func (o *QuoteApprovalResponse) GetApprovalPolicyVersionOk() (*string, bool)`

GetApprovalPolicyVersionOk returns a tuple with the ApprovalPolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalPolicyVersion

`func (o *QuoteApprovalResponse) SetApprovalPolicyVersion(v string)`

SetApprovalPolicyVersion sets ApprovalPolicyVersion field to given value.


### GetApprovalBasisDigest

`func (o *QuoteApprovalResponse) GetApprovalBasisDigest() string`

GetApprovalBasisDigest returns the ApprovalBasisDigest field if non-nil, zero value otherwise.

### GetApprovalBasisDigestOk

`func (o *QuoteApprovalResponse) GetApprovalBasisDigestOk() (*string, bool)`

GetApprovalBasisDigestOk returns a tuple with the ApprovalBasisDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalBasisDigest

`func (o *QuoteApprovalResponse) SetApprovalBasisDigest(v string)`

SetApprovalBasisDigest sets ApprovalBasisDigest field to given value.

### HasApprovalBasisDigest

`func (o *QuoteApprovalResponse) HasApprovalBasisDigest() bool`

HasApprovalBasisDigest returns a boolean if a field has been set.

### GetReasons

`func (o *QuoteApprovalResponse) GetReasons() []map[string]interface{}`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *QuoteApprovalResponse) GetReasonsOk() (*[]map[string]interface{}, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *QuoteApprovalResponse) SetReasons(v []map[string]interface{})`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *QuoteApprovalResponse) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetApprovalId

`func (o *QuoteApprovalResponse) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *QuoteApprovalResponse) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *QuoteApprovalResponse) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.

### HasApprovalId

`func (o *QuoteApprovalResponse) HasApprovalId() bool`

HasApprovalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


