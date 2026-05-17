# SubscriptionActivateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionData** | **map[string]interface{}** | Raw OpenMeter subscription response | 
**CustomerId** | Pointer to **string** | Customer ID | [optional] 
**CustomerKey** | Pointer to **string** | Customer key (for wallet binding) | [optional] 
**SubjectKeys** | Pointer to **[]string** | Usage attribution subject keys | [optional] 
**ActiveFrom** | Pointer to **string** | Subscription start date (ISO 8601) | [optional] 
**ActiveTo** | Pointer to **string** | Subscription end date (ISO 8601) | [optional] 
**BillingCadence** | Pointer to **string** | Billing cadence (ISO 8601 duration) | [optional] 
**PoolOverrides** | Pointer to **map[string]float32** | Per-pool recurring credit overrides: pool_key -&gt; amount. &#39;__global__&#39; overrides global pool. | [optional] 

## Methods

### NewSubscriptionActivateRequest

`func NewSubscriptionActivateRequest(subscriptionData map[string]interface{}, ) *SubscriptionActivateRequest`

NewSubscriptionActivateRequest instantiates a new SubscriptionActivateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionActivateRequestWithDefaults

`func NewSubscriptionActivateRequestWithDefaults() *SubscriptionActivateRequest`

NewSubscriptionActivateRequestWithDefaults instantiates a new SubscriptionActivateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionData

`func (o *SubscriptionActivateRequest) GetSubscriptionData() map[string]interface{}`

GetSubscriptionData returns the SubscriptionData field if non-nil, zero value otherwise.

### GetSubscriptionDataOk

`func (o *SubscriptionActivateRequest) GetSubscriptionDataOk() (*map[string]interface{}, bool)`

GetSubscriptionDataOk returns a tuple with the SubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionData

`func (o *SubscriptionActivateRequest) SetSubscriptionData(v map[string]interface{})`

SetSubscriptionData sets SubscriptionData field to given value.


### GetCustomerId

`func (o *SubscriptionActivateRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionActivateRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionActivateRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionActivateRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerKey

`func (o *SubscriptionActivateRequest) GetCustomerKey() string`

GetCustomerKey returns the CustomerKey field if non-nil, zero value otherwise.

### GetCustomerKeyOk

`func (o *SubscriptionActivateRequest) GetCustomerKeyOk() (*string, bool)`

GetCustomerKeyOk returns a tuple with the CustomerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerKey

`func (o *SubscriptionActivateRequest) SetCustomerKey(v string)`

SetCustomerKey sets CustomerKey field to given value.

### HasCustomerKey

`func (o *SubscriptionActivateRequest) HasCustomerKey() bool`

HasCustomerKey returns a boolean if a field has been set.

### GetSubjectKeys

`func (o *SubscriptionActivateRequest) GetSubjectKeys() []string`

GetSubjectKeys returns the SubjectKeys field if non-nil, zero value otherwise.

### GetSubjectKeysOk

`func (o *SubscriptionActivateRequest) GetSubjectKeysOk() (*[]string, bool)`

GetSubjectKeysOk returns a tuple with the SubjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectKeys

`func (o *SubscriptionActivateRequest) SetSubjectKeys(v []string)`

SetSubjectKeys sets SubjectKeys field to given value.

### HasSubjectKeys

`func (o *SubscriptionActivateRequest) HasSubjectKeys() bool`

HasSubjectKeys returns a boolean if a field has been set.

### GetActiveFrom

`func (o *SubscriptionActivateRequest) GetActiveFrom() string`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionActivateRequest) GetActiveFromOk() (*string, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionActivateRequest) SetActiveFrom(v string)`

SetActiveFrom sets ActiveFrom field to given value.

### HasActiveFrom

`func (o *SubscriptionActivateRequest) HasActiveFrom() bool`

HasActiveFrom returns a boolean if a field has been set.

### GetActiveTo

`func (o *SubscriptionActivateRequest) GetActiveTo() string`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionActivateRequest) GetActiveToOk() (*string, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionActivateRequest) SetActiveTo(v string)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionActivateRequest) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetBillingCadence

`func (o *SubscriptionActivateRequest) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *SubscriptionActivateRequest) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *SubscriptionActivateRequest) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *SubscriptionActivateRequest) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetPoolOverrides

`func (o *SubscriptionActivateRequest) GetPoolOverrides() map[string]float32`

GetPoolOverrides returns the PoolOverrides field if non-nil, zero value otherwise.

### GetPoolOverridesOk

`func (o *SubscriptionActivateRequest) GetPoolOverridesOk() (*map[string]float32, bool)`

GetPoolOverridesOk returns a tuple with the PoolOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOverrides

`func (o *SubscriptionActivateRequest) SetPoolOverrides(v map[string]float32)`

SetPoolOverrides sets PoolOverrides field to given value.

### HasPoolOverrides

`func (o *SubscriptionActivateRequest) HasPoolOverrides() bool`

HasPoolOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


