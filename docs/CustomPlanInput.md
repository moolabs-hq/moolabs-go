# CustomPlanInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name for the resource. Between 1 and 256 characters. | 
**Description** | Pointer to **string** | Optional description of the resource. Maximum 1024 characters. | [optional] 
**Metadata** | Pointer to **interface{}** | Additional metadata for the resource. | [optional] 
**Alignment** | Pointer to [**Alignment**](Alignment.md) | Alignment configuration for the plan. | [optional] 
**Currency** | **string** | The currency code of the plan. | 
**BillingCadence** | **string** | The default billing cadence for subscriptions using this plan. Defines how often customers are billed using ISO8601 duration format. Examples: \&quot;P1M\&quot; (monthly), \&quot;P3M\&quot; (quarterly), \&quot;P1Y\&quot; (annually). | 
**ProRatingConfig** | Pointer to [**ProRatingConfig**](ProRatingConfig.md) | Default pro-rating configuration for subscriptions using this plan. | [optional] 
**CreditConfig** | Pointer to [**CreditConfig**](CreditConfig.md) | Default credit allowance configuration for subscriptions using this plan (first-class column). | [optional] 
**Phases** | [**[]PlanPhase**](PlanPhase.md) | The plan phase or pricing ramp allows changing a plan&#39;s rate cards over time as a subscription progresses. A phase switch occurs only at the end of a billing period, ensuring that a single subscription invoice will not include charges from different phase prices. | 

## Methods

### NewCustomPlanInput

`func NewCustomPlanInput(name string, currency string, billingCadence string, phases []PlanPhase, ) *CustomPlanInput`

NewCustomPlanInput instantiates a new CustomPlanInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPlanInputWithDefaults

`func NewCustomPlanInputWithDefaults() *CustomPlanInput`

NewCustomPlanInputWithDefaults instantiates a new CustomPlanInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomPlanInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPlanInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPlanInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomPlanInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomPlanInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomPlanInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomPlanInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomPlanInput) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomPlanInput) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomPlanInput) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomPlanInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CustomPlanInput) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CustomPlanInput) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAlignment

`func (o *CustomPlanInput) GetAlignment() Alignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *CustomPlanInput) GetAlignmentOk() (*Alignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *CustomPlanInput) SetAlignment(v Alignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *CustomPlanInput) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetCurrency

`func (o *CustomPlanInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomPlanInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomPlanInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBillingCadence

`func (o *CustomPlanInput) GetBillingCadence() string`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *CustomPlanInput) GetBillingCadenceOk() (*string, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *CustomPlanInput) SetBillingCadence(v string)`

SetBillingCadence sets BillingCadence field to given value.


### GetProRatingConfig

`func (o *CustomPlanInput) GetProRatingConfig() ProRatingConfig`

GetProRatingConfig returns the ProRatingConfig field if non-nil, zero value otherwise.

### GetProRatingConfigOk

`func (o *CustomPlanInput) GetProRatingConfigOk() (*ProRatingConfig, bool)`

GetProRatingConfigOk returns a tuple with the ProRatingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProRatingConfig

`func (o *CustomPlanInput) SetProRatingConfig(v ProRatingConfig)`

SetProRatingConfig sets ProRatingConfig field to given value.

### HasProRatingConfig

`func (o *CustomPlanInput) HasProRatingConfig() bool`

HasProRatingConfig returns a boolean if a field has been set.

### GetCreditConfig

`func (o *CustomPlanInput) GetCreditConfig() CreditConfig`

GetCreditConfig returns the CreditConfig field if non-nil, zero value otherwise.

### GetCreditConfigOk

`func (o *CustomPlanInput) GetCreditConfigOk() (*CreditConfig, bool)`

GetCreditConfigOk returns a tuple with the CreditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditConfig

`func (o *CustomPlanInput) SetCreditConfig(v CreditConfig)`

SetCreditConfig sets CreditConfig field to given value.

### HasCreditConfig

`func (o *CustomPlanInput) HasCreditConfig() bool`

HasCreditConfig returns a boolean if a field has been set.

### GetPhases

`func (o *CustomPlanInput) GetPhases() []PlanPhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *CustomPlanInput) GetPhasesOk() (*[]PlanPhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *CustomPlanInput) SetPhases(v []PlanPhase)`

SetPhases sets Phases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


