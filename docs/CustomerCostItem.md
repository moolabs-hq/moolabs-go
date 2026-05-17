# CustomerCostItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**TotalCost** | **float32** |  | 
**RequestCount** | **int32** |  | 
**AvgCostPerRequest** | **float32** |  | 
**TopModel** | **string** |  | 

## Methods

### NewCustomerCostItem

`func NewCustomerCostItem(customerId string, totalCost float32, requestCount int32, avgCostPerRequest float32, topModel string, ) *CustomerCostItem`

NewCustomerCostItem instantiates a new CustomerCostItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCostItemWithDefaults

`func NewCustomerCostItemWithDefaults() *CustomerCostItem`

NewCustomerCostItemWithDefaults instantiates a new CustomerCostItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomerCostItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerCostItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerCostItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTotalCost

`func (o *CustomerCostItem) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *CustomerCostItem) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *CustomerCostItem) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.


### GetRequestCount

`func (o *CustomerCostItem) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *CustomerCostItem) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *CustomerCostItem) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.


### GetAvgCostPerRequest

`func (o *CustomerCostItem) GetAvgCostPerRequest() float32`

GetAvgCostPerRequest returns the AvgCostPerRequest field if non-nil, zero value otherwise.

### GetAvgCostPerRequestOk

`func (o *CustomerCostItem) GetAvgCostPerRequestOk() (*float32, bool)`

GetAvgCostPerRequestOk returns a tuple with the AvgCostPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgCostPerRequest

`func (o *CustomerCostItem) SetAvgCostPerRequest(v float32)`

SetAvgCostPerRequest sets AvgCostPerRequest field to given value.


### GetTopModel

`func (o *CustomerCostItem) GetTopModel() string`

GetTopModel returns the TopModel field if non-nil, zero value otherwise.

### GetTopModelOk

`func (o *CustomerCostItem) GetTopModelOk() (*string, bool)`

GetTopModelOk returns a tuple with the TopModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopModel

`func (o *CustomerCostItem) SetTopModel(v string)`

SetTopModel sets TopModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


