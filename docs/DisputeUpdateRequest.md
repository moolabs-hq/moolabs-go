# DisputeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisputeType** | Pointer to **string** |  | [optional] 
**ClassificationStatus** | Pointer to **string** |  | [optional] 
**AccountingClass** | Pointer to **string** |  | [optional] 
**AccountingReasonCode** | Pointer to **string** |  | [optional] 
**ExternalDisputeId** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**ResolutionDeadline** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDisputeUpdateRequest

`func NewDisputeUpdateRequest() *DisputeUpdateRequest`

NewDisputeUpdateRequest instantiates a new DisputeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeUpdateRequestWithDefaults

`func NewDisputeUpdateRequestWithDefaults() *DisputeUpdateRequest`

NewDisputeUpdateRequestWithDefaults instantiates a new DisputeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputeType

`func (o *DisputeUpdateRequest) GetDisputeType() string`

GetDisputeType returns the DisputeType field if non-nil, zero value otherwise.

### GetDisputeTypeOk

`func (o *DisputeUpdateRequest) GetDisputeTypeOk() (*string, bool)`

GetDisputeTypeOk returns a tuple with the DisputeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeType

`func (o *DisputeUpdateRequest) SetDisputeType(v string)`

SetDisputeType sets DisputeType field to given value.

### HasDisputeType

`func (o *DisputeUpdateRequest) HasDisputeType() bool`

HasDisputeType returns a boolean if a field has been set.

### GetClassificationStatus

`func (o *DisputeUpdateRequest) GetClassificationStatus() string`

GetClassificationStatus returns the ClassificationStatus field if non-nil, zero value otherwise.

### GetClassificationStatusOk

`func (o *DisputeUpdateRequest) GetClassificationStatusOk() (*string, bool)`

GetClassificationStatusOk returns a tuple with the ClassificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationStatus

`func (o *DisputeUpdateRequest) SetClassificationStatus(v string)`

SetClassificationStatus sets ClassificationStatus field to given value.

### HasClassificationStatus

`func (o *DisputeUpdateRequest) HasClassificationStatus() bool`

HasClassificationStatus returns a boolean if a field has been set.

### GetAccountingClass

`func (o *DisputeUpdateRequest) GetAccountingClass() string`

GetAccountingClass returns the AccountingClass field if non-nil, zero value otherwise.

### GetAccountingClassOk

`func (o *DisputeUpdateRequest) GetAccountingClassOk() (*string, bool)`

GetAccountingClassOk returns a tuple with the AccountingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingClass

`func (o *DisputeUpdateRequest) SetAccountingClass(v string)`

SetAccountingClass sets AccountingClass field to given value.

### HasAccountingClass

`func (o *DisputeUpdateRequest) HasAccountingClass() bool`

HasAccountingClass returns a boolean if a field has been set.

### GetAccountingReasonCode

`func (o *DisputeUpdateRequest) GetAccountingReasonCode() string`

GetAccountingReasonCode returns the AccountingReasonCode field if non-nil, zero value otherwise.

### GetAccountingReasonCodeOk

`func (o *DisputeUpdateRequest) GetAccountingReasonCodeOk() (*string, bool)`

GetAccountingReasonCodeOk returns a tuple with the AccountingReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingReasonCode

`func (o *DisputeUpdateRequest) SetAccountingReasonCode(v string)`

SetAccountingReasonCode sets AccountingReasonCode field to given value.

### HasAccountingReasonCode

`func (o *DisputeUpdateRequest) HasAccountingReasonCode() bool`

HasAccountingReasonCode returns a boolean if a field has been set.

### GetExternalDisputeId

`func (o *DisputeUpdateRequest) GetExternalDisputeId() string`

GetExternalDisputeId returns the ExternalDisputeId field if non-nil, zero value otherwise.

### GetExternalDisputeIdOk

`func (o *DisputeUpdateRequest) GetExternalDisputeIdOk() (*string, bool)`

GetExternalDisputeIdOk returns a tuple with the ExternalDisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDisputeId

`func (o *DisputeUpdateRequest) SetExternalDisputeId(v string)`

SetExternalDisputeId sets ExternalDisputeId field to given value.

### HasExternalDisputeId

`func (o *DisputeUpdateRequest) HasExternalDisputeId() bool`

HasExternalDisputeId returns a boolean if a field has been set.

### GetAssignedTo

`func (o *DisputeUpdateRequest) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *DisputeUpdateRequest) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *DisputeUpdateRequest) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *DisputeUpdateRequest) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetResolutionDeadline

`func (o *DisputeUpdateRequest) GetResolutionDeadline() time.Time`

GetResolutionDeadline returns the ResolutionDeadline field if non-nil, zero value otherwise.

### GetResolutionDeadlineOk

`func (o *DisputeUpdateRequest) GetResolutionDeadlineOk() (*time.Time, bool)`

GetResolutionDeadlineOk returns a tuple with the ResolutionDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionDeadline

`func (o *DisputeUpdateRequest) SetResolutionDeadline(v time.Time)`

SetResolutionDeadline sets ResolutionDeadline field to given value.

### HasResolutionDeadline

`func (o *DisputeUpdateRequest) HasResolutionDeadline() bool`

HasResolutionDeadline returns a boolean if a field has been set.

### GetDescription

`func (o *DisputeUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DisputeUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DisputeUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DisputeUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


