# ReconstructionRunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** |  | 
**RootWalletId** | **string** |  | 
**WalletCount** | **int32** |  | 
**AffectedEventsCount** | **int32** |  | 
**CorrectionEntriesCreated** | **[]string** |  | 

## Methods

### NewReconstructionRunResponse

`func NewReconstructionRunResponse(runId string, rootWalletId string, walletCount int32, affectedEventsCount int32, correctionEntriesCreated []string, ) *ReconstructionRunResponse`

NewReconstructionRunResponse instantiates a new ReconstructionRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconstructionRunResponseWithDefaults

`func NewReconstructionRunResponseWithDefaults() *ReconstructionRunResponse`

NewReconstructionRunResponseWithDefaults instantiates a new ReconstructionRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *ReconstructionRunResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ReconstructionRunResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ReconstructionRunResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetRootWalletId

`func (o *ReconstructionRunResponse) GetRootWalletId() string`

GetRootWalletId returns the RootWalletId field if non-nil, zero value otherwise.

### GetRootWalletIdOk

`func (o *ReconstructionRunResponse) GetRootWalletIdOk() (*string, bool)`

GetRootWalletIdOk returns a tuple with the RootWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWalletId

`func (o *ReconstructionRunResponse) SetRootWalletId(v string)`

SetRootWalletId sets RootWalletId field to given value.


### GetWalletCount

`func (o *ReconstructionRunResponse) GetWalletCount() int32`

GetWalletCount returns the WalletCount field if non-nil, zero value otherwise.

### GetWalletCountOk

`func (o *ReconstructionRunResponse) GetWalletCountOk() (*int32, bool)`

GetWalletCountOk returns a tuple with the WalletCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletCount

`func (o *ReconstructionRunResponse) SetWalletCount(v int32)`

SetWalletCount sets WalletCount field to given value.


### GetAffectedEventsCount

`func (o *ReconstructionRunResponse) GetAffectedEventsCount() int32`

GetAffectedEventsCount returns the AffectedEventsCount field if non-nil, zero value otherwise.

### GetAffectedEventsCountOk

`func (o *ReconstructionRunResponse) GetAffectedEventsCountOk() (*int32, bool)`

GetAffectedEventsCountOk returns a tuple with the AffectedEventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEventsCount

`func (o *ReconstructionRunResponse) SetAffectedEventsCount(v int32)`

SetAffectedEventsCount sets AffectedEventsCount field to given value.


### GetCorrectionEntriesCreated

`func (o *ReconstructionRunResponse) GetCorrectionEntriesCreated() []string`

GetCorrectionEntriesCreated returns the CorrectionEntriesCreated field if non-nil, zero value otherwise.

### GetCorrectionEntriesCreatedOk

`func (o *ReconstructionRunResponse) GetCorrectionEntriesCreatedOk() (*[]string, bool)`

GetCorrectionEntriesCreatedOk returns a tuple with the CorrectionEntriesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectionEntriesCreated

`func (o *ReconstructionRunResponse) SetCorrectionEntriesCreated(v []string)`

SetCorrectionEntriesCreated sets CorrectionEntriesCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


