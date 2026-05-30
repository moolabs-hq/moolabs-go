# MeterTestEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtractedValue** | Pointer to **float64** |  | [optional] 
**IsNumeric** | **bool** |  | 
**GroupByValues** | **map[string]string** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewMeterTestEventResponse

`func NewMeterTestEventResponse(isNumeric bool, groupByValues map[string]string, ) *MeterTestEventResponse`

NewMeterTestEventResponse instantiates a new MeterTestEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterTestEventResponseWithDefaults

`func NewMeterTestEventResponseWithDefaults() *MeterTestEventResponse`

NewMeterTestEventResponseWithDefaults instantiates a new MeterTestEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtractedValue

`func (o *MeterTestEventResponse) GetExtractedValue() float64`

GetExtractedValue returns the ExtractedValue field if non-nil, zero value otherwise.

### GetExtractedValueOk

`func (o *MeterTestEventResponse) GetExtractedValueOk() (*float64, bool)`

GetExtractedValueOk returns a tuple with the ExtractedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedValue

`func (o *MeterTestEventResponse) SetExtractedValue(v float64)`

SetExtractedValue sets ExtractedValue field to given value.

### HasExtractedValue

`func (o *MeterTestEventResponse) HasExtractedValue() bool`

HasExtractedValue returns a boolean if a field has been set.

### GetIsNumeric

`func (o *MeterTestEventResponse) GetIsNumeric() bool`

GetIsNumeric returns the IsNumeric field if non-nil, zero value otherwise.

### GetIsNumericOk

`func (o *MeterTestEventResponse) GetIsNumericOk() (*bool, bool)`

GetIsNumericOk returns a tuple with the IsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumeric

`func (o *MeterTestEventResponse) SetIsNumeric(v bool)`

SetIsNumeric sets IsNumeric field to given value.


### GetGroupByValues

`func (o *MeterTestEventResponse) GetGroupByValues() map[string]string`

GetGroupByValues returns the GroupByValues field if non-nil, zero value otherwise.

### GetGroupByValuesOk

`func (o *MeterTestEventResponse) GetGroupByValuesOk() (*map[string]string, bool)`

GetGroupByValuesOk returns a tuple with the GroupByValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByValues

`func (o *MeterTestEventResponse) SetGroupByValues(v map[string]string)`

SetGroupByValues sets GroupByValues field to given value.


### GetReason

`func (o *MeterTestEventResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MeterTestEventResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MeterTestEventResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MeterTestEventResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


