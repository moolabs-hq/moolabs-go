# ListAutoTopupRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AutoTopupRuleResponse**](AutoTopupRuleResponse.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewListAutoTopupRulesResponse

`func NewListAutoTopupRulesResponse(items []AutoTopupRuleResponse, total int32, ) *ListAutoTopupRulesResponse`

NewListAutoTopupRulesResponse instantiates a new ListAutoTopupRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAutoTopupRulesResponseWithDefaults

`func NewListAutoTopupRulesResponseWithDefaults() *ListAutoTopupRulesResponse`

NewListAutoTopupRulesResponseWithDefaults instantiates a new ListAutoTopupRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListAutoTopupRulesResponse) GetItems() []AutoTopupRuleResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListAutoTopupRulesResponse) GetItemsOk() (*[]AutoTopupRuleResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListAutoTopupRulesResponse) SetItems(v []AutoTopupRuleResponse)`

SetItems sets Items field to given value.


### GetTotal

`func (o *ListAutoTopupRulesResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListAutoTopupRulesResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListAutoTopupRulesResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


