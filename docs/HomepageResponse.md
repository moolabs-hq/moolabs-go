# HomepageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Insights** | [**[]InsightCard**](InsightCard.md) |  | 
**QuickActions** | [**[]QuickAction**](QuickAction.md) |  | 
**Greeting** | **string** |  | 

## Methods

### NewHomepageResponse

`func NewHomepageResponse(insights []InsightCard, quickActions []QuickAction, greeting string, ) *HomepageResponse`

NewHomepageResponse instantiates a new HomepageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHomepageResponseWithDefaults

`func NewHomepageResponseWithDefaults() *HomepageResponse`

NewHomepageResponseWithDefaults instantiates a new HomepageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsights

`func (o *HomepageResponse) GetInsights() []InsightCard`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *HomepageResponse) GetInsightsOk() (*[]InsightCard, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *HomepageResponse) SetInsights(v []InsightCard)`

SetInsights sets Insights field to given value.


### GetQuickActions

`func (o *HomepageResponse) GetQuickActions() []QuickAction`

GetQuickActions returns the QuickActions field if non-nil, zero value otherwise.

### GetQuickActionsOk

`func (o *HomepageResponse) GetQuickActionsOk() (*[]QuickAction, bool)`

GetQuickActionsOk returns a tuple with the QuickActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickActions

`func (o *HomepageResponse) SetQuickActions(v []QuickAction)`

SetQuickActions sets QuickActions field to given value.


### GetGreeting

`func (o *HomepageResponse) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *HomepageResponse) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *HomepageResponse) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


