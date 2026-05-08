# GrantUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantId** | **string** | The id of the grant | 
**Usage** | **float64** | The usage in the period | 

## Methods

### NewGrantUsageRecord

`func NewGrantUsageRecord(grantId string, usage float64, ) *GrantUsageRecord`

NewGrantUsageRecord instantiates a new GrantUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantUsageRecordWithDefaults

`func NewGrantUsageRecordWithDefaults() *GrantUsageRecord`

NewGrantUsageRecordWithDefaults instantiates a new GrantUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantId

`func (o *GrantUsageRecord) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *GrantUsageRecord) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *GrantUsageRecord) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.


### GetUsage

`func (o *GrantUsageRecord) GetUsage() float64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GrantUsageRecord) GetUsageOk() (*float64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GrantUsageRecord) SetUsage(v float64)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


