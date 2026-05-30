# CreateQuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRef** | **map[string]interface{}** |  | 
**DealRef** | Pointer to **map[string]interface{}** |  | [optional] 
**BuyerContactRef** | Pointer to **map[string]interface{}** |  | [optional] 
**QuoteType** | Pointer to **string** |  | [optional] [default to "new_subscription"]

## Methods

### NewCreateQuoteRequest

`func NewCreateQuoteRequest(accountRef map[string]interface{}, ) *CreateQuoteRequest`

NewCreateQuoteRequest instantiates a new CreateQuoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuoteRequestWithDefaults

`func NewCreateQuoteRequestWithDefaults() *CreateQuoteRequest`

NewCreateQuoteRequestWithDefaults instantiates a new CreateQuoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRef

`func (o *CreateQuoteRequest) GetAccountRef() map[string]interface{}`

GetAccountRef returns the AccountRef field if non-nil, zero value otherwise.

### GetAccountRefOk

`func (o *CreateQuoteRequest) GetAccountRefOk() (*map[string]interface{}, bool)`

GetAccountRefOk returns a tuple with the AccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRef

`func (o *CreateQuoteRequest) SetAccountRef(v map[string]interface{})`

SetAccountRef sets AccountRef field to given value.


### GetDealRef

`func (o *CreateQuoteRequest) GetDealRef() map[string]interface{}`

GetDealRef returns the DealRef field if non-nil, zero value otherwise.

### GetDealRefOk

`func (o *CreateQuoteRequest) GetDealRefOk() (*map[string]interface{}, bool)`

GetDealRefOk returns a tuple with the DealRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealRef

`func (o *CreateQuoteRequest) SetDealRef(v map[string]interface{})`

SetDealRef sets DealRef field to given value.

### HasDealRef

`func (o *CreateQuoteRequest) HasDealRef() bool`

HasDealRef returns a boolean if a field has been set.

### GetBuyerContactRef

`func (o *CreateQuoteRequest) GetBuyerContactRef() map[string]interface{}`

GetBuyerContactRef returns the BuyerContactRef field if non-nil, zero value otherwise.

### GetBuyerContactRefOk

`func (o *CreateQuoteRequest) GetBuyerContactRefOk() (*map[string]interface{}, bool)`

GetBuyerContactRefOk returns a tuple with the BuyerContactRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerContactRef

`func (o *CreateQuoteRequest) SetBuyerContactRef(v map[string]interface{})`

SetBuyerContactRef sets BuyerContactRef field to given value.

### HasBuyerContactRef

`func (o *CreateQuoteRequest) HasBuyerContactRef() bool`

HasBuyerContactRef returns a boolean if a field has been set.

### GetQuoteType

`func (o *CreateQuoteRequest) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *CreateQuoteRequest) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *CreateQuoteRequest) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *CreateQuoteRequest) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


