# PlanCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional metadata for the resource. | [optional] 
**Key** | **string** | A semi-unique identifier for the resource. | 
**Alignment** | Pointer to [**Alignment**](Alignment.md) | Alignment configuration for the plan. | [optional] 
**Currency** | **string** | The currency code of the plan. | 
**BillingCadence** | **string** | The default billing cadence for subscriptions using this plan. Defines how often customers are billed using ISO8601 duration format. Examples: \&quot;P1M\&quot; (monthly), \&quot;P3M\&quot; (quarterly), \&quot;P1Y\&quot; (annually). | 
**ProRatingConfig** | Pointer to [**ProRatingConfig**](ProRatingConfig.md) | Default pro-rating configuration for subscriptions using this plan. | [optional] 
**CreditConfig** | Pointer to [**CreditConfig**](CreditConfig.md) | Default credit allowance configuration for subscriptions using this plan (first-class column). | [optional] 
**Phases** | [**[]PlanPhase**](PlanPhase.md) | The plan phase or pricing ramp allows changing a plan&#39;s rate cards over time as a subscription progresses. A phase switch occurs only at the end of a billing period, ensuring that a single subscription invoice will not include charges from different phase prices. | 

## Methods

### NewPlanCreate

`func NewPlanCreate(name string, key string, currency string, billingCadence string, phases []PlanPhase, ) *PlanCreate`

NewPlanCreate instantiates a new PlanCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCreateWithDefaults

`func NewPlanCreateWithDefaults() *PlanCreate`

NewPlanCreateWithDefaults instantiates a new PlanCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlanCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlanCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *PlanCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlanCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlanCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlanCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetKey

`func (o *PlanCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlanCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlanCreate) SetKey(v string)`

SetKey sets Key field to given value.


### GetAlignment

`func (o *PlanCreate) GetAlignment() Alignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *PlanCreate) GetAlignmentOk() (*Alignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *PlanCreate) SetAlignment(v Alignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *PlanCreate) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetCurrency

`func (o *PlanCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PlanCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PlanCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBillingCadence

`func (o *PlanCreate) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *PlanCreate) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *PlanCreate) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetProRatingConfig

`func (o *PlanCreate) GetProRatingConfig() ProRatingConfig`

GetProRatingConfig returns the ProRatingConfig field if non-nil, zero value otherwise.

### GetProRatingConfigOk

`func (o *PlanCreate) GetProRatingConfigOk() (*ProRatingConfig, bool)`

GetProRatingConfigOk returns a tuple with the ProRatingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProRatingConfig

`func (o *PlanCreate) SetProRatingConfig(v ProRatingConfig)`

SetProRatingConfig sets ProRatingConfig field to given value.

### HasProRatingConfig

`func (o *PlanCreate) HasProRatingConfig() bool`

HasProRatingConfig returns a boolean if a field has been set.

### GetCreditConfig

`func (o *PlanCreate) GetCreditConfig() CreditConfig`

GetCreditConfig returns the CreditConfig field if non-nil, zero value otherwise.

### GetCreditConfigOk

`func (o *PlanCreate) GetCreditConfigOk() (*CreditConfig, bool)`

GetCreditConfigOk returns a tuple with the CreditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditConfig

`func (o *PlanCreate) SetCreditConfig(v CreditConfig)`

SetCreditConfig sets CreditConfig field to given value.

### HasCreditConfig

`func (o *PlanCreate) HasCreditConfig() bool`

HasCreditConfig returns a boolean if a field has been set.

### GetPhases

`func (o *PlanCreate) GetPhases() []PlanPhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *PlanCreate) GetPhasesOk() (*[]PlanPhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *PlanCreate) SetPhases(v []PlanPhase)`

SetPhases sets Phases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


