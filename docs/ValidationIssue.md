# ValidationIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Id** | **string** | ID of the charge or discount. | [readonly] 
**Severity** | [**ValidationIssueSeverity**](ValidationIssueSeverity.md) | The severity of the issue. | [readonly] 
**Field** | Pointer to **string** | The field that the issue is related to, if available in JSON path format. | [optional] [readonly] 
**Code** | Pointer to **string** | Machine indentifiable code for the issue, if available. | [optional] [readonly] 
**Component** | **string** | Component reporting the issue. | [readonly] 
**Message** | **string** | A human-readable description of the issue. | [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional context for the issue. | [optional] [readonly] 

## Methods

### NewValidationIssue

`func NewValidationIssue(createdAt time.Time, updatedAt time.Time, id string, severity ValidationIssueSeverity, component string, message string, ) *ValidationIssue`

NewValidationIssue instantiates a new ValidationIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationIssueWithDefaults

`func NewValidationIssueWithDefaults() *ValidationIssue`

NewValidationIssueWithDefaults instantiates a new ValidationIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ValidationIssue) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ValidationIssue) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ValidationIssue) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ValidationIssue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ValidationIssue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ValidationIssue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *ValidationIssue) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ValidationIssue) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ValidationIssue) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ValidationIssue) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *ValidationIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidationIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidationIssue) SetId(v string)`

SetId sets Id field to given value.


### GetSeverity

`func (o *ValidationIssue) GetSeverity() ValidationIssueSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationIssue) GetSeverityOk() (*ValidationIssueSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationIssue) SetSeverity(v ValidationIssueSeverity)`

SetSeverity sets Severity field to given value.


### GetField

`func (o *ValidationIssue) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationIssue) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationIssue) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ValidationIssue) HasField() bool`

HasField returns a boolean if a field has been set.

### GetCode

`func (o *ValidationIssue) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationIssue) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationIssue) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ValidationIssue) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetComponent

`func (o *ValidationIssue) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ValidationIssue) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ValidationIssue) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetMessage

`func (o *ValidationIssue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationIssue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationIssue) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetadata

`func (o *ValidationIssue) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ValidationIssue) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ValidationIssue) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ValidationIssue) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


