# CreateQuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRef** | **map[string]interface{}** |  | 
**DealRef** | Pointer to **map[string]interface{}** |  | [optional] 
**BuyerContactRef** | Pointer to **map[string]interface{}** |  | [optional] 
**QuoteType** | Pointer to **string** |  | [optional] [default to "new_subscription"]
**TargetSubscriptionRef** | Pointer to **map[string]interface{}** |  | [optional] 
**CurrentContractSnapshotDigest** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**CoTermBehavior** | Pointer to **string** |  | [optional] 
**ProrationBasis** | Pointer to **map[string]interface{}** |  | [optional] 
**ChangeReason** | Pointer to **string** |  | [optional] 

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

### GetTargetSubscriptionRef

`func (o *CreateQuoteRequest) GetTargetSubscriptionRef() map[string]interface{}`

GetTargetSubscriptionRef returns the TargetSubscriptionRef field if non-nil, zero value otherwise.

### GetTargetSubscriptionRefOk

`func (o *CreateQuoteRequest) GetTargetSubscriptionRefOk() (*map[string]interface{}, bool)`

GetTargetSubscriptionRefOk returns a tuple with the TargetSubscriptionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSubscriptionRef

`func (o *CreateQuoteRequest) SetTargetSubscriptionRef(v map[string]interface{})`

SetTargetSubscriptionRef sets TargetSubscriptionRef field to given value.

### HasTargetSubscriptionRef

`func (o *CreateQuoteRequest) HasTargetSubscriptionRef() bool`

HasTargetSubscriptionRef returns a boolean if a field has been set.

### GetCurrentContractSnapshotDigest

`func (o *CreateQuoteRequest) GetCurrentContractSnapshotDigest() string`

GetCurrentContractSnapshotDigest returns the CurrentContractSnapshotDigest field if non-nil, zero value otherwise.

### GetCurrentContractSnapshotDigestOk

`func (o *CreateQuoteRequest) GetCurrentContractSnapshotDigestOk() (*string, bool)`

GetCurrentContractSnapshotDigestOk returns a tuple with the CurrentContractSnapshotDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentContractSnapshotDigest

`func (o *CreateQuoteRequest) SetCurrentContractSnapshotDigest(v string)`

SetCurrentContractSnapshotDigest sets CurrentContractSnapshotDigest field to given value.

### HasCurrentContractSnapshotDigest

`func (o *CreateQuoteRequest) HasCurrentContractSnapshotDigest() bool`

HasCurrentContractSnapshotDigest returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *CreateQuoteRequest) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CreateQuoteRequest) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CreateQuoteRequest) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *CreateQuoteRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetCoTermBehavior

`func (o *CreateQuoteRequest) GetCoTermBehavior() string`

GetCoTermBehavior returns the CoTermBehavior field if non-nil, zero value otherwise.

### GetCoTermBehaviorOk

`func (o *CreateQuoteRequest) GetCoTermBehaviorOk() (*string, bool)`

GetCoTermBehaviorOk returns a tuple with the CoTermBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoTermBehavior

`func (o *CreateQuoteRequest) SetCoTermBehavior(v string)`

SetCoTermBehavior sets CoTermBehavior field to given value.

### HasCoTermBehavior

`func (o *CreateQuoteRequest) HasCoTermBehavior() bool`

HasCoTermBehavior returns a boolean if a field has been set.

### GetProrationBasis

`func (o *CreateQuoteRequest) GetProrationBasis() map[string]interface{}`

GetProrationBasis returns the ProrationBasis field if non-nil, zero value otherwise.

### GetProrationBasisOk

`func (o *CreateQuoteRequest) GetProrationBasisOk() (*map[string]interface{}, bool)`

GetProrationBasisOk returns a tuple with the ProrationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBasis

`func (o *CreateQuoteRequest) SetProrationBasis(v map[string]interface{})`

SetProrationBasis sets ProrationBasis field to given value.

### HasProrationBasis

`func (o *CreateQuoteRequest) HasProrationBasis() bool`

HasProrationBasis returns a boolean if a field has been set.

### GetChangeReason

`func (o *CreateQuoteRequest) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *CreateQuoteRequest) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *CreateQuoteRequest) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.

### HasChangeReason

`func (o *CreateQuoteRequest) HasChangeReason() bool`

HasChangeReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


