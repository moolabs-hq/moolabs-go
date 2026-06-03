# QuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**AccountRef** | **map[string]interface{}** |  | 
**DealRef** | **map[string]interface{}** |  | 
**BuyerContactRef** | **map[string]interface{}** |  | 
**QuoteType** | Pointer to **string** |  | [optional] [default to "new_subscription"]
**OwnerUserId** | **string** |  | 
**CurrentVersion** | **int32** |  | 
**State** | **string** |  | 
**ExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteResponse

`func NewQuoteResponse(id string, tenantId string, accountRef map[string]interface{}, dealRef map[string]interface{}, buyerContactRef map[string]interface{}, ownerUserId string, currentVersion int32, state string, ) *QuoteResponse`

NewQuoteResponse instantiates a new QuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResponseWithDefaults

`func NewQuoteResponseWithDefaults() *QuoteResponse`

NewQuoteResponseWithDefaults instantiates a new QuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAccountRef

`func (o *QuoteResponse) GetAccountRef() map[string]interface{}`

GetAccountRef returns the AccountRef field if non-nil, zero value otherwise.

### GetAccountRefOk

`func (o *QuoteResponse) GetAccountRefOk() (*map[string]interface{}, bool)`

GetAccountRefOk returns a tuple with the AccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRef

`func (o *QuoteResponse) SetAccountRef(v map[string]interface{})`

SetAccountRef sets AccountRef field to given value.


### GetDealRef

`func (o *QuoteResponse) GetDealRef() map[string]interface{}`

GetDealRef returns the DealRef field if non-nil, zero value otherwise.

### GetDealRefOk

`func (o *QuoteResponse) GetDealRefOk() (*map[string]interface{}, bool)`

GetDealRefOk returns a tuple with the DealRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealRef

`func (o *QuoteResponse) SetDealRef(v map[string]interface{})`

SetDealRef sets DealRef field to given value.


### GetBuyerContactRef

`func (o *QuoteResponse) GetBuyerContactRef() map[string]interface{}`

GetBuyerContactRef returns the BuyerContactRef field if non-nil, zero value otherwise.

### GetBuyerContactRefOk

`func (o *QuoteResponse) GetBuyerContactRefOk() (*map[string]interface{}, bool)`

GetBuyerContactRefOk returns a tuple with the BuyerContactRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerContactRef

`func (o *QuoteResponse) SetBuyerContactRef(v map[string]interface{})`

SetBuyerContactRef sets BuyerContactRef field to given value.


### GetQuoteType

`func (o *QuoteResponse) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *QuoteResponse) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *QuoteResponse) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *QuoteResponse) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *QuoteResponse) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *QuoteResponse) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *QuoteResponse) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.


### GetCurrentVersion

`func (o *QuoteResponse) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *QuoteResponse) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *QuoteResponse) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetState

`func (o *QuoteResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuoteResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuoteResponse) SetState(v string)`

SetState sets State field to given value.


### GetExpiresAt

`func (o *QuoteResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *QuoteResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *QuoteResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *QuoteResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


