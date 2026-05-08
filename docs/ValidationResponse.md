# ValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **string** |  | 
**IsValid** | **bool** |  | 
**BalanceMatches** | **bool** |  | 
**FingerprintMatches** | **bool** |  | 
**StoredBalanceMicros** | **int32** |  | 
**RecalculatedBalanceMicros** | **int32** |  | 
**StoredFingerprint** | **string** |  | 
**RecalculatedFingerprint** | **string** |  | 
**AsOfEffectiveAt** | **string** |  | 
**AsOfRecordedAt** | **string** |  | 

## Methods

### NewValidationResponse

`func NewValidationResponse(snapshotId string, isValid bool, balanceMatches bool, fingerprintMatches bool, storedBalanceMicros int32, recalculatedBalanceMicros int32, storedFingerprint string, recalculatedFingerprint string, asOfEffectiveAt string, asOfRecordedAt string, ) *ValidationResponse`

NewValidationResponse instantiates a new ValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResponseWithDefaults

`func NewValidationResponseWithDefaults() *ValidationResponse`

NewValidationResponseWithDefaults instantiates a new ValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *ValidationResponse) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ValidationResponse) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ValidationResponse) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetIsValid

`func (o *ValidationResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ValidationResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ValidationResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetBalanceMatches

`func (o *ValidationResponse) GetBalanceMatches() bool`

GetBalanceMatches returns the BalanceMatches field if non-nil, zero value otherwise.

### GetBalanceMatchesOk

`func (o *ValidationResponse) GetBalanceMatchesOk() (*bool, bool)`

GetBalanceMatchesOk returns a tuple with the BalanceMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceMatches

`func (o *ValidationResponse) SetBalanceMatches(v bool)`

SetBalanceMatches sets BalanceMatches field to given value.


### GetFingerprintMatches

`func (o *ValidationResponse) GetFingerprintMatches() bool`

GetFingerprintMatches returns the FingerprintMatches field if non-nil, zero value otherwise.

### GetFingerprintMatchesOk

`func (o *ValidationResponse) GetFingerprintMatchesOk() (*bool, bool)`

GetFingerprintMatchesOk returns a tuple with the FingerprintMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintMatches

`func (o *ValidationResponse) SetFingerprintMatches(v bool)`

SetFingerprintMatches sets FingerprintMatches field to given value.


### GetStoredBalanceMicros

`func (o *ValidationResponse) GetStoredBalanceMicros() int32`

GetStoredBalanceMicros returns the StoredBalanceMicros field if non-nil, zero value otherwise.

### GetStoredBalanceMicrosOk

`func (o *ValidationResponse) GetStoredBalanceMicrosOk() (*int32, bool)`

GetStoredBalanceMicrosOk returns a tuple with the StoredBalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredBalanceMicros

`func (o *ValidationResponse) SetStoredBalanceMicros(v int32)`

SetStoredBalanceMicros sets StoredBalanceMicros field to given value.


### GetRecalculatedBalanceMicros

`func (o *ValidationResponse) GetRecalculatedBalanceMicros() int32`

GetRecalculatedBalanceMicros returns the RecalculatedBalanceMicros field if non-nil, zero value otherwise.

### GetRecalculatedBalanceMicrosOk

`func (o *ValidationResponse) GetRecalculatedBalanceMicrosOk() (*int32, bool)`

GetRecalculatedBalanceMicrosOk returns a tuple with the RecalculatedBalanceMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculatedBalanceMicros

`func (o *ValidationResponse) SetRecalculatedBalanceMicros(v int32)`

SetRecalculatedBalanceMicros sets RecalculatedBalanceMicros field to given value.


### GetStoredFingerprint

`func (o *ValidationResponse) GetStoredFingerprint() string`

GetStoredFingerprint returns the StoredFingerprint field if non-nil, zero value otherwise.

### GetStoredFingerprintOk

`func (o *ValidationResponse) GetStoredFingerprintOk() (*string, bool)`

GetStoredFingerprintOk returns a tuple with the StoredFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredFingerprint

`func (o *ValidationResponse) SetStoredFingerprint(v string)`

SetStoredFingerprint sets StoredFingerprint field to given value.


### GetRecalculatedFingerprint

`func (o *ValidationResponse) GetRecalculatedFingerprint() string`

GetRecalculatedFingerprint returns the RecalculatedFingerprint field if non-nil, zero value otherwise.

### GetRecalculatedFingerprintOk

`func (o *ValidationResponse) GetRecalculatedFingerprintOk() (*string, bool)`

GetRecalculatedFingerprintOk returns a tuple with the RecalculatedFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculatedFingerprint

`func (o *ValidationResponse) SetRecalculatedFingerprint(v string)`

SetRecalculatedFingerprint sets RecalculatedFingerprint field to given value.


### GetAsOfEffectiveAt

`func (o *ValidationResponse) GetAsOfEffectiveAt() string`

GetAsOfEffectiveAt returns the AsOfEffectiveAt field if non-nil, zero value otherwise.

### GetAsOfEffectiveAtOk

`func (o *ValidationResponse) GetAsOfEffectiveAtOk() (*string, bool)`

GetAsOfEffectiveAtOk returns a tuple with the AsOfEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfEffectiveAt

`func (o *ValidationResponse) SetAsOfEffectiveAt(v string)`

SetAsOfEffectiveAt sets AsOfEffectiveAt field to given value.


### GetAsOfRecordedAt

`func (o *ValidationResponse) GetAsOfRecordedAt() string`

GetAsOfRecordedAt returns the AsOfRecordedAt field if non-nil, zero value otherwise.

### GetAsOfRecordedAtOk

`func (o *ValidationResponse) GetAsOfRecordedAtOk() (*string, bool)`

GetAsOfRecordedAtOk returns a tuple with the AsOfRecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfRecordedAt

`func (o *ValidationResponse) SetAsOfRecordedAt(v string)`

SetAsOfRecordedAt sets AsOfRecordedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


