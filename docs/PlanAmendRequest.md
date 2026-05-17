# PlanAmendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAmountMicros** | **int32** |  | 
**InstallmentCount** | **int32** |  | 
**Frequency** | Pointer to **string** |  | [optional] [default to "monthly"]
**FirstDueDate** | **string** |  | 
**ProposedBy** | Pointer to **string** |  | [optional] [default to "human"]
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]
**GracePeriodDays** | Pointer to **int32** |  | [optional] [default to 3]
**MaxMissedBeforeDefault** | Pointer to **int32** |  | [optional] [default to 2]

## Methods

### NewPlanAmendRequest

`func NewPlanAmendRequest(totalAmountMicros int32, installmentCount int32, firstDueDate string, ) *PlanAmendRequest`

NewPlanAmendRequest instantiates a new PlanAmendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanAmendRequestWithDefaults

`func NewPlanAmendRequestWithDefaults() *PlanAmendRequest`

NewPlanAmendRequestWithDefaults instantiates a new PlanAmendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAmountMicros

`func (o *PlanAmendRequest) GetTotalAmountMicros() int32`

GetTotalAmountMicros returns the TotalAmountMicros field if non-nil, zero value otherwise.

### GetTotalAmountMicrosOk

`func (o *PlanAmendRequest) GetTotalAmountMicrosOk() (*int32, bool)`

GetTotalAmountMicrosOk returns a tuple with the TotalAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMicros

`func (o *PlanAmendRequest) SetTotalAmountMicros(v int32)`

SetTotalAmountMicros sets TotalAmountMicros field to given value.


### GetInstallmentCount

`func (o *PlanAmendRequest) GetInstallmentCount() int32`

GetInstallmentCount returns the InstallmentCount field if non-nil, zero value otherwise.

### GetInstallmentCountOk

`func (o *PlanAmendRequest) GetInstallmentCountOk() (*int32, bool)`

GetInstallmentCountOk returns a tuple with the InstallmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentCount

`func (o *PlanAmendRequest) SetInstallmentCount(v int32)`

SetInstallmentCount sets InstallmentCount field to given value.


### GetFrequency

`func (o *PlanAmendRequest) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PlanAmendRequest) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PlanAmendRequest) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PlanAmendRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFirstDueDate

`func (o *PlanAmendRequest) GetFirstDueDate() string`

GetFirstDueDate returns the FirstDueDate field if non-nil, zero value otherwise.

### GetFirstDueDateOk

`func (o *PlanAmendRequest) GetFirstDueDateOk() (*string, bool)`

GetFirstDueDateOk returns a tuple with the FirstDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDueDate

`func (o *PlanAmendRequest) SetFirstDueDate(v string)`

SetFirstDueDate sets FirstDueDate field to given value.


### GetProposedBy

`func (o *PlanAmendRequest) GetProposedBy() string`

GetProposedBy returns the ProposedBy field if non-nil, zero value otherwise.

### GetProposedByOk

`func (o *PlanAmendRequest) GetProposedByOk() (*string, bool)`

GetProposedByOk returns a tuple with the ProposedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedBy

`func (o *PlanAmendRequest) SetProposedBy(v string)`

SetProposedBy sets ProposedBy field to given value.

### HasProposedBy

`func (o *PlanAmendRequest) HasProposedBy() bool`

HasProposedBy returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PlanAmendRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PlanAmendRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PlanAmendRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PlanAmendRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetGracePeriodDays

`func (o *PlanAmendRequest) GetGracePeriodDays() int32`

GetGracePeriodDays returns the GracePeriodDays field if non-nil, zero value otherwise.

### GetGracePeriodDaysOk

`func (o *PlanAmendRequest) GetGracePeriodDaysOk() (*int32, bool)`

GetGracePeriodDaysOk returns a tuple with the GracePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodDays

`func (o *PlanAmendRequest) SetGracePeriodDays(v int32)`

SetGracePeriodDays sets GracePeriodDays field to given value.

### HasGracePeriodDays

`func (o *PlanAmendRequest) HasGracePeriodDays() bool`

HasGracePeriodDays returns a boolean if a field has been set.

### GetMaxMissedBeforeDefault

`func (o *PlanAmendRequest) GetMaxMissedBeforeDefault() int32`

GetMaxMissedBeforeDefault returns the MaxMissedBeforeDefault field if non-nil, zero value otherwise.

### GetMaxMissedBeforeDefaultOk

`func (o *PlanAmendRequest) GetMaxMissedBeforeDefaultOk() (*int32, bool)`

GetMaxMissedBeforeDefaultOk returns a tuple with the MaxMissedBeforeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMissedBeforeDefault

`func (o *PlanAmendRequest) SetMaxMissedBeforeDefault(v int32)`

SetMaxMissedBeforeDefault sets MaxMissedBeforeDefault field to given value.

### HasMaxMissedBeforeDefault

`func (o *PlanAmendRequest) HasMaxMissedBeforeDefault() bool`

HasMaxMissedBeforeDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


