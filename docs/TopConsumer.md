# TopConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**TotalCost** | **float32** |  | 
**RequestCount** | **int32** |  | 
**Rank** | **int32** |  | 

## Methods

### NewTopConsumer

`func NewTopConsumer(customerId string, totalCost float32, requestCount int32, rank int32, ) *TopConsumer`

NewTopConsumer instantiates a new TopConsumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopConsumerWithDefaults

`func NewTopConsumerWithDefaults() *TopConsumer`

NewTopConsumerWithDefaults instantiates a new TopConsumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *TopConsumer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TopConsumer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TopConsumer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTotalCost

`func (o *TopConsumer) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *TopConsumer) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *TopConsumer) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### GetRequestCount

`func (o *TopConsumer) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *TopConsumer) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *TopConsumer) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.


### GetRank

`func (o *TopConsumer) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *TopConsumer) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *TopConsumer) SetRank(v int32)`

SetRank sets Rank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


