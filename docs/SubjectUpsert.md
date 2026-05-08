# SubjectUpsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique, human-readable identifier for the subject. This is typically a database ID or a customer key. | 
**DisplayName** | Pointer to **string** | A human-readable display name for the subject. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata for the subject. | [optional] 
**CurrentPeriodStart** | Pointer to **time.Time** | The start of the current period for the subject. | [optional] 
**CurrentPeriodEnd** | Pointer to **time.Time** | The end of the current period for the subject. | [optional] 
**StripeCustomerId** | Pointer to **string** | The Stripe customer ID for the subject. | [optional] 

## Methods

### NewSubjectUpsert

`func NewSubjectUpsert(key string, ) *SubjectUpsert`

NewSubjectUpsert instantiates a new SubjectUpsert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectUpsertWithDefaults

`func NewSubjectUpsertWithDefaults() *SubjectUpsert`

NewSubjectUpsertWithDefaults instantiates a new SubjectUpsert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SubjectUpsert) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SubjectUpsert) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SubjectUpsert) SetKey(v string)`

SetKey sets Key field to given value.


### GetDisplayName

`func (o *SubjectUpsert) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubjectUpsert) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubjectUpsert) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubjectUpsert) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMetadata

`func (o *SubjectUpsert) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubjectUpsert) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubjectUpsert) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubjectUpsert) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *SubjectUpsert) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *SubjectUpsert) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *SubjectUpsert) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *SubjectUpsert) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *SubjectUpsert) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubjectUpsert) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubjectUpsert) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *SubjectUpsert) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *SubjectUpsert) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *SubjectUpsert) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *SubjectUpsert) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *SubjectUpsert) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


