# SubscriptionExpanded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the resource. | [readonly] 
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**DeletedAt** | Pointer to **time.Time** | Timestamp of when the resource was permanently deleted. | [optional] [readonly] 
**ActiveFrom** | **time.Time** | The cadence start of the resource. | 
**ActiveTo** | Pointer to **time.Time** | The cadence end of the resource. | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Set of key-value pairs managed by the system. Cannot be modified by user. | [optional] [readonly] 
**Status** | [**SubscriptionStatus**](SubscriptionStatus.md) | The status of the subscription. | [readonly] 
**CustomerId** | **string** | The customer ID of the subscription. | 
**Plan** | Pointer to [**PlanReference**](PlanReference.md) | The plan of the subscription. | [optional] 
**Currency** | **string** | The currency code of the subscription. Will be revised once we add multi currency support. | 
**BillingCadence** | **string** | The billing cadence for the subscriptions. Defines how often customers are billed using ISO8601 duration format. Examples: \&quot;P1M\&quot; (monthly), \&quot;P3M\&quot; (quarterly), \&quot;P1Y\&quot; (annually). | [readonly] 
**ProRatingConfig** | Pointer to [**ProRatingConfig**](ProRatingConfig.md) | The pro-rating configuration for the subscriptions. | [optional] [readonly] 
**BillingAnchor** | **time.Time** | The normalizedbilling anchor of the subscription. | [readonly] 
**CommercialOverrides** | Pointer to [**CommercialOverrides**](CommercialOverrides.md) | Commercial terms for this subscription (discounts, pool overrides, wallet policy). | [optional] [readonly] 
**Alignment** | Pointer to [**SubscriptionAlignment**](SubscriptionAlignment.md) | Alignment details enriched with the current billing period. | [optional] 
**Phases** | [**[]SubscriptionPhaseExpanded**](SubscriptionPhaseExpanded.md) | The phases of the subscription. | 

## Methods

### NewSubscriptionExpanded

`func NewSubscriptionExpanded(id string, name string, createdAt time.Time, updatedAt time.Time, activeFrom time.Time, status SubscriptionStatus, customerId string, currency string, billingCadence string, billingAnchor time.Time, phases []SubscriptionPhaseExpanded, ) *SubscriptionExpanded`

NewSubscriptionExpanded instantiates a new SubscriptionExpanded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionExpandedWithDefaults

`func NewSubscriptionExpandedWithDefaults() *SubscriptionExpanded`

NewSubscriptionExpandedWithDefaults instantiates a new SubscriptionExpanded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionExpanded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionExpanded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionExpanded) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SubscriptionExpanded) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionExpanded) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionExpanded) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionExpanded) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionExpanded) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionExpanded) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionExpanded) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionExpanded) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionExpanded) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionExpanded) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionExpanded) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionExpanded) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionExpanded) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionExpanded) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionExpanded) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionExpanded) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionExpanded) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *SubscriptionExpanded) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SubscriptionExpanded) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SubscriptionExpanded) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *SubscriptionExpanded) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetActiveFrom

`func (o *SubscriptionExpanded) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *SubscriptionExpanded) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *SubscriptionExpanded) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *SubscriptionExpanded) GetActiveTo() time.Time`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *SubscriptionExpanded) GetActiveToOk() (*time.Time, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *SubscriptionExpanded) SetActiveTo(v time.Time)`

SetActiveTo sets ActiveTo field to given value.

### HasActiveTo

`func (o *SubscriptionExpanded) HasActiveTo() bool`

HasActiveTo returns a boolean if a field has been set.

### GetAnnotations

`func (o *SubscriptionExpanded) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SubscriptionExpanded) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SubscriptionExpanded) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SubscriptionExpanded) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionExpanded) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionExpanded) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionExpanded) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.


### GetCustomerId

`func (o *SubscriptionExpanded) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionExpanded) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionExpanded) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPlan

`func (o *SubscriptionExpanded) GetPlan() PlanReference`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionExpanded) GetPlanOk() (*PlanReference, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionExpanded) SetPlan(v PlanReference)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionExpanded) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionExpanded) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionExpanded) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionExpanded) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBillingCadence

`func (o *SubscriptionExpanded) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *SubscriptionExpanded) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *SubscriptionExpanded) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetProRatingConfig

`func (o *SubscriptionExpanded) GetProRatingConfig() ProRatingConfig`

GetProRatingConfig returns the ProRatingConfig field if non-nil, zero value otherwise.

### GetProRatingConfigOk

`func (o *SubscriptionExpanded) GetProRatingConfigOk() (*ProRatingConfig, bool)`

GetProRatingConfigOk returns a tuple with the ProRatingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProRatingConfig

`func (o *SubscriptionExpanded) SetProRatingConfig(v ProRatingConfig)`

SetProRatingConfig sets ProRatingConfig field to given value.

### HasProRatingConfig

`func (o *SubscriptionExpanded) HasProRatingConfig() bool`

HasProRatingConfig returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *SubscriptionExpanded) GetBillingAnchor() time.Time`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *SubscriptionExpanded) GetBillingAnchorOk() (*time.Time, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *SubscriptionExpanded) SetBillingAnchor(v time.Time)`

SetBillingAnchor sets BillingAnchor field to given value.


### GetCommercialOverrides

`func (o *SubscriptionExpanded) GetCommercialOverrides() CommercialOverrides`

GetCommercialOverrides returns the CommercialOverrides field if non-nil, zero value otherwise.

### GetCommercialOverridesOk

`func (o *SubscriptionExpanded) GetCommercialOverridesOk() (*CommercialOverrides, bool)`

GetCommercialOverridesOk returns a tuple with the CommercialOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialOverrides

`func (o *SubscriptionExpanded) SetCommercialOverrides(v CommercialOverrides)`

SetCommercialOverrides sets CommercialOverrides field to given value.

### HasCommercialOverrides

`func (o *SubscriptionExpanded) HasCommercialOverrides() bool`

HasCommercialOverrides returns a boolean if a field has been set.

### GetAlignment

`func (o *SubscriptionExpanded) GetAlignment() SubscriptionAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *SubscriptionExpanded) GetAlignmentOk() (*SubscriptionAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *SubscriptionExpanded) SetAlignment(v SubscriptionAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *SubscriptionExpanded) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetPhases

`func (o *SubscriptionExpanded) GetPhases() []SubscriptionPhaseExpanded`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *SubscriptionExpanded) GetPhasesOk() (*[]SubscriptionPhaseExpanded, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *SubscriptionExpanded) SetPhases(v []SubscriptionPhaseExpanded)`

SetPhases sets Phases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


