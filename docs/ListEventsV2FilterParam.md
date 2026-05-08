# ListEventsV2FilterParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**FilterString**](FilterString.md) |  | [optional] 
**Source** | Pointer to [**FilterString**](FilterString.md) |  | [optional] 
**Subject** | Pointer to [**FilterString**](FilterString.md) |  | [optional] 
**CustomerId** | Pointer to [**FilterIDExact**](FilterIDExact.md) |  | [optional] 
**Type** | Pointer to [**FilterString**](FilterString.md) |  | [optional] 
**Time** | Pointer to [**FilterTime**](FilterTime.md) |  | [optional] 
**IngestedAt** | Pointer to [**FilterTime**](FilterTime.md) |  | [optional] 

## Methods

### NewListEventsV2FilterParam

`func NewListEventsV2FilterParam() *ListEventsV2FilterParam`

NewListEventsV2FilterParam instantiates a new ListEventsV2FilterParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventsV2FilterParamWithDefaults

`func NewListEventsV2FilterParamWithDefaults() *ListEventsV2FilterParam`

NewListEventsV2FilterParamWithDefaults instantiates a new ListEventsV2FilterParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListEventsV2FilterParam) GetId() FilterString`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListEventsV2FilterParam) GetIdOk() (*FilterString, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListEventsV2FilterParam) SetId(v FilterString)`

SetId sets Id field to given value.

### HasId

`func (o *ListEventsV2FilterParam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *ListEventsV2FilterParam) GetSource() FilterString`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListEventsV2FilterParam) GetSourceOk() (*FilterString, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListEventsV2FilterParam) SetSource(v FilterString)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListEventsV2FilterParam) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *ListEventsV2FilterParam) GetSubject() FilterString`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ListEventsV2FilterParam) GetSubjectOk() (*FilterString, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ListEventsV2FilterParam) SetSubject(v FilterString)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ListEventsV2FilterParam) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetCustomerId

`func (o *ListEventsV2FilterParam) GetCustomerId() FilterIDExact`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListEventsV2FilterParam) GetCustomerIdOk() (*FilterIDExact, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListEventsV2FilterParam) SetCustomerId(v FilterIDExact)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ListEventsV2FilterParam) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetType

`func (o *ListEventsV2FilterParam) GetType() FilterString`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListEventsV2FilterParam) GetTypeOk() (*FilterString, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListEventsV2FilterParam) SetType(v FilterString)`

SetType sets Type field to given value.

### HasType

`func (o *ListEventsV2FilterParam) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTime

`func (o *ListEventsV2FilterParam) GetTime() FilterTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ListEventsV2FilterParam) GetTimeOk() (*FilterTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ListEventsV2FilterParam) SetTime(v FilterTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *ListEventsV2FilterParam) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetIngestedAt

`func (o *ListEventsV2FilterParam) GetIngestedAt() FilterTime`

GetIngestedAt returns the IngestedAt field if non-nil, zero value otherwise.

### GetIngestedAtOk

`func (o *ListEventsV2FilterParam) GetIngestedAtOk() (*FilterTime, bool)`

GetIngestedAtOk returns a tuple with the IngestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedAt

`func (o *ListEventsV2FilterParam) SetIngestedAt(v FilterTime)`

SetIngestedAt sets IngestedAt field to given value.

### HasIngestedAt

`func (o *ListEventsV2FilterParam) HasIngestedAt() bool`

HasIngestedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


