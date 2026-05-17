# TestMeterEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtractedValue** | Pointer to **float32** |  | [optional] 
**IsNumeric** | **bool** |  | 
**GroupByValues** | **map[string]string** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewTestMeterEvent200Response

`func NewTestMeterEvent200Response(isNumeric bool, groupByValues map[string]string, ) *TestMeterEvent200Response`

NewTestMeterEvent200Response instantiates a new TestMeterEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestMeterEvent200ResponseWithDefaults

`func NewTestMeterEvent200ResponseWithDefaults() *TestMeterEvent200Response`

NewTestMeterEvent200ResponseWithDefaults instantiates a new TestMeterEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtractedValue

`func (o *TestMeterEvent200Response) GetExtractedValue() float32`

GetExtractedValue returns the ExtractedValue field if non-nil, zero value otherwise.

### GetExtractedValueOk

`func (o *TestMeterEvent200Response) GetExtractedValueOk() (*float32, bool)`

GetExtractedValueOk returns a tuple with the ExtractedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedValue

`func (o *TestMeterEvent200Response) SetExtractedValue(v float32)`

SetExtractedValue sets ExtractedValue field to given value.

### HasExtractedValue

`func (o *TestMeterEvent200Response) HasExtractedValue() bool`

HasExtractedValue returns a boolean if a field has been set.

### GetIsNumeric

`func (o *TestMeterEvent200Response) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *TestMeterEvent200Response) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *TestMeterEvent200Response) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.


### GetGroupByValues

`func (o *TestMeterEvent200Response) GetGroupByValues() map[string]string`

GetGroupByValues returns the GroupByValues field if non-nil, zero value otherwise.

### GetGroupByValuesOk

`func (o *TestMeterEvent200Response) GetGroupByValuesOk() (*map[string]string, bool)`

GetGroupByValuesOk returns a tuple with the GroupByValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByValues

`func (o *TestMeterEvent200Response) SetGroupByValues(v map[string]string)`

SetGroupByValues sets GroupByValues field to given value.


### GetReason

`func (o *TestMeterEvent200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TestMeterEvent200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TestMeterEvent200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TestMeterEvent200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


