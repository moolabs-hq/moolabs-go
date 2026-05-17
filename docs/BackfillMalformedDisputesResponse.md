# BackfillMalformedDisputesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scanned** | **int32** |  | 
**Candidates** | [**[]BackfillCandidate**](BackfillCandidate.md) |  | 
**SoftMarked** | **int32** |  | 
**Rerouted** | **int32** |  | 

## Methods

### NewBackfillMalformedDisputesResponse

`func NewBackfillMalformedDisputesResponse(scanned int32, candidates []BackfillCandidate, softMarked int32, rerouted int32, ) *BackfillMalformedDisputesResponse`

NewBackfillMalformedDisputesResponse instantiates a new BackfillMalformedDisputesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackfillMalformedDisputesResponseWithDefaults

`func NewBackfillMalformedDisputesResponseWithDefaults() *BackfillMalformedDisputesResponse`

NewBackfillMalformedDisputesResponseWithDefaults instantiates a new BackfillMalformedDisputesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanned

`func (o *BackfillMalformedDisputesResponse) GetScanned() int32`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *BackfillMalformedDisputesResponse) GetScannedOk() (*int32, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *BackfillMalformedDisputesResponse) SetScanned(v int32)`

SetScanned sets Scanned field to given value.


### GetCandidates

`func (o *BackfillMalformedDisputesResponse) GetCandidates() []BackfillCandidate`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *BackfillMalformedDisputesResponse) GetCandidatesOk() (*[]BackfillCandidate, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *BackfillMalformedDisputesResponse) SetCandidates(v []BackfillCandidate)`

SetCandidates sets Candidates field to given value.


### GetSoftMarked

`func (o *BackfillMalformedDisputesResponse) GetSoftMarked() int32`

GetSoftMarked returns the SoftMarked field if non-nil, zero value otherwise.

### GetSoftMarkedOk

`func (o *BackfillMalformedDisputesResponse) GetSoftMarkedOk() (*int32, bool)`

GetSoftMarkedOk returns a tuple with the SoftMarked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftMarked

`func (o *BackfillMalformedDisputesResponse) SetSoftMarked(v int32)`

SetSoftMarked sets SoftMarked field to given value.


### GetRerouted

`func (o *BackfillMalformedDisputesResponse) GetRerouted() int32`

GetRerouted returns the Rerouted field if non-nil, zero value otherwise.

### GetReroutedOk

`func (o *BackfillMalformedDisputesResponse) GetReroutedOk() (*int32, bool)`

GetReroutedOk returns a tuple with the Rerouted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerouted

`func (o *BackfillMalformedDisputesResponse) SetRerouted(v int32)`

SetRerouted sets Rerouted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


