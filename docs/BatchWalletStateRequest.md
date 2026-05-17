# BatchWalletStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallets** | [**[]BatchWalletStateRequestItem**](BatchWalletStateRequestItem.md) | List of wallet identifiers | 
**EffectiveAsOf** | Pointer to **time.Time** | Effective as-of timestamp (business time) for time travel | [optional] 
**RecordedAsOf** | Pointer to **time.Time** | Recorded as-of timestamp (system time) for time travel | [optional] 
**ConsistentView** | Pointer to **bool** | Use strong consistency for reads | [optional] [default to false]
**Consistency** | Pointer to **string** | Consistency level: &#39;eventual&#39; or &#39;STRONG&#39; | [optional] [default to "eventual"]

## Methods

### NewBatchWalletStateRequest

`func NewBatchWalletStateRequest(wallets []BatchWalletStateRequestItem, ) *BatchWalletStateRequest`

NewBatchWalletStateRequest instantiates a new BatchWalletStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWalletStateRequestWithDefaults

`func NewBatchWalletStateRequestWithDefaults() *BatchWalletStateRequest`

NewBatchWalletStateRequestWithDefaults instantiates a new BatchWalletStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallets

`func (o *BatchWalletStateRequest) GetWallets() []BatchWalletStateRequestItem`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *BatchWalletStateRequest) GetWalletsOk() (*[]BatchWalletStateRequestItem, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *BatchWalletStateRequest) SetWallets(v []BatchWalletStateRequestItem)`

SetWallets sets Wallets field to given value.


### GetEffectiveAsOf

`func (o *BatchWalletStateRequest) GetEffectiveAsOf() time.Time`

GetEffectiveAsOf returns the EffectiveAsOf field if non-nil, zero value otherwise.

### GetEffectiveAsOfOk

`func (o *BatchWalletStateRequest) GetEffectiveAsOfOk() (*time.Time, bool)`

GetEffectiveAsOfOk returns a tuple with the EffectiveAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAsOf

`func (o *BatchWalletStateRequest) SetEffectiveAsOf(v time.Time)`

SetEffectiveAsOf sets EffectiveAsOf field to given value.

### HasEffectiveAsOf

`func (o *BatchWalletStateRequest) HasEffectiveAsOf() bool`

HasEffectiveAsOf returns a boolean if a field has been set.

### GetRecordedAsOf

`func (o *BatchWalletStateRequest) GetRecordedAsOf() time.Time`

GetRecordedAsOf returns the RecordedAsOf field if non-nil, zero value otherwise.

### GetRecordedAsOfOk

`func (o *BatchWalletStateRequest) GetRecordedAsOfOk() (*time.Time, bool)`

GetRecordedAsOfOk returns a tuple with the RecordedAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAsOf

`func (o *BatchWalletStateRequest) SetRecordedAsOf(v time.Time)`

SetRecordedAsOf sets RecordedAsOf field to given value.

### HasRecordedAsOf

`func (o *BatchWalletStateRequest) HasRecordedAsOf() bool`

HasRecordedAsOf returns a boolean if a field has been set.

### GetConsistentView

`func (o *BatchWalletStateRequest) GetConsistentView() bool`

GetConsistentView returns the ConsistentView field if non-nil, zero value otherwise.

### GetConsistentViewOk

`func (o *BatchWalletStateRequest) GetConsistentViewOk() (*bool, bool)`

GetConsistentViewOk returns a tuple with the ConsistentView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistentView

`func (o *BatchWalletStateRequest) SetConsistentView(v bool)`

SetConsistentView sets ConsistentView field to given value.

### HasConsistentView

`func (o *BatchWalletStateRequest) HasConsistentView() bool`

HasConsistentView returns a boolean if a field has been set.

### GetConsistency

`func (o *BatchWalletStateRequest) GetConsistency() string`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *BatchWalletStateRequest) GetConsistencyOk() (*string, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *BatchWalletStateRequest) SetConsistency(v string)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *BatchWalletStateRequest) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


