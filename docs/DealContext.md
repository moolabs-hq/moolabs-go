# DealContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bind** | **map[string]string** |  | 
**Rules** | [**[]DealContextRule**](DealContextRule.md) |  | 

## Methods

### NewDealContext

`func NewDealContext(bind map[string]string, rules []DealContextRule, ) *DealContext`

NewDealContext instantiates a new DealContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealContextWithDefaults

`func NewDealContextWithDefaults() *DealContext`

NewDealContextWithDefaults instantiates a new DealContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBind

`func (o *DealContext) GetBind() map[string]string`

GetBind returns the Bind field if non-nil, zero value otherwise.

### GetBindOk

`func (o *DealContext) GetBindOk() (*map[string]string, bool)`

GetBindOk returns a tuple with the Bind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBind

`func (o *DealContext) SetBind(v map[string]string)`

SetBind sets Bind field to given value.


### GetRules

`func (o *DealContext) GetRules() []DealContextRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DealContext) GetRulesOk() (*[]DealContextRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DealContext) SetRules(v []DealContextRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


