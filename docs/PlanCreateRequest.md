# PlanCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAmountMicros** | **int32** |  | 
**InstallmentCount** | **int32** |  | 
**Frequency** | Pointer to **string** |  | [optional] [default to "monthly"]
**FirstDueDate** | **string** |  | 
**ProposedBy** | **string** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**GracePeriodDays** | Pointer to **int32** |  | [optional] [default to 3]
**MaxMissedBeforeDefault** | Pointer to **int32** |  | [optional] [default to 2]

## Methods

### NewPlanCreateRequest

`func NewPlanCreateRequest(totalAmountMicros int32, installmentCount int32, firstDueDate string, proposedBy string, ) *PlanCreateRequest`

NewPlanCreateRequest instantiates a new PlanCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanCreateRequestWithDefaults

`func NewPlanCreateRequestWithDefaults() *PlanCreateRequest`

NewPlanCreateRequestWithDefaults instantiates a new PlanCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAmountMicros

`func (o *PlanCreateRequest) GetTotalAmountMicros() int32`

GetTotalAmountMicros returns the TotalAmountMicros field if non-nil, zero value otherwise.

### GetTotalAmountMicrosOk

`func (o *PlanCreateRequest) GetTotalAmountMicrosOk() (*int32, bool)`

GetTotalAmountMicrosOk returns a tuple with the TotalAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMicros

`func (o *PlanCreateRequest) SetTotalAmountMicros(v int32)`

SetTotalAmountMicros sets TotalAmountMicros field to given value.


### GetInstallmentCount

`func (o *PlanCreateRequest) GetInstallmentCount() int32`

GetInstallmentCount returns the InstallmentCount field if non-nil, zero value otherwise.

### GetInstallmentCountOk

`func (o *PlanCreateRequest) GetInstallmentCountOk() (*int32, bool)`

GetInstallmentCountOk returns a tuple with the InstallmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentCount

`func (o *PlanCreateRequest) SetInstallmentCount(v int32)`

SetInstallmentCount sets InstallmentCount field to given value.


### GetFrequency

`func (o *PlanCreateRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PlanCreateRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PlanCreateRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PlanCreateRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFirstDueDate

`func (o *PlanCreateRequest) GetFirstDueDate() string`

GetFirstDueDate returns the FirstDueDate field if non-nil, zero value otherwise.

### GetFirstDueDateOk

`func (o *PlanCreateRequest) GetFirstDueDateOk() (*string, bool)`

GetFirstDueDateOk returns a tuple with the FirstDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDueDate

`func (o *PlanCreateRequest) SetFirstDueDate(v string)`

SetFirstDueDate sets FirstDueDate field to given value.


### GetProposedBy

`func (o *PlanCreateRequest) GetProposedBy() string`

GetProposedBy returns the ProposedBy field if non-nil, zero value otherwise.

### GetProposedByOk

`func (o *PlanCreateRequest) GetProposedByOk() (*string, bool)`

GetProposedByOk returns a tuple with the ProposedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedBy

`func (o *PlanCreateRequest) SetProposedBy(v string)`

SetProposedBy sets ProposedBy field to given value.


### GetCurrencyCode

`func (o *PlanCreateRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PlanCreateRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PlanCreateRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PlanCreateRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetGracePeriodDays

`func (o *PlanCreateRequest) GetGracePeriodDays() int32`

GetGracePeriodDays returns the GracePeriodDays field if non-nil, zero value otherwise.

### GetGracePeriodDaysOk

`func (o *PlanCreateRequest) GetGracePeriodDaysOk() (*int32, bool)`

GetGracePeriodDaysOk returns a tuple with the GracePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodDays

`func (o *PlanCreateRequest) SetGracePeriodDays(v int32)`

SetGracePeriodDays sets GracePeriodDays field to given value.

### HasGracePeriodDays

`func (o *PlanCreateRequest) HasGracePeriodDays() bool`

HasGracePeriodDays returns a boolean if a field has been set.

### GetMaxMissedBeforeDefault

`func (o *PlanCreateRequest) GetMaxMissedBeforeDefault() int32`

GetMaxMissedBeforeDefault returns the MaxMissedBeforeDefault field if non-nil, zero value otherwise.

### GetMaxMissedBeforeDefaultOk

`func (o *PlanCreateRequest) GetMaxMissedBeforeDefaultOk() (*int32, bool)`

GetMaxMissedBeforeDefaultOk returns a tuple with the MaxMissedBeforeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMissedBeforeDefault

`func (o *PlanCreateRequest) SetMaxMissedBeforeDefault(v int32)`

SetMaxMissedBeforeDefault sets MaxMissedBeforeDefault field to given value.

### HasMaxMissedBeforeDefault

`func (o *PlanCreateRequest) HasMaxMissedBeforeDefault() bool`

HasMaxMissedBeforeDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


