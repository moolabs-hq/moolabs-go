# EntitlementStaticV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Config** | **string** | The JSON parsable config of the entitlement. This value is also returned when checking entitlement access and it is useful for configuring fine-grained access settings to the feature, implemented in your own system. Has to be an object. | 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the feature. | [optional] 
**ActiveFrom** | **time.Time** | The cadence start of the resource. | 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | The annotations of the entitlement. | [optional] [readonly] 
**Id** | **string** | Readonly unique ULID identifier. | [readonly] 
**FeatureKey** | **string** | The feature the subject is entitled to use. | 
**FeatureId** | **string** | The feature the subject is entitled to use. | 
**CurrentUsagePeriod** | Pointer to [**Period**](Period.md) | The current usage period. | [optional] 
**UsagePeriod** | Pointer to [**RecurringPeriod**](RecurringPeriod.md) | The defined usage period of the entitlement | [optional] 
**CustomerKey** | Pointer to **string** | The identifier key unique to the customer | [optional] 
**CustomerId** | **string** | The identifier unique to the customer | 

## Methods

### NewEntitlementStaticV2

`func NewEntitlementStaticV2(type_ string, config string, createdAt time.Time, updatedAt time.Time, activeFrom time.Time, id string, featureKey string, featureId string, customerId string, ) *EntitlementStaticV2`

NewEntitlementStaticV2 instantiates a new EntitlementStaticV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementStaticV2WithDefaults

`func NewEntitlementStaticV2WithDefaults() *EntitlementStaticV2`

NewEntitlementStaticV2WithDefaults instantiates a new EntitlementStaticV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntitlementStaticV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementStaticV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementStaticV2) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *EntitlementStaticV2) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EntitlementStaticV2) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EntitlementStaticV2) SetConfig(v string)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *EntitlementStaticV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntitlementStaticV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntitlementStaticV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EntitlementStaticV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntitlementStaticV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntitlementStaticV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EntitlementStaticV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EntitlementStaticV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EntitlementStaticV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EntitlementStaticV2) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementStaticV2) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementStaticV2) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementStaticV2) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementStaticV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActiveFrom

`func (o *EntitlementStaticV2) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *EntitlementStaticV2) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *EntitlementStaticV2) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *EntitlementStaticV2) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *EntitlementStaticV2) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *EntitlementStaticV2) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *EntitlementStaticV2) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetAnnotations

`func (o *EntitlementStaticV2) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *EntitlementStaticV2) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *EntitlementStaticV2) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *EntitlementStaticV2) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetId

`func (o *EntitlementStaticV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementStaticV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementStaticV2) SetId(v string)`

SetId sets Id field to given value.


### GetFeatureKey

`func (o *EntitlementStaticV2) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *EntitlementStaticV2) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *EntitlementStaticV2) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.


### GetFeatureId

`func (o *EntitlementStaticV2) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *EntitlementStaticV2) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *EntitlementStaticV2) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetCurrentUsagePeriod

`func (o *EntitlementStaticV2) GetCurrentUsagePeriod() Period`

GetCurrentUsagePeriod returns the CurrentUsagePeriod field if non-nil, zero value otherwise.

### GetCurrentUsagePeriodOk

`func (o *EntitlementStaticV2) GetCurrentUsagePeriodOk() (*Period, bool)`

GetCurrentUsagePeriodOk returns a tuple with the CurrentUsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsagePeriod

`func (o *EntitlementStaticV2) SetCurrentUsagePeriod(v Period)`

SetCurrentUsagePeriod sets CurrentUsagePeriod field to given value.

### HasCurrentUsagePeriod

`func (o *EntitlementStaticV2) HasCurrentUsagePeriod() bool`

HasCurrentUsagePeriod returns a boolean if a field has been set.

### GetUsagePeriod

`func (o *EntitlementStaticV2) GetUsagePeriod() RecurringPeriod`

GetUsagePeriod returns the UsagePeriod field if non-nil, zero value otherwise.

### GetUsagePeriodOk

`func (o *EntitlementStaticV2) GetUsagePeriodOk() (*RecurringPeriod, bool)`

GetUsagePeriodOk returns a tuple with the UsagePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePeriod

`func (o *EntitlementStaticV2) SetUsagePeriod(v RecurringPeriod)`

SetUsagePeriod sets UsagePeriod field to given value.

### HasUsagePeriod

`func (o *EntitlementStaticV2) HasUsagePeriod() bool`

HasUsagePeriod returns a boolean if a field has been set.

### GetCustomerKey

`func (o *EntitlementStaticV2) GetCustomerKey() string`

GetCustomerKey returns the CustomerKey field if non-nil, zero value otherwise.

### GetCustomerKeyOk

`func (o *EntitlementStaticV2) GetCustomerKeyOk() (*string, bool)`

GetCustomerKeyOk returns a tuple with the CustomerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerKey

`func (o *EntitlementStaticV2) SetCustomerKey(v string)`

SetCustomerKey sets CustomerKey field to given value.

### HasCustomerKey

`func (o *EntitlementStaticV2) HasCustomerKey() bool`

HasCustomerKey returns a boolean if a field has been set.

### GetCustomerId

`func (o *EntitlementStaticV2) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EntitlementStaticV2) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EntitlementStaticV2) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


