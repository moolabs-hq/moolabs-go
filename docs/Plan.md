# Plan

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
**Key** | **string** | A semi-unique identifier for the resource. | 
**Alignment** | Pointer to [**Alignment**](Alignment.md) | Alignment configuration for the plan. | [optional] 
**Version** | **int32** | Version of the plan. Incremented when the plan is updated. | [readonly] [default to 1]
**Currency** | **string** | The currency code of the plan. | 
**BillingCadence** | **string** | The default billing cadence for subscriptions using this plan. Defines how often customers are billed using ISO8601 duration format. Examples: \&quot;P1M\&quot; (monthly), \&quot;P3M\&quot; (quarterly), \&quot;P1Y\&quot; (annually). | 
**ProRatingConfig** | Pointer to [**ProRatingConfig**](ProRatingConfig.md) | Default pro-rating configuration for subscriptions using this plan. | [optional] 
**CreditConfig** | Pointer to [**CreditConfig**](CreditConfig.md) | Default credit allowance configuration for subscriptions using this plan (first-class column). | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | The date and time when the plan becomes effective. When not specified, the plan is a draft. | [optional] [readonly] 
**EffectiveTo** | Pointer to **time.Time** | The date and time when the plan is no longer effective. When not specified, the plan is effective indefinitely. | [optional] [readonly] 
**Status** | [**PlanStatus**](PlanStatus.md) | The status of the plan. Computed based on the effective start and end dates: - draft &#x3D; no effectiveFrom - active &#x3D; effectiveFrom &lt;&#x3D; now &lt; effectiveTo - archived / inactive &#x3D; effectiveTo &lt;&#x3D; now - scheduled &#x3D; now &lt; effectiveFrom &lt; effectiveTo | [readonly] 
**Phases** | [**[]PlanPhase**](PlanPhase.md) | The plan phase or pricing ramp allows changing a plan&#39;s rate cards over time as a subscription progresses. A phase switch occurs only at the end of a billing period, ensuring that a single subscription invoice will not include charges from different phase prices. | 
**ValidationErrors** | [**[]MeterValidationError**](MeterValidationError.md) | List of validation errors. | 

## Methods

### NewPlan

`func NewPlan(id string, name string, createdAt time.Time, updatedAt time.Time, key string, version int32, currency string, billingCadence string, status PlanStatus, phases []PlanPhase, validationErrors []MeterValidationError, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Plan) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Plan) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Plan) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Plan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Plan) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plan) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plan) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Plan) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Plan) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Plan) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Plan) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Plan) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Plan) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Plan) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetKey

`func (o *Plan) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Plan) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Plan) SetKey(v string)`

SetKey sets Key field to given value.


### GetAlignment

`func (o *Plan) GetAlignment() Alignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *Plan) GetAlignmentOk() (*Alignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *Plan) SetAlignment(v Alignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *Plan) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetVersion

`func (o *Plan) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Plan) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Plan) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCurrency

`func (o *Plan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Plan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Plan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBillingCadence

`func (o *Plan) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *Plan) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *Plan) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetProRatingConfig

`func (o *Plan) GetProRatingConfig() ProRatingConfig`

GetProRatingConfig returns the ProRatingConfig field if non-nil, zero value otherwise.

### GetProRatingConfigOk

`func (o *Plan) GetProRatingConfigOk() (*ProRatingConfig, bool)`

GetProRatingConfigOk returns a tuple with the ProRatingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProRatingConfig

`func (o *Plan) SetProRatingConfig(v ProRatingConfig)`

SetProRatingConfig sets ProRatingConfig field to given value.

### HasProRatingConfig

`func (o *Plan) HasProRatingConfig() bool`

HasProRatingConfig returns a boolean if a field has been set.

### GetCreditConfig

`func (o *Plan) GetCreditConfig() CreditConfig`

GetCreditConfig returns the CreditConfig field if non-nil, zero value otherwise.

### GetCreditConfigOk

`func (o *Plan) GetCreditConfigOk() (*CreditConfig, bool)`

GetCreditConfigOk returns a tuple with the CreditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditConfig

`func (o *Plan) SetCreditConfig(v CreditConfig)`

SetCreditConfig sets CreditConfig field to given value.

### HasCreditConfig

`func (o *Plan) HasCreditConfig() bool`

HasCreditConfig returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *Plan) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *Plan) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *Plan) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *Plan) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *Plan) GetEffectiveTo() time.Time`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *Plan) GetEffectiveToOk() (*time.Time, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *Plan) SetEffectiveTo(v time.Time)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *Plan) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetStatus

`func (o *Plan) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Plan) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Plan) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.


### GetPhases

`func (o *Plan) GetPhases() []PlanPhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *Plan) GetPhasesOk() (*[]PlanPhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *Plan) SetPhases(v []PlanPhase)`

SetPhases sets Phases field to given value.


### GetValidationErrors

`func (o *Plan) GetValidationErrors() []MeterValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *Plan) GetValidationErrorsOk() (*[]MeterValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *Plan) SetValidationErrors(v []MeterValidationError)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


