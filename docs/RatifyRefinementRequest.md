# RatifyRefinementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackId** | **string** |  | 
**IndustryVertical** | **string** |  | 
**Owner** | **string** |  | 
**PositionTier** | Pointer to **string** |  | [optional] [default to "preferred"]

## Methods

### NewRatifyRefinementRequest

`func NewRatifyRefinementRequest(feedbackId string, industryVertical string, owner string, ) *RatifyRefinementRequest`

NewRatifyRefinementRequest instantiates a new RatifyRefinementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatifyRefinementRequestWithDefaults

`func NewRatifyRefinementRequestWithDefaults() *RatifyRefinementRequest`

NewRatifyRefinementRequestWithDefaults instantiates a new RatifyRefinementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackId

`func (o *RatifyRefinementRequest) GetFeedbackId() string`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *RatifyRefinementRequest) GetFeedbackIdOk() (*string, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *RatifyRefinementRequest) SetFeedbackId(v string)`

SetFeedbackId sets FeedbackId field to given value.


### GetIndustryVertical

`func (o *RatifyRefinementRequest) GetIndustryVertical() string`

GetIndustryVertical returns the IndustryVertical field if non-nil, zero value otherwise.

### GetIndustryVerticalOk

`func (o *RatifyRefinementRequest) GetIndustryVerticalOk() (*string, bool)`

GetIndustryVerticalOk returns a tuple with the IndustryVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryVertical

`func (o *RatifyRefinementRequest) SetIndustryVertical(v string)`

SetIndustryVertical sets IndustryVertical field to given value.


### GetOwner

`func (o *RatifyRefinementRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RatifyRefinementRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RatifyRefinementRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPositionTier

`func (o *RatifyRefinementRequest) GetPositionTier() string`

GetPositionTier returns the PositionTier field if non-nil, zero value otherwise.

### GetPositionTierOk

`func (o *RatifyRefinementRequest) GetPositionTierOk() (*string, bool)`

GetPositionTierOk returns a tuple with the PositionTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTier

`func (o *RatifyRefinementRequest) SetPositionTier(v string)`

SetPositionTier sets PositionTier field to given value.

### HasPositionTier

`func (o *RatifyRefinementRequest) HasPositionTier() bool`

HasPositionTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


